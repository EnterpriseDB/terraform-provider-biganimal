package provider

import (
	"context"
	"time"

	"github.com/EnterpriseDB/terraform-provider-biganimal/pkg/api"
	"github.com/EnterpriseDB/terraform-provider-biganimal/pkg/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type ProjectResource struct{}

func NewProjectResource() *ProjectResource {
	return &ProjectResource{}
}

func (p *ProjectResource) Schema() *schema.Resource {
	return &schema.Resource{
		Description: "The project resource is used to manage projects in your organization. " +
			"See [Managing projects](https://www.enterprisedb.com/docs/biganimal/latest/administering_cluster/projects/) for more details.\n\n" +
			"Newly created projects are not automatically connected to your cloud providers. " +
			"Please visit [Connecting your cloud](https://www.enterprisedb.com/docs/biganimal/latest/getting_started/02_connecting_to_your_cloud/) for more details.",

		CreateContext: p.Create,
		ReadContext:   p.Read,
		UpdateContext: p.Update,
		DeleteContext: p.Delete,

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(60 * time.Minute),
			Update: schema.DefaultTimeout(60 * time.Minute),
			Delete: schema.DefaultTimeout(60 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"project_id": {
				Description: "Project ID of the project.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"project_name": {
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
			// We don't have a mechanism to automate the csp connection right now
			// So, the `cloud_providers` value is computed only.
			"cloud_providers": {
				Description: "Enabled Cloud Providers.",
				Type:        schema.TypeSet,
				Computed:    true,
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
	}
}

func (p *ProjectResource) Create(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := api.BuildAPI(meta).ProjectClient()
	project_name := d.Get("project_name").(string)

	projectId, err := client.Create(ctx, project_name)

	if err != nil {
		return fromBigAnimalErr(err)
	}

	d.SetId(projectId)

	// retry until we get success
	err = resource.RetryContext(
		ctx,
		d.Timeout(schema.TimeoutCreate)-time.Minute,
		p.retryFunc(ctx, d, meta, projectId))

	if err != nil {
		return diag.FromErr(err)
	}
	return diag.Diagnostics{}
}

func (p *ProjectResource) Read(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	if err := p.read(ctx, d, meta); err != nil {
		return fromBigAnimalErr(err)
	}
	return diag.Diagnostics{}
}

func (p *ProjectResource) read(ctx context.Context, d *schema.ResourceData, meta any) error {
	// read the given project
	client := api.BuildAPI(meta).ProjectClient()

	projectId := d.Id()
	project, err := client.Read(ctx, projectId)
	if err != nil {
		return err
	}
	utils.SetOrPanic(d, "project_id", project.ProjectId)
	utils.SetOrPanic(d, "project_name", project.ProjectName)
	utils.SetOrPanic(d, "user_count", project.UserCount)
	utils.SetOrPanic(d, "cluster_count", project.ClusterCount)
	utils.SetOrPanic(d, "cloud_providers", project.CloudProviders)

	d.SetId(project.ProjectId)
	return nil
}

func (p *ProjectResource) Update(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	if d.HasChange("project_name") {
		client := api.BuildAPI(meta).ProjectClient()
		projectId := d.Id()
		projectName := d.Get("project_name").(string)
		_, err := client.Update(ctx, projectId, projectName)
		if err != nil {
			return fromBigAnimalErr(err)
		}
		err = resource.RetryContext(
			ctx,
			d.Timeout(schema.TimeoutCreate)-time.Minute,
			p.retryFunc(ctx, d, meta, projectId))

		if err != nil {
			return diag.FromErr(err)
		}
		return diag.Diagnostics{}
	}
	return nil
}

func (p *ProjectResource) Delete(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	// Delete a project
	client := api.BuildAPI(meta).ProjectClient()
	projectId := d.Id()
	if err := client.Delete(ctx, projectId); err != nil {
		return fromBigAnimalErr(err)
	}
	return diag.Diagnostics{}
}

func (p *ProjectResource) retryFunc(ctx context.Context, d *schema.ResourceData, meta any, projectId string) resource.RetryFunc {
	return func() *resource.RetryError {
		if err := p.read(ctx, d, meta); err != nil {
			return resource.NonRetryableError(err)
		}
		return nil
	}
}
