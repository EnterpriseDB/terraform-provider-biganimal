package terraform

import "github.com/hashicorp/terraform-plugin-framework/types"

type Tag struct {
	TagId   types.String `tfsdk:"tag_id"`
	TagName types.String `tfsdk:"tag_name"`
	Color   types.String `tfsdk:"color"`
}
