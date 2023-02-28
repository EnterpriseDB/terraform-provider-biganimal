package provider

import (
	"context"
	"github.com/EnterpriseDB/terraform-provider-biganimal/pkg/api"
	"github.com/EnterpriseDB/terraform-provider-biganimal/pkg/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type AWSConnectionData struct{}

func NewAWSConnectionData() *AWSConnectionData {
	return &AWSConnectionData{}
}

func (d *AWSConnectionData) Schema() *schema.Resource {
	return &schema.Resource{
		Description: "The aws_connection data source shows the BigAnimal AWS Connection.",
		ReadContext: d.Read,
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
				Computed:    true,
			},
			"role_arn": {
				Description: "the AWS iam role used by BigAnimal.",
				Type:        schema.TypeString,
				Computed:    true,
			},
		},
	}
}

func (d *AWSConnectionData) Read(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := api.BuildAPI(meta).ProviderClient()

	projectID := data.Get("project_id").(string)

	conn, err := client.GetAWSConnection(ctx, projectID)
	if err != nil {
		return fromBigAnimalErr(err)
	}

	utils.SetOrPanic(data, "external_id", conn.ExternalID)
	utils.SetOrPanic(data, "role_arn", conn.RoleArn)
	data.SetId(projectID)
	return diag.Diagnostics{}
}
