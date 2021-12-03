# -----------------------------------------------------------------------------
# Creating Base Views From JDBC Data Source
# -----------------------------------------------------------------------------

resource "denodo_base_view" "db_bv" {
  count                    = length(data.terraform_remote_state.ds.outputs.object_list.*)
  data_source_catalog_name = var.data_source_catalog_name
  data_source_database     = data.terraform_remote_state.ds.outputs.data_source_database
  data_source_name         = data.terraform_remote_state.ds.outputs.data_source_name
  data_source_schema_name  = var.data_source_schema_name
  data_source_table_name   = data.terraform_remote_state.ds.outputs.object_list[count.index].object_name
  database                 = data.terraform_remote_state.folder.outputs.database[0]
  folder                   = data.terraform_remote_state.folder.outputs.folder[0]
  name                     = format("bv_%s", data.terraform_remote_state.ds.outputs.object_list[count.index].object_name)
}
