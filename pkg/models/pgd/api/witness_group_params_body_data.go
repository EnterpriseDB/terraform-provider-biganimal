package api

import "github.com/EnterpriseDB/terraform-provider-biganimal/pkg/models"

type WitnessGroupParamsBodyData struct {
	InstanceType *InstanceType   `json:"instanceType,omitempty"`
	Provider     *CloudProvider  `json:"provider,omitempty"`
	Region       *Region         `json:"region,omitempty"`
	Storage      *models.Storage `json:"storage,omitempty"`
}
