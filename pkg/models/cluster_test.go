package models

import (
	"context"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"gotest.tools/v3/assert"
)

var testResource = &schema.Resource{
	Description: "Create a Postgres Cluster",

	CreateContext: func(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics { return nil },
	ReadContext:   func(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics { return nil },
	UpdateContext: func(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics { return nil },
	DeleteContext: func(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics { return nil },

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
			Description: "Earliest Backup recover time",
			Type:        schema.TypeString,
			Computed:    true,
		},
		"instance_type": {
			Description: "Instance Type",
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
		"phase": {
			Description: "Current Phase of the cluster.",
			Type:        schema.TypeString,
			Computed:    true,
		},
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

func TestMakeThing(t *testing.T) {
	cr := testResource

	testCases := []struct {
		name    string
		in      []interface{}
		kind    string
		want    any
		wantErr bool
	}{
		{
			in: []interface{}{
				map[string]interface{}{
					"id":    "id",
					"nodes": 1,
				},
			},
			want: Architecture{
				ClusterArchitectureId: "id",
				Nodes:                 1.0,
			},
			wantErr: false,
			kind:    "cluster_architecture",
		},
		{
			in: []interface{}{
				map[string]interface{}{
					"name":  "something",
					"value": "1",
				},
			},
			want: KeyValues{
				{
					Name:  "something",
					Value: "1",
				},
			},
			wantErr: false,
			kind:    "pg_config",
		},
		{
			in: []interface{}{
				map[string]interface{}{
					"iops":              "one",
					"size":              "two",
					"throughput":        "three",
					"volume_properties": "four",
					"volume_type":       "five",
				},
			},
			want: Storage{
				Iops:               "one",
				Size:               "two",
				Throughput:         "three",
				VolumePropertiesId: "four",
				VolumeTypeId:       "five",
			},
			wantErr: false,
			kind:    "storage",
		},
		{
			in: []interface{}{
				map[string]interface{}{
					"cidr_block":  "one",
					"description": "two",
				},
			},
			want: []AllowedIpRange{
				{
					CidrBlock:   "one",
					Description: "two",
				},
			},
			wantErr: false,
			kind:    "allowed_ip_ranges",
		},
	}

	for _, tcase := range testCases {
		t.Logf("testing MakeThing on %s", tcase.kind)
		key := tcase.kind
		config := map[string]interface{}{
			key: tcase.in,
		}

		d := schema.TestResourceDataRaw(t, cr.Schema, config)
		var (
			a   any
			err error
		)

		switch tcase.kind {
		case "cluster_architecture":
			a, err = MakeThing[Architecture](d.Get(key))
		case "pg_config":
			a, err = MakeThing[KeyValues](d.Get(key))
		case "storage":
			a, err = MakeThing[Storage](d.Get(key))
		case "allowed_ip_ranges":
			a, err = MakeThing[[]AllowedIpRange](d.Get(key))
		}

		assert.DeepEqual(t, a, tcase.want)
		assert.Equal(t, err != nil, tcase.wantErr)
	}
}
