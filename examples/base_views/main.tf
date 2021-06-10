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
# Reading from Remote State File from Data Source
# -----------------------------------------------------------------------------

data "terraform_remote_state" "ds" {
  backend = "local"
  config = {
    path = "../jdbc_data_source/terraform.tfstate"
  }
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
# Creating Base Views From JDBC Data Source
# -----------------------------------------------------------------------------

resource "denodo_base_view" "db_bv" {
  count                    = length(data.terraform_remote_state.ds.outputs.object_list.*)
  data_source_catalog_name = var.data_source_catalog_name
  data_source_database     = data.terraform_remote_state.ds.outputs.data_source_database
  data_source_name         = data.terraform_remote_state.ds.outputs.data_source_name
  data_source_schema_name  = var.data_source_schema_name
  data_source_table_name   = data.terraform_remote_state.ds.outputs.object_list[count.index].object_name
  database                 = data.terraform_remote_state.folder.outputs.data[1].database
  folder                   = data.terraform_remote_state.folder.outputs.data[1].folder
  name                     = format("bv_%s", data.terraform_remote_state.ds.outputs.object_list[count.index].object_name)
}
