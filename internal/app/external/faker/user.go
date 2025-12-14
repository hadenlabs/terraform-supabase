package faker

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"strings"

	"github.com/bxcodec/faker/v3"
	fakerTag "github.com/bxcodec/faker/v3"
	"github.com/lithammer/shortuuid/v3"

	"github.com/hadenlabs/terraform-supabase/internal/errors"
)

type FakeUser interface {
	UserName() string // username
	Email() string    // email
	FullName() string // fullName
	LastName() string // LastName
	Password() string // password
}

type fakeUser struct{}

func User() FakeUser {
	return fakeUser{}
}

var (
	userNames = []string{"optimusprime", "wheeljack", "bumblebee"}
)

func (n fakeUser) UserName() string {
	num, err := rand.Int(rand.Reader, big.NewInt(int64(len(userNames))))
	if err != nil {
		panic(errors.New(errors.ErrorUnknown, err.Error()))
	}
	nameuuid := fmt.Sprintf("%s-%s", userNames[num.Int64()], shortuuid.New())
	return strings.ToLower(nameuuid)
}

func (n fakeUser) Email() string {
	return fakerTag.Email()
}

func (n fakeUser) Password() string {
	return shortuuid.New()
}

func (n fakeUser) FullName() string {
	return faker.Name()
}

func (n fakeUser) LastName() string {
	return faker.LastName()
}
