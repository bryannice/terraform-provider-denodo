# -----------------------------------------------------------------------------
# Create Dervived View
# -----------------------------------------------------------------------------

resource "denodo_dervived_view" "dv" {
  database  = data.terraform_remote_state.folder.outputs.database[2]
  directory = "test_files"
}
