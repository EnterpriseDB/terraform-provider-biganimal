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
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type ClusterResource struct{}

func NewClusterResource() *ClusterResource {
	return &ClusterResource{}
}

func (c *ClusterResource) Schema() *schema.Resource {
	return &schema.Resource{
		Description: "The cluster resource is used to manage BigAnimal clusters. See [Creating a cluster](https://www.enterprisedb.com/docs/biganimal/latest/getting_started/creating_a_cluster/) for more details.",

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
				Type:        schema.TypeList,
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
			"cluster_architecture": {
				Description: "Cluster architecture. See [Supported cluster types](https://www.enterprisedb.com/docs/biganimal/latest/overview/02_high_availability/) for details.",
				Type:        schema.TypeList,
				Required:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Description: "Cluster architecture ID. For example, \"single\", \"ha\", or \"eha\".",
							Type:        schema.TypeString,
							Required:    true,
						},
						"name": {
							Description: "Name.",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"nodes": {
							Description: "Node count.",
							Type:        schema.TypeInt,
							Required:    true,
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
				Description: "Instance type. For example, \"azure:Standard_D2s_v3\" or \"aws:c5.large\".",
				Type:        schema.TypeString,
				Required:    true,
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
			"password": {
				Description: "Password for the user edb_admin. It must be 12 characters or more.",
				Type:        schema.TypeString,
				Required:    true,
				Sensitive:   true,
			},
			"pg_config": {
				Description: "Database configuration parameters. See [Modifying database configuration parameters](https://www.enterprisedb.com/docs/biganimal/latest/using_cluster/03_modifying_your_cluster/05_db_configuration_parameters/) for details.",
				Type:        schema.TypeList,
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
			"pg_type": {
				Description: "Postgres type. For example, \"epas\", \"pgextended\", or \"postgres\".",
				Type:        schema.TypeString,
				Required:    true,
			},
			"pg_version": {
				Description: "Postgres version. See [Supported Postgres types and versions](https://www.enterprisedb.com/docs/biganimal/latest/overview/05_database_version_policy/#supported-postgres-types-and-versions) for supported Postgres types and versions.",
				Type:        schema.TypeString,
				Required:    true,
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
			"cloud_provider": {
				Description: "Cloud provider. For example, \"aws\" or \"azure\".",
				Type:        schema.TypeString,
				Required:    true,
			},
			"read_only_connections": {
				Description: "Is read only connection enabled.",
				Type:        schema.TypeBool,
				Optional:    true,
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
						},
						"size": {
							Description: "Size of the volume. It can be set to different values depending on your volume type and properties.",
							Type:        schema.TypeString,
							Required:    true,
						},
						"throughput": {
							Description: "Throughput.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"volume_properties": {
							Description: "Volume properties in accordance with the selected volume type.",
							Type:        schema.TypeString,
							Required:    true,
						},
						"volume_type": {
							Description: "Volume type. For Azure: \"azurepremiumstorage\" or \"ultradisk\". For AWS: \"gp3\", \"io2\", or \"io2-block-express\".",
							Type:        schema.TypeString,
							Required:    true,
						},
					},
				},
			},
		},
	}
}

func (c *ClusterResource) Create(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := api.BuildAPI(meta).ClusterClient()

	cluster, err := models.NewClusterForCreate(d)
	if err != nil {
		return diag.FromErr(err)
	}

	clusterId, err := client.Create(ctx, *cluster)
	if err != nil {
		return fromBigAnimalErr(err)
	}

	d.SetId(clusterId)

	// retry until we get success
	err = resource.RetryContext(
		ctx,
		d.Timeout(schema.TimeoutCreate)-time.Minute,
		c.retryFunc(ctx, d, meta, clusterId))
	if err != nil {
		return diag.FromErr(err)
	}
	return diag.Diagnostics{}
}

func (c *ClusterResource) Read(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	if err := c.read(ctx, d, meta); err != nil {
		return fromBigAnimalErr(err)
	}
	return diag.Diagnostics{}
}

func (c *ClusterResource) read(ctx context.Context, d *schema.ResourceData, meta any) error {
	client := api.BuildAPI(meta).ClusterClient()

	clusterId := d.Id()
	cluster, err := client.Read(ctx, clusterId)
	if err != nil {
		return err
	}

	connection, err := client.ConnectionString(ctx, clusterId)
	if err != nil {
		return err
	}

	// set the outputs
	utils.SetOrPanic(d, "allowed_ip_ranges", cluster.AllowedIpRanges)
	utils.SetOrPanic(d, "backup_retention_period", cluster.BackupRetentionPeriod)
	utils.SetOrPanic(d, "cluster_architecture", cluster.ClusterArchitecture)
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
	utils.SetOrPanic(d, "read_only_connections", cluster.ReadOnlyConnections)
	utils.SetOrPanic(d, "csp_auth", cluster.CSPAuth)
	utils.SetOrPanic(d, "region", cluster.Region)
	utils.SetOrPanic(d, "storage", cluster.Storage)
	utils.SetOrPanic(d, "resizing_pvc", cluster.ResizingPvc)
	utils.SetOrPanic(d, "cluster_id", cluster.ClusterId)
	utils.SetOrPanic(d, "connection_uri", connection.PgUri)
	utils.SetOrPanic(d, "ro_connection_uri", connection.ReadOnlyPgUri)

	d.SetId(*cluster.ClusterId)
	return nil
}

func (c *ClusterResource) Update(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := api.BuildAPI(meta).ClusterClient()
	// short circuit early for these types of changes
	if d.HasChange("pg_type") {
		return diag.FromErr(errors.New("pg_type is immutable"))
	}
	if d.HasChange("pg_version") {
		return diag.FromErr(errors.New("pg_version is immutable"))
	}
	if d.HasChange("provider") {
		return diag.FromErr(errors.New("cloud provider is immutable"))
	}
	if d.HasChange("region") {
		return diag.FromErr(errors.New("region is immutable"))
	}
	if d.HasChange("password") {
		return diag.FromErr(errors.New("password is immutable for now"))
	}

	cluster, err := models.NewClusterForUpdate(d)
	if err != nil {
		return diag.FromErr(err)
	}

	clusterId := d.Id()

	_, err = client.Update(ctx, cluster, d.Id())
	if err != nil {
		return fromBigAnimalErr(err)
	}

	// retry until we get success
	err = resource.RetryContext(
		ctx,
		d.Timeout(schema.TimeoutUpdate)-time.Minute,
		c.retryFunc(ctx, d, meta, clusterId))
	if err != nil {
		return diag.FromErr(err)
	}
	return diag.Diagnostics{}
}

func (c *ClusterResource) Delete(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := api.BuildAPI(meta).ClusterClient()
	clusterId := d.Id()
	if err := client.Delete(ctx, clusterId); err != nil {
		return fromBigAnimalErr(err)
	}
	return diag.Diagnostics{}
}

func (c *ClusterResource) retryFunc(ctx context.Context, d *schema.ResourceData, meta any, clusterId string) resource.RetryFunc {
	client := api.BuildAPI(meta).ClusterClient()
	return func() *resource.RetryError {
		cluster, err := client.Read(ctx, clusterId)
		if err != nil {
			return resource.NonRetryableError(fmt.Errorf("Error describing instance: %s", err))
		}

		if !cluster.IsHealthy() {
			return resource.RetryableError(errors.New("Instance not yet ready"))
		}

		if err := c.read(ctx, d, meta); err != nil {
			return resource.NonRetryableError(err)
		}
		return nil
	}
}
