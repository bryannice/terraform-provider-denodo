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
# Creating Roles to Northwind
# -----------------------------------------------------------------------------

/*
module "roles" {
  count               = length(local.data.roles.*)
  database_name       = module.database.id
  name                = local.data.roles[count.index].name
  admin               = local.data.roles[count.index].admin
  connect             = local.data.roles[count.index].connect
  create              = local.data.roles[count.index].create
  create_data_service = local.data.roles[count.index].create_data_service
  create_data_source  = local.data.roles[count.index].create_data_source
  create_folder       = local.data.roles[count.index].create_folder
  create_view         = local.data.roles[count.index].create_view
  execute             = local.data.roles[count.index].execute
  file                = local.data.roles[count.index].file
  grant               = local.data.roles[count.index].grant
  meta_data           = local.data.roles[count.index].meta_data
  monitor_admin       = local.data.roles[count.index].monitor_admin
  write               = local.data.roles[count.index].write
  source              = "./modules/roles"
}
*/

# -----------------------------------------------------------------------------
# Creating Users to Northwind
# -----------------------------------------------------------------------------

/*
module "users" {
  count       = length(local.data.users.*)
  admin       = local.data.users[count.index].admin
  description = local.data.users[count.index].description
  password    = local.data.users[count.index].password
  roles       = module.roles[count.index].role
  username    = local.data.users[count.index].name
  source      = "./modules/users"
}
*/

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

# -----------------------------------------------------------------------------
# Fetching Table List From JDBC Data Source
# -----------------------------------------------------------------------------

data "denodo_jdbc_data_source_object" "data_bv" {
  catalog_name = local.data.base_views.catalog_name
  database     = module.database.id
  name         = module.data_source.data_source_name
  schema_name  = local.data.base_views.schema_name
}

# -----------------------------------------------------------------------------
# Reading from Remote State File from Data Source
# -----------------------------------------------------------------------------
data "terraform_remote_state" "ds" {
  backend = "local"
  config = {
    path = "./terraform.tfstate"
  }
}

# -----------------------------------------------------------------------------
# Creating Base Views in Northwind
# -----------------------------------------------------------------------------

module "base_views" {
  count                    = length(data.terraform_remote_state.ds.outputs.object_list.*)
  data_source_catalog_name = local.data.base_views.catalog_name
  data_source_database     = data.terraform_remote_state.ds.outputs.data_source_database
  data_source_name         = data.terraform_remote_state.ds.outputs.data_source_name
  data_source_schema_name  = data.terraform_remote_state.ds.outputs.object_list[count.index].schema_name
  data_source_table_name   = data.terraform_remote_state.ds.outputs.object_list[count.index].object_name
  database                 = data.terraform_remote_state.ds.outputs.data_source_database
  folder                   = local.data.folders[1]
  name                     = format("bv_%s", data.terraform_remote_state.ds.outputs.object_list[count.index].object_name)
  source                   = "./modules/base_views"
}

module "custom_base_views" {
  database = data.terraform_remote_state.ds.outputs.data_source_database
  name     = "bv_order_details_orders"
  source   = "./modules/base_views"
  vql      = file("./sql/base_view/bv_order_details_orders.sql")
}

# -----------------------------------------------------------------------------
# Creating Derived Views in Northwind
# -----------------------------------------------------------------------------

module "derived_views" {
  for_each = fileset("./sql/derived_view", "*.sql")
  database = module.custom_base_views.database
  name     = replace(each.value, ".sql", "")
  source   = "./modules/derived_views"
  vql      = file(format("./sql/derived_view/%s", each.value))
}
