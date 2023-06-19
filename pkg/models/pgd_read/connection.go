package pgd_read

import (
	"encoding/json"

	"github.com/EnterpriseDB/terraform-provider-biganimal/pkg/models"
)

type ClusterConnection string

// UnmarshalJSON to implement json.Unmarshaler for custom unmarshalling
func (recv *ClusterConnection) UnmarshalJSON(d []byte) error {
	var apiResult models.ClusterConnection
	if err := json.Unmarshal(d, &apiResult); err != nil {
		return err
	}

	clusterConnection := ClusterConnection(apiResult.PgUri)
	*recv = clusterConnection
	return nil
}
