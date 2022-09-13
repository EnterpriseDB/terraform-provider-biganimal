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

type ClusterUsageClusterArchitecture struct {
	ClusterArchitectureId   *string `json:"clusterArchitectureId,omitempty"`
	ClusterArchitectureName *string `json:"clusterArchitectureName,omitempty"`
}

func (m ClusterUsageClusterArchitecture) String() string {
	data, _ := json.MarshalIndent(m, "", "\t")
	return string(data)
}
