package trojan

import (
	"context"
	"time"

	"github.com/NamiraNet/xray-core/common"
	"github.com/NamiraNet/xray-core/common/buf"
	"github.com/NamiraNet/xray-core/common/errors"
	"github.com/NamiraNet/xray-core/common/net"
	"github.com/NamiraNet/xray-core/common/protocol"
	"github.com/NamiraNet/xray-core/common/retry"
	"github.com/NamiraNet/xray-core/common/session"
	"github.com/NamiraNet/xray-core/common/signal"
	"github.com/NamiraNet/xray-core/common/task"
	core "github.com/NamiraNet/xray-core/core"
	"github.com/NamiraNet/xray-core/features/policy"
	"github.com/NamiraNet/xray-core/transport"
	"github.com/NamiraNet/xray-core/transport/internet"
	"github.com/NamiraNet/xray-core/transport/internet/stat"
)

// Client is a inbound handler for trojan protocol
type Client struct {
	serverPicker  protocol.ServerPicker
	policyManager policy.Manager
}

// NewClient create a new trojan client.
func NewClient(ctx context.Context, config *ClientConfig) (*Client, error) {
	serverList := protocol.NewServerList()
	for _, rec := range config.Server {
		s, err := protocol.NewServerSpecFromPB(rec)
		if err != nil {
			return nil, errors.New("failed to parse server spec").Base(err)
		}
		serverList.AddServer(s)
	}
	if serverList.Size() == 0 {
		return nil, errors.New("0 server")
	}

	v := core.MustFromContext(ctx)
	client := &Client{
		serverPicker:  protocol.NewRoundRobinServerPicker(serverList),
		policyManager: v.GetFeature(policy.ManagerType()).(policy.Manager),
	}
	return client, nil
}

// Process implements OutboundHandler.Process().
func (c *Client) Process(ctx context.Context, link *transport.Link, dialer internet.Dialer) error {
	outbounds := session.OutboundsFromContext(ctx)
	ob := outbounds[len(outbounds)-1]
	if !ob.Target.IsValid() {
		return errors.New("target not specified")
	}
	ob.Name = "trojan"
	ob.CanSpliceCopy = 3
	destination := ob.Target
	network := destination.Network

	var server *protocol.ServerSpec
	var conn stat.Connection

	err := retry.ExponentialBackoff(5, 100).On(func() error {
		server = c.serverPicker.PickServer()
		rawConn, err := dialer.Dial(ctx, server.Destination())
		if err != nil {
			return err
		}

		conn = rawConn
		return nil
	})
	if err != nil {
		return errors.New("failed to find an available destination").AtWarning().Base(err)
	}
	errors.LogInfo(ctx, "tunneling request to ", destination, " via ", server.Destination().NetAddr())

	defer conn.Close()

	user := server.PickUser()
	account, ok := user.Account.(*MemoryAccount)
	if !ok {
		return errors.New("user account is not valid")
	}

	var newCtx context.Context
	var newCancel context.CancelFunc
	if session.TimeoutOnlyFromContext(ctx) {
		newCtx, newCancel = context.WithCancel(context.Background())
	}

	sessionPolicy := c.policyManager.ForLevel(user.Level)
	ctx, cancel := context.WithCancel(ctx)
	timer := signal.CancelAfterInactivity(ctx, func() {
		cancel()
		if newCancel != nil {
			newCancel()
		}
	}, sessionPolicy.Timeouts.ConnectionIdle)

	postRequest := func() error {
		defer timer.SetTimeout(sessionPolicy.Timeouts.DownlinkOnly)

		bufferWriter := buf.NewBufferedWriter(buf.NewWriter(conn))

		connWriter := &ConnWriter{
			Writer:  bufferWriter,
			Target:  destination,
			Account: account,
		}

		var bodyWriter buf.Writer
		if destination.Network == net.Network_UDP {
			bodyWriter = &PacketWriter{Writer: connWriter, Target: destination}
		} else {
			bodyWriter = connWriter
		}

		// write some request payload to buffer
		if err = buf.CopyOnceTimeout(link.Reader, bodyWriter, time.Millisecond*100); err != nil && err != buf.ErrNotTimeoutReader && err != buf.ErrReadTimeout {
			return errors.New("failed to write A request payload").Base(err).AtWarning()
		}

		// Flush; bufferWriter.WriteMultiBuffer now is bufferWriter.writer.WriteMultiBuffer
		if err = bufferWriter.SetBuffered(false); err != nil {
			return errors.New("failed to flush payload").Base(err).AtWarning()
		}

		// Send header if not sent yet
		if _, err = connWriter.Write([]byte{}); err != nil {
			return err.(*errors.Error).AtWarning()
		}

		if err = buf.Copy(link.Reader, bodyWriter, buf.UpdateActivity(timer)); err != nil {
			return errors.New("failed to transfer request payload").Base(err).AtInfo()
		}

		return nil
	}

	getResponse := func() error {
		defer timer.SetTimeout(sessionPolicy.Timeouts.UplinkOnly)

		var reader buf.Reader
		if network == net.Network_UDP {
			reader = &PacketReader{
				Reader: conn,
			}
		} else {
			reader = buf.NewReader(conn)
		}
		return buf.Copy(reader, link.Writer, buf.UpdateActivity(timer))
	}

	if newCtx != nil {
		ctx = newCtx
	}

	responseDoneAndCloseWriter := task.OnSuccess(getResponse, task.Close(link.Writer))
	if err := task.Run(ctx, postRequest, responseDoneAndCloseWriter); err != nil {
		return errors.New("connection ends").Base(err)
	}

	return nil
}

func init() {
	common.Must(common.RegisterConfig((*ClientConfig)(nil), func(ctx context.Context, config interface{}) (interface{}, error) {
		return NewClient(ctx, config.(*ClientConfig))
	}))
}
