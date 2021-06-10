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
# Creating Database
# -----------------------------------------------------------------------------

resource "denodo_database" "db" {
  authentication       = var.denodo_database_authentication
  char_set             = var.denodo_database_char_set
  cost_optimization    = var.denodo_database_cost_optimization
  description          = var.denodo_database_description
  name                 = var.denodo_database_name
  summary_rewrite      = var.denodo_database_summary_rewrite
  query_simplification = var.denodo_database_query_simplification
}
