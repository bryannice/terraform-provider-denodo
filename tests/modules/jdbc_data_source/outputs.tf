output "data_source_database" {
  value = denodo_jdbc_data_source.db_ds.denodo_database
}

output "data_source_name" {
  value = denodo_jdbc_data_source.db_ds.id
}
