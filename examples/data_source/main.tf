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
# Fetch List of Objects in JDBC Data Source
# -----------------------------------------------------------------------------

data "denodo_jdbc_data_source_table" "jdst" {
  catalog_name = var.data_source_catalog_name
  database     = var.data_source_database
  name         = var.data_source_name
  schema_name  = var.data_source_schema_name
}
