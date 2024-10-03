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
// this modifier can merge state with config and return as plan(e.g. state is tags set in UI and config is tags set in terraform config)
func (m assignTagsModifier) PlanModifySet(ctx context.Context, req planmodifier.SetRequest, resp *planmodifier.SetResponse) {
	state := req.StateValue
	plan := resp.PlanValue
	config := req.ConfigValue

	// This is on creation.
	// Do nothing if there is no state value.
	if req.StateValue.IsNull() {
		return
	}

	// Below is everything else ie update with tags set in config.

	// This is on update and tags are not set in config so just use state as plan
	if req.PlanValue.IsUnknown() {
		resp.PlanValue = basetypes.NewSetValueMust(req.ConfigValue.ElementType(ctx), state.Elements())
		return
	}

	// check for tag duplicates in config
	checkDupes := make(map[string]interface{})
	for _, configTag := range config.Elements() {
		tagName := configTag.(basetypes.ObjectValue).Attributes()["tag_name"].(basetypes.StringValue).ValueString()
		checkDupes[tagName] = nil
	}

	// if checkDupes is not equal to plan.Elements() then there are duplicates
	if len(checkDupes) != len(config.Elements()) {
		resp.Diagnostics.AddError("Duplicate tag_name not allowed", "Please remove duplicate tag_name")
	}

	// merge plan into newPlan (plan is from config) and merge state in newPlan (state is from read)
	newPlan := state.Elements()
	for _, planTag := range plan.Elements() {
		existing := false
		for _, newPlanElem := range newPlan {
			if planTag.Equal(newPlanElem) {
				existing = true
				continue
			}
		}
		if !existing {
			newPlan = append(newPlan, planTag)
		}
	}

	resp.PlanValue = basetypes.NewSetValueMust(req.ConfigValue.ElementType(ctx), newPlan)
}

func ContainsTag(s []attr.Value, e attr.Value) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
