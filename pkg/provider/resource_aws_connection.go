package provider

import (
	"context"
	"fmt"
	"time"

	"github.com/EnterpriseDB/terraform-provider-biganimal/pkg/api"
	"github.com/EnterpriseDB/terraform-provider-biganimal/pkg/models"
	"github.com/EnterpriseDB/terraform-provider-biganimal/pkg/utils"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type AWSConnectionResource struct{}

func NewAWSConnectionResource() *AWSConnectionResource {
	return &AWSConnectionResource{}
}

func (a *AWSConnectionResource) Schema() *schema.Resource {
	return &schema.Resource{
		Description: `The aws_connection resource is used to make connection between your project and AWS.
o obtain the necessary input parameters, please refer to [Connect CSP](https://www.enterprisedb.com/docs/biganimal/latest/getting_started/02_connecting_to_your_cloud/connecting_aws/).`,

		CreateContext: a.Create,
		ReadContext:   a.Read,
		UpdateContext: a.Update,
		DeleteContext: a.Delete,

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(60 * time.Minute),
			Update: schema.DefaultTimeout(60 * time.Minute),
			Delete: schema.DefaultTimeout(60 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"project_id": {
				Description:      "Project ID of the project.",
				Type:             schema.TypeString,
				Required:         true,
				ValidateDiagFunc: validateProjectId,
			},
			"external_id": {
				Description: "The AWS external ID provided by BigAnimal.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"role_arn": {
				Description:      "The AWS IAM role used by BigAnimal.",
				Type:             schema.TypeString,
				Required:         true,
				ValidateDiagFunc: validateARN,
			},
		},
	}
}

func (a *AWSConnectionResource) Create(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	return a.create(ctx, d, meta)
}

func (a *AWSConnectionResource) create(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	projectID := d.Get("project_id").(string)
	externalID := d.Get("external_id").(string)
	roleArn := d.Get("role_arn").(string)
	client := api.BuildAPI(meta).CloudProviderClient()
	var model models.AWSConnection

	model.RoleArn = roleArn
	model.ExternalID = externalID

	tflog.Debug(ctx, fmt.Sprintf("role_arn: %s, external_id: %s", roleArn, externalID))

	if err := client.CreateAWSConnection(ctx, projectID, model); err != nil {
		return fromBigAnimalErr(err)
	}

	// azure_connection is unique for one project
	d.SetId(projectID)

	if err := a.read(ctx, d, meta); err != nil {
		return diag.FromErr(err)
	}
	return diag.Diagnostics{}
}

func (a *AWSConnectionResource) Read(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	if err := a.read(ctx, d, meta); err != nil {
		return fromBigAnimalErr(err)
	}
	return diag.Diagnostics{}
}

func (a *AWSConnectionResource) read(ctx context.Context, d *schema.ResourceData, meta any) error {
	client := api.BuildAPI(meta).CloudProviderClient()

	projectId := d.Id()
	conn, err := client.GetAWSConnection(ctx, projectId)
	if err != nil {
		return err
	}

	utils.SetOrPanic(d, "external_id", conn.ExternalID)
	utils.SetOrPanic(d, "role_arn", conn.RoleArn)
	return nil
}

func (a *AWSConnectionResource) Update(ctx context.Context, data *schema.ResourceData, meta any) diag.Diagnostics {
	return unsupportedWarning("aws_connection can't be updated")
}

func (a *AWSConnectionResource) Delete(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {
	return unsupportedWarning("aws_connection can't be deleted")
}
