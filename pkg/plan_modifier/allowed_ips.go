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
	if req.StateValue.IsNull() {
		return
	}

	if !req.StateValue.IsNull() {
		configValue := req.ConfigValue.Elements()

		// sort state the same as config
		stateValue := req.StateValue.Elements()

		newState := []attr.Value{}
		for _, v := range configValue {
			if allowedIpContains(stateValue, v) {
				newState = append(newState, v)
			}
		}

		req.StateValue = basetypes.NewListValueMust(newState[0].Type(ctx), newState)

		// sort plan the same as config
		planValue := resp.PlanValue.Elements()

		newPlan := []attr.Value{}
		for _, v := range configValue {
			if allowedIpContains(planValue, v) {
				newPlan = append(newPlan, v)
			}
		}

		resp.PlanValue = basetypes.NewListValueMust(newPlan[0].Type(ctx), newPlan)

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

func allowedIpContains(s []attr.Value, e attr.Value) bool {
	for _, a := range s {
		aCidr := a.(basetypes.ObjectValue).Attributes()["cidr_block"].String()
		eCidr := e.(basetypes.ObjectValue).Attributes()["cidr_block"].String()
		aDesc := a.(basetypes.ObjectValue).Attributes()["description"].String()
		eDesc := e.(basetypes.ObjectValue).Attributes()["description"].String()
		if aCidr == eCidr && aDesc == eDesc {
			return true
		}
	}
	return false
}
