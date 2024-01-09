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
	networkingAllowedIpRangesSetFunc := func(description string) basetypes.SetValue {
		defaultAttrs := map[string]attr.Value{"cidr_block": basetypes.NewStringValue("0.0.0.0/0"), "description": basetypes.NewStringValue(description)}
		defaultAttrTypes := map[string]attr.Type{"cidr_block": defaultAttrs["cidr_block"].Type(ctx), "description": defaultAttrs["description"].Type(ctx)}

		defaultObjectValue := basetypes.NewObjectValueMust(defaultAttrTypes, defaultAttrs)
		setOfObjects := []attr.Value{}
		setOfObjects = append(setOfObjects, defaultObjectValue)
		return basetypes.NewSetValueMust(defaultObjectValue.Type(ctx), setOfObjects)
	}

	if req.StateValue.IsNull() {

		newNetworkingPlan := []attr.Value{}
		for _, pDg := range resp.PlanValue.Elements() {
			pDgAttrTypes := types.ObjectNull(pDg.(basetypes.ObjectValue).AttributeTypes(ctx))
			pDgAttrValues := pDg.(basetypes.ObjectValue).Attributes()

			if pDgAttrValues["private_networking"].(basetypes.BoolValue).ValueBool() {
				// fix to set the correct allowed ip ranges to allow all if a PGD data group has private networking set as true
				pDgAttrValues["allowed_ip_ranges"] = networkingAllowedIpRangesSetFunc("To allow all access")
				dgOb, diag := types.ObjectValue(pDgAttrTypes.AttributeTypes(ctx), pDgAttrValues)
				if diag.HasError() {
					resp.Diagnostics.Append(diag...)
					return
				}

				newNetworkingPlan = append(newNetworkingPlan, dgOb)
			} else if len(pDgAttrValues["allowed_ip_ranges"].(types.Set).Elements()) == 0 {
				// fix to set the correct allowed ip ranges for PGD data group if allowed ip ranges length is 0
				pDgAttrValues["allowed_ip_ranges"] = networkingAllowedIpRangesSetFunc("")
				dgOb, diag := types.ObjectValue(pDgAttrTypes.AttributeTypes(ctx), pDgAttrValues)
				if diag.HasError() {
					resp.Diagnostics.Append(diag...)
					return
				}

				newNetworkingPlan = append(newNetworkingPlan, dgOb)
			}
		}

		if len(newNetworkingPlan) != 0 {
			resp.PlanValue = basetypes.NewSetValueMust(newNetworkingPlan[0].Type(ctx), newNetworkingPlan)
		}

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
				// set the unknowns manually for delete and add group.
				// if we don't set manually and it is set the same way as useStateForUnknown,
				// then when it puts the state in plan value it will be set by plan dg index
				// against state dg index which will be in wrong order if delete a group.
				pDgAttrTypes := types.ObjectNull(planDgs[0].(basetypes.ObjectValue).AttributeTypes(ctx))
				pDgAttrValues := pDg.(basetypes.ObjectValue).Attributes()
				sDgAttrValues := sDg.(basetypes.ObjectValue).Attributes()

				pDgClusterArch := pDgAttrValues["cluster_architecture"].(basetypes.ObjectValue).Attributes()
				sDgClusterArch := sDgAttrValues["cluster_architecture"].(basetypes.ObjectValue).Attributes()
				clusterArchAttrTypes := types.ObjectNull(pDgAttrValues["cluster_architecture"].(basetypes.ObjectValue).AttributeTypes(ctx))
				ClusterArchAttrValues := pDgClusterArch
				ClusterArchAttrValues["cluster_architecture_name"] = sDgClusterArch["cluster_architecture_name"]
				ClusterArchAttrValues["witness_nodes"] = sDgClusterArch["witness_nodes"]

				pDgAttrValues["cluster_architecture"] = types.ObjectValueMust(clusterArchAttrTypes.AttributeTypes(ctx), ClusterArchAttrValues)

				pDgAttrValues["cluster_name"] = sDgAttrValues["cluster_name"]
				pDgAttrValues["cluster_type"] = sDgAttrValues["cluster_type"]
				pDgAttrValues["conditions"] = sDgAttrValues["conditions"]
				pDgAttrValues["connection_uri"] = sDgAttrValues["connection_uri"]
				pDgAttrValues["created_at"] = sDgAttrValues["created_at"]
				pDgAttrValues["group_id"] = sDgAttrValues["group_id"]
				pDgAttrValues["logs_url"] = sDgAttrValues["logs_url"]
				pDgAttrValues["metrics_url"] = sDgAttrValues["metrics_url"]
				pDgAttrValues["phase"] = sDgAttrValues["phase"]
				pDgAttrValues["resizing_pvc"] = sDgAttrValues["resizing_pvc"]

				pDgStorage := pDgAttrValues["storage"].(basetypes.ObjectValue).Attributes()
				sDgStorage := sDgAttrValues["storage"].(basetypes.ObjectValue).Attributes()
				storageAttrTypes := types.ObjectNull(pDgAttrValues["storage"].(basetypes.ObjectValue).AttributeTypes(ctx))
				storageAttrValues := pDgStorage
				storageAttrValues["iops"] = sDgStorage["iops"]
				storageAttrValues["throughput"] = sDgStorage["throughput"]

				pDgAttrValues["storage"] = types.ObjectValueMust(storageAttrTypes.AttributeTypes(ctx), storageAttrValues)

				// fix to set the correct allowed ip ranges to allow all if a PGD data group has private networking set as true
				if pDgAttrValues["private_networking"].(basetypes.BoolValue).ValueBool() {
					pDgAttrValues["allowed_ip_ranges"] = networkingAllowedIpRangesSetFunc("To allow all access")
					// fix to set the correct allowed ip ranges for PGD data group if allowed ip ranges length is 0
				} else if len(pDgAttrValues["allowed_ip_ranges"].(basetypes.SetValue).Elements()) == 0 {
					pDgAttrValues["allowed_ip_ranges"] = networkingAllowedIpRangesSetFunc("")
				}

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
		var stateDgRegion string
		for k := range stateDgs {
			if stateDgs[k].(basetypes.ObjectValue).Attributes()["region"].Equal(planDg.(basetypes.ObjectValue).Attributes()["region"]) {
				k := k
				stateDgKey = &k
				stateDgRegion = stateDgs[k].(basetypes.ObjectValue).Attributes()["region"].String()
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
				resp.Diagnostics.AddWarning("Allowed IP ranges changed", fmt.Sprintf("Allowed IP ranges have changed from %v to %v for data group with region %v",
					stateAllowedIps,
					planAllowedIps,
					stateDgRegion))
			}

			// backup retention period
			planBackupRetention := planDg.(basetypes.ObjectValue).Attributes()["backup_retention_period"]
			stateBackupRetention := stateDgs[*stateDgKey].(basetypes.ObjectValue).Attributes()["backup_retention_period"]

			if !planBackupRetention.Equal(stateBackupRetention) {
				resp.Diagnostics.AddWarning("Backup retention changed", fmt.Sprintf("backup retention period has changed from %v to %v for data group with region %v",
					stateBackupRetention,
					planBackupRetention,
					stateDgRegion))
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
				resp.Diagnostics.AddWarning("Cluster architecture changed", fmt.Sprintf("Cluster architecture changed from %v to %v for data group with region %v",
					stateArch,
					planArch,
					stateDgRegion))
			}

			// csp auth
			planCspAuth := planDg.(basetypes.ObjectValue).Attributes()["csp_auth"]
			stateCspAuth := stateDgs[*stateDgKey].(basetypes.ObjectValue).Attributes()["csp_auth"]

			if !planCspAuth.Equal(stateCspAuth) {
				resp.Diagnostics.AddWarning("CSP auth changed", fmt.Sprintf("CSP auth changed from %v to %v for data group with region %v",
					stateCspAuth,
					planCspAuth,
					stateDgRegion))
			}

			// instance type
			planInstanceType := planDg.(basetypes.ObjectValue).Attributes()["instance_type"]
			stateInstanceType := stateDgs[*stateDgKey].(basetypes.ObjectValue).Attributes()["instance_type"]

			if !planInstanceType.Equal(stateInstanceType) {
				resp.Diagnostics.AddWarning("Instance type changed", fmt.Sprintf("Instance type changed from %v to %v for data group with region %v",
					stateInstanceType,
					planInstanceType,
					stateDgRegion))
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
				resp.Diagnostics.AddWarning("Storage changed", fmt.Sprintf("Storage changed from %v to %v for data group with region %v",
					stateStorage,
					planStorage,
					stateDgRegion))
			}

			// pg type
			planPGType := planDg.(basetypes.ObjectValue).Attributes()["pg_type"]
			statePGType := stateDgs[*stateDgKey].(basetypes.ObjectValue).Attributes()["pg_type"]

			if !planPGType.Equal(statePGType) {
				resp.Diagnostics.AddError("PG type cannot be changed",
					fmt.Sprintf("PG type cannot be changed. PG type changed from expected value %v to %v in config for data group with region %v",
						statePGType,
						planPGType,
						stateDgRegion))
				return
			}

			// pg version
			planPGVersion := planDg.(basetypes.ObjectValue).Attributes()["pg_version"]
			statePGVersion := stateDgs[*stateDgKey].(basetypes.ObjectValue).Attributes()["pg_version"]

			if !planPGVersion.Equal(statePGVersion) {
				resp.Diagnostics.AddError("PG version cannot be changed",
					fmt.Sprintf("PG version cannot be changed. PG version changed from expected value %v to %v in config for data group with region %v",
						statePGVersion,
						planPGVersion,
						stateDgRegion))
				return
			}

			// networking
			planNetworking := planDg.(basetypes.ObjectValue).Attributes()["private_networking"]
			stateNetworking := stateDgs[*stateDgKey].(basetypes.ObjectValue).Attributes()["private_networking"]

			if !planNetworking.Equal(stateNetworking) {
				resp.Diagnostics.AddWarning("Private networking changed", fmt.Sprintf("Private networking changed from %v to %v for data group with region %v",
					stateNetworking,
					planNetworking,
					stateDgRegion))
			}

			// cloud provider
			planCloudProvider := planDg.(basetypes.ObjectValue).Attributes()["cloud_provider"]
			stateCloudProvider := stateDgs[*stateDgKey].(basetypes.ObjectValue).Attributes()["cloud_provider"]

			if !planCloudProvider.Equal(stateCloudProvider) {
				resp.Diagnostics.AddError("Cloud provider cannot be changed",
					fmt.Sprintf("Cloud provider cannot be changed. Cloud provider changed from expected value: %v to %v in config for data group with region %v",
						stateCloudProvider,
						planCloudProvider,
						stateDgRegion))
				return
			}

			// region
			planRegion := planDg.(basetypes.ObjectValue).Attributes()["region"]
			stateRegion := stateDgs[*stateDgKey].(basetypes.ObjectValue).Attributes()["region"]

			if !planRegion.Equal(stateRegion) {
				resp.Diagnostics.AddWarning("Region changed", fmt.Sprintf("Region changed from %v to %v for data group with region %v",
					stateRegion,
					planRegion,
					stateDgRegion))
			}

			// maintenance window
			planMW := planDg.(basetypes.ObjectValue).Attributes()["maintenance_window"]
			stateMw := stateDgs[*stateDgKey].(basetypes.ObjectValue).Attributes()["maintenance_window"]

			if !planMW.Equal(stateMw) {
				resp.Diagnostics.AddWarning("Maintenance window changed", fmt.Sprintf("Maintenance window changed from %v to %v for data group with region %v",
					stateMw,
					planMW,
					stateDgRegion))
			}
		}

	}
}
