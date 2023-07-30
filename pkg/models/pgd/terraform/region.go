package terraform

import "github.com/hashicorp/terraform-plugin-framework/types"

type Region struct {
	RegionId types.String `tfsdk:"region_id"`
}
