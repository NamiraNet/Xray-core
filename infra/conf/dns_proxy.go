package conf

import (
	"github.com/NamiraNet/xray-core/common/errors"
	"github.com/NamiraNet/xray-core/common/net"
	"github.com/NamiraNet/xray-core/proxy/dns"
	"google.golang.org/protobuf/proto"
)

type DNSOutboundConfig struct {
	Network    Network  `json:"network"`
	Address    *Address `json:"address"`
	Port       uint16   `json:"port"`
	UserLevel  uint32   `json:"userLevel"`
	NonIPQuery string   `json:"nonIPQuery"`
	BlockTypes []int32  `json:"blockTypes"`
}

func (c *DNSOutboundConfig) Build() (proto.Message, error) {
	config := &dns.Config{
		Server: &net.Endpoint{
			Network: c.Network.Build(),
			Port:    uint32(c.Port),
		},
		UserLevel: c.UserLevel,
	}
	if c.Address != nil {
		config.Server.Address = c.Address.Build()
	}
	switch c.NonIPQuery {
	case "":
		c.NonIPQuery = "drop"
	case "drop", "skip", "reject":
	default:
		return nil, errors.New(`unknown "nonIPQuery": `, c.NonIPQuery)
	}
	config.Non_IPQuery = c.NonIPQuery
	config.BlockTypes = c.BlockTypes
	return config, nil
}
