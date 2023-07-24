package plan_modifier

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
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

	newPlan := []attr.Value{}

	// hack need to sort plan we are using a slice instead of type.Set. This is so the compare and value setting is correct
	// https://developer.hashicorp.com/terraform/plugin/framework/resources/plan-modification#caveats
	for _, sDg := range stateDgs {
		for _, pDg := range planDgs {
			stateRegion := sDg.(basetypes.ObjectValue).Attributes()["region"]
			planRegion := pDg.(basetypes.ObjectValue).Attributes()["region"]
			if stateRegion.Equal(planRegion) {
				newPlan = append(newPlan, pDg)
			}
		}
	}

	// if the config/plan dgs count does not match with the expected count in the state dgs
	if len(newPlan) != len(stateDgs) {
		stateRegions := []attr.Value{}
		planRegions := []attr.Value{}
		for _, sDg := range stateDgs {
			stateRegion := sDg.(basetypes.ObjectValue).Attributes()["region"]
			stateRegions = append(stateRegions, stateRegion)
		}
		for _, pDg := range planDgs {
			planRegion := pDg.(basetypes.ObjectValue).Attributes()["region"]
			planRegions = append(planRegions, planRegion)
		}
		resp.Diagnostics.AddError("Regions in config not matching state", fmt.Sprintf("config regions: %v expected to match state regions: %v", planRegions, stateRegions))
		return
	}

	resp.PlanValue = basetypes.NewSetValueMust(newPlan[0].Type(ctx), newPlan)

	if len(stateDgs) > len(planDgs) {
		resp.Diagnostics.AddError("Upscaling not supported", "Upscaling data groups and witness groups currently not supported")
		return
	}

	for _, planDg := range planDgs {
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

		if stateDgKey == nil {
			resp.Diagnostics.AddWarning("Data group not found", fmt.Sprintf("data group with region %v not found", planDg.(basetypes.ObjectValue).Attributes()["region"].String()))
			continue
		}

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
			resp.Diagnostics.AddWarning("PG type changed", fmt.Sprintf("PG type changed from %v to %v",
				statePGType,
				planPGType))
		}

		// pg version
		planPGVersion := planDg.(basetypes.ObjectValue).Attributes()["pg_version"]
		statePGVersion := stateDgs[*stateDgKey].(basetypes.ObjectValue).Attributes()["pg_version"]

		if !planPGVersion.Equal(statePGVersion) {
			resp.Diagnostics.AddWarning("PG version changed", fmt.Sprintf("PG version changed from %v to %v",
				statePGVersion,
				planPGVersion))
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
			resp.Diagnostics.AddWarning("Cloud provider changed", fmt.Sprintf("Cloud provider changed from %v to %v",
				stateCloudProvider,
				planCloudProvider))
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
