---
page_title: "genesyscloud_journey_segment Resource - terraform-provider-genesyscloud"
subcategory: ""
description: |-
Genesys Cloud Journey Segment
---
# genesyscloud_journey_segment (Resource)

Genesys Cloud Journey Segment

## API Usage
The following Genesys Cloud APIs are used by this resource. Ensure your OAuth Client has been granted the necessary scopes and permissions to perform these operations:

* [GET /api/v2/journey/segments](https://developer.genesys.cloud/commdigital/digital/webmessaging/journey/journey-apis#get-api-v2-journey-segments)
* [POST /api/v2/journey/segments](https://developer.genesys.cloud/commdigital/digital/webmessaging/journey/journey-apis#post-api-v2-journey-segments)
* [GET /api/v2/journey/segments/{segmentId}](https://developer.genesys.cloud/commdigital/digital/webmessaging/journey/journey-apis#get-api-v2-journey-segments--segmentId-)
* [PATCH /api/v2/journey/segments/{segmentId}](https://developer.genesys.cloud/commdigital/digital/webmessaging/journey/journey-apis#patch-api-v2-journey-segments--segmentId-)
* [DELETE /api/v2/journey/segments/{segmentId}](https://developer.genesys.cloud/commdigital/digital/webmessaging/journey/journey-apis#delete-api-v2-journey-segments--segmentId-)

## Example Usage

```terraform
resource "genesyscloud_journey_segment" "test_journey_segment" {
  display_name = "journey_segment_1"
  description = "Description of Journey Segment"
  color = "008000"
  scope = "Customer"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `display_name` (String) The display name of the segment.
- `color` (String) The hexadecimal color value of the segment.
- `scope` (String) The target entity that a segment applies to.Valid values: Session, Customer.

### Optional

- `description` (String) A description of the segment.

### Read-Only

- `id` (String) The ID of this resource.