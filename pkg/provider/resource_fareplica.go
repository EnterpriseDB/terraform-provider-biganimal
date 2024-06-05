package provider

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/EnterpriseDB/terraform-provider-biganimal/pkg/api"
	"github.com/EnterpriseDB/terraform-provider-biganimal/pkg/models"
	"github.com/EnterpriseDB/terraform-provider-biganimal/pkg/utils"
	"github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
)

type FAReplicaResource struct {
	client *api.ClusterClient
}

type FAReplicaResourceModel struct {
	ID                     types.String                   `tfsdk:"id"`
	CspAuth                types.Bool                     `tfsdk:"csp_auth"`
	Region                 types.String                   `tfsdk:"region"`
	InstanceType           types.String                   `tfsdk:"instance_type"`
	ResizingPvc            types.List                     `tfsdk:"resizing_pvc"`
	MetricsUrl             *string                        `tfsdk:"metrics_url"`
	ClusterId              *string                        `tfsdk:"cluster_id"`
	ReplicaSourceClusterId *string                        `tfsdk:"source_cluster_id"`
	Phase                  *string                        `tfsdk:"phase"`
	ConnectionUri          types.String                   `tfsdk:"connection_uri"`
	ClusterName            types.String                   `tfsdk:"cluster_name"`
	Storage                *StorageResourceModel          `tfsdk:"storage"`
	PgConfig               []PgConfigResourceModel        `tfsdk:"pg_config"`
	ProjectId              string                         `tfsdk:"project_id"`
	LogsUrl                *string                        `tfsdk:"logs_url"`
	BackupRetentionPeriod  types.String                   `tfsdk:"backup_retention_period"`
	PrivateNetworking      types.Bool                     `tfsdk:"private_networking"`
	AllowedIpRanges        []AllowedIpRangesResourceModel `tfsdk:"allowed_ip_ranges"`
	CreatedAt              types.String                   `tfsdk:"created_at"`
	ServiceAccountIds      types.Set                      `tfsdk:"service_account_ids"`
	PeAllowedPrincipalIds  types.Set                      `tfsdk:"pe_allowed_principal_ids"`

	Timeouts timeouts.Value `tfsdk:"timeouts"`
}

func NewFAReplicaResource() resource.Resource {
	return &FAReplicaResource{}
}

func (r *FAReplicaResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
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
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.UseStateForUnknown(),
				},
			},
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
				Description: "Instance type. For example, \"azure:Standard_D2s_v3\", \"aws:c5.large\" or \"gcp:e2-highcpu-4\".",
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
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"project_id": schema.StringAttribute{
				Description: "BigAnimal Project ID.",
				Optional:    true,
				Validators: []validator.String{
					ProjectIdValidator(),
				},
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
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
				PlanModifiers: []planmodifier.List{
					listplanmodifier.UseStateForUnknown(),
				},
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
		},
	}
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

	if err := r.ensureClusterIsHealthy(ctx, config, timeout); err != nil {
		if !appendDiagFromBAErr(err, &resp.Diagnostics) {
			resp.Diagnostics.AddError("Error waiting for the cluster is ready ", err.Error())
		}
		return
	}

	if err := r.read(ctx, &config); err != nil {
		if !appendDiagFromBAErr(err, &resp.Diagnostics) {
			resp.Diagnostics.AddError("Error reading cluster", err.Error())
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

	if err := r.read(ctx, &state); err != nil {
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

	if err := r.ensureClusterIsHealthy(ctx, plan, timeout); err != nil {
		if !appendDiagFromBAErr(err, &resp.Diagnostics) {
			resp.Diagnostics.AddError("Error waiting faraway replica to be ready ", err.Error())
		}
		return
	}

	if err := r.read(ctx, &plan); err != nil {
		if !appendDiagFromBAErr(err, &resp.Diagnostics) {
			resp.Diagnostics.AddError("Error reading faraway replica", err.Error())
		}
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

func (r *FAReplicaResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
}

func (r FAReplicaResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
}

func (r *FAReplicaResource) read(ctx context.Context, fAReplicaResourceModel *FAReplicaResourceModel) error {
	apiCluster, err := r.client.Read(ctx, fAReplicaResourceModel.ProjectId, *fAReplicaResourceModel.ClusterId)
	if err != nil {
		return err
	}

	connection, err := r.client.ConnectionString(ctx, fAReplicaResourceModel.ProjectId, *fAReplicaResourceModel.ClusterId)
	if err != nil {
		return err
	}

	fAReplicaResourceModel.ID = types.StringValue(fmt.Sprintf("%s/%s", fAReplicaResourceModel.ProjectId, *fAReplicaResourceModel.ClusterId))
	fAReplicaResourceModel.ClusterId = apiCluster.ClusterId
	fAReplicaResourceModel.ClusterName = types.StringPointerValue(apiCluster.ClusterName)
	fAReplicaResourceModel.Phase = apiCluster.Phase
	fAReplicaResourceModel.Region = types.StringValue(apiCluster.Region.Id)
	fAReplicaResourceModel.InstanceType = types.StringValue(apiCluster.InstanceType.InstanceTypeId)
	fAReplicaResourceModel.Storage = &StorageResourceModel{
		VolumeType:       types.StringPointerValue(apiCluster.Storage.VolumeTypeId),
		VolumeProperties: types.StringPointerValue(apiCluster.Storage.VolumePropertiesId),
		Size:             types.StringPointerValue(apiCluster.Storage.Size),
		Iops:             types.StringPointerValue(apiCluster.Storage.Iops),
		Throughput:       types.StringPointerValue(apiCluster.Storage.Throughput),
	}
	fAReplicaResourceModel.ResizingPvc = StringSliceToList(apiCluster.ResizingPvc)
	fAReplicaResourceModel.ConnectionUri = types.StringPointerValue(&connection.PgUri)
	fAReplicaResourceModel.CspAuth = types.BoolPointerValue(apiCluster.CSPAuth)
	fAReplicaResourceModel.LogsUrl = apiCluster.LogsUrl
	fAReplicaResourceModel.MetricsUrl = apiCluster.MetricsUrl
	fAReplicaResourceModel.BackupRetentionPeriod = types.StringPointerValue(apiCluster.BackupRetentionPeriod)
	fAReplicaResourceModel.PrivateNetworking = types.BoolPointerValue(apiCluster.PrivateNetworking)

	// pgConfig. If tf resource pg config elem matches with api response pg config elem then add the elem to tf resource pg config
	newPgConfig := []PgConfigResourceModel{}
	if configs := apiCluster.PgConfig; configs != nil {
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

	fAReplicaResourceModel.AllowedIpRanges = []AllowedIpRangesResourceModel{}
	if allowedIpRanges := apiCluster.AllowedIpRanges; allowedIpRanges != nil {
		for _, ipRange := range *allowedIpRanges {
			fAReplicaResourceModel.AllowedIpRanges = append(fAReplicaResourceModel.AllowedIpRanges, AllowedIpRangesResourceModel{
				CidrBlock:   ipRange.CidrBlock,
				Description: types.StringValue(ipRange.Description),
			})
		}
	}

	if pt := apiCluster.CreatedAt; pt != nil {
		fAReplicaResourceModel.CreatedAt = types.StringValue(pt.String())
	}

	if apiCluster.PeAllowedPrincipalIds != nil {
		fAReplicaResourceModel.PeAllowedPrincipalIds = StringSliceToSet(utils.ToValue(&apiCluster.PeAllowedPrincipalIds))
	}

	if apiCluster.ServiceAccountIds != nil {
		fAReplicaResourceModel.ServiceAccountIds = StringSliceToSet(utils.ToValue(&apiCluster.ServiceAccountIds))
	}

	return nil
}

func (r *FAReplicaResource) ensureClusterIsHealthy(ctx context.Context, cluster FAReplicaResourceModel, timeout time.Duration) error {
	return retry.RetryContext(
		ctx,
		timeout,
		func() *retry.RetryError {
			resp, err := r.client.Read(ctx, cluster.ProjectId, *cluster.ClusterId)
			if err != nil {
				return retry.NonRetryableError(err)
			}

			if !resp.IsHealthy() {
				return retry.RetryableError(errors.New("faraway-replica not yet ready"))
			}
			return nil
		})
}

func (r *FAReplicaResource) buildRequestBah(ctx context.Context, fAReplicaResourceModel FAReplicaResourceModel) (svAccIds, principalIds *[]string, err error) {
	sourceCluster, err := r.client.Read(ctx, fAReplicaResourceModel.ProjectId, *fAReplicaResourceModel.ReplicaSourceClusterId)
	if err != nil {
		return nil, nil, err
	}

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
			plist = append(plist, strings.Replace(peId.String(), "\"", "", -1))
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
				slist = append(slist, strings.Replace(saId.String(), "\"", "", -1))
			}

			svAccIds = utils.ToPointer(slist)
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
	}

	allowedIpRanges := []models.AllowedIpRange{}
	for _, ipRange := range fAReplicaResourceModel.AllowedIpRanges {
		allowedIpRanges = append(allowedIpRanges, models.AllowedIpRange{
			CidrBlock:   ipRange.CidrBlock,
			Description: ipRange.Description.ValueString(),
		})
	}
	cluster.AllowedIpRanges = &allowedIpRanges

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

	return cluster, nil
}

func (r *FAReplicaResource) makeFaReplicaForUpdate(ctx context.Context, fAReplicaResourceModel FAReplicaResourceModel) (*models.Cluster, error) {
	fAReplicaModel, err := r.generateGenericFAReplicaModel(ctx, fAReplicaResourceModel)
	if err != nil {
		return nil, err
	}
	fAReplicaModel.Region = nil
	fAReplicaModel.ReplicaSourceClusterId = nil
	return &fAReplicaModel, nil
}
