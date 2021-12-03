# -----------------------------------------------------------------------------
# Creating Read Database Role
# -----------------------------------------------------------------------------

resource "denodo_database_role" "db_role_read" {
  database_name       = data.terraform_remote_state.vbd.outputs.id
  name                = format("%s_%s", data.terraform_remote_state.vbd.outputs.id, "read")
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
  database_name       = data.terraform_remote_state.vbd.outputs.id
  name                = format("%s_%s", data.terraform_remote_state.vbd.outputs.id, "dev")
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
