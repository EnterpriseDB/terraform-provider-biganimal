package cluster

import (
	"context"
	"errors"
	"time"

	"github.com/EnterpriseDB/terraform-provider-biganimal/pkg/api"
	"github.com/EnterpriseDB/terraform-provider-biganimal/pkg/apiv2"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceCluster() *schema.Resource {
	return &schema.Resource{
		Description: "Sample cluster data source in the BigAnimal terraform provider .",
		ReadContext: DataSourceClusterRead,
		Schema: map[string]*schema.Schema{
			"allowed_ip_ranges": {
				Description: "Allowed IP ranges",
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cidr_block": {
							Description: "CIDR Block",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"description": {
							Description: "CIDR Block Description",
							Type:        schema.TypeString,
							Computed:    true,
						},
					},
				},
			},
			"backup_retention_period": {
				Description: "Backup Retention Period.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"cluster_architecture": {
				Description: "Cluster Architecture",
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Description: "ID",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"name": {
							Description: "Name",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"nodes": {
							Description: "Node Count",
							Type:        schema.TypeInt,
							Computed:    true,
						},
					},
				},
			},

			"cluster_name": {
				Description: "Name of the cluster.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"created_at": {
				Description: "Cluster Creation Time",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"deleted_at": {
				Description: "Cluster Deletion Time",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"expired_at": {
				Description: "Cluster Expiry Time",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"first_recoverability_point_at": {
				Description: "Cluster Expiry Time",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"instance_type": {
				Description: "Instance Type",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"category": {
							Description: "Instance category",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"cpu": {
							Description: "core count",
							Type:        schema.TypeInt,
							Computed:    true,
						},
						"family_name": {
							Description: "Family Name",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"instance_type_id": {
							Description: "Instance ID",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"instance_type_name": {
							Description: "Instance Name",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"ram": {
							Description: "Memory in Mb",
							Type:        schema.TypeFloat,
							Computed:    true,
						},
					},
				},
			},
			"id": {
				Description: "cluster ID",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"pg_config": {
				Description: "Instance Type",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Description: "GUC Name",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"value": {
							Description: "GUC Value",
							Type:        schema.TypeString,
							Computed:    true,
						},
					},
				},
			},
			"pg_type": {
				Description: "Allowed IP ranges",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Description: "ID",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"name": {
							Description: "Name",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"supported_cluster_architecture_ids": {
							Description: "Supported Cluster Architectures",
							Type:        schema.TypeList,
							Computed:    true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
			"pg_version": {
				Description: "Postgres type",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"phase": {
				Description: "Current Phase of the cluster.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"private_networking": {
				Description: "Is private networking enabled",
				Type:        schema.TypeBool,
				Computed:    true,
			},
			"cloud_provider": {
				Description: "Cloud Provider",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"region": {
				Description: "Region",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"replicas": {
				Description: "Replicas",
				Type:        schema.TypeInt,
				Computed:    true,
			},
			"resizing_pvc": {
				Description: "Resizing PVC",
				Type:        schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Computed: true,
			},
			"storage": {
				Description: "Storage",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"iops": {
							Description: "IOPS",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"size": {
							Description: "Size",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"throughput": {
							Description: "Throughput",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"volume_properties": {
							Description: "Volume Properties",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"volume_type": {
							Description: "Volume Type",
							Type:        schema.TypeString,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func DataSourceClusterRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	diags := diag.Diagnostics{}
	client := api.BuildAPI(meta, api.ClusterClientType).ClusterClient()

	clusterName, ok := d.Get("cluster_name").(string)
	if !ok {
		return diag.FromErr(errors.New("Unable to find cluster id"))
	}

	cluster, err := client.ReadByName(ctx, clusterName)
	if err != nil {
		return diag.FromErr(err)
	}

	// set the outputs
	d.Set("backup_retention_period", cluster.BackupRetentionPeriod)
	d.Set("cluster_architecture", getClusterArchitectureData(cluster))
	d.Set("created_at", pointInTimeToString(*cluster.CreatedAt))
	d.Set("cluster_name", &cluster.ClusterName)

	if cluster.DeletedAt != (*apiv2.PointInTime)(nil) {
		d.Set("deleted_at", pointInTimeToString(*cluster.DeletedAt))
	}

	if cluster.ExpiredAt != (*apiv2.PointInTime)(nil) {
		d.Set("expired_at", pointInTimeToString(*cluster.ExpiredAt))
	}

	if cluster.FirstRecoverabilityPointAt != (*apiv2.PointInTime)(nil) {
		d.Set("first_recoverability_point_at", pointInTimeToString(*cluster.FirstRecoverabilityPointAt))
	}

	d.Set("instance_type", getInstanceTypeData(cluster))
	d.Set("id", cluster.ClusterId)
	d.Set("pg_config", getPgConfigData(cluster))
	d.Set("pg_type", getPgTypeData(cluster))
	d.Set("pg_version", cluster.PgVersion.PgVersionName)
	d.Set("phase", cluster.Phase)
	d.Set("private_networking", cluster.PrivateNetworking)
	d.Set("cloud_provider", cluster.Provider.CloudProviderId)
	d.Set("region", cluster.Region.RegionId)
	d.Set("replicas", cluster.Replicas)
	d.Set("storage", getStorageData(cluster))
	d.Set("resizing_pvc", cluster.ResizingPvc)

	d.SetId(cluster.ClusterId)

	return diags
}

func pointInTimeToString(p apiv2.PointInTime) string {
	return time.Unix(int64(p.Seconds), int64(p.Nanos)).String()
}

func getClusterArchitectureData(cluster *apiv2.ClusterDetail) []interface{} {
	propMap := map[string]interface{}{}
	propMap["id"] = cluster.ClusterArchitecture.ClusterArchitectureId
	propMap["name"] = cluster.ClusterArchitecture.ClusterArchitectureName
	propMap["nodes"] = cluster.ClusterArchitecture.Nodes

	return []interface{}{propMap}
}

func getInstanceTypeData(cluster *apiv2.ClusterDetail) []interface{} {
	propMap := map[string]interface{}{}
	propMap["category"] = cluster.InstanceType.Category
	propMap["cpu"] = cluster.InstanceType.Cpu
	propMap["family_name"] = cluster.InstanceType.FamilyName
	propMap["instance_type_id"] = cluster.InstanceType.InstanceTypeId
	propMap["instance_type_name"] = cluster.InstanceType.InstanceTypeName
	propMap["ram"] = cluster.InstanceType.Ram

	return []interface{}{propMap}
}

func getPgConfigData(cluster *apiv2.ClusterDetail) []interface{} {
	list := []interface{}{}

	for _, guc := range *cluster.PgConfig {
		propMap := map[string]interface{}{}
		propMap["name"] = guc.Name
		propMap["value"] = guc.Value
		list = append(list, propMap)
	}

	return list
}

func getPgTypeData(cluster *apiv2.ClusterDetail) []interface{} {
	propMap := map[string]interface{}{}
	propMap["id"] = cluster.PgType.PgTypeId
	propMap["name"] = cluster.PgType.PgTypeName
	propMap["supported_cluster_architecture_ids"] = cluster.PgType.SupportedClusterArchitectureIds

	return []interface{}{propMap}
}

func getStorageData(cluster *apiv2.ClusterDetail) []interface{} {
	propMap := map[string]interface{}{}
	propMap["iops"] = cluster.Storage.Iops
	propMap["size"] = cluster.Storage.Size
	propMap["throughput"] = cluster.Storage.Throughput
	propMap["volume_properties"] = cluster.Storage.VolumePropertiesId
	propMap["volume_type"] = cluster.Storage.VolumeTypeId

	return []interface{}{propMap}
}
