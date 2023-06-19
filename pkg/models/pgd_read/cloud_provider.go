package pgd_read

import (
	"encoding/json"

	"github.com/EnterpriseDB/terraform-provider-biganimal/pkg/models"
)

type CloudProvider string

// UnmarshalJSON to implement json.Unmarshaler for custom unmarshalling
func (recv *CloudProvider) UnmarshalJSON(d []byte) error {
	var apiResult models.CloudProvider
	if err := json.Unmarshal(d, &apiResult); err != nil {
		return err
	}

	cloudProvider := CloudProvider(apiResult.CloudProviderId)
	*recv = cloudProvider
	return nil
}
