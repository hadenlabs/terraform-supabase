package faker

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFakeApiKeyName(t *testing.T) {
	name := ApiKey().Name()
	assert.NotEmpty(t, name, name)
}

func TestFakeApiKeyDescription(t *testing.T) {
	description := ApiKey().Description()
	assert.NotEmpty(t, description, description)
}
