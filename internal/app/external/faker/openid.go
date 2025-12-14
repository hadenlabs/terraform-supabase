package faker

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"strings"

	"github.com/lithammer/shortuuid/v3"

	"github.com/hadenlabs/terraform-supabase/internal/errors"
)

type FakeOpenID interface {
	ClientID() string // Name ClientID()
	Name() string     // Name OpenID
}

type fakeOpenID struct{}

func OpenID() FakeOpenID {
	return fakeOpenID{}
}

var (
	openIDNames = []string{"api", "mobile", "front"}
)

func (n fakeOpenID) ClientID() string {
	num, err := rand.Int(rand.Reader, big.NewInt(int64(len(openIDNames))))
	if err != nil {
		panic(errors.New(errors.ErrorUnknown, err.Error()))
	}
	nameuuid := fmt.Sprintf("%s-%s", names[num.Int64()], shortuuid.New())
	return strings.ToLower(nameuuid)
}

func (n fakeOpenID) Name() string {
	num, err := rand.Int(rand.Reader, big.NewInt(int64(len(openIDNames))))
	if err != nil {
		panic(errors.New(errors.ErrorUnknown, err.Error()))
	}
	nameuuid := fmt.Sprintf("%s-%s", names[num.Int64()], shortuuid.New())
	return strings.ToLower(nameuuid)
}
