---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "airbyte_source_amplitude Data Source - terraform-provider-airbyte"
subcategory: ""
description: |-
  SourceAmplitude DataSource
---

# airbyte_source_amplitude (Data Source)

SourceAmplitude DataSource

## Example Usage

```terraform
data "airbyte_source_amplitude" "my_source_amplitude" {
  secret_id = "...my_secret_id..."
  source_id = "...my_source_id..."
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `source_id` (String)

### Optional

- `secret_id` (String) Optional secretID obtained through the public API OAuth redirect flow.

### Read-Only

- `configuration` (Attributes) (see [below for nested schema](#nestedatt--configuration))
- `name` (String)
- `workspace_id` (String)

<a id="nestedatt--configuration"></a>
### Nested Schema for `configuration`

Read-Only:

- `api_key` (String) Amplitude API Key. See the <a href="https://docs.airbyte.com/integrations/sources/amplitude#setup-guide">setup guide</a> for more information on how to obtain this key.
- `data_region` (String) must be one of ["Standard Server", "EU Residency Server"]
Amplitude data region server
- `request_time_range` (Number) According to <a href="https://www.docs.developers.amplitude.com/analytics/apis/export-api/#considerations">Considerations</a> too big time range in request can cause a timeout error. In this case, set shorter time interval in hours.
- `secret_key` (String) Amplitude Secret Key. See the <a href="https://docs.airbyte.com/integrations/sources/amplitude#setup-guide">setup guide</a> for more information on how to obtain this key.
- `source_type` (String) must be one of ["amplitude"]
- `start_date` (String) UTC date and time in the format 2021-01-25T00:00:00Z. Any data before this date will not be replicated.

