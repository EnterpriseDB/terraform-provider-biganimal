package plan_modifier

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func CustomWitnessGroupDiffConfig() planmodifier.List {
	return CustomWitnessGroupDiffModifier{}
}

// CustomWitnessGroupModifier implements the plan modifier.
type CustomWitnessGroupDiffModifier struct{}

// Description returns a human-readable description of the plan modifier.
func (m CustomWitnessGroupDiffModifier) Description(_ context.Context) string {
	return "Once set, the value of this attribute in state will not change."
}

// MarkdownDescription returns a markdown description of the plan modifier.
func (m CustomWitnessGroupDiffModifier) MarkdownDescription(_ context.Context) string {
	return "Once set, the value of this attribute in state will not change."
}

// PlanModifyList implements the plan modification logic.
func (m CustomWitnessGroupDiffModifier) PlanModifyList(ctx context.Context, req planmodifier.ListRequest, resp *planmodifier.ListResponse) {
	if req.StateValue.IsNull() {
		return
	}

	planWgs := resp.PlanValue.Elements()
	stateWgs := req.StateValue.Elements()

	newPlan := []attr.Value{}

	// Need to sort the plan according to the state this is so the compare and setting unknowns are correct
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
		planRegion := pWg.(basetypes.ObjectValue).Attributes()["region"]
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

	// validations
	for pk, pWg := range planWgs {
		for sk, sWg := range stateWgs {
			planRegion := pWg.(basetypes.ObjectValue).Attributes()["region"]
			stateRegion := sWg.(basetypes.ObjectValue).Attributes()["region"]

			if !stateRegion.Equal(planRegion) && pk == sk {
				resp.Diagnostics.AddError("Witness group region cannot be changed",
					fmt.Sprintf("Witness group region cannot be changed. Witness group region changed from expected value %v to %v in config",
						stateRegion,
						planRegion))
				return
			}

			planCloudProvider := pWg.(basetypes.ObjectValue).Attributes()["cloud_provider"]
			stateCloudProvider := sWg.(basetypes.ObjectValue).Attributes()["cloud_provider"]

			if !planCloudProvider.Equal(stateCloudProvider) && pk == sk {
				resp.Diagnostics.AddError("Witness group cloud provider cannot be changed",
					fmt.Sprintf("witness group cloud provider cannot be changed. witness group cloud provider changed from expected value: %v to %v in config",
						stateCloudProvider,
						planCloudProvider))
				return
			}
		}
	}

	// remove groups
	// for _, sWg := range stateWgs {
	// 	stateGroupExistsInPlanGroups := false
	// 	stateRegion := sWg.(basetypes.ObjectValue).Attributes()["region"]
	// 	for _, pWg := range planWgs {
	// 		planRegion := pWg.(basetypes.ObjectValue).Attributes()["region"]
	// 		if stateRegion.Equal(planRegion) {
	// 			stateGroupExistsInPlanGroups = true
	// 			break
	// 		}
	// 	}

	// 	if !stateGroupExistsInPlanGroups {
	// 		resp.Diagnostics.AddWarning("Removing witness group", fmt.Sprintf("Removing witness group with region %v", stateRegion))
	// 	}
	// }
	if len(newPlan) != 0 {
		resp.PlanValue = basetypes.NewListValueMust(newPlan[0].Type(ctx), newPlan)
	} else {
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
}
