variable "data_source_class_path" {
  type = string
}

variable "data_source_database" {
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

variable "data_source_folder" {
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

variable "data_source_username" {
  type      = string
  sensitive = true
}
