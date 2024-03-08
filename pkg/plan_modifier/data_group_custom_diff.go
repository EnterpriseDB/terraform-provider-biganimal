package plan_modifier

import (
	"context"
	"fmt"
	"reflect"

	"github.com/EnterpriseDB/terraform-provider-biganimal/pkg/models/pgd/terraform"
	"github.com/EnterpriseDB/terraform-provider-biganimal/pkg/utils"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func CustomDataGroupDiffConfig() planmodifier.Set {
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

// PlanModifySet implements the plan modification logic.
func (m CustomDataGroupDiffModifier) PlanModifySet(ctx context.Context, req planmodifier.SetRequest, resp *planmodifier.SetResponse) {
	if req.StateValue.IsNull() {
		// private networking case when doing create
		var planDgsObs []terraform.DataGroup
		diag := resp.PlanValue.ElementsAs(ctx, &planDgsObs, false)
		if diag.ErrorsCount() > 0 {
			resp.Diagnostics.Append(diag...)
			return
		}

		for _, pDg := range planDgsObs {
			// fix to set the correct allowed ip ranges to allow all if a PGD data group has private networking set as true
			if pDg.PrivateNetworking != nil && *pDg.PrivateNetworking {
				pDg.AllowedIpRanges = types.SetValueMust(pDg.AllowedIpRanges.ElementType(ctx), []attr.Value{
					types.ObjectValueMust(
						pDg.AllowedIpRanges.ElementType(ctx).(types.ObjectType).AttributeTypes(),
						map[string]attr.Value{
							"cidr_block":  types.StringValue("0.0.0.0/0"),
							"description": types.StringValue("To allow all access"),
						}),
				})
				// fix to set the correct allowed ip ranges for PGD data group if allowed ip ranges length is 0
			} else if pDg.AllowedIpRanges.IsNull() || len(pDg.AllowedIpRanges.Elements()) == 0 {
				pDg.AllowedIpRanges = types.SetValueMust(pDg.AllowedIpRanges.ElementType(ctx), []attr.Value{
					types.ObjectValueMust(
						pDg.AllowedIpRanges.ElementType(ctx).(types.ObjectType).AttributeTypes(),
						map[string]attr.Value{
							"cidr_block":  types.StringValue("0.0.0.0/0"),
							"description": types.StringValue(""),
						}),
				})
			}
		}

		mapState := tfsdk.State{Schema: req.Plan.Schema, Raw: req.Plan.Raw}
		diag = mapState.SetAttribute(ctx, path.Root("data_groups"), planDgsObs)
		if diag.ErrorsCount() > 0 {
			resp.Diagnostics.Append(diag...)
			return
		}

		tfDgsMap := new(types.Set)
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

				// fix to set the correct allowed ip ranges to allow all if a PGD data group has private networking set as true
				if pDg.PrivateNetworking != nil && *pDg.PrivateNetworking {
					pDg.AllowedIpRanges = types.SetValueMust(pDg.AllowedIpRanges.ElementType(ctx), []attr.Value{
						types.ObjectValueMust(
							pDg.AllowedIpRanges.ElementType(ctx).(types.ObjectType).AttributeTypes(),
							map[string]attr.Value{
								"cidr_block":  types.StringValue("0.0.0.0/0"),
								"description": types.StringValue("To allow all access"),
							}),
					})
					// fix to set the correct allowed ip ranges for PGD data group if allowed ip ranges length is 0
				} else if pDg.AllowedIpRanges.IsNull() || len(pDg.AllowedIpRanges.Elements()) == 0 {
					pDg.AllowedIpRanges = types.SetValueMust(pDg.AllowedIpRanges.ElementType(ctx), []attr.Value{
						types.ObjectValueMust(
							pDg.AllowedIpRanges.ElementType(ctx).(types.ObjectType).AttributeTypes(),
							map[string]attr.Value{
								"cidr_block":  types.StringValue("0.0.0.0/0"),
								"description": types.StringValue(""),
							}),
					})
				}

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

	tfDgsMap := new(types.Set)
	mapState.GetAttribute(ctx, path.Root("data_groups"), tfDgsMap)

	resp.PlanValue = *tfDgsMap

	for _, pDg := range newDgPlan {
		if len(stateDgsObs) == 0 {
			return
		}
		var foundStateDg *terraform.DataGroup
		for _, sDg := range stateDgsObs {
			if reflect.DeepEqual(sDg.Region, pDg.Region) {
				foundStateDg = &sDg
				break
			}
		}

		// data group may not exist in state because user is adding a new group with a new region
		if foundStateDg == nil {
			continue
		}

		if foundStateDg != nil {

			// allowed ips
			if !reflect.DeepEqual(pDg.AllowedIpRanges, foundStateDg.AllowedIpRanges) {
				resp.Diagnostics.AddWarning("Allowed IP ranges changed", fmt.Sprintf("Allowed IP ranges have changed from %v to %v for data group with region %v",
					foundStateDg.AllowedIpRanges.String(),
					pDg.AllowedIpRanges.String(),
					foundStateDg.Region.RegionId))
			}

			// backup retention period
			if !reflect.DeepEqual(pDg.BackupRetentionPeriod, foundStateDg.BackupRetentionPeriod) {
				resp.Diagnostics.AddWarning("Backup retention changed", fmt.Sprintf("backup retention period has changed from %v to %v for data group with region %v",
					*foundStateDg.BackupRetentionPeriod,
					*pDg.BackupRetentionPeriod,
					foundStateDg.Region.RegionId))
			}

			// cluster architecture
			if pDg.ClusterArchitecture.ClusterArchitectureId != foundStateDg.ClusterArchitecture.ClusterArchitectureId ||
				pDg.ClusterArchitecture.Nodes != foundStateDg.ClusterArchitecture.Nodes {
				resp.Diagnostics.AddWarning("Cluster architecture changed", fmt.Sprintf("Cluster architecture changed from %v to %v for data group with region %v",
					*foundStateDg.ClusterArchitecture,
					*pDg.ClusterArchitecture,
					foundStateDg.Region.RegionId))
			}

			// csp auth
			if !reflect.DeepEqual(pDg.CspAuth, foundStateDg.CspAuth) {
				resp.Diagnostics.AddWarning("CSP auth changed", fmt.Sprintf("CSP auth changed from %v to %v for data group with region %v",
					*foundStateDg.CspAuth,
					*pDg.CspAuth,
					foundStateDg.Region.RegionId))
			}

			// instance type
			if !reflect.DeepEqual(pDg.InstanceType, foundStateDg.InstanceType) {
				resp.Diagnostics.AddWarning("Instance type changed", fmt.Sprintf("Instance type changed from %v to %v for data group with region %v",
					*foundStateDg.InstanceType,
					*pDg.InstanceType,
					foundStateDg.Region.RegionId))
			}

			// storage
			if pDg.Storage.VolumeTypeId != foundStateDg.Storage.VolumeTypeId ||
				pDg.Storage.VolumePropertiesId != foundStateDg.Storage.VolumePropertiesId ||
				pDg.Storage.Size != foundStateDg.Storage.Size {
				resp.Diagnostics.AddWarning("Storage changed", fmt.Sprintf("Storage changed from %v to %v for data group with region %v",
					*foundStateDg.Storage,
					*pDg.Storage,
					foundStateDg.Region.RegionId))
			}

			// pg type
			if !reflect.DeepEqual(pDg.PgType, foundStateDg.PgType) {
				resp.Diagnostics.AddError("PG type cannot be changed",
					fmt.Sprintf("PG type cannot be changed. PG type changed from expected value %v to %v in config for data group with region %v",
						*foundStateDg.PgType,
						*pDg.PgType,
						foundStateDg.Region.RegionId))
				return
			}

			// pg version
			if !reflect.DeepEqual(pDg.PgVersion, foundStateDg.PgVersion) {
				resp.Diagnostics.AddError("PG version cannot be changed",
					fmt.Sprintf("PG version cannot be changed. PG version changed from expected value %v to %v in config for data group with region %v",
						*foundStateDg.PgVersion,
						*pDg.PgVersion,
						foundStateDg.Region.RegionId))
				return
			}

			// networking
			if !reflect.DeepEqual(pDg.PrivateNetworking, foundStateDg.PrivateNetworking) {
				resp.Diagnostics.AddWarning("Private networking changed", fmt.Sprintf("Private networking changed from %v to %v for data group with region %v",
					*foundStateDg.PrivateNetworking,
					*pDg.PrivateNetworking,
					foundStateDg.Region.RegionId))
			}

			// cloud provider
			if !reflect.DeepEqual(pDg.Provider, foundStateDg.Provider) {
				resp.Diagnostics.AddError("Cloud provider cannot be changed",
					fmt.Sprintf("Cloud provider cannot be changed. Cloud provider changed from expected value: %v to %v in config for data group with region %v",
						utils.PrintJson(*foundStateDg.Provider),
						utils.PrintJson(*pDg.Provider),
						foundStateDg.Region.RegionId))
				return
			}

			// region
			if !reflect.DeepEqual(pDg.Region, foundStateDg.Region) {
				resp.Diagnostics.AddWarning("Region changed", fmt.Sprintf("Region changed from %v to %v for data group with region %v",
					*foundStateDg.Region,
					*pDg.Region,
					foundStateDg.Region.RegionId))
			}

			// maintenance window
			if !reflect.DeepEqual(pDg.MaintenanceWindow, foundStateDg.MaintenanceWindow) {
				resp.Diagnostics.AddWarning("Maintenance window changed", fmt.Sprintf("Maintenance window changed from %v to %v for data group with region %v",
					utils.PrintJson(*foundStateDg.MaintenanceWindow),
					utils.PrintJson(*pDg.MaintenanceWindow),
					foundStateDg.Region.RegionId))
			}

			// pg config
			if !reflect.DeepEqual(pDg.PgConfig, foundStateDg.PgConfig) {
				resp.Diagnostics.AddWarning("Pg config changed", fmt.Sprintf("Pg config changed from %v to %v for data group with region %v",
					*foundStateDg.PgConfig,
					*pDg.PgConfig,
					foundStateDg.Region.RegionId))
			}
		}
	}
}
