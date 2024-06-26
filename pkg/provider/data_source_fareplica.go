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

type FAReplicaData struct{}

func NewFAReplicaData() *FAReplicaData {
	return &FAReplicaData{}
}

func (c *FAReplicaData) Schema() *schema.Resource {
	return &schema.Resource{
		Description: "The faraway replica cluster data source describes a BigAnimal faraway replica connected to the cluster. The data source requires faraway replica cluster ID.",
		ReadContext: c.Read,
		Schema: map[string]*schema.Schema{
			"allowed_ip_ranges": {
				Description: "Allowed IP ranges.",
				Type:        schema.TypeSet,
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
			"source_cluster_id": {
				Description: "Source Cluster ID.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"cluster_type": {
				Description: "Type of the Specified Cluster .",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"cluster_name": {
				Description: "Name of the faraway replica cluster.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"most_recent": {
				Description: "Show the most recent cluster when there are multiple clusters with the same name.",
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
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
			"logs_url": {
				Description: "The URL to find the logs of this cluster.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"metrics_url": {
				Description: "The URL to find the metrics of this cluster.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"cluster_id": {
				Description: "Faraway Replica Cluster ID.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"connection_uri": {
				Description: "Cluster connection URI.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"pg_config": {
				Description: "Database configuration parameters.",
				Type:        schema.TypeSet,
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
			"project_id": {
				Description:      "BigAnimal Project ID.",
				Type:             schema.TypeString,
				Required:         true,
				ValidateDiagFunc: validateProjectId,
			},
			"cloud_provider": {
				Description: "Cloud provider.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"csp_auth": {
				Description: "Is authentication handled by the cloud service provider.",
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
			"pgvector": {
				Description: "Is pgvector extension enabled. Adds support for vector storage and vector similarity search to Postgres.",
				Computed:    true,
				Type:        schema.TypeBool,
			},
			"post_gis": {
				Description: "Is postGIS extension enabled. PostGIS extends the capabilities of the PostgreSQL relational database by adding support storing, indexing and querying geographic data.",
				Computed:    true,
				Type:        schema.TypeBool,
			},
		},
	}
}

func (c *FAReplicaData) Read(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	diags := diag.Diagnostics{}
	client := api.BuildAPI(meta).ClusterClient()

	clusterId, ok := d.Get("cluster_id").(string)
	if !ok {
		return diag.FromErr(errors.New("unable to find cluster ID"))
	}
	projectId := d.Get("project_id").(string)

	cluster, err := client.Read(ctx, projectId, clusterId)
	if err != nil {
		return fromBigAnimalErr(err)
	}
	tflog.Debug(ctx, pretty.Sprint(cluster))

	if *cluster.ClusterType != "faraway_replica" {
		return diag.FromErr(errors.New("specified cluster is not a 'faraway replica', please use 'biganimal_cluster' data source to fetch details about this cluster"))
	}

	connection, err := client.ConnectionString(ctx, projectId, *cluster.ClusterId)
	if err != nil {
		return diag.FromErr(err)
	}

	// set the outputs
	utils.SetOrPanic(d, "source_cluster_id", cluster.ReplicaSourceClusterId)
	utils.SetOrPanic(d, "cluster_type", cluster.ClusterType)
	utils.SetOrPanic(d, "allowed_ip_ranges", cluster.AllowedIpRanges)
	utils.SetOrPanic(d, "backup_retention_period", cluster.BackupRetentionPeriod)
	utils.SetOrPanic(d, "cluster_architecture", *cluster.ClusterArchitecture)
	utils.SetOrPanic(d, "created_at", cluster.CreatedAt)
	utils.SetOrPanic(d, "csp_auth", cluster.CSPAuth)
	utils.SetOrPanic(d, "deleted_at", cluster.DeletedAt)
	utils.SetOrPanic(d, "expired_at", cluster.ExpiredAt)
	utils.SetOrPanic(d, "cluster_name", cluster.ClusterName)
	utils.SetOrPanic(d, "first_recoverability_point_at", cluster.FirstRecoverabilityPointAt)
	utils.SetOrPanic(d, "instance_type", cluster.InstanceType)
	utils.SetOrPanic(d, "logs_url", cluster.LogsUrl)
	utils.SetOrPanic(d, "metrics_url", cluster.MetricsUrl)
	utils.SetOrPanic(d, "pg_config", cluster.PgConfig)
	utils.SetOrPanic(d, "pg_type", cluster.PgType)
	utils.SetOrPanic(d, "pg_version", cluster.PgVersion)
	utils.SetOrPanic(d, "phase", cluster.Phase)
	utils.SetOrPanic(d, "private_networking", cluster.PrivateNetworking)
	utils.SetOrPanic(d, "cloud_provider", cluster.Provider)
	utils.SetOrPanic(d, "region", cluster.Region)
	utils.SetOrPanic(d, "storage", *cluster.Storage)
	utils.SetOrPanic(d, "resizing_pvc", cluster.ResizingPvc)
	utils.SetOrPanic(d, "cluster_id", cluster.ClusterId)
	utils.SetOrPanic(d, "connection_uri", connection.PgUri)

	if *cluster.Extensions != nil {
		for _, v := range *cluster.Extensions {
			if v.Enabled && v.ExtensionId == "pgvector" {
				utils.SetOrPanic(d, "pgvector", true) // Computed
			}
			if v.Enabled && v.ExtensionId == "postgis" {
				utils.SetOrPanic(d, "post_gis", true) // Computed
			}
		}
	}

	d.SetId(*cluster.ClusterId)

	return diags
}
