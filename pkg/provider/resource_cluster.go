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
	ConnectionUri              *string                            `tfsdk:"connection_uri"`
	ClusterName                types.String                       `tfsdk:"cluster_name"`
	RoConnectionUri            *string                            `tfsdk:"ro_connection_uri"`
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
	ServiceAccountIds          []string                           `tfsdk:"service_account_ids"`
	PeAllowedPrincipalIds      []string                           `tfsdk:"pe_allowed_principal_ids"`

	Timeouts timeouts.Value `tfsdk:"timeouts"`
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
						MarkdownDescription: "Throughput is automatically calculated by BigAnimal based on the IOPS input.",
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
				PlanModifiers:       []planmodifier.Set{plan_modifier.CustomPGConfig()},
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
				PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"cluster_name": schema.StringAttribute{
				MarkdownDescription: "Name of the cluster.",
				Required:            true,
			},
			"phase": schema.StringAttribute{
				MarkdownDescription: "Current phase of the cluster.",
				Computed:            true,
				PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},

			"ro_connection_uri": schema.StringAttribute{
				MarkdownDescription: "Cluster read-only connection URI. Only available for high availability clusters.",
				Computed:            true,
				PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
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
				Description: "Cloud provider. For example, \"aws\", \"azure\" or \"gcp\".",
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
			"service_account_ids": schema.ListAttribute{
				MarkdownDescription: "A Google Cloud Service Account is used for logs. If you leave this blank, then you will be unable to access log details for this cluster. Required when cluster is deployed on BigAnimal's cloud account.",
				Optional:            true,
				Computed:            true,
			},

			"pe_allowed_principal_ids": schema.ListAttribute{
				MarkdownDescription: "Cloud provider subscription/account ID, need to be specified when cluster is deployed on BigAnimal's cloud account.",
				Optional:            true,
				Computed:            true,
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

	clusterId, err := c.client.Create(ctx, config.ProjectId, c.makeClusterForCreate(ctx, config))
	if err != nil {
		if !appendDiagFromBAErr(err, &resp.Diagnostics) {
			resp.Diagnostics.AddError("Error creating cluster", err.Error())
		}
		return
	}

	config.ClusterId = &clusterId

	timeout, diagnostics := config.Timeouts.Create(ctx, time.Minute*60)
	resp.Diagnostics.Append(diagnostics...)
	if resp.Diagnostics.HasError() {
		return
	}

	if err := c.ensureClusterIsHealthy(ctx, config, timeout); err != nil {
		if !appendDiagFromBAErr(err, &resp.Diagnostics) {
			resp.Diagnostics.AddError("Error waiting for the cluster is ready ", err.Error())
		}
		return
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
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	_, err := c.client.Update(ctx, c.makeClusterFoUpdate(ctx, plan), plan.ProjectId, *plan.ClusterId)
	if err != nil {
		if !appendDiagFromBAErr(err, &resp.Diagnostics) {
			resp.Diagnostics.AddError("Error updating cluster", err.Error())
		}
		return
	}

	timeout, diagnostics := plan.Timeouts.Update(ctx, time.Minute*60)
	resp.Diagnostics.Append(diagnostics...)
	if err := c.ensureClusterIsHealthy(ctx, plan, timeout); err != nil {
		if !appendDiagFromBAErr(err, &resp.Diagnostics) {
			resp.Diagnostics.AddError("Error waiting for the cluster is ready ", err.Error())
		}
		return
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

func (c *clusterResource) read(ctx context.Context, clusterResource *ClusterResourceModel) error {
	cluster, err := c.client.Read(ctx, clusterResource.ProjectId, *clusterResource.ClusterId)
	if err != nil {
		return err
	}

	connection, err := c.client.ConnectionString(ctx, clusterResource.ProjectId, *clusterResource.ClusterId)
	if err != nil {
		return err
	}

	clusterResource.ID = types.StringValue(fmt.Sprintf("%s/%s", clusterResource.ProjectId, *clusterResource.ClusterId))
	clusterResource.ClusterId = cluster.ClusterId
	clusterResource.ClusterName = types.StringPointerValue(cluster.ClusterName)
	clusterResource.ClusterType = cluster.ClusterType
	clusterResource.Phase = cluster.Phase
	clusterResource.CloudProvider = types.StringValue(cluster.Provider.CloudProviderId)
	clusterResource.ClusterArchitecture = &ClusterArchitectureResourceModel{
		Id:    cluster.ClusterArchitecture.ClusterArchitectureId,
		Nodes: cluster.ClusterArchitecture.Nodes,
		Name:  types.StringValue(cluster.ClusterArchitecture.ClusterArchitectureName),
	}
	clusterResource.Region = types.StringValue(cluster.Region.Id)
	clusterResource.InstanceType = types.StringValue(cluster.InstanceType.InstanceTypeId)
	clusterResource.Storage = &StorageResourceModel{
		VolumeType:       types.StringPointerValue(cluster.Storage.VolumeTypeId),
		VolumeProperties: types.StringPointerValue(cluster.Storage.VolumePropertiesId),
		Size:             types.StringPointerValue(cluster.Storage.Size),
		Iops:             types.StringPointerValue(cluster.Storage.Iops),
		Throughput:       types.StringPointerValue(cluster.Storage.Throughput),
	}
	clusterResource.ResizingPvc = StringSliceToList(cluster.ResizingPvc)
	clusterResource.ReadOnlyConnections = types.BoolPointerValue(cluster.ReadOnlyConnections)
	clusterResource.ConnectionUri = &connection.PgUri
	clusterResource.RoConnectionUri = &connection.ReadOnlyPgUri
	clusterResource.CspAuth = types.BoolPointerValue(cluster.CSPAuth)
	clusterResource.LogsUrl = cluster.LogsUrl
	clusterResource.MetricsUrl = cluster.MetricsUrl
	clusterResource.BackupRetentionPeriod = types.StringPointerValue(cluster.BackupRetentionPeriod)
	clusterResource.PgVersion = types.StringValue(cluster.PgVersion.PgVersionId)
	clusterResource.PgType = types.StringValue(cluster.PgType.PgTypeId)
	clusterResource.FarawayReplicaIds = StringSliceToSet(cluster.FarawayReplicaIds)
	clusterResource.PrivateNetworking = types.BoolPointerValue(cluster.PrivateNetworking)

	if cluster.FirstRecoverabilityPointAt != nil {
		firstPointAt := cluster.FirstRecoverabilityPointAt.String()
		clusterResource.FirstRecoverabilityPointAt = &firstPointAt
	}

	clusterResource.PgConfig = []PgConfigResourceModel{}
	if configs := cluster.PgConfig; configs != nil {
		for _, kv := range *configs {
			clusterResource.PgConfig = append(clusterResource.PgConfig, PgConfigResourceModel{
				Name:  kv.Name,
				Value: kv.Value,
			})
		}
	}

	clusterResource.AllowedIpRanges = []AllowedIpRangesResourceModel{}
	if allowedIpRanges := cluster.AllowedIpRanges; allowedIpRanges != nil {
		for _, ipRange := range *allowedIpRanges {
			clusterResource.AllowedIpRanges = append(clusterResource.AllowedIpRanges, AllowedIpRangesResourceModel{
				CidrBlock:   ipRange.CidrBlock,
				Description: types.StringValue(ipRange.Description),
			})
		}
	}

	if pt := cluster.CreatedAt; pt != nil {
		clusterResource.CreatedAt = types.StringValue(pt.String())
	}

	if cluster.MaintenanceWindow != nil {
		clusterResource.MaintenanceWindow = &commonTerraform.MaintenanceWindow{
			IsEnabled: cluster.MaintenanceWindow.IsEnabled,
			StartDay:  types.Int64PointerValue(utils.ToPointer(int64(*cluster.MaintenanceWindow.StartDay))),
			StartTime: types.StringPointerValue(cluster.MaintenanceWindow.StartTime),
		}
	}

	if cluster.PeAllowedPrincipalIds != nil {
		var pids []string
		for _, v := range *cluster.PeAllowedPrincipalIds {
			pids = append(pids, v)
		}
		clusterResource.PeAllowedPrincipalIds = pids
	}

	if cluster.ServiceAccountIds != nil {
		var saIds []string
		for _, v := range *cluster.ServiceAccountIds {
			saIds = append(saIds, v)
		}
		clusterResource.PeAllowedPrincipalIds = saIds
	}

	return nil
}

func (c *clusterResource) ensureClusterIsHealthy(ctx context.Context, cluster ClusterResourceModel, timeout time.Duration) error {
	return retry.RetryContext(
		ctx,
		timeout,
		func() *retry.RetryError {
			resp, err := c.client.Read(ctx, cluster.ProjectId, *cluster.ClusterId)
			if err != nil {
				return retry.NonRetryableError(err)
			}

			if !resp.IsHealthy() {
				return retry.RetryableError(errors.New("cluster not yet ready"))
			}
			return nil
		})
}

func (c *clusterResource) makeClusterForCreate(ctx context.Context, clusterResource ClusterResourceModel) models.Cluster {
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

	if strings.Contains(clusterResource.CloudProvider.ValueString(), "bah") {
		pids, _ := c.client.GetPeAllowedPrincipalIds(ctx, clusterResource.ProjectId, clusterResource.CloudProvider.ValueString(), clusterResource.Region.ValueString())

		if clusterResource.PeAllowedPrincipalIds != nil {
			cluster.PeAllowedPrincipalIds = utils.ToPointer(clusterResource.PeAllowedPrincipalIds)
		} else if len(pids.Data) != 0 {
			cluster.PeAllowedPrincipalIds = utils.ToPointer(pids.Data)
		}

		if clusterResource.CloudProvider.ValueString() == "bah:gcp" {
			sids, _ := c.client.GetServiceAccountIds(ctx, clusterResource.ProjectId, clusterResource.CloudProvider.ValueString(), clusterResource.Region.ValueString())
			if clusterResource.ServiceAccountIds != nil {
				cluster.ServiceAccountIds = utils.ToPointer(clusterResource.ServiceAccountIds)
			} else if len(sids.Data) != 0 {
				cluster.ServiceAccountIds = utils.ToPointer(sids.Data)
			}
		}
	}

	if clusterResource.PeAllowedPrincipalIds != nil {
		cluster.PeAllowedPrincipalIds = utils.ToPointer(clusterResource.PeAllowedPrincipalIds)
	}

	return cluster
}

func (c *clusterResource) makeClusterFoUpdate(ctx context.Context, clusterResource ClusterResourceModel) *models.Cluster {
	cluster := c.makeClusterForCreate(ctx, clusterResource)
	cluster.ClusterId = nil
	cluster.PgType = nil
	cluster.PgVersion = nil
	cluster.Provider = nil
	cluster.Region = nil
	return &cluster
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
