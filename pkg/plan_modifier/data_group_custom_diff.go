package plan_modifier

import (
	"context"
	"fmt"

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

// PlanModifyList implements the plan modification logic.
func (m customDataGroupDiffModifier) PlanModifySet(ctx context.Context, req planmodifier.SetRequest, resp *planmodifier.SetResponse) {
	planDgs := resp.PlanValue.Elements()
	stateDgs := req.StateValue.Elements()

	// newState := []attr.Value{}
	// for _, v := range configValue {
	// 	if dataGroupContains(stateValue, v) {
	// 		newState = append(newState, v)
	// 	}
	// }

	if len(stateDgs) > len(planDgs) {
		resp.Diagnostics.AddError("Upscaling not supported", "Upscaling data groups and witness groups currently not supported")
		return
	}

	for k, planDg := range planDgs {
		planBackupRetention := planDg.(basetypes.ObjectValue).Attributes()["backup_retention_period"].String()
		stateBackupRetention := stateDgs[k].(basetypes.ObjectValue).Attributes()["backup_retention_period"].String()

		if planBackupRetention != stateBackupRetention {
			resp.Diagnostics.AddWarning("Backup retention changed", fmt.Sprintf("backup retention period has changed from %v to %v",
				planBackupRetention,
				stateBackupRetention))
		}

	}

	_ = planDgs
	_ = stateDgs
}

// func dataGroupContains(s []attr.Value, e attr.Value) bool {
// 	for _, a := range s {
// 		aName := a.(basetypes.ObjectValue).Attributes()["name"].String()
// 		eName := e.(basetypes.ObjectValue).Attributes()["name"].String()
// 		if aName == eName {
// 			return true
// 		}
// 	}
// 	return false
// }
