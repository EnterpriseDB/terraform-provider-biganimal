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
	_ datasource.DataSource              = &analyticsClusterDataSource{}
	_ datasource.DataSourceWithConfigure = &analyticsClusterDataSource{}
)

type analyticsClusterDataSourceModel struct {
	analyticsClusterResourceModel
}
type analyticsClusterDataSource struct {
	client *api.ClusterClient
}

func (r *analyticsClusterDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*api.API).ClusterClient()
}

func (r *analyticsClusterDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_analytics_cluster"
}

func (r *analyticsClusterDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "The analytics cluster resource is used to manage BigAnimal analytics clusters.",
		// using Blocks for backward compatible
		Blocks: map[string]schema.Block{
			"timeouts": timeouts.Block(ctx,
				timeouts.Opts{Create: true, Delete: true, Update: true},
			),
		},
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "Resource ID of the cluster.",
				Computed:            true,
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
			"cluster_id": schema.StringAttribute{
				MarkdownDescription: "Cluster ID.",
				Required:            true,
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
			"project_id": schema.StringAttribute{
				MarkdownDescription: "BigAnimal Project ID.",
				Required:            true,
				Validators:          []validator.String{ProjectIdValidator()},
			},
			"logs_url": schema.StringAttribute{
				MarkdownDescription: "The URL to find the logs of this cluster.",
				Computed:            true,
			},
			"backup_retention_period": schema.StringAttribute{
				MarkdownDescription: "Backup retention period. For example, \"7d\", \"2w\", or \"3m\".",
				Optional:            true,
				Computed:            true,
				Validators:          []validator.String{BackupRetentionPeriodValidator()},
			},
			"cloud_provider": schema.StringAttribute{
				Description: "Cloud provider. For example, \"aws\" or \"bah:aws\".",
				Computed:    true,
			},
			"pg_type": schema.StringAttribute{
				MarkdownDescription: "Postgres type. For example, \"epas\" or \"pgextended\".",
				Computed:            true,
				Validators:          []validator.String{stringvalidator.OneOf("epas", "pgextended", "postgres")},
			},
			"first_recoverability_point_at": schema.StringAttribute{
				MarkdownDescription: "Earliest backup recover time.",
				Computed:            true,
			},
			"pg_version": schema.StringAttribute{
				MarkdownDescription: "Postgres version. For example 16",
				Computed:            true,
			},
			"private_networking": schema.BoolAttribute{
				MarkdownDescription: "Is private networking enabled.",
				Computed:            true,
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
				MarkdownDescription: "Instance type. For example, \"azure:Standard_D2s_v3\", \"aws:c5.large\" or \"gcp:e2-highcpu-4\".",
				Computed:            true,
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
				MarkdownDescription: "Is authentication handled by the cloud service provider.",
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
						Validators:          []validator.Int64{int64validator.Between(0, 6)},
					},
					"start_time": schema.StringAttribute{
						MarkdownDescription: "Start time. \"hh:mm\", for example: \"23:59\".",
						Optional:            true,
						Computed:            true,
						Validators:          []validator.String{startTimeValidator()},
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
			"pause": schema.BoolAttribute{
				MarkdownDescription: "Pause cluster. If true it will put the cluster on pause and set the phase as paused, if false it will resume the cluster and set the phase as healthy. " +
					"Pausing a cluster allows you to save on compute costs without losing data or cluster configuration settings. " +
					"While paused, clusters aren't upgraded or patched, but changes are applied when the cluster resumes. " +
					"Pausing a high availability cluster shuts down all cluster nodes",
				Optional: true,
			},
			"tags": schema.SetNestedAttribute{
				Description: "Assign existing tags or create tags to assign to this resource",
				Computed:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"tag_id": schema.StringAttribute{
							Computed: true,
						},
						"tag_name": schema.StringAttribute{
							Computed: true,
						},
						"color": schema.StringAttribute{
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func (r *analyticsClusterDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data analyticsClusterDataSourceModel
	diags := req.Config.Get(ctx, &data.analyticsClusterResourceModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	if err := readAnalyticsCluster(ctx, r.client, &data.analyticsClusterResourceModel); err != nil {
		if !appendDiagFromBAErr(err, &resp.Diagnostics) {
			resp.Diagnostics.AddError("Error reading cluster", err.Error())
		}
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, data.analyticsClusterResourceModel)...)
}

func NewAnalyticsClusterDataSource() datasource.DataSource {
	return &analyticsClusterDataSource{}
}
