// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
)

type SourceMailchimpAuthenticationAPIKeyAuthTypeEnum string

const (
	SourceMailchimpAuthenticationAPIKeyAuthTypeEnumApikey SourceMailchimpAuthenticationAPIKeyAuthTypeEnum = "apikey"
)

func (e SourceMailchimpAuthenticationAPIKeyAuthTypeEnum) ToPointer() *SourceMailchimpAuthenticationAPIKeyAuthTypeEnum {
	return &e
}

func (e *SourceMailchimpAuthenticationAPIKeyAuthTypeEnum) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "apikey":
		*e = SourceMailchimpAuthenticationAPIKeyAuthTypeEnum(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceMailchimpAuthenticationAPIKeyAuthTypeEnum: %v", v)
	}
}

type SourceMailchimpAuthenticationAPIKey struct {
	// Mailchimp API Key. See the <a href="https://docs.airbyte.com/integrations/sources/mailchimp">docs</a> for information on how to generate this key.
	Apikey   string                                          `json:"apikey"`
	AuthType SourceMailchimpAuthenticationAPIKeyAuthTypeEnum `json:"auth_type"`
}

type SourceMailchimpAuthenticationOAuth20AuthTypeEnum string

const (
	SourceMailchimpAuthenticationOAuth20AuthTypeEnumOauth20 SourceMailchimpAuthenticationOAuth20AuthTypeEnum = "oauth2.0"
)

func (e SourceMailchimpAuthenticationOAuth20AuthTypeEnum) ToPointer() *SourceMailchimpAuthenticationOAuth20AuthTypeEnum {
	return &e
}

func (e *SourceMailchimpAuthenticationOAuth20AuthTypeEnum) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "oauth2.0":
		*e = SourceMailchimpAuthenticationOAuth20AuthTypeEnum(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceMailchimpAuthenticationOAuth20AuthTypeEnum: %v", v)
	}
}

type SourceMailchimpAuthenticationOAuth20 struct {
	// An access token generated using the above client ID and secret.
	AccessToken string                                           `json:"access_token"`
	AuthType    SourceMailchimpAuthenticationOAuth20AuthTypeEnum `json:"auth_type"`
	// The Client ID of your OAuth application.
	ClientID *string `json:"client_id,omitempty"`
	// The Client Secret of your OAuth application.
	ClientSecret *string `json:"client_secret,omitempty"`
}

type SourceMailchimpAuthenticationType string

const (
	SourceMailchimpAuthenticationTypeSourceMailchimpAuthenticationOAuth20 SourceMailchimpAuthenticationType = "source-mailchimp_Authentication_OAuth2.0"
	SourceMailchimpAuthenticationTypeSourceMailchimpAuthenticationAPIKey  SourceMailchimpAuthenticationType = "source-mailchimp_Authentication_API Key"
)

type SourceMailchimpAuthentication struct {
	SourceMailchimpAuthenticationOAuth20 *SourceMailchimpAuthenticationOAuth20
	SourceMailchimpAuthenticationAPIKey  *SourceMailchimpAuthenticationAPIKey

	Type SourceMailchimpAuthenticationType
}

func CreateSourceMailchimpAuthenticationSourceMailchimpAuthenticationOAuth20(sourceMailchimpAuthenticationOAuth20 SourceMailchimpAuthenticationOAuth20) SourceMailchimpAuthentication {
	typ := SourceMailchimpAuthenticationTypeSourceMailchimpAuthenticationOAuth20

	return SourceMailchimpAuthentication{
		SourceMailchimpAuthenticationOAuth20: &sourceMailchimpAuthenticationOAuth20,
		Type:                                 typ,
	}
}

func CreateSourceMailchimpAuthenticationSourceMailchimpAuthenticationAPIKey(sourceMailchimpAuthenticationAPIKey SourceMailchimpAuthenticationAPIKey) SourceMailchimpAuthentication {
	typ := SourceMailchimpAuthenticationTypeSourceMailchimpAuthenticationAPIKey

	return SourceMailchimpAuthentication{
		SourceMailchimpAuthenticationAPIKey: &sourceMailchimpAuthenticationAPIKey,
		Type:                                typ,
	}
}

func (u *SourceMailchimpAuthentication) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	sourceMailchimpAuthenticationOAuth20 := new(SourceMailchimpAuthenticationOAuth20)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&sourceMailchimpAuthenticationOAuth20); err == nil {
		u.SourceMailchimpAuthenticationOAuth20 = sourceMailchimpAuthenticationOAuth20
		u.Type = SourceMailchimpAuthenticationTypeSourceMailchimpAuthenticationOAuth20
		return nil
	}

	sourceMailchimpAuthenticationAPIKey := new(SourceMailchimpAuthenticationAPIKey)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&sourceMailchimpAuthenticationAPIKey); err == nil {
		u.SourceMailchimpAuthenticationAPIKey = sourceMailchimpAuthenticationAPIKey
		u.Type = SourceMailchimpAuthenticationTypeSourceMailchimpAuthenticationAPIKey
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u SourceMailchimpAuthentication) MarshalJSON() ([]byte, error) {
	if u.SourceMailchimpAuthenticationOAuth20 != nil {
		return json.Marshal(u.SourceMailchimpAuthenticationOAuth20)
	}

	if u.SourceMailchimpAuthenticationAPIKey != nil {
		return json.Marshal(u.SourceMailchimpAuthenticationAPIKey)
	}

	return nil, nil
}

type SourceMailchimpMailchimpEnum string

const (
	SourceMailchimpMailchimpEnumMailchimp SourceMailchimpMailchimpEnum = "mailchimp"
)

func (e SourceMailchimpMailchimpEnum) ToPointer() *SourceMailchimpMailchimpEnum {
	return &e
}

func (e *SourceMailchimpMailchimpEnum) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "mailchimp":
		*e = SourceMailchimpMailchimpEnum(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceMailchimpMailchimpEnum: %v", v)
	}
}

// SourceMailchimp - The values required to configure the source.
type SourceMailchimp struct {
	CampaignID  *string                        `json:"campaign_id,omitempty"`
	Credentials *SourceMailchimpAuthentication `json:"credentials,omitempty"`
	SourceType  SourceMailchimpMailchimpEnum   `json:"sourceType"`
}
