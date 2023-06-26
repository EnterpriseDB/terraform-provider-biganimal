package pgd

import "github.com/EnterpriseDB/terraform-provider-biganimal/pkg/models"

type ClusterCreateDataGroup struct {
	AllowedIpRanges       []models.AllowedIpRange `json:"allowedIpRanges"`
	BackupRetentionPeriod *string                 `json:"backupRetentionPeriod,omitempty"`
	ClusterArchitecture   *ClusterArchitecture    `json:"clusterArchitecture"`
	ClusterType           string                  `json:"clusterType"`
	CspAuth               *bool                   `json:"cspAuth,omitempty"`
	InstanceType          *InstanceType           `json:"instanceType"`
	PgConfig              *[]models.KeyValue      `json:"pgConfig"`
	PgType                *PgType                 `json:"pgType"`
	PgVersion             *PgVersion              `json:"pgVersion"`
	PrivateNetworking     bool                    `json:"privateNetworking"`
	Provider              *CloudProvider          `json:"provider"`
	Region                *Region                 `json:"region"`
	Storage               *models.Storage         `json:"storage"`
}
