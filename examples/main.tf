terraform {
  required_providers {
    denodo = {
      version = "0.1"
      source  = "custom.com/bryannice/denodo"
    }
  }
}

provider "denodo" {
  database = var.denodo_database
  host = var.denodo_host
  password = var.denodo_password
  port = var.denodo_port
  ssl_mode = var.denodo_ssl_mode
  username = var.denodo_username
}

resource "denodo_database" "db" {
  authentication = var.denodo_database_authentication
  char_set = var.denodo_database_char_set
  cost_optimization = var.denodo_database_cost_optimization
  description = var.denodo_database_description
  name = var.denodo_database_name
  summary_rewrite = var.summary_rewrite
  query_simplification = var.denodo_database_query_simplification
}

resource "denodo_database_role" "db_role" {

}

data "denodo_jdbc_data_source_table" "jdst" {
  catalog_name = var.data_source_catalog_name
  database = var.data_source_database
  name = var.data_source_name
  schema_name = var.data_source_schema_name
}