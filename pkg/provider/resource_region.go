package provider

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/EnterpriseDB/terraform-provider-biganimal/pkg/api"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceRegion() *schema.Resource {
	return &schema.Resource{
		Description: "Manage a region",

		CreateContext: ResourceRegionCreate,
		ReadContext:   ResourceRegionRead,
		UpdateContext: ResourceRegionUpdate,
		DeleteContext: ResourceRegionDelete,

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(60 * time.Minute),
			Update: schema.DefaultTimeout(60 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"cloud_provider": {
				Description: "Cloud Provider",
				Type:        schema.TypeString,
				Required:    true,
			},
			"region_id": {
				Description: "Region ID",
				Type:        schema.TypeString,
				Required:    true,
			},
			"name": {
				Description: "Region Name",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"status": {
				Description: "Region Status",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     api.REGION_ACTIVE,
			},
			"continent": {
				Description: "Continent",
				Type:        schema.TypeString,
				Computed:    true,
			},
		},
	}
}

func ResourceRegionCreate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	return ResourceRegionUpdate(ctx, d, meta)
}

func ResourceRegionRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	if err := resourceRegionRead(ctx, d, meta); err != nil {
		return diag.FromErr(err)
	}
	return diag.Diagnostics{}
}

func resourceRegionRead(ctx context.Context, d *schema.ResourceData, meta any) error {
	client := api.BuildAPI(meta).RegionClient()
	cloud_provider := d.Get("cloud_provider").(string)

	id := d.Get("region_id").(string)
	region, err := client.Read(ctx, cloud_provider, id)
	if err != nil {
		return err
	}

	d.Set("name", region.Name)
	d.Set("status", region.Status)
	d.Set("continent", region.Continent)

	d.SetId(fmt.Sprintf("%s/%s", cloud_provider, id))

	return nil
}

func ResourceRegionUpdate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := api.BuildAPI(meta).RegionClient()

	cloudProvider := d.Get("cloud_provider").(string)
	id := d.Get("region_id").(string)
	desiredState := d.Get("status").(string)

	region, err := client.Read(ctx, cloudProvider, id)
	if err != nil {
		return diag.FromErr(err)
	}

	d.Set("name", region.Name)
	d.Set("continent", region.Continent)
	d.Set("status", desiredState)
	d.SetId(fmt.Sprintf("%s/%s", cloudProvider, id))

	if desiredState == region.Status { // no change, exit early
		return diag.Diagnostics{}
	}

	tflog.Debug(ctx, fmt.Sprintf("updating region from %s to %s", region.Status, desiredState))
	if err = client.Update(ctx, desiredState, cloudProvider, id); err != nil {
		return diag.FromErr(err)
	}

	// retry until we get success
	err = resource.RetryContext(
		ctx,
		d.Timeout(schema.TimeoutCreate)-time.Minute,
		retryRegionFunc(ctx, d, meta, cloudProvider, id, desiredState))
	if err != nil {
		return diag.FromErr(err)
	}
	return diag.Diagnostics{}
}

func ResourceRegionDelete(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := api.BuildAPI(meta).RegionClient()

	cloudProvider := d.Get("cloud_provider").(string)
	id := d.Get("region_id").(string)
	desiredState := api.REGION_INACTIVE
	if err := client.Update(ctx, api.REGION_INACTIVE, cloudProvider, id); err != nil {
		return diag.FromErr(err)
	}

	// retry until we get success
	err := resource.RetryContext(
		ctx,
		d.Timeout(schema.TimeoutDelete)-time.Minute,
		retryRegionFunc(ctx, d, meta, cloudProvider, id, desiredState))
	if err != nil {
		return diag.FromErr(err)
	}

	return diag.Diagnostics{}
}

func retryRegionFunc(ctx context.Context, d *schema.ResourceData, meta any, cloudProvider, regionId, desiredState string) resource.RetryFunc {
	client := api.BuildAPI(meta).RegionClient()
	return func() *resource.RetryError {
		region, err := client.Read(ctx, cloudProvider, regionId)
		if err != nil {
			return resource.NonRetryableError(fmt.Errorf("Error describing instance: %s", err))
		}

		if region.Status != desiredState {
			return resource.RetryableError(errors.New("Operation incomplete"))
		}

		if err := resourceRegionRead(ctx, d, meta); err != nil {
			return resource.NonRetryableError(err)
		}

		return nil
	}
}
