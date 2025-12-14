package faker

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"strings"

	"github.com/lithammer/shortuuid/v3"

	"github.com/hadenlabs/terraform-supabase/internal/errors"
)

type FakeRealm interface {
	Name() string // Name server
}

type fakeRealm struct{}

func Realm() FakeRealm {
	return fakeRealm{}
}

var (
	names = []string{"optimusprime", "wheeljack", "bumblebee"}
)

func (n fakeRealm) Name() string {
	num, err := rand.Int(rand.Reader, big.NewInt(int64(len(names))))
	if err != nil {
		panic(errors.New(errors.ErrorUnknown, err.Error()))
	}
	nameuuid := fmt.Sprintf("%s-%s", names[num.Int64()], shortuuid.New())
	return strings.ToLower(nameuuid)
}
