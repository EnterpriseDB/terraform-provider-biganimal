package plan_modifier

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func CustomAllowedIps() planmodifier.List {
	return customAllowedIpsModifier{}
}

// customAllowedIpsModifier implements the plan modifier.
type customAllowedIpsModifier struct{}

// Description returns a human-readable description of the plan modifier.
func (m customAllowedIpsModifier) Description(_ context.Context) string {
	return "Once set, the value of this attribute in state will not change."
}

// MarkdownDescription returns a markdown description of the plan modifier.
func (m customAllowedIpsModifier) MarkdownDescription(_ context.Context) string {
	return "Once set, the value of this attribute in state will not change."
}

// PlanModifyList implements the plan modification logic.
func (m customAllowedIpsModifier) PlanModifyList(ctx context.Context, req planmodifier.ListRequest, resp *planmodifier.ListResponse) {
	if len(resp.PlanValue.Elements()) == 0 {
		// if plan value is [] the api will return 0.0.0.0/0
		defaultAttrs := map[string]attr.Value{"cidr_block": basetypes.NewStringValue("0.0.0.0/0"), "description": basetypes.NewStringValue("")}
		defaultAttrTypes := map[string]attr.Type{"cidr_block": defaultAttrs["cidr_block"].Type(ctx), "description": defaultAttrs["description"].Type(ctx)}

		defaultObjectValue := basetypes.NewObjectValueMust(defaultAttrTypes, defaultAttrs)
		listOfObjects := []attr.Value{}
		listOfObjects = append(listOfObjects, defaultObjectValue)
		listValue := basetypes.NewListValueMust(defaultObjectValue.Type(ctx), listOfObjects)
		resp.PlanValue = listValue
		return
	}

	if req.StateValue.IsNull() {
		return
	}

	// Do nothing if there is a known planned value.
	if !req.PlanValue.IsUnknown() {
		return
	}

	// Do nothing if there is an unknown configuration value, otherwise interpolation gets messed up.
	if req.ConfigValue.IsUnknown() {
		return
	}

	resp.PlanValue = req.StateValue
}
