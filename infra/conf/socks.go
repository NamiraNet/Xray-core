package conf

import (
	"encoding/json"

	"github.com/NamiraNet/xray-core/common/errors"
	"github.com/NamiraNet/xray-core/common/protocol"
	"github.com/NamiraNet/xray-core/common/serial"
	"github.com/NamiraNet/xray-core/proxy/socks"
	"google.golang.org/protobuf/proto"
)

type SocksAccount struct {
	Username string `json:"user"`
	Password string `json:"pass"`
}

func (v *SocksAccount) Build() *socks.Account {
	return &socks.Account{
		Username: v.Username,
		Password: v.Password,
	}
}

const (
	AuthMethodNoAuth   = "noauth"
	AuthMethodUserPass = "password"
)

type SocksServerConfig struct {
	AuthMethod string          `json:"auth"`
	Accounts   []*SocksAccount `json:"accounts"`
	UDP        bool            `json:"udp"`
	Host       *Address        `json:"ip"`
	UserLevel  uint32          `json:"userLevel"`
}

func (v *SocksServerConfig) Build() (proto.Message, error) {
	config := new(socks.ServerConfig)
	switch v.AuthMethod {
	case AuthMethodNoAuth:
		config.AuthType = socks.AuthType_NO_AUTH
	case AuthMethodUserPass:
		config.AuthType = socks.AuthType_PASSWORD
	default:
		// errors.New("unknown socks auth method: ", v.AuthMethod, ". Default to noauth.").AtWarning().WriteToLog()
		config.AuthType = socks.AuthType_NO_AUTH
	}

	if len(v.Accounts) > 0 {
		config.Accounts = make(map[string]string, len(v.Accounts))
		for _, account := range v.Accounts {
			config.Accounts[account.Username] = account.Password
		}
	}

	config.UdpEnabled = v.UDP
	if v.Host != nil {
		config.Address = v.Host.Build()
	}

	config.UserLevel = v.UserLevel
	return config, nil
}

type SocksRemoteConfig struct {
	Address *Address          `json:"address"`
	Port    uint16            `json:"port"`
	Users   []json.RawMessage `json:"users"`
}

type SocksClientConfig struct {
	Servers []*SocksRemoteConfig `json:"servers"`
}

func (v *SocksClientConfig) Build() (proto.Message, error) {
	config := new(socks.ClientConfig)
	config.Server = make([]*protocol.ServerEndpoint, len(v.Servers))
	for idx, serverConfig := range v.Servers {
		server := &protocol.ServerEndpoint{
			Address: serverConfig.Address.Build(),
			Port:    uint32(serverConfig.Port),
		}
		for _, rawUser := range serverConfig.Users {
			user := new(protocol.User)
			if err := json.Unmarshal(rawUser, user); err != nil {
				return nil, errors.New("failed to parse Socks user").Base(err).AtError()
			}
			account := new(SocksAccount)
			if err := json.Unmarshal(rawUser, account); err != nil {
				return nil, errors.New("failed to parse socks account").Base(err).AtError()
			}
			user.Account = serial.ToTypedMessage(account.Build())
			server.User = append(server.User, user)
		}
		config.Server[idx] = server
	}
	return config, nil
}
