package provider

import (
	"context"
	"errors"
	"fmt"

	"github.com/EnterpriseDB/terraform-provider-biganimal/pkg/api"
	"github.com/EnterpriseDB/terraform-provider-biganimal/pkg/models"
	"github.com/EnterpriseDB/terraform-provider-biganimal/pkg/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/thoas/go-funk"
)

type ClusterData struct{}

func NewClusterData() *ClusterData {
	return &ClusterData{}
}

func (c *ClusterData) Schema() *schema.Resource {
	return &schema.Resource{
		Description: "Sample cluster data source in the BigAnimal terraform provider .",
		ReadContext: c.Read,
		Schema: map[string]*schema.Schema{
			"allowed_ip_ranges": {
				Description: "Allowed IP ranges",
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cidr_block": {
							Description: "CIDR Block",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"description": {
							Description: "CIDR Block Description",
							Type:        schema.TypeString,
							Computed:    true,
						},
					},
				},
			},
			"backup_retention_period": {
				Description: "Backup Retention Period.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"cluster_architecture": {
				Description: "Cluster Architecture",
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Description: "ID",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"name": {
							Description: "Name",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"nodes": {
							Description: "Node Count",
							Type:        schema.TypeInt,
							Computed:    true,
						},
					},
				},
			},
			"cluster_name": {
				Description: "Name of the cluster.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"created_at": {
				Description: "Cluster Creation Time",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"deleted_at": {
				Description: "Cluster Deletion Time",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"expired_at": {
				Description: "Cluster Expiry Time",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"first_recoverability_point_at": {
				Description: "Earliest Backup recover time",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"instance_type": {
				Description: "InstanceType",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"cluster_id": {
				Description: "cluster ID",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"pg_config": {
				Description: "Instance Type",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Description: "GUC Name",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"value": {
							Description: "GUC Value",
							Type:        schema.TypeString,
							Computed:    true,
						},
					},
				},
			},
			"pg_type": {
				Description: "Postgres type",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"pg_version": {
				Description: "Postgres version",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"phase": {
				Description: "Current Phase of the cluster.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"private_networking": {
				Description: "Is private networking enabled",
				Type:        schema.TypeBool,
				Computed:    true,
			},
			"cloud_provider": {
				Description: "Cloud Provider",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"region": {
				Description: "Region",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"replicas": {
				Description: "Replicas",
				Type:        schema.TypeInt,
				Computed:    true,
			},
			"resizing_pvc": {
				Description: "Resizing PVC",
				Type:        schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Computed: true,
			},
			"storage": {
				Description: "Storage",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"iops": {
							Description: "IOPS",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"size": {
							Description: "Size",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"throughput": {
							Description: "Throughput",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"volume_properties": {
							Description: "Volume Properties",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"volume_type": {
							Description: "Volume Type",
							Type:        schema.TypeString,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func (c *ClusterData) Read(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	diags := diag.Diagnostics{}
	client := api.BuildAPI(meta).ClusterClient()

	clusterName, ok := d.Get("cluster_name").(string)
	if !ok {
		return diag.FromErr(errors.New("Unable to find cluster name"))
	}

	cluster, err := client.ReadByName(ctx, clusterName)
	if err != nil {
		return diag.FromErr(err)
	}

	// set the outputs
	SetOrPanic(d, "backup_retention_period", cluster.BackupRetentionPeriod)
	SetOrPanic(d, "cluster_architecture", cluster.ArchitecturePropList())
	SetOrPanic(d, "created_at", cluster.CreatedAt)
	SetOrPanic(d, "deleted_at", cluster.DeletedAt)
	SetOrPanic(d, "expired_at", cluster.ExpiredAt)
	SetOrPanic(d, "cluster_name", cluster.ClusterName)
	SetOrPanic(d, "first_recoverability_point_at", cluster.FirstRecoverabilityPointAt)
	SetOrPanic(d, "instance_type", cluster.InstanceType)
	SetOrPanic(d, "pg_config", cluster.PgConfig.PropList())
	SetOrPanic(d, "pg_type", cluster.PgType)
	SetOrPanic(d, "pg_version", cluster.PgType)
	SetOrPanic(d, "phase", cluster.Phase)
	SetOrPanic(d, "private_networking", cluster.PrivateNetworking)
	SetOrPanic(d, "cloud_provider", cluster.Provider)
	SetOrPanic(d, "region", cluster.Region)
	SetOrPanic(d, "replicas", cluster.Replicas)
	SetOrPanic(d, "storage", cluster.Storage.PropList())
	SetOrPanic(d, "resizing_pvc", cluster.ResizingPvc)
	SetOrPanic(d, "cluster_id", cluster.ClusterId)

	d.SetId(*cluster.ClusterId)

	return diags
}

func SetOrPanic(d *schema.ResourceData, key string, value interface{}) {
	if funk.IsEmpty(value) {
		return // empty value
	}

	var err error
	switch v := value.(type) {
	case []interface{}, models.PropList:
		err = d.Set(key, v)
	case int, bool, string:
		err = d.Set(key, v)
	case *float64:
		err = d.Set(key, *v)
	case *int:
		err = d.Set(key, *v)
	case *bool:
		err = d.Set(key, *v)
	case *string:
		err = d.Set(key, utils.DerefString(v))

	default:
		stringer, ok := value.(fmt.Stringer)
		if !ok {
			panic(fmt.Sprintf(" don't know how to handle %T", value))
		}

		err = d.Set(key, stringer.String())
	}

	// if d.Set fails
	if err != nil {
		panic(err)
	}
}