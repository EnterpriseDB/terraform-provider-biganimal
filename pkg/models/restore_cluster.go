package models

import (
	commonApi "github.com/EnterpriseDB/terraform-provider-biganimal/pkg/models/common/api"
)

type RestoreCluster struct {
	AllowedIpRanges       *[]AllowedIpRange            `json:"allowedIpRanges,omitempty"`
	BackupRetentionPeriod *string                      `json:"backupRetentionPeriod,omitempty"`
	ClusterArchitecture   *Architecture                `json:"clusterArchitecture,omitempty" mapstructure:"cluster_architecture"`
	ClusterName           *string                      `json:"clusterName,omitempty"`
	ClusterType           *string                      `json:"clusterType,omitempty"`
	CSPAuth               *bool                        `json:"cspAuth,omitempty"`
	InstanceType          *InstanceType                `json:"instanceType,omitempty"`
	Password              *string                      `json:"password,omitempty"`
	PgConfig              *[]KeyValue                  `json:"pgConfig,omitempty"`
	Phase                 *string                      `json:"phase,omitempty"`
	ReadOnlyConnections   *bool                        `json:"readOnlyConnections,omitempty"`
	Region                *Region                      `json:"region,omitempty"`
	ResizingPvc           []string                     `json:"resizingPvc,omitempty"`
	Storage               *Storage                     `json:"storage,omitempty"`
	MaintenanceWindow     *commonApi.MaintenanceWindow `json:"maintenanceWindow,omitempty"`
	ServiceAccountIds     *[]string                    `json:"serviceAccountIds,omitempty"`
	PeAllowedPrincipalIds *[]string                    `json:"peAllowedPrincipalIds,omitempty"`
	SuperuserAccess       *bool                        `json:"superuserAccess,omitempty"`
	RestorePoint          *string                      `json:"selectedRestorePointInTime,omitempty"`
}
