package models

import (
	"fmt"
	"reflect"

	"github.com/EnterpriseDB/terraform-provider-biganimal/pkg/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/mitchellh/mapstructure"
)

func NewCluster(d *schema.ResourceData) (*Cluster, error) {
	allowedIpRanges, err := MakeThing[[]AllowedIpRange](d.Get("allowed_ip_ranges"))
	if err != nil {
		return nil, err
	}

	clusterArchitecture, err := MakeThing[Architecture](d.Get("cluster_architecture"))
	if err != nil {
		return nil, err
	}

	pgConfig, err := MakeThing[KeyValues](d.Get("pg_config"))
	if err != nil {
		return nil, err
	}

	storage, err := MakeThing[Storage](d.Get("storage"))
	if err != nil {
		return nil, err
	}

	cluster := &Cluster{
		AllowedIpRanges:       allowedIpRanges,
		BackupRetentionPeriod: utils.GetStringP(d, "backup_retention_period"),
		ClusterArchitecture:   &clusterArchitecture,
		ClusterId:             utils.GetStringP(d, "cluster_id"),
		ClusterName:           utils.GetStringP(d, "cluster_name"),

		//  these are readonly attributes, that come from the cluster api,
		// and end up in the resourceData.  we don't set these from the
		// resource data into the cluster

		// Conditions
		// CreatedAt
		// DeletedAt
		// ExpiredAt
		// FirstRecoverabilityPointAt
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
			RegionId: utils.GetString(d, "region"),
		},
		Replicas: utils.GetIntP(d, "replicas"),
		Storage:  &storage,
	}

	return cluster, nil
}

// everything is omitempty,
// and everything is either nullable, or empty-able
type Cluster struct {
	AllowedIpRanges            []AllowedIpRange `json:"allowedIpRanges,omitempty"`
	BackupRetentionPeriod      *string          `json:"backupRetentionPeriod,omitempty"`
	ClusterArchitecture        *Architecture    `json:"clusterArchitecture,omitempty" mapstructure:"cluster_architecture"`
	ClusterId                  *string          `json:"clusterId,omitempty"`
	ClusterName                *string          `json:"clusterName,omitempty"`
	Conditions                 []Condition      `json:"conditions,omitempty"`
	CreatedAt                  *PointInTime     `json:"createdAt,omitempty"`
	DeletedAt                  *PointInTime     `json:"deletedAt,omitempty"`
	ExpiredAt                  *PointInTime     `json:"expiredAt,omitempty"`
	FirstRecoverabilityPointAt *PointInTime     `json:"firstRecoverabilityPointAt,omitempty"`
	InstanceType               *InstanceType    `json:"instanceType,omitempty"`
	Password                   *string          `json:"password,omitempty"`
	PgConfig                   *KeyValues       `json:"pgConfig,omitempty"`
	PgType                     *PgType          `json:"pgType,omitempty"`
	PgVersion                  *PgVersion       `json:"pgVersion,omitempty"`
	Phase                      *string          `json:"phase,omitempty"`
	PrivateNetworking          *bool            `json:"privateNetworking,omitempty"`
	Provider                   *Provider        `json:"provider,omitempty"`
	Region                     *Region          `json:"region,omitempty"`
	Replicas                   *int             `json:"replicas,omitempty"`
	ResizingPvc                []string         `json:"resizingPvc,omitempty"`
	Storage                    *Storage         `json:"storage,omitempty"`
}

func (c Cluster) ArchitecturePropList() PropList {
	if c.ClusterArchitecture == nil {
		return PropList{}
	}
	propMap := map[string]interface{}{}
	propMap["id"] = c.ClusterArchitecture.ClusterArchitectureId
	propMap["name"] = c.ClusterArchitecture.ClusterArchitectureName
	propMap["nodes"] = c.ClusterArchitecture.Nodes

	return PropList{propMap}
}

func MakeThing[S any](blobs any) (S, error) {
	lst, ok := blobs.([]interface{})
	var s S
	var thing any

	if !ok {
		return s, fmt.Errorf("wrong type of block, need list, got %T", lst)
	}

	if reflect.TypeOf(s).Kind() == reflect.Slice || reflect.TypeOf(s).Kind() == reflect.Array {
		thing = lst
	} else {
		if len(lst) != 1 {
			return s, fmt.Errorf("%T needs exactly one block", s)
		}

		thing = lst[0]
	}

	err := mapstructure.Decode(thing, &s)
	return s, err
}
