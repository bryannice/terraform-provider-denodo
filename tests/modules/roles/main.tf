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
# Creating Database Role
# -----------------------------------------------------------------------------

resource "denodo_database_role" "role" {
  database_name       = var.database_name
  name                = var.name
  admin               = var.admin
  connect             = var.connect
  create              = var.create
  create_data_service = var.create_data_service
  create_data_source  = var.create_data_source
  create_folder       = var.create_folder
  create_view         = var.create_view
  execute             = var.execute
  file                = var.file
  grant               = var.grant
  meta_data           = var.meta_data
  monitor_admin       = var.monitor_admin
  write               = var.write
}
