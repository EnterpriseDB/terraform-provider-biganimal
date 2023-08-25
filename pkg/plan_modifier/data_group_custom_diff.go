package plan_modifier

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func CustomDataGroupDiffConfig() planmodifier.Set {
	return customDataGroupDiffModifier{}
}

// customDataGroupModifier implements the plan modifier.
type customDataGroupDiffModifier struct{}

// Description returns a human-readable description of the plan modifier.
func (m customDataGroupDiffModifier) Description(_ context.Context) string {
	return "Once set, the value of this attribute in state will not change."
}

// MarkdownDescription returns a markdown description of the plan modifier.
func (m customDataGroupDiffModifier) MarkdownDescription(_ context.Context) string {
	return "Once set, the value of this attribute in state will not change."
}

// PlanModifySet implements the plan modification logic.
func (m customDataGroupDiffModifier) PlanModifySet(ctx context.Context, req planmodifier.SetRequest, resp *planmodifier.SetResponse) {
	if req.StateValue.IsNull() {
		return
	}

	planDgs := resp.PlanValue.Elements()
	stateDgs := req.StateValue.Elements()

	if len(planDgs) == 0 {
		resp.Diagnostics.AddWarning("No data groups in config", "No data groups in config please add at least 1 data group")
		return
	}

	newPlan := []attr.Value{}

	// Need to sort the plan according to the state this is so the compare and setting unknowns are correct
	// https://developer.hashicorp.com/terraform/plugin/framework/resources/plan-modification#caveats
	// sort the order of the plan the same as the state, state is from the read and plan is from the config
	// plan will compare against state from read() and plan will also verify it is the same as the config via schema types
	for _, sDg := range stateDgs {
		stateRegion := sDg.(basetypes.ObjectValue).Attributes()["region"]
		for _, pDg := range planDgs {
			planRegion := pDg.(basetypes.ObjectValue).Attributes()["region"]
			if stateRegion.Equal(planRegion) {
				// set the unknowns manually mainly for delete group.
				// if we don't set manually they will be set automatically if you put the state in plan value
				// and they will be set by plan dg index against state dg index which will be in wrong order
				// if deleting a group.
				pDgAttrTypes := types.ObjectNull(planDgs[0].(basetypes.ObjectValue).AttributeTypes(ctx))
				pDgAttrValues := pDg.(basetypes.ObjectValue).Attributes()
				pDgAttrValues["cluster_name"] = sDg.(basetypes.ObjectValue).Attributes()["cluster_name"]
				pDgAttrValues["cluster_type"] = sDg.(basetypes.ObjectValue).Attributes()["cluster_type"]
				pDgAttrValues["conditions"] = sDg.(basetypes.ObjectValue).Attributes()["conditions"]
				pDgAttrValues["connection_uri"] = sDg.(basetypes.ObjectValue).Attributes()["connection_uri"]
				pDgAttrValues["created_at"] = sDg.(basetypes.ObjectValue).Attributes()["created_at"]
				pDgAttrValues["group_id"] = sDg.(basetypes.ObjectValue).Attributes()["group_id"]
				pDgAttrValues["logs_url"] = sDg.(basetypes.ObjectValue).Attributes()["logs_url"]
				pDgAttrValues["metrics_url"] = sDg.(basetypes.ObjectValue).Attributes()["metrics_url"]
				pDgAttrValues["phase"] = sDg.(basetypes.ObjectValue).Attributes()["phase"]
				pDgAttrValues["resizing_pvc"] = sDg.(basetypes.ObjectValue).Attributes()["resizing_pvc"]

				pDgStorage := pDg.(basetypes.ObjectValue).Attributes()["storage"].(basetypes.ObjectValue).Attributes()
				sDgStorage := sDg.(basetypes.ObjectValue).Attributes()["storage"].(basetypes.ObjectValue).Attributes()
				storageAttrTypes := types.ObjectNull(pDgAttrValues["storage"].(basetypes.ObjectValue).AttributeTypes(ctx))
				storageAttrValues := pDgStorage
				storageAttrValues["iops"] = sDgStorage["iops"]
				storageAttrValues["throughput"] = sDgStorage["throughput"]

				pDgAttrValues["storage"] = types.ObjectValueMust(storageAttrTypes.AttributeTypes(ctx), storageAttrValues)

				dgOb, diag := types.ObjectValue(pDgAttrTypes.AttributeTypes(ctx), pDgAttrValues)
				if diag.HasError() {
					resp.Diagnostics.Append(diag...)
					return
				}

				newPlan = append(newPlan, dgOb)
			}
		}
	}

	// add new groups
	for _, pDg := range planDgs {
		planGroupExistsInStateGroups := false
		planRegion := pDg.(basetypes.ObjectValue).Attributes()["region"]
		for _, sDg := range stateDgs {
			stateRegion := sDg.(basetypes.ObjectValue).Attributes()["region"]
			if stateRegion.Equal(planRegion) {
				planGroupExistsInStateGroups = true
				break
			}
		}

		if !planGroupExistsInStateGroups {
			newPlan = append(newPlan, pDg)
			resp.Diagnostics.AddWarning("Adding new data group", fmt.Sprintf("Adding new data group with region %v", planRegion))
		}
	}

	// remove groups
	for _, sDg := range stateDgs {
		stateGroupExistsInPlanGroups := false
		stateRegion := sDg.(basetypes.ObjectValue).Attributes()["region"]
		for _, pDg := range planDgs {
			planRegion := pDg.(basetypes.ObjectValue).Attributes()["region"]
			if stateRegion.Equal(planRegion) {
				stateGroupExistsInPlanGroups = true
				break
			}
		}

		if !stateGroupExistsInPlanGroups {
			resp.Diagnostics.AddWarning("Removing data group", fmt.Sprintf("Removing data group with region %v", stateRegion))
		}
	}

	if len(newPlan) == 0 {
		resp.Diagnostics.AddWarning("Plan data group generation error", "Plan data group error: regions may not be matching, regions missing in config or no data groups in config")
		return
	}
	resp.PlanValue = basetypes.NewSetValueMust(newPlan[0].Type(ctx), newPlan)

	for _, planDg := range resp.PlanValue.Elements() {
		if stateDgs == nil {
			return
		}
		var stateDgKey *int
		for k := range stateDgs {
			if stateDgs[k].(basetypes.ObjectValue).Attributes()["region"].Equal(planDg.(basetypes.ObjectValue).Attributes()["region"]) {
				k := k
				stateDgKey = &k
				break
			}
		}

		// data group may not exist in state because user is adding a new group with a new region
		if stateDgKey == nil {
			continue
		}

		if stateDgKey != nil {

			// allowed ips
			planAllowedIps := planDg.(basetypes.ObjectValue).Attributes()["allowed_ip_ranges"]
			stateAllowedIps := stateDgs[*stateDgKey].(basetypes.ObjectValue).Attributes()["allowed_ip_ranges"]

			if !planAllowedIps.Equal(stateAllowedIps) {
				resp.Diagnostics.AddWarning("Allowed IP ranges changed", fmt.Sprintf("Allowed IP ranges have changed from %v to %v",
					stateAllowedIps,
					planAllowedIps))
			}

			// backup retention period
			planBackupRetention := planDg.(basetypes.ObjectValue).Attributes()["backup_retention_period"]
			stateBackupRetention := stateDgs[*stateDgKey].(basetypes.ObjectValue).Attributes()["backup_retention_period"]

			if !planBackupRetention.Equal(stateBackupRetention) {
				resp.Diagnostics.AddWarning("Backup retention changed", fmt.Sprintf("backup retention period has changed from %v to %v",
					stateBackupRetention,
					planBackupRetention))
			}

			// cluster architecture
			planArch := planDg.(basetypes.ObjectValue).Attributes()["cluster_architecture"]
			stateArch := stateDgs[*stateDgKey].(basetypes.ObjectValue).Attributes()["cluster_architecture"]

			pArchId := planArch.(basetypes.ObjectValue).Attributes()["cluster_architecture_id"]
			pArchWitnessNodes := planArch.(basetypes.ObjectValue).Attributes()["witness_nodes"]
			pArchNodes := planArch.(basetypes.ObjectValue).Attributes()["nodes"]

			sArchId := stateArch.(basetypes.ObjectValue).Attributes()["cluster_architecture_id"]
			sArchWitnessNodes := stateArch.(basetypes.ObjectValue).Attributes()["witness_nodes"]
			sArchNodes := stateArch.(basetypes.ObjectValue).Attributes()["nodes"]

			if !pArchId.Equal(sArchId) || !pArchWitnessNodes.Equal(sArchWitnessNodes) || !pArchNodes.Equal(sArchNodes) {
				resp.Diagnostics.AddWarning("Cluster architecture changed", fmt.Sprintf("Cluster architecture changed from %v to %v",
					stateArch,
					planArch))
			}

			// csp auth
			planCspAuth := planDg.(basetypes.ObjectValue).Attributes()["csp_auth"]
			stateCspAuth := stateDgs[*stateDgKey].(basetypes.ObjectValue).Attributes()["csp_auth"]

			if !planCspAuth.Equal(stateCspAuth) {
				resp.Diagnostics.AddWarning("CSP auth changed", fmt.Sprintf("CSP auth changed from %v to %v",
					stateCspAuth,
					planCspAuth))
			}

			// instance type
			planInstanceType := planDg.(basetypes.ObjectValue).Attributes()["instance_type"]
			stateInstanceType := stateDgs[*stateDgKey].(basetypes.ObjectValue).Attributes()["instance_type"]

			if !planInstanceType.Equal(stateInstanceType) {
				resp.Diagnostics.AddWarning("Instance type changed", fmt.Sprintf("Instance type changed from %v to %v",
					stateInstanceType,
					planInstanceType))
			}

			// pg config
			// TODO: We should add the default pg_config values to the plan, so that we show the correct drift.
			// For details, please check pkg/plan_modifier/pg_config.go
			planPgConfig := planDg.(basetypes.ObjectValue).Attributes()["pg_config"]
			statePgConfig := stateDgs[*stateDgKey].(basetypes.ObjectValue).Attributes()["pg_config"]

			if !planPgConfig.Equal(statePgConfig) {
				resp.Diagnostics.AddWarning("PG config changed", fmt.Sprintf("PG config changed from %v to %v",
					statePgConfig,
					planPgConfig))
			}

			// storage
			planStorage := planDg.(basetypes.ObjectValue).Attributes()["storage"]
			stateStorage := stateDgs[*stateDgKey].(basetypes.ObjectValue).Attributes()["storage"]

			pStorageType := planStorage.(basetypes.ObjectValue).Attributes()["volume_type"]
			pStorageProperties := planStorage.(basetypes.ObjectValue).Attributes()["volume_properties"]
			pStorageSize := planStorage.(basetypes.ObjectValue).Attributes()["size"]

			sStorageType := stateStorage.(basetypes.ObjectValue).Attributes()["volume_type"]
			sStorageProperties := stateStorage.(basetypes.ObjectValue).Attributes()["volume_properties"]
			sStorageSize := stateStorage.(basetypes.ObjectValue).Attributes()["size"]

			if !pStorageType.Equal(sStorageType) || !pStorageProperties.Equal(sStorageProperties) || !pStorageSize.Equal(sStorageSize) {
				resp.Diagnostics.AddWarning("Storage changed", fmt.Sprintf("Storage changed from %v to %v",
					stateStorage,
					planStorage))
			}

			// pg type
			planPGType := planDg.(basetypes.ObjectValue).Attributes()["pg_type"]
			statePGType := stateDgs[*stateDgKey].(basetypes.ObjectValue).Attributes()["pg_type"]

			if !planPGType.Equal(statePGType) {
				resp.Diagnostics.AddError("PG type cannot be changed",
					fmt.Sprintf("PG type cannot be changed. PG type changed from expected value %v to %v in config",
						statePGType,
						planPGType))
				return
			}

			// pg version
			planPGVersion := planDg.(basetypes.ObjectValue).Attributes()["pg_version"]
			statePGVersion := stateDgs[*stateDgKey].(basetypes.ObjectValue).Attributes()["pg_version"]

			if !planPGVersion.Equal(statePGVersion) {
				resp.Diagnostics.AddError("PG version cannot be changed",
					fmt.Sprintf("PG version cannot be changed. PG version changed from expected value %v to %v in config",
						statePGVersion,
						planPGVersion))
				return
			}

			// networking
			planNetworking := planDg.(basetypes.ObjectValue).Attributes()["private_networking"]
			stateNetworking := stateDgs[*stateDgKey].(basetypes.ObjectValue).Attributes()["private_networking"]

			if !planNetworking.Equal(stateNetworking) {
				resp.Diagnostics.AddWarning("Private networking changed", fmt.Sprintf("Private networking changed from %v to %v",
					stateNetworking,
					planNetworking))
			}

			// cloud provider
			planCloudProvider := planDg.(basetypes.ObjectValue).Attributes()["cloud_provider"]
			stateCloudProvider := stateDgs[*stateDgKey].(basetypes.ObjectValue).Attributes()["cloud_provider"]

			if !planCloudProvider.Equal(stateCloudProvider) {
				resp.Diagnostics.AddError("Cloud provider cannot be changed",
					fmt.Sprintf("Cloud provider cannot be changed. Cloud provider changed from expected value: %v to %v in config",
						stateCloudProvider,
						planCloudProvider))
				return
			}

			// region
			planRegion := planDg.(basetypes.ObjectValue).Attributes()["region"]
			stateRegion := stateDgs[*stateDgKey].(basetypes.ObjectValue).Attributes()["region"]

			if !planRegion.Equal(stateRegion) {
				resp.Diagnostics.AddWarning("Region changed", fmt.Sprintf("Region changed from %v to %v",
					stateRegion,
					planRegion))
			}

			// maintenance window
			planMW := planDg.(basetypes.ObjectValue).Attributes()["maintenance_window"]
			stateMw := stateDgs[*stateDgKey].(basetypes.ObjectValue).Attributes()["maintenance_window"]

			if !planMW.Equal(stateMw) {
				resp.Diagnostics.AddWarning("Maintenance window changed", fmt.Sprintf("Maintenance window changed from %v to %v",
					stateMw,
					planMW))
			}
		}

	}
}
