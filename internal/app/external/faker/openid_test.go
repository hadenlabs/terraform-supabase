package faker

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFakeOpenIDName(t *testing.T) {
	name := OpenID().Name()
	namePrefix := strings.Split(name, "-")[0]
	assert.Contains(t, names, namePrefix, namePrefix)
}

func TestFakeOpenIDClientID(t *testing.T) {
	clientID := OpenID().ClientID()
	namePrefix := strings.Split(clientID, "-")[0]
	assert.Contains(t, names, namePrefix, namePrefix)
}
