package terraform

import "github.com/hashicorp/terraform-plugin-framework/types"

type Condition struct {
	ConditionStatus types.String `tfsdk:"condition_status"`
	Type_           types.String `tfsdk:"type"`
}
