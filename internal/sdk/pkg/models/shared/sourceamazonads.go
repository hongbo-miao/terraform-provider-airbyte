// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type SourceAmazonAdsAuthTypeEnum string

const (
	SourceAmazonAdsAuthTypeEnumOauth20 SourceAmazonAdsAuthTypeEnum = "oauth2.0"
)

func (e SourceAmazonAdsAuthTypeEnum) ToPointer() *SourceAmazonAdsAuthTypeEnum {
	return &e
}

func (e *SourceAmazonAdsAuthTypeEnum) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "oauth2.0":
		*e = SourceAmazonAdsAuthTypeEnum(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceAmazonAdsAuthTypeEnum: %v", v)
	}
}

// SourceAmazonAdsRegionEnum - Region to pull data from (EU/NA/FE). See <a href="https://advertising.amazon.com/API/docs/en-us/info/api-overview#api-endpoints">docs</a> for more details.
type SourceAmazonAdsRegionEnum string

const (
	SourceAmazonAdsRegionEnumNa SourceAmazonAdsRegionEnum = "NA"
	SourceAmazonAdsRegionEnumEu SourceAmazonAdsRegionEnum = "EU"
	SourceAmazonAdsRegionEnumFe SourceAmazonAdsRegionEnum = "FE"
)

func (e SourceAmazonAdsRegionEnum) ToPointer() *SourceAmazonAdsRegionEnum {
	return &e
}

func (e *SourceAmazonAdsRegionEnum) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "NA":
		fallthrough
	case "EU":
		fallthrough
	case "FE":
		*e = SourceAmazonAdsRegionEnum(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceAmazonAdsRegionEnum: %v", v)
	}
}

type SourceAmazonAdsReportRecordTypesEnum string

const (
	SourceAmazonAdsReportRecordTypesEnumAdGroups      SourceAmazonAdsReportRecordTypesEnum = "adGroups"
	SourceAmazonAdsReportRecordTypesEnumAsins         SourceAmazonAdsReportRecordTypesEnum = "asins"
	SourceAmazonAdsReportRecordTypesEnumAsinsKeywords SourceAmazonAdsReportRecordTypesEnum = "asins_keywords"
	SourceAmazonAdsReportRecordTypesEnumAsinsTargets  SourceAmazonAdsReportRecordTypesEnum = "asins_targets"
	SourceAmazonAdsReportRecordTypesEnumCampaigns     SourceAmazonAdsReportRecordTypesEnum = "campaigns"
	SourceAmazonAdsReportRecordTypesEnumKeywords      SourceAmazonAdsReportRecordTypesEnum = "keywords"
	SourceAmazonAdsReportRecordTypesEnumProductAds    SourceAmazonAdsReportRecordTypesEnum = "productAds"
	SourceAmazonAdsReportRecordTypesEnumTargets       SourceAmazonAdsReportRecordTypesEnum = "targets"
)

func (e SourceAmazonAdsReportRecordTypesEnum) ToPointer() *SourceAmazonAdsReportRecordTypesEnum {
	return &e
}

func (e *SourceAmazonAdsReportRecordTypesEnum) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "adGroups":
		fallthrough
	case "asins":
		fallthrough
	case "asins_keywords":
		fallthrough
	case "asins_targets":
		fallthrough
	case "campaigns":
		fallthrough
	case "keywords":
		fallthrough
	case "productAds":
		fallthrough
	case "targets":
		*e = SourceAmazonAdsReportRecordTypesEnum(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceAmazonAdsReportRecordTypesEnum: %v", v)
	}
}

type SourceAmazonAdsAmazonAdsEnum string

const (
	SourceAmazonAdsAmazonAdsEnumAmazonAds SourceAmazonAdsAmazonAdsEnum = "amazon-ads"
)

func (e SourceAmazonAdsAmazonAdsEnum) ToPointer() *SourceAmazonAdsAmazonAdsEnum {
	return &e
}

func (e *SourceAmazonAdsAmazonAdsEnum) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "amazon-ads":
		*e = SourceAmazonAdsAmazonAdsEnum(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceAmazonAdsAmazonAdsEnum: %v", v)
	}
}

type SourceAmazonAdsStateFilterEnum string

const (
	SourceAmazonAdsStateFilterEnumEnabled  SourceAmazonAdsStateFilterEnum = "enabled"
	SourceAmazonAdsStateFilterEnumPaused   SourceAmazonAdsStateFilterEnum = "paused"
	SourceAmazonAdsStateFilterEnumArchived SourceAmazonAdsStateFilterEnum = "archived"
)

func (e SourceAmazonAdsStateFilterEnum) ToPointer() *SourceAmazonAdsStateFilterEnum {
	return &e
}

func (e *SourceAmazonAdsStateFilterEnum) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "enabled":
		fallthrough
	case "paused":
		fallthrough
	case "archived":
		*e = SourceAmazonAdsStateFilterEnum(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceAmazonAdsStateFilterEnum: %v", v)
	}
}

// SourceAmazonAds - The values required to configure the source.
type SourceAmazonAds struct {
	AuthType *SourceAmazonAdsAuthTypeEnum `json:"auth_type,omitempty"`
	// The client ID of your Amazon Ads developer application. See the <a href="https://advertising.amazon.com/API/docs/en-us/get-started/generate-api-tokens#retrieve-your-client-id-and-client-secret">docs</a> for more information.
	ClientID string `json:"client_id"`
	// The client secret of your Amazon Ads developer application. See the <a href="https://advertising.amazon.com/API/docs/en-us/get-started/generate-api-tokens#retrieve-your-client-id-and-client-secret">docs</a> for more information.
	ClientSecret string `json:"client_secret"`
	// The amount of days to go back in time to get the updated data from Amazon Ads
	LookBackWindow *int64 `json:"look_back_window,omitempty"`
	// Profile IDs you want to fetch data for. See <a href="https://advertising.amazon.com/API/docs/en-us/concepts/authorization/profiles">docs</a> for more details.
	Profiles []int64 `json:"profiles,omitempty"`
	// Amazon Ads refresh token. See the <a href="https://advertising.amazon.com/API/docs/en-us/get-started/generate-api-tokens">docs</a> for more information on how to obtain this token.
	RefreshToken string `json:"refresh_token"`
	// Region to pull data from (EU/NA/FE). See <a href="https://advertising.amazon.com/API/docs/en-us/info/api-overview#api-endpoints">docs</a> for more details.
	Region *SourceAmazonAdsRegionEnum `json:"region,omitempty"`
	// Optional configuration which accepts an array of string of record types. Leave blank for default behaviour to pull all report types. Use this config option only if you want to pull specific report type(s). See <a href="https://advertising.amazon.com/API/docs/en-us/reporting/v2/report-types">docs</a> for more details
	ReportRecordTypes []SourceAmazonAdsReportRecordTypesEnum `json:"report_record_types,omitempty"`
	SourceType        SourceAmazonAdsAmazonAdsEnum           `json:"sourceType"`
	// The Start date for collecting reports, should not be more than 60 days in the past. In YYYY-MM-DD format
	StartDate *string `json:"start_date,omitempty"`
	// Reflects the state of the Display, Product, and Brand Campaign streams as enabled, paused, or archived. If you do not populate this field, it will be ignored completely.
	StateFilter []SourceAmazonAdsStateFilterEnum `json:"state_filter,omitempty"`
}
