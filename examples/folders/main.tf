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
# Reading from Remote State File from Virtual Database
# -----------------------------------------------------------------------------

data "terraform_remote_state" "vbd" {
  backend = "local"
  config = {
    path = "../virtual_database/terraform.tfstate"
  }
}

# -----------------------------------------------------------------------------
# Creating Data Source Folder in Database
# -----------------------------------------------------------------------------

resource "denodo_database_folder" "db_folder_ds" {
  database    = data.terraform_remote_state.vbd.outputs.id
  folder_path = "/data_source"
}

# -----------------------------------------------------------------------------
# Creating Base View Folder in Database
# -----------------------------------------------------------------------------

resource "denodo_database_folder" "db_folder_bv" {
  database    = data.terraform_remote_state.vbd.outputs.id
  folder_path = "/base_view"
}