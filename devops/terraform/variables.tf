variable "location" {
  type    = string
  default = "Central US"
}

variable "database_server" {
  type    = string
  default = "database-habits"
}

variable "database_resource_group" {
  type    = string
  default = null
}

variable "lets_encrypt_azure" {
  type    = string
}