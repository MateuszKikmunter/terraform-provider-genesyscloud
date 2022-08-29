resource "genesyscloud_journey_action_map" "terraform_test_-TEST-CASE-" {
  display_name          = "terraform_test_-TEST-CASE-_updated"
  trigger_with_segments = [genesyscloud_journey_segment.terraform_test_-TEST-CASE-_action_map_dependency.id]
  activation {
    type = "immediate"
  }
  action {
    media_type = "webMessagingOffer"
    web_messaging_offer_fields {
      offer_text        = "Offer text"
      architect_flow_id = "669c20cd-e7f9-4e31-9556-2d4432b3e02b" # This is a random hardcoded value!
    }
  }
  start_date = "2022-07-04T12:00:00.000000"

  depends_on = [genesyscloud_journey_segment.terraform_test_-TEST-CASE-_action_map_dependency]
}

resource "genesyscloud_journey_segment" "terraform_test_-TEST-CASE-_action_map_dependency" {
  display_name            = "terraform_test_-TEST-CASE-_action_map_dependency"
  color                   = "#008000"
  scope                   = "Customer"
  should_display_to_agent = false
  external_segment {
    id     = "4654654654"
    name   = "external segment name"
    source = "AdobeExperiencePlatform"
  }
}