package pgd

import (
	"encoding/json"

	"github.com/EnterpriseDB/terraform-provider-biganimal/pkg/models"
)

type CloudProvider struct {
	CloudProviderId string `json:"cloudProviderId" tfsdk:"cloud_provider_id"`
}

// UnmarshalJSON to implement json.Unmarshaler for custom unmarshalling
func (recv *CloudProvider) UnmarshalJSON(d []byte) error {
	var apiResult models.CloudProvider
	if err := json.Unmarshal(d, &apiResult); err != nil {
		return err
	}

	*recv = CloudProvider{
		CloudProviderId: apiResult.CloudProviderId,
	}
	return nil
}
