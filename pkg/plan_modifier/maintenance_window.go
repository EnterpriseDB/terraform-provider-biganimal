package plan_modifier

import (
	"context"

	"github.com/EnterpriseDB/terraform-provider-biganimal/pkg/models"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
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
	if !req.PlanValue.IsUnknown() {
		var planObject map[string]tftypes.Value

		err := req.Plan.Raw.As(&planObject)
		if err != nil {
			resp.Diagnostics.AddError("Mapping plan object in custom maintenance window plan modifier error", err.Error())
			return
		}

		mwOb := models.MaintenanceWindow{}
		err = planObject["maintenance_window"].As(&mwOb)
		if err != nil {
			resp.Diagnostics.AddError("Mapping maintenance window object in maintenance window plan modifier error", err.Error())
			return
		}

		if mwOb.IsEnabled != nil && !*mwOb.IsEnabled {
			if (mwOb.StartDay != nil && *mwOb.StartDay != 0) ||
				(mwOb.StartTime != nil && *mwOb.StartTime != "00:00") {
				resp.Diagnostics.AddError("Maintenance window start_day and start_time cannot be set if is_enabled is false", "Please either remove or comment out start_time and start_day values or the whole maintenance_window block.")
			}

			return
		}
	}

	if resp.PlanValue.IsUnknown() {
		resp.PlanValue = basetypes.NewObjectValueMust(resp.PlanValue.AttributeTypes(ctx), map[string]attr.Value{
			"is_enabled": basetypes.NewBoolValue(false),
			"start_day":  basetypes.NewInt64Value(0),
			"start_time": basetypes.NewStringValue("00:00"),
		})
	}
}
