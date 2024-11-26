package plan_modifier

import (
	"context"
	"fmt"
	"reflect"

	"github.com/EnterpriseDB/terraform-provider-biganimal/pkg/models/pgd/terraform"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func CustomDataGroupDiffConfig() planmodifier.List {
	return CustomDataGroupDiffModifier{}
}

// CustomDataGroupModifier implements the plan modifier.
type CustomDataGroupDiffModifier struct{}

// Description returns a human-readable description of the plan modifier.
func (m CustomDataGroupDiffModifier) Description(_ context.Context) string {
	return "Once set, the value of this attribute in state will not change."
}

// MarkdownDescription returns a markdown description of the plan modifier.
func (m CustomDataGroupDiffModifier) MarkdownDescription(_ context.Context) string {
	return "Once set, the value of this attribute in state will not change."
}

// PlanModifyList implements the plan modification logic.
func (m CustomDataGroupDiffModifier) PlanModifyList(ctx context.Context, req planmodifier.ListRequest, resp *planmodifier.ListResponse) {
	if req.StateValue.IsNull() {
		// private networking case when doing create
		var planDgsObs []terraform.DataGroup
		diag := resp.PlanValue.ElementsAs(ctx, &planDgsObs, false)
		if diag.ErrorsCount() > 0 {
			resp.Diagnostics.Append(diag...)
			return
		}

		mapState := tfsdk.State{Schema: req.Plan.Schema, Raw: req.Plan.Raw}
		diag = mapState.SetAttribute(ctx, path.Root("data_groups"), planDgsObs)
		if diag.ErrorsCount() > 0 {
			resp.Diagnostics.Append(diag...)
			return
		}

		tfDgsMap := new(types.List)
		mapState.GetAttribute(ctx, path.Root("data_groups"), tfDgsMap)

		resp.PlanValue = *tfDgsMap

		return
	}

	if len(resp.PlanValue.Elements()) == 0 {
		resp.Diagnostics.AddWarning("No data groups in config", "No data groups in config please add at least 1 data group")
		return
	}

	newDgPlan := []terraform.DataGroup{}

	var stateDgsObs []terraform.DataGroup
	diag := req.StateValue.ElementsAs(ctx, &stateDgsObs, false)
	if diag.ErrorsCount() > 0 {
		resp.Diagnostics.Append(diag...)
		return
	}

	var planDgsObs []terraform.DataGroup
	diag = resp.PlanValue.ElementsAs(ctx, &planDgsObs, false)
	if diag.ErrorsCount() > 0 {
		resp.Diagnostics.Append(diag...)
		return
	}

	// Need to sort the plan according to the state this is so the compare and setting unknowns are correct
	// https://developer.hashicorp.com/terraform/plugin/framework/resources/plan-modification#caveats
	// sort the order of the plan the same as the state, state is from the read and plan is from the config
	// plan will compare against state from read() and plan will also verify it is the same as the config via schema types
	for _, sDg := range stateDgsObs {
		for _, pDg := range planDgsObs {
			// set the unknowns manually for delete and add group.
			// if we don't set manually and it is set the same way as useStateForUnknown,
			// then when it puts the state in plan value it will be set by plan dg index
			// against state dg index which will be in wrong order if delete a group.
			if reflect.DeepEqual(sDg.Region, pDg.Region) {
				pDg.ClusterArchitecture.ClusterArchitectureName = sDg.ClusterArchitecture.ClusterArchitectureName
				pDg.ClusterArchitecture.WitnessNodes = sDg.ClusterArchitecture.WitnessNodes
				pDg.ClusterName = sDg.ClusterName
				pDg.ClusterType = sDg.ClusterType
				pDg.Connection = sDg.Connection
				pDg.CreatedAt = sDg.CreatedAt
				pDg.GroupId = sDg.GroupId
				pDg.LogsUrl = sDg.LogsUrl
				pDg.MetricsUrl = sDg.MetricsUrl
				pDg.Phase = sDg.Phase
				pDg.ResizingPvc = sDg.ResizingPvc
				pDg.Storage.Iops = sDg.Storage.Iops
				pDg.Storage.Throughput = sDg.Storage.Throughput

				// if private networking has change then connection string will change
				if sDg.PrivateNetworking != pDg.PrivateNetworking {
					pDg.Connection = types.StringUnknown()
				}

				newDgPlan = append(newDgPlan, pDg)
			}
		}
	}

	// add new groups
	for _, pDg := range planDgsObs {
		planGroupExistsInStateGroups := false
		for _, sDg := range stateDgsObs {
			if reflect.DeepEqual(sDg.Region, pDg.Region) {
				planGroupExistsInStateGroups = true
				break
			}
		}

		if !planGroupExistsInStateGroups {
			newDgPlan = append(newDgPlan, pDg)
			resp.Diagnostics.AddWarning("Adding new data group", fmt.Sprintf("Adding new data group with region %v", pDg.Region.RegionId))
		}
	}

	// remove groups
	for _, sDg := range stateDgsObs {
		stateGroupExistsInPlanGroups := false
		for _, pDg := range planDgsObs {
			if reflect.DeepEqual(sDg.Region, pDg.Region) {
				stateGroupExistsInPlanGroups = true
				break
			}
		}

		if !stateGroupExistsInPlanGroups {
			resp.Diagnostics.AddWarning("Removing data group", fmt.Sprintf("Removing data group with region %v", sDg.Region.RegionId))
		}
	}

	if len(newDgPlan) == 0 {
		resp.Diagnostics.AddWarning("Plan data group generation error", "Plan data group error: regions may not be matching, regions missing in config or no data groups in config")
		return
	}

	mapState := tfsdk.State{Schema: req.Plan.Schema, Raw: req.Plan.Raw}
	diag = mapState.SetAttribute(ctx, path.Root("data_groups"), newDgPlan)
	if diag.ErrorsCount() > 0 {
		resp.Diagnostics.Append(diag...)
		return
	}

	tfDgsMap := new(types.List)
	mapState.GetAttribute(ctx, path.Root("data_groups"), tfDgsMap)

	resp.PlanValue = *tfDgsMap
}
