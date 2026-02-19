#
# ----------------------------------------------------------------------------------------------------------------------
# REQUIRED PARAMETERS
# These variables must be set when using this module.
# ----------------------------------------------------------------------------------------------------------------------

variable "database_password" {
  type        = string
  description = "(Required, Sensitive) Password for the project database"
  sensitive   = true
}

variable "name" {
  type        = string
  description = "(Required) Name of the project"
}

variable "organization_id" {
  type        = string
  description = "(Required) Organization slug (found in the Supabase dashboard URL or organization settings)"
}

variable "region" {
  type        = string
  description = "(Required) Region where the project is located"
  default     = "us-east-1"
}

# ----------------------------------------------------------------------------------------------------------------------
# OPTIONAL PARAMETERS
# These variables have defaults, but may be overridden.
# ----------------------------------------------------------------------------------------------------------------------

variable "instance_size" {
  type        = string
  description = "(Optional) Desired instance size of the project"
  default     = "micro"
}

variable "legacy_api_keys_enabled" {
  type        = bool
  description = <<-EOT
    (Optional, Deprecated) Controls whether anon and service_role JWT-based api keys should be enabled.
    Please note: these keys are no longer recommended (more information here).
  EOT
  default     = null
}

# ----------------------------------------------------------------------------------------------------------------------
# MODULE CONFIGURATION PARAMETERS
# These variables are used to configure the module.
# See https://medium.com/mineiros/the-ultimate-guide-on-how-to-write-terraform-modules-part-1-81f86d31f02
# ----------------------------------------------------------------------------------------------------------------------

variable "module_enabled" {
  type        = bool
  description = "(Optional) Whether to create resources within the module or not. Default is true."
  default     = true
}
