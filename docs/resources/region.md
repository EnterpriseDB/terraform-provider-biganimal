---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "biganimal_region Resource - terraform-provider-biganimal"
subcategory: ""
description: |-
  Manage a region
---

# biganimal_region (Resource)

Manage a region



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `cloud_provider` (String) Cloud Provider
- `region_id` (String) Region ID

### Optional

- `status` (String) Region Status
- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `continent` (String) Continent
- `id` (String) The ID of this resource.
- `name` (String) Region Name

<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String)
- `delete` (String)
- `update` (String)