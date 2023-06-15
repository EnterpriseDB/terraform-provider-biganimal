package tf

import (
	"encoding/json"

	"github.com/EnterpriseDB/terraform-provider-biganimal/pkg/models/api"
)

type ClusterStorageResponse struct {
	Iops               string `tfsdk:"iops"`
	Size               string `tfsdk:"size"`
	Throughput         string `tfsdk:"throughput"`
	VolumePropertiesId string `tfsdk:"volume_properties"`
	VolumeTypeId       string `tfsdk:"volume_type"`
}

func (r ClusterStorageResponse) String() string {
	return r.VolumePropertiesId
}

// UnmarshalJSON to implement json.Unmarshaler for custom unmarshalling
func (recv *ClusterStorageResponse) UnmarshalJSON(d []byte) error {
	var apiResult api.ClusterStorageResponse
	if err := json.Unmarshal(d, &apiResult); err != nil {
		return err
	}

	clusterStorageResponse := ClusterStorageResponse{
		Iops:               apiResult.Iops,
		Size:               apiResult.Size,
		Throughput:         apiResult.Throughput,
		VolumePropertiesId: apiResult.VolumePropertiesId,
		VolumeTypeId:       apiResult.VolumeTypeId,
	}
	*recv = clusterStorageResponse
	return nil
}
