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
# Reading from Remote State File from Folder
# -----------------------------------------------------------------------------

data "terraform_remote_state" "folder" {
  backend = "local"
  config = {
    path = "../folders/terraform.tfstate"
  }
}

# -----------------------------------------------------------------------------
# Create Dervived View
# -----------------------------------------------------------------------------

resource "denodo_dervived_view" "dv" {
  database  = data.terraform_remote_state.folder.outputs.database[2]
  directory = "test_files"
}
