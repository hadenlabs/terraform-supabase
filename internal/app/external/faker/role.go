package faker

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"strings"

	"github.com/lithammer/shortuuid/v3"

	"github.com/hadenlabs/terraform-supabase/internal/errors"
)

type FakeRole interface {
	Name() string // Name Role
}

type fakeRole struct{}

func Role() FakeRole {
	return fakeRole{}
}

var (
	roleNames = []string{"admin", "sre", "developer"}
)

func (n fakeRole) Name() string {
	num, err := rand.Int(rand.Reader, big.NewInt(int64(len(roleNames))))
	if err != nil {
		panic(errors.New(errors.ErrorUnknown, err.Error()))
	}
	nameuuid := fmt.Sprintf("%s-%s", names[num.Int64()], shortuuid.New())
	return strings.ToLower(nameuuid)
}
