---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "airbyte_source_mailjet_sms Data Source - terraform-provider-airbyte"
subcategory: ""
description: |-
  SourceMailjetSms DataSource
---

# airbyte_source_mailjet_sms (Data Source)

SourceMailjetSms DataSource

## Example Usage

```terraform
data "airbyte_source_mailjet_sms" "my_source_mailjetsms" {
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

- `end_date` (Number) Retrieve SMS messages created before the specified timestamp. Required format - Unix timestamp.
- `source_type` (String) must be one of ["mailjet-sms"]
- `start_date` (Number) Retrieve SMS messages created after the specified timestamp. Required format - Unix timestamp.
- `token` (String) Your access token. See <a href="https://dev.mailjet.com/sms/reference/overview/authentication">here</a>.

