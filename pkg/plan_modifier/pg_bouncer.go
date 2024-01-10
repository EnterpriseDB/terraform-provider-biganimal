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
