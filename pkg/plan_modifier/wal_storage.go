package plan_modifier

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func WalStorageForUnknown() planmodifier.Object {
	return WalStorageForUnknownModifier{}
}

// MaintenanceWindowForUnknownModifier implements the plan modifier.
type WalStorageForUnknownModifier struct{}

// Description returns a human-readable description of the plan modifier.
func (r WalStorageForUnknownModifier) Description(_ context.Context) string {
	return "Once set, the value of this attribute in state will not change."
}

// MarkdownDescription returns a markdown description of the plan modifier.
func (r WalStorageForUnknownModifier) MarkdownDescription(_ context.Context) string {
	return "Once set, the value of this attribute in state will not change."
}

// PlanModifyObject implements the plan modification logic.
func (r WalStorageForUnknownModifier) PlanModifyObject(ctx context.Context, req planmodifier.ObjectRequest, resp *planmodifier.ObjectResponse) {
	state := req.StateValue

	// use state for unknown
	if resp.PlanValue.IsUnknown() {
		resp.PlanValue = basetypes.NewObjectValueMust(state.AttributeTypes(ctx), state.Attributes())
	}
}
