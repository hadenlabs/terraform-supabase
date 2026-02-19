# ----------------------------------------------------------------------------------------------------------------------
# OUTPUT CALCULATED VARIABLES (prefer full objects)
# ----------------------------------------------------------------------------------------------------------------------

output "id" {
  description = "API key identifier"
  value       = local.outputs.module_enabled ? one(supabase_apikey.this.*.id) : null
}

output "api_key" {
  description = "API key (sensitive)"
  value       = local.outputs.module_enabled ? one(supabase_apikey.this.*.api_key) : null
  sensitive   = true
}

output "type" {
  description = "Type of the API key"
  value       = local.outputs.module_enabled ? one(supabase_apikey.this.*.type) : null
}

output "secret_jwt_template" {
  description = "Secret JWT template"
  value       = local.outputs.module_enabled ? one(supabase_apikey.this.*.secret_jwt_template) : null
}

# ----------------------------------------------------------------------------------------------------------------------
# OUTPUT ALL RESOURCES AS FULL OBJECTS
# ----------------------------------------------------------------------------------------------------------------------

output "apikey" {
  description = "All attributes of the created API key resource"
  value       = local.outputs.module_enabled ? one(supabase_apikey.this.*) : null
}

# ----------------------------------------------------------------------------------------------------------------------
# OUTPUT ALL INPUT VARIABLES
# ----------------------------------------------------------------------------------------------------------------------

output "project_ref" {
  description = "Project reference ID"
  value       = local.outputs.project_ref
}

output "name" {
  description = "Name of the API key"
  value       = local.outputs.name
}

output "description" {
  description = "Description of the API key"
  value       = local.outputs.description
}

# ----------------------------------------------------------------------------------------------------------------------
# OUTPUT MODULE CONFIGURATION
# ----------------------------------------------------------------------------------------------------------------------

output "module_enabled" {
  description = "Whether the module is enabled."
  value       = local.outputs.module_enabled
}
