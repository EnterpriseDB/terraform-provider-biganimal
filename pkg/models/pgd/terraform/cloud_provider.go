package terraform

import "github.com/hashicorp/terraform-plugin-framework/types"

type CloudProvider struct {
	CloudProviderId types.String `tfsdk:"cloud_provider_id"`
}
