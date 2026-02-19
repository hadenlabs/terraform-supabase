package supabase

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewProject(t *testing.T) {
	t.Parallel()

	project := NewProject()

	assert.Equal(t, "hadenlabs", project.OrganizationID, "OrganizationID should default to 'hadenlabs'")
	// Other fields should be faker-generated
	assert.NotEmpty(t, project.DatabasePassword, "DatabasePassword should be faker-generated")
	assert.NotEmpty(t, project.Name, "Name should be faker-generated")
	assert.NotEmpty(t, project.Region, "Region should be faker-generated")
	assert.NotEmpty(t, project.InstanceSize, "InstanceSize should be faker-generated")
}

func TestProject_WithOrganizationID(t *testing.T) {
	t.Parallel()

	project := NewProject().
		WithOrganizationID("custom-org")

	assert.Equal(t, "custom-org", project.OrganizationID, "OrganizationID should be set to custom value")
}

func TestProject_WithOrganizationID_Chaining(t *testing.T) {
	t.Parallel()

	// Test that chaining works correctly
	project := NewProject().
		WithOrganizationID("first").
		WithOrganizationID("second")

	assert.Equal(t, "second", project.OrganizationID, "Last chained value should be used")
}

func TestProject_ToMap_Basic(t *testing.T) {
	t.Parallel()

	project := NewProject()
	result := project.ToMap()

	expected := map[string]interface{}{
		"organization_id": "hadenlabs",
	}

	assert.Equal(t, expected["organization_id"], result["organization_id"], "ToMap should return correct map structure")
}

func TestProject_ToMap_WithCustomOrganizationID(t *testing.T) {
	t.Parallel()

	project := NewProject().
		WithOrganizationID("my-organization")

	result := project.ToMap()

	assert.Equal(t, "my-organization", result["organization_id"], "ToMap should reflect custom organization ID")
}

func TestProject_DefaultValues(t *testing.T) {
	t.Parallel()

	// Test multiple instances have same defaults
	project1 := NewProject()
	project2 := NewProject()

	assert.Equal(t, project1.OrganizationID, project2.OrganizationID, "Multiple instances should have same defaults")
	assert.Equal(t, "hadenlabs", project1.OrganizationID, "Default OrganizationID should be 'hadenlabs'")
	assert.Equal(t, "hadenlabs", project2.OrganizationID, "Default OrganizationID should be 'hadenlabs'")
}

func TestProject_ImmutableOperations(t *testing.T) {
	t.Parallel()

	// Test that operations return new instances (builder pattern)
	original := NewProject()
	modified := original.WithOrganizationID("modified")

	// Original should remain unchanged
	assert.Equal(t, "hadenlabs", original.OrganizationID, "Original instance should remain unchanged")
	assert.Equal(t, "modified", modified.OrganizationID, "Modified instance should have new value")
}

func TestProject_EmptyStringOrganizationID(t *testing.T) {
	t.Parallel()

	project := NewProject().
		WithOrganizationID("")

	assert.Equal(t, "", project.OrganizationID, "Should allow empty string as organization ID")
}

func TestProject_SpecialCharactersOrganizationID(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		orgID    string
		expected string
	}{
		{"with hyphen", "org-123", "org-123"},
		{"with underscore", "org_456", "org_456"},
		{"with numbers", "org789", "org789"},
		{"mixed case", "OrgName", "OrgName"},
		{"with dot", "org.name", "org.name"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			project := NewProject().WithOrganizationID(tc.orgID)
			assert.Equal(t, tc.expected, project.OrganizationID, "Should handle special characters in organization ID")
		})
	}
}

func TestNewProjectWithFaker(t *testing.T) {
	t.Parallel()

	project := NewProjectWithFaker()

	// All fields should be faker-generated
	assert.NotEmpty(t, project.OrganizationID)
	assert.NotEmpty(t, project.DatabasePassword)
	assert.NotEmpty(t, project.Name)
	assert.NotEmpty(t, project.Region)
	assert.NotEmpty(t, project.InstanceSize)

	// OrganizationID should be faker-generated (not "hadenlabs")
	assert.NotEqual(t, "hadenlabs", project.OrganizationID)
	assert.Contains(t, project.OrganizationID, "org-")
}

func TestProject_WithMethods(t *testing.T) {
	t.Parallel()

	project := NewProject().
		WithOrganizationID("custom-org").
		WithDatabasePassword("CustomPass123!").
		WithName("custom-name").
		WithRegion("eu-central-1").
		WithInstanceSize("large")

	assert.Equal(t, "custom-org", project.OrganizationID)
	assert.Equal(t, "CustomPass123!", project.DatabasePassword)
	assert.Equal(t, "custom-name", project.Name)
	assert.Equal(t, "eu-central-1", project.Region)
	assert.Equal(t, "large", project.InstanceSize)
}

func TestProject_ToMap_Complete(t *testing.T) {
	t.Parallel()

	project := NewProject()
	result := project.ToMap()

	assert.Equal(t, "hadenlabs", result["organization_id"])
	assert.NotEmpty(t, result["database_password"])
	assert.NotEmpty(t, result["name"])
	assert.NotEmpty(t, result["region"])
	assert.NotEmpty(t, result["instance_size"])
	assert.Equal(t, false, result["legacy_api_keys_enabled"])
	assert.Equal(t, true, result["module_enabled"])
}

func TestProject_ToMapWithCustomValues(t *testing.T) {
	t.Parallel()

	project := NewProject()
	result := project.ToMapWithCustomValues(true, false)

	assert.Equal(t, "hadenlabs", result["organization_id"])
	assert.NotEmpty(t, result["database_password"])
	assert.NotEmpty(t, result["name"])
	assert.NotEmpty(t, result["region"])
	assert.NotEmpty(t, result["instance_size"])
	assert.Equal(t, true, result["legacy_api_keys_enabled"])
	assert.Equal(t, false, result["module_enabled"])
}

func TestProject_ImmutableWithMethods(t *testing.T) {
	t.Parallel()

	original := NewProject()
	modified := original.
		WithOrganizationID("new-org").
		WithName("new-name")

	// Original should remain unchanged
	assert.Equal(t, "hadenlabs", original.OrganizationID)
	assert.NotEqual(t, "new-name", original.Name)

	// Modified should have new values
	assert.Equal(t, "new-org", modified.OrganizationID)
	assert.Equal(t, "new-name", modified.Name)

	// Other fields should be copied from original
	assert.Equal(t, original.DatabasePassword, modified.DatabasePassword)
	assert.Equal(t, original.Region, modified.Region)
	assert.Equal(t, original.InstanceSize, modified.InstanceSize)
}
