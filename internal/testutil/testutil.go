package testutil

import (
	"github.com/gruntwork-io/terratest/modules/terraform"

	"github.com/hadenlabs/terraform-supabase/internal/testutil/supabase"
)

// Default returns a new Project instance with default values from supabase package
func Default() *supabase.Project {
	return supabase.Default()
}

// DefaultWithOrganizationID creates a new Project with a specific organization ID
func DefaultWithOrganizationID(orgID string) *supabase.Project {
	return supabase.DefaultWithOrganizationID(orgID)
}

// DefaultForModule creates Terraform options with default values for a specific module
func DefaultForModule(moduleDir string) *terraform.Options {
	return supabase.DefaultForModule(moduleDir)
}

// DefaultForModuleWithOrganizationID creates Terraform options with custom organization ID
func DefaultForModuleWithOrganizationID(moduleDir, orgID string) *terraform.Options {
	return supabase.DefaultForModuleWithOrganizationID(moduleDir, orgID)
}

// MergeProjectValues merges project values with custom values, with custom values taking precedence
func MergeProjectValues(customValues map[string]interface{}) map[string]interface{} {
	return supabase.MergeProjectValues(customValues)
}

// MergeProjectValuesWithOrganizationID merges custom values with specific organization ID
func MergeProjectValuesWithOrganizationID(orgID string, customValues map[string]interface{}) map[string]interface{} {
	return supabase.MergeProjectValuesWithOrganizationID(orgID, customValues)
}

// TerraformOptions creates Terraform options with merged project values
func TerraformOptions(moduleDir string, customValues map[string]interface{}) *terraform.Options {
	return supabase.TerraformOptions(moduleDir, customValues)
}

// TerraformOptionsWithOrganizationID creates Terraform options with specific organization ID
func TerraformOptionsWithOrganizationID(moduleDir, orgID string, customValues map[string]interface{}) *terraform.Options {
	return supabase.TerraformOptionsWithOrganizationID(moduleDir, orgID, customValues)
}

// GetOrganizationID returns the organization ID from a map of Terraform variables
func GetOrganizationID(vars map[string]interface{}) string {
	return supabase.GetOrganizationID(vars)
}

// SetOrganizationID sets the organization ID in a map of Terraform variables
func SetOrganizationID(vars map[string]interface{}, orgID string) map[string]interface{} {
	return supabase.SetOrganizationID(vars, orgID)
}

// IsDefaultOrganizationID checks if the organization ID is the default value
func IsDefaultOrganizationID(orgID string) bool {
	return supabase.IsDefaultOrganizationID(orgID)
}

// ValidateOrganizationID validates that an organization ID is not empty
func ValidateOrganizationID(orgID string) bool {
	return supabase.ValidateOrganizationID(orgID)
}
