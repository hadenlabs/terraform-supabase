package faker

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"strings"

	"github.com/lithammer/shortuuid/v3"

	"github.com/hadenlabs/terraform-supabase/internal/errors"
)

// projectNames is a list of possible project name prefixes
var projectNames = []string{
	"backend",
	"frontend",
	"api",
	"mobile",
	"web",
	"dashboard",
	"admin",
	"app",
	"service",
	"platform",
	"portal",
	"cms",
	"ecommerce",
	"blog",
	"forum",
	"docs",
	"analytics",
	"monitoring",
	"auth",
	"database",
}

// regionNames is a list of possible Supabase regions
var regionNames = []string{
	"us-east-1",
	"us-west-1",
	"eu-west-1",
	"ap-southeast-1",
	"eu-central-1",
	"ap-northeast-1",
}

// instanceSizes is a list of possible instance sizes
var instanceSizes = []string{
	"micro",
	"small",
	"medium",
	"large",
	"xlarge",
}

// FakeProject interface defines methods for generating fake project data
type FakeProject interface {
	Name() string             // Name generates a fake project name
	OrganizationID() string   // OrganizationID generates a fake organization ID
	Region() string           // Region generates a fake region
	InstanceSize() string     // InstanceSize generates a fake instance size
	DatabasePassword() string // DatabasePassword generates a fake database password
}

type fakeProject struct{}

// Project returns a new FakeProject instance
func Project() FakeProject {
	return fakeProject{}
}

// Name generates a fake project name
func (p fakeProject) Name() string {
	num, err := rand.Int(rand.Reader, big.NewInt(int64(len(projectNames))))
	if err != nil {
		panic(errors.New(errors.ErrorUnknown, err.Error()))
	}
	nameuuid := fmt.Sprintf("%s-%s", projectNames[num.Int64()], shortuuid.New())
	return strings.ToLower(nameuuid)
}

// OrganizationID generates a fake organization ID
func (p fakeProject) OrganizationID() string {
	// Organization IDs in Supabase are typically like "org-xxxxxx"
	return fmt.Sprintf("org-%s", shortuuid.New()[:8])
}

// Region generates a fake region
func (p fakeProject) Region() string {
	num, err := rand.Int(rand.Reader, big.NewInt(int64(len(regionNames))))
	if err != nil {
		panic(errors.New(errors.ErrorUnknown, err.Error()))
	}
	return regionNames[num.Int64()]
}

// InstanceSize generates a fake instance size
func (p fakeProject) InstanceSize() string {
	num, err := rand.Int(rand.Reader, big.NewInt(int64(len(instanceSizes))))
	if err != nil {
		panic(errors.New(errors.ErrorUnknown, err.Error()))
	}
	return instanceSizes[num.Int64()]
}

// DatabasePassword generates a fake database password
func (p fakeProject) DatabasePassword() string {
	// Generate a secure password with mix of characters
	const passwordLength = 16
	const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*"

	password := make([]byte, passwordLength)
	for i := range password {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(chars))))
		if err != nil {
			panic(errors.New(errors.ErrorUnknown, err.Error()))
		}
		password[i] = chars[num.Int64()]
	}

	return string(password)
}
