data "biganimal_projects" "this" {}

output "projects" {
  value = data.biganimal_projects.this.projects
}
