package create

import (
	"github.com/EnterpriseDB/terraform-provider-biganimal/pkg/models"
	"github.com/EnterpriseDB/terraform-provider-biganimal/pkg/models/pgd"
)

type ClusterCreateDataGroup struct {
	AllowedIpRanges       []models.AllowedIpRange            `json:"allowedIpRanges" tfsdk:"allowed_ip_ranges"`
	BackupRetentionPeriod *string                            `json:"backupRetentionPeriod,omitempty" tfsdk:"backup_retention_period"`
	ClusterArchitecture   *ClusterClusterArchitecture        `json:"clusterArchitecture" tfsdk:"cluster_architecture"`
	ClusterType           string                             `json:"clusterType" tfsdk:"cluster_type"`
	CspAuth               *bool                              `json:"cspAuth,omitempty" tfsdk:"csp_auth"`
	InstanceType          *ClusterInstanceType               `json:"instanceType" tfsdk:"instance_type"`
	MaintenanceWindow     *ClusterMaintenanceWindow          `json:"maintenanceWindow,omitempty" tfsdk:"maintenance_window"`
	PgConfig              []pgd.ArrayOfNameValueObjectsInner `json:"pgConfig" tfsdk:"pg_config"`
	PgType                *ClusterPgType                     `json:"pgType" tfsdk:"pg_type"`
	PgVersion             *ClusterPgVersion                  `json:"pgVersion" tfsdk:"pg_version"`
	PrivateNetworking     bool                               `json:"privateNetworking" tfsdk:"private_networking"`
	Provider              *ClusterCloudProvider              `json:"provider" tfsdk:"provider"`
	Region                *ClusterRegion                     `json:"region" tfsdk:"region"`
	Storage               *ClusterStorage                    `json:"storage" tfsdk:"storage"`
}
