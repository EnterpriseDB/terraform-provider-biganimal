package pgd

import (
	"encoding/json"

	"github.com/EnterpriseDB/terraform-provider-biganimal/pkg/models"
)

type InstanceType string

// UnmarshalJSON to implement json.Unmarshaler for custom unmarshalling
func (recv *InstanceType) UnmarshalJSON(d []byte) error {
	var apiResult models.InstanceType
	if err := json.Unmarshal(d, &apiResult); err != nil {
		return err
	}

	instanceType := InstanceType(apiResult.InstanceTypeId)
	*recv = instanceType
	return nil
}
