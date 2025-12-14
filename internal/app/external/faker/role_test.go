package faker

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFakeRoleName(t *testing.T) {
	name := Role().Name()
	namePrefix := strings.Split(name, "-")[0]
	assert.Contains(t, names, namePrefix, namePrefix)
}
