package vmess

import (
	"strings"

	"github.com/NamiraNet/xray-core/common/errors"
	"github.com/NamiraNet/xray-core/common/protocol"
	"github.com/NamiraNet/xray-core/common/uuid"
	"google.golang.org/protobuf/proto"
)

// MemoryAccount is an in-memory form of VMess account.
type MemoryAccount struct {
	// ID is the main ID of the account.
	ID *protocol.ID
	// Security type of the account. Used for client connections.
	Security protocol.SecurityType

	AuthenticatedLengthExperiment bool
	NoTerminationSignal           bool
}

// Equals implements protocol.Account.
func (a *MemoryAccount) Equals(account protocol.Account) bool {
	vmessAccount, ok := account.(*MemoryAccount)
	if !ok {
		return false
	}
	return a.ID.Equals(vmessAccount.ID)
}

func (a *MemoryAccount) ToProto() proto.Message {
	test := ""
	if a.AuthenticatedLengthExperiment {
		test = "AuthenticatedLength|"
	}
	if a.NoTerminationSignal {
		test = test + "NoTerminationSignal"
	}
	return &Account{
		Id:               a.ID.String(),
		TestsEnabled:     test,
		SecuritySettings: &protocol.SecurityConfig{Type: a.Security},
	}
}

// AsAccount implements protocol.Account.
func (a *Account) AsAccount() (protocol.Account, error) {
	id, err := uuid.ParseString(a.Id)
	if err != nil {
		return nil, errors.New("failed to parse ID").Base(err).AtError()
	}
	protoID := protocol.NewID(id)
	var AuthenticatedLength, NoTerminationSignal bool
	if strings.Contains(a.TestsEnabled, "AuthenticatedLength") {
		AuthenticatedLength = true
	}
	if strings.Contains(a.TestsEnabled, "NoTerminationSignal") {
		NoTerminationSignal = true
	}
	return &MemoryAccount{
		ID:                            protoID,
		Security:                      a.SecuritySettings.GetSecurityType(),
		AuthenticatedLengthExperiment: AuthenticatedLength,
		NoTerminationSignal:           NoTerminationSignal,
	}, nil
}
