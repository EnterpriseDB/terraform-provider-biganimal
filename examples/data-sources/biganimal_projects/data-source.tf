data "biganimal_projects" "this" {
  query = var.query
}

output "projects" {
  value = data.biganimal_projects.this.projects
}

output "number_of_projects" {
  value = length(data.biganimal_projects.this.projects)
}

variable "query" {
  type        = string
  description = "Query string for the projects"
  default     = ""
}
