package cluster

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/EnterpriseDB/terraform-provider-biganimal/pkg/api"
	"github.com/EnterpriseDB/terraform-provider-biganimal/pkg/apiv2"
	"github.com/EnterpriseDB/terraform-provider-biganimal/pkg/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceCluster() *schema.Resource {
	return &schema.Resource{
		Description: "Create a Postgres Cluster",

		CreateContext: ResourceClusterCreate,
		ReadContext:   ResourceClusterRead,
		UpdateContext: ResourceClusterUpdate,
		DeleteContext: ResourceClusterDelete,

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(45 * time.Minute),
		},

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
							Required:    true,
						},
						"description": {
							Description: "CIDR Block Description",
							Type:        schema.TypeString,
							Optional:    true,
						},
					},
				},
			},
			"backup_retention_period": {
				Description: "Backup Retention Period.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"cluster_architecture": {
				Description: "Cluster Architecture",
				Type:        schema.TypeList,
				Required:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Description: "ID",
							Type:        schema.TypeString,
							Required:    true,
						},
						"name": {
							Description: "Name",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"nodes": {
							Description: "Node Count",
							Type:        schema.TypeInt,
							Required:    true,
						},
					},
				},
			},
			"cluster_name": {
				Description: "Name of the cluster.",
				Type:        schema.TypeString,
				Required:    true,
			},

			// I don't think we need this on the *resource* side.  skip for now
			// "created_at": {
			// 	Description: "Cluster Creation Time",
			// 	Type:        schema.TypeString,
			// 	Computed:    true,
			// },
			// "deleted_at": {
			// 	Description: "Cluster Deletion Time",
			// 	Type:        schema.TypeString,
			// 	Computed:    true,
			// },
			// "expired_at": {
			// 	Description: "Cluster Expiry Time",
			// 	Type:        schema.TypeString,
			// 	Computed:    true,
			// },
			"first_recoverability_point_at": {
				Description: "Cluster Expiry Time",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"instance_type_id": {
				Description: "Cluster Expiry Time",
				Type:        schema.TypeString,
				Required:    true,
			},
			"id": {
				Description: "cluster ID",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"password": {
				Description: "Password",
				Type:        schema.TypeString,
				Required:    true,
				Sensitive:   true,
			},
			"pg_config": {
				Description: "Instance Type",
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Description: "GUC Name",
							Type:        schema.TypeString,
							Required:    true,
						},
						"value": {
							Description: "GUC Value",
							Type:        schema.TypeString,
							Required:    true,
						},
					},
				},
			},
			"pg_type": {
				Description: "Postgres type",
				Type:        schema.TypeString,
				Required:    true,
			},
			"pg_version": {
				Description: "Postgres Version",
				Type:        schema.TypeString,
				Required:    true,
			},
			// "phase": {
			// 	Description: "Current Phase of the cluster.",
			// 	Type:        schema.TypeString,
			// 	Computed:    true,
			// },
			"private_networking": {
				Description: "Is private networking enabled",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"cloud_provider": {
				Description: "Cloud Provider",
				Type:        schema.TypeString,
				Required:    true,
			},
			"region": {
				Description: "Region",
				Type:        schema.TypeString,
				Required:    true,
			},
			"replicas": {
				Description: "Replicas",
				Type:        schema.TypeInt,
				Required:    true,
			},
			// "resizing_pvc": {
			// 	Description: "Resizing PVC",
			// 	Type:        schema.TypeList,
			// 	Elem: &schema.Schema{
			// 		Type: schema.TypeString,
			// 	},
			// 	Computed: true,
			// },
			"storage": {
				Description: "Storage",
				Type:        schema.TypeList,
				Required:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"iops": {
							Description: "IOPS",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"size": {
							Description: "Size",
							Type:        schema.TypeString,
							Required:    true,
						},
						"throughput": {
							Description: "Throughput",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"volume_properties": {
							Description: "Volume Properties",
							Type:        schema.TypeString,
							Required:    true,
						},
						"volume_type": {
							Description: "Volume Type",
							Type:        schema.TypeString,
							Required:    true,
						},
					},
				},
			},
		},
	}
}

func ResourceClusterCreate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := api.BuildAPI(meta, api.ClusterClientType).ClusterClient()
	cluster := apiv2.ClustersBody{}

	cluster.ClusterName = d.Get("cluster_name").(string)
	cluster.BackupRetentionPeriod = utils.StringRef(d.Get("backup_retention_period").(string))
	cluster.AllowedIpRanges = makeAllowedIpRanges(d.Get("allowed_ip_ranges").([]interface{}))

	cluster_architecture := d.Get("cluster_architecture").([]interface{})
	if len(cluster_architecture) != 1 {
		return diag.FromErr(errors.New("require exactly 1 cluster_architecture"))
	}
	cluster.ClusterArchitecture = makeClusterArchitecture(cluster_architecture[0].(map[string]interface{}))
	cluster.InstanceType = &apiv2.ClustersInstanceType{
		InstanceTypeId: d.Get("instance_type_id").(string),
	}
	cluster.Password = d.Get("password").(string)
	cluster.PgConfig = makePgConfig(d.Get("pg_config").([]interface{}))
	cluster.PgType = &apiv2.ClustersPgType{
		PgTypeId: d.Get("pg_type").(string),
	}
	cluster.PgVersion = &apiv2.ClustersPgVersion{
		PgVersionId: d.Get("pg_version").(string),
	}
	cluster.PrivateNetworking = d.Get("private_networking").(bool)
	cluster.Provider = &apiv2.ClustersProvider{
		CloudProviderId: d.Get("cloud_provider").(string),
	}
	cluster.Region = &apiv2.ClustersRegion{
		RegionId: d.Get("region").(string),
	}
	cluster.Replicas = utils.F64Ref(float64((d.Get("replicas").(int))))
	storage := d.Get("storage").([]interface{})
	if len(storage) != 1 {
		return diag.FromErr(errors.New("require exactly 1 storage stanza"))
	}
	cluster.Storage = makeStorage(storage[0].(map[string]interface{}))

	clusterId, err := client.Create(ctx, cluster)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(clusterId)

	// retry until we get success
	err = resource.RetryContext(ctx, d.Timeout(schema.TimeoutCreate)-time.Minute, func() *resource.RetryError {
		clusterId := clusterId
		cluster, err := client.Read(ctx, clusterId)
		if err != nil {
			return resource.NonRetryableError(fmt.Errorf("Error describing instance: %s", err))
		}

		if !client.HasOkCondition(cluster.Conditions) {
			return resource.RetryableError(errors.New("Instance not yet ready"))
		}

		if err := resourceClusterRead(ctx, d, meta); err != nil {
			return resource.NonRetryableError(err)
		}
		return nil
	})
	if err != nil {
		return diag.FromErr(err)
	}
	return diag.Diagnostics{}
}

func ResourceClusterRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	if err := resourceClusterRead(ctx, d, meta); err != nil {
		return diag.FromErr(err)
	}
	return diag.Diagnostics{}
}

func resourceClusterRead(ctx context.Context, d *schema.ResourceData, meta any) error {
	client := api.BuildAPI(meta, api.ClusterClientType).ClusterClient()

	clusterId := d.Id()
	cluster, err := client.Read(ctx, clusterId)
	if err != nil {
		return err
	}

	// set the outputs
	d.Set("backup_retention_period", cluster.BackupRetentionPeriod)
	d.Set("cluster_architecture", getClusterArchitectureData(cluster))
	d.Set("cluster_name", &cluster.ClusterName)

	if cluster.FirstRecoverabilityPointAt != (*apiv2.PointInTime)(nil) {
		d.Set("first_recoverability_point_at", pointInTimeToString(*cluster.FirstRecoverabilityPointAt))
	}

	d.Set("instance_type", getInstanceTypeData(cluster))
	d.Set("id", cluster.ClusterId)
	d.Set("pg_config", cluster.PgType.PgTypeId)
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
	return nil
}

func ResourceClusterUpdate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := api.BuildAPI(meta, api.ClusterClientType).ClusterClient()
	// short circuit early for these types of changes
	if d.HasChange("pg_type") {
		return diag.FromErr(errors.New("pg_type is immutable"))
	}
	if d.HasChange("pg_version") {
		return diag.FromErr(errors.New("pg_version is immutable"))
	}
	if d.HasChange("provider") {
		return diag.FromErr(errors.New("cloud provider is immutable"))
	}
	if d.HasChange("region") {
		return diag.FromErr(errors.New("region is immutable"))
	}
	if d.HasChange("password") {
		return diag.FromErr(errors.New("password is immutable for now"))
	}

	cluster := apiv2.ClustersClusterIdBody{}
	cluster.ClusterName = d.Get("cluster_name").(string)
	cluster.BackupRetentionPeriod = utils.StringRef(d.Get("backup_retention_period").(string))
	cluster.AllowedIpRanges = makeAllowedIpRanges(d.Get("allowed_ip_ranges").([]interface{}))

	cluster_architecture := d.Get("cluster_architecture").([]interface{})
	if len(cluster_architecture) != 1 {
		return diag.FromErr(errors.New("require exactly 1 cluster_architecture"))
	}
	cluster.ClusterArchitecture = makeClusterArchitecture(cluster_architecture[0].(map[string]interface{}))
	cluster.InstanceType = &apiv2.ClustersInstanceType{
		InstanceTypeId: d.Get("instance_type_id").(string),
	}

	// cluster.Password = d.Get("password").(string)

	cluster.PgConfig = makePgConfig(d.Get("pg_config").([]interface{}))

	cluster.PrivateNetworking = d.Get("private_networking").(bool)
	cluster.Replicas = utils.F64Ref(float64((d.Get("replicas").(int))))
	storage := d.Get("storage").([]interface{})
	if len(storage) != 1 {
		return diag.FromErr(errors.New("require exactly 1 storage stanza"))
	}
	cluster.Storage = makeStorage(storage[0].(map[string]interface{}))

	newCluster, err := client.Update(ctx, cluster, d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	// retry until we get success
	err = resource.RetryContext(ctx, d.Timeout(schema.TimeoutCreate)-time.Minute, func() *resource.RetryError {
		clusterId := newCluster.ClusterId
		cluster, err := client.Read(ctx, clusterId)
		if err != nil {
			return resource.NonRetryableError(fmt.Errorf("Error describing instance: %s", err))
		}

		if !client.HasOkCondition(cluster.Conditions) {
			return resource.RetryableError(errors.New("Instance not yet ready"))
		}

		if err := resourceClusterRead(ctx, d, meta); err != nil {
			return resource.NonRetryableError(err)
		}
		return nil
	})
	if err != nil {
		return diag.FromErr(err)
	}
	return diag.Diagnostics{}
}

func ResourceClusterDelete(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := api.BuildAPI(meta, api.ClusterClientType).ClusterClient()
	clusterId := d.Id()
	if err := client.Delete(ctx, clusterId); err != nil {
		return diag.FromErr(err)
	}
	return diag.Diagnostics{}
}

func makeClusterArchitecture(blob map[string]interface{}) *apiv2.ClustersClusterArchitecture {
	return &apiv2.ClustersClusterArchitecture{
		ClusterArchitectureId: blob["id"].(string),
		Nodes:                 float64(blob["nodes"].(int)),
	}
}

func makeAllowedIpRanges(list []interface{}) []apiv2.AllowedIpRange {
	data := []apiv2.AllowedIpRange{}

	for _, v := range list {
		blob := v.(map[string]interface{})
		data = append(data, apiv2.AllowedIpRange{
			CidrBlock:   blob["cidr_block"].(string),
			Description: blob["description"].(string),
		})
	}
	return data
}

func makePgConfig(list []interface{}) []apiv2.ClustersClusterArchitectureParams {
	data := []apiv2.ClustersClusterArchitectureParams{}

	for _, v := range list {
		blob := v.(map[string]interface{})
		data = append(data, apiv2.ClustersClusterArchitectureParams{
			Name:  utils.StringRef(blob["name"].(string)),
			Value: utils.StringRef(blob["value"].(string)),
		})
	}
	return data
}

func makeStorage(blob map[string]interface{}) *apiv2.ClustersStorage {
	storage := &apiv2.ClustersStorage{
		Size:               utils.StringRef(blob["size"].(string)),
		VolumePropertiesId: blob["volume_properties"].(string),
		VolumeTypeId:       blob["volume_type"].(string),
	}

	if iops, ok := blob["iops"].(string); ok {
		storage.Iops = utils.StringRef(iops)
	} else {
		storage.Iops = utils.StringRef("")
	}

	if throughput, ok := blob["throughput"].(string); ok {
		storage.Throughput = utils.StringRef(throughput)
	} else {
		storage.Throughput = utils.StringRef("")
	}

	return storage
}

// {
// 	"clusterName": "My Cluster",
// 	"password": "perfectenschlag",
// 	"privateNetworking": false,
// 	"replicas": 1,
// 	"allowedIpRanges": [
// 	  {
// 		"cidrBlock": "252.1.1.1/24",
// 		"description": "New York"
// 	  },
// 	  {
// 		"cidrBlock": "167.3.2.1/32",
// 		"description": "Boston"
// 	  }
// 	],
// 	"pgConfig": [
// 	  {
// 		"name": "application_name",
// 		"value": "restore-test"
// 	  },
// 	  {
// 		"name": "array_nulls",
// 		"value": "off"
// 	  }
// 	],
// 	"pgType": {
// 	  "pgTypeId": "epas"
// 	},
// 	"pgVersion": {
// 	  "pgVersionId": "13"
// 	},
// 	"provider": {
// 	  "cloudProviderId": "azure"
// 	},
// 	"readOnlyConnections": false,
// 	"region": {
// 	  "regionId": "australiaeast"
// 	},
// 	"instanceType": {
// 	  "instanceTypeId": "azure:Standard_D16s_v3"
// 	},
// 	"storage": {
// 	  "volumeTypeId": "azurepremiumstorage",
// 	  "volumePropertiesId": "P1",
// 	  "iops": "120",
// 	  "size": "4 Gi"
// 	},
// 	"clusterArchitecture": {
// 	  "clusterArchitectureId": "single",
// 	  "nodes": 1
// 	},
// 	"backupRetentionPeriod": "6d"
//   }
