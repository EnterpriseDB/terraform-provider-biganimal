package provider

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/EnterpriseDB/terraform-provider-biganimal/pkg/api"
	"github.com/EnterpriseDB/terraform-provider-biganimal/pkg/models"
	"github.com/EnterpriseDB/terraform-provider-biganimal/pkg/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type FAReplicaResource struct{}

func NewFAReplicaResource() *FAReplicaResource {
	return &FAReplicaResource{}
}

func (c *FAReplicaResource) Schema() *schema.Resource {
	return &schema.Resource{
		Description: "The faraway replica resource is used to manage cluster faraway-replicas on different active regions in the cloud. See [Managing replicas](https://www.enterprisedb.com/docs/biganimal/latest/using_cluster/managing_replicas/) for more details.",

		CreateContext: c.Create,
		ReadContext:   c.Read,
		UpdateContext: c.Update,
		DeleteContext: c.Delete,

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(45 * time.Minute),
		},

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
							Required:    true,
						},
						"description": {
							Description: "CIDR block description.",
							Type:        schema.TypeString,
							Optional:    true,
						},
					},
				},
			},
			"backup_retention_period": {
				Description: "Backup retention period. For example, \"7d\", \"2w\", or \"3m\".",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"csp_auth": {
				Description: "Is authentication handled by the cloud service provider. Available for AWS only, See [Authentication](https://www.enterprisedb.com/docs/biganimal/latest/getting_started/creating_a_cluster/#authentication) for details.",
				Type:        schema.TypeBool,
				Optional:    true,
			},

			"source_cluster_id": {
				Description: "Source cluster ID.",
				Type:        schema.TypeString,
				Required:    true,
			},

			"cluster_type": {
				Description: "Type of the cluster. For example, \"cluster\" for biganimal_cluster resources, or \"faraway_replica\" for biganimal_faraway_replica resources.",
				Type:        schema.TypeString,
				Computed:    true,
			},

			"cluster_name": {
				Description: "Name of the faraway replica cluster.",
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
			//"first_recoverability_point_at": {
			//	Description: "Earliest backup recover time.",
			//	Type:        schema.TypeString,
			//	Computed:    true,
			//},
			"instance_type": {
				Description: "Instance type. For example, \"azure:Standard_D2s_v3\", \"aws:c5.large\" or \"gcp:e2-highcpu-4\".",
				Type:        schema.TypeString,
				Required:    true,
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
				Description: "Cluster ID.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"connection_uri": {
				Description: "Cluster connection URI.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"pg_config": {
				Description: "Database configuration parameters. See [Modifying database configuration parameters](https://www.enterprisedb.com/docs/biganimal/latest/using_cluster/03_modifying_your_cluster/05_db_configuration_parameters/) for details.",
				Type:        schema.TypeSet,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Description: "GUC name.",
							Type:        schema.TypeString,
							Required:    true,
						},
						"value": {
							Description: "GUC value.",
							Type:        schema.TypeString,
							Required:    true,
						},
					},
				},
			},
			"phase": {
				Description: "Current phase of the cluster.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"private_networking": {
				Description: "Is private networking enabled.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"project_id": {
				Description:      "BigAnimal Project ID.",
				Type:             schema.TypeString,
				Required:         true,
				ValidateDiagFunc: validateProjectId,
			},
			"region": {
				Description: "Region to deploy the cluster. See [Supported regions](https://www.enterprisedb.com/docs/biganimal/latest/overview/03a_region_support/) for supported regions.",
				Type:        schema.TypeString,
				Required:    true,
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
				Required:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"iops": {
							Description: "IOPS for the selected volume. It can be set to different values depending on your volume type and properties.",
							Type:        schema.TypeString,
							Optional:    true,
							DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
								// iops is an optional field.
								// If there is already a value set (old != "")
								// and there is no new value (new == ""),
								// we can suppress this Diff
								return new == "" && old != ""
							},
						},
						"size": {
							Description: "Size of the volume. It can be set to different values depending on your volume type and properties.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"throughput": {
							Description: "Throughput is automatically calculated by BigAnimal based on the IOPS input.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"volume_properties": {
							Description: "Volume properties in accordance with the selected volume type.",
							Type:        schema.TypeString,
							Required:    true,
						},
						"volume_type": {
							Description: "Volume type. For Azure: \"azurepremiumstorage\" or \"ultradisk\". For AWS: \"gp3\", \"io2\", or \"io2-block-express\". For Google Cloud: only \"pd-ssd\".",
							Type:        schema.TypeString,
							Required:    true,
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

func (c *FAReplicaResource) Create(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := api.BuildAPI(meta).ClusterClient()

	err := d.Set("cluster_type", "faraway_replica")
	if err != nil {
		return diag.FromErr(err)
	}
	cluster, err := models.NewClusterForCreate(d)
	if err != nil {
		return diag.FromErr(err)
	}

	projectId := d.Get("project_id").(string)

	clusterId, err := client.Create(ctx, projectId, *cluster)
	if err != nil {
		return fromBigAnimalErr(err)
	}

	d.SetId(clusterId)

	// retry until we get success
	err = retry.RetryContext(
		ctx,
		d.Timeout(schema.TimeoutCreate)-time.Minute,
		c.retryFunc(ctx, d, meta, clusterId))
	if err != nil {
		return diag.FromErr(err)
	}
	return diag.Diagnostics{}
}

func (c *FAReplicaResource) Read(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	if err := c.read(ctx, d, meta); err != nil {
		return fromBigAnimalErr(err)
	}
	return diag.Diagnostics{}
}

func (c *FAReplicaResource) read(ctx context.Context, d *schema.ResourceData, meta any) error {
	client := api.BuildAPI(meta).ClusterClient()

	clusterId := d.Id()
	projectId := d.Get("project_id").(string)
	cluster, err := client.Read(ctx, projectId, clusterId)

	// return error if faraway-replica is promoted to a cluster
	if *cluster.ClusterType != "faraway_replica" {
		return fmt.Errorf("the specified cluster is no longer a 'faraway replica' and has likely been promoted to a standalone cluster. Please use the 'biganimal_cluster' data source to retrieve information about this cluster")
	}

	if err != nil {
		return err
	}

	connection, err := client.ConnectionString(ctx, projectId, clusterId)
	if err != nil {
		return err
	}

	// set the outputs
	utils.SetOrPanic(d, "cluster_type", "faraway_replica")
	utils.SetOrPanic(d, "source_cluster_id", cluster.ReplicaSourceClusterId)      // Required
	utils.SetOrPanic(d, "allowed_ip_ranges", cluster.AllowedIpRanges)             // optional
	utils.SetOrPanic(d, "backup_retention_period", cluster.BackupRetentionPeriod) // optional
	utils.SetOrPanic(d, "created_at", cluster.CreatedAt)                          // Computed
	utils.SetOrPanic(d, "deleted_at", cluster.DeletedAt)                          // Computed
	utils.SetOrPanic(d, "expired_at", cluster.ExpiredAt)                          // Computed
	utils.SetOrPanic(d, "cluster_name", cluster.ClusterName)                      // Required
	// utils.SetOrPanic(d, "first_recoverability_point_at", cluster.FirstRecoverabilityPointAt) // Computed
	utils.SetOrPanic(d, "instance_type", cluster.InstanceType)           // Required
	utils.SetOrPanic(d, "logs_url", cluster.LogsUrl)                     // Computed
	utils.SetOrPanic(d, "metrics_url", cluster.MetricsUrl)               // Computed
	utils.SetOrPanic(d, "pg_config", cluster.PgConfig)                   // Optional
	utils.SetOrPanic(d, "phase", cluster.Phase)                          // Computed
	utils.SetOrPanic(d, "private_networking", cluster.PrivateNetworking) // Optional

	utils.SetOrPanic(d, "csp_auth", cluster.CSPAuth)         // optional
	utils.SetOrPanic(d, "region", cluster.Region)            // Required
	utils.SetOrPanic(d, "storage", cluster.Storage)          // Required
	utils.SetOrPanic(d, "resizing_pvc", cluster.ResizingPvc) // Computed
	utils.SetOrPanic(d, "cluster_id", cluster.ClusterId)     // Computed
	utils.SetOrPanic(d, "connection_uri", connection.PgUri)  // Computed

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
	return nil
}

func (c *FAReplicaResource) Update(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := api.BuildAPI(meta).ClusterClient()
	err := d.Set("cluster_type", "faraway_replica")
	if err != nil {
		return diag.FromErr(err)
	}
	// short circuit early for these types of changes
	if d.HasChange("region") {
		return diag.FromErr(errors.New("region is immutable"))
	}

	if d.HasChange("backup_retention_period") {
		return diag.FromErr(errors.New("backup retention period is immutable"))
	}

	cluster, err := models.NewClusterForUpdate(d)
	if err != nil {
		return diag.FromErr(err)
	}

	clusterId := d.Id()
	projectId := d.Get("project_id").(string)
	_, err = client.Update(ctx, cluster, projectId, clusterId)
	if err != nil {
		return fromBigAnimalErr(err)
	}

	// retry until we get success
	err = retry.RetryContext(
		ctx,
		d.Timeout(schema.TimeoutUpdate)-time.Minute,
		c.retryFunc(ctx, d, meta, clusterId))
	if err != nil {
		return diag.FromErr(err)
	}
	return diag.Diagnostics{}
}

func (c *FAReplicaResource) Delete(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := api.BuildAPI(meta).ClusterClient()
	clusterId := d.Id()
	projectId := d.Get("project_id").(string)
	if err := client.Delete(ctx, projectId, clusterId); err != nil {
		return fromBigAnimalErr(err)
	}
	return diag.Diagnostics{}
}

func (c *FAReplicaResource) retryFunc(ctx context.Context, d *schema.ResourceData, meta any, clusterId string) retry.RetryFunc {
	client := api.BuildAPI(meta).ClusterClient()
	return func() *retry.RetryError {
		projectId := d.Get("project_id").(string)
		cluster, err := client.Read(ctx, projectId, clusterId)
		if err != nil {
			return retry.NonRetryableError(fmt.Errorf("error describing instance: %s", err))
		}

		if !cluster.IsHealthy() {
			return retry.RetryableError(errors.New("instance not yet ready"))
		}

		if err := c.read(ctx, d, meta); err != nil {
			return retry.NonRetryableError(err)
		}
		return nil
	}
}
