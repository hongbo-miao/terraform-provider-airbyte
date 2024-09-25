---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "airbyte_destination_timeplus Data Source - terraform-provider-airbyte"
subcategory: ""
description: |-
  DestinationTimeplus DataSource
---

# airbyte_destination_timeplus (Data Source)

DestinationTimeplus DataSource

## Example Usage

```terraform
data "airbyte_destination_timeplus" "my_destination_timeplus" {
  destination_id = "...my_destination_id..."
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `destination_id` (String)

### Read-Only

- `configuration` (String) The values required to configure the destination. Parsed as JSON.
- `destination_type` (String)
- `name` (String)
- `workspace_id` (String)

