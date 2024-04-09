package plan_modifier

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func CustomAssignTags() planmodifier.Set {
	return assignTagsModifier{}
}

// assignTagsModifier implements the plan modifier.
type assignTagsModifier struct{}

// Description returns a human-readable description of the plan modifier.
func (m assignTagsModifier) Description(_ context.Context) string {
	return "Once set, the value of this attribute in state will not change."
}

// MarkdownDescription returns a markdown description of the plan modifier.
func (m assignTagsModifier) MarkdownDescription(_ context.Context) string {
	return "Once set, the value of this attribute in state will not change."
}

// PlanModifySet implements the plan modification logic.
func (m assignTagsModifier) PlanModifySet(ctx context.Context, req planmodifier.SetRequest, resp *planmodifier.SetResponse) {
	state := req.StateValue
	plan := resp.PlanValue

	newPlan := []attr.Value{}

	// This is on creation.
	// Do nothing if there is no state value.
	if req.StateValue.IsNull() {
		return
	}

	// This is on update and tags are not set in config so just plan for state
	if req.PlanValue.IsUnknown() {
		resp.PlanValue = basetypes.NewSetValueMust(req.ConfigValue.ElementType(ctx), state.Elements())
		return
	}

	// This is for anything else ie update with tags set in config.

	// merge plan with state in newPlan
	for _, v := range state.Elements() {
		newPlan = append(newPlan, v)
	}

	// merge plan with state in newPlan
	for _, v := range plan.Elements() {
		newPlan = append(newPlan, v)
	}

	resp.PlanValue = basetypes.NewSetValueMust(req.ConfigValue.ElementType(ctx), newPlan)
}
