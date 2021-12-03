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
# Get data from yml config
# -----------------------------------------------------------------------------

locals {
  data = yamldecode(file("tf_vars.yml"))
}


# -----------------------------------------------------------------------------
# Setting the Denodo Provider
# -----------------------------------------------------------------------------

provider "denodo" {
  host     = local.data.host
  password = local.data.password
  port     = local.data.port
  ssl_mode = "disable"
  username = local.data.username
}

# -----------------------------------------------------------------------------
# Setting the Create Virtual Database Northwind
# -----------------------------------------------------------------------------

module "database" {
  source               = "./modules/virtual_database"
  authentication       = local.data.database.authentication
  char_set             = local.data.database.char_set
  cost_optimization    = local.data.database.cost_optimization
  description          = local.data.database.description
  name                 = local.data.database.name
  query_simplification = local.data.database.query_simplification
  summary_rewrite      = local.data.database.summary_rewrite
}

# -----------------------------------------------------------------------------
# Creating Folders
# -----------------------------------------------------------------------------

module "folders" {
  count       = length(local.data.folders.*)
  database    = module.database.id
  folder_path = local.data.folders[count.index]
  source      = "./modules/folders"
}

# -----------------------------------------------------------------------------
# Creating Data Source to Northwind
# -----------------------------------------------------------------------------

module "data_source" {
  data_source_class_path        = local.data.data_source.class_path
  data_source_database_type     = local.data.data_source.data_source_database_type
  data_source_database_version  = local.data.data_source.data_source_database_version
  data_source_database_uri      = local.data.data_source.database_uri
  data_source_database          = module.database.id
  data_source_driver_class_name = local.data.data_source.driver_class_name
  data_source_folder            = module.folders[0].folder
  data_source_name              = local.data.data_source.name
  data_source_password          = local.data.data_source.password
  data_source_username          = local.data.data_source.username
  source                        = "./modules/jdbc_data_source"
}