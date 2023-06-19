package pgd_read

import (
	"encoding/json"

	"github.com/EnterpriseDB/terraform-provider-biganimal/pkg/models"
)

type Region string

// UnmarshalJSON to implement json.Unmarshaler for custom unmarshalling
func (recv *Region) UnmarshalJSON(d []byte) error {
	var apiResult models.Region
	if err := json.Unmarshal(d, &apiResult); err != nil {
		return err
	}

	region := Region(apiResult.Id)
	*recv = region
	return nil
}
