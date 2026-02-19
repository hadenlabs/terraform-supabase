# ----------------------------------------------------------------------------------------------------------------------
# SET TERRAFORM AND PROVIDER REQUIREMENTS FOR RUNNING THIS MODULE
# ----------------------------------------------------------------------------------------------------------------------

terraform {
  required_version = ">= 1.0.0"

  required_providers {
    supabase = {
      source  = "supabase/supabase"
      version = "1.7.0"
    }

    null = {
      source  = "hashicorp/null"
      version = ">= 2.0"
    }
  }
}
