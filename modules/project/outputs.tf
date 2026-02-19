# ----------------------------------------------------------------------------------------------------------------------
# OUTPUT CALCULATED VARIABLES (prefer full objects)
# ----------------------------------------------------------------------------------------------------------------------

output "id" {
  description = "id of user"
  value       = local.outputs.module_enabled ? one(supabase_project.this.*.id) : null
}

# ----------------------------------------------------------------------------------------------------------------------
# OUTPUT ALL RESOURCES AS FULL OBJECTS
# ----------------------------------------------------------------------------------------------------------------------


# OUTPUT ALL RESOURCES AS FULL OBJECTS

# OUTPUT ALL INPUT VARIABLES

# OUTPUT MODULE CONFIGURATION


# ----------------------------------------------------------------------------------------------------------------------
# OUTPUT MODULE CONFIGURATION
# ----------------------------------------------------------------------------------------------------------------------

output "module_enabled" {
  description = "Whether the module is enabled."
  value       = local.outputs.module_enabled
}
