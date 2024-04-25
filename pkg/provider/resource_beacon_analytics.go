package provider

import (
	"context"

	"github.com/EnterpriseDB/terraform-provider-biganimal/pkg/api"
	commonTerraform "github.com/EnterpriseDB/terraform-provider-biganimal/pkg/models/common/terraform"
	"github.com/EnterpriseDB/terraform-provider-biganimal/pkg/plan_modifier"
	"github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ resource.Resource              = &beaconAnalyticsResource{}
	_ resource.ResourceWithConfigure = &beaconAnalyticsResource{}
)

type BeaconAnalyticsResourceModel struct {
	ID                         types.String                       `tfsdk:"id"`
	CspAuth                    types.Bool                         `tfsdk:"csp_auth"`
	Region                     types.String                       `tfsdk:"region"`
	InstanceType               types.String                       `tfsdk:"instance_type"`
	ResizingPvc                types.List                         `tfsdk:"resizing_pvc"`
	MetricsUrl                 *string                            `tfsdk:"metrics_url"`
	ClusterId                  *string                            `tfsdk:"cluster_id"`
	Phase                      *string                            `tfsdk:"phase"`
	ClusterArchitecture        *ClusterArchitectureResourceModel  `tfsdk:"cluster_architecture"`
	ConnectionUri              types.String                       `tfsdk:"connection_uri"`
	ClusterName                types.String                       `tfsdk:"cluster_name"`
	PgConfig                   []PgConfigResourceModel            `tfsdk:"pg_config"`
	FirstRecoverabilityPointAt *string                            `tfsdk:"first_recoverability_point_at"`
	ProjectId                  string                             `tfsdk:"project_id"`
	LogsUrl                    *string                            `tfsdk:"logs_url"`
	BackupRetentionPeriod      types.String                       `tfsdk:"backup_retention_period"`
	ClusterType                *string                            `tfsdk:"cluster_type"`
	CloudProvider              types.String                       `tfsdk:"cloud_provider"`
	PgType                     types.String                       `tfsdk:"pg_type"`
	Password                   types.String                       `tfsdk:"password"`
	PgVersion                  types.String                       `tfsdk:"pg_version"`
	PrivateNetworking          types.Bool                         `tfsdk:"private_networking"`
	AllowedIpRanges            []AllowedIpRangesResourceModel     `tfsdk:"allowed_ip_ranges"`
	CreatedAt                  types.String                       `tfsdk:"created_at"`
	MaintenanceWindow          *commonTerraform.MaintenanceWindow `tfsdk:"maintenance_window"`
	ServiceAccountIds          types.Set                          `tfsdk:"service_account_ids"`
	PeAllowedPrincipalIds      types.Set                          `tfsdk:"pe_allowed_principal_ids"`
	PgBouncer                  *PgBouncerModel                    `tfsdk:"pg_bouncer"`
	Pause                      types.Bool                         `tfsdk:"pause"`

	Timeouts timeouts.Value `tfsdk:"timeouts"`
}

type beaconAnalyticsResource struct {
	client *api.BeaconAnalyticsClient
}

func (tr *beaconAnalyticsResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	tr.client = req.ProviderData.(*api.API).BeaconAnalyticsClient()
}

func (tr *beaconAnalyticsResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_beacon_analytics"
}

func (tf *beaconAnalyticsResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "The beacon analystics cluster resource is used to manage BigAnimal beacon analystics clusters.",
		// using Blocks for backward compatible
		Blocks: map[string]schema.Block{
			"timeouts": timeouts.Block(ctx,
				timeouts.Opts{Create: true, Delete: true, Update: true},
			),
			"allowed_ip_ranges": schema.SetNestedBlock{
				MarkdownDescription: "Allowed IP ranges.",
				PlanModifiers:       []planmodifier.Set{plan_modifier.CustomAllowedIps()},
				NestedObject: schema.NestedBlockObject{
					Attributes: map[string]schema.Attribute{
						"cidr_block": schema.StringAttribute{
							MarkdownDescription: "CIDR block.",
							Required:            true,
						},
						"description": schema.StringAttribute{
							MarkdownDescription: "CIDR block description.",
							Optional:            true,
						},
					},
				},
			},
			"cluster_architecture": schema.SingleNestedBlock{
				MarkdownDescription: "Cluster architecture. See [Supported cluster types](https://www.enterprisedb.com/docs/biganimal/latest/overview/02_high_availability/) for details.",
				Attributes: map[string]schema.Attribute{
					"nodes": schema.Int64Attribute{
						MarkdownDescription: "Node count.",
						Required:            true,
					},
					"id": schema.StringAttribute{
						MarkdownDescription: "Cluster architecture ID. For example, \"single\" or \"ha\".For Extreme High Availability clusters, please use the [biganimal_pgd](https://registry.terraform.io/providers/EnterpriseDB/biganimal/latest/docs/resources/pgd) resource.",
						Required:            true,
						Validators: []validator.String{
							stringvalidator.OneOf("single", "ha"),
						},
					},
					"name": schema.StringAttribute{
						MarkdownDescription: "Name.",
						Optional:            true,
						Computed:            true,
						PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
				},
			},
		},
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "Resource ID of the cluster.",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"cluster_id": schema.StringAttribute{
				MarkdownDescription: "Cluster ID.",
				Computed:            true,
				PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"connection_uri": schema.StringAttribute{
				MarkdownDescription: "Cluster connection URI.",
				Computed:            true,
				PlanModifiers:       []planmodifier.String{plan_modifier.CustomConnection()},
			},
			"cluster_name": schema.StringAttribute{
				MarkdownDescription: "Name of the cluster.",
				Required:            true,
			},
			"phase": schema.StringAttribute{
				MarkdownDescription: "Current phase of the cluster.",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					plan_modifier.CustomPhaseForUnknown(),
				},
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
				PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"backup_retention_period": schema.StringAttribute{
				MarkdownDescription: "Backup retention period. For example, \"7d\", \"2w\", or \"3m\".",
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					BackupRetentionPeriodValidator(),
				},
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"cluster_type": schema.StringAttribute{
				MarkdownDescription: "Type of the cluster. For example, \"cluster\" for biganimal_cluster resources, or \"faraway_replica\" for biganimal_faraway_replica resources.",
				Computed:            true,
				PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"cloud_provider": schema.StringAttribute{
				Description: "Cloud provider. For example, \"aws\", \"azure\", \"gcp\" or \"bah:aws\", \"bah:gcp\".",
				Required:    true,
			},
			"pg_type": schema.StringAttribute{
				MarkdownDescription: "Postgres type. For example, \"epas\", \"pgextended\", or \"postgres\".",
				Required:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("epas", "pgextended", "postgres"),
				},
			},
			"first_recoverability_point_at": schema.StringAttribute{
				MarkdownDescription: "Earliest backup recover time.",
				Computed:            true,
				PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"pg_version": schema.StringAttribute{
				MarkdownDescription: "Postgres version. See [Supported Postgres types and versions](https://www.enterprisedb.com/docs/biganimal/latest/overview/05_database_version_policy/#supported-postgres-types-and-versions) for supported Postgres types and versions.",
				Required:            true,
			},
			"private_networking": schema.BoolAttribute{
				MarkdownDescription: "Is private networking enabled.",
				Optional:            true,
			},
			"password": schema.StringAttribute{
				MarkdownDescription: "Password for the user edb_admin. It must be 12 characters or more.",
				Required:            true,
			},
			"created_at": schema.StringAttribute{
				MarkdownDescription: "Cluster creation time.",
				Computed:            true,
				PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"region": schema.StringAttribute{
				MarkdownDescription: "Region to deploy the cluster. See [Supported regions](https://www.enterprisedb.com/docs/biganimal/latest/overview/03a_region_support/) for supported regions.",
				Required:            true,
			},
			"instance_type": schema.StringAttribute{
				MarkdownDescription: "Instance type. For example, \"azure:Standard_D2s_v3\", \"aws:c5.large\" or \"gcp:e2-highcpu-4\".",
				Required:            true,
			},
			"resizing_pvc": schema.ListAttribute{
				MarkdownDescription: "Resizing PVC.",
				Computed:            true,
				PlanModifiers:       []planmodifier.List{listplanmodifier.UseStateForUnknown()},
				ElementType:         types.StringType,
			},
			"metrics_url": schema.StringAttribute{
				MarkdownDescription: "The URL to find the metrics of this cluster.",
				Computed:            true,
				PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"csp_auth": schema.BoolAttribute{
				MarkdownDescription: "Is authentication handled by the cloud service provider. Available for AWS only, See [Authentication](https://www.enterprisedb.com/docs/biganimal/latest/getting_started/creating_a_cluster/#authentication) for details.",
				Optional:            true,
				Computed:            true,
				PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
			},
			"maintenance_window": schema.SingleNestedAttribute{
				MarkdownDescription: "Custom maintenance window.",
				Optional:            true,
				Computed:            true,
				PlanModifiers: []planmodifier.Object{
					plan_modifier.MaintenanceWindowForUnknown(),
				},
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
				PlanModifiers:       []planmodifier.Set{setplanmodifier.UseStateForUnknown()},
			},

			"pe_allowed_principal_ids": schema.SetAttribute{
				MarkdownDescription: "Cloud provider subscription/account ID, need to be specified when cluster is deployed on BigAnimal's cloud account.",
				Optional:            true,
				Computed:            true,
				ElementType:         types.StringType,
				PlanModifiers:       []planmodifier.Set{setplanmodifier.UseStateForUnknown()},
			},
			"pause": schema.BoolAttribute{
				MarkdownDescription: "Pause cluster. If true it will put the cluster on pause and set the phase as paused, if false it will resume the cluster and set the phase as healthy. " +
					"Pausing a cluster allows you to save on compute costs without losing data or cluster configuration settings. " +
					"While paused, clusters aren't upgraded or patched, but changes are applied when the cluster resumes. " +
					"Pausing a high availability cluster shuts down all cluster nodes",
				Optional:      true,
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
			},
		},
	}
}

func (tr *beaconAnalyticsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
}

func (c *beaconAnalyticsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
}

func (c *beaconAnalyticsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
}

func (tr *beaconAnalyticsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
}

func (tr *beaconAnalyticsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
}

func NewBeaconAnalyticsResource() resource.Resource {
	return &beaconAnalyticsResource{}
}
