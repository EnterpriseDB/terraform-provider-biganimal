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

// RegionResource is a struct to namespace all the functions
// involved in the Region Resource.  When multiple resources and objects
// are in the same pkg/provider, then it's difficult to namespace things well
type RegionResource struct{}

func NewRegionResource() *RegionResource {
	return &RegionResource{}
}

func (r *RegionResource) Schema() *schema.Resource {
	return &schema.Resource{
		Description: "Manage a region",

		CreateContext: r.Create,
		ReadContext:   r.Read,
		UpdateContext: r.Update,
		DeleteContext: r.Delete,

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

func (r *RegionResource) Create(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	return r.Update(ctx, d, meta)
}

func (r *RegionResource) Read(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	if err := r.read(ctx, d, meta); err != nil {
		return diag.FromErr(err)
	}
	return diag.Diagnostics{}
}

func (r *RegionResource) read(ctx context.Context, d *schema.ResourceData, meta any) error {
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

func (r *RegionResource) Update(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
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
		r.retryFunc(ctx, d, meta, cloudProvider, id, desiredState))
	if err != nil {
		return diag.FromErr(err)
	}
	return diag.Diagnostics{}
}

func (r *RegionResource) Delete(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
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
		r.retryFunc(ctx, d, meta, cloudProvider, id, desiredState))
	if err != nil {
		return diag.FromErr(err)
	}

	return diag.Diagnostics{}
}

func (r *RegionResource) retryFunc(ctx context.Context, d *schema.ResourceData, meta any, cloudProvider, regionId, desiredState string) resource.RetryFunc {
	client := api.BuildAPI(meta).RegionClient()
	return func() *resource.RetryError {
		region, err := client.Read(ctx, cloudProvider, regionId)
		if err != nil {
			return resource.NonRetryableError(fmt.Errorf("Error describing instance: %s", err))
		}

		if region.Status != desiredState {
			return resource.RetryableError(errors.New("Operation incomplete"))
		}

		if err := r.read(ctx, d, meta); err != nil {
			return resource.NonRetryableError(err)
		}

		return nil
	}
}
