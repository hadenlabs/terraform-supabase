package faker

import (
	"reflect"

	fakerTag "github.com/bxcodec/faker/v3"
)

func Generator() {
	_ = fakerTag.AddProvider("ProjectNameFaker", func(v reflect.Value) (interface{}, error) {
		return Project().Name(), nil
	})
	_ = fakerTag.AddProvider("ProjectOrganizationIDFaker", func(v reflect.Value) (interface{}, error) {
		return Project().OrganizationID(), nil
	})
	_ = fakerTag.AddProvider("ProjectRegionFaker", func(v reflect.Value) (interface{}, error) {
		return Project().Region(), nil
	})
	_ = fakerTag.AddProvider("ProjectInstanceSizeFaker", func(v reflect.Value) (interface{}, error) {
		return Project().InstanceSize(), nil
	})
	_ = fakerTag.AddProvider("ProjectDatabasePasswordFaker", func(v reflect.Value) (interface{}, error) {
		return Project().DatabasePassword(), nil
	})
	_ = fakerTag.AddProvider("UserUserNameFaker", func(v reflect.Value) (interface{}, error) {
		return User().UserName(), nil
	})
	_ = fakerTag.AddProvider("UserFullNameFaker", func(v reflect.Value) (interface{}, error) {
		return User().FullName(), nil
	})
	_ = fakerTag.AddProvider("UserPasswordFaker", func(v reflect.Value) (interface{}, error) {
		return User().Password(), nil
	})
	_ = fakerTag.AddProvider("UserEmailFaker", func(v reflect.Value) (interface{}, error) {
		return User().Email(), nil
	})
	_ = fakerTag.AddProvider("OpenIDName", func(v reflect.Value) (interface{}, error) {
		return OpenID().Name(), nil
	})
	_ = fakerTag.AddProvider("OpenIDClientID", func(v reflect.Value) (interface{}, error) {
		return OpenID().ClientID(), nil
	})
	_ = fakerTag.AddProvider("RoleName", func(v reflect.Value) (interface{}, error) {
		return Role().Name(), nil
	})
}
