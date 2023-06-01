---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "airbyte_source_bigquery Resource - terraform-provider-airbyte"
subcategory: ""
description: |-
  SourceBigquery Resource
---

# airbyte_source_bigquery (Resource)

SourceBigquery Resource



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

- `credentials_json` (String)
- `project_id` (String)
- `source_type` (String)

Optional:

- `dataset_id` (String)

