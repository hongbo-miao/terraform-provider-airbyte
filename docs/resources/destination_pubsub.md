---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "airbyte_destination_pubsub Resource - terraform-provider-airbyte"
subcategory: ""
description: |-
  DestinationPubsub Resource
---

# airbyte_destination_pubsub (Resource)

DestinationPubsub Resource



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `configuration` (Attributes) (see [below for nested schema](#nestedatt--configuration))
- `name` (String)
- `workspace_id` (String)

### Read-Only

- `destination_id` (String)
- `destination_type` (String)

<a id="nestedatt--configuration"></a>
### Nested Schema for `configuration`

Required:

- `batching_enabled` (Boolean)
- `credentials_json` (String)
- `destination_type` (String)
- `ordering_enabled` (Boolean)
- `project_id` (String)
- `topic_id` (String)

Optional:

- `batching_delay_threshold` (Number)
- `batching_element_count_threshold` (Number)
- `batching_request_bytes_threshold` (Number)

