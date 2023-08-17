package plan_modifier

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func CustomPGConfig() planmodifier.Set {
	return customPGConfigModifier{}
}

// customPGConfigModifier implements the plan modifier.
type customPGConfigModifier struct{}

// Description returns a human-readable description of the plan modifier.
func (m customPGConfigModifier) Description(_ context.Context) string {
	return "Once set, the value of this attribute in state will not change."
}

// MarkdownDescription returns a markdown description of the plan modifier.
func (m customPGConfigModifier) MarkdownDescription(_ context.Context) string {
	return "Once set, the value of this attribute in state will not change."
}

// PlanModifyList implements the plan modification logic.
func (m customPGConfigModifier) PlanModifySet(ctx context.Context, req planmodifier.SetRequest, resp *planmodifier.SetResponse) {
	defaults := map[string]string{
		"autovacuum_max_workers":       "5",
		"autovacuum_vacuum_cost_limit": "3000",
		"checkpoint_completion_target": "0.9",
		"checkpoint_timeout":           "15min",
		"cpu_tuple_cost":               "0.03",
		"effective_cache_size":         "0.75 * ram",
		"maintenance_work_mem":         "(0.15 * (ram - shared_buffers) / autovacuum_max_workers) > 1GB ? 1GB : (0.15 * (ram - shared_buffers) / autovacuum_max_workers)",
		"random_page_cost":             "1.1",
		"shared_buffers":               "((0.25 * ram) > 80GB) ? 80GB : (0.25 * ram)",
		"tcp_keepalives_idle":          "120",
		"tcp_keepalives_interval":      "30",
		"wal_buffers":                  "64MB",
		"wal_compression":              "on",
	}

	elementTypeAttrTypes := req.StateValue.ElementType(ctx).(basetypes.ObjectType).AttrTypes
	if req.StateValue.IsNull() {
		setOfObjects := resp.PlanValue.Elements()

		for k, v := range defaults {
			if !pgConfigNameExists(resp.PlanValue.Elements(), k) {
				defaultAttrs := map[string]attr.Value{"name": basetypes.NewStringValue(k), "value": basetypes.NewStringValue(v)}
				defaultObjectValue := basetypes.NewObjectValueMust(elementTypeAttrTypes, defaultAttrs)
				setOfObjects = append(setOfObjects, defaultObjectValue)
			}
		}

		setValue := basetypes.NewSetValueMust(req.StateValue.ElementType(ctx), setOfObjects)
		resp.PlanValue = setValue
		return
	}

	if !req.StateValue.IsNull() {
		setOfObjects := resp.PlanValue.Elements()

		for _, v := range req.StateValue.Elements() {
			statePgConfigName := strings.Replace(v.(basetypes.ObjectValue).Attributes()["name"].String(), "\"", "", -1)
			if !pgConfigNameExists(resp.PlanValue.Elements(), statePgConfigName) {
				defaultObjectValue := basetypes.NewObjectValueMust(elementTypeAttrTypes, v.(basetypes.ObjectValue).Attributes())
				setOfObjects = append(setOfObjects, defaultObjectValue)
			}
		}

		setValue := basetypes.NewSetValueMust(req.StateValue.ElementType(ctx), setOfObjects)
		resp.PlanValue = setValue
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

func pgConfigNameExists(s []attr.Value, e string) bool {
	for _, a := range s {
		_, ok := a.(basetypes.ObjectValue).Attributes()[e]
		return ok
	}
	return false
}
