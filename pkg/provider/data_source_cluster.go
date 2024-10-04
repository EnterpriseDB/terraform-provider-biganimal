package provider

import (
	"context"

	"github.com/EnterpriseDB/terraform-provider-biganimal/pkg/api"
	"github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ datasource.DataSource              = &clusterDataSource{}
	_ datasource.DataSourceWithConfigure = &clusterDataSource{}
)

type PgConfigDatasourceModel struct {
	Value types.String `tfsdk:"value"`
	Name  types.String `tfsdk:"name"`
}

type StorageDatasourceModel struct {
	Throughput       types.String `tfsdk:"throughput"`
	VolumeProperties types.String `tfsdk:"volume_properties"`
	VolumeType       types.String `tfsdk:"volume_type"`
	Iops             types.String `tfsdk:"iops"`
	Size             types.String `tfsdk:"size"`
}

type ClusterArchitectureDatasourceModel struct {
	Nodes types.Int64  `tfsdk:"nodes"`
	Id    types.String `tfsdk:"id"`
	Name  types.String `tfsdk:"name"`
}

type clusterDatasourceModel struct {
	ClusterResourceModel
}

type AllowedIpRangesDatasourceModel struct {
	CidrBlock   types.String `tfsdk:"cidr_block"`
	Description types.String `tfsdk:"description"`
}

type clusterDataSource struct {
	client *api.ClusterClient
}

func (c *clusterDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_cluster"
}

// Configure adds the provider configured client to the data source.
func (c *clusterDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	c.client = req.ProviderData.(*api.API).ClusterClient()
}

func (c *clusterDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "The cluster data source describes a BigAnimal cluster. The data source requires your cluster ID.",
		// using Blocks for backward compatible
		Blocks: map[string]schema.Block{
			"timeouts": timeouts.Block(ctx,
				timeouts.Opts{Create: true, Delete: true, Update: true},
			),
		},
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "Data source ID.",
				Computed:            true,
			},
			"cluster_id": schema.StringAttribute{
				MarkdownDescription: "Cluster ID.",
				Required:            true,
			},
			"cluster_architecture": schema.SingleNestedAttribute{
				Description: "Cluster architecture.",
				Computed:    true,
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Description: "Cluster architecture ID. For example, \"single\" or \"ha\".For Extreme High Availability clusters, please use the [biganimal_pgd](https://registry.terraform.io/providers/EnterpriseDB/biganimal/latest/docs/resources/pgd) resource.",
						Required:    true,
					},
					"name": schema.StringAttribute{
						Description: "Name.",
						Computed:    true,
					},
					"nodes": schema.Float64Attribute{
						Description: "Node count.",
						Required:    true,
					},
				},
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
							Optional:    true,
						},
					},
				},
			},
			"pg_config": schema.SetNestedAttribute{
				Description: "Database configuration parameters. See [Modifying database configuration parameters](https://www.enterprisedb.com/docs/biganimal/latest/using_cluster/03_modifying_your_cluster/05_db_configuration_parameters/) for details.",
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
			"storage": schema.SingleNestedAttribute{
				Description: "Storage.",
				Computed:    true,
				Attributes: map[string]schema.Attribute{
					"iops": schema.StringAttribute{
						Description: "IOPS for the selected volume. It can be set to different values depending on your volume type and properties.",
						Optional:    true,
						Computed:    true,
					},
					"size": schema.StringAttribute{
						Description: "Size of the volume. It can be set to different values depending on your volume type and properties.",
						Required:    true,
					},
					"throughput": schema.StringAttribute{
						Description: "Throughput is automatically calculated by BigAnimal based on the IOPS input if it's not provided.",
						Optional:    true,
						Computed:    true,
					},
					"volume_properties": schema.StringAttribute{
						Description: "Volume properties in accordance with the selected volume type.",
						Required:    true,
					},
					"volume_type": schema.StringAttribute{
						Description: "Volume type. For Azure: \"azurepremiumstorage\" or \"ultradisk\". For AWS: \"gp3\", \"io2\", org s \"io2-block-express\". For Google Cloud: only \"pd-ssd\".",
						Required:    true,
					},
				},
			},
			"connection_uri": schema.StringAttribute{
				MarkdownDescription: "Cluster connection URI.",
				Computed:            true,
			},
			"cluster_name": schema.StringAttribute{
				MarkdownDescription: "Name of the cluster.",
				Computed:            true,
			},
			"phase": schema.StringAttribute{
				MarkdownDescription: "Current phase of the cluster.",
				Computed:            true,
			},

			"ro_connection_uri": schema.StringAttribute{
				MarkdownDescription: "Cluster read-only connection URI. Only available for high availability clusters.",
				Computed:            true,
			},
			"project_id": schema.StringAttribute{
				MarkdownDescription: "BigAnimal Project ID.",
				Required:            true,
				Validators: []validator.String{
					ProjectIdValidator(),
				},
			},
			"logs_url": schema.StringAttribute{
				MarkdownDescription: "The URL to find the logs of this cluster.",
				Computed:            true,
			},
			"backup_retention_period": schema.StringAttribute{
				MarkdownDescription: "Backup retention period. For example, \"7d\", \"2w\", or \"3m\".",
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					BackupRetentionPeriodValidator(),
				},
			},
			"cluster_type": schema.StringAttribute{
				MarkdownDescription: "Type of the cluster. For example, \"cluster\" for biganimal_cluster resources, or \"faraway_replica\" for biganimal_faraway_replica resources.",
				Computed:            true,
			},
			"cloud_provider": schema.StringAttribute{
				Description: "Cloud provider. For example, \"aws\", \"azure\", \"gcp\" or \"bah:aws\", \"bah:gcp\".",
				Computed:    true,
			},
			"pg_type": schema.StringAttribute{
				MarkdownDescription: "Postgres type. For example, \"epas\", \"pgextended\", or \"postgres\".",
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("epas", "pgextended", "postgres"),
				},
			},
			"first_recoverability_point_at": schema.StringAttribute{
				MarkdownDescription: "Earliest backup recover time.",
				Computed:            true,
			},
			"faraway_replica_ids": schema.SetAttribute{
				Computed:    true,
				ElementType: types.StringType,
			},
			"pg_version": schema.StringAttribute{
				MarkdownDescription: "Postgres version. See [Supported Postgres types and versions](https://www.enterprisedb.com/docs/biganimal/latest/overview/05_database_version_policy/#supported-postgres-types-and-versions) for supported Postgres types and versions.",
				Computed:            true,
			},
			"private_networking": schema.BoolAttribute{
				MarkdownDescription: "Is private networking enabled.",
				Optional:            true,
			},
			"password": schema.StringAttribute{
				MarkdownDescription: "Password for the user edb_admin. It must be 12 characters or more.",
				Computed:            true,
			},
			"created_at": schema.StringAttribute{
				MarkdownDescription: "Cluster creation time.",
				Computed:            true,
			},
			"region": schema.StringAttribute{
				MarkdownDescription: "Region to deploy the cluster. See [Supported regions](https://www.enterprisedb.com/docs/biganimal/latest/overview/03a_region_support/) for supported regions.",
				Computed:            true,
			},
			"instance_type": schema.StringAttribute{
				MarkdownDescription: "Instance type. For example, \"azure:Standard_D2s_v3\", \"aws:c6i.large\" or \"gcp:e2-highcpu-4\".",
				Computed:            true,
			},
			"read_only_connections": schema.BoolAttribute{
				MarkdownDescription: "Is read only connection enabled.",
				Optional:            true,
			},
			"resizing_pvc": schema.ListAttribute{
				MarkdownDescription: "Resizing PVC.",
				Computed:            true,
				ElementType:         types.StringType,
			},
			"metrics_url": schema.StringAttribute{
				MarkdownDescription: "The URL to find the metrics of this cluster.",
				Computed:            true,
			},
			"csp_auth": schema.BoolAttribute{
				MarkdownDescription: "Is authentication handled by the cloud service provider. Available for AWS only, See [Authentication](https://www.enterprisedb.com/docs/biganimal/latest/getting_started/creating_a_cluster/#authentication) for details.",
				Optional:            true,
				Computed:            true,
			},
			"maintenance_window": schema.SingleNestedAttribute{
				MarkdownDescription: "Custom maintenance window.",
				Optional:            true,
				Computed:            true,
				Attributes: map[string]schema.Attribute{
					"is_enabled": schema.BoolAttribute{
						MarkdownDescription: "Is maintenance window enabled.",
						Required:            true,
					},
					"start_day": schema.Int64Attribute{
						MarkdownDescription: "The day of week, 0 represents Sunday, 1 is Monday, and so on.",
						Optional:            true,
						Computed:            true,
						Validators: []validator.Int64{
							int64validator.Between(0, 6),
						},
					},
					"start_time": schema.StringAttribute{
						MarkdownDescription: "Start time. \"hh:mm\", for example: \"23:59\".",
						Optional:            true,
						Computed:            true,
						Validators: []validator.String{
							startTimeValidator(),
						},
					},
				},
			},
			"service_account_ids": schema.SetAttribute{
				MarkdownDescription: "A Google Cloud Service Account is used for logs. If you leave this blank, then you will be unable to access log details for this cluster. Required when cluster is deployed on BigAnimal's cloud account.",
				Optional:            true,
				Computed:            true,
				ElementType:         types.StringType,
			},

			"pe_allowed_principal_ids": schema.SetAttribute{
				MarkdownDescription: "Cloud provider subscription/account ID, need to be specified when cluster is deployed on BigAnimal's cloud account.",
				Optional:            true,
				Computed:            true,
				ElementType:         types.StringType,
			},

			"superuser_access": schema.BoolAttribute{
				MarkdownDescription: "Enable to grant superuser access to the edb_admin role.",
				Optional:            true,
				Computed:            true,
			},
			"pgvector": schema.BoolAttribute{
				MarkdownDescription: "Is pgvector extension enabled. Adds support for vector storage and vector similarity search to Postgres.",
				Optional:            true,
				Computed:            true,
			},
			"post_gis": schema.BoolAttribute{
				MarkdownDescription: "Is postGIS extension enabled. PostGIS extends the capabilities of the PostgreSQL relational database by adding support storing, indexing and querying geographic data.",
				Optional:            true,
				Computed:            true,
			},
			"pg_bouncer": schema.SingleNestedAttribute{
				MarkdownDescription: "Pg bouncer.",
				Optional:            true,
				Computed:            true,
				Attributes: map[string]schema.Attribute{
					"is_enabled": schema.BoolAttribute{
						MarkdownDescription: "Is pg bouncer enabled.",
						Required:            true,
					},
					"settings": schema.SetNestedAttribute{
						Description: "PgBouncer Configuration Settings.",
						Optional:    true,
						Computed:    true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"name": schema.StringAttribute{
									Description: "Name.",
									Required:    true,
								},
								"operation": schema.StringAttribute{
									Description: "Operation.",
									Required:    true,
									Validators: []validator.String{
										stringvalidator.OneOf("read-write", "read-only"),
									},
								},
								"value": schema.StringAttribute{
									Description: "Value.",
									Required:    true,
								},
							},
						},
					},
				},
			},
			"pause": schema.BoolAttribute{
				MarkdownDescription: "Pause cluster. If true it will put the cluster on pause and set the phase as paused, if false it will resume the cluster and set the phase as healthy. " +
					"Pausing a cluster allows you to save on compute costs without losing data or cluster configuration settings. " +
					"While paused, clusters aren't upgraded or patched, but changes are applied when the cluster resumes. " +
					"Pausing a high availability cluster shuts down all cluster nodes",
				Optional: true,
			},
			"transparent_data_encryption": schema.SingleNestedAttribute{
				MarkdownDescription: "Transparent Data Encryption (TDE) key",
				Optional:            true,
				Attributes: map[string]schema.Attribute{
					"key_id": schema.StringAttribute{
						MarkdownDescription: "Transparent Data Encryption (TDE) key ID.",
						Computed:            true,
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
				Optional:            true,
				Computed:            true,
			},
			"volume_snapshot_backup": schema.BoolAttribute{
				MarkdownDescription: "Volume snapshot.",
				Optional:            true,
			},
			"transparent_data_encryption_action": schema.StringAttribute{
				MarkdownDescription: "Transparent data encryption action.",
				Computed:            true,
			},
		},
	}
}

func (c *clusterDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data clusterDatasourceModel
	diags := req.Config.Get(ctx, &data.ClusterResourceModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	if err := readCluster(ctx, c.client, &data.ClusterResourceModel); err != nil {
		if !appendDiagFromBAErr(err, &resp.Diagnostics) {
			resp.Diagnostics.AddError("Error reading cluster", err.Error())
		}
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, data.ClusterResourceModel)...)
}

func NewClusterDataSource() datasource.DataSource {
	return &clusterDataSource{}
}
