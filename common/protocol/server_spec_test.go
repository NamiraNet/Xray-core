package protocol_test

import (
	"strings"
	"testing"
	"time"

	"github.com/NamiraNet/xray-core/common"
	"github.com/NamiraNet/xray-core/common/net"
	. "github.com/NamiraNet/xray-core/common/protocol"
	"github.com/NamiraNet/xray-core/common/uuid"
	"github.com/NamiraNet/xray-core/proxy/vmess"
)

func TestAlwaysValidStrategy(t *testing.T) {
	strategy := AlwaysValid()
	if !strategy.IsValid() {
		t.Error("strategy not valid")
	}
	strategy.Invalidate()
	if !strategy.IsValid() {
		t.Error("strategy not valid")
	}
}

func TestTimeoutValidStrategy(t *testing.T) {
	strategy := BeforeTime(time.Now().Add(2 * time.Second))
	if !strategy.IsValid() {
		t.Error("strategy not valid")
	}
	time.Sleep(3 * time.Second)
	if strategy.IsValid() {
		t.Error("strategy is valid")
	}

	strategy = BeforeTime(time.Now().Add(2 * time.Second))
	strategy.Invalidate()
	if strategy.IsValid() {
		t.Error("strategy is valid")
	}
}

func TestUserInServerSpec(t *testing.T) {
	uuid1 := uuid.New()
	uuid2 := uuid.New()

	toAccount := func(a *vmess.Account) Account {
		account, err := a.AsAccount()
		common.Must(err)
		return account
	}

	spec := NewServerSpec(net.Destination{}, AlwaysValid(), &MemoryUser{
		Email:   "test1@example.com",
		Account: toAccount(&vmess.Account{Id: uuid1.String()}),
	})
	if spec.HasUser(&MemoryUser{
		Email:   "test1@example.com",
		Account: toAccount(&vmess.Account{Id: uuid2.String()}),
	}) {
		t.Error("has user: ", uuid2)
	}

	spec.AddUser(&MemoryUser{Email: "test2@example.com"})
	if !spec.HasUser(&MemoryUser{
		Email:   "test1@example.com",
		Account: toAccount(&vmess.Account{Id: uuid1.String()}),
	}) {
		t.Error("not having user: ", uuid1)
	}
}

func TestPickUser(t *testing.T) {
	spec := NewServerSpec(net.Destination{}, AlwaysValid(), &MemoryUser{Email: "test1@example.com"}, &MemoryUser{Email: "test2@example.com"}, &MemoryUser{Email: "test3@example.com"})
	user := spec.PickUser()
	if !strings.HasSuffix(user.Email, "@example.com") {
		t.Error("user: ", user.Email)
	}
}
