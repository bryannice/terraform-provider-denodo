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
# Creating Database
# -----------------------------------------------------------------------------

resource "denodo_database" "db" {
  authentication       = var.authentication
  char_set             = var.char_set
  cost_optimization    = var.cost_optimization
  description          = var.description
  name                 = var.name
  summary_rewrite      = var.summary_rewrite
  query_simplification = var.query_simplification
}
