package models

import "github.com/EnterpriseDB/terraform-provider-biganimal/pkg/models/api"

// type ClusterDataGroup struct {
// 	GroupId                    string                                  `json:"groupId" tfsdk:"group_id"`
// 	AllowedIpRanges            *[]AllowedIpRange                       `json:"allowedIpRanges" tfsdk:"allowed_ip_ranges"`
// 	BackupRetentionPeriod      string                                  `json:"backupRetentionPeriod" tfsdk:"backup_retention_period"`
// 	ClusterArchitecture        *api.ClusterClusterArchitectureResponse `json:"clusterArchitecture,omitempty" tfsdk:"cluster_architecture"`
// 	ClusterName                string                                  `json:"clusterName" tfsdk:"cluster_name"`
// 	ClusterType                string                                  `json:"clusterType" tfsdk:"cluster_type"`
// 	Connection                 *ClusterConnection                      `json:"connection,omitempty" tfsdk:"connection_uri"`
// 	CreatedAt                  *PointInTime                            `json:"createdAt,omitempty" tfsdk:"created_at"`
// 	CspAuth                    *bool                                   `json:"cspAuth,omitempty" tfsdk:"csp_auth"`
// 	DeletedAt                  *PointInTime                            `json:"deletedAt,omitempty" tfsdk:"deleted_at"`
// 	ExpiredAt                  *PointInTime                            `json:"expiredAt,omitempty" tfsdk:"expired_at"`
// 	FirstRecoverabilityPointAt *PointInTime                            `json:"firstRecoverabilityPointAt,omitempty" tfsdk:"first_recoverability_point_at"`
// 	InstanceType               *InstanceType                           `json:"instanceType,omitempty" tfsdk:"instance_type"`
// 	LogsUrl                    *string                                 `json:"logsUrl,omitempty" tfsdk:"logs_url"`
// 	MetricsUrl                 *string                                 `json:"metricsUrl,omitempty" tfsdk:"metrics_url"`
// 	PgConfig                   *[]KeyValue                             `json:"pgConfig,omitempty" tfsdk:"pg_config"`
// 	PgType                     *PgType                                 `json:"pgType,omitempty" tfsdk:"pg_type"`
// 	PgVersion                  *PgVersion                              `json:"pgVersion,omitempty" tfsdk:"pg_version"`
// 	Phase                      string                                  `json:"phase" tfsdk:"phase"`
// 	PrivateNetworking          bool                                    `json:"privateNetworking" tfsdk:"private_networking"`
// 	Provider                   *CloudProvider                          `json:"provider,omitempty" tfsdk:"cloud_provider"`
// 	Region                     *Region                                 `json:"region,omitempty" tfsdk:"region"`
// 	ResizingPvc                *[]string                               `json:"resizingPvc,omitempty" tfsdk:"resizing_pvc"`
// 	Storage                    *Storage                                `json:"storage,omitempty" tfsdk:"storage"`
// }

type ClusterDataGroup struct {
	GroupId                    string                                  `json:"groupId" tfsdk:"group_id"`
	AllowedIpRanges            *[]AllowedIpRange                       `json:"allowedIpRanges" tfsdk:"allowed_ip_ranges"`
	BackupRetentionPeriod      string                                  `json:"backupRetentionPeriod" tfsdk:"backup_retention_period"`
	ClusterArchitecture        *api.ClusterClusterArchitectureResponse `json:"clusterArchitecture,omitempty" tfsdk:"cluster_architecture"`
	ClusterName                string                                  `json:"clusterName" tfsdk:"cluster_name"`
	ClusterType                string                                  `json:"clusterType" tfsdk:"cluster_type"`
	Connection                 *ClusterConnection                      `json:"connection,omitempty" tfsdk:"connection_uri"`
	CreatedAt                  *PointInTime                            `json:"createdAt,omitempty" tfsdk:"created_at"`
	CspAuth                    *bool                                   `json:"cspAuth,omitempty" tfsdk:"csp_auth"`
	DeletedAt                  *PointInTime                            `json:"deletedAt,omitempty" tfsdk:"deleted_at"`
	ExpiredAt                  *PointInTime                            `json:"expiredAt,omitempty" tfsdk:"expired_at"`
	FirstRecoverabilityPointAt *PointInTime                            `json:"firstRecoverabilityPointAt,omitempty" tfsdk:"first_recoverability_point_at"`
	InstanceType               *InstanceType                           `json:"instanceType,omitempty" tfsdk:"instance_type"`
	LogsUrl                    *string                                 `json:"logsUrl,omitempty" tfsdk:"logs_url"`
	MetricsUrl                 *string                                 `json:"metricsUrl,omitempty" tfsdk:"metrics_url"`
	PgConfig                   *[]KeyValue                             `json:"pgConfig,omitempty" tfsdk:"pg_config"`
	PgType                     *PgType                                 `json:"pgType,omitempty" tfsdk:"pg_type"`
	PgVersion                  *PgVersion                              `json:"pgVersion,omitempty" tfsdk:"pg_version"`
	Phase                      string                                  `json:"phase" tfsdk:"phase"`
	PrivateNetworking          bool                                    `json:"privateNetworking" tfsdk:"private_networking"`
	Provider                   *CloudProvider                          `json:"provider,omitempty" tfsdk:"cloud_provider"`
	Region                     *Region                                 `json:"region,omitempty" tfsdk:"region"`
	ResizingPvc                *[]string                               `json:"resizingPvc,omitempty" tfsdk:"resizing_pvc"`
	Storage                    *Storage                                `json:"storage,omitempty" tfsdk:"storage"`
}
