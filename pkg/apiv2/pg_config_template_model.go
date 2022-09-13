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

type PgConfigTemplate struct {
	PgConfigTemplateId  string  `json:"pgConfigTemplateId"`
	PgTypeId            string  `json:"pgTypeId"`
	PgTypeName          *string `json:"pgTypeName,omitempty"`
	PgVersionId         string  `json:"pgVersionId"`
	PgVersionName       *string `json:"pgVersionName,omitempty"`
	TemplateDescription *string `json:"templateDescription,omitempty"`
	TemplateName        string  `json:"templateName"`
}

func (m PgConfigTemplate) String() string {
	data, _ := json.MarshalIndent(m, "", "\t")
	return string(data)
}
