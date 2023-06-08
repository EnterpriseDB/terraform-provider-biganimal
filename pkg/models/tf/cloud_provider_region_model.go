package tf

import (
	"encoding/json"

	"github.com/EnterpriseDB/terraform-provider-biganimal/pkg/models/api"
)

type CloudProviderRegion string

// UnmarshalJSON to implement json.Unmarshaler for custom unmarshalling
func (recv *CloudProviderRegion) UnmarshalJSON(d []byte) error {
	var apiResult api.CloudProviderRegion
	if err := json.Unmarshal(d, &apiResult); err != nil {
		return err
	}

	cloudProviderRegion := CloudProviderRegion(apiResult.RegionId)
	*recv = cloudProviderRegion
	return nil
}
