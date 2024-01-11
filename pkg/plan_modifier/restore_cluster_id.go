package plan_modifier

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
)

// CustomRestoreClusterId returns a plan modifier that copies a known prior state
// value into the planned value. Use this when it is known that an unconfigured
// value will remain the same after a resource update.
//
// To prevent Terraform errors, the framework automatically sets unconfigured
// and Computed attributes to an unknown value "(known after apply)" on update.
// Using this plan modifier will instead display the prior state value in the
// plan, unless a prior plan modifier adjusts the value.
func CustomRestoreClusterId() planmodifier.String {
	return customRestoreClusterIdModifier{}
}

// customRestoreClusterIdModifier implements the plan modifier.
type customRestoreClusterIdModifier struct{}

// Description returns a human-readable description of the plan modifier.
func (m customRestoreClusterIdModifier) Description(_ context.Context) string {
	return "Once set, the value of this attribute in state will not change."
}

// MarkdownDescription returns a markdown description of the plan modifier.
func (m customRestoreClusterIdModifier) MarkdownDescription(_ context.Context) string {
	return "Once set, the value of this attribute in state will not change."
}

// PlanModifyString implements the plan modification logic.
func (m customRestoreClusterIdModifier) PlanModifyString(ctx context.Context, req planmodifier.StringRequest, resp *planmodifier.StringResponse) {
	if !req.PlanValue.IsNull() {
		resp.Diagnostics.AddWarning(
			"You are restoring a cluster",
			"You are restoring a cluster.  The config fields 'cloud_provider', 'pg_type', 'pg_version' and 'private_networking' are unmodifiable and have to match the source cluster. After restoring the cluster, please remove fields 'restore_cluster_id','restore_from_deleted' and 'restore_point' from the config",
		)
	}
}
