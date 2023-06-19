package pgd_read

import (
	"encoding/json"

	"github.com/EnterpriseDB/terraform-provider-biganimal/pkg/models"
)

type PgType string

// UnmarshalJSON to implement json.Unmarshaler for custom unmarshalling
func (recv *PgType) UnmarshalJSON(d []byte) error {
	var apiResult models.PgType
	if err := json.Unmarshal(d, &apiResult); err != nil {
		return err
	}

	pgType := PgType(apiResult.PgTypeId)
	*recv = pgType
	return nil
}
