---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "genesyscloud_outbound_filespecificationtemplate Data Source - terraform-provider-genesyscloud"
subcategory: ""
description: |-
  Data source for Genesys Cloud Outbound File Specification Template. Select a file specification template by name.
---

# genesyscloud_outbound_filespecificationtemplate (Data Source)

Data source for Genesys Cloud Outbound File Specification Template. Select a file specification template by name.

## Example Usage

```terraform
data "genesyscloud_outbound_filespecificationtemplate" "file_specification_template" {
  name = "Example File Specification Template"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) File Specification Template name.

### Read-Only

- `id` (String) The ID of this resource.