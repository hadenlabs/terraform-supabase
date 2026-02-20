# Integration Tests for Supabase Project Module

This directory contains integration tests for the Supabase Project Terraform module using Terratest.

## Test Structure

### Test Structure

- **Test Files**: Go files with `//go:build integration` build tag
- **Test Directories**: Terraform configurations for each test scenario
  - `project-basic/` - Basic test with module enabled and all required variables
  - `project-disabled/` - Test with module disabled (`module_enabled = false`)

### Test Files

- `project_basic_test.go` - Tests basic project creation (integration tag)
- `project_disabled_test.go` - Tests module when disabled (integration tag)
- `dummy.go` - Makes directory a valid Go package

## Running Tests

### Prerequisites

1. **Go 1.24+** installed
2. **Terraform 1.0.0+** installed
3. **Supabase Access Token** - Required for the Supabase provider
   - Set as environment variable: `SUPABASE_ACCESS_TOKEN`
   - Or configure in provider configuration

### Environment Setup

```bash
# Set Supabase access token (required for tests)
export SUPABASE_ACCESS_TOKEN="your-supabase-access-token"

# From project root, run integration tests for this module
go test -tags=integration -race -v ./modules/project/test/... -timeout 60m

# Or run all integration tests
go test -tags=integration -race -v ./... -timeout 60m
```

### Running Tests

```bash
# Run integration tests for this module
go test -tags=integration -race -v ./modules/project/test/... -timeout 60m

# Run specific integration test
go test -tags=integration -race -v ./modules/project/test/... -run TestProjectBasicSuccess
go test -tags=integration -race -v ./modules/project/test/... -run TestProjectDisabledSuccess

# Run tests in parallel
go test -tags=integration -race -v -parallel 10 ./modules/project/test/... -timeout 60m
```

### Using Taskfile

The main project Taskfile includes test commands:

```bash
# From project root, run all tests
task test

# Run tests for specific module
go test -race -v ./modules/project/... -coverprofile cover.out -timeout 60m
```

## Test Details

**Note**: These are integration tests that require actual Supabase credentials and will create real resources.

### TestProjectBasicSuccess

Tests the basic functionality of the module with:

- All required variables provided
- Module enabled (`module_enabled = true`)
- Optional variables with values
- Verifies that:
  - Project ID is returned (not empty)
  - Module enabled output is `true`

### TestProjectDisabledSuccess

Tests the module when disabled with:

- All required variables provided
- Module disabled (`module_enabled = false`)
- Optional variables as `null`
- Verifies that:
  - Project ID is empty (no resource created)
  - Module enabled output is `false`

## Test Data Generation

Tests use the `faker` package to generate realistic test data:

- **Project Names**: Random project names with prefixes like `backend-`, `api-`, etc.
- **Organization IDs**: Generated as `org-xxxxxx` format
- **Regions**: Valid Supabase regions
- **Instance Sizes**: Valid instance sizes (small, medium, large, xlarge)
- **Database Passwords**: Secure passwords with mixed characters

## Test Cleanup

All tests use Terratest's `defer terraform.Destroy()` pattern to ensure resources are cleaned up after tests, even if tests fail.

## Build Tags

Tests in this directory use the `integration` build tag (`//go:build integration`) to separate them from unit tests. This allows:

- Running only unit tests: `go test ./...`
- Running only integration tests: `go test -tags=integration ./...`
- Running all tests: `go test -tags=integration ./...` (plus unit tests separately)

## Adding New Tests

When adding new integration tests:

1. Create a new test directory with Terraform configuration
2. Create a corresponding Go test file with `//go:build integration` at the top
3. Follow the existing patterns for:
   - Test function naming (`TestXxxSuccess`)
   - Parallel execution (`t.Parallel()`)
   - Resource cleanup (`defer terraform.Destroy()`)
   - Assertions using `testify/assert`
4. Use the `faker` package for generating test data

## Troubleshooting

### Common Issues

1. **Provider Authentication Errors**

   ```
   Error: error configuring Supabase provider: access token is required
   ```

   Solution: Set `SUPABASE_ACCESS_TOKEN` environment variable

2. **Organization Not Found**

   ```
   Error: organization not found
   ```

   Solution: Use a valid organization ID that exists in your Supabase account

3. **Region Not Available**

   ```
   Error: region not available
   ```

   Solution: Check available regions in your Supabase organization

4. **Build Tag Issues**
   ```
   no packages to test
   ```
   Solution: Use `-tags=integration` flag when running integration tests

### Debugging Tests

```bash
# Run integration tests with detailed output
go test -tags=integration -v -count=1 ./modules/project/test/...

# Run with race detector
go test -tags=integration -v -race ./modules/project/test/...

# Generate test coverage report
go test -tags=integration -v -coverprofile=coverage.out ./modules/project/test/...
go tool cover -html=coverage.out
```

## Dependencies

- **Terratest**: For Terraform integration testing
- **Testify**: For assertions and test utilities
- **Faker**: For generating test data (internal package)

Dependencies are managed in the root `go.mod` file, not in this directory.

## Notes

- Tests create real Supabase resources and may incur costs
- Always ensure tests clean up resources properly
- Use unique names for resources to avoid conflicts
- Consider rate limits when running tests frequently
- Integration tests are separated from unit tests using build tags
- Tests follow the project's Go 1.24 and modular structure requirements
