package pgd

import (
	"encoding/json"

	"github.com/EnterpriseDB/terraform-provider-biganimal/pkg/models"
)

type PgType struct {
	PgTypeId string `json:"pgTypeId" tfsdk:"pg_type_id"`
}

// UnmarshalJSON to implement json.Unmarshaler for custom unmarshalling
func (recv *PgType) UnmarshalJSON(d []byte) error {
	var apiResult models.PgType
	if err := json.Unmarshal(d, &apiResult); err != nil {
		return err
	}

	*recv = PgType{
		PgTypeId: apiResult.PgTypeId,
	}
	return nil
}
