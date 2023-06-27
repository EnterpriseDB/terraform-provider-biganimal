package pgd

import "github.com/EnterpriseDB/terraform-provider-biganimal/pkg/models"

type DataGroup struct {
	GroupId                    *string                  `json:"groupId,omitempty" tfsdk:"group_id"`
	AllowedIpRanges            *[]models.AllowedIpRange `json:"allowedIpRanges,omitempty" tfsdk:"allowed_ip_ranges"`
	BackupRetentionPeriod      *string                  `json:"backupRetentionPeriod,omitempty" tfsdk:"backup_retention_period"`
	ClusterArchitecture        *ClusterArchitecture     `json:"clusterArchitecture,omitempty" tfsdk:"cluster_architecture"`
	ClusterName                *string                  `json:"clusterName,omitempty" tfsdk:"cluster_name"`
	ClusterType                *string                  `json:"clusterType,omitempty" tfsdk:"cluster_type"`
	Connection                 *ClusterConnection       `json:"connection,omitempty" tfsdk:"connection_uri"`
	CreatedAt                  *PointInTime             `json:"createdAt,omitempty" tfsdk:"created_at"`
	CspAuth                    *bool                    `json:"cspAuth,omitempty" tfsdk:"csp_auth"`
	DeletedAt                  *PointInTime             `json:"deletedAt,omitempty" tfsdk:"deleted_at"`
	ExpiredAt                  *PointInTime             `json:"expiredAt,omitempty" tfsdk:"expired_at"`
	FirstRecoverabilityPointAt *PointInTime             `json:"firstRecoverabilityPointAt,omitempty" tfsdk:"first_recoverability_point_at"`
	InstanceType               *InstanceType            `json:"instanceType,omitempty" tfsdk:"instance_type"`
	LogsUrl                    *string                  `json:"logsUrl,omitempty" tfsdk:"logs_url"`
	MetricsUrl                 *string                  `json:"metricsUrl,omitempty" tfsdk:"metrics_url"`
	PgConfig                   *[]models.KeyValue       `json:"pgConfig,omitempty" tfsdk:"pg_config"`
	PgType                     *PgType                  `json:"pgType,omitempty" tfsdk:"pg_type"`
	PgVersion                  *PgVersion               `json:"pgVersion,omitempty" tfsdk:"pg_version"`
	Phase                      *string                  `json:"phase,omitempty" tfsdk:"phase"`
	PrivateNetworking          *bool                    `json:"privateNetworking,omitempty" tfsdk:"private_networking"`
	Provider                   *CloudProvider           `json:"provider,omitempty" tfsdk:"cloud_provider"`
	Region                     *Region                  `json:"region,omitempty" tfsdk:"region"`
	ResizingPvc                *[]string                `json:"resizingPvc,omitempty" tfsdk:"resizing_pvc"`
	Storage                    *models.Storage          `json:"storage,omitempty" tfsdk:"storage"`
}
