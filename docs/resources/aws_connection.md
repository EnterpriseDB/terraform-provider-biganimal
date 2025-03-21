---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "biganimal_aws_connection Resource - terraform-provider-biganimal"
subcategory: ""
description: |-
  The awsconnection resource is used to make connection between your project and AWS.
  o obtain the necessary input parameters, please refer to [Connect CSP](https://www.enterprisedb.com/docs/biganimal/latest/gettingstarted/02connectingtoyourcloud/connecting_aws/).
---

# biganimal_aws_connection (Resource)

The aws_connection resource is used to make connection between your project and AWS.
o obtain the necessary input parameters, please refer to [Connect CSP](https://www.enterprisedb.com/docs/biganimal/latest/getting_started/02_connecting_to_your_cloud/connecting_aws/).

## Example Usage

```terraform
terraform {
  required_providers {
    biganimal = {
      source  = "EnterpriseDB/biganimal"
      version = "2.0.0"
    }
  }
}


variable "external_id" {
  type = string
}
variable "project_id" {
  type = string
}
variable "role_arn" {
  type = string
}


resource "biganimal_aws_connection" "project_aws_conn" {
  project_id  = var.project_id
  role_arn    = var.role_arn
  external_id = var.external_id
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `external_id` (String) The AWS external ID provided by BigAnimal.
- `project_id` (String) Project ID of the project.
- `role_arn` (String) The AWS IAM role used by BigAnimal.

### Optional

- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String)
- `delete` (String)
- `update` (String)
