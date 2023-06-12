data "biganimal_projects" "this" {}

output "projects" {
  value = data.biganimal_projects.this.projects
}

output "number_of_projects" {
  value = length(data.biganimal_projects.this.projects)
}
