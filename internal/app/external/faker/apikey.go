package faker

import (
	"strings"

	fakerTag "github.com/bxcodec/faker/v3"
)

// FakeApiKey interface defines methods for generating fake ApiKey data
type FakeApiKey interface {
	Name() string        // Name generates a fake ApiKey name
	Description() string // Description generates a fake Description
}

type fakeApiKey struct{}

// ApiKey returns a new FakeApiKey instance
func ApiKey() FakeApiKey {
	return fakeApiKey{}
}

// Name generates a fake ApiKey name
func (p fakeApiKey) Name() string {
	return strings.ToLower(fakerTag.Name())
}

// OrganizationID generates a fake organization ID
func (p fakeApiKey) Description() string {
	return fakerTag.Sentence()
}
