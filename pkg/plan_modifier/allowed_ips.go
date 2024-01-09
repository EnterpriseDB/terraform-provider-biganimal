package plan_modifier

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

func CustomAllowedIps() planmodifier.Set {
	return customAllowedIpsModifier{}
}

// customAllowedIpsModifier implements the plan modifier.
type customAllowedIpsModifier struct{}

// Description returns a human-readable description of the plan modifier.
func (m customAllowedIpsModifier) Description(_ context.Context) string {
	return "Once set, the value of this attribute in state will not change."
}

// MarkdownDescription returns a markdown description of the plan modifier.
func (m customAllowedIpsModifier) MarkdownDescription(_ context.Context) string {
	return "Once set, the value of this attribute in state will not change."
}

// PlanModifySet implements the plan modification logic.
func (m customAllowedIpsModifier) PlanModifySet(ctx context.Context, req planmodifier.SetRequest, resp *planmodifier.SetResponse) {
	var planObject map[string]tftypes.Value

	err := req.Plan.Raw.As(&planObject)
	if err != nil {
		resp.Diagnostics.AddError("Mapping plan object in allowed ip ranges plan modifier error", err.Error())
		return
	}

	var privateNetworking bool
	err = planObject["private_networking"].As(&privateNetworking)
	if err != nil {
		resp.Diagnostics.AddError("Mapping private networking object in allowed ip ranges plan modifier error", err.Error())
		return
	}

	allowedIpRangesSetValueFunc := func(description string) basetypes.SetValue {
		defaultAttrs := map[string]attr.Value{"cidr_block": basetypes.NewStringValue("0.0.0.0/0"), "description": basetypes.NewStringValue(description)}
		defaultAttrTypes := map[string]attr.Type{"cidr_block": defaultAttrs["cidr_block"].Type(ctx), "description": defaultAttrs["description"].Type(ctx)}

		defaultObjectValue := basetypes.NewObjectValueMust(defaultAttrTypes, defaultAttrs)
		setOfObjects := []attr.Value{}
		setOfObjects = append(setOfObjects, defaultObjectValue)
		setValue := basetypes.NewSetValueMust(defaultObjectValue.Type(ctx), setOfObjects)

		return setValue
	}

	// if private networking set allowed IP ranges to cidr_block:"0.0.0.0/0" description:"To allow all access"
	if privateNetworking {
		resp.PlanValue = allowedIpRangesSetValueFunc("To allow all access")
		return
	}

	// if allowed IP ranges plan value is [] set allowed IP ranges cidr_block:"0.0.0.0/0" description:""
	if len(resp.PlanValue.Elements()) == 0 {
		resp.PlanValue = allowedIpRangesSetValueFunc("")
		return
	}

	if req.StateValue.IsNull() {
		return
	}

	// Do nothing if there is a known planned value.
	if !req.PlanValue.IsUnknown() {
		return
	}

	// Do nothing if there is an unknown configuration value, otherwise interpolation gets messed up.
	if req.ConfigValue.IsUnknown() {
		return
	}

	resp.PlanValue = req.StateValue
}
