package faker

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFakeProjectName(t *testing.T) {
	name := Project().Name()
	namePrefix := strings.Split(name, "-")[0]
	assert.Contains(t, projectNames, namePrefix, namePrefix)
}

func TestFakeProjectOrganizationID(t *testing.T) {
	orgID := Project().OrganizationID()
	assert.True(t, strings.HasPrefix(orgID, "org-"), "Organization ID should start with 'org-'")
	assert.Len(t, orgID, 12, "Organization ID should be 12 characters long (org- + 8 chars)")
}

func TestFakeProjectRegion(t *testing.T) {
	region := Project().Region()
	assert.Contains(t, regionNames, region, region)
}

func TestFakeProjectInstanceSize(t *testing.T) {
	instanceSize := Project().InstanceSize()
	assert.Contains(t, instanceSizes, instanceSize, instanceSize)
}

func TestFakeProjectDatabasePassword(t *testing.T) {
	password := Project().DatabasePassword()
	assert.Len(t, password, 16, "Password should be 16 characters long")

	// Check that password contains at least some special characters
	hasSpecialChar := false
	specialChars := "!@#$%^&*"
	for _, char := range password {
		if strings.ContainsRune(specialChars, char) {
			hasSpecialChar = true
			break
		}
	}
	assert.True(t, hasSpecialChar, "Password should contain at least one special character")

	// Check that password contains at least one uppercase letter
	hasUppercase := false
	for _, char := range password {
		if char >= 'A' && char <= 'Z' {
			hasUppercase = true
			break
		}
	}
	assert.True(t, hasUppercase, "Password should contain at least one uppercase letter")

	// Check that password contains at least one lowercase letter
	hasLowercase := false
	for _, char := range password {
		if char >= 'a' && char <= 'z' {
			hasLowercase = true
			break
		}
	}
	assert.True(t, hasLowercase, "Password should contain at least one lowercase letter")

	// Check that password contains at least one digit
	hasDigit := false
	for _, char := range password {
		if char >= '0' && char <= '9' {
			hasDigit = true
			break
		}
	}
	assert.True(t, hasDigit, "Password should contain at least one digit")
}
