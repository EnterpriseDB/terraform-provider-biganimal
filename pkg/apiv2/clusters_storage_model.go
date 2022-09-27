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

type ClustersStorage struct {
	Iops *string `json:"iops,omitempty"`
	Size *string `json:"size,omitempty"`
	// Unused
	Throughput         *string `json:"throughput,omitempty"`
	VolumePropertiesId string  `json:"volumePropertiesId"`
	VolumeTypeId       string  `json:"volumeTypeId"`
}

func (m ClustersStorage) String() string {
	data, _ := json.MarshalIndent(m, "", "\t")
	return string(data)
}