package udp

import (
	"context"

	"github.com/NamiraNet/xray-core/common"
	"github.com/NamiraNet/xray-core/common/net"
	"github.com/NamiraNet/xray-core/transport/internet"
	"github.com/NamiraNet/xray-core/transport/internet/stat"
)

func init() {
	common.Must(internet.RegisterTransportDialer(protocolName,
		func(ctx context.Context, dest net.Destination, streamSettings *internet.MemoryStreamConfig) (stat.Connection, error) {
			var sockopt *internet.SocketConfig
			if streamSettings != nil {
				sockopt = streamSettings.SocketSettings
			}
			conn, err := internet.DialSystem(ctx, dest, sockopt)
			if err != nil {
				return nil, err
			}
			// TODO: handle dialer options
			return stat.Connection(conn), nil
		}))
}
