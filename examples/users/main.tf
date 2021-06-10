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
# Setting the Denodo Provider
# -----------------------------------------------------------------------------

provider "denodo" {
  database = var.denodo_database
  host     = var.denodo_host
  password = var.denodo_password
  port     = var.denodo_port
  ssl_mode = "require"
  username = var.denodo_username
}

# -----------------------------------------------------------------------------
# Reading from Remote State File from Roles
# -----------------------------------------------------------------------------

data "terraform_remote_state" "roles" {
  backend = "local"
  config = {
    path = "../roles/terraform.tfstate"
  }
}

# -----------------------------------------------------------------------------
# Creating Read User Account
# -----------------------------------------------------------------------------

resource "denodo_user" "db_usr_read" {
  admin       = false
  description = "test read user"
  password    = var.denodo_read_user_password
  roles       = data.terraform_remote_state.roles.outputs.roles[0].role_name
  username    = "test_read_user"
}

# -----------------------------------------------------------------------------
# Creating Read Developer Account
# -----------------------------------------------------------------------------

resource "denodo_user" "db_usr_dev" {
  admin       = false
  description = "test dev user"
  password    = var.denodo_dev_user_password
  roles       = data.terraform_remote_state.roles.outputs.roles[1].role_name
  username    = "test_dev_user"
}
