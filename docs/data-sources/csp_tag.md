---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "biganimal_csp_tag Data Source - terraform-provider-biganimal"
subcategory: ""
description: |-
  CSP Tags will enable users to categorize and organize resources across types and improve the efficiency of resource retrieval
---

# biganimal_csp_tag (Data Source)

CSP Tags will enable users to categorize and organize resources across types and improve the efficiency of resource retrieval

## Example Usage

```terraform
terraform {
  required_providers {
    biganimal = {
      source  = "EnterpriseDB/biganimal"
      version = "1.1.1"
    }
  }
}

data "biganimal_csp_tag" "this" {
  project_id        = "<ex-project-id>"        # ex: "prj_12345"
  cloud_provider_id = "<ex-cloud-provider-id>" # ex: "aws"
}

output "csp_tags" {
  value = data.biganimal_csp_tag.this.csp_tags
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `cloud_provider_id` (String)
- `project_id` (String)

### Optional

- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `add_tags` (Attributes List) (see [below for nested schema](#nestedatt--add_tags))
- `csp_tags` (Attributes List) CSP Tags on cluster (see [below for nested schema](#nestedatt--csp_tags))
- `delete_tags` (List of String)
- `edit_tags` (Attributes List) (see [below for nested schema](#nestedatt--edit_tags))
- `id` (String) The ID of this resource.

<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m" (minutes), "h" (hours).
- `delete` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m" (minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are saved into state before the destroy operation occurs.
- `update` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m" (minutes), "h" (hours).


<a id="nestedatt--add_tags"></a>
### Nested Schema for `add_tags`

Read-Only:

- `csp_tag_key` (String)
- `csp_tag_value` (String)


<a id="nestedatt--csp_tags"></a>
### Nested Schema for `csp_tags`

Read-Only:

- `csp_tag_id` (String)
- `csp_tag_key` (String)
- `csp_tag_value` (String)
- `status` (String)


<a id="nestedatt--edit_tags"></a>
### Nested Schema for `edit_tags`

Read-Only:

- `csp_tag_id` (String)
- `csp_tag_key` (String)
- `csp_tag_value` (String)
- `status` (String)