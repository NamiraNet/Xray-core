package tcp

import (
	"github.com/NamiraNet/xray-core/common"
	"github.com/NamiraNet/xray-core/common/net"
)

// PickPort returns an unused TCP port in the system. The port returned is highly likely to be unused, but not guaranteed.
func PickPort() net.Port {
	listener, err := net.Listen("tcp4", "127.0.0.1:0")
	common.Must(err)
	defer listener.Close()

	addr := listener.Addr().(*net.TCPAddr)
	return net.Port(addr.Port)
}
