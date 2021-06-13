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
