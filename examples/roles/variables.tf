variable "denodo_database" {
  type = string
}

variable "denodo_host" {
  type = string
}

variable "denodo_password" {
  type = string
}

variable "denodo_port" {
  type = number
}

variable "denodo_username" {
  type = string
}

variable "denodo_database_authentication" {
  type    = string
  default = "LOCAL"
}

variable "denodo_database_char_set" {
  type    = string
  default = "DEFAULT"
}

variable "denodo_database_cost_optimization" {
  type    = string
  default = "DEFAULT"
}

variable "denodo_database_description" {
  type    = string
  default = "testing example"
}

variable "denodo_database_name" {
  type    = string
  default = "test_database"
}

variable "denodo_database_summary_rewrite" {
  type    = string
  default = "DEFAULT"
}

variable "denodo_database_query_simplification" {
  type    = string
  default = "DEFAULT"
}

variable "denodo_dev_user_password" {
  type      = string
  sensitive = true
}

variable "denodo_read_user_password" {
  type      = string
  sensitive = true
}

variable "data_source_catalog_name" {
  type = string
}

variable "data_source_class_path" {
  type = string
}

variable "data_source_database_type" {
  type = string
}

variable "data_source_database_version" {
  type = string
}

variable "data_source_database_uri" {
  type = string
}

variable "data_source_driver_class_name" {
  type = string
}

variable "data_source_name" {
  type    = string
  default = "test_data_source"
}

variable "data_source_password" {
  type      = string
  sensitive = true
}

variable "data_source_schema_name" {
  type = string
}

variable "data_source_username" {
  type      = string
  sensitive = true
}
