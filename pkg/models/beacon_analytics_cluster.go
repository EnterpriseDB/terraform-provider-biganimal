package models

import (
	commonApi "github.com/EnterpriseDB/terraform-provider-biganimal/pkg/models/common/api"
)

type BeaconAnalyticsCluster struct {
	ClusterType                *string                      `json:"clusterType,omitempty"`
	AllowedIpRanges            *[]AllowedIpRange            `json:"allowedIpRanges,omitempty"`
	BackupRetentionPeriod      *string                      `json:"backupRetentionPeriod,omitempty"`
	ClusterId                  *string                      `json:"clusterId,omitempty"`
	ClusterName                *string                      `json:"clusterName,omitempty"`
	Conditions                 []Condition                  `json:"conditions,omitempty"`
	CreatedAt                  *PointInTime                 `json:"createdAt,omitempty"`
	CSPAuth                    *bool                        `json:"cspAuth,omitempty"`
	DeletedAt                  *PointInTime                 `json:"deletedAt,omitempty"`
	ExpiredAt                  *PointInTime                 `json:"expiredAt,omitempty"`
	FirstRecoverabilityPointAt *PointInTime                 `json:"firstRecoverabilityPointAt,omitempty"`
	InstanceType               *InstanceType                `json:"instanceType,omitempty"`
	LogsUrl                    *string                      `json:"logsUrl,omitempty"`
	MetricsUrl                 *string                      `json:"metricsUrl,omitempty"`
	Password                   *string                      `json:"password,omitempty"`
	PgType                     *PgType                      `json:"pgType,omitempty"`
	PgVersion                  *PgVersion                   `json:"pgVersion,omitempty"`
	Phase                      *string                      `json:"phase,omitempty"`
	PrivateNetworking          *bool                        `json:"privateNetworking,omitempty"`
	Provider                   *Provider                    `json:"provider,omitempty"`
	Region                     *Region                      `json:"region,omitempty"`
	ResizingPvc                []string                     `json:"resizingPvc,omitempty"`
	MaintenanceWindow          *commonApi.MaintenanceWindow `json:"maintenanceWindow,omitempty"`
	ServiceAccountIds          *[]string                    `json:"serviceAccountIds,omitempty"`
	PeAllowedPrincipalIds      *[]string                    `json:"peAllowedPrincipalIds,omitempty"`
}
