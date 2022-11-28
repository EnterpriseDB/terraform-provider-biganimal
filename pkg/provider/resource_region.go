package provider

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/EnterpriseDB/terraform-provider-biganimal/pkg/api"
	"github.com/EnterpriseDB/terraform-provider-biganimal/pkg/utils"
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
		Description: "The region resource is used to manage regions for a given cloud provider. See [Activating regions](https://www.enterprisedb.com/docs/biganimal/latest/getting_started/activating_regions/) for more details.",

		CreateContext: r.Create,
		ReadContext:   r.Read,
		UpdateContext: r.Update,
		DeleteContext: r.Delete,

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(60 * time.Minute),
			Update: schema.DefaultTimeout(60 * time.Minute),
			Delete: schema.DefaultTimeout(60 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"cloud_provider": {
				Description: "Cloud Provider. For example, \"aws\" or \"azure\".",
				Type:        schema.TypeString,
				Required:    true,
			},
			"region_id": {
				Description: "Region ID of the region. For example, \"germanywestcentral\" in the Azure cloud provider or \"eu-west-1\" in the AWS cloud provider.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"name": {
				Description: "Region Name of the region. For example, \"Germany West Central\" or \"EU West 1\".",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"status": {
				Description: "Region Status of the region. For example, \"ACTIVE\", \"INACTIVE\", or \"SUSPENDED\".",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     api.REGION_ACTIVE,
			},
			"continent": {
				Description: "Continent that region belongs to. For example, \"Asia\", \"Australia\", or \"Europe\".",
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

	utils.SetOrPanic(d, "name", region.Name)
	utils.SetOrPanic(d, "status", region.Status)
	utils.SetOrPanic(d, "continent", region.Continent)
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

	utils.SetOrPanic(d, "name", region.Name)
	utils.SetOrPanic(d, "continent", region.Continent)
	utils.SetOrPanic(d, "status", desiredState)
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
