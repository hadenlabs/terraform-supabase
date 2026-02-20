locals {
  defaults = {}

  input = {
    project_ref = var.project_id
    name        = var.name
    description = var.description

    # module
    module_enabled = var.module_enabled
  }

  generated = {
    project_ref = local.input.project_ref
    name        = local.input.name
    description = local.input.description

    # module
    module_enabled = local.input.module_enabled
  }

  outputs = {
    project_ref = local.generated.project_ref
    name        = local.generated.name
    description = local.generated.description

    # module
    module_enabled = local.generated.module_enabled
  }

}

resource "supabase_apikey" "this" {
  count = local.outputs.module_enabled ? 1 : 0

  project_ref = local.outputs.project_ref
  name        = local.outputs.name
  description = local.outputs.description
}
