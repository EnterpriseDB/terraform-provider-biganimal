package provider

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/EnterpriseDB/terraform-provider-biganimal/pkg/api"
	"github.com/EnterpriseDB/terraform-provider-biganimal/pkg/models"
	"github.com/EnterpriseDB/terraform-provider-biganimal/pkg/models/pgd"
	"github.com/EnterpriseDB/terraform-provider-biganimal/pkg/plan_modifier"
	"github.com/EnterpriseDB/terraform-provider-biganimal/pkg/utils"
	"github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
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
						"first_recoverability_point_at": schema.StringAttribute{
							Description: "Earliest backup recover time.",
							Optional:    true,
							Computed:    true,
							PlanModifiers: []planmodifier.String{
								plan_modifier.CustomStringForUnknown(),
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
							Description: "Cluster connection URI.",
							Computed:    true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"phase": schema.StringAttribute{
							Description: "Current phase of the cluster group.",
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
									Description: "Name.",
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
							Description: "Phase.",
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
							Optional:    true,
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
	ID            *string            `tfsdk:"id"`
	ProjectId     string             `tfsdk:"project_id"`
	ClusterId     *string            `tfsdk:"cluster_id"`
	ClusterName   *string            `tfsdk:"cluster_name"`
	MostRecent    *bool              `tfsdk:"most_recent"`
	Password      *string            `tfsdk:"password"`
	Timeouts      timeouts.Value     `tfsdk:"timeouts"`
	DataGroups    []pgd.DataGroup    `tfsdk:"data_groups"`
	WitnessGroups []pgd.WitnessGroup `tfsdk:"witness_groups"`
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

	witnessGroupParamsBody := pgd.WitnessGroupParamsBody{}
	for _, v := range config.DataGroups {
		v.ClusterType = utils.ToPointer("data_group")

		buildStorageAs(v.Storage)

		witnessGroupParamsBody.Groups = append(witnessGroupParamsBody.Groups, pgd.WitnessGroupParamsBodyData{
			InstanceType: v.InstanceType,
			Provider:     v.Provider,
			Region:       v.Region,
			Storage:      v.Storage,
		})
		if v.AllowedIpRanges == nil {
			v.AllowedIpRanges = &[]models.AllowedIpRange{}
		}
		if v.PgConfig == nil {
			v.PgConfig = &[]models.KeyValue{}
		}
		*clusterReqBody.Groups = append(*clusterReqBody.Groups, v)
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
			v.ClusterArchitecture = &pgd.ClusterArchitecture{
				ClusterArchitectureId: "pgd",
				Nodes:                 config.DataGroups[0].ClusterArchitecture.Nodes,
			}
			v.ClusterType = utils.ToPointer("witness_group")
			v.Provider = &pgd.CloudProvider{
				CloudProviderId: config.DataGroups[0].Provider.CloudProviderId,
			}
			v.InstanceType = calWitnessResp.InstanceType
			v.Storage = calWitnessResp.Storage
			*clusterReqBody.Groups = append(*clusterReqBody.Groups, v)
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
		p.retryFuncAs(ctx, &config),
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

	// config.DataGroups = []pgd.DataGroup{}
	// config.WitnessGroups = []pgd.WitnessGroup{}

	// if err = buildGroupsToTypeAs(*clusterResp, &config.DataGroups, &config.WitnessGroups); err != nil {
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

	if err = buildGroupsToTypeAs(*clusterResp, &state.DataGroups, &state.WitnessGroups); err != nil {
		resp.Diagnostics.AddError("Resource read error", fmt.Sprintf("Unable to copy group, got error: %s", err))
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

	witnessGroupParamsBody := pgd.WitnessGroupParamsBody{}
	for _, v := range plan.DataGroups {
		v.ClusterType = utils.ToPointer("data_group")

		buildStorageAs(v.Storage)

		witnessGroupParamsBody.Groups = append(witnessGroupParamsBody.Groups, pgd.WitnessGroupParamsBodyData{
			InstanceType: v.InstanceType,
			Provider:     v.Provider,
			Region:       v.Region,
			Storage:      v.Storage,
		})

		// only allow fields which are able to be modifed in the request
		reqDg := pgd.DataGroup{
			GroupId:               v.GroupId,
			ClusterType:           v.ClusterType,
			AllowedIpRanges:       v.AllowedIpRanges,
			BackupRetentionPeriod: v.BackupRetentionPeriod,
			CspAuth:               v.CspAuth,
			InstanceType:          v.InstanceType,
			PgConfig:              v.PgConfig,
			PrivateNetworking:     v.PrivateNetworking,
			Storage:               v.Storage,
			MaintenanceWindow:     v.MaintenanceWindow,
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
			v.ClusterArchitecture = &pgd.ClusterArchitecture{
				ClusterArchitectureId: "pgd",
				Nodes:                 plan.DataGroups[0].ClusterArchitecture.Nodes,
			}
			v.ClusterType = utils.ToPointer("witness_group")
			v.Provider = &pgd.CloudProvider{
				CloudProviderId: plan.DataGroups[0].Provider.CloudProviderId,
			}
			v.InstanceType = calWitnessResp.InstanceType
			v.Storage = calWitnessResp.Storage
			v.Phase = nil
			v.Region = nil
			v.Provider = nil
			*clusterReqBody.Groups = append(*clusterReqBody.Groups, v)
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

	plan.ID = plan.ClusterId

	timeout, _ := plan.Timeouts.Update(ctx, 60*time.Minute)

	err = retry.RetryContext(
		ctx,
		timeout-time.Minute,
		p.retryFuncAs(ctx, &plan),
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

func (p pgdResource) isHealthy(ctx context.Context, dgs []pgd.DataGroup, wgs []pgd.WitnessGroup) (bool, error) {
	healthy := true

	for _, v := range dgs {
		if *v.Phase != models.PHASE_HEALTHY {
			healthy = false
		}

		for _, cond := range *v.Conditions {
			if *cond.Type_ != models.CONDITION_DEPLOYED && *cond.ConditionStatus != "True" {
				healthy = false
			}
		}
	}

	for _, v := range wgs {
		if *v.Phase != models.PHASE_HEALTHY {
			healthy = false
		}
	}

	return healthy, nil
}

func (p *pgdResource) retryFuncAs(ctx context.Context, state *PGD) retry.RetryFunc {
	return func() *retry.RetryError {
		pgdResp, err := p.client.Read(ctx, state.ProjectId, *state.ClusterId)
		if err != nil {
			return retry.NonRetryableError(fmt.Errorf("error describing instance: %s", err))
		}

		if err := buildGroupsToTypeAs(*pgdResp, &state.DataGroups, &state.WitnessGroups); err != nil {
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

func buildGroupsToTypeAs(clusterResp models.Cluster, dgs *[]pgd.DataGroup, wgs *[]pgd.WitnessGroup) error {
	*dgs = []pgd.DataGroup{}
	*wgs = []pgd.WitnessGroup{}

	for _, v := range *clusterResp.Groups {
		switch apiGroupResp := v.(type) {
		case map[string]interface{}:
			if apiGroupResp["clusterType"] == "data_group" {
				model := pgd.DataGroup{}

				if err := utils.CopyObjectJson(apiGroupResp, &model); err != nil {
					if err != nil {
						return fmt.Errorf("unable to copy data group, got error: %s", err)
					}
				}

				*dgs = append(*dgs, model)
			}

			if apiGroupResp["clusterType"] == "witness_group" {
				model := pgd.WitnessGroup{}

				if err := utils.CopyObjectJson(apiGroupResp, &model); err != nil {
					if err != nil {
						return fmt.Errorf("unable to copy witness group, got error: %s", err)
					}
				}

				*wgs = append(*wgs, model)
			}
		}
	}

	return nil
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

func buildStorageAs(storage *models.Storage) {
	// azurepremiumstorage only needs volume type, properties and size
	// other values will cause an unhelpful error on the API
	if *storage.VolumeTypeId == "azurepremiumstorage" {
		storage.Iops = nil
		storage.Throughput = nil
	}
}

func NewPgdResource() resource.Resource {
	return &pgdResource{}
}
