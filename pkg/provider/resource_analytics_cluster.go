package provider

import (
	"context"
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
	"github.com/hashicorp/terraform-plugin-framework/path"
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
	_ resource.Resource              = &analyticsClusterResource{}
	_ resource.ResourceWithConfigure = &analyticsClusterResource{}
)

type analyticsClusterResourceModel struct {
	ID                         types.String                       `tfsdk:"id"`
	CspAuth                    types.Bool                         `tfsdk:"csp_auth"`
	Region                     types.String                       `tfsdk:"region"`
	InstanceType               types.String                       `tfsdk:"instance_type"`
	ResizingPvc                types.List                         `tfsdk:"resizing_pvc"`
	MetricsUrl                 *string                            `tfsdk:"metrics_url"`
	ClusterId                  *string                            `tfsdk:"cluster_id"`
	Phase                      *string                            `tfsdk:"phase"`
	ConnectionUri              types.String                       `tfsdk:"connection_uri"`
	ClusterName                types.String                       `tfsdk:"cluster_name"`
	FirstRecoverabilityPointAt *string                            `tfsdk:"first_recoverability_point_at"`
	ProjectId                  string                             `tfsdk:"project_id"`
	LogsUrl                    *string                            `tfsdk:"logs_url"`
	BackupRetentionPeriod      types.String                       `tfsdk:"backup_retention_period"`
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
	Pause                      types.Bool                         `tfsdk:"pause"`

	Timeouts timeouts.Value `tfsdk:"timeouts"`
}

func (r analyticsClusterResourceModel) projectId() string {
	return r.ProjectId
}

func (r analyticsClusterResourceModel) clusterId() string {
	return *r.ClusterId
}

func (c *analyticsClusterResourceModel) setPhase(phase string) {
	c.Phase = &phase
}

func (c *analyticsClusterResourceModel) setPgIdentity(pgIdentity string) {
}

func (c *analyticsClusterResourceModel) setCloudProvider(cloudProvider string) {
	c.CloudProvider = types.StringValue(cloudProvider)
}

type analyticsClusterResource struct {
	client *api.ClusterClient
}

func (r *analyticsClusterResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*api.API).ClusterClient()
}

func (r *analyticsClusterResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_analytics_cluster"
}

func (r *analyticsClusterResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
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
				PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
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
				PlanModifiers:       []planmodifier.String{plan_modifier.CustomPhaseForUnknown()},
			},
			"project_id": schema.StringAttribute{
				MarkdownDescription: "BigAnimal Project ID.",
				Required:            true,
				Validators:          []validator.String{ProjectIdValidator()},
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
				Validators:          []validator.String{BackupRetentionPeriodValidator()},
				PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"cloud_provider": schema.StringAttribute{
				Description: "Cloud provider. For example, \"aws\" or \"bah:aws\".",
				Required:    true,
			},
			"pg_type": schema.StringAttribute{
				MarkdownDescription: "Postgres type. For example, \"epas\" or \"pgextended\".",
				Required:            true,
				Validators:          []validator.String{stringvalidator.OneOf("epas", "pgextended", "postgres")},
			},
			"first_recoverability_point_at": schema.StringAttribute{
				MarkdownDescription: "Earliest backup recover time.",
				Computed:            true,
				PlanModifiers:       []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"pg_version": schema.StringAttribute{
				MarkdownDescription: "Postgres version. For example 16",
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
				MarkdownDescription: "Is authentication handled by the cloud service provider.",
				Optional:            true,
				Computed:            true,
				PlanModifiers:       []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
			},
			"maintenance_window": schema.SingleNestedAttribute{
				MarkdownDescription: "Custom maintenance window.",
				Optional:            true,
				Computed:            true,
				PlanModifiers:       []planmodifier.Object{plan_modifier.MaintenanceWindowForUnknown()},
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

func (r *analyticsClusterResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var config analyticsClusterResourceModel
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	clusterModel, err := generateAnalyticsClusterModelCreate(ctx, r.client, config)
	if err != nil {
		if !appendDiagFromBAErr(err, &resp.Diagnostics) {
			resp.Diagnostics.AddError("Error creating cluster", err.Error())
		}
		return
	}

	// consume cluster create with analytics request
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

	// keep retrying until cluster is healthy
	if err := ensureClusterIsEndStateAs(ctx, r.client, &config, timeout); err != nil {
		if !appendDiagFromBAErr(err, &resp.Diagnostics) {
			resp.Diagnostics.AddError("Error waiting for the cluster is ready ", err.Error())
		}
		return
	}

	if config.Pause.ValueBool() {
		_, err = r.client.ClusterPause(ctx, config.ProjectId, *config.ClusterId)
		if err != nil {
			if !appendDiagFromBAErr(err, &resp.Diagnostics) {
				resp.Diagnostics.AddError("Error pausing cluster API request", err.Error())
			}
			return
		}

		// keep retrying until cluster is paused
		if err := ensureClusterIsPaused(ctx, r.client, &config, timeout); err != nil {
			if !appendDiagFromBAErr(err, &resp.Diagnostics) {
				resp.Diagnostics.AddError("Error waiting for the cluster to pause", err.Error())
			}
			return
		}
	}

	// after cluster is in the correct state (healthy/paused) then get the cluster and save into state
	if err := read(ctx, r.client, &config); err != nil {
		if !appendDiagFromBAErr(err, &resp.Diagnostics) {
			resp.Diagnostics.AddError("Error reading cluster", err.Error())
		}
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, config)...)
}

// uses generateAnalyticsClusterModelCreate but just comments out some fields that are not needed for update
func generateAnalyticsClusterModelUpdate(ctx context.Context, client *api.ClusterClient, clusterResource analyticsClusterResourceModel) (models.Cluster, error) {
	cluster, err := generateAnalyticsClusterModelCreate(ctx, client, clusterResource)
	if err != nil {
		return models.Cluster{}, err
	}

	cluster.ClusterId = nil
	cluster.PgType = nil
	cluster.PgVersion = nil
	cluster.Provider = nil
	cluster.Region = nil

	return cluster, nil
}

// used for create operation
func generateAnalyticsClusterModelCreate(ctx context.Context, client *api.ClusterClient, clusterResource analyticsClusterResourceModel) (models.Cluster, error) {
	cluster := models.Cluster{
		ClusterType:           utils.ToPointer("analytical"),
		ClusterName:           clusterResource.ClusterName.ValueStringPointer(),
		Password:              clusterResource.Password.ValueStringPointer(),
		Provider:              &models.Provider{CloudProviderId: clusterResource.CloudProvider.ValueString()},
		Region:                &models.Region{Id: clusterResource.Region.ValueString()},
		InstanceType:          &models.InstanceType{InstanceTypeId: clusterResource.InstanceType.ValueString()},
		PgType:                &models.PgType{PgTypeId: clusterResource.PgType.ValueString()},
		PgVersion:             &models.PgVersion{PgVersionId: clusterResource.PgVersion.ValueString()},
		CSPAuth:               clusterResource.CspAuth.ValueBoolPointer(),
		PrivateNetworking:     clusterResource.PrivateNetworking.ValueBoolPointer(),
		BackupRetentionPeriod: clusterResource.BackupRetentionPeriod.ValueStringPointer(),
	}

	cluster.ClusterId = nil
	cluster.PgConfig = nil

	allowedIpRanges := []models.AllowedIpRange{}
	for _, ipRange := range clusterResource.AllowedIpRanges {
		allowedIpRanges = append(allowedIpRanges, models.AllowedIpRange{
			CidrBlock:   ipRange.CidrBlock,
			Description: ipRange.Description.ValueString(),
		})
	}
	cluster.AllowedIpRanges = &allowedIpRanges

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
		clusterRscCSP := clusterResource.CloudProvider
		clusterRscPrincipalIds := clusterResource.PeAllowedPrincipalIds
		clusterRscSvcAcntIds := clusterResource.ServiceAccountIds

		// If there is an existing Principal Account Id for that Region, use that one.
		pids, err := client.GetPeAllowedPrincipalIds(ctx, clusterResource.ProjectId, clusterRscCSP.ValueString(), clusterResource.Region.ValueString())
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
			sids, _ := client.GetServiceAccountIds(ctx, clusterResource.ProjectId, clusterResource.CloudProvider.ValueString(), clusterResource.Region.ValueString())
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
	}

	return cluster, nil
}

func read(ctx context.Context, client *api.ClusterClient, tfClusterResource *analyticsClusterResourceModel) error {
	apiCluster, err := client.Read(ctx, tfClusterResource.ProjectId, *tfClusterResource.ClusterId)
	if err != nil {
		return err
	}

	connection, err := client.ConnectionString(ctx, tfClusterResource.ProjectId, *tfClusterResource.ClusterId)
	if err != nil {
		return err
	}

	tfClusterResource.ID = types.StringValue(fmt.Sprintf("%s/%s", tfClusterResource.ProjectId, *tfClusterResource.ClusterId))
	tfClusterResource.ClusterId = apiCluster.ClusterId
	tfClusterResource.ClusterName = types.StringPointerValue(apiCluster.ClusterName)
	tfClusterResource.Phase = apiCluster.Phase
	tfClusterResource.CloudProvider = types.StringValue(apiCluster.Provider.CloudProviderId)
	tfClusterResource.Region = types.StringValue(apiCluster.Region.Id)
	tfClusterResource.InstanceType = types.StringValue(apiCluster.InstanceType.InstanceTypeId)
	tfClusterResource.ResizingPvc = StringSliceToList(apiCluster.ResizingPvc)
	tfClusterResource.ConnectionUri = types.StringPointerValue(&connection.PgUri)
	tfClusterResource.CspAuth = types.BoolPointerValue(apiCluster.CSPAuth)
	tfClusterResource.LogsUrl = apiCluster.LogsUrl
	tfClusterResource.MetricsUrl = apiCluster.MetricsUrl
	tfClusterResource.BackupRetentionPeriod = types.StringPointerValue(apiCluster.BackupRetentionPeriod)
	tfClusterResource.PgVersion = types.StringValue(apiCluster.PgVersion.PgVersionId)
	tfClusterResource.PgType = types.StringValue(apiCluster.PgType.PgTypeId)
	tfClusterResource.PrivateNetworking = types.BoolPointerValue(apiCluster.PrivateNetworking)

	if apiCluster.FirstRecoverabilityPointAt != nil {
		firstPointAt := apiCluster.FirstRecoverabilityPointAt.String()
		tfClusterResource.FirstRecoverabilityPointAt = &firstPointAt
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

	return nil
}

func (r *analyticsClusterResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state analyticsClusterResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	if err := read(ctx, r.client, &state); err != nil {
		if !appendDiagFromBAErr(err, &resp.Diagnostics) {
			resp.Diagnostics.AddError("Error reading cluster", err.Error())
		}
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}

func (r *analyticsClusterResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan analyticsClusterResourceModel

	timeout, diagnostics := plan.Timeouts.Update(ctx, time.Minute*60)
	resp.Diagnostics.Append(diagnostics...)

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state analyticsClusterResourceModel
	diags = req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// cluster = pause,   tf pause = true, it will error and say you will need to set pause = false to update
	// cluster = pause,   tf pause = false, it will resume then update
	// cluster = healthy, tf pause = true, it will update then pause
	// cluster = healthy, tf pause = false, it will update
	if *state.Phase != constants.PHASE_HEALTHY && *state.Phase != constants.PHASE_PAUSED {
		resp.Diagnostics.AddError("Cluster not ready please wait", "Cluster not ready for update operation please wait")
		return
	}

	if *state.Phase == constants.PHASE_PAUSED {
		if plan.Pause.ValueBool() {
			resp.Diagnostics.AddError("Error cannot update paused cluster", "cannot update paused cluster, please set pause = false to resume cluster")
			return
		}

		if !plan.Pause.ValueBool() {
			_, err := r.client.ClusterResume(ctx, plan.ProjectId, *plan.ClusterId)
			if err != nil {
				if !appendDiagFromBAErr(err, &resp.Diagnostics) {
					resp.Diagnostics.AddError("Error resuming cluster API request", err.Error())
				}
				return
			}

			if err := ensureClusterIsEndStateAs(ctx, r.client, &plan, timeout); err != nil {
				if !appendDiagFromBAErr(err, &resp.Diagnostics) {
					resp.Diagnostics.AddError("Error waiting for the cluster is ready ", err.Error())
				}
				return
			}
		}
	}

	clusterModel, err := generateAnalyticsClusterModelUpdate(ctx, r.client.ClusterClient(), plan)
	if err != nil {
		if !appendDiagFromBAErr(err, &resp.Diagnostics) {
			resp.Diagnostics.AddError("Error updating cluster", err.Error())
		}
		return
	}

	_, err = r.client.Update(ctx, &clusterModel, plan.ProjectId, *plan.ClusterId)
	if err != nil {
		if !appendDiagFromBAErr(err, &resp.Diagnostics) {
			resp.Diagnostics.AddError("Error updating cluster API request", err.Error())
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

	if plan.Pause.ValueBool() {
		_, err = r.client.ClusterPause(ctx, plan.ProjectId, *plan.ClusterId)
		if err != nil {
			if !appendDiagFromBAErr(err, &resp.Diagnostics) {
				resp.Diagnostics.AddError("Error pausing cluster API request", err.Error())
			}
			return
		}

		if err := ensureClusterIsPaused(ctx, r.client, &plan, timeout); err != nil {
			if !appendDiagFromBAErr(err, &resp.Diagnostics) {
				resp.Diagnostics.AddError("Error waiting for the cluster to pause", err.Error())
			}
			return
		}
	}

	if err := read(ctx, r.client, &plan); err != nil {
		if !appendDiagFromBAErr(err, &resp.Diagnostics) {
			resp.Diagnostics.AddError("Error reading cluster", err.Error())
		}
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

func (r *analyticsClusterResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state analyticsClusterResourceModel

	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	err := r.client.Delete(ctx, state.ProjectId, *state.ClusterId)
	if err != nil {
		if !appendDiagFromBAErr(err, &resp.Diagnostics) {
			resp.Diagnostics.AddError("Error deleting cluster", err.Error())
		}
		return
	}
}

func (r *analyticsClusterResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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

func NewAnalyticsClusterResource() resource.Resource {
	return &analyticsClusterResource{}
}
