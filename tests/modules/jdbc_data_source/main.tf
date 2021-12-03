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
# Creating JDBC Data Source
# -----------------------------------------------------------------------------

resource "denodo_jdbc_data_source" "db_ds" {
  class_path                   = var.data_source_class_path
  data_source_database_type    = var.data_source_database_type
  data_source_database_version = var.data_source_database_version
  database_uri                 = var.data_source_database_uri
  denodo_database              = var.data_source_database
  driver_class_name            = var.data_source_driver_class_name
  folder                       = var.data_source_folder
  name                         = var.data_source_name
  password                     = var.data_source_password
  username                     = var.data_source_username
}
