package tf

import (
	"encoding/json"

	"github.com/EnterpriseDB/terraform-provider-biganimal/pkg/models/api"
)

type CloudProvider string

// UnmarshalJSON to implement json.Unmarshaler for custom unmarshalling
func (recv *CloudProvider) UnmarshalJSON(d []byte) error {
	var apiResult api.CloudProvider
	if err := json.Unmarshal(d, &apiResult); err != nil {
		return err
	}

	cloudProvider := CloudProvider(apiResult.CloudProviderId)
	*recv = cloudProvider
	return nil
}
