package faker

import (
	"reflect"

	fakerTag "github.com/bxcodec/faker/v3"
)

func Generator() {
	_ = fakerTag.AddProvider("ProjectNameFaker", func(v reflect.Value) (any, error) {
		return Project().Name(), nil
	})
	_ = fakerTag.AddProvider("ProjectOrganizationIDFaker", func(v reflect.Value) (any, error) {
		return Project().OrganizationID(), nil
	})
	_ = fakerTag.AddProvider("ProjectRegionFaker", func(v reflect.Value) (any, error) {
		return Project().Region(), nil
	})
	_ = fakerTag.AddProvider("ProjectInstanceSizeFaker", func(v reflect.Value) (any, error) {
		return Project().InstanceSize(), nil
	})
	_ = fakerTag.AddProvider("ProjectDatabasePasswordFaker", func(v reflect.Value) (any, error) {
		return Project().DatabasePassword(), nil
	})

	_ = fakerTag.AddProvider("ApiKeyNameFaker", func(v reflect.Value) (any, error) {
		return ApiKey().Name(), nil
	})

	_ = fakerTag.AddProvider("ApiKeyDescriptionFaker", func(v reflect.Value) (any, error) {
		return ApiKey().Description(), nil
	})
}
