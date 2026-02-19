# Terraform Module: Supabase API Key

This Terraform module creates and manages Supabase API keys.

## Usage

```hcl
module "supabase_apikey" {
  source = "github.com/hadenlabs/terraform-supabase//modules/apikey"

  project_ref = "your-project-ref"
  name        = "my-api-key"
  description = "API key for external service integration"

  module_enabled = true
}
```

## Requirements

| Name      | Version  |
| --------- | -------- |
| terraform | >= 1.0.0 |
| supabase  | 1.7.0    |

## Providers

| Name     | Version |
| -------- | ------- |
| supabase | 1.7.0   |

## Resources

| Name                 | Type     |
| -------------------- | -------- |
| supabase_apikey.this | resource |

## Inputs

| Name | Description | Type | Default | Required |
| --- | --- | --- | --- | :-: |
| description | Description of the API key | `string` | `null` | no |
| module_enabled | Whether to create resources within the module or not. Default is true. | `bool` | `true` | no |
| name | Name of the API key | `string` | n/a | yes |
| project_ref | Project reference ID | `string` | n/a | yes |

## Outputs

| Name                | Description                                    |
| ------------------- | ---------------------------------------------- |
| api_key             | API key (sensitive)                            |
| apikey              | All attributes of the created API key resource |
| description         | Description of the API key                     |
| id                  | API key identifier                             |
| module_enabled      | Whether the module is enabled.                 |
| name                | Name of the API key                            |
| project_ref         | Project reference ID                           |
| secret_jwt_template | Secret JWT template                            |
| type                | Type of the API key                            |

## Examples

### Basic Example

```hcl
module "supabase_apikey" {
  source = "github.com/hadenlabs/terraform-supabase//modules/apikey"

  project_ref = "mayuaycdtijbctgqbycg"
  name        = "production-api-key"
}
```

### API Key with Description

```hcl
module "supabase_apikey" {
  source = "github.com/hadenlabs/terraform-supabase//modules/apikey"

  project_ref = "mayuaycdtijbctgqbycg"
  name        = "external-service-key"
  description = "API key for external service integration"
}
```

### Conditional Creation

```hcl
module "supabase_apikey" {
  source = "github.com/hadenlabs/terraform-supabase//modules/apikey"

  project_ref    = "mayuaycdtijbctgqbycg"
  name           = "conditional-key"
  description    = "Created only when needed"
  module_enabled = var.create_api_key
}
```

## Import

API keys can be imported using the following syntax:

```shell
# The ID is the project reference and a unique identifier of the key separated by '/'
terraform import module.supabase_apikey.supabase_apikey.this <project_ref>/<key_id>
```

## Notes

- The `api_key` output is marked as sensitive and will not be displayed in Terraform output.
- API keys are specific to a Supabase project identified by the `project_ref`.
- The `secret_jwt_template` contains the role configuration for JWT-based authentication.
