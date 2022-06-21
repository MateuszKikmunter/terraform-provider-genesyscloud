package genesyscloud

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/mypurecloud/platform-client-sdk-go/v72/platformclientv2"
)

type journeySegmentStruct struct {
	resourceID  string
	displayName string
	color       string
	scope       string
	context     string
	journey     string
}

type contextStruct struct {
	key              string
	values           string
	operator         string
	shouldIgnoreCase bool
	entityType       string
}

func TestAccResourceJourneySegmentBasic(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	resourcePrefix := "genesyscloud_journey_segment."
	journeySegmentIdPrefix := "terraform_test_"
	journeySegmentId := journeySegmentIdPrefix + strconv.Itoa(rand.Intn(1000))
	displayName1 := journeySegmentId
	displayName2 := journeySegmentId + "_updated"
	color1 := "#008000"
	color2 := "#308000"
	scope1 := "Customer"
	//scope2 := "Session"
	contextPatternCriteriaKey1 := "geolocation.postalCode"
	contextPatternCriteriaValues1 := "something"
	contextPatternCriteriaOperator1 := "equal"
	contextPatternCriteriaShouldIgnoreCase1 := true
	contextPatternCriteriaEntityType1 := "visit"
	journey :=
		`journey {
			patterns {
				criteria {
					key = "page.hostname"
					values = ["something_else", "more"]
					operator = "equal"
					should_ignore_case = false
				}
				count = 1
				stream_type = "Web"
				session_type = "web"
			} 
		}`

	err := authorizeSdk()
	if err != nil {
		t.Fatal(err)
	}

	cleanupJourneySegments(journeySegmentIdPrefix)

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				// Create
				Config: generateJourneySegmentResource(&journeySegmentStruct{
					journeySegmentId,
					displayName1,
					color1,
					scope1,
					generateContext(&contextStruct{
						contextPatternCriteriaKey1,
						contextPatternCriteriaValues1,
						contextPatternCriteriaOperator1,
						contextPatternCriteriaShouldIgnoreCase1,
						contextPatternCriteriaEntityType1,
					}),
					journey,
				}),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourcePrefix+journeySegmentId, "display_name", displayName1),
					resource.TestCheckResourceAttr(resourcePrefix+journeySegmentId, "color", color1),
					resource.TestCheckResourceAttr(resourcePrefix+journeySegmentId, "scope", scope1),
				),
			},
			{
				// Update
				Config: generateJourneySegmentResource(&journeySegmentStruct{
					journeySegmentId,
					displayName2,
					color2,
					scope1,
					generateContext(&contextStruct{
						contextPatternCriteriaKey1,
						contextPatternCriteriaValues1,
						contextPatternCriteriaOperator1,
						contextPatternCriteriaShouldIgnoreCase1,
						contextPatternCriteriaEntityType1,
					}),
					journey,
				}),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourcePrefix+journeySegmentId, "display_name", displayName2),
					resource.TestCheckResourceAttr(resourcePrefix+journeySegmentId, "color", color2),
					resource.TestCheckResourceAttr(resourcePrefix+journeySegmentId, "scope", scope1),
				),
			},
			{
				// Import/Read
				ResourceName:      resourcePrefix + journeySegmentId,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
		CheckDestroy: testVerifyJourneySegmentsDestroyed,
	})
}

func cleanupJourneySegments(journeySegmentIdPrefix string) {
	journeyApi := platformclientv2.NewJourneyApiWithConfig(sdkConfig)

	pageCount := 1 // Needed because of broken journey common paging
	for pageNum := 1; pageNum <= pageCount; pageNum++ {
		const pageSize = 100
		journeySegments, _, getErr := journeyApi.GetJourneySegments("", pageSize, pageNum, true, nil, nil, "")
		if getErr != nil {
			return
		}

		if journeySegments.Entities == nil || len(*journeySegments.Entities) == 0 {
			break
		}

		for _, journeySegment := range *journeySegments.Entities {
			if journeySegment.DisplayName != nil && strings.HasPrefix(*journeySegment.DisplayName, journeySegmentIdPrefix) {
				_, delErr := journeyApi.DeleteJourneySegment(*journeySegment.Id)
				if delErr != nil {
					diag.Errorf("failed to delete journey segment %s (%s): %s", *journeySegment.Id, *journeySegment.DisplayName, delErr)
					return
				}
				log.Printf("Deleted journey segment %s (%s)", *journeySegment.Id, *journeySegment.DisplayName)
			}
		}

		pageCount = *journeySegments.PageCount
	}
}

func generateJourneySegmentResource(journeySegment *journeySegmentStruct) string {
	return fmt.Sprintf(`resource "genesyscloud_journey_segment" "%s" {
		display_name = "%s"
		color = "%s"
		scope = "%s"
		%s
		%s
	}`, journeySegment.resourceID,
		journeySegment.displayName,
		journeySegment.color,
		journeySegment.scope,
		journeySegment.context,
		journeySegment.scope)
}

func generateContext(context *contextStruct) string {
	return fmt.Sprintf(`context {
		patterns {
			criteria {
				key = "%s"
				values = ["%s"]
				operator = "%s"
				should_ignore_case = %t
				entity_type = "%s"
			}
		}
	}`, context.key,
		context.values,
		context.operator,
		context.shouldIgnoreCase,
		context.entityType,
	)
}

func testVerifyJourneySegmentsDestroyed(state *terraform.State) error {
	journeyApi := platformclientv2.NewJourneyApiWithConfig(sdkConfig)
	for _, rs := range state.RootModule().Resources {
		if rs.Type != "genesyscloud_journey_segment" {
			continue
		}

		journeySegment, resp, err := journeyApi.GetJourneySegment(rs.Primary.ID)
		if journeySegment != nil {
			return fmt.Errorf("journey segment (%s) still exists", rs.Primary.ID)
		}

		if isStatus404(resp) {
			// Journey segment not found as expected
			continue
		}

		// Unexpected error
		return fmt.Errorf("unexpected error: %s", err)
	}
	// Success. All Journey segment destroyed
	return nil
}
