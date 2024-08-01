package plan_modifier

import (
	"context"

	"github.com/EnterpriseDB/terraform-provider-biganimal/pkg/constants"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

func CustomTDEAction() planmodifier.String {
	return CustomTDEActionModifier{}
}

// customStringForUnknownModifier implements the plan modifier.
type CustomTDEActionModifier struct{}

// Description returns a human-readable description of the plan modifier.
func (m CustomTDEActionModifier) Description(_ context.Context) string {
	return "Once set, the value of this attribute in state will not change."
}

// MarkdownDescription returns a markdown description of the plan modifier.
func (m CustomTDEActionModifier) MarkdownDescription(_ context.Context) string {
	return "Once set, the value of this attribute in state will not change."
}

// PlanModifyString implements the plan modification logic.
func (m CustomTDEActionModifier) PlanModifyString(ctx context.Context, req planmodifier.StringRequest, resp *planmodifier.StringResponse) {
	var planObject map[string]tftypes.Value
	var stateObject map[string]tftypes.Value

	err := req.Plan.Raw.As(&planObject)
	if err != nil {
		resp.Diagnostics.AddError("Mapping plan object in custom phase plan modifier error", err.Error())
		return
	}

	err = req.State.Raw.As(&stateObject)
	if err != nil {
		resp.Diagnostics.AddError("Mapping state object in custom phase plan modifier error", err.Error())
		return
	}

	var tde map[string]tftypes.Value
	err = planObject["transparent_data_encryption"].As(&tde)
	if err != nil {
		resp.Diagnostics.AddError("Mapping transparent data encryption in custom phase plan modifier error", err.Error())
		return
	}

	var phase string
	err = stateObject["phase"].As(&phase)
	if err != nil {
		resp.Diagnostics.AddError("Mapping phase in custom phase plan modifier error", err.Error())
		return
	}

	// this is on create. If state phase is empty and plan tde key_id is set then display tde action check warning
	if phase == "" {
		if tde["key_id"].String() != "" {
			resp.Diagnostics.AddWarning("Transparent data encryption info", constants.TDE_CHECK_ACTION)
		}
		return
	}

	// this is on update. If state phase is healthy then tde action should be null
	if phase == constants.PHASE_HEALTHY {
		resp.PlanValue = basetypes.NewStringNull()
		return
	}

	// this is on update. If state phase is waiting for tde key access then display tde action check warning
	if phase == constants.PHASE_WAITING_FOR_ACCESS_TO_ENCRYPTION_KEY {
		resp.Diagnostics.AddWarning("Transparent data encryption info", constants.TDE_CHECK_ACTION)
		return
	}

	if !req.PlanValue.IsUnknown() {
		return
	}

	// Do nothing if there is an unknown configuration value, otherwise interpolation gets messed up.
	if req.ConfigValue.IsUnknown() {
		return
	}

	resp.PlanValue = req.StateValue
}
