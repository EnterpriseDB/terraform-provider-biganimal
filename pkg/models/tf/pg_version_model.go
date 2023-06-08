package tf

import (
	"encoding/json"

	"github.com/EnterpriseDB/terraform-provider-biganimal/pkg/models/api"
)

type PgVersion string

// UnmarshalJSON to implement json.Unmarshaler for custom unmarshalling
func (recv *PgVersion) UnmarshalJSON(d []byte) error {
	var apiResult api.PgVersion
	if err := json.Unmarshal(d, &apiResult); err != nil {
		return err
	}

	pgVersion := PgVersion(apiResult.PgVersionId)
	*recv = pgVersion
	return nil
}
