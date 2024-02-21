package plan_modifier

import (
	"context"
	"strings"

	"github.com/EnterpriseDB/terraform-provider-biganimal/pkg/models"
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

	var planObject map[string]tftypes.Value

	err := req.Plan.Raw.As(&planObject)
	if err != nil {
		resp.Diagnostics.AddError("Mapping plan object in custom phase plan modifier error", err.Error())
		return
	}

	var pause bool
	err = planObject["pause"].As(&pause)
	if err != nil {
		resp.Diagnostics.AddError("Mapping bool pause in custom phase plan modifier error", err.Error())
		return
	}

	if pause {
		resp.PlanValue = basetypes.NewStringPointerValue(utils.ToPointer(models.PHASE_PAUSED))
	} else {
		resp.PlanValue = basetypes.NewStringPointerValue(utils.ToPointer(models.PHASE_HEALTHY))
	}

	if !strings.Contains(resp.PlanValue.String(), models.PHASE_HEALTHY) && !strings.Contains(resp.PlanValue.String(), models.PHASE_PAUSED) {
		resp.Diagnostics.AddError("Cluster not in not ready for update operations", "Cluster not in healthy state for update operations please wait...")
		return
	}
}
