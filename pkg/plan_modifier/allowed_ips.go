package plan_modifier

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
)

func CustomAllowedIps() planmodifier.Set {
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
func (m customAllowedIpsModifier) PlanModifySet(ctx context.Context, req planmodifier.SetRequest, resp *planmodifier.SetResponse) {
	// if len(resp.PlanValue.Elements()) == 0 {
	// 	resp.PlanValue = types.SetNull(resp.PlanValue.ElementType(ctx))
	// 	return
	// }

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
