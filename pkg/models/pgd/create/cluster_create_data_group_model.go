package create

import (
	"github.com/EnterpriseDB/terraform-provider-biganimal/pkg/models"
	"github.com/EnterpriseDB/terraform-provider-biganimal/pkg/models/pgd"
)

type ClusterCreateDataGroup struct {
	AllowedIpRanges       []models.AllowedIpRange            `json:"allowedIpRanges"`
	BackupRetentionPeriod *string                            `json:"backupRetentionPeriod,omitempty"`
	ClusterArchitecture   *ClusterClusterArchitecture        `json:"clusterArchitecture"`
	ClusterType           string                             `json:"clusterType"`
	CspAuth               *bool                              `json:"cspAuth,omitempty"`
	InstanceType          *ClusterInstanceType               `json:"instanceType"`
	MaintenanceWindow     *ClusterMaintenanceWindow          `json:"maintenanceWindow,omitempty"`
	PgConfig              []pgd.ArrayOfNameValueObjectsInner `json:"pgConfig"`
	PgType                *ClusterPgType                     `json:"pgType"`
	PgVersion             *ClusterPgVersion                  `json:"pgVersion"`
	PrivateNetworking     bool                               `json:"privateNetworking"`
	Provider              *ClusterCloudProvider              `json:"provider"`
	Region                *ClusterRegion                     `json:"region"`
	Storage               *ClusterStorage                    `json:"storage"`
}
