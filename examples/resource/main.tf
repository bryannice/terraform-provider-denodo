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

# -----------------------------------------------------------------------------
# Creating Data Source Folder in Database
# -----------------------------------------------------------------------------

resource "denodo_database_folder" "db_folder_ds" {
  database    = denodo_database.db.id
  folder_path = "/data_source"
}

# -----------------------------------------------------------------------------
# Creating Base View Folder in Database
# -----------------------------------------------------------------------------

resource "denodo_database_folder" "db_folder_bv" {
  database    = denodo_database.db.id
  folder_path = "/base_view"
}

# -----------------------------------------------------------------------------
# Creating JDBC Data Source
# -----------------------------------------------------------------------------

resource "denodo_jdbc_data_source" "db_ds" {
  class_path                   = var.data_source_class_path
  data_source_database_type    = var.data_source_database_type
  data_source_database_version = var.data_source_database_version
  database_uri                 = var.data_source_database_uri
  denodo_database              = denodo_database.db.id
  driver_class_name            = var.data_source_driver_class_name
  folder                       = denodo_database_folder.db_folder_ds.id
  name                         = var.data_source_name
  password                     = var.data_source_password
  username                     = var.data_source_username
}

# -----------------------------------------------------------------------------
# Fetch List of Objects in JDBC Data Source
# -----------------------------------------------------------------------------

data "denodo_jdbc_data_source_table" "jdst" {
  catalog_name = var.data_source_catalog_name
  database     = denodo_database.db.id
  name         = denodo_jdbc_data_source.db_ds.id
  schema_name  = var.data_source_schema_name
}

# -----------------------------------------------------------------------------
# Creating Base Views From JDBC Data Source
# -----------------------------------------------------------------------------

resource "denodo_base_view" "db_bv" {
  data_source_catalog_name = var.data_source_catalog_name
  data_source_database     = denodo_database.db.id
  data_source_name         = denodo_jdbc_data_source.db_ds.id
  data_source_schema_name  = var.data_source_schema_name
  data_source_table_name   = data.denodo_jdbc_data_source_table.jdst.tables[0].object_name
  database                 = denodo_database.db.id
  folder                   = denodo_database_folder.db_folder_bv.id
  name                     = format("bv_%s", data.denodo_jdbc_data_source_table.jdst.tables[0].object_name)
}

# -----------------------------------------------------------------------------
# Creating Read Database Role
# -----------------------------------------------------------------------------

resource "denodo_database_role" "db_role_read" {
  database_name       = denodo_database.db.id
  name                = format("%s_%s", denodo_database.db.id, "read")
  admin               = false
  connect             = true
  create              = false
  create_data_service = false
  create_data_source  = false
  create_folder       = false
  create_view         = false
  execute             = true
  file                = false
  grant               = false
  meta_data           = true
  monitor_admin       = false
  write               = false
}

# -----------------------------------------------------------------------------
# Creating Developer Database Role
# -----------------------------------------------------------------------------

resource "denodo_database_role" "db_role_dev" {
  database_name       = denodo_database.db.id
  name                = format("%s_%s", denodo_database.db.id, "dev")
  admin               = false
  connect             = true
  create              = false
  create_data_service = false
  create_data_source  = true
  create_folder       = true
  create_view         = true
  execute             = true
  file                = false
  grant               = false
  meta_data           = true
  monitor_admin       = true
  write               = true
}

# -----------------------------------------------------------------------------
# Creating Read User Account
# -----------------------------------------------------------------------------

resource "denodo_user" "db_usr_read" {
  admin       = false
  description = "test read user"
  password    = var.denodo_read_user_password
  roles       = denodo_database_role.db_role_read.id
  username    = "test_read_user"
}

# -----------------------------------------------------------------------------
# Creating Read Developer Account
# -----------------------------------------------------------------------------

resource "denodo_user" "db_usr_dev" {
  admin       = false
  description = "test dev user"
  password    = var.denodo_dev_user_password
  roles       = denodo_database_role.db_role_dev.id
  username    = "test_dev_user"
}
