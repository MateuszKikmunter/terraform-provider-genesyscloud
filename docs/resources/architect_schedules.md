---
page_title: "genesyscloud_architect_schedules Resource - terraform-provider-genesyscloud"
subcategory: ""
description: |-
  Genesys Cloud Architect Schedules
---
# genesyscloud_architect_schedules (Resource)

Genesys Cloud Architect Schedules

## API Usage
The following Genesys Cloud APIs are used by this resource. Ensure your OAuth Client has been granted the necessary scopes and permissions to perform these operations:

**No APIs**



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **end** (String) Date time is represented as an ISO-8601 string without a timezone.
- **name** (String) Name of the schedule.
- **rrule** (String) An iCal Recurrence Rule (RRULE) string.
- **start** (String) Date time is represented as an ISO-8601 string without a timezone.

### Optional

- **description** (String) Description of the schedule.
- **id** (String) The ID of this resource.
- **state** (String) Indicates if the schedule is active, inactive or deleted.
- **version** (Number) Schedule's current version.
