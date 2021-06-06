terraform {
  required_providers {
    denodo = {
      version = "0.1"
      source  = "custom.com/bryannice/denodo"
    }
  }
}

provider "denodo" {
  database = var.database
  host = var.host
  password = var.password
  port = var.port
  ssl_mode = var.ssl_mode
  username = var.username
}

data "denodo_jdbc_data_source_table" "jdst" {
  catalog_name = var.data_source_catalog_name
  database = var.data_source_database
  name = var.data_source_name
  schema_name = var.data_source_schema_name
}