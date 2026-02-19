# TestUtil Implementation Summary

## Overview

The `testutil` package provides a comprehensive set of utilities for testing Terraform modules, specifically designed for the Supabase project. It offers consistent default values, easy-to-use test configurations, and integration with the existing faker package.

## Core Components

### 1. DefaultValues Struct

The `DefaultValues` struct is the foundation of the package, providing strongly-typed default values for Supabase project testing:

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

### 2. Default Values

The package provides sensible defaults for all Supabase project variables:

| Field                  | Default Value          | Description               |
| ---------------------- | ---------------------- | ------------------------- |
| `DatabasePassword`     | `"SecurePassword123!"` | Secure database password  |
| `Name`                 | `"test-project"`       | Project name              |
| `OrganizationID`       | `"org-test123"`        | Organization ID           |
| `Region`               | `"us-east-1"`          | AWS region                |
| `InstanceSize`         | `"small"`              | Instance size             |
| `LegacyAPIKeysEnabled` | `false`                | Legacy API keys disabled  |
| `ModuleEnabled`        | `true`                 | Module enabled by default |

### 3. Builder Pattern Methods

The struct uses a builder pattern with fluent interface methods:

```go
defaults := NewDefaultValues().
    WithName("custom-project").
    WithRegion("eu-west-1").
    WithModuleEnabled(false)
```

All `With*` methods return `*DefaultValues` for method chaining.

### 4. Conversion to Terraform Variables

The `ToMap()` method converts the struct to a `map[string]interface{}` compatible with Terraform options:

```go
func (d *DefaultValues) ToMap() map[string]interface{} {
    return map[string]interface{}{
        "database_password":       d.DatabasePassword,
        "name":                    d.Name,
        "organization_id":         d.OrganizationID,
        "region":                  d.Region,
        "instance_size":           d.InstanceSize,
        "legacy_api_keys_enabled": d.LegacyAPIKeysEnabled,
        "module_enabled":          d.ModuleEnabled,
    }
}
```

## Helper Functions

### 1. Basic Utilities

- `Default()` - Returns a new `DefaultValues` instance with static defaults
- `DefaultWithFaker()` - Returns defaults with faker-generated values
- `DefaultForModule(dir)` - Creates Terraform options with defaults for a module
- `DefaultForModuleWithFaker(dir)` - Creates Terraform options with faker-generated values

### 2. Value Merging

- `MergeDefaults(customValues)` - Merges custom values with static defaults
- `MergeDefaultsWithFaker(customValues)` - Merges custom values with faker-generated defaults
- `TerraformOptions(dir, customValues)` - Creates Terraform options with merged values
- `TerraformOptionsWithFaker(dir, customValues)` - Creates Terraform options with faker-generated merged values

## Usage Examples

### Basic Usage

```go
// Get default values
defaults := testutil.Default()

// Customize
customDefaults := defaults.
    WithName("my-project").
    WithRegion("eu-west-1")

// Convert to Terraform variables
tfVars := customDefaults.ToMap()
```

### In Tests

```go
func TestProjectBasic(t *testing.T) {
    t.Parallel()

    // Create Terraform options with defaults
    terraformOptions := testutil.DefaultForModule("project-basic").
        WithName("test-project").
        WithModuleEnabled(true)

    defer terraform.Destroy(t, terraformOptions)
    terraform.InitAndApply(t, terraformOptions)

    // Test assertions...
}
```

### With Faker Integration

```go
func TestProjectWithFaker(t *testing.T) {
    t.Parallel()

    // Use faker-generated values
    terraformOptions := testutil.DefaultForModuleWithFaker("project-basic")

    defer terraform.Destroy(t, terraformOptions)
    terraform.InitAndApply(t, terraformOptions)

    // Test assertions...
}
```

### With Custom Values

```go
func TestProjectCustom(t *testing.T) {
    t.Parallel()

    customValues := map[string]interface{}{
        "name":   "integration-test",
        "region": "eu-central-1",
    }

    terraformOptions := testutil.TerraformOptions("project-basic", customValues)

    defer terraform.Destroy(t, terraformOptions)
    terraform.InitAndApply(t, terraformOptions)

    // Test assertions...
}
```

## Benefits

### 1. Consistency

- All tests use the same base defaults
- Reduces duplication of test configuration code
- Ensures consistent test behavior

### 2. Maintainability

- Changes to defaults only need to be made in one place
- Clear separation between test logic and configuration
- Easy to update when module variables change

### 3. Flexibility

- Can use static defaults or faker-generated values
- Easy to override specific values while keeping others
- Supports both simple and complex test scenarios

### 4. Type Safety

- Compile-time checking of field names
- IDE autocompletion for all methods
- Prevents runtime errors from typos

### 5. Integration

- Works seamlessly with existing faker package
- Compatible with Terratest framework
- Easy to extend for new modules

## File Structure

```
internal/testutil/
├── structs.go           # DefaultValues struct and methods
├── testutil.go          # Helper functions
├── structs_test.go      # Unit tests for DefaultValues
├── testutil_test.go     # Unit tests for helper functions
├── example_test.go      # Usage examples
├── README.md           # Documentation
└── IMPLEMENTATION.md   # This file
```

## Testing

The package includes comprehensive tests:

1. **Unit Tests**: Test all struct methods and helper functions
2. **Example Tests**: Demonstrate real-world usage patterns
3. **Integration Examples**: Show how to use in actual Terraform tests

Run tests with:

```bash
go test ./internal/testutil/... -v
```

## Future Extensions

The architecture is designed to be extensible:

1. **Additional Modules**: Can add support for other Terraform modules
2. **Environment-Specific Defaults**: Could add presets for dev/staging/prod
3. **Validation**: Add validation methods for test values
4. **Serialization**: Add JSON/YAML serialization for external configuration

## Conclusion

The `testutil` package provides a robust, type-safe foundation for testing Terraform modules in the Supabase project. It simplifies test configuration, ensures consistency, and integrates seamlessly with existing testing infrastructure.
