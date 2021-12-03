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
# Creating Folders in Database
# -----------------------------------------------------------------------------

resource "denodo_database_folder" "db_folder" {
  database    = var.database
  folder_path = var.folder_path
}
