package pgd

import (
	"encoding/json"

	"github.com/EnterpriseDB/terraform-provider-biganimal/pkg/models"
)

type PgVersion struct {
	PgVersionId string `json:"pgVersionId" tfsdk:"pg_version_id"`
}

// UnmarshalJSON to implement json.Unmarshaler for custom unmarshalling
func (recv *PgVersion) UnmarshalJSON(d []byte) error {
	var apiResult models.PgVersion
	if err := json.Unmarshal(d, &apiResult); err != nil {
		return err
	}

	*recv = PgVersion{
		PgVersionId: apiResult.PgVersionId,
	}
	return nil
}
