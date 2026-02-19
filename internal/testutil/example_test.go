package testutil

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestExampleBasicUsage demonstrates basic usage of the testutil package
func TestExampleBasicUsage(t *testing.T) {
	t.Parallel()

	// Get default project
	project := Default()
	fmt.Printf("Default Organization ID: %s\n", project.OrganizationID)

	// Customize the organization ID
	customProject := project.WithOrganizationID("my-company")
	fmt.Printf("Custom Organization ID: %s\n", customProject.OrganizationID)

	// Convert to Terraform variables
	tfVars := customProject.ToMap()
	fmt.Printf("Terraform Variables: %v\n", tfVars)

	// Assertions
	assert.Equal(t, "hadenlabs", project.OrganizationID, "Original project should remain unchanged")
	assert.Equal(t, "my-company", customProject.OrganizationID, "New project should have custom value")
	assert.Equal(t, "my-company", tfVars["organization_id"], "Map should reflect custom value")
}

// TestExampleUtilityFunctions demonstrates using utility functions
func TestExampleUtilityFunctions(t *testing.T) {
	t.Parallel()

	// Get default project
	defaultProject := Default()
	fmt.Printf("Default: %s\n", defaultProject.OrganizationID)

	// Get project with custom organization ID
	customProject := DefaultWithOrganizationID("acme-corp")
	fmt.Printf("Custom: %s\n", customProject.OrganizationID)

	// Check if it's the default value
	isDefault := IsDefaultOrganizationID("hadenlabs")
	fmt.Printf("Is 'hadenlabs' default? %v\n", isDefault)

	// Validate an organization ID
	isValid := ValidateOrganizationID("valid-org")
	fmt.Printf("Is 'valid-org' valid? %v\n", isValid)

	// Assertions
	assert.Equal(t, "hadenlabs", defaultProject.OrganizationID)
	assert.Equal(t, "acme-corp", customProject.OrganizationID)
	assert.True(t, isDefault)
	assert.True(t, isValid)
}

// TestExampleTerraformOptions demonstrates creating Terraform options
func TestExampleTerraformOptions(t *testing.T) {
	t.Parallel()

	moduleDir := "modules/project"

	// Option 1: Default organization ID
	options1 := DefaultForModule(moduleDir)
	fmt.Printf("Options 1 - Organization ID: %s\n", options1.Vars["organization_id"])

	// Option 2: Custom organization ID
	options2 := DefaultForModuleWithOrganizationID(moduleDir, "my-org")
	fmt.Printf("Options 2 - Organization ID: %s\n", options2.Vars["organization_id"])

	// Option 3: Merge with custom values
	customValues := map[string]interface{}{
		"name":   "test-project",
		"region": "us-east-1",
	}
	options3 := TerraformOptions(moduleDir, customValues)
	fmt.Printf("Options 3 - Organization ID: %s, Name: %s\n",
		options3.Vars["organization_id"], options3.Vars["name"])

	// Option 4: Specific organization ID with custom values
	options4 := TerraformOptionsWithOrganizationID(moduleDir, "specific-org", customValues)
	fmt.Printf("Options 4 - Organization ID: %s, Name: %s\n",
		options4.Vars["organization_id"], options4.Vars["name"])

	// Assertions
	assert.Equal(t, "hadenlabs", options1.Vars["organization_id"])
	assert.Equal(t, "my-org", options2.Vars["organization_id"])
	assert.Equal(t, "hadenlabs", options3.Vars["organization_id"])
	assert.Equal(t, "test-project", options3.Vars["name"])
	assert.Equal(t, "specific-org", options4.Vars["organization_id"])
	assert.Equal(t, "test-project", options4.Vars["name"])
}

// TestProjectIntegrationExample demonstrates real test using the testutil package
func TestProjectIntegrationExample(t *testing.T) {
	t.Parallel()

	// Create Terraform options with default organization ID
	terraformOptions := DefaultForModule("modules/project/test/project-basic")

	// Customize other variables if needed
	terraformOptions.Vars["name"] = "test-integration-project"
	terraformOptions.Vars["module_enabled"] = true

	// In a real test, you would use:
	// defer terraform.Destroy(t, terraformOptions)
	// terraform.InitAndApply(t, terraformOptions)

	// Verify the options
	assert.Equal(t, "hadenlabs", terraformOptions.Vars["organization_id"])
	assert.Equal(t, "test-integration-project", terraformOptions.Vars["name"])
	assert.Equal(t, true, terraformOptions.Vars["module_enabled"])
}

// TestProjectWithCustomOrganization demonstrates test with custom organization ID
func TestProjectWithCustomOrganization(t *testing.T) {
	t.Parallel()

	// Use a custom organization ID
	terraformOptions := DefaultForModuleWithOrganizationID(
		"modules/project/test/project-basic",
		"my-company-org",
	)

	// Add other required variables
	terraformOptions.Vars["name"] = "company-project"
	terraformOptions.Vars["database_password"] = "SecurePass123!"
	terraformOptions.Vars["region"] = "us-west-1"
	terraformOptions.Vars["module_enabled"] = true

	// Verify
	assert.Equal(t, "my-company-org", terraformOptions.Vars["organization_id"])
	assert.Equal(t, "company-project", terraformOptions.Vars["name"])
	assert.Equal(t, "SecurePass123!", terraformOptions.Vars["database_password"])
	assert.Equal(t, "us-west-1", terraformOptions.Vars["region"])
	assert.Equal(t, true, terraformOptions.Vars["module_enabled"])
}

// TestProjectWithMergedValues demonstrates using merge functions
func TestProjectWithMergedValues(t *testing.T) {
	t.Parallel()

	moduleDir := "modules/project/test/project-basic"

	// Define all custom values
	customValues := map[string]interface{}{
		"organization_id":         "merged-org",
		"name":                    "merged-project",
		"database_password":       "MergedPass456!",
		"region":                  "eu-west-1",
		"instance_size":           "medium",
		"legacy_api_keys_enabled": false,
		"module_enabled":          true,
	}

	// Merge with defaults (though we're overriding all values)
	terraformOptions := TerraformOptions(moduleDir, customValues)

	// Verify all values
	assert.Equal(t, "merged-org", terraformOptions.Vars["organization_id"])
	assert.Equal(t, "merged-project", terraformOptions.Vars["name"])
	assert.Equal(t, "MergedPass456!", terraformOptions.Vars["database_password"])
	assert.Equal(t, "eu-west-1", terraformOptions.Vars["region"])
	assert.Equal(t, "medium", terraformOptions.Vars["instance_size"])
	assert.Equal(t, false, terraformOptions.Vars["legacy_api_keys_enabled"])
	assert.Equal(t, true, terraformOptions.Vars["module_enabled"])
}

// TestProjectOrganizationSuite demonstrates test suite with different organization configurations
func TestProjectOrganizationSuite(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name        string
		orgID       string
		description string
	}{
		{
			name:        "default-organization",
			orgID:       "hadenlabs",
			description: "Using default organization ID",
		},
		{
			name:        "custom-organization",
			orgID:       "acme-corporation",
			description: "Using custom organization ID",
		},
		{
			name:        "numeric-organization",
			orgID:       "org-12345",
			description: "Using numeric organization ID",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			// Create Terraform options for this test case
			terraformOptions := DefaultForModuleWithOrganizationID(
				"modules/project/test/project-basic",
				tc.orgID,
			)

			// Add common variables
			terraformOptions.Vars["name"] = fmt.Sprintf("test-%s", tc.name)
			terraformOptions.Vars["module_enabled"] = true

			// Verify
			assert.Equal(t, tc.orgID, terraformOptions.Vars["organization_id"])
			assert.Equal(t, fmt.Sprintf("test-%s", tc.name), terraformOptions.Vars["name"])
			assert.Equal(t, true, terraformOptions.Vars["module_enabled"])

			// Validate the organization ID
			assert.True(t, ValidateOrganizationID(tc.orgID), "Organization ID should be valid")
		})
	}
}

// TestHelperFunctions demonstrates helper function usage
func TestHelperFunctions(t *testing.T) {
	t.Parallel()

	// Test GetOrganizationID
	vars := map[string]interface{}{
		"name":            "test-project",
		"organization_id": "test-org",
		"region":          "us-east-1",
	}

	orgID := GetOrganizationID(vars)
	assert.Equal(t, "test-org", orgID)

	// Test SetOrganizationID
	newVars := SetOrganizationID(vars, "new-org")
	assert.Equal(t, "new-org", newVars["organization_id"])
	assert.Equal(t, "test-project", newVars["name"])
	assert.Equal(t, "us-east-1", newVars["region"])

	// Test MergeProjectValues
	merged := MergeProjectValues(map[string]interface{}{
		"organization_id": "merged-org",
		"extra_field":     "extra-value",
	})
	assert.Equal(t, "merged-org", merged["organization_id"])
	assert.Equal(t, "extra-value", merged["extra_field"])
}

// TestEdgeCases demonstrates edge cases and validation
func TestEdgeCases(t *testing.T) {
	t.Parallel()

	// Test empty organization ID
	project := Default().WithOrganizationID("")
	assert.Equal(t, "", project.OrganizationID)
	assert.False(t, ValidateOrganizationID(""), "Empty organization ID should be invalid")

	// Test default value detection
	assert.True(t, IsDefaultOrganizationID("hadenlabs"))
	assert.False(t, IsDefaultOrganizationID("Hadenlabs"))  // Case sensitive
	assert.False(t, IsDefaultOrganizationID("hadenlabs ")) // Trailing space

	// Test with special characters
	specialProject := Default().WithOrganizationID("org-name_with.mixed-chars123")
	assert.Equal(t, "org-name_with.mixed-chars123", specialProject.OrganizationID)
	assert.True(t, ValidateOrganizationID(specialProject.OrganizationID))
}

// TestCompleteIntegrationPattern demonstrates complete integration test pattern
func TestCompleteIntegrationPattern(t *testing.T) {
	t.Parallel()

	// This shows the complete pattern for an integration test
	moduleDir := "modules/project/test/project-basic"

	// Step 1: Create Terraform options with desired organization ID
	terraformOptions := DefaultForModuleWithOrganizationID(moduleDir, "integration-test-org")

	// Step 2: Set all required variables
	terraformOptions.Vars["name"] = "integration-test-project"
	terraformOptions.Vars["database_password"] = "IntegrationPass789!"
	terraformOptions.Vars["region"] = "eu-central-1"
	terraformOptions.Vars["instance_size"] = "small"
	terraformOptions.Vars["legacy_api_keys_enabled"] = false
	terraformOptions.Vars["module_enabled"] = true

	// Step 3: Validate all values are set
	assert.Equal(t, "integration-test-org", terraformOptions.Vars["organization_id"])
	assert.Equal(t, "integration-test-project", terraformOptions.Vars["name"])
	assert.Equal(t, "IntegrationPass789!", terraformOptions.Vars["database_password"])
	assert.Equal(t, "eu-central-1", terraformOptions.Vars["region"])
	assert.Equal(t, "small", terraformOptions.Vars["instance_size"])
	assert.Equal(t, false, terraformOptions.Vars["legacy_api_keys_enabled"])
	assert.Equal(t, true, terraformOptions.Vars["module_enabled"])

	// Step 4: In a real test, you would now run Terraform
	// defer terraform.Destroy(t, terraformOptions)
	// terraform.InitAndApply(t, terraformOptions)

	// Step 5: Verify outputs
	// outputProjectID := terraform.Output(t, terraformOptions, "project_id")
	// assert.NotEmpty(t, outputProjectID)
}
