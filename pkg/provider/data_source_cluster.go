package provider

import (
	"context"
	"fmt"
	"regexp"

	"github.com/EnterpriseDB/terraform-provider-biganimal/pkg/api"
	terraformCommon "github.com/EnterpriseDB/terraform-provider-biganimal/pkg/models/common/terraform"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var (
	_ datasource.DataSource              = &clusterDataSource{}
	_ datasource.DataSourceWithConfigure = &clusterDataSource{}
)

type PgConfigDatasourceModel struct {
	Value types.String `tfsdk:"value"`
	Name  types.String `tfsdk:"name"`
}

type StorageDatasourceModel struct {
	Throughput       types.String `tfsdk:"throughput"`
	VolumeProperties types.String `tfsdk:"volume_properties"`
	VolumeType       types.String `tfsdk:"volume_type"`
	Iops             types.String `tfsdk:"iops"`
	Size             types.String `tfsdk:"size"`
}

type ClusterArchitectureDatasourceModel struct {
	Nodes types.Int64  `tfsdk:"nodes"`
	Id    types.String `tfsdk:"id"`
	Name  types.String `tfsdk:"name"`
}

type clusterDatasourceModel struct {
	ID                         types.String                        `tfsdk:"id"`
	AllowedIpRanges            []AllowedIpRangesDatasourceModel    `tfsdk:"allowed_ip_ranges"`
	BackupRetentionPeriod      types.String                        `tfsdk:"backup_retention_period"`
	CreatedAt                  types.String                        `tfsdk:"created_at"`
	InstanceType               types.String                        `tfsdk:"instance_type"`
	ClusterName                types.String                        `tfsdk:"cluster_name"`
	FarawayReplicaIds          []string                            `tfsdk:"faraway_replica_ids"`
	LogsUrl                    types.String                        `tfsdk:"logs_url"`
	MostRecent                 types.Bool                          `tfsdk:"most_recent"`
	ClusterId                  types.String                        `tfsdk:"cluster_id"`
	PgConfig                   []PgConfigDatasourceModel           `tfsdk:"pg_config"`
	ExpiredAt                  types.String                        `tfsdk:"expired_at"`
	ReadOnlyConnections        types.Bool                          `tfsdk:"read_only_connections"`
	ResizingPvc                []string                            `tfsdk:"resizing_pvc"`
	CspAuth                    types.Bool                          `tfsdk:"csp_auth"`
	MetricsUrl                 types.String                        `tfsdk:"metrics_url"`
	CloudProvider              types.String                        `tfsdk:"cloud_provider"`
	Storage                    *StorageDatasourceModel             `tfsdk:"storage"`
	RoConnectionUri            types.String                        `tfsdk:"ro_connection_uri"`
	PrivateNetworking          types.Bool                          `tfsdk:"private_networking"`
	PgType                     types.String                        `tfsdk:"pg_type"`
	PgVersion                  types.String                        `tfsdk:"pg_version"`
	ClusterArchitecture        *ClusterArchitectureDatasourceModel `tfsdk:"cluster_architecture"`
	ProjectId                  types.String                        `tfsdk:"project_id"`
	ClusterType                types.String                        `tfsdk:"cluster_type"`
	Phase                      types.String                        `tfsdk:"phase"`
	DeletedAt                  types.String                        `tfsdk:"deleted_at"`
	ConnectionUri              types.String                        `tfsdk:"connection_uri"`
	Region                     types.String                        `tfsdk:"region"`
	FirstRecoverabilityPointAt types.String                        `tfsdk:"first_recoverability_point_at"`
	MaintenanceWindow          *terraformCommon.MaintenanceWindow  `tfsdk:"maintenance_window"`
	ServiceAccountIds          types.Set                           `tfsdk:"service_account_ids"`
	PeAllowedPrincipalIds      types.Set                           `tfsdk:"pe_allowed_principal_ids"`
	SuperuserAccess            types.Bool                          `tfsdk:"superuser_access"`
}

type AllowedIpRangesDatasourceModel struct {
	CidrBlock   types.String `tfsdk:"cidr_block"`
	Description types.String `tfsdk:"description"`
}

type clusterDataSource struct {
	client *api.ClusterClient
}

func (c *clusterDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_cluster"
}

// Configure adds the provider configured client to the data source.
func (c *clusterDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	c.client = req.ProviderData.(*api.API).ClusterClient()
}

func (c *clusterDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "The cluster data source describes a BigAnimal cluster. The data source requires your cluster name.",
		Attributes: map[string]schema.Attribute{
			"project_id": schema.StringAttribute{
				MarkdownDescription: "BigAnimal Project ID.",
				Required:            true,
				Validators: []validator.String{
					ProjectIdValidator(),
				},
			},

			"cluster_name": schema.StringAttribute{
				MarkdownDescription: "Name of the cluster.",
				Required:            true,
				Validators: []validator.String{
					stringvalidator.RegexMatches(regexp.MustCompile(`^\S+$`), "Cluster name should not be an empty string"),
				},
			},
			"most_recent": schema.BoolAttribute{
				MarkdownDescription: "Show the most recent cluster when there are multiple clusters with the same name.",
				Optional:            true,
			},

			"id": schema.StringAttribute{
				Description: "Datasource ID.",
				Computed:    true,
			},

			"allowed_ip_ranges": schema.SetNestedAttribute{
				MarkdownDescription: "Allowed IP ranges.",
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"cidr_block": schema.StringAttribute{
							MarkdownDescription: "CIDR block.",
							Computed:            true,
						},
						"description": schema.StringAttribute{
							MarkdownDescription: "CIDR block description.",
							Computed:            true,
						},
					},
				},
			},

			"backup_retention_period": schema.StringAttribute{
				MarkdownDescription: "Backup retention period.",
				Computed:            true,
			},
			"created_at": schema.StringAttribute{
				MarkdownDescription: "Cluster creation time.",
				Computed:            true,
			},
			"instance_type": schema.StringAttribute{
				MarkdownDescription: "Instance type.",
				Computed:            true,
			},
			"faraway_replica_ids": schema.SetAttribute{
				Computed:    true,
				Optional:    true,
				ElementType: types.StringType,
			},
			"logs_url": schema.StringAttribute{
				MarkdownDescription: "The URL to find the logs of this cluster.",
				Computed:            true,
			},

			"cluster_id": schema.StringAttribute{
				MarkdownDescription: "Cluster ID.",
				Computed:            true,
			},
			"pg_config": schema.SetNestedAttribute{
				MarkdownDescription: "Database configuration parameters.",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							MarkdownDescription: "GUC name.",
							Computed:            true,
						},
						"value": schema.StringAttribute{
							MarkdownDescription: "GUC value.",
							Computed:            true,
						},
					},
				},
			},
			"expired_at": schema.StringAttribute{
				MarkdownDescription: "Cluster expiry time.",
				Computed:            true,
			},
			"read_only_connections": schema.BoolAttribute{
				MarkdownDescription: "Is read only connection enabled.",
				Computed:            true,
			},
			"resizing_pvc": schema.ListAttribute{
				MarkdownDescription: "Resizing PVC.",
				Computed:            true,
				ElementType:         types.StringType,
			},
			"csp_auth": schema.BoolAttribute{
				MarkdownDescription: "Is authentication handled by the cloud service provider.",
				Computed:            true,
			},
			"metrics_url": schema.StringAttribute{
				MarkdownDescription: "The URL to find the metrics of this cluster.",
				Computed:            true,
			},
			"cloud_provider": schema.StringAttribute{
				MarkdownDescription: "Cloud provider.",
				Computed:            true,
			},
			"storage": schema.SingleNestedAttribute{
				MarkdownDescription: "Storage.",
				Computed:            true,
				Attributes: map[string]schema.Attribute{
					"volume_properties": schema.StringAttribute{
						MarkdownDescription: "Volume properties.",
						Computed:            true,
					},
					"volume_type": schema.StringAttribute{
						MarkdownDescription: "Volume type.",
						Computed:            true,
					},
					"iops": schema.StringAttribute{
						MarkdownDescription: "IOPS for the selected volume.",
						Computed:            true,
					},
					"size": schema.StringAttribute{
						MarkdownDescription: "Size of the volume.",
						Computed:            true,
					},
					"throughput": schema.StringAttribute{
						MarkdownDescription: "Throughput.",
						Computed:            true,
					},
				},
			},

			"ro_connection_uri": schema.StringAttribute{
				MarkdownDescription: "Cluster read-only connection URI. Only available for high availability clusters.",
				Computed:            true,
			},
			"private_networking": schema.BoolAttribute{
				MarkdownDescription: "Is private networking enabled.",
				Computed:            true,
			},
			"pg_type": schema.StringAttribute{
				MarkdownDescription: "Postgres type.",
				Computed:            true,
			},
			"pg_version": schema.StringAttribute{
				MarkdownDescription: "Postgres version.",
				Computed:            true,
			},
			"cluster_architecture": schema.SingleNestedAttribute{
				Computed:            true,
				MarkdownDescription: "Cluster architecture.",
				Attributes: map[string]schema.Attribute{
					"nodes": schema.Int64Attribute{
						MarkdownDescription: "Node count.",
						Computed:            true,
					},
					"id": schema.StringAttribute{
						MarkdownDescription: "Cluster architecture ID.",
						Computed:            true,
					},
					"name": schema.StringAttribute{
						MarkdownDescription: "Name.",
						Computed:            true,
					},
				},
			},
			"cluster_type": schema.StringAttribute{
				MarkdownDescription: "Type of the Specified Cluster.",
				Computed:            true,
			},
			"phase": schema.StringAttribute{
				MarkdownDescription: "Current phase of the cluster.",
				Computed:            true,
			},
			"deleted_at": schema.StringAttribute{
				MarkdownDescription: "Cluster deletion time.",
				Computed:            true,
			},
			"connection_uri": schema.StringAttribute{
				MarkdownDescription: "Cluster connection URI.",
				Computed:            true,
			},
			"region": schema.StringAttribute{
				MarkdownDescription: "Region to deploy the cluster.",
				Computed:            true,
			},
			"first_recoverability_point_at": schema.StringAttribute{
				MarkdownDescription: "Earliest backup recover time.",
				Computed:            true,
			},
			"maintenance_window": schema.SingleNestedAttribute{
				Computed:            true,
				MarkdownDescription: "Custom maintenance window.",
				Attributes: map[string]schema.Attribute{
					"is_enabled": schema.BoolAttribute{
						MarkdownDescription: "Is maintenance window enabled.",
						Computed:            true,
					},
					"start_day": schema.Int64Attribute{
						MarkdownDescription: "The day of week, 0 represents Sunday, 1 is Monday, and so on.",
						Computed:            true,
						Validators: []validator.Int64{
							int64validator.Between(0, 6),
						},
					},
					"start_time": schema.StringAttribute{
						MarkdownDescription: "Start time. \"hh:mm\", for example: \"23:59\".",
						Computed:            true,
						Validators: []validator.String{
							startTimeValidator(),
						},
					},
				},
			},
			"service_account_ids": schema.SetAttribute{
				Computed:    true,
				Optional:    true,
				ElementType: types.StringType,
			},
			"pe_allowed_principal_ids": schema.SetAttribute{
				Computed:    true,
				Optional:    true,
				ElementType: types.StringType,
			},
			"superuser_access": schema.BoolAttribute{
				MarkdownDescription: "Is superuser access enabled.",
				Computed:            true,
			},
			"pgvector": schema.BoolAttribute{
				MarkdownDescription: "Is pgvector extension enabled. Adds support for vector storage and vector similarity search to Postgres.",
				Computed:            true,
			},
			"pg_bouncer": schema.SingleNestedAttribute{
				MarkdownDescription: "Pg bouncer.",
				Optional:            true,
				Computed:            true,
				Attributes: map[string]schema.Attribute{
					"is_enabled": schema.BoolAttribute{
						MarkdownDescription: "Is pg bouncer enabled.",
						Required:            true,
					},
					"settings": schema.SetNestedAttribute{
						Description: "PgBouncer Configuration Settings.",
						Optional:    true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"name": schema.StringAttribute{
									Description: "Name.",
									Required:    true,
								},
								"operation": schema.StringAttribute{
									Description: "Operation.",
									Required:    true,
								},
								"value": schema.StringAttribute{
									Description: "Value.",
									Required:    true,
								},
							},
						},
					},
				},
			},
		},
	}
}

func (c *clusterDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	tflog.Error(ctx, "starting")
	var data clusterDatasourceModel
	diags := req.Config.Get(ctx, &data)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "read cluster data source")

	cluster, err := c.client.ReadByName(ctx, data.ProjectId.ValueString(), data.ClusterName.ValueString(), data.MostRecent.ValueBool())
	if err != nil {
		if !appendDiagFromBAErr(err, &resp.Diagnostics) {
			resp.Diagnostics.AddError("Error reading cluster", err.Error())
		}
		return
	}

	connection, err := c.client.ConnectionString(ctx, data.ProjectId.ValueString(), data.ClusterId.ValueString())
	if err != nil {
		if !appendDiagFromBAErr(err, &resp.Diagnostics) {
			resp.Diagnostics.AddError("Error reading cluster", err.Error())
		}
		return
	}

	data.ID = types.StringValue(fmt.Sprintf("%s/%s", data.ProjectId.ValueString(), data.ClusterId.ValueString()))
	data.ClusterId = types.StringPointerValue(cluster.ClusterId)
	data.ClusterName = types.StringPointerValue(cluster.ClusterName)
	data.ClusterType = types.StringPointerValue(cluster.ClusterType)
	data.Phase = types.StringPointerValue(cluster.Phase)
	data.CloudProvider = types.StringValue(cluster.Provider.CloudProviderId)
	data.ClusterArchitecture = &ClusterArchitectureDatasourceModel{
		Id:    types.StringValue(cluster.ClusterArchitecture.ClusterArchitectureId),
		Nodes: types.Int64Value(int64(cluster.ClusterArchitecture.Nodes)),
		Name:  types.StringValue(cluster.ClusterArchitecture.ClusterArchitectureName),
	}
	data.Region = types.StringValue(cluster.Region.Id)
	data.InstanceType = types.StringValue(cluster.InstanceType.InstanceTypeId)
	data.Storage = &StorageDatasourceModel{
		VolumeType:       types.StringPointerValue(cluster.Storage.VolumeTypeId),
		VolumeProperties: types.StringPointerValue(cluster.Storage.VolumePropertiesId),
		Size:             types.StringPointerValue(cluster.Storage.Size),
		Iops:             types.StringPointerValue(cluster.Storage.Iops),
		Throughput:       types.StringPointerValue(cluster.Storage.Throughput),
	}
	data.ResizingPvc = cluster.ResizingPvc
	data.ReadOnlyConnections = types.BoolPointerValue(cluster.ReadOnlyConnections)
	data.ConnectionUri = types.StringValue(connection.PgUri)
	data.RoConnectionUri = types.StringValue(connection.ReadOnlyPgUri)
	data.CspAuth = types.BoolPointerValue(cluster.CSPAuth)
	data.LogsUrl = types.StringPointerValue(cluster.LogsUrl)
	data.MetricsUrl = types.StringPointerValue(cluster.MetricsUrl)
	data.BackupRetentionPeriod = types.StringPointerValue(cluster.BackupRetentionPeriod)
	data.PgVersion = types.StringValue(cluster.PgVersion.PgVersionId)
	data.PgType = types.StringValue(cluster.PgType.PgTypeId)
	data.PrivateNetworking = types.BoolPointerValue(cluster.PrivateNetworking)
	data.SuperuserAccess = types.BoolPointerValue(cluster.SuperuserAccess)

	if cluster.FarawayReplicaIds != nil {
		data.FarawayReplicaIds = *cluster.FarawayReplicaIds
	}

	if cluster.FirstRecoverabilityPointAt != nil {
		data.FirstRecoverabilityPointAt = types.StringValue(cluster.FirstRecoverabilityPointAt.String())
	}

	data.PgConfig = []PgConfigDatasourceModel{}
	if configs := cluster.PgConfig; configs != nil {
		for _, kv := range *configs {
			data.PgConfig = append(data.PgConfig, PgConfigDatasourceModel{
				Name:  types.StringValue(kv.Name),
				Value: types.StringValue(kv.Value),
			})
		}
	}

	data.AllowedIpRanges = []AllowedIpRangesDatasourceModel{}
	if allowedIpRanges := cluster.AllowedIpRanges; allowedIpRanges != nil {
		for _, ipRange := range *allowedIpRanges {
			data.AllowedIpRanges = append(data.AllowedIpRanges, AllowedIpRangesDatasourceModel{
				CidrBlock:   types.StringValue(ipRange.CidrBlock),
				Description: types.StringValue(ipRange.Description),
			})
		}
	}

	if pt := cluster.CreatedAt; pt != nil {
		data.CreatedAt = types.StringValue(pt.String())
	}

	data.ExpiredAt = types.StringNull()
	if pt := cluster.ExpiredAt; pt != nil {
		data.ExpiredAt = types.StringValue(pt.String())
	}

	data.DeletedAt = types.StringNull()
	if pt := cluster.DeletedAt; pt != nil {
		data.DeletedAt = types.StringValue(pt.String())
	}

	if cluster.ServiceAccountIds != nil {
		serviceAccountIds := []attr.Value{}

		for _, v := range *cluster.ServiceAccountIds {
			serviceAccountIds = append(serviceAccountIds, types.StringValue(v))
		}

		data.ServiceAccountIds = types.SetValueMust(types.StringType, serviceAccountIds)
	}

	if cluster.PeAllowedPrincipalIds != nil {
		peAllowedPrincipalIds := []attr.Value{}
		for _, v := range *cluster.PeAllowedPrincipalIds {
			peAllowedPrincipalIds = append(peAllowedPrincipalIds, types.StringValue(v))
		}

		data.PeAllowedPrincipalIds = types.SetValueMust(types.StringType, peAllowedPrincipalIds)
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func NewClusterDataSource() datasource.DataSource {
	return &clusterDataSource{}
}
