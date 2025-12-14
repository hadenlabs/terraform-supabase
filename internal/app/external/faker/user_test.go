package faker

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFakeUserUserName(t *testing.T) {
	userNameFake := User().UserName()
	prefix := strings.Split(userNameFake, "-")[0]
	assert.Contains(t, userNames, prefix, prefix)
}

func TestFakeUserPassword(t *testing.T) {
	password := User().Password()
	assert.NotEmpty(t, password, password)
}

func TestFakeUserEmail(t *testing.T) {
	email := User().Email()
	assert.NotEmpty(t, email, email)
}

func TestFakeUserFullName(t *testing.T) {
	fullName := User().FullName()
	assert.NotEmpty(t, fullName, fullName)
}
