package plan_modifier

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

func CustomConnection() planmodifier.String {
	return customConnectionModifier{}
}

// customConnectionModifier implements the plan modifier.
type customConnectionModifier struct{}

// Description returns a human-readable description of the plan modifier.
func (m customConnectionModifier) Description(_ context.Context) string {
	return "Once set, the value of this attribute in state will not change."
}

// MarkdownDescription returns a markdown description of the plan modifier.
func (m customConnectionModifier) MarkdownDescription(_ context.Context) string {
	return "Once set, the value of this attribute in state will not change."
}

// PlanModifyString implements the plan modification logic.
func (m customConnectionModifier) PlanModifyString(ctx context.Context, req planmodifier.StringRequest, resp *planmodifier.StringResponse) {
	var planObject map[string]tftypes.Value

	err := req.Plan.Raw.As(&planObject)
	if err != nil {
		resp.Diagnostics.AddError("Mapping plan object in allowed ip ranges plan modifier error", err.Error())
		return
	}

	var planPrivateNetworking bool
	err = planObject["private_networking"].As(&planPrivateNetworking)
	if err != nil {
		resp.Diagnostics.AddError("Mapping private networking object in allowed ip ranges plan modifier error", err.Error())
		return
	}

	var stateObject map[string]tftypes.Value

	err = req.State.Raw.As(&stateObject)
	if err != nil {
		resp.Diagnostics.AddError("Mapping state object in allowed ip ranges plan modifier error", err.Error())
		return
	}

	var statePrivateNetworking bool
	err = stateObject["private_networking"].As(&statePrivateNetworking)
	if err != nil {
		resp.Diagnostics.AddError("Mapping private networking object in allowed ip ranges plan modifier error", err.Error())
		return
	}

	// private networking has changed so connection uri will change
	if planPrivateNetworking != statePrivateNetworking {
		resp.PlanValue = basetypes.NewStringUnknown()
		return
	}

	// Do nothing if there is no state value.
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
