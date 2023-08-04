package plan_modifier

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func CustomWitnessGroupDiffConfig() planmodifier.Set {
	return customWitnessGroupDiffModifier{}
}

// customWitnessGroupModifier implements the plan modifier.
type customWitnessGroupDiffModifier struct{}

// Description returns a human-readable description of the plan modifier.
func (m customWitnessGroupDiffModifier) Description(_ context.Context) string {
	return "Once set, the value of this attribute in state will not change."
}

// MarkdownDescription returns a markdown description of the plan modifier.
func (m customWitnessGroupDiffModifier) MarkdownDescription(_ context.Context) string {
	return "Once set, the value of this attribute in state will not change."
}

// PlanModifySet implements the plan modification logic.
func (m customWitnessGroupDiffModifier) PlanModifySet(ctx context.Context, req planmodifier.SetRequest, resp *planmodifier.SetResponse) {
	if req.StateValue.IsNull() {
		return
	}

	planWgs := resp.PlanValue.Elements()
	stateWgs := req.StateValue.Elements()

	newPlan := []attr.Value{}

	// hack need to sort plan we are using a slice instead of type.Set. This is so the compare and value setting is correct
	// https://developer.hashicorp.com/terraform/plugin/framework/resources/plan-modification#caveats
	// sort the order of the plan the same as the state, state is from the read and plan is from the config
	for _, sWg := range stateWgs {
		stateRegion := sWg.(basetypes.ObjectValue).Attributes()["region"]
		for _, pWg := range planWgs {
			planRegion := pWg.(basetypes.ObjectValue).Attributes()["region"]
			if stateRegion.Equal(planRegion) {
				newPlan = append(newPlan, pWg)
			}
		}
	}

	// add new groups
	for _, pWg := range planWgs {
		planGroupExistsInStateGroups := false
		var planRegion attr.Value
		planRegion = pWg.(basetypes.ObjectValue).Attributes()["region"]
		for _, sWg := range stateWgs {
			stateRegion := sWg.(basetypes.ObjectValue).Attributes()["region"]
			if stateRegion.Equal(planRegion) {
				planGroupExistsInStateGroups = true
				break
			}
		}

		if !planGroupExistsInStateGroups {
			newPlan = append(newPlan, pWg)
			resp.Diagnostics.AddWarning("Adding new witness group", fmt.Sprintf("Adding new witness group with region %v", planRegion))
		}
	}

	// remove groups
	for _, sWg := range stateWgs {
		stateGroupExistsInPlanGroups := false
		var stateRegion attr.Value
		stateRegion = sWg.(basetypes.ObjectValue).Attributes()["region"]
		for _, pWg := range planWgs {
			planRegion := pWg.(basetypes.ObjectValue).Attributes()["region"]
			if stateRegion.Equal(planRegion) {
				stateGroupExistsInPlanGroups = true
				break
			}
		}

		if !stateGroupExistsInPlanGroups {
			resp.Diagnostics.AddWarning("Removing witness group", fmt.Sprintf("Removing witness group with region %v", stateRegion))
		}
	}
	if len(newPlan) != 0 {
		resp.PlanValue = basetypes.NewSetValueMust(newPlan[0].Type(ctx), newPlan)
	}
}
