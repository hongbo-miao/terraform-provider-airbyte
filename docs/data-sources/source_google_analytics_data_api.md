---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "airbyte_source_google_analytics_data_api Data Source - terraform-provider-airbyte"
subcategory: ""
description: |-
  SourceGoogleAnalyticsDataAPI DataSource
---

# airbyte_source_google_analytics_data_api (Data Source)

SourceGoogleAnalyticsDataAPI DataSource

## Example Usage

```terraform
data "airbyte_source_google_analytics_data_api" "my_source_googleanalyticsdataapi" {
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

- `credentials` (Attributes) Credentials for the service (see [below for nested schema](#nestedatt--configuration--credentials))
- `custom_reports` (String) A JSON array describing the custom reports you want to sync from Google Analytics. See <a href="https://docs.airbyte.com/integrations/sources/google-analytics-v4/#custom-reports">the docs</a> for more information about the exact format you can use to fill out this field.
- `date_ranges_start_date` (String) The start date from which to replicate report data in the format YYYY-MM-DD. Data generated before this date will not be included in the report. Not applied to custom Cohort reports.
- `property_id` (String) A Google Analytics GA4 property identifier whose events are tracked. Specified in the URL path and not the body such as "123...". See <a href="https://developers.google.com/analytics/devguides/reporting/data/v1/property-id#what_is_my_property_id">the docs</a> for more details.
- `source_type` (String) must be one of ["google-analytics-data-api"]
- `window_in_days` (Number) The time increment used by the connector when requesting data from the Google Analytics API. More information is available in the <a href="https://docs.airbyte.com/integrations/sources/google-analytics-v4/#sampling-in-reports">the docs</a>. The bigger this value is, the faster the sync will be, but the more likely that sampling will be applied to your data, potentially causing inaccuracies in the returned results. We recommend setting this to 1 unless you have a hard requirement to make the sync faster at the expense of accuracy. The minimum allowed value for this field is 1, and the maximum is 364. Not applied to custom Cohort reports.

<a id="nestedatt--configuration--credentials"></a>
### Nested Schema for `configuration.credentials`

Read-Only:

- `source_google_analytics_data_api_credentials_authenticate_via_google_oauth` (Attributes) Credentials for the service (see [below for nested schema](#nestedatt--configuration--credentials--source_google_analytics_data_api_credentials_authenticate_via_google_oauth))
- `source_google_analytics_data_api_credentials_service_account_key_authentication` (Attributes) Credentials for the service (see [below for nested schema](#nestedatt--configuration--credentials--source_google_analytics_data_api_credentials_service_account_key_authentication))
- `source_google_analytics_data_api_update_credentials_authenticate_via_google_oauth` (Attributes) Credentials for the service (see [below for nested schema](#nestedatt--configuration--credentials--source_google_analytics_data_api_update_credentials_authenticate_via_google_oauth))
- `source_google_analytics_data_api_update_credentials_service_account_key_authentication` (Attributes) Credentials for the service (see [below for nested schema](#nestedatt--configuration--credentials--source_google_analytics_data_api_update_credentials_service_account_key_authentication))

<a id="nestedatt--configuration--credentials--source_google_analytics_data_api_credentials_authenticate_via_google_oauth"></a>
### Nested Schema for `configuration.credentials.source_google_analytics_data_api_credentials_authenticate_via_google_oauth`

Read-Only:

- `access_token` (String) Access Token for making authenticated requests.
- `auth_type` (String) must be one of ["Client"]
- `client_id` (String) The Client ID of your Google Analytics developer application.
- `client_secret` (String) The Client Secret of your Google Analytics developer application.
- `refresh_token` (String) The token for obtaining a new access token.


<a id="nestedatt--configuration--credentials--source_google_analytics_data_api_credentials_service_account_key_authentication"></a>
### Nested Schema for `configuration.credentials.source_google_analytics_data_api_credentials_service_account_key_authentication`

Read-Only:

- `auth_type` (String) must be one of ["Service"]
- `credentials_json` (String) The JSON key of the service account to use for authorization


<a id="nestedatt--configuration--credentials--source_google_analytics_data_api_update_credentials_authenticate_via_google_oauth"></a>
### Nested Schema for `configuration.credentials.source_google_analytics_data_api_update_credentials_authenticate_via_google_oauth`

Read-Only:

- `access_token` (String) Access Token for making authenticated requests.
- `auth_type` (String) must be one of ["Client"]
- `client_id` (String) The Client ID of your Google Analytics developer application.
- `client_secret` (String) The Client Secret of your Google Analytics developer application.
- `refresh_token` (String) The token for obtaining a new access token.


<a id="nestedatt--configuration--credentials--source_google_analytics_data_api_update_credentials_service_account_key_authentication"></a>
### Nested Schema for `configuration.credentials.source_google_analytics_data_api_update_credentials_service_account_key_authentication`

Read-Only:

- `auth_type` (String) must be one of ["Service"]
- `credentials_json` (String) The JSON key of the service account to use for authorization

