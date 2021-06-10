output "data" {
  value = [
    tomap({
      database = denodo_database_folder.db_folder_ds.database,
      folder   = denodo_database_folder.db_folder_ds.id
    }),
    tomap({
      database = denodo_database_folder.db_folder_bv.database,
      folder   = denodo_database_folder.db_folder_bv.id
    })
  ]
}
