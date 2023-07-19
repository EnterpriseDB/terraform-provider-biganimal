package plan_modifier

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func CustomPGConfig() planmodifier.List {
	return customPGConfigModifier{}
}

// customStringUnknownModifier implements the plan modifier.
type customPGConfigModifier struct{}

// Description returns a human-readable description of the plan modifier.
func (m customPGConfigModifier) Description(_ context.Context) string {
	return "Once set, the value of this attribute in state will not change."
}

// MarkdownDescription returns a markdown description of the plan modifier.
func (m customPGConfigModifier) MarkdownDescription(_ context.Context) string {
	return "Once set, the value of this attribute in state will not change."
}

// PlanModifyList implements the plan modification logic.
func (m customPGConfigModifier) PlanModifyList(ctx context.Context, req planmodifier.ListRequest, resp *planmodifier.ListResponse) {
	if req.StateValue.IsNull() {
		return
	}

	configValue := req.ConfigValue.Elements()

	// sort state the same as config
	stateValue := req.StateValue.Elements()

	newState := []attr.Value{}
	for _, v := range configValue {
		if pgConfigContains(stateValue, v) {
			newState = append(newState, v)
		}
	}

	req.StateValue = basetypes.NewListValueMust(newState[0].Type(ctx), newState)

	// sort plan the same as config
	planValue := req.PlanValue.Elements()

	newPlan := []attr.Value{}
	for _, v := range configValue {
		if pgConfigContains(planValue, v) {
			newPlan = append(newPlan, v)
		}
	}

	resp.PlanValue = basetypes.NewListValueMust(newPlan[0].Type(ctx), newPlan)

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

func pgConfigContains(s []attr.Value, e attr.Value) bool {
	for _, a := range s {
		aName := a.(basetypes.ObjectValue).Attributes()["name"].String()
		eName := e.(basetypes.ObjectValue).Attributes()["name"].String()
		if aName == eName {
			return true
		}
	}
	return false
}
