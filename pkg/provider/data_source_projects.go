package provider

import (
	"context"
	"strconv"
	"time"

	"github.com/EnterpriseDB/terraform-provider-biganimal/pkg/api"
	"github.com/EnterpriseDB/terraform-provider-biganimal/pkg/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type ProjectsData struct{}

func NewProjectsData() *ProjectsData {
	return &ProjectsData{}
}

func (p *ProjectsData) Schema() *schema.Resource {
	return &schema.Resource{
		Description: "The projects data source shows the BigAnimal Projects.",
		ReadContext: p.Read,

		Schema: map[string]*schema.Schema{
			"projects": {
				Description: "List of the organization's projects.",
				Type:        schema.TypeSet,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"project_id": {
							Description: "Project ID of the project.",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"name": {
							Description: "Project Name of the project.",
							Type:        schema.TypeString,
							Required:    true,
						},
						"user_count": {
							Description: "User Count of the project.",
							Type:        schema.TypeInt,
							Computed:    true,
						},
						"cluster_count": {
							Description: "User Count of the project.",
							Type:        schema.TypeInt,
							Computed:    true,
						},
						"cloud_providers": {
							Description: "Enabled Cloud Providers.",
							Type:        schema.TypeSet,
							Computed:    true,
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"cloud_provider_id": {
										Description: "Cloud Provider ID.",
										Type:        schema.TypeString,
										Computed:    true,
									},
									"cloud_provider_name": {
										Description: "Cloud Provider Name.",
										Type:        schema.TypeString,
										Computed:    true,
									},
								},
							},
						},
					},
				},
			},
			"query": {
				Description: "Query to filter project list.",
				Type:        schema.TypeString,
				Optional:    true,
			},
		},
	}
}

func (p *ProjectsData) Read(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	diags := diag.Diagnostics{}
	client := api.BuildAPI(meta).ProjectClient()

	query := d.Get("query").(string)

	projects, err := client.List(ctx, query)
	if err != nil {
		return fromBigAnimalErr(err)
	}

	utils.SetOrPanic(d, "projects", projects)
	// FIXME: Check if there is a better way to set the ID.
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}
