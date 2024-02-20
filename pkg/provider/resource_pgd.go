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
	pgdApi "github.com/EnterpriseDB/terraform-provider-biganimal/pkg/models/pgd/api"
	"github.com/EnterpriseDB/terraform-provider-biganimal/pkg/models/pgd/terraform"
	"github.com/EnterpriseDB/terraform-provider-biganimal/pkg/plan_modifier"
	"github.com/EnterpriseDB/terraform-provider-biganimal/pkg/utils"
	"github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/float64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
)

var (
	_ resource.Resource                = &pgdResource{}
	_ resource.ResourceWithConfigure   = &pgdResource{}
	_ resource.ResourceWithImportState = &pgdResource{}
)

type pgdResource struct {
	client *api.PGDClient
}

func PgdSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		MarkdownDescription: "The PGD cluster data source describes a BigAnimal cluster. The data source requires your PGD cluster name.",
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
			"project_id": schema.StringAttribute{
				Description: "BigAnimal Project ID.",
				Optional:    true,
				Validators: []validator.String{
					ProjectIdValidator(),
				},
			},
			"cluster_id": schema.StringAttribute{
				Description: "Cluster ID.",
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"cluster_name": schema.StringAttribute{
				Description: "cluster name",
				Required:    true,
			},
			"most_recent": schema.BoolAttribute{
				Description: "Show the most recent cluster when there are multiple clusters with the same name",
				Optional:    true,
			},
			"password": schema.StringAttribute{
				Description: "Password for the user edb_admin. It must be 12 characters or more.",
				Required:    true,
				Sensitive:   true,
			},
			"data_groups": schema.SetNestedAttribute{
				Description: "Cluster data groups.",
				Required:    true,
				PlanModifiers: []planmodifier.Set{
					plan_modifier.CustomDataGroupDiffConfig(),
				},
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"group_id": schema.StringAttribute{
							Description: "Group ID of the group.",
							Computed:    true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"backup_retention_period": schema.StringAttribute{
							Description: "Backup retention period",
							Required:    true,
							Validators: []validator.String{
								BackupRetentionPeriodValidator(),
							},
						},
						"cluster_name": schema.StringAttribute{
							Description: "Name of the group.",
							Computed:    true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"cluster_type": schema.StringAttribute{
							Description: "Type of the Specified Cluster",
							Optional:    true,
							Computed:    true,
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
						"connection_uri": schema.StringAttribute{
							Description: "Data group connection URI.",
							Computed:    true,
						},
						"phase": schema.StringAttribute{
							Description: "Current phase of the data group.",
							Computed:    true,
							PlanModifiers: []planmodifier.String{
								plan_modifier.CustomPhaseForUnknown(),
							},
						},
						"private_networking": schema.BoolAttribute{
							Description: "Is private networking enabled.",
							Required:    true,
						},
						"csp_auth": schema.BoolAttribute{
							Description: "Is authentication handled by the cloud service provider.",
							Required:    true,
						},
						"resizing_pvc": schema.SetAttribute{
							Description: "Resizing PVC.",
							Computed:    true,
							ElementType: types.StringType,
							PlanModifiers: []planmodifier.Set{
								setplanmodifier.UseStateForUnknown(),
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
						"pg_config": schema.SetNestedAttribute{
							Description: "Database configuration parameters.",
							Required:    true,
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
						"cluster_architecture": schema.SingleNestedAttribute{
							Description: "Cluster architecture.",
							Required:    true,
							Attributes: map[string]schema.Attribute{
								"cluster_architecture_id": schema.StringAttribute{
									Description: "Cluster architecture ID.",
									Required:    true,
									Validators: []validator.String{
										stringvalidator.OneOf("pgd"),
									},
								},
								"cluster_architecture_name": schema.StringAttribute{
									Description: "Cluster architecture name.",
									Computed:    true,
									PlanModifiers: []planmodifier.String{
										stringplanmodifier.UseStateForUnknown(),
									},
								},
								"nodes": schema.Float64Attribute{
									Description: "Node count.",
									Required:    true,
								},
								"witness_nodes": schema.Int64Attribute{
									Description: "Witness nodes count.",
									Computed:    true,
									PlanModifiers: []planmodifier.Int64{
										int64planmodifier.UseStateForUnknown(),
									},
								},
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
						"pg_type": schema.SingleNestedAttribute{
							Description: "Postgres type.",
							Required:    true,
							Attributes: map[string]schema.Attribute{
								"pg_type_id": schema.StringAttribute{
									Description: "Data group pg type id.",
									Required:    true,
								},
							},
						},
						"pg_version": schema.SingleNestedAttribute{
							Description: "Postgres version.",
							Required:    true,
							Attributes: map[string]schema.Attribute{
								"pg_version_id": schema.StringAttribute{
									Description: "Data group pg version id.",
									Required:    true,
								},
							},
						},
						"cloud_provider": schema.SingleNestedAttribute{
							Description: "Cloud provider.",
							Required:    true,
							Attributes: map[string]schema.Attribute{
								"cloud_provider_id": schema.StringAttribute{
									Description: "Data group cloud provider id.",
									Required:    true,
								},
							},
						},
						"region": schema.SingleNestedAttribute{
							Description: "Region.",
							Required:    true,
							Attributes: map[string]schema.Attribute{
								"region_id": schema.StringAttribute{
									Description: "Data group region id.",
									Required:    true,
								},
							},
						},
						"instance_type": schema.SingleNestedAttribute{
							Description: "Instance type.",
							Required:    true,
							Attributes: map[string]schema.Attribute{
								"instance_type_id": schema.StringAttribute{
									Description: "Data group instance type id.",
									Required:    true,
								},
							},
						},
						"maintenance_window": schema.SingleNestedAttribute{
							Description: "Custom maintenance window.",
							Required:    true,
							Attributes: map[string]schema.Attribute{
								"is_enabled": schema.BoolAttribute{
									Description: "Is maintenance window enabled.",
									Required:    true,
								},
								"start_day": schema.Float64Attribute{
									Description: "Start day.",
									Required:    true,
								},
								"start_time": schema.StringAttribute{
									Description: "Start time.",
									Required:    true,
								},
							},
						},
						"conditions": schema.SetNestedAttribute{
							Description: "Conditions.",
							Computed:    true,
							PlanModifiers: []planmodifier.Set{
								setplanmodifier.UseStateForUnknown(),
							},
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"condition_status": schema.StringAttribute{
										Description: "Condition status",
										Computed:    true,
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.UseStateForUnknown(),
										},
									},
									"type": schema.StringAttribute{
										Description: "Type",
										Computed:    true,
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.UseStateForUnknown(),
										},
									},
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
				},
			},
			"witness_groups": schema.SetNestedAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Set{
					plan_modifier.CustomWitnessGroupDiffConfig(),
				},
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"group_id": schema.StringAttribute{
							Description: "Group id of witness group.",
							Computed:    true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"phase": schema.StringAttribute{
							Description: "Current phase of the witness group.",
							Computed:    true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"cluster_architecture": schema.SingleNestedAttribute{
							Description: "Cluster architecture.",
							Computed:    true,
							PlanModifiers: []planmodifier.Object{
								objectplanmodifier.UseStateForUnknown(),
							},
							Attributes: map[string]schema.Attribute{
								"cluster_architecture_id": schema.StringAttribute{
									Description: "Cluster architecture ID.",
									Computed:    true,
									PlanModifiers: []planmodifier.String{
										stringplanmodifier.UseStateForUnknown(),
									},
								},
								"cluster_architecture_name": schema.StringAttribute{
									Description: "Name.",
									Computed:    true,
									PlanModifiers: []planmodifier.String{
										stringplanmodifier.UseStateForUnknown(),
									},
								},
								"nodes": schema.Float64Attribute{
									Description: "Nodes.",
									Computed:    true,
									PlanModifiers: []planmodifier.Float64{
										float64planmodifier.UseStateForUnknown(),
									},
								},
								"witness_nodes": schema.Float64Attribute{
									Description: "Witness nodes count.",
									Computed:    true,
									PlanModifiers: []planmodifier.Float64{
										float64planmodifier.UseStateForUnknown(),
									},
								},
							},
						},
						"region": schema.SingleNestedAttribute{
							Description: "Region.",
							Required:    true,
							Attributes: map[string]schema.Attribute{
								"region_id": schema.StringAttribute{
									Description: "Region id.",
									Required:    true,
								},
							},
						},
						"cluster_type": schema.StringAttribute{
							Description: "Type of the Specified Cluster",
							Computed:    true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"cloud_provider": schema.SingleNestedAttribute{
							Description: "Witness Group cloud provider id. It can be set during creation only and can be different than the cloud provider of the data groups. Once set, cannot be changed.",
							Optional:    true,
							Computed:    true,
							PlanModifiers: []planmodifier.Object{
								objectplanmodifier.UseStateForUnknown(),
							},
							Attributes: map[string]schema.Attribute{
								"cloud_provider_id": schema.StringAttribute{
									Description: "Cloud provider id.",
									Optional:    true,
									Computed:    true,
									PlanModifiers: []planmodifier.String{
										stringplanmodifier.UseStateForUnknown(),
									},
								},
							},
						},
						"instance_type": schema.SingleNestedAttribute{
							Description: "Instance type.",
							Computed:    true,
							PlanModifiers: []planmodifier.Object{
								objectplanmodifier.UseStateForUnknown(),
							},
							Attributes: map[string]schema.Attribute{
								"instance_type_id": schema.StringAttribute{
									Description: "Witness group instance type id.",
									Computed:    true,
									PlanModifiers: []planmodifier.String{
										stringplanmodifier.UseStateForUnknown(),
									},
								},
							},
						},
						"storage": schema.SingleNestedAttribute{
							Description: "Storage.",
							Computed:    true,
							PlanModifiers: []planmodifier.Object{
								objectplanmodifier.UseStateForUnknown(),
							},
							Attributes: map[string]schema.Attribute{
								"iops": schema.StringAttribute{
									Description: "IOPS for the selected volume.",
									Computed:    true,
									PlanModifiers: []planmodifier.String{
										stringplanmodifier.UseStateForUnknown(),
									},
								},
								"size": schema.StringAttribute{
									Description: "Size of the volume.",
									Computed:    true,
									PlanModifiers: []planmodifier.String{
										stringplanmodifier.UseStateForUnknown(),
									},
								},
								"throughput": schema.StringAttribute{
									Description: "Throughput.",
									Computed:    true,
									PlanModifiers: []planmodifier.String{
										stringplanmodifier.UseStateForUnknown(),
									},
								},
								"volume_properties": schema.StringAttribute{
									Description: "Volume properties.",
									Computed:    true,
									PlanModifiers: []planmodifier.String{
										stringplanmodifier.UseStateForUnknown(),
									},
								},
								"volume_type": schema.StringAttribute{
									Description: "Volume type.",
									Computed:    true,
									PlanModifiers: []planmodifier.String{
										stringplanmodifier.UseStateForUnknown(),
									},
								},
							},
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
					},
				},
			},
		},
	}
}

func (p pgdResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_pgd"
}

func (p pgdResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = PgdSchema(ctx)
}

// Configure adds the provider configured client to the data source.
func (p *pgdResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	p.client = req.ProviderData.(*api.API).PGDClient()
}

type PGD struct {
	ID            *string                  `tfsdk:"id"`
	ProjectId     string                   `tfsdk:"project_id"`
	ClusterId     *string                  `tfsdk:"cluster_id"`
	ClusterName   *string                  `tfsdk:"cluster_name"`
	MostRecent    *bool                    `tfsdk:"most_recent"`
	Password      *string                  `tfsdk:"password"`
	Timeouts      timeouts.Value           `tfsdk:"timeouts"`
	DataGroups    []terraform.DataGroup    `tfsdk:"data_groups"`
	WitnessGroups []terraform.WitnessGroup `tfsdk:"witness_groups"`
}

// Create creates the resource and sets the initial Terraform state.
func (p pgdResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var config PGD
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	clusterReqBody := models.Cluster{
		ClusterName: config.ClusterName,
		ClusterType: utils.ToPointer("cluster"),
		Password:    config.Password,
	}

	clusterReqBody.Groups = &[]any{}

	for _, v := range config.DataGroups {
		v := v

		storage := buildApiStorage(*v.Storage)

		if v.PgConfig == nil {
			v.PgConfig = &[]models.KeyValue{}
		}

		clusterArchName := v.ClusterArchitecture.ClusterArchitectureName.ValueStringPointer()
		if v.ClusterArchitecture.ClusterArchitectureName.IsUnknown() {
			clusterArchName = nil
		}

		clusterArch := &pgdApi.ClusterArchitecture{
			ClusterArchitectureId:   v.ClusterArchitecture.ClusterArchitectureId,
			ClusterArchitectureName: clusterArchName,
			Nodes:                   v.ClusterArchitecture.Nodes,
		}

		svAccIds, principalIds := buildApiBah(ctx, p.client, &resp.Diagnostics, config.ProjectId, v)
		if resp.Diagnostics.HasError() {
			return
		}

		apiDGModel := pgdApi.DataGroup{
			AllowedIpRanges:       buildApiAllowedIpRanges(v.AllowedIpRanges),
			BackupRetentionPeriod: v.BackupRetentionPeriod,
			Provider:              v.Provider,
			ClusterArchitecture:   clusterArch,
			CspAuth:               v.CspAuth,
			ClusterType:           utils.ToPointer("data_group"),
			InstanceType:          v.InstanceType,
			MaintenanceWindow:     v.MaintenanceWindow,
			PgConfig:              v.PgConfig,
			PgType:                v.PgType,
			PgVersion:             v.PgVersion,
			PrivateNetworking:     v.PrivateNetworking,
			Region:                v.Region,
			Storage:               storage,
			ServiceAccountIds:     svAccIds,
			PeAllowedPrincipalIds: principalIds,
		}
		*clusterReqBody.Groups = append(*clusterReqBody.Groups, apiDGModel)
	}

	if len(config.WitnessGroups) > 0 {
		for _, wg := range config.WitnessGroups {
			providerId := *config.DataGroups[0].Provider.CloudProviderId

			if !wg.Provider.IsNull() && !wg.Provider.IsUnknown() {
				providerId = strings.Replace(wg.Provider.Attributes()["cloud_provider_id"].String(), "\"", "", -1)
			}

			calWitnessResp, err := p.client.CalculateWitnessGroupParams(ctx, config.ProjectId, pgdApi.WitnessGroupParamsBody{
				Provider: &pgdApi.CloudProvider{
					CloudProviderId: utils.ToPointer(providerId),
				},
				Region: &pgdApi.Region{
					RegionId: wg.Region.RegionId.String(),
				},
			})
			if err != nil {
				if appendDiagFromBAErr(err, &resp.Diagnostics) {
					return
				}
				resp.Diagnostics.AddError("Error calculating witness group params v2", "Could not calculate witness group params v2, unexpected error: "+err.Error())
				return
			}

			wgReq := pgdApi.WitnessGroup{
				ClusterArchitecture: &pgdApi.ClusterArchitecture{
					ClusterArchitectureId: "pgd",
					Nodes:                 config.DataGroups[0].ClusterArchitecture.Nodes,
				},
				ClusterType: utils.ToPointer("witness_group"),
				Provider: &pgdApi.CloudProvider{
					CloudProviderId: utils.ToPointer(providerId),
				},
				InstanceType: calWitnessResp.InstanceType,
				Storage:      calWitnessResp.Storage,
				Region: &pgdApi.Region{
					RegionId: wg.Region.RegionId.ValueString(),
				},
			}

			if wg.MaintenanceWindow != nil {
				wgReq.MaintenanceWindow = &commonApi.MaintenanceWindow{
					IsEnabled: wg.MaintenanceWindow.IsEnabled,
					StartTime: wg.MaintenanceWindow.StartTime.ValueStringPointer(),
				}

				if !wg.MaintenanceWindow.StartDay.IsUnknown() && !wg.MaintenanceWindow.StartDay.IsNull() {
					wgReq.MaintenanceWindow.StartDay = utils.ToPointer(float64(wg.MaintenanceWindow.StartDay.ValueInt64()))
				}
			}

			*clusterReqBody.Groups = append(*clusterReqBody.Groups, wgReq)
		}
	}

	clusterId, err := p.client.Create(ctx, config.ProjectId, clusterReqBody)
	if err != nil {
		if appendDiagFromBAErr(err, &resp.Diagnostics) {
			return
		}
		resp.Diagnostics.AddError("Error creating PGD cluster", "Could not create PGD cluster, unexpected error: "+err.Error())
		return
	}

	config.ID = &clusterId
	config.ClusterId = &clusterId

	// retry func
	timeout, _ := config.Timeouts.Create(ctx, 60*time.Minute)

	err = retry.RetryContext(
		ctx,
		timeout-time.Minute,
		p.retryFuncAs(ctx, &resp.Diagnostics, resp.State, &config),
	)

	if err != nil {
		if appendDiagFromBAErr(err, &resp.Diagnostics) {
			return
		}
		resp.Diagnostics.AddError("Error retrying PGD cluster", "Could not create PGD cluster, unexpected error: "+err.Error())
		return
	}
	// end of retry func

	// uncomment below and comment retry func above to skip retry for testing purposes
	// clusterResp, err := p.client.Read(ctx, config.ProjectId, clusterId)
	// if err != nil {
	// 	resp.Diagnostics.AddError("Error reading PGD cluster", "Could not read PGD cluster, unexpected error: "+err.Error())
	// 	return
	// }

	// buildTFGroupsAs(ctx, &resp.Diagnostics, resp.State, *clusterResp, &config)
	// if resp.Diagnostics.HasError() {
	// 	return
	// }

	diags = resp.State.Set(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (p pgdResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state PGD
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	clusterResp, err := p.client.Read(ctx, state.ProjectId, *state.ClusterId)
	if err != nil {
		if appendDiagFromBAErr(err, &resp.Diagnostics) {
			return
		}
		resp.Diagnostics.AddError("Error reading PGD cluster", "Could not read PGD cluster, unexpected error: "+err.Error())
		return
	}

	state.ID = clusterResp.ClusterId
	state.ClusterId = clusterResp.ClusterId
	state.ClusterName = clusterResp.ClusterName

	buildTFGroupsAs(ctx, &resp.Diagnostics, resp.State, *clusterResp, &state)
	if resp.Diagnostics.HasError() {
		return
	}

	for _, v := range state.DataGroups {
		if v.Phase.ValueString() != models.PHASE_HEALTHY {
			resp.Diagnostics.AddError("Cluster not ready please wait", "Cluster not ready for update operation please wait")
			return
		}
	}

	for _, v := range state.WitnessGroups {
		if v.Phase.ValueString() != models.PHASE_HEALTHY {
			resp.Diagnostics.AddError("Cluster not ready please wait", "Cluster not ready for update operation please wait")
			return
		}
	}

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (p pgdResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan PGD
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	clusterReqBody := models.Cluster{
		ClusterName: plan.ClusterName,
		ClusterType: utils.ToPointer("cluster"),
		Password:    plan.Password,
	}

	clusterReqBody.Groups = &[]any{}

	for _, v := range plan.DataGroups {
		storage := buildApiStorage(*v.Storage)

		groupId := v.GroupId.ValueStringPointer()
		if v.GroupId.IsUnknown() {
			groupId = nil
		}

		svAccIds, principalIds := buildApiBah(ctx, p.client, &resp.Diagnostics, plan.ProjectId, v)
		if resp.Diagnostics.HasError() {
			return
		}

		// only allow fields which are able to be modified in the request for updating
		reqDg := pgdApi.DataGroup{
			GroupId:               groupId,
			ClusterType:           utils.ToPointer("data_group"),
			AllowedIpRanges:       buildApiAllowedIpRanges(v.AllowedIpRanges),
			BackupRetentionPeriod: v.BackupRetentionPeriod,
			CspAuth:               v.CspAuth,
			InstanceType:          v.InstanceType,
			PgConfig:              v.PgConfig,
			PrivateNetworking:     v.PrivateNetworking,
			Storage:               storage,
			MaintenanceWindow:     v.MaintenanceWindow,
			ServiceAccountIds:     svAccIds,
			PeAllowedPrincipalIds: principalIds,
		}

		// signals that it doesn't have an existing group id so this is a new group to add and needs extra fields
		if reqDg.GroupId == nil {
			reqDg.Provider = v.Provider
			reqDg.Region = v.Region
			reqDg.ClusterArchitecture = &pgdApi.ClusterArchitecture{
				ClusterArchitectureId: v.ClusterArchitecture.ClusterArchitectureId,
				Nodes:                 v.ClusterArchitecture.Nodes,
			}
			reqDg.PgType = v.PgType
			reqDg.PgVersion = v.PgVersion
			reqDg.ServiceAccountIds = svAccIds
			reqDg.PeAllowedPrincipalIds = principalIds
		}

		*clusterReqBody.Groups = append(*clusterReqBody.Groups, reqDg)
	}

	if len(plan.WitnessGroups) > 0 {
		for _, wg := range plan.WitnessGroups {
			providerId := *plan.DataGroups[0].Provider.CloudProviderId

			if !wg.Provider.IsNull() && !wg.Provider.IsUnknown() {
				providerId = strings.Replace(wg.Provider.Attributes()["cloud_provider_id"].String(), "\"", "", -1)
			}

			calWitnessResp, err := p.client.CalculateWitnessGroupParams(ctx, plan.ProjectId, pgdApi.WitnessGroupParamsBody{
				Provider: &pgdApi.CloudProvider{
					CloudProviderId: utils.ToPointer(providerId),
				},
				Region: &pgdApi.Region{
					RegionId: wg.Region.RegionId.String(),
				},
			})
			if err != nil {
				if appendDiagFromBAErr(err, &resp.Diagnostics) {
					return
				}
				resp.Diagnostics.AddError("Error calculating witness group params v2", "Could not calculate witness group params v2, unexpected error: "+err.Error())
				return
			}

			// cannot change anything on witness group, this only allows adding a new witness group
			if wg.GroupId.IsNull() || wg.GroupId.IsUnknown() {
				wgReq := pgdApi.WitnessGroup{
					ClusterArchitecture: &pgdApi.ClusterArchitecture{
						ClusterArchitectureId: "pgd",
						Nodes:                 plan.DataGroups[0].ClusterArchitecture.Nodes,
					},
					ClusterType: utils.ToPointer("witness_group"),
					Provider: &pgdApi.CloudProvider{
						CloudProviderId: utils.ToPointer(providerId),
					},
					InstanceType: calWitnessResp.InstanceType,
					Storage:      calWitnessResp.Storage,
					Region: &pgdApi.Region{
						RegionId: wg.Region.RegionId.ValueString(),
					},
				}

				if wg.MaintenanceWindow != nil {
					wgReq.MaintenanceWindow = &commonApi.MaintenanceWindow{
						IsEnabled: wg.MaintenanceWindow.IsEnabled,
						StartTime: wg.MaintenanceWindow.StartTime.ValueStringPointer(),
					}

					if !wg.MaintenanceWindow.StartDay.IsUnknown() && !wg.MaintenanceWindow.StartDay.IsNull() {
						wgReq.MaintenanceWindow.StartDay = utils.ToPointer(float64(wg.MaintenanceWindow.StartDay.ValueInt64()))
					}
				}

				*clusterReqBody.Groups = append(*clusterReqBody.Groups, wgReq)
			} else {
				// keep existing witness group when updating a data group
				wgReq := pgdApi.WitnessGroup{
					ClusterType: utils.ToPointer("witness_group"),
					GroupId:     wg.GroupId.ValueStringPointer(),
				}
				*clusterReqBody.Groups = append(*clusterReqBody.Groups, wgReq)
			}
		}
	}

	_, err := p.client.Update(ctx, plan.ProjectId, *plan.ClusterId, clusterReqBody)
	if err != nil {
		if appendDiagFromBAErr(err, &resp.Diagnostics) {
			return
		}
		resp.Diagnostics.AddError("Error updating project", "Could not update project, unexpected error: "+err.Error())
		return
	}

	// sleep after update operation as API can incorrectly respond with healthy state when checking the phase
	// this is possibly a bug in the API
	time.Sleep(20 * time.Second)

	plan.ID = plan.ClusterId

	// retry func
	timeout, _ := plan.Timeouts.Update(ctx, 60*time.Minute)

	err = retry.RetryContext(
		ctx,
		timeout-time.Minute,
		p.retryFuncAs(ctx, &resp.Diagnostics, resp.State, &plan),
	)

	if err != nil {
		if appendDiagFromBAErr(err, &resp.Diagnostics) {
			return
		}
		resp.Diagnostics.AddError("Error retrying PGD cluster", "Could not update PGD cluster, unexpected error: "+err.Error())
		return
	}
	// end of retry func

	// uncomment below and comment retry func above to skip retry for testing purposes
	// clusterResp, err := p.client.Read(ctx, plan.ProjectId, *plan.ClusterId)
	// if err != nil {
	// 	resp.Diagnostics.AddError("Error reading PGD cluster", "Could not read PGD cluster, unexpected error: "+err.Error())
	// 	return
	// }

	// buildTFGroupsAs(ctx, &resp.Diagnostics, resp.State, *clusterResp, &plan.DataGroups, &plan.WitnessGroups)
	// if resp.Diagnostics.HasError() {
	// 	return
	// }

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (p pgdResource) isHealthy(ctx context.Context, dgs []terraform.DataGroup, wgs []terraform.WitnessGroup) (bool, error) {
	for _, v := range dgs {
		if *v.Phase.ValueStringPointer() != models.PHASE_HEALTHY {
			return false, nil
		}

		conditions := &[]terraform.Condition{}
		v.Conditions.ElementsAs(ctx, conditions, true)

		if len(*conditions) == 0 {
			return false, nil
		}

		for _, cond := range *conditions {
			if *cond.Type_.ValueStringPointer() != models.CONDITION_DEPLOYED && *cond.ConditionStatus.ValueStringPointer() != "True" {
				return false, nil
			}
		}
	}

	for _, v := range wgs {
		if *v.Phase.ValueStringPointer() != models.PHASE_HEALTHY {
			return false, nil
		}
	}

	return true, nil
}

func (p *pgdResource) retryFuncAs(ctx context.Context, diags *diag.Diagnostics, state tfsdk.State, pgdTfResource *PGD) retry.RetryFunc {
	return func() *retry.RetryError {
		pgdResp, err := p.client.Read(ctx, pgdTfResource.ProjectId, *pgdTfResource.ClusterId)
		if err != nil {
			return retry.NonRetryableError(fmt.Errorf("error describing instance: %s", err))
		}

		buildTFGroupsAs(ctx, diags, state, *pgdResp, pgdTfResource)
		if diags.HasError() {
			return retry.NonRetryableError(fmt.Errorf("unable to copy group, got error: %s", err))
		}

		isHealthy, err := p.isHealthy(ctx, pgdTfResource.DataGroups, pgdTfResource.WitnessGroups)
		if err != nil {
			return retry.NonRetryableError(fmt.Errorf("error getting health: %s", err))
		}

		if !isHealthy {
			return retry.RetryableError(errors.New("instance not yet ready"))
		}

		return nil
	}
}

func buildTFGroupsAs(ctx context.Context, diags *diag.Diagnostics, state tfsdk.State, clusterResp models.Cluster, asPgdTFResource *PGD) {
	originalTFDgs := asPgdTFResource.DataGroups
	asPgdTFResource.DataGroups = []terraform.DataGroup{}
	asPgdTFResource.WitnessGroups = []terraform.WitnessGroup{}

	for _, v := range *clusterResp.Groups {
		switch apiGroupResp := v.(type) {
		case map[string]interface{}:
			if apiGroupResp["clusterType"] == "data_group" {
				apiDGModel := pgdApi.DataGroup{}

				if err := utils.CopyObjectJson(apiGroupResp, &apiDGModel); err != nil {
					if err != nil {
						diags.AddError("unable to copy data group", err.Error())
						return
					}
				}

				dgTFType := new(types.Set)
				state.GetAttribute(ctx, path.Root("data_groups"), dgTFType)

				// cluster arch
				clusterArch := &terraform.ClusterArchitecture{
					ClusterArchitectureId:   apiDGModel.ClusterArchitecture.ClusterArchitectureId,
					ClusterArchitectureName: types.StringPointerValue(apiDGModel.ClusterArchitecture.ClusterArchitectureName),
					Nodes:                   apiDGModel.ClusterArchitecture.Nodes,
					WitnessNodes:            types.Int64Value(int64(*apiDGModel.ClusterArchitecture.WitnessNodes)),
				}

				// conditions
				conditionsPathSteps := tftypes.NewAttributePathWithSteps([]tftypes.AttributePathStep{
					tftypes.AttributeName("conditions"),
				})

				conditionsAttr, _ := dgTFType.ElementType(ctx).ApplyTerraform5AttributePathStep(conditionsPathSteps.NextStep())
				conditionsTFType, ok := conditionsAttr.(types.SetType)
				if !ok {
					diags.AddError("provider casting error", "cannot cast condition response to set type")
					return
				}

				conditionsElemTFType, ok := conditionsTFType.ElemType.(types.ObjectType)
				if !ok {
					diags.AddError("provider casting error", "cannot cast condition element response to object type")
					return
				}

				conditions := []attr.Value{}
				if apiDGModel.Conditions != nil && len(*apiDGModel.Conditions) != 0 {
					for _, v := range *apiDGModel.Conditions {
						ob, diag := types.ObjectValue(conditionsElemTFType.AttrTypes, map[string]attr.Value{
							"condition_status": types.StringValue(*v.ConditionStatus),
							"type":             types.StringValue(*v.Type_),
						})
						if diag.HasError() {
							diags.Append(diag...)
							return
						}
						conditions = append(conditions, ob)
					}
				}

				conditionsElemType := types.ObjectType{AttrTypes: conditionsElemTFType.AttrTypes}
				conditionsSet := types.SetNull(conditionsElemType)
				if len(conditions) > 0 {
					conditionsSet = types.SetValueMust(conditionsElemType, conditions)
				}

				// pgConfig. If tf resource pg config elem matches with api response pg config elem then add the elem to tf resource pg config
				newPgConfig := []models.KeyValue{}
				var tfPgConfig *[]models.KeyValue
				for _, pgdTFResourceDG := range originalTFDgs {
					if pgdTFResourceDG.Region.RegionId == apiDGModel.Region.RegionId {
						tfPgConfig = pgdTFResourceDG.PgConfig
						break
					}
				}

				if tfPgConfig != nil && apiDGModel.PgConfig != nil {
					for _, tfPgConf := range *tfPgConfig {
						for _, apiPgConf := range *apiDGModel.PgConfig {
							if tfPgConf.Name == apiPgConf.Name {
								newPgConfig = append(newPgConfig, models.KeyValue{Name: apiPgConf.Name, Value: apiPgConf.Value})
							}
						}
					}
				}

				// resizing pvc
				resizingPvc := []attr.Value{}
				if apiDGModel.ResizingPvc != nil && len(*apiDGModel.ResizingPvc) != 0 {
					for _, v := range *apiDGModel.ResizingPvc {
						v := v
						resizingPvc = append(resizingPvc, types.StringPointerValue(&v))
					}
				}

				// storage
				storage := &terraform.Storage{
					Size:               types.StringPointerValue(apiDGModel.Storage.Size),
					VolumePropertiesId: types.StringPointerValue(apiDGModel.Storage.VolumePropertiesId),
					VolumeTypeId:       types.StringPointerValue(apiDGModel.Storage.VolumeTypeId),
					Iops:               types.StringPointerValue(apiDGModel.Storage.Iops),
					Throughput:         types.StringPointerValue(apiDGModel.Storage.Throughput),
				}

				// service account ids
				serviceAccIds := []attr.Value{}
				if apiDGModel.ServiceAccountIds != nil && len(*apiDGModel.ServiceAccountIds) != 0 {
					for _, v := range *apiDGModel.ServiceAccountIds {
						v := v
						serviceAccIds = append(serviceAccIds, types.StringPointerValue(&v))
					}
				}

				// pe allowed principal ids account ids
				principalIds := []attr.Value{}
				if apiDGModel.PeAllowedPrincipalIds != nil && len(*apiDGModel.PeAllowedPrincipalIds) != 0 {
					for _, v := range *apiDGModel.PeAllowedPrincipalIds {
						v := v
						principalIds = append(principalIds, types.StringPointerValue(&v))
					}
				}

				// allowed ip ranges
				allwdIpRngsPathSteps := tftypes.NewAttributePathWithSteps([]tftypes.AttributePathStep{
					tftypes.AttributeName("allowed_ip_ranges"),
				})

				allwdIpRngsAttr, _ := dgTFType.ElementType(ctx).ApplyTerraform5AttributePathStep(allwdIpRngsPathSteps.NextStep())
				allwdIpRngsTFType, ok := allwdIpRngsAttr.(types.SetType)
				if !ok {
					diags.AddError("provider type assertion error", "cannot type assert allowed_ip_ranges response to set type")
					return
				}

				allwdIpRngsElemTFType, ok := allwdIpRngsTFType.ElemType.(types.ObjectType)
				if !ok {
					diags.AddError("provider type assertion  error", "cannot type assert allowed_ip_ranges element response to object type")
					return
				}
				allowedIpRanges := []attr.Value{}
				if apiDGModel.AllowedIpRanges != nil && len(*apiDGModel.AllowedIpRanges) > 0 {
					for _, v := range *apiDGModel.AllowedIpRanges {
						v := v
						ob, diag := types.ObjectValue(allwdIpRngsElemTFType.AttrTypes, map[string]attr.Value{
							"cidr_block":  types.StringValue(*&v.CidrBlock),
							"description": types.StringValue(*&v.Description),
						})
						if diag.HasError() {
							diags.Append(diag...)
							return
						}
						allowedIpRanges = append(allowedIpRanges, ob)
					}
				}

				allwdIpRngsElemType := types.ObjectType{AttrTypes: allwdIpRngsElemTFType.AttrTypes}
				allwdIpRngsSet := types.SetNull(allwdIpRngsElemType)
				if len(allowedIpRanges) > 0 {
					allwdIpRngsSet = types.SetValueMust(allwdIpRngsElemType, allowedIpRanges)
				}

				tfDGModel := terraform.DataGroup{
					GroupId:               types.StringPointerValue(apiDGModel.GroupId),
					AllowedIpRanges:       allwdIpRngsSet,
					BackupRetentionPeriod: apiDGModel.BackupRetentionPeriod,
					ClusterArchitecture:   clusterArch,
					ClusterName:           types.StringPointerValue(apiDGModel.ClusterName),
					ClusterType:           types.StringPointerValue(apiDGModel.ClusterType),
					Conditions:            conditionsSet,
					Connection:            types.StringPointerValue((*string)(apiDGModel.Connection)),
					CreatedAt:             types.StringPointerValue((*string)(apiDGModel.CreatedAt)),
					CspAuth:               apiDGModel.CspAuth,
					InstanceType:          apiDGModel.InstanceType,
					LogsUrl:               types.StringPointerValue(apiDGModel.LogsUrl),
					MetricsUrl:            types.StringPointerValue(apiDGModel.MetricsUrl),
					PgConfig:              &newPgConfig,
					PgType:                apiDGModel.PgType,
					PgVersion:             apiDGModel.PgVersion,
					Phase:                 types.StringPointerValue(apiDGModel.Phase),
					PrivateNetworking:     apiDGModel.PrivateNetworking,
					Provider:              apiDGModel.Provider,
					Region:                apiDGModel.Region,
					ResizingPvc:           types.SetValueMust(types.StringType, resizingPvc),
					Storage:               storage,
					MaintenanceWindow:     apiDGModel.MaintenanceWindow,
					ServiceAccountIds:     types.SetValueMust(types.StringType, serviceAccIds),
					PeAllowedPrincipalIds: types.SetValueMust(types.StringType, principalIds),
				}

				asPgdTFResource.DataGroups = append(asPgdTFResource.DataGroups, tfDGModel)
			}

			if apiGroupResp["clusterType"] == "witness_group" {
				apiWGModel := pgdApi.WitnessGroup{}

				if err := utils.CopyObjectJson(apiGroupResp, &apiWGModel); err != nil {
					if err != nil {
						diags.AddError("unable to copy witness group", err.Error())
						return
					}
				}

				wgTFType := new(types.Set)
				state.GetAttribute(ctx, path.Root("witness_groups"), wgTFType)

				// cluster arch
				clusterArchPathSteps := tftypes.NewAttributePathWithSteps([]tftypes.AttributePathStep{
					tftypes.AttributeName("cluster_architecture"),
				})

				clusterArchAttr, _ := wgTFType.ElementType(ctx).ApplyTerraform5AttributePathStep(clusterArchPathSteps.NextStep())
				clusterArchTFType, ok := clusterArchAttr.(types.ObjectType)
				if !ok {
					diags.AddError("cluster arch casting error", "cannot cast cluster architecture response to object type")
					return
				}

				clusterArch := types.ObjectNull(clusterArchTFType.AttrTypes)
				if apiWGModel.ClusterArchitecture != nil {
					ob, diag := types.ObjectValue(clusterArchTFType.AttrTypes, map[string]attr.Value{
						"cluster_architecture_id":   types.StringValue(apiWGModel.ClusterArchitecture.ClusterArchitectureId),
						"cluster_architecture_name": types.StringValue(*apiWGModel.ClusterArchitecture.ClusterArchitectureName),
						"nodes":                     types.Float64Value(apiWGModel.ClusterArchitecture.Nodes),
						"witness_nodes":             types.Float64PointerValue(apiWGModel.ClusterArchitecture.WitnessNodes),
					})
					if diag.HasError() {
						diags.Append(diag...)
						return
					}
					clusterArch = ob
				}

				// instance type
				instanceTypePathSteps := tftypes.NewAttributePathWithSteps([]tftypes.AttributePathStep{
					tftypes.AttributeName("instance_type"),
				})

				instanceTypeAttr, _ := wgTFType.ElementType(ctx).ApplyTerraform5AttributePathStep(instanceTypePathSteps.NextStep())
				instanceTypeTFType, ok := instanceTypeAttr.(types.ObjectType)

				if !ok {
					diags.AddError("provider error", "cannot cast instance type response to object type")
					return
				}
				instanceType := types.ObjectNull(instanceTypeTFType.AttrTypes)
				if apiWGModel.InstanceType != nil {

					ob, diag := types.ObjectValue(instanceTypeTFType.AttrTypes, map[string]attr.Value{
						"instance_type_id": types.StringValue(apiWGModel.InstanceType.InstanceTypeId),
					})
					if diag.HasError() {
						diags.Append(diag...)
						return
					}

					instanceType = ob
				}

				// provider
				providerPathSteps := tftypes.NewAttributePathWithSteps([]tftypes.AttributePathStep{
					tftypes.AttributeName("cloud_provider"),
				})

				providerAttr, _ := wgTFType.ElementType(ctx).ApplyTerraform5AttributePathStep(providerPathSteps.NextStep())
				providerTFType, ok := providerAttr.(types.ObjectType)

				if !ok {
					diags.AddError("provider error", "cannot cast cloud provider response object type")
					return
				}
				provider := types.ObjectNull(providerTFType.AttrTypes)
				if apiWGModel.Provider != nil {

					ob, diag := types.ObjectValue(providerTFType.AttrTypes, map[string]attr.Value{
						"cloud_provider_id": types.StringValue(*apiWGModel.Provider.CloudProviderId),
					})
					if diag.HasError() {
						diags.Append(diag...)
						return
					}
					provider = ob
				}

				// region
				region := &terraform.Region{
					RegionId: types.StringValue(apiWGModel.Region.RegionId),
				}

				// storage
				storagePathSteps := tftypes.NewAttributePathWithSteps([]tftypes.AttributePathStep{
					tftypes.AttributeName("storage"),
				})

				storageAttr, _ := wgTFType.ElementType(ctx).ApplyTerraform5AttributePathStep(storagePathSteps.NextStep())
				storageTFType, ok := storageAttr.(types.ObjectType)

				if !ok {
					diags.AddError("provider error", "cannot cast storage response object type")
					return
				}
				storage := types.ObjectNull(storageTFType.AttrTypes)
				if apiWGModel.Storage != nil {

					ob, diag := types.ObjectValue(storageTFType.AttrTypes, map[string]attr.Value{
						"iops":              types.StringValue(*apiWGModel.Storage.Iops),
						"size":              types.StringValue(*apiWGModel.Storage.Size),
						"throughput":        types.StringValue(*apiWGModel.Storage.Throughput),
						"volume_properties": types.StringValue(*apiWGModel.Storage.VolumePropertiesId),
						"volume_type":       types.StringValue(*apiWGModel.Storage.VolumeTypeId),
					})
					if diag.HasError() {
						diags.Append(diag...)
						return
					}
					storage = ob
				}

				// maintenance window
				var mw *commonTerraform.MaintenanceWindow
				if apiWGModel.MaintenanceWindow != nil {
					mw = &commonTerraform.MaintenanceWindow{
						IsEnabled: apiWGModel.MaintenanceWindow.IsEnabled,
						StartTime: types.StringPointerValue(apiWGModel.MaintenanceWindow.StartTime),
						StartDay:  types.Int64Value(int64(*apiWGModel.MaintenanceWindow.StartDay)),
					}
				}

				tfWGModel := terraform.WitnessGroup{
					GroupId:             types.StringPointerValue(apiWGModel.GroupId),
					ClusterArchitecture: clusterArch,
					ClusterType:         types.StringPointerValue(apiWGModel.ClusterType),
					InstanceType:        instanceType,
					Provider:            provider,
					Region:              region,
					Storage:             storage,
					Phase:               types.StringPointerValue(apiWGModel.Phase),
					MaintenanceWindow:   mw,
				}

				asPgdTFResource.WitnessGroups = append(asPgdTFResource.WitnessGroups, tfWGModel)
			}
		}
	}
}

// Delete deletes the resource and removes the Terraform state on success.
func (p pgdResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// Retrieve values from state
	var state PGD
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	if err := p.client.Delete(ctx, state.ProjectId, *state.ClusterId); err != nil {
		if appendDiagFromBAErr(err, &resp.Diagnostics) {
			return
		}
		resp.Diagnostics.AddError("Error deleting cluster", "Could not delete cluster, unexpected error: "+err.Error())
		return
	}
}

func (p pgdResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	idParts := strings.Split(req.ID, "/")
	if len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf("Expected import identifier with format: project_id/cluster_id. Got: %q", req.ID),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("project_id"), utils.ToPointer(idParts[0]))...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("cluster_id"), utils.ToPointer(idParts[1]))...)
}

func buildApiBah(ctx context.Context, client *api.PGDClient, diags *diag.Diagnostics, projectId string, dg terraform.DataGroup) (svAccIds, principalIds *[]string) {
	if strings.Contains(*dg.Provider.CloudProviderId, "bah") {
		if !dg.PeAllowedPrincipalIds.IsNull() && !dg.PeAllowedPrincipalIds.IsUnknown() && len(dg.PeAllowedPrincipalIds.Elements()) > 0 {
			elemDiag := dg.PeAllowedPrincipalIds.ElementsAs(ctx, &principalIds, false)
			if elemDiag.HasError() {
				diags.Append(elemDiag...)
				return nil, nil
			}
		} else {
			pids, err := client.GetPeAllowedPrincipalIds(ctx, projectId, *dg.Provider.CloudProviderId, dg.Region.RegionId)
			if err != nil {
				diags.AddError("pgd get pe allowed principal ids error", err.Error())
				return nil, nil
			}

			setOfPrincipalIds := []attr.Value{}
			for _, v := range pids.Data {
				setOfPrincipalIds = append(setOfPrincipalIds, basetypes.NewStringValue(v))
			}

			tfPrincipalIds := basetypes.NewSetValueMust(basetypes.StringType{}, setOfPrincipalIds)

			// if it doesn't have any existing service account ids then use config
			if !tfPrincipalIds.IsNull() && !tfPrincipalIds.IsUnknown() && len(tfPrincipalIds.Elements()) > 0 {
				elemDiag := dg.PeAllowedPrincipalIds.ElementsAs(ctx, &principalIds, false)
				if elemDiag.HasError() {
					diags.Append(elemDiag...)
					return nil, nil
				}
			}
		}

		if strings.Contains(*dg.Provider.CloudProviderId, "bah:gcp") {
			if !dg.ServiceAccountIds.IsNull() && !dg.ServiceAccountIds.IsUnknown() && len(dg.ServiceAccountIds.Elements()) > 0 {
				elemDiag := dg.ServiceAccountIds.ElementsAs(ctx, &svAccIds, false)
				if elemDiag.HasError() {
					diags.Append(elemDiag...)
					return nil, nil
				}
			} else {
				sids, err := client.GetServiceAccountIds(ctx, projectId, *dg.Provider.CloudProviderId, dg.Region.RegionId)
				if err != nil {
					diags.AddError("pgd get service account ids error", err.Error())
					return nil, nil
				}

				setOfSvAccIds := []attr.Value{}
				for _, v := range sids.Data {
					setOfSvAccIds = append(setOfSvAccIds, basetypes.NewStringValue(v))
				}

				tfSvAccIds := basetypes.NewSetValueMust(basetypes.StringType{}, setOfSvAccIds)

				// if it doesn't have any existing service account ids then use config
				if !tfSvAccIds.IsNull() && !tfSvAccIds.IsUnknown() && len(tfSvAccIds.Elements()) > 0 {
					elemDiag := dg.ServiceAccountIds.ElementsAs(ctx, &svAccIds, false)
					if elemDiag.HasError() {
						diags.Append(elemDiag...)
						return nil, nil
					}
				}
			}
		}
	}
	return
}

func buildApiAllowedIpRanges(tfAllowedIpRanges basetypes.SetValue) *[]models.AllowedIpRange {
	apiAllowedIpRanges := &[]models.AllowedIpRange{}

	for _, v := range tfAllowedIpRanges.Elements() {
		*apiAllowedIpRanges = append(*apiAllowedIpRanges, models.AllowedIpRange{
			CidrBlock:   v.(types.Object).Attributes()["cidr_block"].(types.String).ValueString(),
			Description: v.(types.Object).Attributes()["description"].(types.String).ValueString(),
		})
	}

	return apiAllowedIpRanges
}

func buildApiStorage(tfStorage terraform.Storage) *models.Storage {
	var iops *string
	var throughput *string

	// azurepremiumstorage only needs volume type, properties and size
	// other values will cause an unhelpful error on the API
	if *tfStorage.VolumeTypeId.ValueStringPointer() == "azurepremiumstorage" {
		iops = nil
		throughput = nil
	} else {
		iops = tfStorage.Iops.ValueStringPointer()
		if tfStorage.Iops.IsUnknown() {
			iops = nil
		}

		throughput = tfStorage.Throughput.ValueStringPointer()
		if tfStorage.Throughput.IsUnknown() {
			throughput = nil
		}
	}

	resultStorage := &models.Storage{
		Size:               tfStorage.Size.ValueStringPointer(),
		VolumePropertiesId: tfStorage.VolumePropertiesId.ValueStringPointer(),
		VolumeTypeId:       tfStorage.VolumeTypeId.ValueStringPointer(),
		Iops:               iops,
		Throughput:         throughput,
	}

	return resultStorage
}

func NewPgdResource() resource.Resource {
	return &pgdResource{}
}
