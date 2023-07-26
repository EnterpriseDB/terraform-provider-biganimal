terraform {
  required_providers {
    biganimal = {
      source  = "EnterpriseDB/biganimal"
      version = "0.5.1"
    }
    random = {
      source  = "hashicorp/random"
      version = "3.5.1"
    }
  }
}

resource "random_pet" "project_name" {
  separator = " "
}

resource "biganimal_project" "this" {
  project_name = format("TF %s", title(random_pet.project_name.id))
}

output "project_name" {
  value = resource.biganimal_project.this.project_name
}

output "project_id" {
  value = resource.biganimal_project.this.id
}

output "project" {
  value = resource.biganimal_project.this
}
