---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "airbyte_source_chargebee Data Source - terraform-provider-airbyte"
subcategory: ""
description: |-
  SourceChargebee DataSource
---

# airbyte_source_chargebee (Data Source)

SourceChargebee DataSource

## Example Usage

```terraform
data "airbyte_source_chargebee" "my_source_chargebee" {
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

- `product_catalog` (String) must be one of ["1.0", "2.0"]
Product Catalog version of your Chargebee site. Instructions on how to find your version you may find <a href="https://apidocs.chargebee.com/docs/api?prod_cat_ver=2">here</a> under `API Version` section.
- `site` (String) The site prefix for your Chargebee instance.
- `site_api_key` (String) Chargebee API Key. See the <a href="https://docs.airbyte.com/integrations/sources/chargebee">docs</a> for more information on how to obtain this key.
- `source_type` (String) must be one of ["chargebee"]
- `start_date` (String) UTC date and time in the format 2021-01-25T00:00:00Z. Any data before this date will not be replicated.

