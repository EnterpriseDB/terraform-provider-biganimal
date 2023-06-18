package provider

import (
	"context"
	"fmt"
	"time"

	"github.com/EnterpriseDB/terraform-provider-biganimal/pkg/api"
	apiModels "github.com/EnterpriseDB/terraform-provider-biganimal/pkg/models/api"
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
							Computed:    true,
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
						"deleted_at": schema.StringAttribute{
							Description: "Cluster deletion time.",
							Computed:    true,
						},
						"expired_at": schema.StringAttribute{
							Description: "Cluster expiry time.",
							Computed:    true,
						},
						"first_recoverability_point_at": schema.StringAttribute{
							Description: "Earliest backup recover time.",
							Computed:    true,
						},
						"instance_type": schema.StringAttribute{
							Description: "Instance type.",
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
							Description: "Cluster connection URI.",
							Computed:    true,
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
							Computed:    true,
						},
						"pg_version": schema.StringAttribute{
							Description: "Postgres version.",
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
						"cloud_provider": schema.StringAttribute{
							Description: "Cloud provider.",
							Computed:    true,
						},
						"csp_auth": schema.BoolAttribute{
							Description: "Is authentication handled by the cloud service provider.",
							Computed:    true,
						},
						"region": schema.StringAttribute{
							Description: "Data group region.",
							Computed:    true,
						},
						"resizing_pvc": schema.SetAttribute{
							ElementType: types.StringType,
							Computed:    true,
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
							Computed:    true,
						},
					},
				},
			},
			"project_id": schema.StringAttribute{
				Description: "BigAnimal Project ID.",
				Required:    true,
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
		},
	}
}

type DataGroupData struct {
	GroupId                    string                                        `tfsdk:"group_id"`
	AllowedIpRanges            []apiModels.ClusterAllowedIpRange             `tfsdk:"allowed_ip_ranges"`
	BackupRetentionPeriod      string                                        `tfsdk:"backup_retention_period"`
	ClusterArchitecture        *apiModels.ClusterClusterArchitectureResponse `tfsdk:"cluster_architecture"`
	ClusterName                string                                        `tfsdk:"cluster_name"`
	ClusterType                string                                        `tfsdk:"cluster_type"`
	CreatedAt                  *string                                       `tfsdk:"created_at"`
	DeletedAt                  *string                                       `tfsdk:"deleted_at"`
	ExpiredAt                  *string                                       `tfsdk:"expired_at"`
	FirstRecoverabilityPointAt *string                                       `tfsdk:"first_recoverability_point_at"`
	InstanceType               *string                                       `tfsdk:"instance_type"`
	LogsUrl                    *string                                       `tfsdk:"logs_url"`
	MetricsUrl                 *string                                       `tfsdk:"metrics_url"`
	Connection                 *string                                       `tfsdk:"connection_uri"`
	PgConfig                   *[]apiModels.ArrayOfNameValueObjectsInner     `tfsdk:"pg_config"`
	PgType                     *string                                       `tfsdk:"pg_type"`
	PgVersion                  *string                                       `tfsdk:"pg_version"`
	Phase                      string                                        `tfsdk:"phase"`
	PrivateNetworking          bool                                          `tfsdk:"private_networking"`
	Provider                   *string                                       `tfsdk:"cloud_provider"`
	CspAuth                    *bool                                         `tfsdk:"csp_auth"`
	Region                     *string                                       `tfsdk:"region"`
	ResizingPvc                *[]string                                     `tfsdk:"resizing_pvc"`
	Storage                    *tf.ClusterStorageResponse                    `tfsdk:"storage"`
}

type WitnessGroupData struct {
	Region *string `tfsdk:"region"`
}

type PGDDataSourceData struct {
	ProjectID     string             `tfsdk:"project_id"`
	ClusterID     *string            `tfsdk:"cluster_id"`
	ClusterName   string             `tfsdk:"cluster_name"`
	MostRecent    *bool              `tfsdk:"most_recent"`
	DataGroups    []DataGroupData    `tfsdk:"data_groups"`
	WitnessGroups []WitnessGroupData `tfsdk:"witness_groups"`
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
		switch apiGroupResp := v.(type) {
		case map[string]interface{}:
			if apiGroupResp["clusterType"] == "data_group" {
				apiDgModel := apiModels.ClusterDataGroup{}
				tfDgModel := DataGroupData{}

				if err := utils.CopyObjectJson(apiGroupResp, &apiDgModel); err != nil {
					if err != nil {
						resp.Diagnostics.AddError("Read Error", fmt.Sprintf("Unable to copy data group, got error: %s", err))
						return
					}
				}

				tfDgModel.GroupId = apiDgModel.GroupId
				tfDgModel.AllowedIpRanges = apiDgModel.AllowedIpRanges
				tfDgModel.BackupRetentionPeriod = apiDgModel.BackupRetentionPeriod
				tfDgModel.ClusterArchitecture = apiDgModel.ClusterArchitecture
				tfDgModel.ClusterName = apiDgModel.ClusterName
				tfDgModel.ClusterType = apiDgModel.ClusterType

				if apiDgModel.CreatedAt != nil {
					tfDgModel.CreatedAt = utils.ToPointer(time.Unix(int64(apiDgModel.CreatedAt.Seconds), int64(apiDgModel.CreatedAt.Nanos)).String())
				}

				if apiDgModel.DeletedAt != nil {
					tfDgModel.DeletedAt = utils.ToPointer(time.Unix(int64(apiDgModel.DeletedAt.Seconds), int64(apiDgModel.DeletedAt.Nanos)).String())
				}

				if apiDgModel.ExpiredAt != nil {
					tfDgModel.ExpiredAt = utils.ToPointer(time.Unix(int64(apiDgModel.ExpiredAt.Seconds), int64(apiDgModel.ExpiredAt.Nanos)).String())
				}

				if apiDgModel.FirstRecoverabilityPointAt != nil {
					tfDgModel.FirstRecoverabilityPointAt = utils.ToPointer(time.Unix(int64(apiDgModel.FirstRecoverabilityPointAt.Seconds), int64(apiDgModel.FirstRecoverabilityPointAt.Nanos)).String())
				}

				if apiDgModel.InstanceType != nil {
					tfDgModel.InstanceType = utils.ToPointer(apiDgModel.InstanceType.InstanceTypeId)
				}

				tfDgModel.LogsUrl = apiDgModel.LogsUrl
				tfDgModel.MetricsUrl = apiDgModel.MetricsUrl

				if apiDgModel.Connection != nil {
					tfDgModel.Connection = apiDgModel.Connection.PgUri
				}

				if apiDgModel.PgConfig != nil {
					tfDgModel.PgConfig = apiDgModel.PgConfig
				}

				if apiDgModel.PgType != nil {
					tfDgModel.PgType = &apiDgModel.PgType.PgTypeId
				}

				tfDgModel.PgVersion = &apiDgModel.PgVersion.PgVersionId

				tfDgModel.Phase = apiDgModel.Phase
				tfDgModel.PrivateNetworking = apiDgModel.PrivateNetworking

				if apiDgModel.Provider != nil {
					tfDgModel.Provider = &apiDgModel.Provider.CloudProviderId
				}

				tfDgModel.CspAuth = apiDgModel.CspAuth

				if apiDgModel.Region != nil {
					tfDgModel.Region = &apiDgModel.Region.RegionId
				}

				if apiDgModel.ResizingPvc != nil {
					tfDgModel.ResizingPvc = apiDgModel.ResizingPvc
				}

				if apiDgModel.Storage != nil {
					tfDgModel.Storage = &tf.ClusterStorageResponse{
						Iops:               apiDgModel.Storage.Iops,
						Size:               apiDgModel.Storage.Size,
						Throughput:         apiDgModel.Storage.Throughput,
						VolumePropertiesId: apiDgModel.Storage.VolumePropertiesId,
						VolumeTypeId:       apiDgModel.Storage.VolumeTypeId,
					}
				}

				data.DataGroups = append(data.DataGroups, tfDgModel)
			}

			if apiGroupResp["clusterType"] == "witness_group" {
				apiWgModel := apiModels.ClusterWitnessGroup{}
				tfWgModel := WitnessGroupData{}

				if err := utils.CopyObjectJson(apiGroupResp, &apiWgModel); err != nil {
					if err != nil {
						resp.Diagnostics.AddError("Read Error", fmt.Sprintf("Unable to copy witness group, got error: %s", err))
						return
					}
				}

				if apiWgModel.Region != nil {
					tfWgModel.Region = &apiWgModel.Region.RegionId
				}

				data.WitnessGroups = append(data.WitnessGroups, tfWgModel)
			}
		}
	}

	diags = resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func NewPgdDataSource() datasource.DataSource {
	return &pgdDataSource{}
}
