<!-- markdown-link-check-disable -->
<!-- BEGIN_TF_DOCS -->
## Requirements

| Name | Version |
|------|---------|
| terraform | >= 1.0.0 |
| null | >= 2.0 |
| supabase | 1.7.0 |

## Providers

| Name | Version |
|------|---------|
| supabase | 1.7.0 |

## Modules

No modules.

## Resources

| Name | Type |
|------|------|
| [supabase_project.this](https://registry.terraform.io/providers/supabase/supabase/1.7.0/docs/resources/project) | resource |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| database\_password | (Required, Sensitive) Password for the project database | `string` | n/a | yes |
| instance\_size | (Optional) Desired instance size of the project | `string` | `"micro"` | no |
| legacy\_api\_keys\_enabled | (Optional, Deprecated) Controls whether anon and service\_role JWT-based api keys should be enabled.<br/>Please note: these keys are no longer recommended (more information here). | `bool` | `null` | no |
| module\_enabled | (Optional) Whether to create resources within the module or not. Default is true. | `bool` | `true` | no |
| name | (Required) Name of the project | `string` | n/a | yes |
| organization\_id | (Required) Organization slug (found in the Supabase dashboard URL or organization settings) | `string` | n/a | yes |
| region | (Required) Region where the project is located | `string` | `"us-east-1"` | no |

## Outputs

| Name | Description |
|------|-------------|
| id | id of user |
| module\_enabled | Whether the module is enabled. |
<!-- END_TF_DOCS -->
<!-- markdown-link-check-enable -->