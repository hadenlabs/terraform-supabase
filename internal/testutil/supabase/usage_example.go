package supabase

import (
	"fmt"
)

// ExampleUsage demonstrates how to use the supabase package
func ExampleUsage() {
	fmt.Println("=== Supabase TestUtil Package Usage Examples ===")
	fmt.Println()

	// Example 1: Create a project with default organization ID and faker-generated other fields
	fmt.Println("Example 1: Default Project (OrganizationID='hadenlabs', others=faker)")
	project1 := NewProject()
	fmt.Printf("  OrganizationID: %s\n", project1.OrganizationID)
	fmt.Printf("  Name: %s\n", project1.Name)
	fmt.Printf("  Region: %s\n", project1.Region)
	fmt.Printf("  InstanceSize: %s\n", project1.InstanceSize)
	fmt.Printf("  DatabasePassword: %s\n", "***hidden***")
	fmt.Println()

	// Example 2: Create a project with all fields from faker
	fmt.Println("Example 2: All Faker Project")
	project2 := NewProjectWithFaker()
	fmt.Printf("  OrganizationID: %s\n", project2.OrganizationID)
	fmt.Printf("  Name: %s\n", project2.Name)
	fmt.Printf("  Region: %s\n", project2.Region)
	fmt.Printf("  InstanceSize: %s\n", project2.InstanceSize)
	fmt.Println()

	// Example 3: Customize a project using builder pattern
	fmt.Println("Example 3: Customized Project")
	project3 := NewProject().
		WithOrganizationID("my-company").
		WithName("production-database").
		WithRegion("us-east-1").
		WithInstanceSize("large").
		WithDatabasePassword("CustomPass123!")

	fmt.Printf("  OrganizationID: %s\n", project3.OrganizationID)
	fmt.Printf("  Name: %s\n", project3.Name)
	fmt.Printf("  Region: %s\n", project3.Region)
	fmt.Printf("  InstanceSize: %s\n", project3.InstanceSize)
	fmt.Println()

	// Example 4: Convert to Terraform variables
	fmt.Println("Example 4: Terraform Variables")
	tfVars := project3.ToMap()
	fmt.Printf("  Map contains %d variables:\n", len(tfVars))
	for key, value := range tfVars {
		if key == "database_password" {
			fmt.Printf("    %s: ***hidden***\n", key)
		} else {
			fmt.Printf("    %s: %v\n", key, value)
		}
	}
	fmt.Println()

	// Example 5: Using utility functions
	fmt.Println("Example 5: Utility Functions")

	// Get default project
	defaultProject := Default()
	fmt.Printf("  Default().OrganizationID: %s\n", defaultProject.OrganizationID)

	// Get project with custom organization ID
	customProject := DefaultWithOrganizationID("acme-corp")
	fmt.Printf("  DefaultWithOrganizationID('acme-corp'): %s\n", customProject.OrganizationID)

	// Get project with all faker fields
	fakerProject := DefaultWithFaker()
	fmt.Printf("  DefaultWithFaker().OrganizationID: %s (faker-generated)\n", fakerProject.OrganizationID)
	fmt.Println()

	// Example 6: Immutability demonstration
	fmt.Println("Example 6: Immutability")
	original := NewProject()
	modified := original.WithOrganizationID("new-org")
	fmt.Printf("  Original OrganizationID: %s\n", original.OrganizationID)
	fmt.Printf("  Modified OrganizationID: %s\n", modified.OrganizationID)
	fmt.Printf("  Original == Modified? %v\n", original.OrganizationID == modified.OrganizationID)
	fmt.Println()

	// Example 7: Complete workflow for testing
	fmt.Println("Example 7: Complete Test Workflow")
	fmt.Println("  // In your test file:")
	fmt.Println("  func TestSupabaseProject(t *testing.T) {")
	fmt.Println("      t.Parallel()")
	fmt.Println("      ")
	fmt.Println("      // Create Terraform options with faker data")
	fmt.Println("      terraformOptions := DefaultForModuleWithFaker(\"modules/project\")")
	fmt.Println("      ")
	fmt.Println("      // Customize for your test")
	fmt.Println("      terraformOptions.Vars[\"module_enabled\"] = true")
	fmt.Println("      terraformOptions.Vars[\"legacy_api_keys_enabled\"] = false")
	fmt.Println("      ")
	fmt.Println("      // Use in test (commented out for example)")
	fmt.Println("      // defer terraform.Destroy(t, terraformOptions)")
	fmt.Println("      // terraform.InitAndApply(t, terraformOptions)")
	fmt.Println("      ")
	fmt.Println("      // Verify values")
	fmt.Println("      // assert.NotEmpty(t, terraformOptions.Vars[\"organization_id\"])")
	fmt.Println("      // assert.NotEmpty(t, terraformOptions.Vars[\"name\"])")
	fmt.Println("  }")
}

// QuickStart provides a quick reference for common use cases
func QuickStart() map[string]string {
	return map[string]string{
		"Basic Project": `
// Create project with default organization ID
project := NewProject()
// OrganizationID="hadenlabs", other fields=faker
`,
		"All Faker": `
// Create project with all fields from faker
project := NewProjectWithFaker()
// All fields are faker-generated
`,
		"Custom Project": `
// Create and customize project
project := NewProject().
    WithOrganizationID("my-org").
    WithName("production").
    WithRegion("us-east-1")
`,
		"Terraform Variables": `
// Convert to Terraform variables
tfVars := project.ToMap()
// Includes: organization_id, database_password, name, region, instance_size,
// legacy_api_keys_enabled=false, module_enabled=true
`,
		"Custom Boolean Values": `
// Convert with custom boolean values
tfVars := project.ToMapWithCustomValues(true, false)
// legacy_api_keys_enabled=true, module_enabled=false
`,
		"Utility Functions": `
// Get default Terraform options
options := DefaultForModule("modules/project")

// Get Terraform options with faker data
options := DefaultForModuleWithFaker("modules/project")

// Merge with custom values
customValues := map[string]interface{}{
    "name": "test-project",
}
options := TerraformOptions("modules/project", customValues)
`,
	}
}

// CommonPatterns shows common testing patterns
func CommonPatterns() []string {
	return []string{
		"Pattern 1: Basic test with defaults",
		"  project := NewProject()",
		"  options := DefaultForModule(\"modules/project\")",
		"",
		"Pattern 2: Test with custom organization",
		"  project := NewProject().WithOrganizationID(\"test-org\")",
		"  options := DefaultForModuleWithOrganizationID(\"modules/project\", \"test-org\")",
		"",
		"Pattern 3: Test with all faker data",
		"  project := NewProjectWithFaker()",
		"  options := DefaultForModuleWithFaker(\"modules/project\")",
		"",
		"Pattern 4: Test suite with different organizations",
		"  orgs := []string{\"hadenlabs\", \"acme-corp\", \"test-org\"}",
		"  for _, org := range orgs {",
		"      options := DefaultForModuleWithOrganizationID(\"modules/project\", org)",
		"      // Run test with this organization",
		"  }",
		"",
		"Pattern 5: Integration test with custom values",
		"  customValues := map[string]interface{}{",
		"      \"name\": \"integration-test\",",
		"      \"region\": \"us-west-1\",",
		"      \"module_enabled\": true,",
		"  }",
		"  options := TerraformOptions(\"modules/project\", customValues)",
	}
}
