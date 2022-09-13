package provider

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	//"fmt"
	"io"
	"net/http"

	"os"

	baapi "github.com/EnterpriseDB/upm-cli/generated/api"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/kr/pretty"
)

func dataSourceScaffolding() *schema.Resource {
	return &schema.Resource{
		// This description is used by the documentation generator and the language server.
		Description: "Sample data source in the Terraform provider scaffolding.",

		ReadContext: dataSourceScaffoldingRead,

		Schema: map[string]*schema.Schema{
			"name": {
				// This description is used by the documentation generator and the language server.
				Description: "Name of the cluster.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"current_primary": {
				// This description is used by the documentation generator and the language server.
				Description: "CurPiri.",
				Type:        schema.TypeString,
				Required:    false,
				Computed:    true,
			},
			"phase": {
				// This description is used by the documentation generator and the language server.
				Description: "Current Phase of the cluster.",
				Type:        schema.TypeString,
				Required:    false,
				Computed:    true,
			},
		},
	}
}

type Clusters struct {
	Data []baapi.Cluster `json:"data"`
}

func dataSourceScaffoldingRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	// use the meta value to retrieve your client from the provider configure method
	// client := meta.(*apiClient)

	//endpoint := "/clusters"
	//client := &http.Client{Timeout: 10 * time.Second}

	var diags diag.Diagnostics

	cluster_name := d.Get("name")

	idFromAPI := "my-id"
	d.SetId(idFromAPI)

	//token := os.Getenv("BA_BEARER_TOKEN")
	baURL := os.Getenv("BA_API_URI")
	token := os.Getenv("BA_BEARER_TOKEN")

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/clusters?name=%s", baURL, cluster_name), nil)

	if err != nil {
		return diag.FromErr(err)
	}

	req.Header.Add("authorization", "Bearer "+token)
	req.Header.Add("content-type", "application/json")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	var clusters Clusters

	if err = json.Unmarshal(body, &clusters); err != nil {
		return diag.FromErr(err)
	}

	if len(clusters.Data) != 1 {
		return diag.FromErr(errors.New("some bullshit here"))
	}

	d.Set("CurrentPrimary", clusters.Data[0].CurrentPrimary)

	d.Set("phase", clusters.Data[0].Phase)
	//fmt.Println(string(result))
	pretty.Println(clusters)



	/* 	var diags diag.Diagnostics

		diags = append(diags, diag.Diagnostic{
	        Severity: diag.Error,
	        Summary:  "Unable to create HashiCups client",
	        Detail:   "Unable to auth user for authenticated HashiCups client",
	      })
	*/
	//return diag.Errorf("not implemented here")
	return diags
}
