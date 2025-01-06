terraform {
  required_providers {
    biganimal = {
      source  = "EnterpriseDB/biganimal"
      version = "1.2.1"
    }
    random = {
      source  = "hashicorp/random"
      version = "3.6.0"
    }
  }
}

resource "random_pet" "project_name" {
  separator = " "
}

resource "biganimal_project" "this" {
  project_name = format("TF %s", title(random_pet.project_name.id))
  #tags = [
  #  {
  #     tag_name  = "<ex_tag_name_1>"
  #     color = "blue"
  #  },
  #  {
  #     tag_name  = "<ex_tag_name_2>"
  #  },
  #]
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
