package provider

import (
	"context"
	"errors"
	"fmt"

	"github.com/EnterpriseDB/terraform-provider-biganimal/pkg/api"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceRegion() *schema.Resource {
	return &schema.Resource{
		Description: "The available regions within a cloud provider",
		ReadContext: dataSourceRegionRead,

		// {
		// 	"regionId": "azure:Canada East",
		// 	"regionName": "Canada East",
		// 	"status": "Active",
		// 	"continent": "Americas"
		//   }
		Schema: map[string]*schema.Schema{
			"regions": {
				Description: "Region ID",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"region_id": {
							Description: "Region ID",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"name": {
							Description: "Region Name",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"status": {
							Description: "Region Status",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"continent": {
							Description: "Continent",
							Type:        schema.TypeString,
							Computed:    true,
						},
					},
				},
			},
			"cloud_provider": {
				Description: "Cloud Provider",
				Type:        schema.TypeString,
				Required:    true,
			},
			"query": {
				Description: "Query to filter region list",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"region_id": {
				Description: "Unique Region ID",
				Type:        schema.TypeString,
				Optional:    true,
			},
		},
	}
}

func dataSourceRegionRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
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
		return diag.FromErr(err)
	}

	if id != "" && len(regions) != 1 {
		return diag.FromErr(errors.New("unable to find a unique region"))
	}

	d.Set("regions", getRegionList(regions))
	d.SetId(fmt.Sprintf("%s/%s", cloud_provider, query))

	return diags
}

func getRegionList(regions []*api.Region) []interface{} {
	list := []interface{}{}

	for _, region := range regions {
		propMap := map[string]interface{}{}
		propMap["region_id"] = region.Id
		propMap["name"] = region.Name
		propMap["status"] = region.Status
		propMap["continent"] = region.Continent
		list = append(list, propMap)
	}

	return list
}
