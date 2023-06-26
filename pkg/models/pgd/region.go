package pgd

import (
	"encoding/json"

	"github.com/EnterpriseDB/terraform-provider-biganimal/pkg/models"
)

type Region struct {
	RegionId string `json:"regionId" tfsdk:"region_id"`
}

// UnmarshalJSON to implement json.Unmarshaler for custom unmarshalling
func (recv *Region) UnmarshalJSON(d []byte) error {
	var apiResult models.Region
	if err := json.Unmarshal(d, &apiResult); err != nil {
		return err
	}

	*recv = Region{
		RegionId: apiResult.Id,
	}
	return nil
}
