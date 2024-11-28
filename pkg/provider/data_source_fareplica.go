package provider

import (
	"context"

	"github.com/EnterpriseDB/terraform-provider-biganimal/pkg/api"
	"github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ datasource.DataSource              = &FAReplicaData{}
	_ datasource.DataSourceWithConfigure = &FAReplicaData{}
)

type FAReplicaData struct {
	client *api.ClusterClient
}

type FAReplicaDataModel struct {
	FAReplicaResourceModel
}

func (c *FAReplicaData) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_faraway_replica"
}

// Configure adds the provider configured client to the data source.
func (c *FAReplicaData) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	c.client = req.ProviderData.(*api.API).ClusterClient()
}

func (c *FAReplicaData) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "The faraway replica cluster data source describes a BigAnimal faraway replica connected to the cluster. The data source requires faraway replica cluster ID.",
		Blocks: map[string]schema.Block{
			"timeouts": timeouts.Block(ctx,
				timeouts.Opts{Create: true, Delete: true, Update: true}),
		},
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"allowed_ip_ranges": schema.SetNestedAttribute{
				Description: "Allowed IP ranges.",
				Optional:    true,
				Computed:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"cidr_block": schema.StringAttribute{
							Description: "CIDR block",
							Required:    true,
						},
						"description": schema.StringAttribute{
							Description: "Description of CIDR block",
							Required:    true,
						},
					},
				},
			},
			"backup_retention_period": schema.StringAttribute{
				Description: "Backup retention period. For example, \"7d\", \"2w\", or \"3m\".",
				Optional:    true,
				Validators: []validator.String{
					BackupRetentionPeriodValidator(),
				},
			},
			"csp_auth": schema.BoolAttribute{
				Description: "Is authentication handled by the cloud service provider.",
				Optional:    true,
			},
			"source_cluster_id": schema.StringAttribute{
				Description: "Source cluster ID.",
				Computed:    true,
			},

			"cluster_name": schema.StringAttribute{
				Description: "Name of the faraway replica cluster.",
				Computed:    true,
			},
			"created_at": schema.StringAttribute{
				Description: "Cluster creation time.",
				Computed:    true,
			},
			"instance_type": schema.StringAttribute{
				Description: "Instance type. For example, \"azure:Standard_D2s_v3\", \"aws:c6i.large\" or \"gcp:e2-highcpu-4\".",
				Computed:    true,
			},
			"logs_url": schema.StringAttribute{
				Description: "The URL to find the logs of this cluster.",
				Computed:    true,
			},
			"metrics_url": schema.StringAttribute{
				Description: "The URL to find the metrics of this cluster.",
				Computed:    true,
			},
			"cluster_id": schema.StringAttribute{
				Description: "Cluster ID.",
				Required:    true,
			},
			"connection_uri": schema.StringAttribute{
				Description: "Cluster connection URI.",
				Computed:    true,
			},
			"pg_config": schema.SetNestedAttribute{
				Description: "Database configuration parameters.",
				Optional:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							Description: "GUC name.",
							Required:    true,
						},
						"value": schema.StringAttribute{
							Description: "GUC value.",
							Required:    true,
						},
					},
				},
			},
			"phase": schema.StringAttribute{
				Description: "Current phase of the cluster.",
				Computed:    true,
			},
			"private_networking": schema.BoolAttribute{
				Description: "Is private networking enabled.",
				Optional:    true,
			},
			"project_id": schema.StringAttribute{
				Description: "BigAnimal Project ID.",
				Optional:    true,
				Validators: []validator.String{
					ProjectIdValidator(),
				},
			},
			"region": schema.StringAttribute{
				Description: "Region to deploy the cluster. See [Supported regions](https://www.enterprisedb.com/docs/biganimal/latest/overview/03a_region_support/) for supported regions.",
				Computed:    true,
			},
			"resizing_pvc": schema.ListAttribute{
				Description: "Resizing PVC.",
				Computed:    true,
				ElementType: types.StringType,
			},
			"storage": schema.SingleNestedAttribute{
				Description: "Storage.",
				Computed:    true,
				Attributes: map[string]schema.Attribute{
					"iops": schema.StringAttribute{
						Description: "IOPS for the selected volume.",
						Optional:    true,
						Computed:    true,
					},
					"size": schema.StringAttribute{
						Description: "Size of the volume.",
						Optional:    true,
						Computed:    true,
					},
					"throughput": schema.StringAttribute{
						Description: "Throughput.",
						Optional:    true,
						Computed:    true,
					},
					"volume_properties": schema.StringAttribute{
						Description: "Volume properties.",
						Required:    true,
					},
					"volume_type": schema.StringAttribute{
						Description: "Volume type.",
						Required:    true,
					},
				},
			},
			"service_account_ids": schema.SetAttribute{
				Description: "A Google Cloud Service Account is used for logs. If you leave this blank, then you will be unable to access log details for this cluster. Required when cluster is deployed on BigAnimal's cloud account.",
				Optional:    true,
				Computed:    true,
				ElementType: types.StringType,
			},
			"pe_allowed_principal_ids": schema.SetAttribute{
				Description: "Cloud provider subscription/account ID, need to be specified when cluster is deployed on BigAnimal's cloud account.",
				Optional:    true,
				Computed:    true,
				ElementType: types.StringType,
			},
			"cluster_type": schema.StringAttribute{
				MarkdownDescription: "Type of the cluster. For example, \"cluster\" for biganimal_cluster resources, or \"faraway_replica\" for biganimal_faraway_replica resources.",
				Computed:            true,
			},
			"cluster_architecture": schema.SingleNestedAttribute{
				Description: "Cluster architecture.",
				Computed:    true,
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Description: "Cluster architecture ID. For example, \"single\" or \"ha\".For Extreme High Availability clusters, please use the [biganimal_pgd](https://registry.terraform.io/providers/EnterpriseDB/biganimal/latest/docs/resources/pgd) resource.",
						Computed:    true,
					},
					"name": schema.StringAttribute{
						Description: "Name.",
						Computed:    true,
					},
					"nodes": schema.Float64Attribute{
						Description: "Node count.",
						Computed:    true,
					},
				},
			},
			"pg_version": schema.StringAttribute{
				MarkdownDescription: "Postgres version. See [Supported Postgres types and versions](https://www.enterprisedb.com/docs/biganimal/latest/overview/05_database_version_policy/#supported-postgres-types-and-versions) for supported Postgres types and versions.",
				Computed:            true,
			},
			"pg_type": schema.StringAttribute{
				MarkdownDescription: "Postgres type. For example, \"epas\", \"pgextended\", or \"postgres\".",
				Computed:            true,
			},
			"cloud_provider": schema.StringAttribute{
				Description: "Cloud provider. For example, \"aws\", \"azure\", \"gcp\" or \"bah:aws\", \"bah:gcp\".",
				Computed:    true,
			},
			"volume_snapshot_backup": schema.BoolAttribute{
				MarkdownDescription: "Enable to take a snapshot of the volume.",
				Computed:            true,
			},
			"transparent_data_encryption": schema.SingleNestedAttribute{
				MarkdownDescription: "Transparent Data Encryption (TDE) key",
				Optional:            true,
				Attributes: map[string]schema.Attribute{
					"key_id": schema.StringAttribute{
						MarkdownDescription: "Transparent Data Encryption (TDE) key ID.",
						Required:            true,
					},
					"key_name": schema.StringAttribute{
						MarkdownDescription: "Key name.",
						Computed:            true,
					},
					"status": schema.StringAttribute{
						MarkdownDescription: "Status.",
						Computed:            true,
					},
				},
			},
			"pg_identity": schema.StringAttribute{
				MarkdownDescription: "PG Identity required to grant key permissions to activate the cluster.",
				Computed:            true,
			},
			"transparent_data_encryption_action": schema.StringAttribute{
				MarkdownDescription: "Transparent data encryption action.",
				Computed:            true,
			},
			"tags": schema.SetNestedAttribute{
				Description: "show tags associated with this resource",
				Computed:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"tag_id": schema.StringAttribute{
							Computed: true,
						},
						"tag_name": schema.StringAttribute{
							Required: true,
						},
						"color": schema.StringAttribute{
							Computed: true,
						},
					},
				},
			},
			"backup_schedule_time": ResourceBackupScheduleTime,
			"wal_storage":          resourceWal,
		},
	}
}

func (c *FAReplicaData) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data FAReplicaDataModel
	diags := req.Config.Get(ctx, &data.FAReplicaResourceModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	if err := readFAReplica(ctx, c.client, &data.FAReplicaResourceModel); err != nil {
		if !appendDiagFromBAErr(err, &resp.Diagnostics) {
			resp.Diagnostics.AddError("Error reading faraway-replica", err.Error())
		}
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, data.FAReplicaResourceModel)...)
}

// func (c *FAReplicaData) Read(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
// 	diags := diag.Diagnostics{}
// 	client := api.BuildAPI(meta).ClusterClient()

// 	clusterId, ok := d.Get("cluster_id").(string)
// 	if !ok {
// 		return diag.FromErr(errors.New("unable to find cluster ID"))
// 	}
// 	projectId := d.Get("project_id").(string)

// 	cluster, err := client.Read(ctx, projectId, clusterId)
// 	if err != nil {
// 		return fromBigAnimalErr(err)
// 	}
// 	tflog.Debug(ctx, pretty.Sprint(cluster))

// 	if *cluster.ClusterType != "faraway_replica" {
// 		return diag.FromErr(errors.New("specified cluster is not a 'faraway replica', please use 'biganimal_cluster' data source to fetch details about this cluster"))
// 	}

// 	connection, err := client.ConnectionString(ctx, projectId, *cluster.ClusterId)
// 	if err != nil {
// 		return diag.FromErr(err)
// 	}

// 	// set the outputs
// 	utils.SetOrPanic(d, "source_cluster_id", cluster.ReplicaSourceClusterId)
// 	utils.SetOrPanic(d, "cluster_type", cluster.ClusterType)
// 	utils.SetOrPanic(d, "allowed_ip_ranges", cluster.AllowedIpRanges)
// 	utils.SetOrPanic(d, "backup_retention_period", cluster.BackupRetentionPeriod)
// 	utils.SetOrPanic(d, "cluster_architecture", *cluster.ClusterArchitecture)
// 	utils.SetOrPanic(d, "created_at", cluster.CreatedAt)
// 	utils.SetOrPanic(d, "csp_auth", cluster.CSPAuth)
// 	utils.SetOrPanic(d, "deleted_at", cluster.DeletedAt)
// 	utils.SetOrPanic(d, "expired_at", cluster.ExpiredAt)
// 	utils.SetOrPanic(d, "cluster_name", cluster.ClusterName)
// 	utils.SetOrPanic(d, "first_recoverability_point_at", cluster.FirstRecoverabilityPointAt)
// 	utils.SetOrPanic(d, "instance_type", cluster.InstanceType)
// 	utils.SetOrPanic(d, "logs_url", cluster.LogsUrl)
// 	utils.SetOrPanic(d, "metrics_url", cluster.MetricsUrl)
// 	utils.SetOrPanic(d, "pg_config", cluster.PgConfig)
// 	utils.SetOrPanic(d, "pg_type", cluster.PgType)
// 	utils.SetOrPanic(d, "pg_version", cluster.PgVersion)
// 	utils.SetOrPanic(d, "phase", cluster.Phase)
// 	utils.SetOrPanic(d, "private_networking", cluster.PrivateNetworking)
// 	utils.SetOrPanic(d, "cloud_provider", cluster.Provider)
// 	utils.SetOrPanic(d, "region", cluster.Region)
// 	utils.SetOrPanic(d, "storage", *cluster.Storage)
// 	utils.SetOrPanic(d, "resizing_pvc", cluster.ResizingPvc)
// 	utils.SetOrPanic(d, "cluster_id", cluster.ClusterId)
// 	utils.SetOrPanic(d, "connection_uri", connection.PgUri)

// 	if *cluster.Extensions != nil {
// 		for _, v := range *cluster.Extensions {
// 			if v.Enabled && v.ExtensionId == "pgvector" {
// 				utils.SetOrPanic(d, "pgvector", true) // Computed
// 			}
// 			if v.Enabled && v.ExtensionId == "postgis" {
// 				utils.SetOrPanic(d, "post_gis", true) // Computed
// 			}
// 		}
// 	}

// 	d.SetId(*cluster.ClusterId)

// 	return diags
// }

func NewFAReplicaDataSource() datasource.DataSource {
	return &FAReplicaData{}
}
