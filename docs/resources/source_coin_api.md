---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "airbyte_source_coin_api Resource - terraform-provider-airbyte"
subcategory: ""
description: |-
  SourceCoinAPI Resource
---

# airbyte_source_coin_api (Resource)

SourceCoinAPI Resource



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `configuration` (Attributes) (see [below for nested schema](#nestedatt--configuration))
- `name` (String)
- `workspace_id` (String)

### Optional

- `secret_id` (String)

### Read-Only

- `source_id` (String)
- `source_type` (String)

<a id="nestedatt--configuration"></a>
### Nested Schema for `configuration`

Required:

- `api_key` (String)
- `environment` (String) The environment to use. Either sandbox or production.
- `period` (String)
- `source_type` (String)
- `start_date` (String)
- `symbol_id` (String)

Optional:

- `end_date` (String)
- `limit` (Number)

