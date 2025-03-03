package provider

import (
	"context"
	"fmt"

	"github.com/EnterpriseDB/terraform-provider-biganimal/pkg/api"
	"github.com/EnterpriseDB/terraform-provider-biganimal/pkg/models/pgd/terraform"
	"github.com/EnterpriseDB/terraform-provider-biganimal/pkg/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var (
	_ datasource.DataSource              = &pgdDataSource{}
	_ datasource.DataSourceWithConfigure = &pgdDataSource{}
)

type pgdDataSource struct {
	client *api.PGDClient
}

func (p pgdDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_pgd"
}

// Configure adds the provider configured client to the data source.
func (d *pgdDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*api.API).PGDClient()
}

func (p pgdDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "The PGD cluster data source describes a BigAnimal cluster. The data source requires your PGD cluster name.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"project_id": schema.StringAttribute{
				Description: "BigAnimal Project ID.",
				Required:    true,
				Validators: []validator.String{
					ProjectIdValidator(),
				},
			},
			"cluster_id": schema.StringAttribute{
				Description: "Cluster ID.",
				Computed:    true,
			},
			"cluster_name": schema.StringAttribute{
				Description: "cluster name",
				Required:    true,
			},
			"most_recent": schema.BoolAttribute{
				Description: "Show the most recent cluster when there are multiple clusters with the same name",
				Optional:    true,
			},
			"data_groups": schema.SetNestedAttribute{
				Description: "Cluster data groups.",
				Computed:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"group_id": schema.StringAttribute{
							Description: "Group ID of the group.",
							Computed:    true,
						},
						"backup_retention_period": schema.StringAttribute{
							Description: "Backup retention period",
							Computed:    true,
						},
						"cluster_name": schema.StringAttribute{
							Description: "Name of the group.",
							Computed:    true,
						},
						"cluster_type": schema.StringAttribute{
							Description: "Type of the Specified Cluster",
							Computed:    true,
						},
						"created_at": schema.StringAttribute{
							Description: "Cluster creation time.",
							Computed:    true,
						},
						"logs_url": schema.StringAttribute{
							Description: "The URL to find the logs of this cluster.",
							Computed:    true,
						},
						"metrics_url": schema.StringAttribute{
							Description: "The URL to find the metrics of this cluster.",
							Computed:    true,
						},
						"connection_uri": schema.StringAttribute{
							Description: "Data group connection URI.",
							Computed:    true,
						},
						"phase": schema.StringAttribute{
							Description: "Current phase of the cluster group.",
							Computed:    true,
						},
						"private_networking": schema.BoolAttribute{
							Description: "Is private networking enabled.",
							Computed:    true,
						},
						"csp_auth": schema.BoolAttribute{
							Description: "Is authentication handled by the cloud service provider.",
							Computed:    true,
						},
						"resizing_pvc": schema.SetAttribute{
							Description: "Resizing PVC.",
							Computed:    true,
							ElementType: types.StringType,
						},
						"allowed_ip_ranges": schema.SetNestedAttribute{
							Description: "Allowed IP ranges.",
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
						"cluster_architecture": schema.SingleNestedAttribute{
							Description: "Cluster architecture.",
							Computed:    true,
							Attributes: map[string]schema.Attribute{
								"cluster_architecture_id": schema.StringAttribute{
									Description: "Cluster architecture ID.",
									Computed:    true,
								},
								"cluster_architecture_name": schema.StringAttribute{
									Description: "Cluster architecture name.",
									Computed:    true,
								},
								"nodes": schema.Float64Attribute{
									Description: "Node count.",
									Computed:    true,
								},
								"witness_nodes": schema.Int64Attribute{
									Description: "Witness nodes count.",
									Computed:    true,
								},
							},
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
									Required:    true,
								},
							},
						},
						"pg_type": schema.SingleNestedAttribute{
							Description: "Postgres type.",
							Computed:    true,
							Attributes: map[string]schema.Attribute{
								"pg_type_id": schema.StringAttribute{
									Description: "Data group pg type id.",
									Computed:    true,
								},
							},
						},
						"pg_version": schema.SingleNestedAttribute{
							Description: "Postgres version.",
							Computed:    true,
							Attributes: map[string]schema.Attribute{
								"pg_version_id": schema.StringAttribute{
									Description: "Data group pg version id.",
									Computed:    true,
								},
							},
						},
						"cloud_provider": schema.SingleNestedAttribute{
							Description: "Cloud provider.",
							Computed:    true,
							Attributes: map[string]schema.Attribute{
								"cloud_provider_id": schema.StringAttribute{
									Description: "Data group cloud provider id.",
									Computed:    true,
								},
							},
						},
						"region": schema.SingleNestedAttribute{
							Description: "Region.",
							Computed:    true,
							Attributes: map[string]schema.Attribute{
								"region_id": schema.StringAttribute{
									Description: "Data group region id.",
									Computed:    true,
								},
							},
						},
						"instance_type": schema.SingleNestedAttribute{
							Description: "Instance type.",
							Computed:    true,
							Attributes: map[string]schema.Attribute{
								"instance_type_id": schema.StringAttribute{
									Description: "Data group instance type id.",
									Computed:    true,
								},
							},
						},
						"maintenance_window": schema.SingleNestedAttribute{
							Description: "Custom maintenance window.",
							Computed:    true,
							Attributes: map[string]schema.Attribute{
								"is_enabled": schema.BoolAttribute{
									Description: "Is maintenance window enabled.",
									Computed:    true,
								},
								"start_day": schema.Float64Attribute{
									Description: "Start day.",
									Computed:    true,
								},
								"start_time": schema.StringAttribute{
									Description: "Start time.",
									Computed:    true,
								},
							},
						},
						"service_account_ids": schema.SetAttribute{
							Description: "A Google Cloud Service Account is used for logs. If you leave this blank, then you will be unable to access log details for this cluster. Required when cluster is deployed on BigAnimal's cloud account.",
							Computed:    true,
							ElementType: types.StringType,
						},

						"pe_allowed_principal_ids": schema.SetAttribute{
							Description: "Cloud provider subscription/account ID, need to be specified when cluster is deployed on BigAnimal's cloud account.",
							Computed:    true,
							ElementType: types.StringType,
						},
						"ro_connection_uri": schema.SetAttribute{
							Description: "Read-only connection URI.",
							Computed:    true,
							ElementType: types.StringType,
						},
						"read_only_connections": schema.BoolAttribute{
							Description: "Is read-only connections enabled.",
							Computed:    true,
						},
						"backup_schedule_time": ResourceBackupScheduleTime,
						"wal_storage":          resourceWal,
						"private_link_service_alias": schema.StringAttribute{
							MarkdownDescription: "Private link service alias.",
							Computed:            true,
						},
						"private_link_service_name": schema.StringAttribute{
							MarkdownDescription: "private link service name.",
							Computed:            true,
						},
					},
				},
			},
			"witness_groups": schema.SetNestedAttribute{
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"group_id": schema.StringAttribute{
							Description: "Group id of witness group.",
							Computed:    true,
						},
						"phase": schema.StringAttribute{
							Description: "Current phase of the witness group.",
							Computed:    true,
						},
						"cluster_architecture": schema.SingleNestedAttribute{
							Description: "Cluster architecture.",
							Computed:    true,
							Attributes: map[string]schema.Attribute{
								"cluster_architecture_id": schema.StringAttribute{
									Description: "Cluster architecture ID.",
									Computed:    true,
								},
								"cluster_architecture_name": schema.StringAttribute{
									Description: "Name.",
									Computed:    true,
								},
								"nodes": schema.Float64Attribute{
									Description: "Nodes.",
									Computed:    true,
								},
								"witness_nodes": schema.Float64Attribute{
									Description: "Witness nodes count.",
									Computed:    true,
								},
							},
						},
						"region": schema.SingleNestedAttribute{
							Description: "Region.",
							Computed:    true,
							Attributes: map[string]schema.Attribute{
								"region_id": schema.StringAttribute{
									Description: "Region id.",
									Computed:    true,
								},
							},
						},
						"cluster_type": schema.StringAttribute{
							Description: "Type of the Specified Cluster",
							Computed:    true,
						},
						"cloud_provider": schema.SingleNestedAttribute{
							Description: "Cloud provider.",
							Computed:    true,
							Attributes: map[string]schema.Attribute{
								"cloud_provider_id": schema.StringAttribute{
									Description: "Cloud provider id.",
									Computed:    true,
								},
							},
						},
						"instance_type": schema.SingleNestedAttribute{
							Description: "Instance type.",
							Computed:    true,
							Attributes: map[string]schema.Attribute{
								"instance_type_id": schema.StringAttribute{
									Description: "Witness group instance type id.",
									Computed:    true,
								},
							},
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
						"maintenance_window": schema.SingleNestedAttribute{
							Description: "Custom maintenance window.",
							Computed:    true,
							Attributes: map[string]schema.Attribute{
								"is_enabled": schema.BoolAttribute{
									Description: "Is maintenance window enabled.",
									Computed:    true,
								},
								"start_day": schema.Int64Attribute{
									Description: "The day of week, 0 represents Sunday, 1 is Monday, and so on.",
									Computed:    true,
								},
								"start_time": schema.StringAttribute{
									Description: "Start time. \"hh:mm\", for example: \"23:59\".",
									Computed:    true,
								},
							},
						},
					},
				},
			},
		},
	}
}

type PGDDataSourceData struct {
	ID            *string                  `tfsdk:"id"`
	ProjectID     string                   `tfsdk:"project_id"`
	ClusterID     *string                  `tfsdk:"cluster_id"`
	ClusterName   string                   `tfsdk:"cluster_name"`
	MostRecent    *bool                    `tfsdk:"most_recent"`
	DataGroups    []terraform.DataGroup    `tfsdk:"data_groups"`
	WitnessGroups []terraform.WitnessGroup `tfsdk:"witness_groups"`
}

func (p pgdDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data PGDDataSourceData

	diags := req.Config.Get(ctx, &data)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Trace(ctx, "read pgd data source")

	if data.MostRecent == nil {
		data.MostRecent = utils.ToPointer(false)
	}

	cluster, err := p.client.PGDClient().ReadByName(ctx, data.ProjectID, data.ClusterName, *data.MostRecent)
	if err != nil {
		if appendDiagFromBAErr(err, &resp.Diagnostics) {
			return
		}
		resp.Diagnostics.AddError("Read error", fmt.Sprintf("Unable to call read cluster, got error: %s", err))
		return
	}

	if cluster.ClusterArchitecture.ClusterArchitectureId != "pgd" {
		resp.Diagnostics.AddError("Wrong cluster architecture error", fmt.Sprintf("Wrong cluster architecture, expected 'pgd' but got: %v", cluster.ClusterArchitecture.ClusterArchitectureId))
		return
	}

	data.ID = cluster.ClusterId
	data.ClusterID = cluster.ClusterId

	buildGroups := PGD{}

	buildTFGroupsAs(ctx, &resp.Diagnostics, resp.State, *cluster, &buildGroups)
	if resp.Diagnostics.HasError() {
		return
	}

	data.DataGroups = buildGroups.DataGroups
	data.WitnessGroups = buildGroups.WitnessGroups

	diags = resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func NewPgdDataSource() datasource.DataSource {
	return &pgdDataSource{}
}
