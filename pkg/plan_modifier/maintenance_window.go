package plan_modifier

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
)

func MaintenanceWindowForUnknown() planmodifier.Object {
	return MaintenanceWindowForUnknownModifier{}
}

// MaintenanceWindowForUnknownModifier implements the plan modifier.
type MaintenanceWindowForUnknownModifier struct{}

// Description returns a human-readable description of the plan modifier.
func (m MaintenanceWindowForUnknownModifier) Description(_ context.Context) string {
	return "Once set, the value of this attribute in state will not change."
}

// MarkdownDescription returns a markdown description of the plan modifier.
func (m MaintenanceWindowForUnknownModifier) MarkdownDescription(_ context.Context) string {
	return "Once set, the value of this attribute in state will not change."
}

// PlanModifyObject implements the plan modification logic.
func (m MaintenanceWindowForUnknownModifier) PlanModifyObject(ctx context.Context, req planmodifier.ObjectRequest, resp *planmodifier.ObjectResponse) {
	if !req.PlanValue.IsUnknown() && !req.StateValue.IsNull() {
		planAttr := req.PlanValue.Attributes()
		stateAttr := req.StateValue.Attributes()

		if strings.Replace(planAttr["is_enabled"].String(), "\"", "", -1) == "false" &&
			planAttr["start_day"] != stateAttr["start_day"] {
			resp.Diagnostics.AddWarning("Maintenance window is_enabled is false but you have changed start day", fmt.Sprintf("maintenance window start day has changed from %v to %v",
				stateAttr["start_day"],
				planAttr["start_day"]))
		}

		if strings.Replace(planAttr["is_enabled"].String(), "\"", "", -1) == "false" &&
			planAttr["start_time"] != stateAttr["start_time"] {
			resp.Diagnostics.AddWarning("Maintenance window is_enabled is false but you have changed start time", fmt.Sprintf("maintenance window start time has changed from %v to %v",
				stateAttr["start_time"],
				planAttr["start_time"]))
		}

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
