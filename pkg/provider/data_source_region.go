package provider

import (
	"context"
	"errors"
	"fmt"

	"github.com/EnterpriseDB/terraform-provider-biganimal/pkg/api"
	"github.com/EnterpriseDB/terraform-provider-biganimal/pkg/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// RegionResource is a struct to namespace all the functions
// involved in the Region Resource.  When multiple resources and objects
// are in the same pkg/provider, then it's difficult to namespace things well
type RegionData struct{}

func NewRegionData() *RegionData {
	return &RegionData{}
}

func (r *RegionData) Schema() *schema.Resource {
	return &schema.Resource{
		Description: "The region data source shows the available regions within a cloud provider.",
		ReadContext: r.Read,

		//{
		//	"regionId": "eu-west-1",
		//	"regionName": "EU West 1",
		//	"status": "ACTIVE",
		//	"continent": "Europe"
		//}
		Schema: map[string]*schema.Schema{
			"regions": {
				Description: "Region information.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"region_id": {
							Description: "Region ID of the region.",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"name": {
							Description: "Region name of the region.",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"status": {
							Description: "Region status of the region.",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"continent": {
							Description: "Continent that region belongs to.",
							Type:        schema.TypeString,
							Computed:    true,
						},
					},
				},
			},
			"cloud_provider": {
				Description: "Cloud provider to list the regions. For example, \"aws\" or \"azure\".",
				Type:        schema.TypeString,
				Required:    true,
			},
			"query": {
				Description: "Query to filter region list.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"region_id": {
				Description: "Unique region ID. For example, \"germanywestcentral\" in the Azure cloud provider, \"eu-west-1\" in the AWS cloud provider.",
				Type:        schema.TypeString,
				Optional:    true,
			},
		},
	}
}

func (r *RegionData) Read(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	diags := diag.Diagnostics{}
	client := api.BuildAPI(meta).RegionClient()
	cloud_provider := d.Get("cloud_provider").(string)
	query := d.Get("query").(string)

	id, ok := d.Get("region_id").(string)
	if ok {
		query = id
	}

	regions, err := client.List(ctx, cloud_provider, query)
	if err != nil {
		return FromBigAnimalErr(err)
	}

	if id != "" && len(regions) != 1 {
		return diag.FromErr(errors.New("unable to find a unique region"))
	}

	utils.SetOrPanic(d, "regions", regions)
	d.SetId(fmt.Sprintf("%s/%s", cloud_provider, query))

	return diags
}
