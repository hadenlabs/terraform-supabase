module "supabase_project" {
  source = "../.."

  # Required variables
  database_password = var.database_password
  name              = var.name
  organization_id   = var.organization_id
  region            = var.region

  # Optional variables
  instance_size           = var.instance_size
  legacy_api_keys_enabled = var.legacy_api_keys_enabled

  # Module configuration
  module_enabled = var.module_enabled
}
