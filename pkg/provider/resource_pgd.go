package provider

import (
	"context"

	"github.com/EnterpriseDB/terraform-provider-biganimal/pkg/api"
	"github.com/EnterpriseDB/terraform-provider-biganimal/pkg/models"
	"github.com/EnterpriseDB/terraform-provider-biganimal/pkg/models/pgd"
	"github.com/EnterpriseDB/terraform-provider-biganimal/pkg/utils"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
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
						// 	"pg_config": schema.SetNestedBlock{
						// 		Description: "Database configuration parameters.",

						// 		NestedObject: schema.NestedBlockObject{
						// 			Attributes: map[string]schema.Attribute{
						// 				"name": schema.StringAttribute{
						// 					Description: "GUC name.",
						// 					Optional:    true,
						// 				},
						// 				"value": schema.StringAttribute{
						// 					Description: "GUC value.",
						// 					Optional:    true,
						// 				},
						// 			},
						// 		},
						// 	},
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
									Description: "Node count.",
									Optional:    true,
								},
								"witness_nodes": schema.Float64Attribute{
									Description: "Witness nodes count.",
									Optional:    true,
								},
							},
						},
						// 	"storage": schema.SingleNestedBlock{
						// 		Description: "Storage.",
						// 		Attributes: map[string]schema.Attribute{
						// 			"iops": schema.StringAttribute{
						// 				Description: "IOPS for the selected volume.",
						// 				Optional:    true,
						// 			},
						// 			"size": schema.StringAttribute{
						// 				Description: "Size of the volume.",
						// 				Optional:    true,
						// 			},
						// 			"throughput": schema.StringAttribute{
						// 				Description: "Throughput.",
						// 				Optional:    true,
						// 			},
						// 			"volume_properties": schema.StringAttribute{
						// 				Description: "Volume properties.",
						// 				Optional:    true,
						// 			},
						// 			"volume_type": schema.StringAttribute{
						// 				Description: "Volume type.",
						// 				Optional:    true,
						// 			},
						// 		},
						// 	},
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
						// "cluster_name": schema.StringAttribute{
						// 	Description: "Name of the group.",
						// 	Optional:    true,
						// },
						// "cluster_type": schema.StringAttribute{
						// 	Description: "Type of the Specified Cluster",
						// 	Optional:    true,
						// },
						// "created_at": schema.StringAttribute{
						// 	Description: "Cluster creation time.",
						// 	Optional:    true,
						// },
						// "deleted_at": schema.StringAttribute{
						// 	Description: "Cluster deletion time.",
						// 	Optional:    true,
						// },
						// "expired_at": schema.StringAttribute{
						// 	Description: "Cluster expiry time.",
						// 	Optional:    true,
						// },
						// "first_recoverability_point_at": schema.StringAttribute{
						// 	Description: "Earliest backup recover time.",
						// 	Optional:    true,
						// },
						// "instance_type": schema.StringAttribute{
						// 	Description: "Instance type.",
						// 	Optional:    true,
						// },
						// "logs_url": schema.StringAttribute{
						// 	Description: "The URL to find the logs of this cluster.",
						// 	Optional:    true,
						// },
						// "metrics_url": schema.StringAttribute{
						// 	Description: "The URL to find the metrics of this cluster.",
						// 	Optional:    true,
						// },
						// "connection_uri": schema.StringAttribute{
						// 	Description: "Cluster connection URI.",
						// 	Optional:    true,
						// },
						// "pg_type": schema.StringAttribute{
						// 	Description: "Postgres type.",
						// 	Optional:    true,
						// },
						// "pg_version": schema.StringAttribute{
						// 	Description: "Postgres version.",
						// 	Optional:    true,
						// },
						// "phase": schema.StringAttribute{
						// 	Description: "Current phase of the cluster group.",
						// 	Optional:    true,
						// },
						// "private_networking": schema.BoolAttribute{
						// 	Description: "Is private networking enabled.",
						// 	Optional:    true,
						// },
						// "cloud_provider": schema.StringAttribute{
						// 	Description: "Cloud provider.",
						// 	Optional:    true,
						// },
						"csp_auth": schema.BoolAttribute{
							Description: "Is authentication handled by the cloud service provider.",
							Optional:    true,
						},
						"region": schema.StringAttribute{
							Description: "Data group region.",
							Optional:    true,
						},
						// "resizing_pvc": schema.SetAttribute{
						// 	ElementType: types.StringType,
						// 	Optional:    true,
						// },
					},
				},
			},
			// "witness_groups": schema.SetNestedBlock{
			// 	Description: "Cluster witness groups.",
			// 	NestedObject: schema.NestedBlockObject{
			// 		Attributes: map[string]schema.Attribute{
			// 			"region": schema.StringAttribute{
			// 				Description: "Witness group region.",
			// 				Optional:    true,
			// 			},
			// 		},
			// 	},
			// },
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
			// "cluster_id": schema.StringAttribute{
			// 	Description: "Cluster ID.",
			// 	Optional:    true,
			// },
			"cluster_name": schema.StringAttribute{
				Description: "cluster name",
				Required:    true,
			},
			// "most_recent": schema.BoolAttribute{
			// 	Description: "Show the most recent cluster when there are multiple clusters with the same name",
			// 	Optional:    true,
			// },
			// "password": schema.StringAttribute{
			// 	Description: "Password for the user edb_admin. It must be 12 characters or more.",
			// 	Required:    true,
			// 	Sensitive:   true,
			// },
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
	ID        *string `tfsdk:"id"`
	ProjectId string  `tfsdk:"project_id"`
	// ClusterId   *string `tfsdk:"cluster_id"`
	ClusterName string `tfsdk:"cluster_name"`
	// MostRecent    *bool              `tfsdk:"most_recent"`
	DataGroups []pgd.DataGroup `tfsdk:"data_groups"`
	// WitnessGroups []pgd.WitnessGroup `tfsdk:"witness_groups"`
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
		Password:    utils.ToPointer("asdfasdfasdfafchgf67sdfds"),
	}

	for _, v := range config.DataGroups {
		*clusterReqBody.Groups = append(*clusterReqBody.Groups, v)
	}
	// for _, v := range config.WitnessGroups {
	// 	*clusterReqBody.Groups = append(*clusterReqBody.Groups, v)
	// }

	// clusterId, err := p.client.Create(ctx, config.ProjectId, clusterReqBody)
	// if err != nil {
	// 	resp.Diagnostics.AddError("Error creating PGD cluster", "Could not create PGD cluster, unexpected error: "+err.Error())
	// 	return
	// }

	// clusterResp, err := p.client.Read(ctx, config.ProjectId, clusterId)
	// if err != nil {
	// 	resp.Diagnostics.AddError("Error reading PGD cluster", "Could not read PGD cluster, unexpected error: "+err.Error())
	// 	return
	// }

	// config.ClusterId = clusterResp.ClusterId

	diags = resp.State.Set(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (p pgdResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
}

func (p pgdResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
}

func (p pgdResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
}

func NewPgdResource() resource.Resource {
	return &pgdResource{}
}
