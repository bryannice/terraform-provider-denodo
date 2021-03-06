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
# Creating Users
# -----------------------------------------------------------------------------

resource "denodo_user" "user" {
  admin       = var.admin
  description = var.description
  password    = var.password
  roles       = var.roles
  username    = var.username
}
