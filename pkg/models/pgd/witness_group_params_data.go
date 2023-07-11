package pgd

import "github.com/EnterpriseDB/terraform-provider-biganimal/pkg/models"

type WitnessGroupParamsData struct {
	InstanceType *InstanceType   `json:"instanceType"`
	Storage      *models.Storage `json:"storage"`
}
