package udp

import (
	"github.com/NamiraNet/xray-core/common/buf"
	"github.com/NamiraNet/xray-core/common/net"
)

// Packet is a UDP packet together with its source and destination address.
type Packet struct {
	Payload *buf.Buffer
	Source  net.Destination
	Target  net.Destination
}
