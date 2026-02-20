output "id" {
  description = "ID of the created apikey"
  value       = module.supabase_apikey.id
}

output "project_id" {
  description = "ID of the created Supabase project"
  value       = module.supabase_project.id
}

output "module_enabled" {
  description = "Whether the module was enabled"
  value       = module.supabase_apikey.module_enabled
}
