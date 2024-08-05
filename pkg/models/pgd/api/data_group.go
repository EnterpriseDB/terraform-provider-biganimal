package api

import "github.com/EnterpriseDB/terraform-provider-biganimal/pkg/models"

type DataGroup struct {
	GroupId               *string                   `json:"groupId,omitempty"`
	AllowedIpRanges       *[]models.AllowedIpRange  `json:"allowedIpRanges,omitempty"`
	BackupRetentionPeriod *string                   `json:"backupRetentionPeriod,omitempty"`
	ClusterArchitecture   *ClusterArchitecture      `json:"clusterArchitecture,omitempty"`
	ClusterName           *string                   `json:"clusterName,omitempty"`
	ClusterType           *string                   `json:"clusterType,omitempty"`
	Conditions            *[]Condition              `json:"conditions,omitempty"`
	Connection            *ClusterConnection        `json:"connection,omitempty"`
	CreatedAt             *PointInTime              `json:"createdAt,omitempty"`
	CspAuth               *bool                     `json:"cspAuth,omitempty"`
	InstanceType          *InstanceType             `json:"instanceType,omitempty"`
	LogsUrl               *string                   `json:"logsUrl,omitempty"`
	MetricsUrl            *string                   `json:"metricsUrl,omitempty"`
	PgConfig              *[]models.KeyValue        `json:"pgConfig,omitempty"`
	PgType                *PgType                   `json:"pgType,omitempty"`
	PgVersion             *PgVersion                `json:"pgVersion,omitempty"`
	Phase                 *string                   `json:"phase,omitempty"`
	PrivateNetworking     *bool                     `json:"privateNetworking,omitempty"`
	Provider              *CloudProvider            `json:"provider,omitempty"`
	Region                *Region                   `json:"region,omitempty"`
	ResizingPvc           *[]string                 `json:"resizingPvc,omitempty"`
	Storage               *models.Storage           `json:"storage,omitempty"`
	MaintenanceWindow     *models.MaintenanceWindow `json:"maintenanceWindow,omitempty"`
	ServiceAccountIds     *[]string                 `json:"serviceAccountIds,omitempty"`
	PeAllowedPrincipalIds *[]string                 `json:"peAllowedPrincipalIds,omitempty"`
	RoConnectionUri       *string                   `json:"roConnectionUri,omitempty"`
}
