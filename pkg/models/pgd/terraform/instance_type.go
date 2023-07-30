package terraform

import "github.com/hashicorp/terraform-plugin-framework/types"

type InstanceType struct {
	InstanceTypeId types.String `tfsdk:"instance_type_id"`
}
