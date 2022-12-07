package provider

import (
	"context"
	"errors"

	"github.com/EnterpriseDB/terraform-provider-biganimal/pkg/api"
	"github.com/EnterpriseDB/terraform-provider-biganimal/pkg/utils"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/kr/pretty"
)

type ClusterData struct{}

func NewClusterData() *ClusterData {
	return &ClusterData{}
}

func (c *ClusterData) Schema() *schema.Resource {
	return &schema.Resource{
		Description: "The cluster data source describes a BigAnimal cluster. The data source requires your cluster name.",
		ReadContext: c.Read,
		Schema: map[string]*schema.Schema{
			"allowed_ip_ranges": {
				Description: "Allowed IP ranges.",
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cidr_block": {
							Description: "CIDR block.",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"description": {
							Description: "CIDR block description.",
							Type:        schema.TypeString,
							Computed:    true,
						},
					},
				},
			},
			"backup_retention_period": {
				Description: "Backup retention period.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"cluster_architecture": {
				Description: "Cluster architecture.",
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Description: "Cluster architecture ID.",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"name": {
							Description: "Name.",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"nodes": {
							Description: "Node count.",
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
				Description: "Cluster creation time.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"deleted_at": {
				Description: "Cluster deletion time.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"expired_at": {
				Description: "Cluster expiry time.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"first_recoverability_point_at": {
				Description: "Earliest backup recover time.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"instance_type": {
				Description: "Instance type.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"cluster_id": {
				Description: "Cluster ID.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"connection_uri": {
				Description: "Cluster connection URI.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"ro_connection_uri": {
				Description: "Cluster read-only connection URI. Only available for high availability clusters.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"pg_config": {
				Description: "Database configuration parameters.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Description: "GUC name.",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"value": {
							Description: "GUC value.",
							Type:        schema.TypeString,
							Computed:    true,
						},
					},
				},
			},
			"pg_type": {
				Description: "Postgres type.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"pg_version": {
				Description: "Postgres version.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"phase": {
				Description: "Current phase of the cluster.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"private_networking": {
				Description: "Is private networking enabled.",
				Type:        schema.TypeBool,
				Computed:    true,
			},
			"cloud_provider": {
				Description: "Cloud provider.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"read_only_connections": {
				Description: "Is read only connection enabled.",
				Type:        schema.TypeBool,
				Computed:    true,
			},
			"region": {
				Description: "Region to deploy the cluster.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"resizing_pvc": {
				Description: "Resizing PVC.",
				Type:        schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Computed: true,
			},
			"storage": {
				Description: "Storage.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"iops": {
							Description: "IOPS for the selected volume.",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"size": {
							Description: "Size of the volume.",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"throughput": {
							Description: "Throughput.",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"volume_properties": {
							Description: "Volume properties.",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"volume_type": {
							Description: "Volume type.",
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
		return BigAnimalFromErr(err)
	}
	tflog.Debug(ctx, pretty.Sprint(cluster))

	connection, err := client.ConnectionString(ctx, *cluster.ClusterId)
	if err != nil {
		return diag.FromErr(err)
	}

	// set the outputs
	utils.SetOrPanic(d, "allowed_ip_ranges", cluster.AllowedIpRanges)
	utils.SetOrPanic(d, "backup_retention_period", cluster.BackupRetentionPeriod)
	utils.SetOrPanic(d, "cluster_architecture", *cluster.ClusterArchitecture)
	utils.SetOrPanic(d, "created_at", cluster.CreatedAt)
	utils.SetOrPanic(d, "deleted_at", cluster.DeletedAt)
	utils.SetOrPanic(d, "expired_at", cluster.ExpiredAt)
	utils.SetOrPanic(d, "cluster_name", cluster.ClusterName)
	utils.SetOrPanic(d, "first_recoverability_point_at", cluster.FirstRecoverabilityPointAt)
	utils.SetOrPanic(d, "instance_type", cluster.InstanceType)
	utils.SetOrPanic(d, "pg_config", cluster.PgConfig)
	utils.SetOrPanic(d, "pg_type", cluster.PgType)
	utils.SetOrPanic(d, "pg_version", cluster.PgVersion)
	utils.SetOrPanic(d, "phase", cluster.Phase)
	utils.SetOrPanic(d, "private_networking", cluster.PrivateNetworking)
	utils.SetOrPanic(d, "cloud_provider", cluster.Provider)
	utils.SetOrPanic(d, "region", cluster.Region)
	utils.SetOrPanic(d, "storage", *cluster.Storage)
	utils.SetOrPanic(d, "read_only_connections", cluster.ReadOnlyConnections)
	utils.SetOrPanic(d, "resizing_pvc", cluster.ResizingPvc)
	utils.SetOrPanic(d, "cluster_id", cluster.ClusterId)
	utils.SetOrPanic(d, "connection_uri", connection.PgUri)
	utils.SetOrPanic(d, "ro_connection_uri", connection.ReadOnlyPgUri)

	d.SetId(*cluster.ClusterId)

	return diags
}
