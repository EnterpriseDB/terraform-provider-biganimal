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

type ClusterConnection struct {
	DatabaseName string `json:"databaseName"`
	Port         string `json:"port"`
	ServiceName  string `json:"serviceName"`
	Username     string `json:"username"`
}

func (m ClusterConnection) String() string {
	data, _ := json.MarshalIndent(m, "", "\t")
	return string(data)
}
