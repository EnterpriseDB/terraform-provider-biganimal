package provider

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/EnterpriseDB/terraform-provider-biganimal/pkg/api"
	"github.com/EnterpriseDB/terraform-provider-biganimal/pkg/models"
	pgdApi "github.com/EnterpriseDB/terraform-provider-biganimal/pkg/models/pgd/api"
	"github.com/EnterpriseDB/terraform-provider-biganimal/pkg/models/pgd/terraform"
	"github.com/EnterpriseDB/terraform-provider-biganimal/pkg/plan_modifier"
	"github.com/EnterpriseDB/terraform-provider-biganimal/pkg/utils"
	"github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/float64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
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

func (p pgdResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_pgd"
}

func (p pgdResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
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
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
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
							Required:    true,
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
								plan_modifier.CustomAllowedIps(),
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
								"witness_nodes": schema.Float64Attribute{
									Description: "Witness nodes count.",
									Optional:    true,
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
					},
				},
			},
			"witness_groups": schema.SetNestedAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.UseStateForUnknown(),
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
							Optional:    true,
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
							Description: "Cloud provider.",
							Computed:    true,
							PlanModifiers: []planmodifier.Object{
								objectplanmodifier.UseStateForUnknown(),
							},
							Attributes: map[string]schema.Attribute{
								"cloud_provider_id": schema.StringAttribute{
									Description: "Cloud provider id.",
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
					},
				},
			},
		},
	}
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

	witnessGroupParamsBody := pgdApi.WitnessGroupParamsBody{}
	for _, v := range config.DataGroups {
		v := v

		storage := buildApiStorage(*v.Storage)

		witnessGroupParamsBody.Groups = append(witnessGroupParamsBody.Groups, pgdApi.WitnessGroupParamsBodyData{
			InstanceType: v.InstanceType,
			Provider:     v.Provider,
			Region:       v.Region,
			Storage:      storage,
		})
		if v.AllowedIpRanges == nil {
			v.AllowedIpRanges = &[]models.AllowedIpRange{}
		}
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
			WitnessNodes:            v.ClusterArchitecture.WitnessNodes,
		}

		apiDGModel := pgdApi.DataGroup{
			AllowedIpRanges:       v.AllowedIpRanges,
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
		}
		*clusterReqBody.Groups = append(*clusterReqBody.Groups, apiDGModel)
	}

	if len(config.WitnessGroups) > 0 {
		calWitnessResp, err := p.client.CalculateWitnessGroupParams(ctx, config.ProjectId, witnessGroupParamsBody)
		if err != nil {
			if appendDiagFromBAErr(err, &resp.Diagnostics) {
				return
			}
			resp.Diagnostics.AddError("Error calculating witness group params", "Could not calculate witness group params, unexpected error: "+err.Error())
			return
		}

		for _, v := range config.WitnessGroups {
			wgReq := pgdApi.WitnessGroup{
				ClusterArchitecture: &pgdApi.ClusterArchitecture{
					ClusterArchitectureId: "pgd",
					Nodes:                 config.DataGroups[0].ClusterArchitecture.Nodes,
				},
				ClusterType: utils.ToPointer("witness_group"),
				Provider: &pgdApi.CloudProvider{
					CloudProviderId: config.DataGroups[0].Provider.CloudProviderId,
				},
				InstanceType: calWitnessResp.InstanceType,
				Storage:      calWitnessResp.Storage,
				Region: &pgdApi.Region{
					RegionId: v.Region.RegionId.ValueString(),
				},
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

	timeout, _ := config.Timeouts.Create(ctx, 60*time.Minute)

	err = retry.RetryContext(
		ctx,
		timeout-time.Minute,
		p.retryFuncAs(ctx, &resp.Diagnostics, &config),
	)

	if err != nil {
		if appendDiagFromBAErr(err, &resp.Diagnostics) {
			return
		}
		resp.Diagnostics.AddError("Error retrying PGD cluster", "Could not create PGD cluster, unexpected error: "+err.Error())
		return
	}

	// clusterResp, err := p.client.Read(ctx, config.ProjectId, clusterId)
	// if err != nil {
	// 	resp.Diagnostics.AddError("Error reading PGD cluster", "Could not read PGD cluster, unexpected error: "+err.Error())
	// 	return
	// }

	// config.DataGroups = []terraform.DataGroup{}
	// config.WitnessGroups = []terraform.WitnessGroup{}

	// if err = buildTFGroupsAs(*clusterResp, &config.DataGroups, &config.WitnessGroups); err != nil {
	// 	resp.Diagnostics.AddError("Resource create error", fmt.Sprintf("Unable to copy group, got error: %s", err))
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

	buildTFGroupsAs(ctx, &resp.Diagnostics, *clusterResp, &state.DataGroups, &state.WitnessGroups)
	if resp.Diagnostics.HasError() {
		return
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

	witnessGroupParamsBody := pgdApi.WitnessGroupParamsBody{}
	for _, v := range plan.DataGroups {
		storage := buildApiStorage(*v.Storage)

		witnessGroupParamsBody.Groups = append(witnessGroupParamsBody.Groups, pgdApi.WitnessGroupParamsBodyData{
			InstanceType: v.InstanceType,
			Provider:     v.Provider,
			Region:       v.Region,
			Storage:      storage,
		})

		groupId := v.GroupId.ValueStringPointer()
		if v.GroupId.IsUnknown() {
			groupId = nil
		}

		// only allow fields which are able to be modifed in the request for updating
		reqDg := pgdApi.DataGroup{
			GroupId:               groupId,
			ClusterType:           utils.ToPointer("data_group"),
			AllowedIpRanges:       v.AllowedIpRanges,
			BackupRetentionPeriod: v.BackupRetentionPeriod,
			CspAuth:               v.CspAuth,
			InstanceType:          v.InstanceType,
			PgConfig:              v.PgConfig,
			PrivateNetworking:     v.PrivateNetworking,
			Storage:               storage,
			MaintenanceWindow:     v.MaintenanceWindow,
		}

		// signals that it doesn't have an existing group id so this is a new group to add needs extra fields
		if reqDg.GroupId == nil {
			reqDg.Provider = v.Provider
			reqDg.Region = v.Region
			reqDg.ClusterArchitecture = &pgdApi.ClusterArchitecture{
				ClusterArchitectureId: v.ClusterArchitecture.ClusterArchitectureId,
				Nodes:                 v.ClusterArchitecture.Nodes,
			}
			reqDg.PgType = v.PgType
			reqDg.PgVersion = v.PgVersion
		}

		*clusterReqBody.Groups = append(*clusterReqBody.Groups, reqDg)
	}

	if len(plan.WitnessGroups) > 0 {
		calWitnessResp, err := p.client.CalculateWitnessGroupParams(ctx, plan.ProjectId, witnessGroupParamsBody)
		if err != nil {
			if appendDiagFromBAErr(err, &resp.Diagnostics) {
				return
			}
			resp.Diagnostics.AddError("Error calculating witness group params", "Could not calculate witness group params, unexpected error: "+err.Error())
			return
		}

		for _, v := range plan.WitnessGroups {
			// cannot change anything on witness group, this only allows adding a new witness group
			if v.GroupId.IsNull() || v.GroupId.IsUnknown() {
				wgReq := pgdApi.WitnessGroup{
					ClusterArchitecture: &pgdApi.ClusterArchitecture{
						ClusterArchitectureId: "pgd",
						Nodes:                 plan.DataGroups[0].ClusterArchitecture.Nodes,
					},
					ClusterType: utils.ToPointer("witness_group"),
					Provider: &pgdApi.CloudProvider{
						CloudProviderId: plan.DataGroups[0].Provider.CloudProviderId,
					},
					InstanceType: calWitnessResp.InstanceType,
					Storage:      calWitnessResp.Storage,
					Region: &pgdApi.Region{
						RegionId: v.Region.RegionId.ValueString(),
					},
				}

				*clusterReqBody.Groups = append(*clusterReqBody.Groups, wgReq)
			}
		}
	}

	res, _ := json.MarshalIndent(clusterReqBody, "", "  ")
	fmt.Println(string(res))
	_, err := p.client.Update(ctx, plan.ProjectId, *plan.ClusterId, clusterReqBody)
	if err != nil {
		if appendDiagFromBAErr(err, &resp.Diagnostics) {
			return
		}
		resp.Diagnostics.AddError("Error updating project", "Could not update project, unexpected error: "+err.Error())
		return
	}

	plan.ID = plan.ClusterId

	timeout, _ := plan.Timeouts.Update(ctx, 60*time.Minute)

	err = retry.RetryContext(
		ctx,
		timeout-time.Minute,
		p.retryFuncAs(ctx, &resp.Diagnostics, &plan),
	)

	if err != nil {
		if appendDiagFromBAErr(err, &resp.Diagnostics) {
			return
		}
		resp.Diagnostics.AddError("Error retrying PGD cluster", "Could not update PGD cluster, unexpected error: "+err.Error())
		return
	}

	// clusterResp, err := p.client.Read(ctx, plan.ProjectId, *plan.ClusterId)
	// if err != nil {
	// 	resp.Diagnostics.AddError("Error reading PGD cluster", "Could not read PGD cluster, unexpected error: "+err.Error())
	// 	return
	// }

	// plan.DataGroups = []pgd.DataGroup{}
	// plan.WitnessGroups = []pgd.WitnessGroup{}

	// if err = buildGroupsToTypeAs(*clusterResp, &plan.DataGroups, &plan.WitnessGroups); err != nil {
	// 	resp.Diagnostics.AddError("Resource create error", fmt.Sprintf("Unable to copy group, got error: %s", err))
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

		if conditions == nil {
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

func (p *pgdResource) retryFuncAs(ctx context.Context, diags *diag.Diagnostics, state *PGD) retry.RetryFunc {
	return func() *retry.RetryError {
		pgdResp, err := p.client.Read(ctx, state.ProjectId, *state.ClusterId)
		if err != nil {
			return retry.NonRetryableError(fmt.Errorf("error describing instance: %s", err))
		}

		buildTFGroupsAs(ctx, diags, *pgdResp, &state.DataGroups, &state.WitnessGroups)
		if diags.HasError() {
			return retry.NonRetryableError(fmt.Errorf("unable to copy group, got error: %s", err))
		}

		isHealthy, err := p.isHealthy(ctx, state.DataGroups, state.WitnessGroups)
		if err != nil {
			return retry.NonRetryableError(fmt.Errorf("error getting health: %s", err))
		}

		if !isHealthy {
			return retry.RetryableError(errors.New("instance not yet ready"))
		}

		return nil
	}
}

func buildTFGroupsAs(ctx context.Context, diags *diag.Diagnostics, clusterResp models.Cluster, dgs *[]terraform.DataGroup, wgs *[]terraform.WitnessGroup) {
	// sourceDgs := dgs
	// t := (*sourceDgs)[0].Conditions.ElementType()

	*dgs = []terraform.DataGroup{}
	*wgs = []terraform.WitnessGroup{}

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

				clusterArch := &terraform.ClusterArchitecture{
					ClusterArchitectureId:   apiDGModel.ClusterArchitecture.ClusterArchitectureId,
					ClusterArchitectureName: types.StringPointerValue(apiDGModel.ClusterArchitecture.ClusterArchitectureName),
					Nodes:                   apiDGModel.ClusterArchitecture.Nodes,
					WitnessNodes:            apiDGModel.ClusterArchitecture.WitnessNodes,
				}

				conditions := []attr.Value{}
				obType := map[string]attr.Type{
					"condition_status": types.StringType,
					"type":             types.StringType,
				}

				if apiDGModel.Conditions != nil && len(*apiDGModel.Conditions) != 0 {
					for _, v := range *apiDGModel.Conditions {
						ob, diag := types.ObjectValue(obType, map[string]attr.Value{
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

				conditionsSet := types.Set{}
				if len(conditions) > 0 {
					conditionsSet = types.SetValueMust(conditions[0].Type(ctx), conditions)
				}

				resizingPvc := []attr.Value{}
				if apiDGModel.ResizingPvc != nil && len(*apiDGModel.ResizingPvc) != 0 {
					for _, v := range *apiDGModel.ResizingPvc {
						v := v
						resizingPvc = append(resizingPvc, types.StringPointerValue(&v))
					}
				}

				storage := &terraform.Storage{
					Size:               types.StringPointerValue(apiDGModel.Storage.Size),
					VolumePropertiesId: types.StringPointerValue(apiDGModel.Storage.VolumePropertiesId),
					VolumeTypeId:       types.StringPointerValue(apiDGModel.Storage.VolumeTypeId),
					Iops:               types.StringPointerValue(apiDGModel.Storage.Iops),
					Throughput:         types.StringPointerValue(apiDGModel.Storage.Throughput),
				}

				tfDGModel := terraform.DataGroup{
					GroupId:               types.StringPointerValue(apiDGModel.GroupId),
					AllowedIpRanges:       apiDGModel.AllowedIpRanges,
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
					PgConfig:              apiDGModel.PgConfig,
					PgType:                apiDGModel.PgType,
					PgVersion:             apiDGModel.PgVersion,
					Phase:                 types.StringPointerValue(apiDGModel.Phase),
					PrivateNetworking:     apiDGModel.PrivateNetworking,
					Provider:              apiDGModel.Provider,
					Region:                apiDGModel.Region,
					ResizingPvc:           types.SetValueMust(types.StringType, resizingPvc),
					Storage:               storage,
					MaintenanceWindow:     apiDGModel.MaintenanceWindow,
				}

				*dgs = append(*dgs, tfDGModel)
			}

			if apiGroupResp["clusterType"] == "witness_group" {
				apiWGModel := pgdApi.WitnessGroup{}

				if err := utils.CopyObjectJson(apiGroupResp, &apiWGModel); err != nil {
					if err != nil {
						diags.AddError("unable to copy witness group", err.Error())
						return
					}
				}

				clusterArch := types.Object{}
				if apiWGModel.ClusterArchitecture != nil {
					obType := map[string]attr.Type{
						"cluster_architecture_id":   types.StringType,
						"cluster_architecture_name": types.StringType,
						"nodes":                     types.StringType,
						"witness_nodes":             types.StringType,
					}

					ob, diag := types.ObjectValue(obType, map[string]attr.Value{
						"cluster_architecture_id":   types.StringValue(apiWGModel.ClusterArchitecture.ClusterArchitectureId),
						"cluster_architecture_name": types.StringValue(*apiWGModel.ClusterArchitecture.ClusterArchitectureName),
						"nodes":                     types.Float64Value(apiWGModel.ClusterArchitecture.Nodes),
						"witness_nodes":             types.Float64Value(*apiWGModel.ClusterArchitecture.WitnessNodes),
					})
					if diag.HasError() {
						diags.Append(diag...)
						return
					}
					clusterArch = ob
				}

				instanceType := types.Object{}
				if apiWGModel.ClusterArchitecture != nil {
					obType := map[string]attr.Type{
						"instance_type_id": types.StringType,
					}

					ob, diag := types.ObjectValue(obType, map[string]attr.Value{
						"instance_type_id": types.StringValue(apiWGModel.InstanceType.InstanceTypeId),
					})
					if diag.HasError() {
						diags.Append(diag...)
						return
					}
					instanceType = ob
				}

				provider := types.Object{}
				if apiWGModel.ClusterArchitecture != nil {
					obType := map[string]attr.Type{
						"cloud_provider_id": types.StringType,
					}

					ob, diag := types.ObjectValue(obType, map[string]attr.Value{
						"cloud_provider_id": types.StringValue(*apiWGModel.Provider.CloudProviderId),
					})
					if diag.HasError() {
						diags.Append(diag...)
						return
					}
					provider = ob
				}

				region := &terraform.Region{
					RegionId: types.StringValue(apiWGModel.Region.RegionId),
				}

				storage := types.Object{}
				if apiWGModel.ClusterArchitecture != nil {
					obType := map[string]attr.Type{
						"iops":              types.StringType,
						"size":              types.StringType,
						"throughput":        types.StringType,
						"volume_properties": types.StringType,
						"volume_type":       types.StringType,
					}

					ob, diag := types.ObjectValue(obType, map[string]attr.Value{
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

				tfWGModel := terraform.WitnessGroup{
					GroupId:             types.StringPointerValue(apiWGModel.GroupId),
					ClusterArchitecture: clusterArch,
					ClusterType:         types.StringPointerValue(apiWGModel.ClusterType),
					InstanceType:        instanceType,
					Provider:            provider,
					Region:              region,
					Storage:             storage,
					Phase:               types.StringPointerValue(apiWGModel.Phase),
				}

				*wgs = append(*wgs, tfWGModel)
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
