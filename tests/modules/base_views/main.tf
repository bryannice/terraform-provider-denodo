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
# Creating Base Views From JDBC Data Source
# -----------------------------------------------------------------------------

resource "denodo_base_view" "bv" {
  data_source_catalog_name = var.data_source_catalog_name
  data_source_database     = var.data_source_database
  data_source_name         = var.data_source_name
  data_source_schema_name  = var.data_source_schema_name
  data_source_table_name   = var.data_source_table_name
  database                 = var.database
  folder                   = var.folder
  name                     = var.name
  vql                      = var.vql
}
