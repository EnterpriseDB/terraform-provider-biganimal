package provider

import (
	"context"
	"fmt"

	"github.com/EnterpriseDB/terraform-provider-biganimal/pkg/api"
	"github.com/EnterpriseDB/terraform-provider-biganimal/pkg/models/tf"
	"github.com/EnterpriseDB/terraform-provider-biganimal/pkg/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var (
	_ datasource.DataSource              = &pgdDataSource{}
	_ datasource.DataSourceWithConfigure = &pgdDataSource{}
)

type pgdDataSource struct {
	client *api.API
}

func (p pgdDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_pgd"
}

// Configure adds the provider configured client to the data source.
func (d *pgdDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*api.API)
}

func (p pgdDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "The PGD cluster data source describes a BigAnimal cluster. The data source requires your PGD cluster name.",
		Attributes: map[string]schema.Attribute{
			"data_groups": schema.SetNestedAttribute{
				Description: "Cluster data groups.",
				Computed:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"group_id": schema.StringAttribute{
							Description: "Group ID of the group.",
							Computed:    true,
						},
						"allowed_ip_ranges": schema.SetNestedAttribute{
							Description: "Allowed ip ranges",
							Computed:    true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"cidr_block": schema.StringAttribute{
										Description: "CIDR block",
										Computed:    true,
									},
									"description": schema.StringAttribute{
										Description: "Description of CIDR block",
										Computed:    true,
									},
								},
							},
						},
						"backup_retention_period": schema.StringAttribute{
							Description: "Backup retention period",
							Optional:    true,
						},
						"cluster_architecture": schema.SingleNestedAttribute{
							Description: "Cluster architecture.",
							Computed:    true,
							Attributes: map[string]schema.Attribute{
								"cluster_architecture_id": schema.StringAttribute{
									Description: "Cluster architecture id.",
									Computed:    true,
								},
								"cluster_architecture_name": schema.StringAttribute{
									Description: "Cluster architecture name.",
									Computed:    true,
								},
								"nodes": schema.Float64Attribute{
									Description: "Nodes.",
									Computed:    true,
								},
								"witness_nodes": schema.Float64Attribute{
									Description: "Witness nodes.",
									Computed:    true,
								},
							},
						},
						"cluster_name": schema.StringAttribute{
							Description: "Name of the group.",
							Optional:    true,
						},
						"cluster_type": schema.StringAttribute{
							Description: "Type of the Specified Cluster",
							Optional:    true,
						},
						"created_at": schema.StringAttribute{
							Description: "Cluster creation time.",
							Optional:    true,
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
						"instance_type": schema.StringAttribute{
							Description: "Instance type.",
							Optional:    true,
						},
						"logs_url": schema.StringAttribute{
							Description: "The URL to find the logs of this cluster.",
							Optional:    true,
						},
						"metrics_url": schema.StringAttribute{
							Description: "The URL to find the metrics of this cluster.",
							Optional:    true,
						},
						"connection_uri": schema.StringAttribute{
							Description: "Cluster connection URI.",
							Optional:    true,
						},
						"pg_config": schema.SetNestedAttribute{
							Description: "Database configuration parameters.",
							Computed:    true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										Description: "GUC name.",
										Computed:    true,
									},
									"value": schema.StringAttribute{
										Description: "GUC value.",
										Computed:    true,
									},
								},
							},
						},
						"pg_type": schema.StringAttribute{
							Description: "Postgres type.",
							Optional:    true,
						},
						"pg_version": schema.StringAttribute{
							Description: "Postgres version.",
							Optional:    true,
						},
						"phase": schema.StringAttribute{
							Description: "Current phase of the cluster group.",
							Optional:    true,
						},
						"private_networking": schema.BoolAttribute{
							Description: "Is private networking enabled.",
							Optional:    true,
						},
						"cloud_provider": schema.StringAttribute{
							Description: "Cloud provider.",
							Optional:    true,
						},
						"csp_auth": schema.BoolAttribute{
							Description: "Is authentication handled by the cloud service provider.",
							Optional:    true,
						},
						"region": schema.StringAttribute{
							Description: "Data group region.",
							Optional:    true,
						},
						"resizing_pvc": schema.SetAttribute{
							ElementType: types.StringType,
							Optional:    true,
						},
						"storage": schema.SingleNestedAttribute{
							Description: "Storage.",
							Computed:    true,
							Attributes: map[string]schema.Attribute{
								"iops": schema.StringAttribute{
									Description: "IOPS for the selected volume.",
									Computed:    true,
								},
								"size": schema.StringAttribute{
									Description: "Size of the volume.",
									Computed:    true,
								},
								"throughput": schema.StringAttribute{
									Description: "Throughput.",
									Computed:    true,
								},
								"volume_properties": schema.StringAttribute{
									Description: "Volume properties.",
									Computed:    true,
								},
								"volume_type": schema.StringAttribute{
									Description: "Volume type.",
									Computed:    true,
								},
							},
						},
					},
				},
			},
			"witness_groups": schema.SetNestedAttribute{
				Description: "Cluster witness groups.",
				Computed:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"region": schema.StringAttribute{
							Description: "Witness group region.",
							Optional:    true,
						},
					},
				},
			},
			"project_id": schema.StringAttribute{
				Description: "BigAnimal Project ID.",
				Optional:    true,
			},
			"cluster_id": schema.StringAttribute{
				Description: "Cluster ID.",
				Optional:    true,
			},
			"cluster_name": schema.StringAttribute{
				Description: "cluster name",
				Optional:    true,
			},
			"most_recent": schema.BoolAttribute{
				Description: "Show the most recent cluster when there are multiple clusters with the same name",
				Optional:    true,
			},
		},
	}
}

func (p pgdDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data tf.PGDDataSourceData
	diags := req.Config.Get(ctx, &data)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Trace(ctx, "read pgd data source")

	if data.MostRecent == nil {
		data.MostRecent = utils.ToPointer(false)
	}

	cluster, err := p.client.ClusterClient().ReadByName(ctx, data.ProjectID, data.ClusterName, *data.MostRecent)
	if err != nil {
		resp.Diagnostics.AddError("Read error", fmt.Sprintf("Unable to call read cluster, got error: %s", err))
		return
	}

	if cluster.ClusterArchitecture.ClusterArchitectureId != "pgd" {
		resp.Diagnostics.AddError("Wrong cluster architecture error", fmt.Sprintf("Wrong cluster architecture, expected 'pgd' but got: %v", cluster.ClusterArchitecture.ClusterArchitectureId))
	}

	data.ClusterID = cluster.ClusterId

	for _, v := range *cluster.Groups {
		switch group := v.(type) {
		case map[string]interface{}:
			if group["clusterType"] == "data_group" {
				tfDg := tf.DataGroupData{}

				if err := utils.CopyObjectJson(group, &tfDg); err != nil {
					if err != nil {
						resp.Diagnostics.AddError("Read Error", fmt.Sprintf("Unable to copy data group, got error: %s", err))
						return
					}
				}
				data.DataGroups = append(data.DataGroups, tfDg)
			}

			if group["clusterType"] == "witness_group" {
				tfWg := tf.WitnessGroupData{}

				if err := utils.CopyObjectJson(group, &tfWg); err != nil {
					if err != nil {
						resp.Diagnostics.AddError("Read Error", fmt.Sprintf("Unable to copy witness group, got error: %s", err))
						return
					}
				}
				data.WitnessGroups = append(data.WitnessGroups, tfWg)
			}
		}
	}

	diags = resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func NewPgdDataSource() datasource.DataSource {
	return &pgdDataSource{}
}
