package provider

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/EnterpriseDB/terraform-provider-biganimal/pkg/api"
	"github.com/EnterpriseDB/terraform-provider-biganimal/pkg/models"
	commonApi "github.com/EnterpriseDB/terraform-provider-biganimal/pkg/models/common/api"
	commonTerraform "github.com/EnterpriseDB/terraform-provider-biganimal/pkg/models/common/terraform"
	"github.com/EnterpriseDB/terraform-provider-biganimal/pkg/plan_modifier"
	"github.com/EnterpriseDB/terraform-provider-biganimal/pkg/utils"
	"github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
)

var (
	_ resource.Resource              = &clusterResource{}
	_ resource.ResourceWithConfigure = &clusterResource{}
)

type ClusterResourceModel struct {
	ID                         types.String                       `tfsdk:"id"`
	CspAuth                    types.Bool                         `tfsdk:"csp_auth"`
	Region                     types.String                       `tfsdk:"region"`
	InstanceType               types.String                       `tfsdk:"instance_type"`
	ReadOnlyConnections        types.Bool                         `tfsdk:"read_only_connections"`
	ResizingPvc                types.List                         `tfsdk:"resizing_pvc"`
	MetricsUrl                 *string                            `tfsdk:"metrics_url"`
	ClusterId                  *string                            `tfsdk:"cluster_id"`
	Phase                      *string                            `tfsdk:"phase"`
	ClusterArchitecture        *ClusterArchitectureResourceModel  `tfsdk:"cluster_architecture"`
	ConnectionUri              types.String                       `tfsdk:"connection_uri"`
	ClusterName                types.String                       `tfsdk:"cluster_name"`
	RoConnectionUri            types.String                       `tfsdk:"ro_connection_uri"`
	Storage                    *StorageResourceModel              `tfsdk:"storage"`
	PgConfig                   []PgConfigResourceModel            `tfsdk:"pg_config"`
	FirstRecoverabilityPointAt *string                            `tfsdk:"first_recoverability_point_at"`
	ProjectId                  string                             `tfsdk:"project_id"`
	LogsUrl                    *string                            `tfsdk:"logs_url"`
	BackupRetentionPeriod      types.String                       `tfsdk:"backup_retention_period"`
	ClusterType                *string                            `tfsdk:"cluster_type"`
	CloudProvider              types.String                       `tfsdk:"cloud_provider"`
	PgType                     types.String                       `tfsdk:"pg_type"`
	Password                   types.String                       `tfsdk:"password"`
	FarawayReplicaIds          types.Set                          `tfsdk:"faraway_replica_ids"`
	PgVersion                  types.String                       `tfsdk:"pg_version"`
	PrivateNetworking          types.Bool                         `tfsdk:"private_networking"`
	AllowedIpRanges            []AllowedIpRangesResourceModel     `tfsdk:"allowed_ip_ranges"`
	CreatedAt                  types.String                       `tfsdk:"created_at"`
	MaintenanceWindow          *commonTerraform.MaintenanceWindow `tfsdk:"maintenance_window"`
	ServiceAccountIds          types.Set                          `tfsdk:"service_account_ids"`
	PeAllowedPrincipalIds      types.Set                          `tfsdk:"pe_allowed_principal_ids"`
	SuperuserAccess            types.Bool                         `tfsdk:"superuser_access"`
	Pgvector                   types.Bool                         `tfsdk:"pgvector"`
	PostGIS                    types.Bool                         `tfsdk:"post_gis"`
	PgBouncer                  *PgBouncerModel                    `tfsdk:"pg_bouncer"`
	Pause                      types.Bool                         `tfsdk:"pause"`
	VolumeSnapshot             types.Bool                         `tfsdk:"volume_snapshot_backup"`
	Timeouts                   timeouts.Value                     `tfsdk:"timeouts"`
}

type ClusterArchitectureResourceModel struct {
	Id    string       `tfsdk:"id"`
	Name  types.String `tfsdk:"name"`
	Nodes int          `tfsdk:"nodes"`
}

type StorageResourceModel struct {
	VolumeType       types.String `tfsdk:"volume_type"`
	VolumeProperties types.String `tfsdk:"volume_properties"`
	Size             types.String `tfsdk:"size"`
	Iops             types.String `tfsdk:"iops"`
	Throughput       types.String `tfsdk:"throughput"`
}

type PgConfigResourceModel struct {
	Name  string `tfsdk:"name"`
	Value string `tfsdk:"value"`
}

type AllowedIpRangesResourceModel struct {
	CidrBlock   string       `tfsdk:"cidr_block"`
	Description types.String `tfsdk:"description"`
}

type PgBouncerModel struct {
	IsEnabled bool      `tfsdk:"is_enabled"`
	Settings  types.Set `tfsdk:"settings"`
}

type PgBouncerSettingsModel struct {
	Name      string `tfsdk:"name"`
	Operation string `tfsdk:"operation"`
	Value     string `tfsdk:"value"`
}

func (c ClusterResourceModel) projectId() string {
	return c.ProjectId
}

func (c ClusterResourceModel) clusterId() string {
	return *c.ClusterId
}

type retryClusterResourceModel interface {
	projectId() string
	clusterId() string
}

type clusterResource struct {
	client *api.ClusterClient
}

func (c *clusterResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	c.client = req.ProviderData.(*api.API).ClusterClient()
}

func (c *clusterResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_cluster"
}

func (c *clusterResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "The cluster resource is used to manage BigAnimal clusters. See [Creating a cluster](https://www.enterprisedb.com/docs/biganimal/latest/getting_started/creating_a_cluster/) for more details.",
		// using Blocks for backward compatible
		Blocks: map[string]schema.Block{
			"timeouts": timeouts.Block(ctx,
				timeouts.Opts{Create: true, Delete: true, Update: true},
			),

			"storage": schema.SingleNestedBlock{
				MarkdownDescription: "Storage.",
				PlanModifiers:       []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"volume_properties": schema.StringAttribute{
						MarkdownDescription: "Volume properties in accordance with the selected volume type.",
						Required:            true,
					},
					"volume_type": schema.StringAttribute{
						MarkdownDescription: "Volume type. For Azure: \"azurepremiumstorage\" or \"ultradisk\". For AWS: \"gp3\", \"io2\", org s \"io2-block-express\". For Google Cloud: only \"pd-ssd\".",
						Required:            true,
					},
					"size": schema.StringAttribute{
						MarkdownDescription: "Size of the volume. It can be set to different values depending on your volume type and properties.",
						Required:            true,
					},
					"iops": schema.StringAttribute{
						MarkdownDescription: "IOPS for the selected volume. It can be set to different values depending on your volume type and properties.",
						Optional:            true,
						Computed:            true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"throughput": schema.StringAttribute{
						MarkdownDescription: "Throughput is automatically calculated by BigAnimal based on the IOPS input if it's not provided.",
						Optional:            true,
						Computed:            true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
				},
			},
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

			"pg_config": schema.SetNestedBlock{
				MarkdownDescription: "Database configuration parameters. See [Modifying database configuration parameters](https://www.enterprisedb.com/docs/biganimal/latest/using_cluster/03_modifying_your_cluster/05_db_configuration_parameters/) for details.",
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.UseStateForUnknown(),
				},
				NestedObject: schema.NestedBlockObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							MarkdownDescription: "GUC name.",
							Required:            true,
						},
						"value": schema.StringAttribute{
							MarkdownDescription: "GUC value.",
							Required:            true,
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
			"ro_connection_uri": schema.StringAttribute{
				MarkdownDescription: "Cluster read-only connection URI. Only available for high availability clusters.",
				Computed:            true,
				PlanModifiers:       []planmodifier.String{plan_modifier.CustomConnection()},
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
			"faraway_replica_ids": schema.SetAttribute{
				Computed:      true,
				PlanModifiers: []planmodifier.Set{setplanmodifier.UseStateForUnknown()},
				ElementType:   types.StringType,
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
			"read_only_connections": schema.BoolAttribute{
				MarkdownDescription: "Is read only connection enabled.",
				Optional:            true,
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

			"superuser_access": schema.BoolAttribute{
				MarkdownDescription: "Enable to grant superuser access to the edb_admin role.",
				Optional:            true,
				Computed:            true,
			},
			"volume_snapshot_backup": schema.BoolAttribute{
				MarkdownDescription: "Enable to take a snapshot of the volume.",
				Optional:            true,
				Computed:            true,
				PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
			},
			"pgvector": schema.BoolAttribute{
				MarkdownDescription: "Is pgvector extension enabled. Adds support for vector storage and vector similarity search to Postgres.",
				Optional:            true,
				Computed:            true,
				PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
			},
			"post_gis": schema.BoolAttribute{
				MarkdownDescription: "Is postGIS extension enabled. PostGIS extends the capabilities of the PostgreSQL relational database by adding support storing, indexing and querying geographic data.",
				Optional:            true,
				Computed:            true,
				PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
			},
			"pg_bouncer": schema.SingleNestedAttribute{
				MarkdownDescription: "Pg bouncer.",
				Optional:            true,
				Computed:            true,
				PlanModifiers: []planmodifier.Object{
					plan_modifier.CustomPgBouncer(),
				},
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
				Optional:      true,
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
			},
		},
	}
}

func (c *clusterResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var config ClusterResourceModel
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	clusterModel, err := c.makeClusterForCreate(ctx, config)
	if err != nil {
		if !appendDiagFromBAErr(err, &resp.Diagnostics) {
			resp.Diagnostics.AddError("Error creating cluster", err.Error())
		}
		return
	}

	clusterId, err := c.client.Create(ctx, config.ProjectId, clusterModel)
	if err != nil {
		if !appendDiagFromBAErr(err, &resp.Diagnostics) {
			resp.Diagnostics.AddError("Error creating cluster API request", err.Error())
		}
		return
	}

	config.ClusterId = &clusterId

	timeout, diagnostics := config.Timeouts.Create(ctx, time.Minute*60)
	resp.Diagnostics.Append(diagnostics...)
	if resp.Diagnostics.HasError() {
		return
	}

	if err := ensureClusterIsHealthy(ctx, c.client, config, timeout); err != nil {
		if !appendDiagFromBAErr(err, &resp.Diagnostics) {
			resp.Diagnostics.AddError("Error waiting for the cluster is ready ", err.Error())
		}
		return
	}

	if config.Pause.ValueBool() {
		_, err = c.client.ClusterPause(ctx, config.ProjectId, *config.ClusterId)
		if err != nil {
			if !appendDiagFromBAErr(err, &resp.Diagnostics) {
				resp.Diagnostics.AddError("Error pausing cluster API request", err.Error())
			}
			return
		}

		if err := ensureClusterIsPaused(ctx, c.client, config, timeout); err != nil {
			if !appendDiagFromBAErr(err, &resp.Diagnostics) {
				resp.Diagnostics.AddError("Error waiting for the cluster to pause", err.Error())
			}
			return
		}
	}

	if err := c.read(ctx, &config); err != nil {
		if !appendDiagFromBAErr(err, &resp.Diagnostics) {
			resp.Diagnostics.AddError("Error reading cluster", err.Error())
		}
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, config)...)
}

func (c *clusterResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state ClusterResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	if err := c.read(ctx, &state); err != nil {
		if !appendDiagFromBAErr(err, &resp.Diagnostics) {
			resp.Diagnostics.AddError("Error reading cluster", err.Error())
		}
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}

func (c *clusterResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan ClusterResourceModel

	timeout, diagnostics := plan.Timeouts.Update(ctx, time.Minute*60)
	resp.Diagnostics.Append(diagnostics...)

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state ClusterResourceModel
	diags = req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// cluster = pause,   tf pause = true, it will error and say you will need to set pause = false to update
	// cluster = pause,   tf pause = false, it will resume then update
	// cluster = healthy, tf pause = true, it will update then pause
	// cluster = healthy, tf pause = false, it will update
	if *state.Phase != models.PHASE_HEALTHY && *state.Phase != models.PHASE_PAUSED {
		resp.Diagnostics.AddError("Cluster not ready please wait", "Cluster not ready for update operation please wait")
		return
	}

	if *state.Phase == models.PHASE_PAUSED {
		if plan.Pause.ValueBool() {
			resp.Diagnostics.AddError("Error cannot update paused cluster", "cannot update paused cluster, please set pause = false to resume cluster")
			return
		}

		if !plan.Pause.ValueBool() {
			_, err := c.client.ClusterResume(ctx, plan.ProjectId, *plan.ClusterId)
			if err != nil {
				if !appendDiagFromBAErr(err, &resp.Diagnostics) {
					resp.Diagnostics.AddError("Error resuming cluster API request", err.Error())
				}
				return
			}

			if err := ensureClusterIsHealthy(ctx, c.client, plan, timeout); err != nil {
				if !appendDiagFromBAErr(err, &resp.Diagnostics) {
					resp.Diagnostics.AddError("Error waiting for the cluster is ready ", err.Error())
				}
				return
			}
		}
	}

	clusterModel, err := c.makeClusterForUpdate(ctx, plan)
	if err != nil {
		if !appendDiagFromBAErr(err, &resp.Diagnostics) {
			resp.Diagnostics.AddError("Error updating cluster", err.Error())
		}
		return
	}

	_, err = c.client.Update(ctx, clusterModel, plan.ProjectId, *plan.ClusterId)
	if err != nil {
		if !appendDiagFromBAErr(err, &resp.Diagnostics) {
			resp.Diagnostics.AddError("Error updating cluster API request", err.Error())
		}
		return
	}

	// sleep after update operation as API can incorrectly respond with healthy state when checking the phase
	// this is possibly a bug in the API
	time.Sleep(20 * time.Second)

	if err := ensureClusterIsHealthy(ctx, c.client, plan, timeout); err != nil {
		if !appendDiagFromBAErr(err, &resp.Diagnostics) {
			resp.Diagnostics.AddError("Error waiting for the cluster is ready ", err.Error())
		}
		return
	}

	if plan.Pause.ValueBool() {
		_, err = c.client.ClusterPause(ctx, plan.ProjectId, *plan.ClusterId)
		if err != nil {
			if !appendDiagFromBAErr(err, &resp.Diagnostics) {
				resp.Diagnostics.AddError("Error pausing cluster API request", err.Error())
			}
			return
		}

		if err := ensureClusterIsPaused(ctx, c.client, plan, timeout); err != nil {
			if !appendDiagFromBAErr(err, &resp.Diagnostics) {
				resp.Diagnostics.AddError("Error waiting for the cluster to pause", err.Error())
			}
			return
		}
	}

	if err := c.read(ctx, &plan); err != nil {
		if !appendDiagFromBAErr(err, &resp.Diagnostics) {
			resp.Diagnostics.AddError("Error reading cluster", err.Error())
		}
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

func (c *clusterResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state ClusterResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	err := c.client.Delete(ctx, state.ProjectId, *state.ClusterId)
	if err != nil {
		if !appendDiagFromBAErr(err, &resp.Diagnostics) {
			resp.Diagnostics.AddError("Error deleting cluster", err.Error())
		}
		return
	}
}

func (c *clusterResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	idParts := strings.Split(req.ID, "/")
	if len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf("Expected import identifier with format: project_id/cluster_id. Got: %q", req.ID),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("project_id"), idParts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("cluster_id"), idParts[1])...)
}

func (c *clusterResource) read(ctx context.Context, tfClusterResource *ClusterResourceModel) error {
	apiCluster, err := c.client.Read(ctx, tfClusterResource.ProjectId, *tfClusterResource.ClusterId)
	if err != nil {
		return err
	}

	connection, err := c.client.ConnectionString(ctx, tfClusterResource.ProjectId, *tfClusterResource.ClusterId)
	if err != nil {
		return err
	}

	tfClusterResource.ID = types.StringValue(fmt.Sprintf("%s/%s", tfClusterResource.ProjectId, *tfClusterResource.ClusterId))
	tfClusterResource.ClusterId = apiCluster.ClusterId
	tfClusterResource.ClusterName = types.StringPointerValue(apiCluster.ClusterName)
	tfClusterResource.ClusterType = apiCluster.ClusterType
	tfClusterResource.Phase = apiCluster.Phase
	tfClusterResource.CloudProvider = types.StringValue(apiCluster.Provider.CloudProviderId)
	tfClusterResource.ClusterArchitecture = &ClusterArchitectureResourceModel{
		Id:    apiCluster.ClusterArchitecture.ClusterArchitectureId,
		Nodes: apiCluster.ClusterArchitecture.Nodes,
		Name:  types.StringValue(apiCluster.ClusterArchitecture.ClusterArchitectureName),
	}
	tfClusterResource.Region = types.StringValue(apiCluster.Region.Id)
	tfClusterResource.InstanceType = types.StringValue(apiCluster.InstanceType.InstanceTypeId)
	tfClusterResource.Storage = &StorageResourceModel{
		VolumeType:       types.StringPointerValue(apiCluster.Storage.VolumeTypeId),
		VolumeProperties: types.StringPointerValue(apiCluster.Storage.VolumePropertiesId),
		Size:             types.StringPointerValue(apiCluster.Storage.Size),
		Iops:             types.StringPointerValue(apiCluster.Storage.Iops),
		Throughput:       types.StringPointerValue(apiCluster.Storage.Throughput),
	}
	tfClusterResource.ResizingPvc = StringSliceToList(apiCluster.ResizingPvc)
	tfClusterResource.ReadOnlyConnections = types.BoolPointerValue(apiCluster.ReadOnlyConnections)
	tfClusterResource.ConnectionUri = types.StringPointerValue(&connection.PgUri)
	tfClusterResource.RoConnectionUri = types.StringPointerValue(&connection.ReadOnlyPgUri)
	tfClusterResource.CspAuth = types.BoolPointerValue(apiCluster.CSPAuth)
	tfClusterResource.LogsUrl = apiCluster.LogsUrl
	tfClusterResource.MetricsUrl = apiCluster.MetricsUrl
	tfClusterResource.BackupRetentionPeriod = types.StringPointerValue(apiCluster.BackupRetentionPeriod)
	tfClusterResource.PgVersion = types.StringValue(apiCluster.PgVersion.PgVersionId)
	tfClusterResource.PgType = types.StringValue(apiCluster.PgType.PgTypeId)
	tfClusterResource.FarawayReplicaIds = StringSliceToSet(apiCluster.FarawayReplicaIds)
	tfClusterResource.PrivateNetworking = types.BoolPointerValue(apiCluster.PrivateNetworking)
	tfClusterResource.SuperuserAccess = types.BoolPointerValue(apiCluster.SuperuserAccess)
	tfClusterResource.VolumeSnapshot = types.BoolPointerValue(apiCluster.VolumeSnapshot)
	if apiCluster.Extensions != nil {
		for _, v := range *apiCluster.Extensions {
			if v.Enabled && v.ExtensionId == "pgvector" {
				tfClusterResource.Pgvector = types.BoolValue(true)
				break
			}
			if v.Enabled && v.ExtensionId == "postgis" {
				tfClusterResource.PostGIS = types.BoolValue(true)
				break
			}
		}
	}

	if apiCluster.FirstRecoverabilityPointAt != nil {
		firstPointAt := apiCluster.FirstRecoverabilityPointAt.String()
		tfClusterResource.FirstRecoverabilityPointAt = &firstPointAt
	}

	// pgConfig. If tf resource pg config elem matches with api response pg config elem then add the elem to tf resource pg config
	newPgConfig := []PgConfigResourceModel{}
	if configs := apiCluster.PgConfig; configs != nil {
		for _, tfCRPgConfig := range tfClusterResource.PgConfig {
			for _, apiConfig := range *configs {
				if tfCRPgConfig.Name == apiConfig.Name {
					newPgConfig = append(newPgConfig, PgConfigResourceModel{
						Name:  apiConfig.Name,
						Value: apiConfig.Value,
					})
				}
			}
		}
	}

	if len(newPgConfig) > 0 {
		tfClusterResource.PgConfig = newPgConfig
	}

	tfClusterResource.AllowedIpRanges = []AllowedIpRangesResourceModel{}
	if allowedIpRanges := apiCluster.AllowedIpRanges; allowedIpRanges != nil {
		for _, ipRange := range *allowedIpRanges {
			tfClusterResource.AllowedIpRanges = append(tfClusterResource.AllowedIpRanges, AllowedIpRangesResourceModel{
				CidrBlock:   ipRange.CidrBlock,
				Description: types.StringValue(ipRange.Description),
			})
		}
	}

	if pt := apiCluster.CreatedAt; pt != nil {
		tfClusterResource.CreatedAt = types.StringValue(pt.String())
	}

	if apiCluster.MaintenanceWindow != nil {
		tfClusterResource.MaintenanceWindow = &commonTerraform.MaintenanceWindow{
			IsEnabled: apiCluster.MaintenanceWindow.IsEnabled,
			StartDay:  types.Int64PointerValue(utils.ToPointer(int64(*apiCluster.MaintenanceWindow.StartDay))),
			StartTime: types.StringPointerValue(apiCluster.MaintenanceWindow.StartTime),
		}
	}

	if apiCluster.PeAllowedPrincipalIds != nil {
		tfClusterResource.PeAllowedPrincipalIds = StringSliceToSet(utils.ToValue(&apiCluster.PeAllowedPrincipalIds))
	}

	if apiCluster.ServiceAccountIds != nil {
		tfClusterResource.ServiceAccountIds = StringSliceToSet(utils.ToValue(&apiCluster.ServiceAccountIds))
	}

	if apiCluster.PgBouncer != nil {
		tfClusterResource.PgBouncer = &PgBouncerModel{}
		*tfClusterResource.PgBouncer = PgBouncerModel{
			IsEnabled: apiCluster.PgBouncer.IsEnabled,
		}

		settingsElemType := map[string]attr.Type{"name": types.StringType, "operation": types.StringType, "value": types.StringType}
		elem := basetypes.NewObjectValueMust(settingsElemType, map[string]attr.Value{
			"name":      basetypes.NewStringValue(""),
			"operation": basetypes.NewStringValue(""),
			"value":     basetypes.NewStringValue(""),
		})

		if !apiCluster.PgBouncer.IsEnabled {
			tfClusterResource.PgBouncer.Settings = basetypes.NewSetNull(elem.Type(ctx))
		} else if apiCluster.PgBouncer.IsEnabled &&
			apiCluster.PgBouncer.Settings != nil &&
			len(*apiCluster.PgBouncer.Settings) == 0 {
			tfClusterResource.PgBouncer.Settings = basetypes.NewSetNull(elem.Type(ctx))
		} else if apiCluster.PgBouncer.Settings != nil && len(*apiCluster.PgBouncer.Settings) > 0 {
			settings := []attr.Value{}

			for _, v := range *apiCluster.PgBouncer.Settings {
				object := basetypes.NewObjectValueMust(settingsElemType, map[string]attr.Value{
					"name":      basetypes.NewStringValue(*v.Name),
					"operation": basetypes.NewStringValue(*v.Operation),
					"value":     basetypes.NewStringValue(*v.Value),
				})
				settings = append(settings, object)
			}
			tfClusterResource.PgBouncer.Settings = basetypes.NewSetValueMust(elem.Type(ctx), settings)
		}
	}

	return nil
}

func ensureClusterIsHealthy(ctx context.Context, client *api.ClusterClient, cluster retryClusterResourceModel, timeout time.Duration) error {
	return retry.RetryContext(
		ctx,
		timeout,
		func() *retry.RetryError {
			resp, err := client.Read(ctx, cluster.projectId(), cluster.clusterId())
			if err != nil {
				return retry.NonRetryableError(err)
			}

			if !resp.IsHealthy() {
				return retry.RetryableError(errors.New("cluster not yet ready"))
			}
			return nil
		})
}

func ensureClusterIsPaused(ctx context.Context, client *api.ClusterClient, cluster retryClusterResourceModel, timeout time.Duration) error {
	return retry.RetryContext(
		ctx,
		timeout,
		func() *retry.RetryError {
			resp, err := client.Read(ctx, cluster.projectId(), cluster.clusterId())
			if err != nil {
				return retry.NonRetryableError(err)
			}

			if *resp.Phase != models.PHASE_PAUSED {
				return retry.RetryableError(errors.New("cluster not yet paused"))
			}
			return nil
		})
}

func (c *clusterResource) makeClusterForCreate(ctx context.Context, clusterResource ClusterResourceModel) (models.Cluster, error) {
	cluster := generateGenericClusterModel(clusterResource)
	// add BAH Code
	if strings.Contains(clusterResource.CloudProvider.ValueString(), "bah") {
		return c.addBAHFields(ctx, cluster, clusterResource)
	} else {
		return cluster, nil
	}
}

func (c *clusterResource) addBAHFields(ctx context.Context, cluster models.Cluster, clusterResource ClusterResourceModel) (models.Cluster, error) {
	clusterRscCSP := clusterResource.CloudProvider
	clusterRscPrincipalIds := clusterResource.PeAllowedPrincipalIds
	clusterRscSvcAcntIds := clusterResource.ServiceAccountIds

	// If there is an existing Principal Account Id for that Region, use that one.
	pids, err := c.client.GetPeAllowedPrincipalIds(ctx, clusterResource.ProjectId, clusterRscCSP.ValueString(), clusterResource.Region.ValueString())
	if err != nil {
		return models.Cluster{}, err
	}
	cluster.PeAllowedPrincipalIds = utils.ToPointer(pids.Data)

	// If there is no existing value, user should provide one
	if cluster.PeAllowedPrincipalIds != nil && len(*cluster.PeAllowedPrincipalIds) == 0 {
		// Here, we prefer to create a non-nil zero length slice, because we need empty JSON array
		// while encoding JSON objects
		// For more info, please visit https://github.com/golang/go/wiki/CodeReviewComments#declaring-empty-slices
		plist := []string{}
		for _, peId := range clusterRscPrincipalIds.Elements() {
			plist = append(plist, strings.Replace(peId.String(), "\"", "", -1))
		}

		cluster.PeAllowedPrincipalIds = utils.ToPointer(plist)
	}

	if clusterRscCSP.ValueString() == "bah:gcp" {
		// If there is an existing Service Account Id for that Region, use that one.
		sids, _ := c.client.GetServiceAccountIds(ctx, clusterResource.ProjectId, clusterResource.CloudProvider.ValueString(), clusterResource.Region.ValueString())
		cluster.ServiceAccountIds = utils.ToPointer(sids.Data)

		// If there is no existing value, user should provide one
		if cluster.ServiceAccountIds != nil && len(*cluster.ServiceAccountIds) == 0 {
			// Here, we prefer to create a non-nil zero length slice, because we need empty JSON array
			// while encoding JSON objects.
			// For more info, please visit https://github.com/golang/go/wiki/CodeReviewComments#declaring-empty-slices
			slist := []string{}
			for _, saId := range clusterRscSvcAcntIds.Elements() {
				slist = append(slist, strings.Replace(saId.String(), "\"", "", -1))
			}

			cluster.ServiceAccountIds = utils.ToPointer(slist)
		}
	}
	return cluster, nil
}

func generateGenericClusterModel(clusterResource ClusterResourceModel) models.Cluster {
	cluster := models.Cluster{
		ClusterName: clusterResource.ClusterName.ValueStringPointer(),
		Password:    clusterResource.Password.ValueStringPointer(),
		ClusterArchitecture: &models.Architecture{
			ClusterArchitectureId: clusterResource.ClusterArchitecture.Id,
			Nodes:                 clusterResource.ClusterArchitecture.Nodes,
		},
		Provider: &models.Provider{CloudProviderId: clusterResource.CloudProvider.ValueString()},
		Region:   &models.Region{Id: clusterResource.Region.ValueString()},
		Storage: &models.Storage{
			VolumePropertiesId: clusterResource.Storage.VolumeProperties.ValueStringPointer(),
			VolumeTypeId:       clusterResource.Storage.VolumeType.ValueStringPointer(),
			Iops:               clusterResource.Storage.Iops.ValueStringPointer(),
			Size:               clusterResource.Storage.Size.ValueStringPointer(),
			Throughput:         clusterResource.Storage.Throughput.ValueStringPointer(),
		},
		InstanceType:          &models.InstanceType{InstanceTypeId: clusterResource.InstanceType.ValueString()},
		PgType:                &models.PgType{PgTypeId: clusterResource.PgType.ValueString()},
		PgVersion:             &models.PgVersion{PgVersionId: clusterResource.PgVersion.ValueString()},
		CSPAuth:               clusterResource.CspAuth.ValueBoolPointer(),
		PrivateNetworking:     clusterResource.PrivateNetworking.ValueBoolPointer(),
		ReadOnlyConnections:   clusterResource.ReadOnlyConnections.ValueBoolPointer(),
		BackupRetentionPeriod: clusterResource.BackupRetentionPeriod.ValueStringPointer(),
		SuperuserAccess:       clusterResource.SuperuserAccess.ValueBoolPointer(),
		VolumeSnapshot:        clusterResource.VolumeSnapshot.ValueBoolPointer(),
	}

	cluster.Extensions = &[]models.ClusterExtension{}
	if clusterResource.Pgvector.ValueBool() {
		*cluster.Extensions = append(*cluster.Extensions, models.ClusterExtension{Enabled: true, ExtensionId: "pgvector"})
	}

	if clusterResource.PostGIS.ValueBool() {
		*cluster.Extensions = append(*cluster.Extensions, models.ClusterExtension{Enabled: true, ExtensionId: "postgis"})
	}

	allowedIpRanges := []models.AllowedIpRange{}
	for _, ipRange := range clusterResource.AllowedIpRanges {
		allowedIpRanges = append(allowedIpRanges, models.AllowedIpRange{
			CidrBlock:   ipRange.CidrBlock,
			Description: ipRange.Description.ValueString(),
		})
	}
	cluster.AllowedIpRanges = &allowedIpRanges

	configs := []models.KeyValue{}
	for _, model := range clusterResource.PgConfig {
		configs = append(configs, models.KeyValue{
			Name:  model.Name,
			Value: model.Value,
		})
	}
	cluster.PgConfig = &configs

	if clusterResource.MaintenanceWindow != nil {
		cluster.MaintenanceWindow = &commonApi.MaintenanceWindow{
			IsEnabled: clusterResource.MaintenanceWindow.IsEnabled,
			StartTime: clusterResource.MaintenanceWindow.StartTime.ValueStringPointer(),
		}

		if !clusterResource.MaintenanceWindow.StartDay.IsNull() && !clusterResource.MaintenanceWindow.StartDay.IsUnknown() {
			cluster.MaintenanceWindow.StartDay = utils.ToPointer(float64(*clusterResource.MaintenanceWindow.StartDay.ValueInt64Pointer()))
		}
	}

	if clusterResource.PgBouncer != nil {
		cluster.PgBouncer = &models.PgBouncer{}
		cluster.PgBouncer.IsEnabled = clusterResource.PgBouncer.IsEnabled
		if !clusterResource.PgBouncer.Settings.IsNull() {
			cluster.PgBouncer.Settings = &[]models.PgBouncerSettings{}
			for _, v := range clusterResource.PgBouncer.Settings.Elements() {
				name := v.(basetypes.ObjectValue).Attributes()["name"].(basetypes.StringValue)
				operation := v.(basetypes.ObjectValue).Attributes()["operation"].(basetypes.StringValue)
				value := v.(basetypes.ObjectValue).Attributes()["value"].(basetypes.StringValue)
				*cluster.PgBouncer.Settings = append(*cluster.PgBouncer.Settings,
					models.PgBouncerSettings{
						Name:      name.ValueStringPointer(),
						Operation: operation.ValueStringPointer(),
						Value:     value.ValueStringPointer(),
					},
				)
			}
		} else if clusterResource.PgBouncer.Settings.IsNull() {
			cluster.PgBouncer.Settings = &[]models.PgBouncerSettings{}
		}
	}

	return cluster
}

func (c *clusterResource) makeClusterForUpdate(ctx context.Context, clusterResource ClusterResourceModel) (*models.Cluster, error) {
	cluster, err := c.makeClusterForCreate(ctx, clusterResource)
	if err != nil {
		return nil, err
	}
	cluster.ClusterId = nil
	cluster.PgType = nil
	cluster.PgVersion = nil
	cluster.Provider = nil
	cluster.Region = nil
	return &cluster, nil
}

func NewClusterResource() resource.Resource {
	return &clusterResource{}
}

func StringSliceToList(items []string) types.List {
	var eles []attr.Value
	for _, item := range items {
		eles = append(eles, types.StringValue(item))
	}

	return types.ListValueMust(types.StringType, eles)
}

func StringSliceToSet(items *[]string) types.Set {
	if items == nil {
		return types.SetNull(types.StringType)
	}

	var eles []attr.Value
	for _, item := range *items {
		eles = append(eles, types.StringValue(item))
	}

	return types.SetValueMust(types.StringType, eles)
}
