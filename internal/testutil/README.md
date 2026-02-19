# TestUtil Package

The `testutil` package provides a comprehensive set of utilities for testing Terraform modules, with a focus on providing consistent default values and easy-to-use test configurations.

## Overview

This package is designed to simplify the creation of test configurations for Terraform modules, particularly for the Supabase project. It provides:

1. **Default Values**: Consistent, reusable default values for testing
2. **Faker Integration**: Integration with the faker package for realistic test data
3. **Helper Functions**: Utilities for creating Terraform options and merging configurations
4. **Type Safety**: Strongly-typed structures for better code completion and error checking

## Installation

The package is part of the internal utilities and can be imported as:

```go
import "github.com/hadenlabs/terraform-supabase/internal/testutil"
```

## Core Components

### DefaultValues Struct

The `DefaultValues` struct provides a centralized way to manage default values for testing:

```go
type DefaultValues struct {
    DatabasePassword     string
    Name                 string
    OrganizationID       string
    Region               string
    InstanceSize         string
    LegacyAPIKeysEnabled bool
    ModuleEnabled        bool
}
```

### Creating Default Values

```go
// Get default values
defaults := testutil.Default()

// Get faker-generated values
fakerDefaults := testutil.DefaultWithFaker()

// Customize defaults
customDefaults := testutil.Default().
    WithName("my-project").
    WithRegion("eu-west-1").
    WithModuleEnabled(false)
```

### Converting to Terraform Variables

```go
// Convert to map for Terraform
tfVars := defaults.ToMap()

// Result:
// map[string]interface{}{
//     "database_password":       "SecurePassword123!",
//     "name":                    "test-project",
//     "organization_id":         "org-test123",
//     "region":                  "us-east-1",
//     "instance_size":           "small",
//     "legacy_api_keys_enabled": false,
//     "module_enabled":          true,
// }
```

## Helper Functions

### Creating Terraform Options

```go
// Default values for a module
options1 := testutil.DefaultForModule("modules/project")

// Faker-generated values for a module
options2 := testutil.DefaultForModuleWithFaker("modules/project")

// Merge custom values with defaults
customValues := map[string]interface{}{
    "name":   "custom-project",
    "region": "eu-central-1",
}
options3 := testutil.TerraformOptions("modules/project", customValues)

// Merge custom values with faker-generated defaults
options4 := testutil.TerraformOptionsWithFaker("modules/project", customValues)
```

### Merging Values

```go
// Merge custom values with defaults
merged1 := testutil.MergeDefaults(map[string]interface{}{
    "name": "override-name",
})

// Merge custom values with faker-generated defaults
merged2 := testutil.MergeDefaultsWithFaker(map[string]interface{}{
    "name": "override-name",
})
```

## Usage Examples

### Basic Test Example

```go
func TestProjectBasic(t *testing.T) {
    t.Parallel()

    // Get Terraform options with defaults
    terraformOptions := testutil.DefaultForModule("project-basic").
        WithName("test-basic-project").
        WithModuleEnabled(true)

    defer terraform.Destroy(t, terraformOptions)
    terraform.InitAndApply(t, terraformOptions)

    // Verify outputs
    outputProjectID := terraform.Output(t, terraformOptions, "project_id")
    assert.NotEmpty(t, outputProjectID)
}
```

### Test with Faker Data

```go
func TestProjectWithFaker(t *testing.T) {
    t.Parallel()

    // Use faker-generated values for more realistic testing
    terraformOptions := testutil.DefaultForModuleWithFaker("project-basic").
        WithModuleEnabled(true)

    defer terraform.Destroy(t, terraformOptions)
    terraform.InitAndApply(t, terraformOptions)

    // Test assertions...
}
```

### Test with Custom Values

```go
func TestProjectCustom(t *testing.T) {
    t.Parallel()

    customValues := map[string]interface{}{
        "name":           "integration-test",
        "instance_size":  "large",
        "module_enabled": true,
    }

    terraformOptions := testutil.TerraformOptions("project-basic", customValues)

    defer terraform.Destroy(t, terraformOptions)
    terraform.InitAndApply(t, terraformOptions)

    // Test assertions...
}
```

## Advanced Usage

### Environment-Specific Configurations

```go
// Development configuration
devConfig := testutil.Default().
    WithName("dev-project").
    WithInstanceSize("small").
    WithLegacyAPIKeysEnabled(false)

// Production configuration
prodConfig := testutil.Default().
    WithName("prod-project").
    WithInstanceSize("xlarge").
    WithRegion("us-east-1").
    WithLegacyAPIKeysEnabled(true)
```

### Test Suites with Different Configurations

```go
func TestProjectSuite(t *testing.T) {
    t.Parallel()

    testCases := []struct {
        name        string
        getOptions  func() *terraform.Options
    }{
        {
            name: "basic",
            getOptions: func() *terraform.Options {
                return testutil.DefaultForModule("project-basic")
            },
        },
        {
            name: "with-faker",
            getOptions: func() *terraform.Options {
                return testutil.DefaultForModuleWithFaker("project-basic")
            },
        },
        {
            name: "custom",
            getOptions: func() *terraform.Options {
                return testutil.TerraformOptions("project-basic", map[string]interface{}{
                    "name": "custom-test",
                })
            },
        },
    }

    for _, tc := range testCases {
        t.Run(tc.name, func(t *testing.T) {
            t.Parallel()
            options := tc.getOptions()
            // Run test with options...
        })
    }
}
```

## Best Practices

1. **Use Defaults for Consistency**: Always start with `testutil.Default()` or `testutil.DefaultWithFaker()` to ensure consistent test data.

2. **Override Only What's Necessary**: Use the `With*` methods to override only the values that are relevant to your test case.

3. **Use Faker for Integration Tests**: For integration tests that create real resources, use `DefaultWithFaker()` to avoid naming conflicts.

4. **Keep Tests Isolated**: Each test should have its own configuration to prevent test pollution.

5. **Clean Up Resources**: Always use `defer terraform.Destroy()` to clean up resources created during tests.

## Default Values Reference

| Field | Default Value | Description |
|-------|---------------|-------------|
| `DatabasePassword` | `"SecurePassword123!"` | Database password for the project |
| `Name` | `"test-project"` | Project name |
| `OrganizationID` | `"org-test123"` | Organization ID |
| `Region` | `"us-east-1"` | AWS region |
| `InstanceSize` | `"small"` | Instance size |
| `LegacyAPIKeysEnabled` | `false` | Whether legacy API keys are enabled |
| `ModuleEnabled` | `true` | Whether the module is enabled |

## Testing

Run the testutil package tests:

```bash
go test ./internal/testutil/... -v
```

## See Also

- [Faker Package](../app/external/faker/) - For generating realistic test data
- [Terratest](https://github.com/gruntwork-io/terratest) - For Terraform testing framework
- [Project Module Tests](../../modules/project/test/) - For integration test examples