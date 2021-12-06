output "object_list" {
  value = data.denodo_jdbc_data_source_object.data_bv.objects.*
}

output "data_source_database" {
  value = module.data_source.data_source_database
}

output "data_source_name" {
  value = module.data_source.data_source_name
}