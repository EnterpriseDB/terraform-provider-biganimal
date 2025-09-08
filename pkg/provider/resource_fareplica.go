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
)

type FAReplicaResource struct {
	client *api.ClusterClient
}

type FAReplicaResourceModel struct {
	ID                              types.String                      `tfsdk:"id"`
	CspAuth                         types.Bool                        `tfsdk:"csp_auth"`
	Region                          types.String                      `tfsdk:"region"`
	InstanceType                    types.String                      `tfsdk:"instance_type"`
	ResizingPvc                     types.List                        `tfsdk:"resizing_pvc"`
	MetricsUrl                      *string                           `tfsdk:"metrics_url"`
	ClusterId                       *string                           `tfsdk:"cluster_id"`
	ReplicaSourceClusterId          *string                           `tfsdk:"source_cluster_id"`
	Phase                           types.String                      `tfsdk:"phase"`
	ConnectionUri                   types.String                      `tfsdk:"connection_uri"`
	ClusterName                     types.String                      `tfsdk:"cluster_name"`
	Storage                         *StorageResourceModel             `tfsdk:"storage"`
	PgConfig                        []PgConfigResourceModel           `tfsdk:"pg_config"`
	ProjectId                       string                            `tfsdk:"project_id"`
	LogsUrl                         *string                           `tfsdk:"logs_url"`
	BackupRetentionPeriod           types.String                      `tfsdk:"backup_retention_period"`
	PrivateNetworking               types.Bool                        `tfsdk:"private_networking"`
	AllowedIpRanges                 types.Set                         `tfsdk:"allowed_ip_ranges"`
	CreatedAt                       types.String                      `tfsdk:"created_at"`
	ServiceAccountIds               types.Set                         `tfsdk:"service_account_ids"`
	PeAllowedPrincipalIds           types.Set                         `tfsdk:"pe_allowed_principal_ids"`
	ClusterArchitecture             *ClusterArchitectureResourceModel `tfsdk:"cluster_architecture"`
	ClusterType                     types.String                      `tfsdk:"cluster_type"`
	PgType                          types.String                      `tfsdk:"pg_type"`
	PgVersion                       types.String                      `tfsdk:"pg_version"`
	CloudProvider                   types.String                      `tfsdk:"cloud_provider"`
	TransparentDataEncryption       *TransparentDataEncryptionModel   `tfsdk:"transparent_data_encryption"`
	PgIdentity                      types.String                      `tfsdk:"pg_identity"`
	TransparentDataEncryptionAction types.String                      `tfsdk:"transparent_data_encryption_action"`
	VolumeSnapshot                  types.Bool                        `tfsdk:"volume_snapshot_backup"`
	Tags                            []commonTerraform.Tag             `tfsdk:"tags"`
	BackupScheduleTime              types.String                      `tfsdk:"backup_schedule_time"`
	WalStorage                      *StorageResourceModel             `tfsdk:"wal_storage"`
	PrivateLinkServiceAlias         types.String                      `tfsdk:"private_link_service_alias"`
	PrivateLinkServiceName          types.String                      `tfsdk:"private_link_service_name"`

	Timeouts timeouts.Value `tfsdk:"timeouts"`
}

func (c FAReplicaResourceModel) projectId() string {
	return c.ProjectId
}

func (c FAReplicaResourceModel) clusterId() string {
	return *c.ClusterId
}

func (c *FAReplicaResourceModel) setPhase(phase string) {
	c.Phase = types.StringValue(phase)
}

func (c *FAReplicaResourceModel) setPgIdentity(pgIdentity string) {
	c.PgIdentity = types.StringValue(pgIdentity)
}

func (c *FAReplicaResourceModel) setCloudProvider(cloudProvider string) {
	c.CloudProvider = types.StringValue(cloudProvider)
}

func NewFAReplicaResource() resource.Resource {
	return &FAReplicaResource{}
}

func fAReplicaSchema(ctx context.Context) *schema.Schema {
	return &schema.Schema{
		MarkdownDescription: "The faraway replica resource is used to manage cluster faraway-replicas on different active regions in the cloud. See [Managing replicas](https://www.enterprisedb.com/docs/biganimal/latest/using_cluster/managing_replicas/) for more details.",
		Blocks: map[string]schema.Block{
			"timeouts": timeouts.Block(ctx,
				timeouts.Opts{Create: true, Delete: true, Update: true}),
		},
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"allowed_ip_ranges": resourceAllowedIpRanges,
			"backup_retention_period": schema.StringAttribute{
				Description: "Backup retention period. For example, \"7d\", \"2w\", or \"3m\".",
				Optional:    true,
				Validators: []validator.String{
					BackupRetentionPeriodValidator(),
				},
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"csp_auth": schema.BoolAttribute{
				Description: "Is authentication handled by the cloud service provider.",
				Optional:    true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"source_cluster_id": schema.StringAttribute{
				Description: "Source cluster ID.",
				Required:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},

			"cluster_name": schema.StringAttribute{
				Description: "Name of the faraway replica cluster.",
				Required:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"created_at": schema.StringAttribute{
				Description: "Cluster creation time.",
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"instance_type": schema.StringAttribute{
				Description: "Instance type. For example, \"azure:Standard_D2s_v3\", \"aws:c6i.large\" or \"gcp:e2-highcpu-4\".",
				Required:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"logs_url": schema.StringAttribute{
				Description: "The URL to find the logs of this cluster.",
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"metrics_url": schema.StringAttribute{
				Description: "The URL to find the metrics of this cluster.",
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"cluster_id": schema.StringAttribute{
				Description: "Cluster ID.",
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"connection_uri": schema.StringAttribute{
				Description: "Cluster connection URI.",
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
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
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.UseStateForUnknown(),
				},
			},
			"phase": schema.StringAttribute{
				Description: "Current phase of the cluster.",
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"private_networking": schema.BoolAttribute{
				Description: "Is private networking enabled.",
				Optional:    true,
				// don't use state for unknown as this field is eventually consistent

			},
			"project_id": schema.StringAttribute{
				Description: "BigAnimal Project ID.",
				Optional:    true,
				Validators: []validator.String{
					ProjectIdValidator(),
				},
				// don't use state for unknown as this field is eventually consistent

			},
			"region": schema.StringAttribute{
				Description: "Region to deploy the cluster. See [Supported regions](https://www.enterprisedb.com/docs/biganimal/latest/overview/03a_region_support/) for supported regions.",
				Required:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"resizing_pvc": schema.ListAttribute{
				Description: "Resizing PVC.",
				Computed:    true,
				ElementType: types.StringType,
			},
			"storage": schema.SingleNestedAttribute{
				Description: "Storage.",
				Required:    true,
				Attributes: map[string]schema.Attribute{
					"iops": schema.StringAttribute{
						Description: "IOPS for the selected volume.",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"size": schema.StringAttribute{
						Description: "Size of the volume.",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"throughput": schema.StringAttribute{
						Description: "Throughput.",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
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
				Description:   "A Google Cloud Service Account is used for logs. If you leave this blank, then you will be unable to access log details for this cluster. Required when cluster is deployed on BigAnimal's cloud account.",
				Optional:      true,
				Computed:      true,
				ElementType:   types.StringType,
				PlanModifiers: []planmodifier.Set{setplanmodifier.UseStateForUnknown()},
			},
			"pe_allowed_principal_ids": schema.SetAttribute{
				Description:   "Cloud provider subscription/account ID, need to be specified when cluster is deployed on BigAnimal's cloud account.",
				Optional:      true,
				Computed:      true,
				ElementType:   types.StringType,
				PlanModifiers: []planmodifier.Set{setplanmodifier.UseStateForUnknown()},
			},
			"cluster_type": schema.StringAttribute{
				MarkdownDescription: "Type of the cluster. For example, \"cluster\" for biganimal_cluster resources, or \"faraway_replica\" for biganimal_faraway_replica resources.",
				Computed:            true,
				PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"cluster_architecture": schema.SingleNestedAttribute{
				Description: "Cluster architecture.",
				Computed:    true,
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Description:   "Cluster architecture ID.",
						Required:      true,
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"nodes": schema.Float64Attribute{
						Description:   "Node count.",
						Required:      true,
						PlanModifiers: []planmodifier.Float64{float64planmodifier.UseStateForUnknown()},
					},
				},
				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
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
			"volume_snapshot_backup": schema.BoolAttribute{
				MarkdownDescription: "Enable to take a snapshot of the volume.",
				Optional:            true,
				PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
			},
			"tags": schema.SetNestedAttribute{
				Description:   "Assign existing tags or create tags to assign to this resource",
				Optional:      true,
				Computed:      true,
				NestedObject:  ResourceTagNestedObject,
				PlanModifiers: []planmodifier.Set{setplanmodifier.UseStateForUnknown()},
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

func (r *FAReplicaResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = *fAReplicaSchema(ctx)
}

// modify plan on at runtime
func (r *FAReplicaResource) ModifyPlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
	ValidateTags(ctx, r.client.TagClient(), req, resp)
}

func (r *FAReplicaResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*api.API).ClusterClient()
}

func (r *FAReplicaResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_faraway_replica"
}

func (r *FAReplicaResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var config FAReplicaResourceModel
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	clusterModel, err := r.generateGenericFAReplicaModel(ctx, config)
	if err != nil {
		if !appendDiagFromBAErr(err, &resp.Diagnostics) {
			resp.Diagnostics.AddError("Error generating faraway-replica create request", err.Error())
		}
		return
	}

	clusterId, err := r.client.Create(ctx, config.ProjectId, clusterModel)
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

	if err := ensureClusterIsEndStateAs(ctx, r.client, &config, timeout); err != nil {
		if !appendDiagFromBAErr(err, &resp.Diagnostics) {
			resp.Diagnostics.AddError("Error waiting for the cluster is ready ", err.Error())
		}

		return
	}

	if config.Phase.ValueString() == constants.PHASE_WAITING_FOR_ACCESS_TO_ENCRYPTION_KEY {
		resp.Diagnostics.AddWarning("Transparent data encryption action", TdeActionInfo(config.CloudProvider.ValueString()))
	}

	if err := readFAReplica(ctx, r.client, &config); err != nil {
		if !appendDiagFromBAErr(err, &resp.Diagnostics) {
			resp.Diagnostics.AddError("Error reading faraway replica", err.Error())
		}
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, config)...)
}

func (r *FAReplicaResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state FAReplicaResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	if err := readFAReplica(ctx, r.client, &state); err != nil {
		if !appendDiagFromBAErr(err, &resp.Diagnostics) {
			resp.Diagnostics.AddError("Error reading faraway-replica", err.Error())
		}
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}

func (r *FAReplicaResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan FAReplicaResourceModel

	timeout, diagnostics := plan.Timeouts.Update(ctx, time.Minute*60)
	resp.Diagnostics.Append(diagnostics...)

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state FAReplicaResourceModel
	diags = req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	fAReplicaModel, err := r.makeFaReplicaForUpdate(ctx, plan)
	if err != nil {
		if !appendDiagFromBAErr(err, &resp.Diagnostics) {
			resp.Diagnostics.AddError("Error updating faraway replica", err.Error())
		}
		return
	}

	_, err = r.client.Update(ctx, fAReplicaModel, plan.ProjectId, *plan.ClusterId)
	if err != nil {
		if !appendDiagFromBAErr(err, &resp.Diagnostics) {
			resp.Diagnostics.AddError("Error updating faraway replica API request", err.Error())
		}
		return
	}

	// sleep after update operation as API can incorrectly respond with healthy state when checking the phase
	// this is possibly a bug in the API
	time.Sleep(20 * time.Second)

	if err := ensureClusterIsEndStateAs(ctx, r.client, &plan, timeout); err != nil {
		if !appendDiagFromBAErr(err, &resp.Diagnostics) {
			resp.Diagnostics.AddError("Error waiting for the cluster is ready ", err.Error())
		}

		return
	}

	if plan.Phase.ValueString() == constants.PHASE_WAITING_FOR_ACCESS_TO_ENCRYPTION_KEY {
		resp.Diagnostics.AddWarning("Transparent data encryption action", TdeActionInfo(plan.CloudProvider.ValueString()))
	}

	if err := readFAReplica(ctx, r.client, &plan); err != nil {
		if !appendDiagFromBAErr(err, &resp.Diagnostics) {
			resp.Diagnostics.AddError("Error reading faraway replica", err.Error())
		}
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

func (r *FAReplicaResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state FAReplicaResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	err := r.client.Delete(ctx, state.ProjectId, *state.ClusterId)
	if err != nil {
		if !appendDiagFromBAErr(err, &resp.Diagnostics) {
			resp.Diagnostics.AddError("Error deleting faraway replica", err.Error())
		}
		return
	}
}

func (r FAReplicaResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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

func readFAReplica(ctx context.Context, client *api.ClusterClient, fAReplicaResourceModel *FAReplicaResourceModel) error {
	responseCluster, err := client.Read(ctx, fAReplicaResourceModel.ProjectId, *fAReplicaResourceModel.ClusterId)
	if err != nil {
		return err
	}

	fAReplicaResourceModel.ID = types.StringValue(fmt.Sprintf("%s/%s", fAReplicaResourceModel.ProjectId, *fAReplicaResourceModel.ClusterId))
	fAReplicaResourceModel.ReplicaSourceClusterId = responseCluster.ReplicaSourceClusterId
	fAReplicaResourceModel.ClusterId = responseCluster.ClusterId
	fAReplicaResourceModel.ClusterName = types.StringPointerValue(responseCluster.ClusterName)
	fAReplicaResourceModel.Phase = types.StringPointerValue(responseCluster.Phase)
	fAReplicaResourceModel.Region = types.StringValue(responseCluster.Region.Id)
	fAReplicaResourceModel.InstanceType = types.StringValue(responseCluster.InstanceType.InstanceTypeId)
	fAReplicaResourceModel.Storage = &StorageResourceModel{
		VolumeType:       types.StringPointerValue(responseCluster.Storage.VolumeTypeId),
		VolumeProperties: types.StringPointerValue(responseCluster.Storage.VolumePropertiesId),
		Size:             types.StringPointerValue(responseCluster.Storage.Size),
		Iops:             types.StringPointerValue(responseCluster.Storage.Iops),
		Throughput:       types.StringPointerValue(responseCluster.Storage.Throughput),
	}
	fAReplicaResourceModel.ResizingPvc = StringSliceToList(responseCluster.ResizingPvc)
	fAReplicaResourceModel.ConnectionUri = types.StringPointerValue(&responseCluster.Connection.PgUri)
	fAReplicaResourceModel.PrivateLinkServiceAlias = types.StringPointerValue(&responseCluster.Connection.PrivateLinkServiceAlias)
	fAReplicaResourceModel.PrivateLinkServiceName = types.StringPointerValue(&responseCluster.Connection.PrivateLinkServiceName)
	fAReplicaResourceModel.CspAuth = types.BoolPointerValue(responseCluster.CSPAuth)
	fAReplicaResourceModel.LogsUrl = responseCluster.LogsUrl
	fAReplicaResourceModel.MetricsUrl = responseCluster.MetricsUrl
	fAReplicaResourceModel.BackupRetentionPeriod = types.StringPointerValue(responseCluster.BackupRetentionPeriod)
	fAReplicaResourceModel.BackupScheduleTime = types.StringPointerValue(responseCluster.BackupScheduleTime)
	fAReplicaResourceModel.PrivateNetworking = types.BoolPointerValue(responseCluster.PrivateNetworking)
	fAReplicaResourceModel.ClusterArchitecture = &ClusterArchitectureResourceModel{
		Id:    responseCluster.ClusterArchitecture.ClusterArchitectureId,
		Nodes: responseCluster.ClusterArchitecture.Nodes,
	}
	fAReplicaResourceModel.ClusterType = types.StringPointerValue(responseCluster.ClusterType)
	fAReplicaResourceModel.CloudProvider = types.StringValue(responseCluster.Provider.CloudProviderId)
	fAReplicaResourceModel.PgVersion = types.StringValue(responseCluster.PgVersion.PgVersionId)
	fAReplicaResourceModel.PgType = types.StringValue(responseCluster.PgType.PgTypeId)
	fAReplicaResourceModel.VolumeSnapshot = types.BoolPointerValue(responseCluster.VolumeSnapshot)
	if responseCluster.WalStorage != nil {
		fAReplicaResourceModel.WalStorage = &StorageResourceModel{
			VolumeType:       types.StringPointerValue(responseCluster.WalStorage.VolumeTypeId),
			VolumeProperties: types.StringPointerValue(responseCluster.WalStorage.VolumePropertiesId),
			Size:             types.StringPointerValue(responseCluster.WalStorage.Size),
			Iops:             types.StringPointerValue(responseCluster.WalStorage.Iops),
			Throughput:       types.StringPointerValue(responseCluster.WalStorage.Throughput),
		}
	}

	// pgConfig. If tf resource pg config elem matches with api response pg config elem then add the elem to tf resource pg config
	newPgConfig := []PgConfigResourceModel{}
	if configs := responseCluster.PgConfig; configs != nil {
		for _, tfCRPgConfig := range fAReplicaResourceModel.PgConfig {
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
		fAReplicaResourceModel.PgConfig = newPgConfig
	}

	allowedIpRanges, diag := buildTFRsrcAllowedIpRanges(responseCluster.AllowedIpRanges)
	if diag.HasError() {
		return errors.New("error building allowed_ip_ranges")
	}

	fAReplicaResourceModel.AllowedIpRanges = allowedIpRanges

	if pt := responseCluster.CreatedAt; pt != nil {
		fAReplicaResourceModel.CreatedAt = types.StringValue(pt.String())
	}

	if responseCluster.PeAllowedPrincipalIds != nil {
		fAReplicaResourceModel.PeAllowedPrincipalIds = StringSliceToSet(utils.ToValue(&responseCluster.PeAllowedPrincipalIds))
	}

	if responseCluster.ServiceAccountIds != nil {
		fAReplicaResourceModel.ServiceAccountIds = StringSliceToSet(utils.ToValue(&responseCluster.ServiceAccountIds))
	}

	fAReplicaResourceModel.PgIdentity = types.StringPointerValue(responseCluster.PgIdentity)
	if responseCluster.EncryptionKeyResp != nil && *responseCluster.Phase != constants.PHASE_HEALTHY {
		if !fAReplicaResourceModel.PgIdentity.IsNull() && fAReplicaResourceModel.PgIdentity.ValueString() != "" {
			fAReplicaResourceModel.TransparentDataEncryptionAction = types.StringValue(TdeActionInfo(responseCluster.Provider.CloudProviderId))
		}
	}

	if responseCluster.EncryptionKeyResp != nil {
		fAReplicaResourceModel.TransparentDataEncryption = &TransparentDataEncryptionModel{}
		fAReplicaResourceModel.TransparentDataEncryption.KeyId = types.StringValue(responseCluster.EncryptionKeyResp.KeyId)
		fAReplicaResourceModel.TransparentDataEncryption.KeyName = types.StringValue(responseCluster.EncryptionKeyResp.KeyName)
		fAReplicaResourceModel.TransparentDataEncryption.Status = types.StringValue(responseCluster.EncryptionKeyResp.Status)
	}

	fAReplicaResourceModel.Tags = []commonTerraform.Tag{}
	for _, v := range responseCluster.Tags {
		fAReplicaResourceModel.Tags = append(fAReplicaResourceModel.Tags, commonTerraform.Tag{
			TagName: types.StringValue(v.TagName),
			Color:   basetypes.NewStringPointerValue(v.Color),
		})
	}

	return nil
}

func (r *FAReplicaResource) buildRequestBah(ctx context.Context, fAReplicaResourceModel FAReplicaResourceModel) (svAccIds, principalIds *[]string, err error) {
	sourceCluster, err := r.client.Read(ctx, fAReplicaResourceModel.ProjectId, *fAReplicaResourceModel.ReplicaSourceClusterId)
	if err != nil {
		return nil, nil, err
	}

	if strings.Contains(sourceCluster.Provider.CloudProviderId, "bah") {
		// If there is an existing Principal Account Id for that Region, use that one.
		pids, err := r.client.GetPeAllowedPrincipalIds(ctx, fAReplicaResourceModel.ProjectId, sourceCluster.Provider.CloudProviderId, fAReplicaResourceModel.Region.ValueString())
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
			for _, peId := range fAReplicaResourceModel.PeAllowedPrincipalIds.Elements() {
				plist = append(plist, peId.(basetypes.StringValue).ValueString())
			}

			principalIds = utils.ToPointer(plist)
		}

		if sourceCluster.Provider.CloudProviderId == "bah:gcp" {
			// If there is an existing Service Account Id for that Region, use that one.
			sids, _ := r.client.GetServiceAccountIds(ctx, fAReplicaResourceModel.ProjectId, sourceCluster.Provider.CloudProviderId, fAReplicaResourceModel.Region.ValueString())
			svAccIds = utils.ToPointer(sids.Data)

			// If there is no existing value, user should provide one
			if svAccIds != nil && len(*svAccIds) == 0 {
				// Here, we prefer to create a non-nil zero length slice, because we need empty JSON array
				// while encoding JSON objects.
				// For more info, please visit https://github.com/golang/go/wiki/CodeReviewComments#declaring-empty-slices
				slist := []string{}
				for _, saId := range fAReplicaResourceModel.ServiceAccountIds.Elements() {
					slist = append(slist, saId.(basetypes.StringValue).ValueString())
				}

				svAccIds = utils.ToPointer(slist)
			}
		}
	}
	return
}

func (r *FAReplicaResource) generateGenericFAReplicaModel(ctx context.Context, fAReplicaResourceModel FAReplicaResourceModel) (models.Cluster, error) {
	cluster := models.Cluster{
		ReplicaSourceClusterId: fAReplicaResourceModel.ReplicaSourceClusterId,
		ClusterName:            fAReplicaResourceModel.ClusterName.ValueStringPointer(),
		ClusterType:            utils.ToPointer("faraway_replica"),
		Region:                 &models.Region{Id: fAReplicaResourceModel.Region.ValueString()},
		Storage: &models.Storage{
			VolumePropertiesId: fAReplicaResourceModel.Storage.VolumeProperties.ValueStringPointer(),
			VolumeTypeId:       fAReplicaResourceModel.Storage.VolumeType.ValueStringPointer(),
			Iops:               fAReplicaResourceModel.Storage.Iops.ValueStringPointer(),
			Size:               fAReplicaResourceModel.Storage.Size.ValueStringPointer(),
			Throughput:         fAReplicaResourceModel.Storage.Throughput.ValueStringPointer(),
		},
		InstanceType:          &models.InstanceType{InstanceTypeId: fAReplicaResourceModel.InstanceType.ValueString()},
		CSPAuth:               fAReplicaResourceModel.CspAuth.ValueBoolPointer(),
		PrivateNetworking:     fAReplicaResourceModel.PrivateNetworking.ValueBoolPointer(),
		BackupRetentionPeriod: fAReplicaResourceModel.BackupRetentionPeriod.ValueStringPointer(),
		BackupScheduleTime:    fAReplicaResourceModel.BackupScheduleTime.ValueStringPointer(),
		VolumeSnapshot:        fAReplicaResourceModel.VolumeSnapshot.ValueBoolPointer(),
	}

	if fAReplicaResourceModel.WalStorage != nil {
		cluster.WalStorage = &models.Storage{
			VolumePropertiesId: fAReplicaResourceModel.WalStorage.VolumeProperties.ValueStringPointer(),
			VolumeTypeId:       fAReplicaResourceModel.WalStorage.VolumeType.ValueStringPointer(),
			Iops:               fAReplicaResourceModel.WalStorage.Iops.ValueStringPointer(),
			Size:               fAReplicaResourceModel.WalStorage.Size.ValueStringPointer(),
			Throughput:         fAReplicaResourceModel.WalStorage.Throughput.ValueStringPointer(),
		}
	}

	cluster.AllowedIpRanges = buildRequestAllowedIpRanges(fAReplicaResourceModel.AllowedIpRanges)

	configs := []models.KeyValue{}
	for _, model := range fAReplicaResourceModel.PgConfig {
		configs = append(configs, models.KeyValue{
			Name:  model.Name,
			Value: model.Value,
		})
	}
	cluster.PgConfig = &configs

	svAccIds, principalIds, err := r.buildRequestBah(ctx, fAReplicaResourceModel)
	if err != nil {
		return models.Cluster{}, err
	}

	cluster.ServiceAccountIds = svAccIds
	cluster.PeAllowedPrincipalIds = principalIds

	if fAReplicaResourceModel.TransparentDataEncryption != nil {
		if !fAReplicaResourceModel.TransparentDataEncryption.KeyId.IsNull() {
			cluster.EncryptionKeyIdReq = fAReplicaResourceModel.TransparentDataEncryption.KeyId.ValueStringPointer()
		}
	}

	tags := []commonApi.Tag{}
	for _, tag := range fAReplicaResourceModel.Tags {
		tags = append(tags, commonApi.Tag{
			Color:   tag.Color.ValueStringPointer(),
			TagName: tag.TagName.ValueString(),
		})
	}
	cluster.Tags = tags

	return cluster, nil
}

func (r *FAReplicaResource) makeFaReplicaForUpdate(ctx context.Context, fAReplicaResourceModel FAReplicaResourceModel) (*models.Cluster, error) {
	fAReplicaModel, err := r.generateGenericFAReplicaModel(ctx, fAReplicaResourceModel)
	if err != nil {
		return nil, err
	}
	fAReplicaModel.Region = nil
	fAReplicaModel.ReplicaSourceClusterId = nil
	fAReplicaModel.EncryptionKeyIdReq = nil
	return &fAReplicaModel, nil
}
