package plan_modifier

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
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

	if !strings.Contains(resp.PlanValue.String(), "Cluster in healthy state") {
		resp.Diagnostics.AddError("Cluster not in not ready for update operations", "Cluster not in healthy state for update operations please wait...")
		return
	}
}