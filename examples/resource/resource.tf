terraform {
  required_providers {
    biganimal = {
      source  = "biganimal"
      version = "0.3.1"
    }
    random = {
      source  = "hashicorp/random"
      version = "3.3.1"
    }
  }
}

resource "random_password" "password" {
  length           = 16
  special          = true
  override_special = "!#$%&*()-_=+[]{}<>:?"
}

# resource "biganimal_cluster" "this_resource" {
#   cluster_name = "tf-test-4"

#   allowed_ip_ranges {
#     cidr_block = "127.0.0.1/32"
#     description = "localhost"
#   }

#   allowed_ip_ranges {
#     cidr_block = "192.168.0.1/32"
#     description = "description!"
#   }

#   backup_retention_period = "6d"
#   cluster_architecture {
#     id = "single"
#     nodes = 1
#   }

#   instance_type_id = "azure:Standard_D2s_v3"
#   password = resource.random_password.password.result
#   pg_config {
#     name  = "application_name"
#     value = "created through terraform"
#   }

#   pg_config {
#     name  = "array_nulls"
#     value = "off"
#   }

#   storage {
#     volume_type = "azurepremiumstorage"
#     volume_properties = "P1"
#     size = "4 Gi"
#   }

#   pg_type = "epas"
#   pg_version = "14"
#   private_networking = false
#   cloud_provider = "azure"
#   region = "australiaeast"
#   replicas = 1
# }