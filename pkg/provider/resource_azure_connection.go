package provider

import (
	"context"
	"github.com/EnterpriseDB/terraform-provider-biganimal/pkg/api"
	"github.com/EnterpriseDB/terraform-provider-biganimal/pkg/models"
	"github.com/EnterpriseDB/terraform-provider-biganimal/pkg/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"time"
)

type AzureConnectionResource struct{}

func NewAzureConnectionResource() *AzureConnectionResource {
	return &AzureConnectionResource{}
}

func (a *AzureConnectionResource) Schema() *schema.Resource {
	return &schema.Resource{
		Description: `The azure_connection resource is used to make connection between your project and Azure. 
To obtain the necessary input parameters, please refer to [Connect CSP](https://www.enterprisedb.com/docs/biganimal/latest/getting_started/02_connecting_to_your_cloud/connecting_azure/).`,

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
				Description: "Project ID of the project.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"tenant_id": {
				Description: "Your Azure tenant identity.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"subscription_id": {
				Description: "Your Azure subscription identity.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"client_id": {
				Description: "Azure APP client identity for BigAnimal.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"client_secret": {
				Description: "Azure APP client secret for BigAnimal.",
				Type:        schema.TypeString,
				Required:    true,
			},
		},
	}
}

func (a *AzureConnectionResource) Create(ctx context.Context, data *schema.ResourceData, meta any) diag.Diagnostics {
	projectID := data.Get("project_id").(string)
	clientID := data.Get("client_id").(string)
	clientSecret := data.Get("client_secret").(string)
	subscriptionID := data.Get("subscription_id").(string)
	tenantID := data.Get("tenant_id").(string)

	client := api.BuildAPI(meta).ProviderClient()
	model := models.AzureConnection{
		ClientId:       clientID,
		ClientSecret:   clientSecret,
		SubscriptionId: subscriptionID,
		TenantId:       tenantID,
	}

	if err := client.RegisterAzure(ctx, projectID, model); err != nil {
		return fromBigAnimalErr(err)
	}
	// azure_connection is unique for one project
	data.SetId(projectID)

	utils.SetOrPanic(data, "client_id", clientID)
	utils.SetOrPanic(data, "client_secret", clientSecret)
	utils.SetOrPanic(data, "subscription_id", subscriptionID)
	utils.SetOrPanic(data, "tenant_id", tenantID)
	return diag.Diagnostics{}
}

func (a *AzureConnectionResource) Read(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {
	return unsupportedWarning("azure_connection can't be read")
}

func (a *AzureConnectionResource) Update(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {
	return unsupportedWarning("azure_connection can't be updated")
}

func (a *AzureConnectionResource) Delete(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {
	return unsupportedWarning("azure_connection can't be deleted")
}
