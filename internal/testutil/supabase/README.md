# Supabase TestUtil Package

A lightweight, focused package for managing Supabase project configurations in tests, specifically designed for the `terraform-supabase` project.

## Overview

The `supabase` package provides a simple, immutable structure for managing Supabase project organization IDs with sensible defaults and utility functions for test configuration.

## Installation

```go
import "github.com/hadenlabs/terraform-supabase/internal/testutil/supabase"
```

## Core Structure

### Project Struct

```go
type Project struct {
    // OrganizationID is the organization identifier for Supabase projects
    OrganizationID string
}
```

### Default Values

- **OrganizationID**: `"hadenlabs"` (default organization ID)

## Basic Usage

### Creating a Project

```go
// Create with default values
project := supabase.NewProject()
// OrganizationID: "hadenlabs"

// Create with custom organization ID
customProject := supabase.NewProject().WithOrganizationID("my-company")
// OrganizationID: "my-company"
```

### Immutable Builder Pattern

The package uses an immutable builder pattern:

```go
original := supabase.NewProject()           // OrganizationID: "hadenlabs"
modified := original.WithOrganizationID("custom")  // Returns NEW instance

// Original remains unchanged
fmt.Println(original.OrganizationID)        // "hadenlabs"
fmt.Println(modified.OrganizationID)        // "custom"
```

### Converting to Terraform Variables

```go
project := supabase.NewProject().WithOrganizationID("acme-corp")
tfVars := project.ToMap()
// map[string]interface{}{"organization_id": "acme-corp"}
```

## Utility Functions

### Basic Utilities

```go
// Get default project
defaultProject := supabase.Default()

// Get project with custom organization ID
customProject := supabase.DefaultWithOrganizationID("my-org")

// Check if organization ID is default
isDefault := supabase.IsDefaultOrganizationID("hadenlabs") // true

// Validate organization ID
isValid := supabase.ValidateOrganizationID("valid-org") // true
```

### Terraform Integration

```go
// Create Terraform options with default organization ID
options1 := supabase.DefaultForModule("modules/project")

// Create Terraform options with custom organization ID
options2 := supabase.DefaultForModuleWithOrganizationID("modules/project", "my-org")

// Merge custom values with defaults
customValues := map[string]interface{}{
    "name":   "test-project",
    "region": "us-east-1",
}
options3 := supabase.TerraformOptions("modules/project", customValues)

// Merge with specific organization ID
options4 := supabase.TerraformOptionsWithOrganizationID("modules/project", "specific-org", customValues)
```

### Value Management

```go
// Get organization ID from variables
orgID := supabase.GetOrganizationID(tfVars)

// Set organization ID in variables
newVars := supabase.SetOrganizationID(tfVars, "new-org")

// Merge values
merged := supabase.MergeProjectValues(customValues)
mergedWithOrg := supabase.MergeProjectValuesWithOrganizationID("custom-org", customValues)
```

## Usage Examples

### Basic Test Example

```go
func TestProjectBasic(t *testing.T) {
    t.Parallel()

    // Create Terraform options with default organization ID
    terraformOptions := supabase.DefaultForModule("modules/project/test/project-basic")

    // Customize other variables
    terraformOptions.Vars["name"] = "test-project"
    terraformOptions.Vars["module_enabled"] = true

    // Use in test
    defer terraform.Destroy(t, terraformOptions)
    terraform.InitAndApply(t, terraformOptions)

    // Verify
    assert.Equal(t, "hadenlabs", terraformOptions.Vars["organization_id"])
}
```

### Test with Custom Organization

```go
func TestProjectCustomOrganization(t *testing.T) {
    t.Parallel()

    terraformOptions := supabase.DefaultForModuleWithOrganizationID(
        "modules/project/test/project-basic",
        "my-company-org",
    )

    // Add other required variables
    terraformOptions.Vars["name"] = "company-project"
    terraformOptions.Vars["database_password"] = "SecurePass123!"

    // Test logic...
}
```

### Test Suite with Multiple Organizations

```go
func TestProjectOrganizationSuite(t *testing.T) {
    t.Parallel()

    testCases := []struct {
        name string
        orgID string
    }{
        {"default", "hadenlabs"},
        {"custom", "acme-corp"},
        {"numeric", "org-12345"},
    }

    for _, tc := range testCases {
        t.Run(tc.name, func(t *testing.T) {
            t.Parallel()

            options := supabase.DefaultForModuleWithOrganizationID(
                "modules/project/test/project-basic",
                tc.orgID,
            )

            // Test with this organization...
        })
    }
}
```

## Features

### 1. **Simplicity**

- Single-purpose structure focused on organization ID management
- Minimal API surface area
- Easy to understand and use

### 2. **Immutability**

- All operations return new instances
- Prevents accidental mutation of shared configurations
- Thread-safe by design

### 3. **Consistency**

- Consistent default values across all tests
- Predictable behavior
- Easy to maintain and update

### 4. **Terraform Integration**

- Seamless integration with Terratest
- Easy conversion to Terraform variables
- Support for merging with custom values

### 5. **Validation**

- Built-in validation functions
- Default value detection
- Input sanitization

## Best Practices

1. **Use Defaults When Possible**: Start with `supabase.Default()` or `supabase.NewProject()` for consistency.

2. **Immutable Operations**: Remember that `WithOrganizationID()` returns a new instance - assign it to a variable.

3. **Validate Inputs**: Use `ValidateOrganizationID()` for user-provided organization IDs.

4. **Clear Naming**: Use descriptive variable names for different project configurations.

5. **Test Organization Variations**: Test with different organization IDs to ensure compatibility.

## File Structure

```
internal/testutil/supabase/
├── structs.go           # Project struct and methods
├── util.go              # Utility functions
├── structs_test.go      # Unit tests for Project struct
├── util_test.go         # Unit tests for utility functions
├── example_test.go      # Usage examples
└── README.md           # This file
```

## Testing

Run all tests:

```bash
go test ./internal/testutil/supabase/... -v
```

## See Also

- [Main TestUtil Package](../README.md) - Comprehensive test utilities
- [Terraform Supabase Project](../../../../README.md) - Main project documentation
- [Terratest Documentation](https://github.com/gruntwork-io/terratest) - Terraform testing framework

## License

Part of the `terraform-supabase` project.
