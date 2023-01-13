package genesyscloud

import (
	"context"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/mypurecloud/platform-client-sdk-go/v89/platformclientv2"
)

func dataSourceResponseManagamentResponseAsset() *schema.Resource {
	return &schema.Resource{
		Description: "Data source for Genesys Cloud Response Management Response Assets. Select a response asset by name.",
		ReadContext: readWithPooledClient(dataSourceResponseManagamentResponseAssetRead),
		Schema: map[string]*schema.Schema{
			"name": {
				Description: "Response asset name.",
				Type:        schema.TypeString,
				Required:    true,
			},
		},
	}
}

func dataSourceResponseManagamentResponseAssetRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var (
		name    = d.Get("name").(string)
		field   = "name"
		fields  = []string{field}
		varType = "TERM"
		filter  = platformclientv2.Responseassetfilter{
			Fields:  &fields,
			Value:   &name,
			VarType: &varType,
		}
		body = platformclientv2.Responseassetsearchrequest{
			Query:  &[]platformclientv2.Responseassetfilter{filter},
			SortBy: &field,
		}
	)

	sdkConfig := m.(*providerMeta).ClientConfig
	respManagementApi := platformclientv2.NewResponseManagementApiWithConfig(sdkConfig)

	return withRetries(ctx, 15*time.Second, func() *resource.RetryError {
		responseData, _, getErr := respManagementApi.PostResponsemanagementResponseassetsSearch(body, nil)
		if getErr != nil {
			return resource.NonRetryableError(fmt.Errorf("Error requesting response asset %s: %s", name, getErr))
		}
		if responseData.Results == nil || len(*responseData.Results) == 0 {
			return resource.RetryableError(fmt.Errorf("No response asset found with name %s", name))
		}
		asset := (*responseData.Results)[0]
		d.SetId(*asset.Id)
		return nil
	})
}
