// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

type SourceQuickbooksAuthorizationMethodOAuth20AuthType string

const (
	SourceQuickbooksAuthorizationMethodOAuth20AuthTypeOauth20 SourceQuickbooksAuthorizationMethodOAuth20AuthType = "oauth2.0"
)

func (e SourceQuickbooksAuthorizationMethodOAuth20AuthType) ToPointer() *SourceQuickbooksAuthorizationMethodOAuth20AuthType {
	return &e
}

func (e *SourceQuickbooksAuthorizationMethodOAuth20AuthType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "oauth2.0":
		*e = SourceQuickbooksAuthorizationMethodOAuth20AuthType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceQuickbooksAuthorizationMethodOAuth20AuthType: %v", v)
	}
}

type SourceQuickbooksAuthorizationMethodOAuth20 struct {
	// Access token fot making authenticated requests.
	AccessToken string                                              `json:"access_token"`
	AuthType    *SourceQuickbooksAuthorizationMethodOAuth20AuthType `json:"auth_type,omitempty"`
	// Identifies which app is making the request. Obtain this value from the Keys tab on the app profile via My Apps on the developer site. There are two versions of this key: development and production.
	ClientID string `json:"client_id"`
	//  Obtain this value from the Keys tab on the app profile via My Apps on the developer site. There are two versions of this key: development and production.
	ClientSecret string `json:"client_secret"`
	// Labeled Company ID. The Make API Calls panel is populated with the realm id and the current access token.
	RealmID string `json:"realm_id"`
	// A token used when refreshing the access token.
	RefreshToken string `json:"refresh_token"`
	// The date-time when the access token should be refreshed.
	TokenExpiryDate time.Time `json:"token_expiry_date"`
}

type SourceQuickbooksAuthorizationMethodType string

const (
	SourceQuickbooksAuthorizationMethodTypeSourceQuickbooksAuthorizationMethodOAuth20 SourceQuickbooksAuthorizationMethodType = "source-quickbooks_Authorization Method_OAuth2.0"
)

type SourceQuickbooksAuthorizationMethod struct {
	SourceQuickbooksAuthorizationMethodOAuth20 *SourceQuickbooksAuthorizationMethodOAuth20

	Type SourceQuickbooksAuthorizationMethodType
}

func CreateSourceQuickbooksAuthorizationMethodSourceQuickbooksAuthorizationMethodOAuth20(sourceQuickbooksAuthorizationMethodOAuth20 SourceQuickbooksAuthorizationMethodOAuth20) SourceQuickbooksAuthorizationMethod {
	typ := SourceQuickbooksAuthorizationMethodTypeSourceQuickbooksAuthorizationMethodOAuth20

	return SourceQuickbooksAuthorizationMethod{
		SourceQuickbooksAuthorizationMethodOAuth20: &sourceQuickbooksAuthorizationMethodOAuth20,
		Type: typ,
	}
}

func (u *SourceQuickbooksAuthorizationMethod) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	sourceQuickbooksAuthorizationMethodOAuth20 := new(SourceQuickbooksAuthorizationMethodOAuth20)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&sourceQuickbooksAuthorizationMethodOAuth20); err == nil {
		u.SourceQuickbooksAuthorizationMethodOAuth20 = sourceQuickbooksAuthorizationMethodOAuth20
		u.Type = SourceQuickbooksAuthorizationMethodTypeSourceQuickbooksAuthorizationMethodOAuth20
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u SourceQuickbooksAuthorizationMethod) MarshalJSON() ([]byte, error) {
	if u.SourceQuickbooksAuthorizationMethodOAuth20 != nil {
		return json.Marshal(u.SourceQuickbooksAuthorizationMethodOAuth20)
	}

	return nil, nil
}

type SourceQuickbooksQuickbooks string

const (
	SourceQuickbooksQuickbooksQuickbooks SourceQuickbooksQuickbooks = "quickbooks"
)

func (e SourceQuickbooksQuickbooks) ToPointer() *SourceQuickbooksQuickbooks {
	return &e
}

func (e *SourceQuickbooksQuickbooks) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "quickbooks":
		*e = SourceQuickbooksQuickbooks(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceQuickbooksQuickbooks: %v", v)
	}
}

type SourceQuickbooks struct {
	Credentials SourceQuickbooksAuthorizationMethod `json:"credentials"`
	// Determines whether to use the sandbox or production environment.
	Sandbox    bool                       `json:"sandbox"`
	SourceType SourceQuickbooksQuickbooks `json:"sourceType"`
	// The default value to use if no bookmark exists for an endpoint (rfc3339 date string). E.g, 2021-03-20T00:00:00+00:00. Any data before this date will not be replicated.
	StartDate time.Time `json:"start_date"`
}