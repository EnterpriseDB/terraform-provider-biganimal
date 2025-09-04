package plan_modifier

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func SetForceUnknownUpdate() planmodifier.Set {
	return forceUnknownModifier{}
}

// forceUnknownModifier implements the plan modifier.
type forceUnknownModifier struct{}

// Description returns a human-readable description of the plan modifier.
func (m forceUnknownModifier) Description(_ context.Context) string {
	return "Once set, the value of this attribute in state will not change."
}

// MarkdownDescription returns a markdown description of the plan modifier.
func (m forceUnknownModifier) MarkdownDescription(_ context.Context) string {
	return "Once set, the value of this attribute in state will not change."
}

// PlanModifySet implements the plan modification logic.
func (m forceUnknownModifier) PlanModifySet(ctx context.Context, req planmodifier.SetRequest, resp *planmodifier.SetResponse) {
	// Do nothing if there is no state value.
	// for create use plan value
	if req.StateValue.IsNull() {
		return
	}

	// Do nothing if there is a known planned value.
	// for update if not unknown or not null then use plan value
	if !req.PlanValue.IsUnknown() && !req.PlanValue.IsNull() {
		return
	}

	// if it reaches here, it typically means it sets the plan value to unknown if config is set to null
	resp.PlanValue = types.SetUnknown(req.PlanValue.ElementType(ctx))
}
