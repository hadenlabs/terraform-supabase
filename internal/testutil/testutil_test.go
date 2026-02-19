package testutil

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestDefault(t *testing.T) {
	t.Parallel()

	project := Default()
	assert.NotNil(t, project)
	assert.Equal(t, "hadenlabs", project.OrganizationID)
}

func TestDefaultWithOrganizationID(t *testing.T) {
	t.Parallel()

	project := DefaultWithOrganizationID("custom-org")
	assert.NotNil(t, project)
	assert.Equal(t, "custom-org", project.OrganizationID)
}

func TestDefaultForModule(t *testing.T) {
	t.Parallel()

	moduleDir := "modules/project"
	options := DefaultForModule(moduleDir)

	assert.Equal(t, moduleDir, options.TerraformDir)
	assert.True(t, options.Upgrade)
	assert.NotNil(t, options.Vars)
	assert.Equal(t, "hadenlabs", options.Vars["organization_id"])
}

func TestDefaultForModuleWithOrganizationID(t *testing.T) {
	t.Parallel()

	moduleDir := "modules/project"
	orgID := "test-organization"
	options := DefaultForModuleWithOrganizationID(moduleDir, orgID)

	assert.Equal(t, moduleDir, options.TerraformDir)
	assert.True(t, options.Upgrade)
	assert.NotNil(t, options.Vars)
	assert.Equal(t, orgID, options.Vars["organization_id"])
}

func TestMergeProjectValues(t *testing.T) {
	t.Parallel()

	customValues := map[string]interface{}{
		"organization_id": "custom-org",
		"extra_field":     "extra-value",
	}

	result := MergeProjectValues(customValues)

	assert.Equal(t, "custom-org", result["organization_id"])
	assert.Equal(t, "extra-value", result["extra_field"])
}

func TestMergeProjectValues_EmptyCustom(t *testing.T) {
	t.Parallel()

	result := MergeProjectValues(map[string]interface{}{})
	assert.Equal(t, "hadenlabs", result["organization_id"])
}

func TestMergeProjectValues_NilCustom(t *testing.T) {
	t.Parallel()

	result := MergeProjectValues(nil)
	assert.Equal(t, "hadenlabs", result["organization_id"])
}

func TestMergeProjectValuesWithOrganizationID(t *testing.T) {
	t.Parallel()

	orgID := "specific-org"
	customValues := map[string]interface{}{
		"extra_field": "extra-value",
	}

	result := MergeProjectValuesWithOrganizationID(orgID, customValues)

	assert.Equal(t, orgID, result["organization_id"])
	assert.Equal(t, "extra-value", result["extra_field"])
}

func TestTerraformOptions(t *testing.T) {
	t.Parallel()

	moduleDir := "modules/project"
	customValues := map[string]interface{}{
		"organization_id": "test-org",
		"name":            "test-project",
	}

	options := TerraformOptions(moduleDir, customValues)

	assert.IsType(t, &terraform.Options{}, options)
	assert.Equal(t, moduleDir, options.TerraformDir)
	assert.True(t, options.Upgrade)
	assert.Equal(t, "test-org", options.Vars["organization_id"])
	assert.Equal(t, "test-project", options.Vars["name"])
}

func TestTerraformOptionsWithOrganizationID(t *testing.T) {
	t.Parallel()

	moduleDir := "modules/project"
	orgID := "custom-org-id"
	customValues := map[string]interface{}{
		"name": "test-project",
	}

	options := TerraformOptionsWithOrganizationID(moduleDir, orgID, customValues)

	assert.IsType(t, &terraform.Options{}, options)
	assert.Equal(t, moduleDir, options.TerraformDir)
	assert.True(t, options.Upgrade)
	assert.Equal(t, orgID, options.Vars["organization_id"])
	assert.Equal(t, "test-project", options.Vars["name"])
}

func TestGetOrganizationID(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		vars     map[string]interface{}
		expected string
	}{
		{
			name: "with organization_id",
			vars: map[string]interface{}{
				"organization_id": "my-org",
				"name":            "test",
			},
			expected: "my-org",
		},
		{
			name:     "without organization_id",
			vars:     map[string]interface{}{"name": "test"},
			expected: "hadenlabs",
		},
		{
			name:     "empty map",
			vars:     map[string]interface{}{},
			expected: "hadenlabs",
		},
		{
			name:     "nil map",
			vars:     nil,
			expected: "hadenlabs",
		},
		{
			name: "non-string organization_id",
			vars: map[string]interface{}{
				"organization_id": 123,
			},
			expected: "hadenlabs",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			result := GetOrganizationID(tc.vars)
			assert.Equal(t, tc.expected, result)
		})
	}
}

func TestSetOrganizationID(t *testing.T) {
	t.Parallel()

	originalVars := map[string]interface{}{
		"name":   "test-project",
		"region": "us-east-1",
	}

	newOrgID := "new-org"
	result := SetOrganizationID(originalVars, newOrgID)

	assert.Equal(t, newOrgID, result["organization_id"])
	assert.Equal(t, "test-project", result["name"])
	assert.Equal(t, "us-east-1", result["region"])
}

func TestSetOrganizationID_OverrideExisting(t *testing.T) {
	t.Parallel()

	originalVars := map[string]interface{}{
		"organization_id": "old-org",
		"name":            "test-project",
	}

	newOrgID := "new-org"
	result := SetOrganizationID(originalVars, newOrgID)

	assert.Equal(t, newOrgID, result["organization_id"])
	assert.Equal(t, "test-project", result["name"])
}

func TestIsDefaultOrganizationID(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		orgID    string
		expected bool
	}{
		{"hadenlabs", true},
		{"Hadenlabs", false},
		{"hadenlabs ", false},
		{"", false},
		{"other-org", false},
	}

	for _, tc := range testCases {
		t.Run(tc.orgID, func(t *testing.T) {
			t.Parallel()
			result := IsDefaultOrganizationID(tc.orgID)
			assert.Equal(t, tc.expected, result)
		})
	}
}

func TestValidateOrganizationID(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		orgID    string
		expected bool
	}{
		{"hadenlabs", true},
		{"my-org", true},
		{"org123", true},
		{"", false},
		{"   ", true}, // Spaces are considered non-empty
	}

	for _, tc := range testCases {
		t.Run(tc.orgID, func(t *testing.T) {
			t.Parallel()
			result := ValidateOrganizationID(tc.orgID)
			assert.Equal(t, tc.expected, result)
		})
	}
}

func TestIntegrationExample(t *testing.T) {
	t.Parallel()

	// Example of how to use the package in a test
	moduleDir := "modules/project/test/project-basic"

	// Option 1: Use default organization ID
	options1 := DefaultForModule(moduleDir)
	assert.Equal(t, "hadenlabs", options1.Vars["organization_id"])

	// Option 2: Use custom organization ID
	options2 := DefaultForModuleWithOrganizationID(moduleDir, "my-company")
	assert.Equal(t, "my-company", options2.Vars["organization_id"])

	// Option 3: Merge with custom values
	customValues := map[string]interface{}{
		"name":   "integration-test",
		"region": "us-west-1",
	}
	options3 := TerraformOptions(moduleDir, customValues)
	assert.Equal(t, "hadenlabs", options3.Vars["organization_id"])
	assert.Equal(t, "integration-test", options3.Vars["name"])
	assert.Equal(t, "us-west-1", options3.Vars["region"])

	// Option 4: Specific organization ID with custom values
	options4 := TerraformOptionsWithOrganizationID(moduleDir, "specific-org", customValues)
	assert.Equal(t, "specific-org", options4.Vars["organization_id"])
	assert.Equal(t, "integration-test", options4.Vars["name"])
	assert.Equal(t, "us-west-1", options4.Vars["region"])
}
