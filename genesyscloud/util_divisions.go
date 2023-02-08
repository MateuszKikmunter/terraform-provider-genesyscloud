package genesyscloud

import (
	"log"
	"sync"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/mypurecloud/platform-client-sdk-go/v92/platformclientv2"
)

type jsonMap map[string]interface{}

// Attempt to get the home division once during a provider run
var divOnce sync.Once
var homeDivID string
var homeDivErr diag.Diagnostics

func getHomeDivisionID() (string, diag.Diagnostics) {
	divOnce.Do(func() {
		authAPI := platformclientv2.NewAuthorizationApi()
		homeDiv, _, err := authAPI.GetAuthorizationDivisionsHome()
		if err != nil {
			homeDivErr = diag.Errorf("Failed to query home division: %s", err)
			return
		}
		homeDivID = *homeDiv.Id
	})

	if homeDivErr != nil {
		return "", homeDivErr
	}
	return homeDivID, nil
}

func updateObjectDivision(d *schema.ResourceData, objType string, sdkConfig *platformclientv2.Configuration) diag.Diagnostics {
	if d.HasChange("division_id") {
		authAPI := platformclientv2.NewAuthorizationApiWithConfig(sdkConfig)
		divisionID := d.Get("division_id").(string)
		if divisionID == "" {
			// Default to home division
			homeDivision, diagErr := getHomeDivisionID()
			if diagErr != nil {
				return diagErr
			}
			divisionID = homeDivision
		}
		log.Printf("Updating division for %s %s to %s", objType, d.Id(), divisionID)
		_, divErr := authAPI.PostAuthorizationDivisionObject(divisionID, objType, []string{d.Id()})
		if divErr != nil {
			return diag.Errorf("Failed to update division for %s %s: %s", objType, d.Id(), divErr)
		}
	}
	return nil
}
