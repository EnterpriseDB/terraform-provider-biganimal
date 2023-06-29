package provider

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/EnterpriseDB/terraform-provider-biganimal/pkg/api"
	"github.com/EnterpriseDB/terraform-provider-biganimal/pkg/models"
	"github.com/EnterpriseDB/terraform-provider-biganimal/pkg/models/pgd"
	"github.com/EnterpriseDB/terraform-provider-biganimal/pkg/utils"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
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
			"data_groups": schema.SetNestedBlock{
				Description: "Cluster data groups.",
				NestedObject: schema.NestedBlockObject{
					Blocks: map[string]schema.Block{
						"allowed_ip_ranges": schema.SetNestedBlock{
							Description: "Allowed IP ranges.",
							NestedObject: schema.NestedBlockObject{
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
						},
						"pg_config": schema.SetNestedBlock{
							Description: "Database configuration parameters.",

							NestedObject: schema.NestedBlockObject{
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										Description: "GUC name.",
										Optional:    true,
									},
									"value": schema.StringAttribute{
										Description: "GUC value.",
										Optional:    true,
									},
								},
							},
						},
						"cluster_architecture": schema.SingleNestedBlock{
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
						"storage": schema.SingleNestedBlock{
							Description: "Storage.",
							Attributes: map[string]schema.Attribute{
								"iops": schema.StringAttribute{
									Description: "IOPS for the selected volume.",
									Optional:    true,
									Computed:    true,
								},
								"size": schema.StringAttribute{
									Description: "Size of the volume.",
									Optional:    true,
									Computed:    true,
								},
								"throughput": schema.StringAttribute{
									Description: "Throughput.",
									Optional:    true,
									Computed:    true,
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
						"pg_type": schema.SingleNestedBlock{
							Description: "Postgres type.",
							Attributes: map[string]schema.Attribute{
								"pg_type_id": schema.StringAttribute{
									Description: "Data group pg type id.",
									Optional:    true,
								},
							},
						},
						"pg_version": schema.SingleNestedBlock{
							Description: "Postgres version.",
							Attributes: map[string]schema.Attribute{
								"pg_version_id": schema.StringAttribute{
									Description: "Data group pg version id.",
									Optional:    true,
								},
							},
						},
						"cloud_provider": schema.SingleNestedBlock{
							Description: "Cloud provider.",
							Attributes: map[string]schema.Attribute{
								"cloud_provider_id": schema.StringAttribute{
									Description: "Data group cloud provider id.",
									Optional:    true,
								},
							},
						},
						"region": schema.SingleNestedBlock{
							Description: "Region.",
							Attributes: map[string]schema.Attribute{
								"region_id": schema.StringAttribute{
									Description: "Data group region id.",
									Optional:    true,
								},
							},
						},
						"instance_type": schema.SingleNestedBlock{
							Description: "Instance type.",
							Attributes: map[string]schema.Attribute{
								"instance_type_id": schema.StringAttribute{
									Description: "Data group instance type id.",
									Optional:    true,
								},
							},
						},
						// "maintenance_window": schema.SingleNestedBlock{
						// 	Description: "Custom maintenance window.",
						// 	Attributes: map[string]schema.Attribute{
						// 		"is_enabled": schema.BoolAttribute{
						// 			Description: "Is maintenance window enabled.",
						// 			Optional:    true,
						// 			Computed:    true,
						// 		},
						// 		"start_day": schema.StringAttribute{
						// 			Description: "Start day.",
						// 			Optional:    true,
						// 			Computed:    true,
						// 		},
						// 		"start_time": schema.StringAttribute{
						// 			Description: "Start time.",
						// 			Optional:    true,
						// 			Computed:    true,
						// 		},
						// 	},
						// },
					},
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
							Optional:    true,
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
						},
						"created_at": schema.StringAttribute{
							Description: "Cluster creation time.",
							Computed:    true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"deleted_at": schema.StringAttribute{
							Description: "Cluster deletion time.",
							Optional:    true,
						},
						"expired_at": schema.StringAttribute{
							Description: "Cluster expiry time.",
							Optional:    true,
						},
						"first_recoverability_point_at": schema.StringAttribute{
							Description: "Earliest backup recover time.",
							Optional:    true,
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
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"private_networking": schema.BoolAttribute{
							Description: "Is private networking enabled.",
							Required:    true,
						},
						"csp_auth": schema.BoolAttribute{
							Description: "Is authentication handled by the cloud service provider.",
							Optional:    true,
						},
						"resizing_pvc": schema.SetAttribute{
							Description: "Resizing PVC.",
							Computed:    true,
							ElementType: types.StringType,
						},
					},
				},
			},
			"witness_groups": schema.SetNestedBlock{
				Description: "Cluster witness groups.",
				NestedObject: schema.NestedBlockObject{
					Blocks: map[string]schema.Block{
						"cluster_architecture": schema.SingleNestedBlock{
							Attributes: map[string]schema.Attribute{
								"cluster_architecture_id": schema.StringAttribute{
									Description: "Cluster architecture ID.",
									Optional:    true,
								},
								"cluster_architecture_name": schema.StringAttribute{
									Description: "Name.",
									Optional:    true,
								},
								"nodes": schema.Float64Attribute{
									Description: "Nodes.",
									Optional:    true,
								},
								"witness_nodes": schema.Float64Attribute{
									Description: "Witness nodes count.",
									Optional:    true,
								},
							},
						},
						"region": schema.SingleNestedBlock{
							Description: "Region.",
							Attributes: map[string]schema.Attribute{
								"region_id": schema.StringAttribute{
									Description: "Witness group region id.",
									Required:    true,
								},
							},
						},
						"cloud_provider": schema.SingleNestedBlock{
							Description: "Cloud provider.",
							Attributes: map[string]schema.Attribute{
								"cloud_provider_id": schema.StringAttribute{
									Description: "Witness group cloud provider id.",
									Computed:    true,
									PlanModifiers: []planmodifier.String{
										stringplanmodifier.UseStateForUnknown(),
									},
								},
							},
						},
						"instance_type": schema.SingleNestedBlock{
							Description: "Instance type.",
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
						"storage": schema.SingleNestedBlock{
							Description: "Storage.",
							Attributes: map[string]schema.Attribute{
								"iops": schema.StringAttribute{
									Description: "IOPS for the selected volume.",
									Optional:    true,
								},
								"size": schema.StringAttribute{
									Description: "Size of the volume.",
									Optional:    true,
								},
								"throughput": schema.StringAttribute{
									Description: "Throughput.",
									Optional:    true,
								},
								"volume_properties": schema.StringAttribute{
									Description: "Volume properties.",
									Optional:    true,
								},
								"volume_type": schema.StringAttribute{
									Description: "Volume type.",
									Optional:    true,
								},
							},
						},
					},
					Attributes: map[string]schema.Attribute{
						"cluster_type": schema.StringAttribute{
							Description: "Type of the Specified Cluster",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
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
	ClusterName   string             `tfsdk:"cluster_name"`
	MostRecent    *bool              `tfsdk:"most_recent"`
	Password      string             `tfsdk:"password"`
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
		ClusterName: &config.ClusterName,
		ClusterType: utils.ToPointer("cluster"),
		Password:    &config.Password,
	}

	clusterReqBody.Groups = &[]any{}

	for _, v := range config.DataGroups {
		v.ClusterType = utils.ToPointer("data_group")
		*clusterReqBody.Groups = append(*clusterReqBody.Groups, v)
	}
	for _, v := range config.WitnessGroups {
		v := v
		v.ClusterArchitecture = &pgd.ClusterArchitecture{
			ClusterArchitectureId: "pgd",
			Nodes:                 2,
		}
		v.ClusterType = utils.ToPointer("witness_group")
		v.Provider = &pgd.CloudProvider{
			CloudProviderId: utils.ToPointer("azure"),
		}
		v.InstanceType = &pgd.InstanceType{
			InstanceTypeId: "azure:Standard_D2s_v3",
		}
		v.Storage = &models.Storage{
			VolumeTypeId:       "azurepremiumstorage",
			VolumePropertiesId: "P1",
			Size:               utils.ToPointer("4 Gi"),
		}
		*clusterReqBody.Groups = append(*clusterReqBody.Groups, v)
	}

	clusterId, err := p.client.Create(ctx, config.ProjectId, clusterReqBody)
	if err != nil {
		resp.Diagnostics.AddError("Error creating PGD cluster", "Could not create PGD cluster, unexpected error: "+err.Error())
		return
	}

	clusterResp, err := p.client.Read(ctx, config.ProjectId, clusterId)
	if err != nil {
		resp.Diagnostics.AddError("Error reading PGD cluster", "Could not read PGD cluster, unexpected error: "+err.Error())
		return
	}

	config.ID = clusterResp.ClusterId
	config.ClusterId = clusterResp.ClusterId
	config.DataGroups = []pgd.DataGroup{}
	config.WitnessGroups = []pgd.WitnessGroup{}

	if err = buildStateGroups(&resp.Diagnostics, *clusterResp, &config.DataGroups, &config.WitnessGroups); err != nil {
		resp.Diagnostics.AddError("Resource create error", fmt.Sprintf("Unable to copy group, got error: %s", err))
		return
	}

	bb, err := json.MarshalIndent(config.WitnessGroups, "", "  ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Print(string(bb))

	//  config.WitnessGroups[0].ClusterArchitecture = &pgd.ClusterArchitecture{
	// 	ClusterArchitectureId:   "pgd",
	// 	Nodes:                   2,
	// 	ClusterArchitectureName: utils.ToPointer("pgd"),
	// 	WitnessNodes:            utils.ToPointer(float64(2)),
	// }

	www := config.WitnessGroups[0]

	config.WitnessGroups = []pgd.WitnessGroup{
		{
			Region: &pgd.Region{
				RegionId: "canadacentral",
			},
			ClusterType: www.ClusterType,
			Provider: &pgd.CloudProvider{
				CloudProviderId: utils.ToPointer("azure"),
			},
		},
	}

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

	clusterResp, err := p.client.Read(ctx, state.ProjectId, *state.ID)
	if err != nil {
		resp.Diagnostics.AddError("Error reading PGD cluster", "Could not read PGD cluster, unexpected error: "+err.Error())
		return
	}

	state.ID = clusterResp.ClusterId
	state.ClusterId = clusterResp.ClusterId

	if err = buildStateGroups(&resp.Diagnostics, *clusterResp, &state.DataGroups, &state.WitnessGroups); err != nil {
		resp.Diagnostics.AddError("Resource read error", fmt.Sprintf("Unable to copy group, got error: %s", err))
		return
	}

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (p pgdResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
}

func (p pgdResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
}

type diagnostics interface {
	AddError(summary string, detail string)
}

func buildStateGroups(diag diagnostics, clusterResp models.Cluster, dgs *[]pgd.DataGroup, wgs *[]pgd.WitnessGroup) error {
	*dgs = []pgd.DataGroup{}
	*wgs = []pgd.WitnessGroup{}

	for _, v := range *clusterResp.Groups {
		switch apiGroupResp := v.(type) {
		case map[string]interface{}:
			if apiGroupResp["clusterType"] == "data_group" {
				model := pgd.DataGroup{}

				if err := utils.CopyObjectJson(apiGroupResp, &model); err != nil {
					if err != nil {
						diag.AddError("Read Error", fmt.Sprintf("Unable to copy data group, got error: %s", err))
						return err
					}
				}

				*dgs = append(*dgs, model)
			}

			if apiGroupResp["clusterType"] == "witness_group" {
				model := pgd.WitnessGroup{}

				if err := utils.CopyObjectJson(apiGroupResp, &model); err != nil {
					if err != nil {
						diag.AddError("Read Error", fmt.Sprintf("Unable to copy witness group, got error: %s", err))
						return err
					}
				}

				*wgs = append(*wgs, model)
			}
		}
	}

	return nil
}

func NewPgdResource() resource.Resource {
	return &pgdResource{}
}
