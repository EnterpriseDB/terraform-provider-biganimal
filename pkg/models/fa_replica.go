package models

import (
	"github.com/EnterpriseDB/terraform-provider-biganimal/pkg/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func NewFAReplica(d *schema.ResourceData) (*FAReplicaCluster, error) {
	allowedIpRanges, err := utils.StructFromProps[[]AllowedIpRange](d.Get("allowed_ip_ranges"))
	if err != nil {
		return nil, err
	}

	clusterArchitecture, err := utils.StructFromProps[Architecture](d.Get("cluster_architecture"))
	if err != nil {
		return nil, err
	}

	pgConfig, err := utils.StructFromProps[[]KeyValue](d.Get("pg_config"))
	if err != nil {
		return nil, err
	}

	storage, err := utils.StructFromProps[Storage](d.Get("storage"))
	if err != nil {
		return nil, err
	}

	cluster := &FAReplicaCluster{
		AllowedIpRanges:       &allowedIpRanges,
		BackupRetentionPeriod: utils.GetStringP(d, "backup_retention_period"),
		ClusterArchitecture:   &clusterArchitecture,
		ClusterId:             utils.GetStringP(d, "cluster_id"),
		ClusterName:           utils.GetStringP(d, "cluster_name"),
		CSPAuth:               utils.GetBoolP(d, "csp_auth"),

		//  these are readonly attributes, that come from the cluster api,
		// and end up in the resourceData.  we don't set these from the
		// resource data into the cluster

		// Conditions
		// CreatedAt
		// DeletedAt
		// ExpiredAt
		// FirstRecoverabilityPointAt
		// LogsUrl
		// MetricsUrl
		// Phase
		// ResizingPvc

		InstanceType: &InstanceType{
			InstanceTypeId: utils.GetString(d, "instance_type"),
		},
		Password: utils.GetStringP(d, "password"),
		PgConfig: &pgConfig,
		PgType: &PgType{
			PgTypeId: utils.GetString(d, "pg_type"),
		},
		PgVersion: &PgVersion{
			PgVersionId: d.Get("pg_version").(string),
		},
		PrivateNetworking: utils.GetBoolP(d, "private_networking"),
		Provider: &Provider{
			CloudProviderId: utils.GetString(d, "cloud_provider"),
		},
		Region: &Region{
			Id: utils.GetString(d, "region"),
		},
		ReadOnlyConnections: utils.GetBoolP(d, "read_only_connections"),
		Storage:             &storage,
	}

	return cluster, nil
}

// the following two methods create slightly different
// versions of clusters for write operations
// this is awkward, and should be replaced soon.
//
// we need to be able to unset this set of values.
// the api doesn't like being sent this information.
// because these fields are readonly in the api.
// we didn't see this when we were using
// the openapi because we had different struct types
// and these fields were omitted from some of those types

func NewFAReplicaForCreate(d *schema.ResourceData) (*FAReplicaCluster, error) {
	c, err := NewFAReplica(d)
	c.ClusterId = nil
	return c, err
}

func NewFAReplicaForUpdate(d *schema.ResourceData) (*FAReplicaCluster, error) {
	c, err := NewFAReplica(d)
	c.ClusterId = nil
	c.PgType = nil
	c.PgVersion = nil
	c.Provider = nil
	c.Region = nil
	return c, err
}

// everything is omitempty,
// and everything is either nullable, or empty-able
type FAReplicaCluster struct {
	AllowedIpRanges            *[]AllowedIpRange `json:"allowedIpRanges,omitempty"`
	BackupRetentionPeriod      *string           `json:"backupRetentionPeriod,omitempty"`
	ClusterArchitecture        *Architecture     `json:"clusterArchitecture,omitempty" mapstructure:"cluster_architecture"`
	ClusterId                  *string           `json:"clusterId,omitempty"`
	ClusterName                *string           `json:"clusterName,omitempty"`
	Conditions                 []Condition       `json:"conditions,omitempty"`
	CreatedAt                  *PointInTime      `json:"createdAt,omitempty"`
	CSPAuth                    *bool             `json:"cspAuth,omitempty"`
	DeletedAt                  *PointInTime      `json:"deletedAt,omitempty"`
	ExpiredAt                  *PointInTime      `json:"expiredAt,omitempty"`
	FirstRecoverabilityPointAt *PointInTime      `json:"firstRecoverabilityPointAt,omitempty"`
	InstanceType               *InstanceType     `json:"instanceType,omitempty"`
	LogsUrl                    *string           `json:"logsUrl,omitempty"`
	MetricsUrl                 *string           `json:"metricsUrl,omitempty"`
	Password                   *string           `json:"password,omitempty"`
	PgConfig                   *[]KeyValue       `json:"pgConfig,omitempty"`
	PgType                     *PgType           `json:"pgType,omitempty"`
	PgVersion                  *PgVersion        `json:"pgVersion,omitempty"`
	Phase                      *string           `json:"phase,omitempty"`
	PrivateNetworking          *bool             `json:"privateNetworking,omitempty"`
	Provider                   *Provider         `json:"provider,omitempty"`
	ReadOnlyConnections        *bool             `json:"readOnlyConnections,omitempty"`
	Region                     *Region           `json:"region,omitempty"`
	ResizingPvc                []string          `json:"resizingPvc,omitempty"`
	Storage                    *Storage          `json:"storage,omitempty"`
}

// IsHealthy checks to see if the cluster has the right condition 'biganimal.com/deployed'
// as well as the correct 'healthy' phase.  '
func (c FAReplicaCluster) IsHealthy() bool {
	if *c.Phase != PHASE_HEALTHY {
		return false
	}
	for _, cond := range c.Conditions {
		if *cond.Type_ == CONDITION_DEPLOYED && *cond.ConditionStatus == "True" {
			return true
		}
	}
	return false
}
