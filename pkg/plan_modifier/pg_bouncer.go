package plan_modifier

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func CustomPgBouncer() planmodifier.Object {
	return CustomPgBouncerModifier{}
}

// CustomPgBouncerModifier implements the plan modifier.
type CustomPgBouncerModifier struct{}

// Description returns a human-readable description of the plan modifier.
func (m CustomPgBouncerModifier) Description(_ context.Context) string {
	return "Once set, the value of this attribute in state will not change."
}

// MarkdownDescription returns a markdown description of the plan modifier.
func (m CustomPgBouncerModifier) MarkdownDescription(_ context.Context) string {
	return "Once set, the value of this attribute in state will not change."
}

// PlanModifyObject implements the plan modification logic.
func (m CustomPgBouncerModifier) PlanModifyObject(ctx context.Context, req planmodifier.ObjectRequest, resp *planmodifier.ObjectResponse) {
	if req.PlanValue.IsUnknown() {
		resp.PlanValue = basetypes.NewObjectValueMust(
			req.StateValue.AttributeTypes(ctx),
			map[string]attr.Value{
				"is_enabled": basetypes.NewBoolValue(false),
				"settings":   basetypes.NewSetNull(req.StateValue.AttributeTypes(ctx)["settings"].(types.SetType).ElemType),
			},
		)
		return
	}

	if !req.PlanValue.IsNull() {
		reqPlanSettings := req.PlanValue.Attributes()["settings"].(basetypes.SetValue)
		// have plan settings combine with state settings if planned settings len > 0
		if len(reqPlanSettings.Elements()) > 0 {
			reqPlanIsEnabled := req.PlanValue.Attributes()["is_enabled"].(basetypes.BoolValue)
			if !reqPlanIsEnabled.ValueBool() {
				if !req.PlanValue.Attributes()["settings"].(basetypes.SetValue).IsNull() {
					resp.Diagnostics.AddError("if pg_bouncer.is_enabled = false then pg_bouncer.settings should be removed", "please remove pg_bouncer.settings if pg_bouncer.is_enabled = false")
					return
				}
			} else if reqPlanIsEnabled.ValueBool() {
				respPlanSettings := resp.PlanValue.Attributes()["settings"].(basetypes.SetValue)
				newPlanWithPrefilledPlannedSettings := respPlanSettings.Elements()
				stateSettings := []attr.Value{}
				if len(req.StateValue.Attributes()) != 0 {
					stateSettings = req.StateValue.Attributes()["settings"].(basetypes.SetValue).Elements()
				}

				// combine state settings with plan settings
				for _, sSetting := range stateSettings {
					stateSettingName := sSetting.(basetypes.ObjectValue).Attributes()["name"]
					for _, pSetting := range newPlanWithPrefilledPlannedSettings {
						planSettingName := pSetting.(basetypes.ObjectValue).Attributes()["name"]
						if stateSettingName.Equal(planSettingName) {
							continue
						}

						newPlanWithPrefilledPlannedSettings = append(newPlanWithPrefilledPlannedSettings, sSetting)
					}
				}

				resp.PlanValue = basetypes.NewObjectValueMust(
					req.StateValue.AttributeTypes(ctx),
					map[string]attr.Value{
						"is_enabled": basetypes.NewBoolValue(true),
						"settings":   basetypes.NewSetValueMust(req.StateValue.AttributeTypes(ctx)["settings"].(types.SetType).ElemType, newPlanWithPrefilledPlannedSettings),
					},
				)

				return
			}
		}

		// if is_enabled = true and settings = []
		reqPlanIsEnabled := req.PlanValue.Attributes()["is_enabled"].(basetypes.BoolValue)
		if reqPlanIsEnabled.ValueBool() && !reqPlanSettings.IsUnknown() && len(reqPlanSettings.Elements()) == 0 {
			resp.Diagnostics.AddError("if pg_bouncer.is_enabled = true then pg_bouncer.settings cannot be []", "please remove pg_bouncer.settings or set pg_bouncer.settings")
			return
		}

		// if is_enabled = false and settings is null and state setting is null then use state value for unknown
		if !reqPlanIsEnabled.ValueBool() &&
			req.ConfigValue.Attributes()["settings"].(basetypes.SetValue).IsNull() &&
			len(req.StateValue.Attributes()) != 0 &&
			req.StateValue.Attributes()["settings"].(basetypes.SetValue).IsNull() {
			resp.PlanValue = req.StateValue
			return
		}
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
