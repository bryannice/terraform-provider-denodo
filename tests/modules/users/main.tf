# -----------------------------------------------------------------------------
# Creating Read User Account
# -----------------------------------------------------------------------------

resource "denodo_user" "db_usr_read" {
  admin       = false
  description = "test read user"
  password    = var.denodo_read_user_password
  roles       = data.terraform_remote_state.roles.outputs.roles[0].role_name
  username    = "test_read_user"
}

# -----------------------------------------------------------------------------
# Creating Read Developer Account
# -----------------------------------------------------------------------------

resource "denodo_user" "db_usr_dev" {
  admin       = false
  description = "test dev user"
  password    = var.denodo_dev_user_password
  roles       = data.terraform_remote_state.roles.outputs.roles[1].role_name
  username    = "test_dev_user"
}
