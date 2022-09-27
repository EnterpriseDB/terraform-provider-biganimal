// Code generated by scripts/api/main.go. DO NOT EDIT.
/*
 * BigAnimal
 *
 * BigAnimal REST API v2 <br /><br /> Please visit [API v2 Changelog page](/api/docs/v2migration.html) for information about migrating from API v1.
 *
 * API version: 17.28.4
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package apiv2

import "encoding/json"

type ClusterIdRestoreBody1 struct {
	AllowedIpRanges            []AllowedIpRange                    `json:"allowedIpRanges"`
	ClusterArchitecture        *ClustersClusterArchitecture        `json:"clusterArchitecture,omitempty"`
	ClusterName                string                              `json:"clusterName"`
	InstanceType               *ClustersInstanceType               `json:"instanceType"`
	Password                   string                              `json:"password"`
	PgConfig                   []ClustersClusterArchitectureParams `json:"pgConfig"`
	Region                     *ClustersRegion                     `json:"region"`
	Replicas                   *float64                            `json:"replicas,omitempty"`
	SelectedRestorePointInTime *string                             `json:"selectedRestorePointInTime,omitempty"`
	Storage                    *ClustersStorage                    `json:"storage"`
}

func (m ClusterIdRestoreBody1) String() string {
	data, _ := json.MarshalIndent(m, "", "\t")
	return string(data)
}