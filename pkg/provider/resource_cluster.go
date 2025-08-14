package provider

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/EnterpriseDB/terraform-provider-biganimal/pkg/api"
	"github.com/EnterpriseDB/terraform-provider-biganimal/pkg/constants"
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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/float64planmodifier"
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
	ID                              types.String                       `tfsdk:"id"`
	CspAuth                         types.Bool                         `tfsdk:"csp_auth"`
	Region                          types.String                       `tfsdk:"region"`
	InstanceType                    types.String                       `tfsdk:"instance_type"`
	ReadOnlyConnections             types.Bool                         `tfsdk:"read_only_connections"`
	ResizingPvc                     types.List                         `tfsdk:"resizing_pvc"`
	MetricsUrl                      *string                            `tfsdk:"metrics_url"`
	ClusterId                       *string                            `tfsdk:"cluster_id"`
	Phase                           types.String                       `tfsdk:"phase"`
	ClusterArchitecture             *ClusterArchitectureResourceModel  `tfsdk:"cluster_architecture"`
	ConnectionUri                   types.String                       `tfsdk:"connection_uri"`
	ClusterName                     types.String                       `tfsdk:"cluster_name"`
	RoConnectionUri                 types.String                       `tfsdk:"ro_connection_uri"`
	Storage                         *StorageResourceModel              `tfsdk:"storage"`
	PgConfig                        []PgConfigResourceModel            `tfsdk:"pg_config"`
	FirstRecoverabilityPointAt      types.String                       `tfsdk:"first_recoverability_point_at"`
	ProjectId                       string                             `tfsdk:"project_id"`
	LogsUrl                         *string                            `tfsdk:"logs_url"`
	BackupRetentionPeriod           types.String                       `tfsdk:"backup_retention_period"`
	ClusterType                     *string                            `tfsdk:"cluster_type"`
	CloudProvider                   types.String                       `tfsdk:"cloud_provider"`
	PgType                          types.String                       `tfsdk:"pg_type"`
	Password                        types.String                       `tfsdk:"password"`
	FarawayReplicaIds               types.Set                          `tfsdk:"faraway_replica_ids"`
	PgVersion                       types.String                       `tfsdk:"pg_version"`
	PrivateNetworking               types.Bool                         `tfsdk:"private_networking"`
	AllowedIpRanges                 []AllowedIpRangesResourceModel     `tfsdk:"allowed_ip_ranges"`
	CreatedAt                       types.String                       `tfsdk:"created_at"`
	MaintenanceWindow               *commonTerraform.MaintenanceWindow `tfsdk:"maintenance_window"`
	ServiceAccountIds               types.Set                          `tfsdk:"service_account_ids"`
	PeAllowedPrincipalIds           types.Set                          `tfsdk:"pe_allowed_principal_ids"`
	SuperuserAccess                 types.Bool                         `tfsdk:"superuser_access"`
	Pgvector                        types.Bool                         `tfsdk:"pgvector"`
	PostGIS                         types.Bool                         `tfsdk:"post_gis"`
	PgBouncer                       *PgBouncerModel                    `tfsdk:"pg_bouncer"`
	Pause                           types.Bool                         `tfsdk:"pause"`
	TransparentDataEncryption       *TransparentDataEncryptionModel    `tfsdk:"transparent_data_encryption"`
	PgIdentity                      types.String                       `tfsdk:"pg_identity"`
	TransparentDataEncryptionAction types.String                       `tfsdk:"transparent_data_encryption_action"`
	VolumeSnapshot                  types.Bool                         `tfsdk:"volume_snapshot_backup"`
	Tags                            []commonTerraform.Tag              `tfsdk:"tags"`
	ServiceName                     types.String                       `tfsdk:"service_name"`
	BackupScheduleTime              types.String                       `tfsdk:"backup_schedule_time"`
	WalStorage                      types.Object                       `tfsdk:"wal_storage"`
	PrivateLinkServiceAlias         types.String                       `tfsdk:"private_link_service_alias"`
	PrivateLinkServiceName          types.String                       `tfsdk:"private_link_service_name"`

	Timeouts timeouts.Value `tfsdk:"timeouts"`
}

type ClusterArchitectureResourceModel struct {
	Id    string `tfsdk:"id"`
	Nodes int    `tfsdk:"nodes"`
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

type TransparentDataEncryptionModel struct {
	KeyId   types.String `tfsdk:"key_id"`
	KeyName types.String `tfsdk:"key_name"`
	Status  types.String `tfsdk:"status"`
}

func (c ClusterResourceModel) projectId() string {
	return c.ProjectId
}

func (c ClusterResourceModel) clusterId() string {
	return *c.ClusterId
}

func (c *ClusterResourceModel) setPhase(phase string) {
	c.Phase = types.StringValue(phase)
}

func (c *ClusterResourceModel) setPgIdentity(pgIdentity string) {
	c.PgIdentity = types.StringValue(pgIdentity)
}

func (c *ClusterResourceModel) setCloudProvider(cloudProvider string) {
	c.CloudProvider = types.StringValue(cloudProvider)
}

type retryClusterResourceModel interface {
	projectId() string
	clusterId() string
	setPhase(string)
	setPgIdentity(string)
	setCloudProvider(string)
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
		},
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "Resource ID of the cluster.",
				Computed:            true,
				PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"cluster_id": schema.StringAttribute{
				MarkdownDescription: "Cluster ID.",
				Computed:            true,
				PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"cluster_architecture": schema.SingleNestedAttribute{
				Description: "Cluster architecture.",
				Required:    true,
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Description:   "Cluster architecture ID. For example, \"single\" or \"ha\".For Extreme High Availability clusters, please use the [biganimal_pgd](https://registry.terraform.io/providers/EnterpriseDB/biganimal/latest/docs/resources/pgd) resource.",
						Required:      true,
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"nodes": schema.Float64Attribute{
						Description:   "Node count.",
						Required:      true,
						PlanModifiers: []planmodifier.Float64{float64planmodifier.UseStateForUnknown()},
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
				PlanModifiers: []planmodifier.Set{setplanmodifier.UseStateForUnknown()},
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
				PlanModifiers: []planmodifier.Set{setplanmodifier.UseStateForUnknown()},
			},
			"storage": schema.SingleNestedAttribute{
				Description: "Storage.",
				Required:    true,
				Attributes: map[string]schema.Attribute{
					"iops": schema.StringAttribute{
						Description:   "IOPS for the selected volume. It can be set to different values depending on your volume type and properties.",
						Optional:      true,
						Computed:      true,
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"size": schema.StringAttribute{
						Description:   "Size of the volume. It can be set to different values depending on your volume type and properties.",
						Required:      true,
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"throughput": schema.StringAttribute{
						Description:   "Throughput is automatically calculated by BigAnimal based on the IOPS input if it's not provided.",
						Optional:      true,
						Computed:      true,
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"volume_properties": schema.StringAttribute{
						Description: "Volume properties in accordance with the selected volume type.",
						Required:    true,
					},
					"volume_type": schema.StringAttribute{
						Description: "Volume type. For Azure: \"azurepremiumstorage\" or \"ultradisk\". For AWS: \"gp3\", \"io2\", or \"io2-block-express\". For Google Cloud: only \"pd-ssd\".",
						Required:    true,
					},
				},
			},
			"connection_uri": schema.StringAttribute{
				MarkdownDescription: "Cluster connection URI.",
				Computed:            true,
				PlanModifiers:       []planmodifier.String{plan_modifier.CustomPrivateNetworking()},
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
				PlanModifiers:       []planmodifier.String{plan_modifier.CustomPrivateNetworking()},
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
				Description:   "Cloud provider. For example, \"aws\", \"azure\", \"gcp\" or \"bah:aws\", \"bah:gcp\".",
				Required:      true,
				PlanModifiers: []planmodifier.String{plan_modifier.CustomClusterCloudProvider()},
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
				Computed:    true,
				ElementType: types.StringType,
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
				MarkdownDescription: "Instance type. For example, \"azure:Standard_D2s_v3\", \"aws:c6i.large\" or \"gcp:e2-highcpu-4\".",
				Required:            true,
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
				PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
			},
			"pgvector": schema.BoolAttribute{
				MarkdownDescription: "Is pgvector extension enabled. Adds support for vector storage and vector similarity search to Postgres.",
				Optional:            true,
				PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
			},
			"post_gis": schema.BoolAttribute{
				MarkdownDescription: "Is postGIS extension enabled. PostGIS extends the capabilities of the PostgreSQL relational database by adding support storing, indexing and querying geographic data.",
				Optional:            true,
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
			"tags": schema.SetNestedAttribute{
				Description:   "Assign existing tags or create tags to assign to this resource",
				Optional:      true,
				Computed:      true,
				NestedObject:  ResourceTagNestedObject,
				PlanModifiers: []planmodifier.Set{setplanmodifier.UseStateForUnknown()},
			},
			"transparent_data_encryption": schema.SingleNestedAttribute{
				MarkdownDescription: "Transparent Data Encryption (TDE) key",
				Optional:            true,
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.UseStateForUnknown(),
				},
				Attributes: map[string]schema.Attribute{
					"key_id": schema.StringAttribute{
						MarkdownDescription: "Transparent Data Encryption (TDE) key ID.",
						Required:            true,
						PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"key_name": schema.StringAttribute{
						MarkdownDescription: "Key name.",
						Computed:            true,
						PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"status": schema.StringAttribute{
						MarkdownDescription: "Status.",
						Computed:            true,
						PlanModifiers:       []planmodifier.String{plan_modifier.CustomTDEStatus()},
					},
				},
			},
			"pg_identity": schema.StringAttribute{
				MarkdownDescription: "PG Identity required to grant key permissions to activate the cluster.",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"transparent_data_encryption_action": schema.StringAttribute{
				MarkdownDescription: "Transparent data encryption action.",
				Computed:            true,
				PlanModifiers:       []planmodifier.String{plan_modifier.CustomTDEAction()},
			},
			"service_name": schema.StringAttribute{
				MarkdownDescription: "Cluster connection service name.",
				Computed:            true,
				PlanModifiers:       []planmodifier.String{plan_modifier.CustomPrivateNetworking()},
			},
			"backup_schedule_time": ResourceBackupScheduleTime,
			"wal_storage":          resourceWal,
			"private_link_service_alias": schema.StringAttribute{
				MarkdownDescription: "Private link service alias.",
				Computed:            true,
				// don't use state for unknown as this field is eventually consistent
			},
			"private_link_service_name": schema.StringAttribute{
				MarkdownDescription: "private link service name.",
				Computed:            true,
				// don't use state for unknown as this field is eventually consistent
			},
		},
	}
}

// modify plan on at runtime
func (c *clusterResource) ModifyPlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
	ValidateTags(ctx, c.client.TagClient(), req, resp)
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

	if err := ensureClusterIsEndStateAs(ctx, c.client, &config, timeout); err != nil {
		if !appendDiagFromBAErr(err, &resp.Diagnostics) {
			resp.Diagnostics.AddError("Error waiting for the cluster is ready ", err.Error())
		}

		return
	}

	if config.Phase.ValueString() == constants.PHASE_WAITING_FOR_ACCESS_TO_ENCRYPTION_KEY {
		resp.Diagnostics.AddWarning("Transparent data encryption action", TdeActionInfo(config.CloudProvider.ValueString()))
	}

	if config.Pause.ValueBool() {
		_, err = c.client.ClusterPause(ctx, config.ProjectId, *config.ClusterId)
		if err != nil {
			if !appendDiagFromBAErr(err, &resp.Diagnostics) {
				resp.Diagnostics.AddError("Error pausing cluster API request", err.Error())
			}
			return
		}

		if err := ensureClusterIsPaused(ctx, c.client, &config, timeout); err != nil {
			if !appendDiagFromBAErr(err, &resp.Diagnostics) {
				resp.Diagnostics.AddError("Error waiting for the cluster to pause", err.Error())
			}
			return
		}
	}

	if err := readCluster(ctx, c.client, &config); err != nil {
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

	if err := readCluster(ctx, c.client, &state); err != nil {
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
	if state.Phase.ValueString() != constants.PHASE_HEALTHY &&
		state.Phase.ValueString() != constants.PHASE_PAUSED &&
		state.Phase.ValueString() != constants.PHASE_WAITING_FOR_ACCESS_TO_ENCRYPTION_KEY {
		resp.Diagnostics.AddError("Cluster not ready please wait", "Cluster not ready for update operation please wait")
		return
	}

	if state.Phase.ValueString() == constants.PHASE_PAUSED {
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

			if err := ensureClusterIsEndStateAs(ctx, c.client, &plan, timeout); err != nil {
				if !appendDiagFromBAErr(err, &resp.Diagnostics) {
					resp.Diagnostics.AddError("Error waiting for the cluster is ready ", err.Error())
				}
				return
			}

			if plan.Phase.ValueString() == constants.PHASE_WAITING_FOR_ACCESS_TO_ENCRYPTION_KEY {
				resp.Diagnostics.AddWarning("Transparent data encryption action", TdeActionInfo(plan.CloudProvider.ValueString()))
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

	if err := ensureClusterIsEndStateAs(ctx, c.client, &plan, timeout); err != nil {
		if !appendDiagFromBAErr(err, &resp.Diagnostics) {
			resp.Diagnostics.AddError("Error waiting for the cluster is ready ", err.Error())
		}
		return
	}

	if plan.Phase.ValueString() == constants.PHASE_WAITING_FOR_ACCESS_TO_ENCRYPTION_KEY {
		resp.Diagnostics.AddWarning("Transparent data encryption action", TdeActionInfo(plan.CloudProvider.ValueString()))
	}

	if plan.Pause.ValueBool() {
		_, err = c.client.ClusterPause(ctx, plan.ProjectId, *plan.ClusterId)
		if err != nil {
			if !appendDiagFromBAErr(err, &resp.Diagnostics) {
				resp.Diagnostics.AddError("Error pausing cluster API request", err.Error())
			}
			return
		}

		if err := ensureClusterIsPaused(ctx, c.client, &plan, timeout); err != nil {
			if !appendDiagFromBAErr(err, &resp.Diagnostics) {
				resp.Diagnostics.AddError("Error waiting for the cluster to pause", err.Error())
			}
			return
		}
	}

	if err := readCluster(ctx, c.client, &plan); err != nil {
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

func readCluster(ctx context.Context, client *api.ClusterClient, tfClusterResource *ClusterResourceModel) error {
	responseCluster, err := client.Read(ctx, tfClusterResource.ProjectId, *tfClusterResource.ClusterId)
	if err != nil {
		return err
	}

	tfClusterResource.ID = types.StringValue(fmt.Sprintf("%s/%s", tfClusterResource.ProjectId, *tfClusterResource.ClusterId))
	tfClusterResource.ClusterId = responseCluster.ClusterId
	tfClusterResource.ClusterName = types.StringPointerValue(responseCluster.ClusterName)
	tfClusterResource.ClusterType = responseCluster.ClusterType
	tfClusterResource.Phase = types.StringPointerValue(responseCluster.Phase)
	tfClusterResource.CloudProvider = types.StringValue(responseCluster.Provider.CloudProviderId)
	tfClusterResource.ClusterArchitecture = &ClusterArchitectureResourceModel{
		Id:    responseCluster.ClusterArchitecture.ClusterArchitectureId,
		Nodes: responseCluster.ClusterArchitecture.Nodes,
	}
	tfClusterResource.Region = types.StringValue(responseCluster.Region.Id)
	tfClusterResource.InstanceType = types.StringValue(responseCluster.InstanceType.InstanceTypeId)
	tfClusterResource.Storage = &StorageResourceModel{
		VolumeType:       types.StringPointerValue(responseCluster.Storage.VolumeTypeId),
		VolumeProperties: types.StringPointerValue(responseCluster.Storage.VolumePropertiesId),
		Size:             types.StringPointerValue(responseCluster.Storage.Size),
		Iops:             types.StringPointerValue(responseCluster.Storage.Iops),
		Throughput:       types.StringPointerValue(responseCluster.Storage.Throughput),
	}
	tfClusterResource.ResizingPvc = StringSliceToList(responseCluster.ResizingPvc)
	tfClusterResource.ReadOnlyConnections = types.BoolPointerValue(responseCluster.ReadOnlyConnections)
	tfClusterResource.ConnectionUri = types.StringPointerValue(&responseCluster.Connection.PgUri)
	tfClusterResource.RoConnectionUri = types.StringPointerValue(&responseCluster.Connection.ReadOnlyPgUri)
	tfClusterResource.ServiceName = types.StringPointerValue(&responseCluster.Connection.ServiceName)
	tfClusterResource.PrivateLinkServiceAlias = types.StringPointerValue(&responseCluster.Connection.PrivateLinkServiceAlias)
	tfClusterResource.PrivateLinkServiceName = types.StringPointerValue(&responseCluster.Connection.PrivateLinkServiceName)
	tfClusterResource.CspAuth = types.BoolPointerValue(responseCluster.CSPAuth)
	tfClusterResource.LogsUrl = responseCluster.LogsUrl
	tfClusterResource.MetricsUrl = responseCluster.MetricsUrl
	tfClusterResource.BackupRetentionPeriod = types.StringPointerValue(responseCluster.BackupRetentionPeriod)
	tfClusterResource.BackupScheduleTime = types.StringPointerValue(responseCluster.BackupScheduleTime)
	tfClusterResource.PgVersion = types.StringValue(responseCluster.PgVersion.PgVersionId)
	tfClusterResource.PgType = types.StringValue(responseCluster.PgType.PgTypeId)
	tfClusterResource.FarawayReplicaIds = StringSliceToSet(responseCluster.FarawayReplicaIds)
	tfClusterResource.PrivateNetworking = types.BoolPointerValue(responseCluster.PrivateNetworking)
	tfClusterResource.SuperuserAccess = types.BoolPointerValue(responseCluster.SuperuserAccess)
	tfClusterResource.PgIdentity = types.StringPointerValue(responseCluster.PgIdentity)
	tfClusterResource.VolumeSnapshot = types.BoolPointerValue(responseCluster.VolumeSnapshot)

	if responseCluster.WalStorage != nil {
		tfClusterResource.WalStorage = BuildTfRsrcWalStorage(responseCluster.WalStorage)
	}

	if responseCluster.EncryptionKeyResp != nil && *responseCluster.Phase != constants.PHASE_HEALTHY {
		if !tfClusterResource.PgIdentity.IsNull() && tfClusterResource.PgIdentity.ValueString() != "" {
			tfClusterResource.TransparentDataEncryptionAction = types.StringValue(TdeActionInfo(responseCluster.Provider.CloudProviderId))
		}
	}

	if responseCluster.Extensions != nil {
		for _, v := range *responseCluster.Extensions {
			switch v.ExtensionId {
			case "pgvector":
				tfClusterResource.Pgvector = types.BoolValue(v.Enabled)
			case "postgis":
				tfClusterResource.PostGIS = types.BoolValue(v.Enabled)
			default:
			}
		}
	}

	if responseCluster.FirstRecoverabilityPointAt != nil {
		firstPointAt := responseCluster.FirstRecoverabilityPointAt.String()
		tfClusterResource.FirstRecoverabilityPointAt = basetypes.NewStringValue(firstPointAt)
	}

	// pgConfig. If tf resource pg config elem matches with api response pg config elem then add the elem to tf resource pg config
	newPgConfig := []PgConfigResourceModel{}
	if configs := responseCluster.PgConfig; configs != nil {
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
	if allowedIpRanges := responseCluster.AllowedIpRanges; allowedIpRanges != nil {
		for _, ipRange := range *allowedIpRanges {
			description := ipRange.Description

			// if cidr block is 0.0.0.0/0 then set description to empty string
			// setting private networking and leaving allowed ip ranges as empty will return
			// cidr block as 0.0.0.0/0 and description as "To allow all access"
			// so we need to set description to empty string to keep it consistent with the tf resource
			if ipRange.CidrBlock == "0.0.0.0/0" {
				description = ""
			}
			tfClusterResource.AllowedIpRanges = append(tfClusterResource.AllowedIpRanges, AllowedIpRangesResourceModel{
				CidrBlock:   ipRange.CidrBlock,
				Description: types.StringValue(description),
			})
		}
	}

	if pt := responseCluster.CreatedAt; pt != nil {
		tfClusterResource.CreatedAt = types.StringValue(pt.String())
	}

	if responseCluster.MaintenanceWindow != nil {
		tfClusterResource.MaintenanceWindow = &commonTerraform.MaintenanceWindow{
			IsEnabled: responseCluster.MaintenanceWindow.IsEnabled,
			StartDay:  types.Int64PointerValue(utils.ToPointer(int64(*responseCluster.MaintenanceWindow.StartDay))),
			StartTime: types.StringPointerValue(responseCluster.MaintenanceWindow.StartTime),
		}
	}

	if responseCluster.PeAllowedPrincipalIds != nil {
		tfClusterResource.PeAllowedPrincipalIds = StringSliceToSet(utils.ToValue(&responseCluster.PeAllowedPrincipalIds))
	}

	if responseCluster.ServiceAccountIds != nil {
		tfClusterResource.ServiceAccountIds = StringSliceToSet(utils.ToValue(&responseCluster.ServiceAccountIds))
	}

	if responseCluster.PgBouncer != nil {
		tfClusterResource.PgBouncer = &PgBouncerModel{}
		*tfClusterResource.PgBouncer = PgBouncerModel{
			IsEnabled: responseCluster.PgBouncer.IsEnabled,
		}

		settingsElemType := map[string]attr.Type{"name": types.StringType, "operation": types.StringType, "value": types.StringType}
		elem := basetypes.NewObjectValueMust(settingsElemType, map[string]attr.Value{
			"name":      basetypes.NewStringValue(""),
			"operation": basetypes.NewStringValue(""),
			"value":     basetypes.NewStringValue(""),
		})

		if !responseCluster.PgBouncer.IsEnabled {
			tfClusterResource.PgBouncer.Settings = basetypes.NewSetNull(elem.Type(ctx))
		} else if responseCluster.PgBouncer.IsEnabled &&
			responseCluster.PgBouncer.Settings != nil &&
			len(*responseCluster.PgBouncer.Settings) == 0 {
			tfClusterResource.PgBouncer.Settings = basetypes.NewSetNull(elem.Type(ctx))
		} else if responseCluster.PgBouncer.Settings != nil && len(*responseCluster.PgBouncer.Settings) > 0 {
			settings := []attr.Value{}

			for _, v := range *responseCluster.PgBouncer.Settings {
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

	buildTfRsrcTagsAs(&tfClusterResource.Tags, responseCluster.Tags)

	if responseCluster.EncryptionKeyResp != nil {
		tfClusterResource.TransparentDataEncryption = &TransparentDataEncryptionModel{}
		tfClusterResource.TransparentDataEncryption.KeyId = types.StringValue(responseCluster.EncryptionKeyResp.KeyId)
		tfClusterResource.TransparentDataEncryption.KeyName = types.StringValue(responseCluster.EncryptionKeyResp.KeyName)
		tfClusterResource.TransparentDataEncryption.Status = types.StringValue(responseCluster.EncryptionKeyResp.Status)
	}

	return nil
}

func ensureClusterIsEndStateAs(ctx context.Context, client *api.ClusterClient, outCluster retryClusterResourceModel, timeout time.Duration) error {
	return retry.RetryContext(
		ctx,
		timeout,
		func() *retry.RetryError {
			resp, err := client.Read(ctx, outCluster.projectId(), outCluster.clusterId())
			if err != nil {
				return retry.NonRetryableError(err)
			}

			outCluster.setPhase(*resp.Phase)
			outCluster.setCloudProvider(resp.Provider.CloudProviderId)
			// if waiting for access to encryption key and pgIdentity is not "", return non-retryable error
			if *resp.Phase == constants.PHASE_WAITING_FOR_ACCESS_TO_ENCRYPTION_KEY && resp.PgIdentity != nil && *resp.PgIdentity != "" {
				outCluster.setPgIdentity(*resp.PgIdentity)
				return nil
			} else if !resp.IsHealthy() {
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

			if *resp.Phase != constants.PHASE_PAUSED {
				return retry.RetryableError(errors.New("cluster not yet paused"))
			}
			return nil
		})
}

func (c *clusterResource) makeClusterForCreate(ctx context.Context, clusterResource ClusterResourceModel) (models.Cluster, error) {
	clusterModel, err := c.generateGenericClusterModel(ctx, clusterResource)
	if err != nil {
		return models.Cluster{}, err
	}
	return clusterModel, nil
}

// note: if private networking is true, it will require A peAllowedPrincipalId
func (c *clusterResource) buildRequestBah(ctx context.Context, clusterResourceModel ClusterResourceModel) (svAccIds, principalIds *[]string, err error) {
	if strings.Contains(clusterResourceModel.CloudProvider.ValueString(), "bah") {
		// If there is an existing Principal Account Id for that Region, use that one.
		pids, err := c.client.GetPeAllowedPrincipalIds(ctx, clusterResourceModel.ProjectId, clusterResourceModel.CloudProvider.ValueString(), clusterResourceModel.Region.ValueString())
		if err != nil {
			return nil, nil, err
		}
		principalIds = utils.ToPointer(pids.Data)

		// If there is no existing value, user should provide one
		if principalIds != nil && len(*principalIds) == 0 {
			// Here, we prefer to create a non-nil zero length slice, because we need empty JSON array
			// while encoding JSON objects
			// For more info, please visit https://github.com/golang/go/wiki/CodeReviewComments#declaring-empty-slices
			plist := []string{}
			for _, peId := range clusterResourceModel.PeAllowedPrincipalIds.Elements() {
				plist = append(plist, peId.(basetypes.StringValue).ValueString())
			}

			principalIds = utils.ToPointer(plist)
		}

		if clusterResourceModel.CloudProvider.ValueString() == "bah:gcp" {
			// If there is an existing Service Account Id for that Region, use that one.
			sids, _ := c.client.GetServiceAccountIds(ctx, clusterResourceModel.ProjectId, clusterResourceModel.CloudProvider.ValueString(), clusterResourceModel.Region.ValueString())
			svAccIds = utils.ToPointer(sids.Data)

			// If there is no existing value, user should provide one
			if svAccIds != nil && len(*svAccIds) == 0 {
				// Here, we prefer to create a non-nil zero length slice, because we need empty JSON array
				// while encoding JSON objects.
				// For more info, please visit https://github.com/golang/go/wiki/CodeReviewComments#declaring-empty-slices
				slist := []string{}
				for _, saId := range clusterResourceModel.ServiceAccountIds.Elements() {
					slist = append(slist, saId.(basetypes.StringValue).ValueString())
				}

				svAccIds = utils.ToPointer(slist)
			}
		}
	}
	return
}

func (c *clusterResource) generateGenericClusterModel(ctx context.Context, clusterResource ClusterResourceModel) (models.Cluster, error) {
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
		BackupScheduleTime:    clusterResource.BackupScheduleTime.ValueStringPointer(),
		SuperuserAccess:       clusterResource.SuperuserAccess.ValueBoolPointer(),
		VolumeSnapshot:        clusterResource.VolumeSnapshot.ValueBoolPointer(),
	}

	if !clusterResource.WalStorage.IsNull() {
		cluster.WalStorage = BuildRequestWalStorage(clusterResource.WalStorage)
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

	tags := []commonApi.Tag{}
	for _, tag := range clusterResource.Tags {
		tags = append(tags, commonApi.Tag{
			Color:   tag.Color.ValueStringPointer(),
			TagName: tag.TagName.ValueString(),
		})
	}
	cluster.Tags = tags

	svAccIds, principalIds, err := c.buildRequestBah(ctx, clusterResource)
	if err != nil {
		return models.Cluster{}, err
	}

	cluster.ServiceAccountIds = svAccIds
	cluster.PeAllowedPrincipalIds = principalIds

	cluster.Tags = buildApiReqTags(clusterResource.Tags)

	if clusterResource.TransparentDataEncryption != nil {
		if !clusterResource.TransparentDataEncryption.KeyId.IsNull() {
			cluster.EncryptionKeyIdReq = clusterResource.TransparentDataEncryption.KeyId.ValueStringPointer()
		}
	}

	return cluster, nil
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
	cluster.EncryptionKeyIdReq = nil
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
