package plan_modifier

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

func CustomClusterCloudProvider() planmodifier.String {
	return customCloudProviderModifier{}
}

type customCloudProviderModifier struct{}

func (m customCloudProviderModifier) Description(_ context.Context) string {
	return "Once set, the value of this attribute in state will not change."
}

func (m customCloudProviderModifier) MarkdownDescription(_ context.Context) string {
	return "Once set, the value of this attribute in state will not change."
}

func (m customCloudProviderModifier) PlanModifyString(ctx context.Context, req planmodifier.StringRequest, resp *planmodifier.StringResponse) {
	cloudProviderConfig := req.ConfigValue.ValueString()
	var configObject map[string]tftypes.Value

	err := req.Config.Raw.As(&configObject)
	if err != nil {
		resp.Diagnostics.AddError("Mapping config object in custom cloud provider modifier error", err.Error())
		return
	}

	if !strings.Contains(cloudProviderConfig, "bah") {
		peIds, ok := configObject["pe_allowed_principal_ids"]
		if ok && !peIds.IsNull() {
			resp.Diagnostics.AddError("your cloud account 'pe_allowed_principal_ids' field not allowed error",
				"field 'pe_allowed_principal_ids' should only be set if you are using BigAnimal's cloud account e.g. 'bah:aws', please remove 'pe_allowed_principal_ids'")
			return
		}

		saIds, ok := configObject["service_account_ids"]
		if ok && !saIds.IsNull() {
			resp.Diagnostics.AddError("your cloud account 'service_account_ids' field not allowed error",
				"field 'service_account_ids' should only be set if you are using BigAnimal's cloud account 'bah:gcp', please remove 'service_account_ids'")
			return
		}
	}

	if strings.Contains(cloudProviderConfig, "bah") && !strings.Contains(cloudProviderConfig, "bah:gcp") {
		saIds, ok := configObject["service_account_ids"]
		if ok && !saIds.IsNull() {
			resp.Diagnostics.AddError("your cloud account 'service_account_ids' field not allowed error",
				"you are not using cloud provider 'bah:gcp', field 'service_account_ids' should only be set if you are using cloud provider 'bah:gcp', please remove 'service_account_ids'")
			return
		}
	}
}
