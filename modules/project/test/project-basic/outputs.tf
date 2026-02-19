output "project_id" {
  description = "ID of the created Supabase project"
  value       = module.supabase_project.id
}

output "module_enabled" {
  description = "Whether the module was enabled"
  value       = module.supabase_project.module_enabled
}
