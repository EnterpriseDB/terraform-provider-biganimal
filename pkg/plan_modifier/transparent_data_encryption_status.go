package plan_modifier

import (
	"context"

	"github.com/EnterpriseDB/terraform-provider-biganimal/pkg/constants"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
)

func CustomTDEStatus() planmodifier.String {
	return CustomTDEStatusModifier{}
}

// CustomTDEStatusModifier implements the plan modifier.
type CustomTDEStatusModifier struct{}

// Description returns a human-readable description of the plan modifier.
func (m CustomTDEStatusModifier) Description(_ context.Context) string {
	return "Once set, the value of this attribute in state will not change."
}

// MarkdownDescription returns a markdown description of the plan modifier.
func (m CustomTDEStatusModifier) MarkdownDescription(_ context.Context) string {
	return "Once set, the value of this attribute in state will not change."
}

// PlanModifyString implements the plan modification logic.
func (m CustomTDEStatusModifier) PlanModifyString(ctx context.Context, req planmodifier.StringRequest, resp *planmodifier.StringResponse) {
	if req.StateValue.ValueString() == constants.TDE_STATUS {
		resp.PlanValue = req.StateValue
	}
}
