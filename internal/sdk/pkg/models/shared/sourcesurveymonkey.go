// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"time"
)

type SourceSurveymonkeySurveyMonkeyAuthorizationMethodAuthMethodEnum string

const (
	SourceSurveymonkeySurveyMonkeyAuthorizationMethodAuthMethodEnumOauth20 SourceSurveymonkeySurveyMonkeyAuthorizationMethodAuthMethodEnum = "oauth2.0"
)

func (e SourceSurveymonkeySurveyMonkeyAuthorizationMethodAuthMethodEnum) ToPointer() *SourceSurveymonkeySurveyMonkeyAuthorizationMethodAuthMethodEnum {
	return &e
}

func (e *SourceSurveymonkeySurveyMonkeyAuthorizationMethodAuthMethodEnum) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "oauth2.0":
		*e = SourceSurveymonkeySurveyMonkeyAuthorizationMethodAuthMethodEnum(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceSurveymonkeySurveyMonkeyAuthorizationMethodAuthMethodEnum: %v", v)
	}
}

// SourceSurveymonkeySurveyMonkeyAuthorizationMethod - The authorization method to use to retrieve data from SurveyMonkey
type SourceSurveymonkeySurveyMonkeyAuthorizationMethod struct {
	// Access Token for making authenticated requests. See the <a href="https://docs.airbyte.io/integrations/sources/surveymonkey">docs</a> for information on how to generate this key.
	AccessToken string                                                          `json:"access_token"`
	AuthMethod  SourceSurveymonkeySurveyMonkeyAuthorizationMethodAuthMethodEnum `json:"auth_method"`
	// The Client ID of the SurveyMonkey developer application.
	ClientID *string `json:"client_id,omitempty"`
	// The Client Secret of the SurveyMonkey developer application.
	ClientSecret *string `json:"client_secret,omitempty"`
}

// SourceSurveymonkeyOriginDatacenterOfTheSurveyMonkeyAccountEnum - Depending on the originating datacenter of the SurveyMonkey account, the API access URL may be different.
type SourceSurveymonkeyOriginDatacenterOfTheSurveyMonkeyAccountEnum string

const (
	SourceSurveymonkeyOriginDatacenterOfTheSurveyMonkeyAccountEnumUsa    SourceSurveymonkeyOriginDatacenterOfTheSurveyMonkeyAccountEnum = "USA"
	SourceSurveymonkeyOriginDatacenterOfTheSurveyMonkeyAccountEnumEurope SourceSurveymonkeyOriginDatacenterOfTheSurveyMonkeyAccountEnum = "Europe"
	SourceSurveymonkeyOriginDatacenterOfTheSurveyMonkeyAccountEnumCanada SourceSurveymonkeyOriginDatacenterOfTheSurveyMonkeyAccountEnum = "Canada"
)

func (e SourceSurveymonkeyOriginDatacenterOfTheSurveyMonkeyAccountEnum) ToPointer() *SourceSurveymonkeyOriginDatacenterOfTheSurveyMonkeyAccountEnum {
	return &e
}

func (e *SourceSurveymonkeyOriginDatacenterOfTheSurveyMonkeyAccountEnum) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "USA":
		fallthrough
	case "Europe":
		fallthrough
	case "Canada":
		*e = SourceSurveymonkeyOriginDatacenterOfTheSurveyMonkeyAccountEnum(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceSurveymonkeyOriginDatacenterOfTheSurveyMonkeyAccountEnum: %v", v)
	}
}

type SourceSurveymonkeySurveymonkeyEnum string

const (
	SourceSurveymonkeySurveymonkeyEnumSurveymonkey SourceSurveymonkeySurveymonkeyEnum = "surveymonkey"
)

func (e SourceSurveymonkeySurveymonkeyEnum) ToPointer() *SourceSurveymonkeySurveymonkeyEnum {
	return &e
}

func (e *SourceSurveymonkeySurveymonkeyEnum) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "surveymonkey":
		*e = SourceSurveymonkeySurveymonkeyEnum(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceSurveymonkeySurveymonkeyEnum: %v", v)
	}
}

// SourceSurveymonkey - The values required to configure the source.
type SourceSurveymonkey struct {
	// The authorization method to use to retrieve data from SurveyMonkey
	Credentials *SourceSurveymonkeySurveyMonkeyAuthorizationMethod `json:"credentials,omitempty"`
	// Depending on the originating datacenter of the SurveyMonkey account, the API access URL may be different.
	Origin     *SourceSurveymonkeyOriginDatacenterOfTheSurveyMonkeyAccountEnum `json:"origin,omitempty"`
	SourceType SourceSurveymonkeySurveymonkeyEnum                              `json:"sourceType"`
	// UTC date and time in the format 2017-01-25T00:00:00Z. Any data before this date will not be replicated.
	StartDate time.Time `json:"start_date"`
	// IDs of the surveys from which you'd like to replicate data. If left empty, data from all boards to which you have access will be replicated.
	SurveyIds []string `json:"survey_ids,omitempty"`
}
