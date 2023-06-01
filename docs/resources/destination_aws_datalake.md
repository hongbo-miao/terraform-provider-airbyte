---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "airbyte_destination_aws_datalake Resource - terraform-provider-airbyte"
subcategory: ""
description: |-
  DestinationAwsDatalake Resource
---

# airbyte_destination_aws_datalake (Resource)

DestinationAwsDatalake Resource



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

- `bucket_name` (String)
- `credentials` (Attributes) (see [below for nested schema](#nestedatt--configuration--credentials))
- `destination_type` (String)
- `lakeformation_database_name` (String)
- `region` (String) The region of the S3 bucket. See <a href="https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/using-regions-availability-zones.html#concepts-available-regions">here</a> for all region codes.

Optional:

- `aws_account_id` (String)
- `bucket_prefix` (String)
- `format` (Attributes) (see [below for nested schema](#nestedatt--configuration--format))
- `glue_catalog_float_as_decimal` (Boolean)
- `lakeformation_database_default_tag_key` (String)
- `lakeformation_database_default_tag_values` (String)
- `lakeformation_governed_tables` (Boolean)
- `partitioning` (String) Partition data by cursor fields when a cursor field is a date

<a id="nestedatt--configuration--credentials"></a>
### Nested Schema for `configuration.credentials`

Optional:

- `destination_aws_datalake_authentication_mode_iam_role` (Attributes) Choose How to Authenticate to AWS. (see [below for nested schema](#nestedatt--configuration--credentials--destination_aws_datalake_authentication_mode_iam_role))
- `destination_aws_datalake_authentication_mode_iam_user` (Attributes) Choose How to Authenticate to AWS. (see [below for nested schema](#nestedatt--configuration--credentials--destination_aws_datalake_authentication_mode_iam_user))

Read-Only:

- `destination_aws_datalake_update_authentication_mode_iam_role` (Attributes) Choose How to Authenticate to AWS. (see [below for nested schema](#nestedatt--configuration--credentials--destination_aws_datalake_update_authentication_mode_iam_role))
- `destination_aws_datalake_update_authentication_mode_iam_user` (Attributes) Choose How to Authenticate to AWS. (see [below for nested schema](#nestedatt--configuration--credentials--destination_aws_datalake_update_authentication_mode_iam_user))

<a id="nestedatt--configuration--credentials--destination_aws_datalake_authentication_mode_iam_role"></a>
### Nested Schema for `configuration.credentials.destination_aws_datalake_authentication_mode_iam_role`

Required:

- `credentials_title` (String) Name of the credentials
- `role_arn` (String)


<a id="nestedatt--configuration--credentials--destination_aws_datalake_authentication_mode_iam_user"></a>
### Nested Schema for `configuration.credentials.destination_aws_datalake_authentication_mode_iam_user`

Required:

- `aws_access_key_id` (String)
- `aws_secret_access_key` (String)
- `credentials_title` (String) Name of the credentials


<a id="nestedatt--configuration--credentials--destination_aws_datalake_update_authentication_mode_iam_role"></a>
### Nested Schema for `configuration.credentials.destination_aws_datalake_update_authentication_mode_iam_role`

Read-Only:

- `credentials_title` (String) Name of the credentials
- `role_arn` (String)


<a id="nestedatt--configuration--credentials--destination_aws_datalake_update_authentication_mode_iam_user"></a>
### Nested Schema for `configuration.credentials.destination_aws_datalake_update_authentication_mode_iam_user`

Read-Only:

- `aws_access_key_id` (String)
- `aws_secret_access_key` (String)
- `credentials_title` (String) Name of the credentials



<a id="nestedatt--configuration--format"></a>
### Nested Schema for `configuration.format`

Optional:

- `destination_aws_datalake_output_format_wildcard_json_lines_newline_delimited_json` (Attributes) Format of the data output. (see [below for nested schema](#nestedatt--configuration--format--destination_aws_datalake_output_format_wildcard_json_lines_newline_delimited_json))
- `destination_aws_datalake_output_format_wildcard_parquet_columnar_storage` (Attributes) Format of the data output. (see [below for nested schema](#nestedatt--configuration--format--destination_aws_datalake_output_format_wildcard_parquet_columnar_storage))

Read-Only:

- `destination_aws_datalake_update_output_format_wildcard_json_lines_newline_delimited_json` (Attributes) Format of the data output. (see [below for nested schema](#nestedatt--configuration--format--destination_aws_datalake_update_output_format_wildcard_json_lines_newline_delimited_json))
- `destination_aws_datalake_update_output_format_wildcard_parquet_columnar_storage` (Attributes) Format of the data output. (see [below for nested schema](#nestedatt--configuration--format--destination_aws_datalake_update_output_format_wildcard_parquet_columnar_storage))

<a id="nestedatt--configuration--format--destination_aws_datalake_output_format_wildcard_json_lines_newline_delimited_json"></a>
### Nested Schema for `configuration.format.destination_aws_datalake_output_format_wildcard_json_lines_newline_delimited_json`

Required:

- `format_type` (String)

Optional:

- `compression_codec` (String) The compression algorithm used to compress data.


<a id="nestedatt--configuration--format--destination_aws_datalake_output_format_wildcard_parquet_columnar_storage"></a>
### Nested Schema for `configuration.format.destination_aws_datalake_output_format_wildcard_parquet_columnar_storage`

Required:

- `format_type` (String)

Optional:

- `compression_codec` (String) The compression algorithm used to compress data.


<a id="nestedatt--configuration--format--destination_aws_datalake_update_output_format_wildcard_json_lines_newline_delimited_json"></a>
### Nested Schema for `configuration.format.destination_aws_datalake_update_output_format_wildcard_json_lines_newline_delimited_json`

Read-Only:

- `compression_codec` (String) The compression algorithm used to compress data.
- `format_type` (String)


<a id="nestedatt--configuration--format--destination_aws_datalake_update_output_format_wildcard_parquet_columnar_storage"></a>
### Nested Schema for `configuration.format.destination_aws_datalake_update_output_format_wildcard_parquet_columnar_storage`

Read-Only:

- `compression_codec` (String) The compression algorithm used to compress data.
- `format_type` (String)

