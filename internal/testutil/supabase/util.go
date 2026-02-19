package supabase

import (
	"github.com/gruntwork-io/terratest/modules/terraform"
)

// Default returns a new Project instance with default values
func Default() *Project {
	return NewProject()
}

// DefaultWithFaker returns a new Project instance with all fields from faker
func DefaultWithFaker() *Project {
	return NewProjectWithFaker()
}

// DefaultWithOrganizationID creates a new Project with a specific organization ID
func DefaultWithOrganizationID(orgID string) *Project {
	return NewProject().WithOrganizationID(orgID)
}

// DefaultForModule creates Terraform options with default values for a specific module
func DefaultForModule(moduleDir string) *terraform.Options {
	project := Default()
	return &terraform.Options{
		TerraformDir: moduleDir,
		Upgrade:      true,
		Vars:         project.ToMap(),
	}
}

// DefaultForModuleWithFaker creates Terraform options with faker-generated values
func DefaultForModuleWithFaker(moduleDir string) *terraform.Options {
	project := DefaultWithFaker()
	return &terraform.Options{
		TerraformDir: moduleDir,
		Upgrade:      true,
		Vars:         project.ToMap(),
	}
}

// DefaultForModuleWithOrganizationID creates Terraform options with custom organization ID
func DefaultForModuleWithOrganizationID(moduleDir, orgID string) *terraform.Options {
	project := DefaultWithOrganizationID(orgID)
	return &terraform.Options{
		TerraformDir: moduleDir,
		Upgrade:      true,
		Vars:         project.ToMap(),
	}
}

// MergeProjectValues merges project values with custom values, with custom values taking precedence
func MergeProjectValues(customValues map[string]interface{}) map[string]interface{} {
	project := Default().ToMap()

	// Create a new map starting with project defaults
	result := make(map[string]interface{})
	for k, v := range project {
		result[k] = v
	}

	// Override with custom values
	for k, v := range customValues {
		result[k] = v
	}

	return result
}

// MergeDefaultsWithFaker merges custom values with faker-generated defaults
func MergeDefaultsWithFaker(customValues map[string]interface{}) map[string]interface{} {
	defaults := DefaultWithFaker().ToMap()

	// Create a new map starting with defaults
	result := make(map[string]interface{})
	for k, v := range defaults {
		result[k] = v
	}

	// Override with custom values
	for k, v := range customValues {
		result[k] = v
	}

	return result
}

// MergeProjectValuesWithOrganizationID merges custom values with specific organization ID
func MergeProjectValuesWithOrganizationID(orgID string, customValues map[string]interface{}) map[string]interface{} {
	project := DefaultWithOrganizationID(orgID).ToMap()

	// Create a new map starting with project defaults
	result := make(map[string]interface{})
	for k, v := range project {
		result[k] = v
	}

	// Override with custom values
	for k, v := range customValues {
		result[k] = v
	}

	return result
}

// TerraformOptions creates Terraform options with merged project values
func TerraformOptions(moduleDir string, customValues map[string]interface{}) *terraform.Options {
	mergedValues := MergeProjectValues(customValues)
	return &terraform.Options{
		TerraformDir: moduleDir,
		Upgrade:      true,
		Vars:         mergedValues,
	}
}

// TerraformOptionsWithFaker creates Terraform options with faker-generated merged values
func TerraformOptionsWithFaker(moduleDir string, customValues map[string]interface{}) *terraform.Options {
	mergedValues := MergeDefaultsWithFaker(customValues)
	return &terraform.Options{
		TerraformDir: moduleDir,
		Upgrade:      true,
		Vars:         mergedValues,
	}
}

// TerraformOptionsWithOrganizationID creates Terraform options with specific organization ID
func TerraformOptionsWithOrganizationID(moduleDir, orgID string, customValues map[string]interface{}) *terraform.Options {
	mergedValues := MergeProjectValuesWithOrganizationID(orgID, customValues)
	return &terraform.Options{
		TerraformDir: moduleDir,
		Upgrade:      true,
		Vars:         mergedValues,
	}
}

// GetOrganizationID returns the organization ID from a map of Terraform variables
func GetOrganizationID(vars map[string]interface{}) string {
	if orgID, ok := vars["organization_id"]; ok {
		if str, ok := orgID.(string); ok {
			return str
		}
	}
	return "hadenlabs" // Default fallback
}

// SetOrganizationID sets the organization ID in a map of Terraform variables
func SetOrganizationID(vars map[string]interface{}, orgID string) map[string]interface{} {
	result := make(map[string]interface{})

	// Copy all existing values
	for k, v := range vars {
		result[k] = v
	}

	// Set organization_id
	result["organization_id"] = orgID

	return result
}

// IsDefaultOrganizationID checks if the organization ID is the default value
func IsDefaultOrganizationID(orgID string) bool {
	return orgID == "hadenlabs"
}

// ValidateOrganizationID validates that an organization ID is not empty
func ValidateOrganizationID(orgID string) bool {
	return orgID != ""
}
