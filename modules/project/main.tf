locals {
  defaults = {}

  input = {

    database_password = var.database_password
    name              = var.name
    organization_id   = var.organization_id
    region            = var.region

    instance_size           = var.instance_size
    legacy_api_keys_enabled = var.legacy_api_keys_enabled

    # module
    module_enabled = var.module_enabled
  }

  generated = {
    database_password = local.input.database_password
    name              = local.input.name
    organization_id   = local.input.organization_id
    region            = local.input.region

    instance_size           = local.input.instance_size
    legacy_api_keys_enabled = local.input.legacy_api_keys_enabled

    # module
    module_enabled = local.input.module_enabled
  }

  outputs = {
    database_password = local.generated.database_password
    name              = local.generated.name
    organization_id   = local.generated.organization_id
    region            = local.generated.region

    instance_size           = local.generated.instance_size
    legacy_api_keys_enabled = local.generated.legacy_api_keys_enabled

    # module
    module_enabled = local.generated.module_enabled
  }

}

resource "supabase_project" "this" {
  count = local.outputs.module_enabled ? 1 : 0

  organization_id         = local.outputs.organization_id
  name                    = local.outputs.name
  database_password       = local.outputs.database_password
  region                  = local.outputs.region
  instance_size           = local.outputs.instance_size
  legacy_api_keys_enabled = local.outputs.legacy_api_keys_enabled
}
