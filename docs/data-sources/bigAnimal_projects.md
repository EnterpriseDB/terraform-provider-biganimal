---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "bigAnimal_projects Data Source - terraform-provider-biganimal"
subcategory: ""
description: |-
  The projects data source shows the BigAnimal Projects.
---

# bigAnimal_projects (Data Source)

The projects data source shows the BigAnimal Projects.

## Example Usage

```terraform
data "biganimal_projects" "this" {}

output "projects" {
  value = data.biganimal_projects.this.projects
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `query` (String) Query to filter project list.

### Read-Only

- `projects` (Attributes Set) List of the organization's projects. (see [below for nested schema](#nestedatt--projects))

<a id="nestedatt--projects"></a>
### Nested Schema for `projects`

Required:

- `name` (String) Project Name of the project.

Optional:

- `cloud_providers` (Attributes Set) Enabled Cloud Providers. (see [below for nested schema](#nestedatt--projects--cloud_providers))

Read-Only:

- `cluster_count` (Number) User Count of the project.
- `project_id` (String) Project ID of the project.
- `user_count` (Number) User Count of the project.

<a id="nestedatt--projects--cloud_providers"></a>
### Nested Schema for `projects.cloud_providers`

Read-Only:

- `cloud_provider_id` (String) Cloud Provider ID.
- `cloud_provider_name` (String) Cloud Provider Name.