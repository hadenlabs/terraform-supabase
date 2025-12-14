package faker

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFakeRealmName(t *testing.T) {
	name := Realm().Name()
	namePrefix := strings.Split(name, "-")[0]
	assert.Contains(t, names, namePrefix, namePrefix)
}
