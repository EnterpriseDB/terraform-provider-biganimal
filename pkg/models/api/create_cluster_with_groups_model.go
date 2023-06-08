package api

import "encoding/json"

type CreateClusterWithGroups struct {
	ClusterName string                          `json:"clusterName"`
	ClusterType string                          `json:"clusterType"`
	Groups      []AnyOfclusterCreateGroupsItems `json:"groups"`
	Password    string                          `json:"password"`
}

func (m CreateClusterWithGroups) String() string {
	data, _ := json.MarshalIndent(m, "", "\t")
	return string(data)
}
