output "roles" {
  value = [
    tomap({
      database  = denodo_database_role.db_role_read.database_name,
      role_name = denodo_database_role.db_role_read.id
    }),
    tomap({
      database  = denodo_database_role.db_role_dev.database_name,
      role_name = denodo_database_role.db_role_dev.id
    })
  ]
}
