package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"

	"github.com/hadenlabs/terraform-supabase/internal/testutil/supabase"
)

func TestProjectBasicSuccess(t *testing.T) {
	t.Parallel()

	// Generate fake data for the test
	project := supabase.NewProject()

	organizationID := project.OrganizationID
	databasePassword := project.DatabasePassword
	name := project.Name
	region := project.Region

	terraformOptions := &terraform.Options{
		// The path to where your Terraform code is located
		TerraformDir: "project-basic",
		Upgrade:      true,
		Vars: map[string]interface{}{
			"database_password":       databasePassword,
			"name":                    name,
			"organization_id":         organizationID,
			"region":                  region,
			"legacy_api_keys_enabled": false,
			"module_enabled":          true,
		},
	}

	// At the end of the test, run `terraform destroy` to clean up any resources that were created
	defer terraform.Destroy(t, terraformOptions)

	// This will run `terraform init` and `terraform apply` and fail the test if there are any errors
	terraform.InitAndApply(t, terraformOptions)

	// Verify outputs
	outputProjectID := terraform.Output(t, terraformOptions, "project_id")
	outputModuleEnabled := terraform.Output(t, terraformOptions, "module_enabled")

	// Assertions
	assert.NotEmpty(t, outputProjectID, "Project ID should not be empty")
	assert.Equal(t, "true", outputModuleEnabled, "Module should be enabled")
}
