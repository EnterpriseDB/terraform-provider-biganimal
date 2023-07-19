package plan_modifier

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
)

func CustomStringDiff() planmodifier.String {
	return customStringDiffModifier{}
}

// CustomStringDiffModifier implements the plan modifier.
type customStringDiffModifier struct{}

// Description returns a human-readable description of the plan modifier.
func (m customStringDiffModifier) Description(_ context.Context) string {
	return "Once set, the value of this attribute in state will not change."
}

// MarkdownDescription returns a markdown description of the plan modifier.
func (m customStringDiffModifier) MarkdownDescription(_ context.Context) string {
	return "Once set, the value of this attribute in state will not change."
}

// PlanModifyList implements the plan modification logic.
func (m customStringDiffModifier) PlanModifyString(ctx context.Context, req planmodifier.StringRequest, resp *planmodifier.StringResponse) {
	if !req.StateValue.IsNull() {
		bstate := req.StateValue.String()
		bplan := resp.PlanValue.String()

		if bplan != bstate {
			resp.Diagnostics.AddWarning("Backup retention changed", fmt.Sprintf("backup retention period has changed from %v to %v",
				bstate,
				bplan))
		}
	}
}
