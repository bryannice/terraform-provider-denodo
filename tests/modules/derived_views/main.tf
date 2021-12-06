# -----------------------------------------------------------------------------
# This is a testing example of the Denodo Provider
# -----------------------------------------------------------------------------

terraform {
  required_providers {
    denodo = {
      version = "0.1"
      source  = "custom.com/bryannice/denodo"
    }
  }
}

# -----------------------------------------------------------------------------
# Create Derived View
# -----------------------------------------------------------------------------

resource "denodo_dervived_view" "dv" {
  database = var.database
  name     = var.name
  vql      = var.vql
}
