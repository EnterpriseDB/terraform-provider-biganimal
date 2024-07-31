package plan_modifier

import (
	"context"

	"github.com/EnterpriseDB/terraform-provider-biganimal/pkg/constants"
	"github.com/EnterpriseDB/terraform-provider-biganimal/pkg/utils"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

func CustomPhaseForUnknown() planmodifier.String {
	return customPhaseForUnknownModifier{}
}

// customStringForUnknownModifier implements the plan modifier.
type customPhaseForUnknownModifier struct{}

// Description returns a human-readable description of the plan modifier.
func (m customPhaseForUnknownModifier) Description(_ context.Context) string {
	return "Once set, the value of this attribute in state will not change."
}

// MarkdownDescription returns a markdown description of the plan modifier.
func (m customPhaseForUnknownModifier) MarkdownDescription(_ context.Context) string {
	return "Once set, the value of this attribute in state will not change."
}

// PlanModifyString implements the plan modification logic.
func (m customPhaseForUnknownModifier) PlanModifyString(ctx context.Context, req planmodifier.StringRequest, resp *planmodifier.StringResponse) {
	// always set phase
	var planObject map[string]tftypes.Value

	err := req.Plan.Raw.As(&planObject)
	if err != nil {
		resp.Diagnostics.AddError("Mapping plan object in custom phase plan modifier error", err.Error())
		return
	}

	// this is a create, if no state it should be unknown
	if req.StateValue.IsNull() {
		return
	}

	// this is update, if phase is waiting for tde key access then set to unknown
	if !req.PlanValue.IsUnknown() {
		if req.PlanValue.ValueString() == constants.PHASE_WAITING_FOR_ACCESS_TO_ENCRYPTION_KEY {
			resp.PlanValue = basetypes.NewStringUnknown()
			return
		}
	}

	var pause bool
	err = planObject["pause"].As(&pause)
	if err != nil {
		resp.Diagnostics.AddError("Mapping bool pause in custom phase plan modifier error", err.Error())
		return
	}

	if pause {
		resp.PlanValue = basetypes.NewStringPointerValue(utils.ToPointer(constants.PHASE_PAUSED))
		return
	} else {
		resp.PlanValue = basetypes.NewStringPointerValue(utils.ToPointer(constants.PHASE_HEALTHY))
		return
	}
}
