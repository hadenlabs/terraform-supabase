package supabase

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

const (
	moduleProjectDir = "modules/project"
)

func TestDefault(t *testing.T) {
	t.Parallel()

	project := Default()
	assert.NotNil(t, project)
	assert.Equal(t, "hadenlabs", project.OrganizationID)
	assert.NotEmpty(t, project.DatabasePassword)
	assert.NotEmpty(t, project.Name)
	assert.NotEmpty(t, project.Region)
	assert.NotEmpty(t, project.InstanceSize)
}

func TestDefaultWithOrganizationID(t *testing.T) {
	t.Parallel()

	project := DefaultWithOrganizationID("custom-org")
	assert.NotNil(t, project)
	assert.Equal(t, "custom-org", project.OrganizationID)
	assert.NotEmpty(t, project.DatabasePassword)
	assert.NotEmpty(t, project.Name)
	assert.NotEmpty(t, project.Region)
	assert.NotEmpty(t, project.InstanceSize)
}

func TestDefaultForModule(t *testing.T) {
	t.Parallel()

	options := DefaultForModule(moduleProjectDir)

	assert.Equal(t, moduleProjectDir, options.TerraformDir)
	assert.True(t, options.Upgrade)
	assert.NotNil(t, options.Vars)
	assert.Equal(t, "hadenlabs", options.Vars["organization_id"])
	assert.NotEmpty(t, options.Vars["database_password"])
	assert.NotEmpty(t, options.Vars["name"])
	assert.NotEmpty(t, options.Vars["region"])
	assert.NotEmpty(t, options.Vars["instance_size"])
	assert.Equal(t, false, options.Vars["legacy_api_keys_enabled"])
	assert.Equal(t, true, options.Vars["module_enabled"])
}

func TestDefaultForModuleWithOrganizationID(t *testing.T) {
	t.Parallel()

	orgID := "test-organization"
	options := DefaultForModuleWithOrganizationID(moduleProjectDir, orgID)

	assert.Equal(t, moduleProjectDir, options.TerraformDir)
	assert.True(t, options.Upgrade)
	assert.NotNil(t, options.Vars)
	assert.Equal(t, orgID, options.Vars["organization_id"])
	assert.NotEmpty(t, options.Vars["database_password"])
	assert.NotEmpty(t, options.Vars["name"])
	assert.NotEmpty(t, options.Vars["region"])
	assert.NotEmpty(t, options.Vars["instance_size"])
}

func TestMergeProjectValues(t *testing.T) {
	t.Parallel()

	customValues := map[string]any{
		"organization_id": "custom-org",
		"extra_field":     "extra-value",
	}

	result := MergeProjectValues(customValues)

	assert.Equal(t, "custom-org", result["organization_id"])
	assert.Equal(t, "extra-value", result["extra_field"])
	assert.NotEmpty(t, result["database_password"])
	assert.NotEmpty(t, result["name"])
	assert.NotEmpty(t, result["region"])
	assert.NotEmpty(t, result["instance_size"])
}

func TestMergeProjectValues_EmptyCustom(t *testing.T) {
	t.Parallel()

	result := MergeProjectValues(map[string]any{})
	assert.Equal(t, "hadenlabs", result["organization_id"])
	assert.NotEmpty(t, result["database_password"])
	assert.NotEmpty(t, result["name"])
	assert.NotEmpty(t, result["region"])
	assert.NotEmpty(t, result["instance_size"])
}

func TestMergeProjectValues_NilCustom(t *testing.T) {
	t.Parallel()

	result := MergeProjectValues(nil)
	assert.Equal(t, "hadenlabs", result["organization_id"])
	assert.NotEmpty(t, result["database_password"])
	assert.NotEmpty(t, result["name"])
	assert.NotEmpty(t, result["region"])
	assert.NotEmpty(t, result["instance_size"])
}

func TestMergeProjectValuesWithOrganizationID(t *testing.T) {
	t.Parallel()

	orgID := "specific-org"
	customValues := map[string]any{
		"extra_field": "extra-value",
	}

	result := MergeProjectValuesWithOrganizationID(orgID, customValues)

	assert.Equal(t, orgID, result["organization_id"])
	assert.Equal(t, "extra-value", result["extra_field"])
	assert.NotEmpty(t, result["database_password"])
	assert.NotEmpty(t, result["name"])
	assert.NotEmpty(t, result["region"])
	assert.NotEmpty(t, result["instance_size"])
}

func TestTerraformOptions(t *testing.T) {
	t.Parallel()

	customValues := map[string]any{
		"organization_id": "test-org",
		"name":            "test-project",
	}

	options := TerraformOptions(moduleProjectDir, customValues)

	assert.IsType(t, &terraform.Options{}, options)
	assert.Equal(t, moduleProjectDir, options.TerraformDir)
	assert.True(t, options.Upgrade)
	assert.Equal(t, "test-org", options.Vars["organization_id"])
	assert.Equal(t, "test-project", options.Vars["name"])
	assert.NotEmpty(t, options.Vars["database_password"])
	assert.NotEmpty(t, options.Vars["region"])
	assert.NotEmpty(t, options.Vars["instance_size"])
}

func TestTerraformOptionsWithOrganizationID(t *testing.T) {
	t.Parallel()

	orgID := "custom-org-id"
	customValues := map[string]any{
		"name": "test-project",
	}

	options := TerraformOptionsWithOrganizationID(moduleProjectDir, orgID, customValues)

	assert.IsType(t, &terraform.Options{}, options)
	assert.Equal(t, moduleProjectDir, options.TerraformDir)
	assert.True(t, options.Upgrade)
	assert.Equal(t, orgID, options.Vars["organization_id"])
	assert.Equal(t, "test-project", options.Vars["name"])
	assert.NotEmpty(t, options.Vars["database_password"])
	assert.NotEmpty(t, options.Vars["region"])
	assert.NotEmpty(t, options.Vars["instance_size"])
}

func TestGetOrganizationID(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		vars     map[string]any
		expected string
	}{
		{
			name: "with organization_id",
			vars: map[string]any{
				"organization_id": "my-org",
				"name":            "test",
			},
			expected: "my-org",
		},
		{
			name:     "without organization_id",
			vars:     map[string]any{"name": "test"},
			expected: "hadenlabs",
		},
		{
			name:     "empty map",
			vars:     map[string]any{},
			expected: "hadenlabs",
		},
		{
			name:     "nil map",
			vars:     nil,
			expected: "hadenlabs",
		},
		{
			name: "non-string organization_id",
			vars: map[string]any{
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

	originalVars := map[string]any{
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

	originalVars := map[string]any{
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
	customValues := map[string]any{
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
	assert.NotEmpty(t, options4.Vars["database_password"])
	assert.NotEmpty(t, options4.Vars["instance_size"])
}

func TestDefaultWithFaker(t *testing.T) {
	t.Parallel()

	project := DefaultWithFaker()
	assert.NotNil(t, project)

	// All fields should be faker-generated
	assert.NotEmpty(t, project.OrganizationID)
	assert.NotEmpty(t, project.DatabasePassword)
	assert.NotEmpty(t, project.Name)
	assert.NotEmpty(t, project.Region)
	assert.NotEmpty(t, project.InstanceSize)

	// OrganizationID should not be default
	assert.NotEqual(t, "hadenlabs", project.OrganizationID)
}

func TestDefaultForModuleWithFaker(t *testing.T) {
	t.Parallel()

	options := DefaultForModuleWithFaker(moduleProjectDir)

	assert.Equal(t, moduleProjectDir, options.TerraformDir)
	assert.True(t, options.Upgrade)
	assert.NotNil(t, options.Vars)

	// All fields should be faker-generated
	assert.NotEmpty(t, options.Vars["organization_id"])
	assert.NotEmpty(t, options.Vars["database_password"])
	assert.NotEmpty(t, options.Vars["name"])
	assert.NotEmpty(t, options.Vars["region"])
	assert.NotEmpty(t, options.Vars["instance_size"])
	assert.Equal(t, false, options.Vars["legacy_api_keys_enabled"])
	assert.Equal(t, true, options.Vars["module_enabled"])
}

func TestMergeDefaultsWithFaker(t *testing.T) {
	t.Parallel()

	customValues := map[string]any{
		"name":           "custom-name-override",
		"module_enabled": false,
	}

	result := MergeDefaultsWithFaker(customValues)

	// Custom values should override faker defaults
	assert.Equal(t, "custom-name-override", result["name"])
	assert.Equal(t, false, result["module_enabled"])

	// Other fields should be faker-generated
	assert.NotEmpty(t, result["organization_id"])
	assert.NotEmpty(t, result["database_password"])
	assert.NotEmpty(t, result["region"])
	assert.NotEmpty(t, result["instance_size"])
	assert.Equal(t, false, result["legacy_api_keys_enabled"])
}

func TestTerraformOptionsWithFaker(t *testing.T) {
	t.Parallel()

	customValues := map[string]any{
		"name":   "faker-custom-project",
		"region": "ap-southeast-1",
	}

	options := TerraformOptionsWithFaker(moduleProjectDir, customValues)

	assert.IsType(t, &terraform.Options{}, options)
	assert.Equal(t, moduleProjectDir, options.TerraformDir)
	assert.True(t, options.Upgrade)

	// Custom values should be used
	assert.Equal(t, "faker-custom-project", options.Vars["name"])
	assert.Equal(t, "ap-southeast-1", options.Vars["region"])

	// Other fields should be faker-generated
	assert.NotEmpty(t, options.Vars["organization_id"])
	assert.NotEmpty(t, options.Vars["database_password"])
	assert.NotEmpty(t, options.Vars["instance_size"])
	assert.Equal(t, false, options.Vars["legacy_api_keys_enabled"])
	assert.Equal(t, true, options.Vars["module_enabled"])
}

func TestIntegrationWithFaker(t *testing.T) {
	t.Parallel()

	// Test the complete flow with faker
	moduleDir := "modules/project/test/project-basic"

	// Option 1: All faker-generated
	options1 := DefaultForModuleWithFaker(moduleDir)
	assert.NotEmpty(t, options1.Vars["organization_id"])
	assert.NotEmpty(t, options1.Vars["name"])

	// Option 2: Faker with custom overrides
	customValues := map[string]any{
		"name":                    "integration-test",
		"legacy_api_keys_enabled": true,
	}
	options2 := TerraformOptionsWithFaker(moduleDir, customValues)
	assert.Equal(t, "integration-test", options2.Vars["name"])
	assert.Equal(t, true, options2.Vars["legacy_api_keys_enabled"])
	assert.NotEmpty(t, options2.Vars["organization_id"]) // Still faker-generated
}
