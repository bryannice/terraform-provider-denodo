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
# Creating JDBC Data Source
# -----------------------------------------------------------------------------

resource "denodo_jdbc_data_source" "db_ds" {
  class_path                   = var.data_source_class_path
  data_source_database_type    = var.data_source_database_type
  data_source_database_version = var.data_source_database_version
  database_uri                 = var.data_source_database_uri
  denodo_database              = data.terraform_remote_state.folder.outputs.database[1]
  driver_class_name            = var.data_source_driver_class_name
  folder                       = data.terraform_remote_state.folder.outputs.folder[1]
  name                         = var.data_source_name
  password                     = var.data_source_password
  username                     = var.data_source_username
}

# -----------------------------------------------------------------------------
# Fetch List of Objects in JDBC Data Source
# -----------------------------------------------------------------------------

data "denodo_jdbc_data_source_object" "jdst" {
  catalog_name = var.data_source_catalog_name
  database     = data.terraform_remote_state.folder.outputs.database[1]
  name         = denodo_jdbc_data_source.db_ds.id
  schema_name  = var.data_source_schema_name
}
