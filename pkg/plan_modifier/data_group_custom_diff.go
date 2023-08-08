package plan_modifier

import (
	"context"
	"fmt"
	"sort"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func CustomDataGroupDiffConfig() planmodifier.List {
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

// PlanModifyList implements the plan modification logic.
func (m customDataGroupDiffModifier) PlanModifyList(ctx context.Context, req planmodifier.ListRequest, resp *planmodifier.ListResponse) {
	if req.StateValue.IsNull() {
		return
	}

	// for list data groups, allowed ip ranges and pg config we have to for new plan:
	// add existing elements
	// sort it
	// add new elements
	// note: after plan is set it compares with config

	// plan is from config
	planDgs := resp.PlanValue.Elements()
	// state is from read
	stateDgs := req.StateValue.Elements()

	if len(planDgs) == 0 {
		resp.Diagnostics.AddWarning("No data groups in config", "No data groups in config please add at least 1 data group")
		return
	}

	newPlan := []attr.Value{}

	// hack need to sort plan we are using a slice instead of type.Set. This is so the compare and value setting is correct
	// https://developer.hashicorp.com/terraform/plugin/framework/resources/plan-modification#caveats
	// sort the order of the plan the same as the state, state is from the read and plan is from the config
	for _, sDg := range stateDgs {
		stateRegion := sDg.(basetypes.ObjectValue).Attributes()["region"]
		for _, pDg := range planDgs {
			planRegion := pDg.(basetypes.ObjectValue).Attributes()["region"]
			if stateRegion.Equal(planRegion) {
				newPlan = append(newPlan, pDg)
			}
		}
	}

	// sort data groups
	sort.Slice(newPlan, func(i, j int) bool {
		iString := newPlan[i].(basetypes.ObjectValue).Attributes()["region"].String()
		jString := newPlan[j].(basetypes.ObjectValue).Attributes()["region"].String()
		return iString < jString
	})

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
		}
	}

	// for allowed ip ranges and pg config
	// add the existing values from state to new allowed ip ranges var state.allowed.cidr and plan.allowed.cidr matches
	// sort the new allowed ip ranges var
	// add new allowed ip range from source plan to new allowed ip ranges var

	// newPlan is sorted according to statedDgs
	// add existing to new plan
	for k, sDg := range stateDgs {
		if k > (len(newPlan) - 1) {
			// no more state dgs matching plan dgs
			return
		}

		newAllowedIps := []attr.Value{}
		newPgConfig := []attr.Value{}

		pA := newPlan[k].(basetypes.ObjectValue).Attributes()["allowed_ip_ranges"].(basetypes.ListValue).Elements()
		sA := sDg.(basetypes.ObjectValue).Attributes()["allowed_ip_ranges"].(basetypes.ListValue).Elements()

		for _, pa := range pA {
			for _, sa := range sA {
				if pa.Equal(sa) {
					newAllowedIps = append(newAllowedIps, pa)
				}
			}
		}

		newPlan[k].(basetypes.ObjectValue).Attributes()["allowed_ip_ranges"] = basetypes.NewSetValueMust(newAllowedIps[0].Type(ctx), newAllowedIps)

		pPg := newPlan[k].(basetypes.ObjectValue).Attributes()["pg_config"].(basetypes.ListValue).Elements()
		sPg := sDg.(basetypes.ObjectValue).Attributes()["pg_config"].(basetypes.ListValue).Elements()

		for _, ppg := range pPg {
			for _, spg := range sPg {
				if ppg.Equal(spg) {
					newPgConfig = append(newPgConfig, ppg)
				}
			}
		}

		newPlan[k].(basetypes.ObjectValue).Attributes()["pg_config"] = basetypes.NewSetValueMust(newPgConfig[0].Type(ctx), newPgConfig)

	}

	// sort new plan
	for _, dg := range newPlan {
		allowedIps := dg.(basetypes.ObjectValue).Attributes()["allowed_ip_ranges"].(basetypes.ListValue).Elements()
		sort.Slice(allowedIps, func(i, j int) bool {
			iString := allowedIps[i].(basetypes.ObjectValue).Attributes()["cidr_block"].String()
			jString := allowedIps[j].(basetypes.ObjectValue).Attributes()["cidr_block"].String()
			return iString < jString
		})

		if len(allowedIps) != 0 {
			dg.(basetypes.ObjectValue).Attributes()["allowed_ip_ranges"] = basetypes.NewListValueMust(allowedIps[0].Type(ctx), allowedIps)
		}

		pgConfig := dg.(basetypes.ObjectValue).Attributes()["pg_config"].(basetypes.ListValue).Elements()
		sort.Slice(pgConfig, func(i, j int) bool {
			iString := pgConfig[i].(basetypes.ObjectValue).Attributes()["name"].String()
			jString := pgConfig[j].(basetypes.ObjectValue).Attributes()["name"].String()
			return iString < jString
		})

		if len(pgConfig) != 0 {
			dg.(basetypes.ObjectValue).Attributes()["pg_config"] = basetypes.NewListValueMust(pgConfig[0].Type(ctx), pgConfig)
		}
	}

	// adding new allowed ips and pg config to new plan
	for k, stateDg := range stateDgs {
		if k > (len(newPlan) - 1) {
			// no more state dgs matching plan dgs
			return
		}

		newPlanAllowedIps := newPlan[k].(basetypes.ObjectValue).Attributes()["allowed_ip_ranges"].(basetypes.ListValue).Elements()
		stateAllowedIps := stateDg.(basetypes.ObjectValue).Attributes()["allowed_ip_ranges"].(basetypes.ListValue).Elements()

		for _, pa := range newPlanAllowedIps {
			allowedIpExists := false
			for _, sa := range stateAllowedIps {
				paCidr := pa.(basetypes.ObjectValue).Attributes()["cidr_block"].(basetypes.StringValue).String()
				saCidr := sa.(basetypes.ObjectValue).Attributes()["cidr_block"].(basetypes.StringValue).String()
				if paCidr == saCidr {
					allowedIpExists = true
					break
				}
			}

			if !allowedIpExists {
				newPlanAllowedIps = append(newPlanAllowedIps, pa)
			}
		}

		if len(newPlanAllowedIps) != 0 {
			newPlan[k].(basetypes.ObjectValue).Attributes()["allowed_ip_ranges"] = basetypes.NewSetValueMust(newPlanAllowedIps[0].Type(ctx), newPlanAllowedIps)
		}

		newPlanPgConfig := newPlan[k].(basetypes.ObjectValue).Attributes()["pg_config"].(basetypes.ListValue).Elements()
		statePgConfig := stateDg.(basetypes.ObjectValue).Attributes()["pg_config"].(basetypes.ListValue).Elements()

		for _, pa := range newPlanPgConfig {
			pgConfigExists := false
			for _, sa := range statePgConfig {
				pName := pa.(basetypes.ObjectValue).Attributes()["name"].(basetypes.StringValue).String()
				sName := sa.(basetypes.ObjectValue).Attributes()["name"].(basetypes.StringValue).String()
				if pName == sName {
					pgConfigExists = true
					break
				}
			}

			if !pgConfigExists {
				newPlanPgConfig = append(newPlanPgConfig, pa)
			}
		}

		if len(newPlanPgConfig) != 0 {
			newPlan[k].(basetypes.ObjectValue).Attributes()["pg_config"] = basetypes.NewSetValueMust(newPlanPgConfig[0].Type(ctx), newPlanPgConfig)
		}
	}

	if len(newPlan) == 0 {
		resp.Diagnostics.AddWarning("Plan data group generation error", "Plan data group error: regions may not be matching, regions missing in config or no data groups in config")
		return
	}
	resp.PlanValue = basetypes.NewListValueMust(newPlan[0].Type(ctx), newPlan)

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
			// planAllowedIps := planDg.(basetypes.ObjectValue).Attributes()["allowed_ip_ranges"]
			// stateAllowedIps := stateDgs[*stateDgKey].(basetypes.ObjectValue).Attributes()["allowed_ip_ranges"]

			// if !planAllowedIps.Equal(stateAllowedIps) {
			// 	resp.Diagnostics.AddWarning("Allowed IP ranges changed", fmt.Sprintf("Allowed IP ranges have changed from %v to %v",
			// 		stateAllowedIps,
			// 		planAllowedIps))
			// }

			// backup retention period
			// planBackupRetention := planDg.(basetypes.ObjectValue).Attributes()["backup_retention_period"]
			// stateBackupRetention := stateDgs[*stateDgKey].(basetypes.ObjectValue).Attributes()["backup_retention_period"]

			// if !planBackupRetention.Equal(stateBackupRetention) {
			// 	resp.Diagnostics.AddWarning("Backup retention changed", fmt.Sprintf("backup retention period has changed from %v to %v",
			// 		stateBackupRetention,
			// 		planBackupRetention))
			// }

			// cluster architecture
			// planArch := planDg.(basetypes.ObjectValue).Attributes()["cluster_architecture"]
			// stateArch := stateDgs[*stateDgKey].(basetypes.ObjectValue).Attributes()["cluster_architecture"]

			// pArchId := planArch.(basetypes.ObjectValue).Attributes()["cluster_architecture_id"]
			// pArchWitnessNodes := planArch.(basetypes.ObjectValue).Attributes()["witness_nodes"]
			// pArchNodes := planArch.(basetypes.ObjectValue).Attributes()["nodes"]

			// sArchId := stateArch.(basetypes.ObjectValue).Attributes()["cluster_architecture_id"]
			// sArchWitnessNodes := stateArch.(basetypes.ObjectValue).Attributes()["witness_nodes"]
			// sArchNodes := stateArch.(basetypes.ObjectValue).Attributes()["nodes"]

			// if !pArchId.Equal(sArchId) || !pArchWitnessNodes.Equal(sArchWitnessNodes) || !pArchNodes.Equal(sArchNodes) {
			// 	resp.Diagnostics.AddWarning("Cluster architecture changed", fmt.Sprintf("Cluster architecture changed from %v to %v",
			// 		stateArch,
			// 		planArch))
			// }

			// csp auth
			// planCspAuth := planDg.(basetypes.ObjectValue).Attributes()["csp_auth"]
			// stateCspAuth := stateDgs[*stateDgKey].(basetypes.ObjectValue).Attributes()["csp_auth"]

			// if !planCspAuth.Equal(stateCspAuth) {
			// 	resp.Diagnostics.AddWarning("CSP auth changed", fmt.Sprintf("CSP auth changed from %v to %v",
			// 		stateCspAuth,
			// 		planCspAuth))
			// }

			// instance type
			// planInstanceType := planDg.(basetypes.ObjectValue).Attributes()["instance_type"]
			// stateInstanceType := stateDgs[*stateDgKey].(basetypes.ObjectValue).Attributes()["instance_type"]

			// if !planInstanceType.Equal(stateInstanceType) {
			// 	resp.Diagnostics.AddWarning("Instance type changed", fmt.Sprintf("Instance type changed from %v to %v",
			// 		stateInstanceType,
			// 		planInstanceType))
			// }

			// pg config
			// planPgConfig := planDg.(basetypes.ObjectValue).Attributes()["pg_config"]
			// statePgConfig := stateDgs[*stateDgKey].(basetypes.ObjectValue).Attributes()["pg_config"]

			// if !planPgConfig.Equal(statePgConfig) {
			// 	resp.Diagnostics.AddWarning("PG config changed", fmt.Sprintf("PG config changed from %v to %v",
			// 		statePgConfig,
			// 		planPgConfig))
			// }

			// storage
			// planStorage := planDg.(basetypes.ObjectValue).Attributes()["storage"]
			// stateStorage := stateDgs[*stateDgKey].(basetypes.ObjectValue).Attributes()["storage"]

			// pStorageType := planStorage.(basetypes.ObjectValue).Attributes()["volume_type"]
			// pStorageProperties := planStorage.(basetypes.ObjectValue).Attributes()["volume_properties"]
			// pStorageSize := planStorage.(basetypes.ObjectValue).Attributes()["size"]

			// sStorageType := stateStorage.(basetypes.ObjectValue).Attributes()["volume_type"]
			// sStorageProperties := stateStorage.(basetypes.ObjectValue).Attributes()["volume_properties"]
			// sStorageSize := stateStorage.(basetypes.ObjectValue).Attributes()["size"]

			// if !pStorageType.Equal(sStorageType) || !pStorageProperties.Equal(sStorageProperties) || !pStorageSize.Equal(sStorageSize) {
			// 	resp.Diagnostics.AddWarning("Storage changed", fmt.Sprintf("Storage changed from %v to %v",
			// 		stateStorage,
			// 		planStorage))
			// }

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
			// planNetworking := planDg.(basetypes.ObjectValue).Attributes()["private_networking"]
			// stateNetworking := stateDgs[*stateDgKey].(basetypes.ObjectValue).Attributes()["private_networking"]

			// if !planNetworking.Equal(stateNetworking) {
			// 	resp.Diagnostics.AddWarning("Private networking changed", fmt.Sprintf("Private networking changed from %v to %v",
			// 		stateNetworking,
			// 		planNetworking))
			// }

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
			// planMW := planDg.(basetypes.ObjectValue).Attributes()["maintenance_window"]
			// stateMw := stateDgs[*stateDgKey].(basetypes.ObjectValue).Attributes()["maintenance_window"]

			// if !planMW.Equal(stateMw) {
			// 	resp.Diagnostics.AddWarning("Maintenance window changed", fmt.Sprintf("Maintenance window changed from %v to %v",
			// 		stateMw,
			// 		planMW))
			// }
		}

	}
}
