variable "database_password" {
  type        = string
  description = "Password for the project database"
  sensitive   = true
}

variable "name" {
  type        = string
  description = "Name of the project"
}

variable "organization_id" {
  type        = string
  description = "Organization slug"
}

variable "region" {
  type        = string
  description = "Region where the project is located"
}

variable "instance_size" {
  type        = string
  description = "Desired instance size of the project"
  default     = null
}

variable "legacy_api_keys_enabled" {
  type        = bool
  description = "Controls whether anon and service_role JWT-based api keys should be enabled"
  default     = null
}

variable "module_enabled" {
  type        = bool
  description = "Whether to create resources within the module or not"
  default     = true
}
