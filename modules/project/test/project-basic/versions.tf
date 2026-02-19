# ----------------------------------------------------------------------------------------------------------------------
# SET TERRAFORM AND PROVIDER REQUIREMENTS FOR RUNNING THIS TEST
# ----------------------------------------------------------------------------------------------------------------------

terraform {
  required_version = ">= 1.0.0"

  required_providers {
    supabase = {
      source  = "supabase/supabase"
      version = "1.7.0"
    }
  }
}

provider "supabase" {
  # Configure the Supabase provider
  # Access token should be provided via environment variable SUPABASE_ACCESS_TOKEN
  # or via terraform.tfvars
}
