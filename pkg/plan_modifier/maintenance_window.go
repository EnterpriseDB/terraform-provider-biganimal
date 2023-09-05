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

		if strings.Replace(planAttr["is_enabled"].String(), "\"", "", -1) == "false" {
			startDayAttr := planAttr["start_day"]
			startTimeAttr := planAttr["start_time"]

			if strings.Replace(startDayAttr.String(), "\"", "", -1) != "0" && !startDayAttr.IsNull() && !startDayAttr.IsUnknown() {
				resp.Diagnostics.AddError("Maintenance window start_day cannot be set if is_enabled is false", fmt.Sprintf("maintenance window start_day has changed to %v, please set to 0 or remove",
					startDayAttr))
			}

			if strings.Replace(startTimeAttr.String(), "\"", "", -1) != "00:00" && !startTimeAttr.IsNull() && !startTimeAttr.IsUnknown() {
				resp.Diagnostics.AddError("Maintenance window start_time cannot be set if is_enabled is false", fmt.Sprintf("maintenance window start_time has changed to %v, please set to 00:00 or remove",
					startTimeAttr))
			}

			return
		}
	}

	// // Do nothing if there is no state value.
	// if req.StateValue.IsNull() {
	// 	return
	// }

	// // Do nothing if there is a known planned value.
	// if !req.PlanValue.IsUnknown() {
	// 	return
	// }

	// // Do nothing if there is an unknown configuration value, otherwise interpolation gets messed up.
	// if req.ConfigValue.IsUnknown() {
	// 	return
	// }

	// resp.PlanValue = req.StateValue
}
