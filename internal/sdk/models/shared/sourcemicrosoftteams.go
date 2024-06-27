// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/internal/utils"
)

type SourceMicrosoftTeamsSchemasAuthType string

const (
	SourceMicrosoftTeamsSchemasAuthTypeToken SourceMicrosoftTeamsSchemasAuthType = "Token"
)

func (e SourceMicrosoftTeamsSchemasAuthType) ToPointer() *SourceMicrosoftTeamsSchemasAuthType {
	return &e
}
func (e *SourceMicrosoftTeamsSchemasAuthType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Token":
		*e = SourceMicrosoftTeamsSchemasAuthType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceMicrosoftTeamsSchemasAuthType: %v", v)
	}
}

type SourceMicrosoftTeamsAuthenticateViaMicrosoft struct {
	authType *SourceMicrosoftTeamsSchemasAuthType `const:"Token" json:"auth_type"`
	// The Client ID of your Microsoft Teams developer application.
	ClientID string `json:"client_id"`
	// The Client Secret of your Microsoft Teams developer application.
	ClientSecret string `json:"client_secret"`
	// A globally unique identifier (GUID) that is different than your organization name or domain. Follow these steps to obtain: open one of the Teams where you belong inside the Teams Application -> Click on the … next to the Team title -> Click on Get link to team -> Copy the link to the team and grab the tenant ID form the URL
	TenantID string `json:"tenant_id"`
}

func (s SourceMicrosoftTeamsAuthenticateViaMicrosoft) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceMicrosoftTeamsAuthenticateViaMicrosoft) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *SourceMicrosoftTeamsAuthenticateViaMicrosoft) GetAuthType() *SourceMicrosoftTeamsSchemasAuthType {
	return SourceMicrosoftTeamsSchemasAuthTypeToken.ToPointer()
}

func (o *SourceMicrosoftTeamsAuthenticateViaMicrosoft) GetClientID() string {
	if o == nil {
		return ""
	}
	return o.ClientID
}

func (o *SourceMicrosoftTeamsAuthenticateViaMicrosoft) GetClientSecret() string {
	if o == nil {
		return ""
	}
	return o.ClientSecret
}

func (o *SourceMicrosoftTeamsAuthenticateViaMicrosoft) GetTenantID() string {
	if o == nil {
		return ""
	}
	return o.TenantID
}

type SourceMicrosoftTeamsAuthType string

const (
	SourceMicrosoftTeamsAuthTypeClient SourceMicrosoftTeamsAuthType = "Client"
)

func (e SourceMicrosoftTeamsAuthType) ToPointer() *SourceMicrosoftTeamsAuthType {
	return &e
}
func (e *SourceMicrosoftTeamsAuthType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Client":
		*e = SourceMicrosoftTeamsAuthType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceMicrosoftTeamsAuthType: %v", v)
	}
}

type SourceMicrosoftTeamsAuthenticateViaMicrosoftOAuth20 struct {
	authType *SourceMicrosoftTeamsAuthType `const:"Client" json:"auth_type"`
	// The Client ID of your Microsoft Teams developer application.
	ClientID string `json:"client_id"`
	// The Client Secret of your Microsoft Teams developer application.
	ClientSecret string `json:"client_secret"`
	// A Refresh Token to renew the expired Access Token.
	RefreshToken string `json:"refresh_token"`
	// A globally unique identifier (GUID) that is different than your organization name or domain. Follow these steps to obtain: open one of the Teams where you belong inside the Teams Application -> Click on the … next to the Team title -> Click on Get link to team -> Copy the link to the team and grab the tenant ID form the URL
	TenantID string `json:"tenant_id"`
}

func (s SourceMicrosoftTeamsAuthenticateViaMicrosoftOAuth20) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceMicrosoftTeamsAuthenticateViaMicrosoftOAuth20) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *SourceMicrosoftTeamsAuthenticateViaMicrosoftOAuth20) GetAuthType() *SourceMicrosoftTeamsAuthType {
	return SourceMicrosoftTeamsAuthTypeClient.ToPointer()
}

func (o *SourceMicrosoftTeamsAuthenticateViaMicrosoftOAuth20) GetClientID() string {
	if o == nil {
		return ""
	}
	return o.ClientID
}

func (o *SourceMicrosoftTeamsAuthenticateViaMicrosoftOAuth20) GetClientSecret() string {
	if o == nil {
		return ""
	}
	return o.ClientSecret
}

func (o *SourceMicrosoftTeamsAuthenticateViaMicrosoftOAuth20) GetRefreshToken() string {
	if o == nil {
		return ""
	}
	return o.RefreshToken
}

func (o *SourceMicrosoftTeamsAuthenticateViaMicrosoftOAuth20) GetTenantID() string {
	if o == nil {
		return ""
	}
	return o.TenantID
}

type SourceMicrosoftTeamsAuthenticationMechanismType string

const (
	SourceMicrosoftTeamsAuthenticationMechanismTypeSourceMicrosoftTeamsAuthenticateViaMicrosoftOAuth20 SourceMicrosoftTeamsAuthenticationMechanismType = "source-microsoft-teams_Authenticate via Microsoft (OAuth 2.0)"
	SourceMicrosoftTeamsAuthenticationMechanismTypeSourceMicrosoftTeamsAuthenticateViaMicrosoft        SourceMicrosoftTeamsAuthenticationMechanismType = "source-microsoft-teams_Authenticate via Microsoft"
)

// SourceMicrosoftTeamsAuthenticationMechanism - Choose how to authenticate to Microsoft
type SourceMicrosoftTeamsAuthenticationMechanism struct {
	SourceMicrosoftTeamsAuthenticateViaMicrosoftOAuth20 *SourceMicrosoftTeamsAuthenticateViaMicrosoftOAuth20
	SourceMicrosoftTeamsAuthenticateViaMicrosoft        *SourceMicrosoftTeamsAuthenticateViaMicrosoft

	Type SourceMicrosoftTeamsAuthenticationMechanismType
}

func CreateSourceMicrosoftTeamsAuthenticationMechanismSourceMicrosoftTeamsAuthenticateViaMicrosoftOAuth20(sourceMicrosoftTeamsAuthenticateViaMicrosoftOAuth20 SourceMicrosoftTeamsAuthenticateViaMicrosoftOAuth20) SourceMicrosoftTeamsAuthenticationMechanism {
	typ := SourceMicrosoftTeamsAuthenticationMechanismTypeSourceMicrosoftTeamsAuthenticateViaMicrosoftOAuth20

	return SourceMicrosoftTeamsAuthenticationMechanism{
		SourceMicrosoftTeamsAuthenticateViaMicrosoftOAuth20: &sourceMicrosoftTeamsAuthenticateViaMicrosoftOAuth20,
		Type: typ,
	}
}

func CreateSourceMicrosoftTeamsAuthenticationMechanismSourceMicrosoftTeamsAuthenticateViaMicrosoft(sourceMicrosoftTeamsAuthenticateViaMicrosoft SourceMicrosoftTeamsAuthenticateViaMicrosoft) SourceMicrosoftTeamsAuthenticationMechanism {
	typ := SourceMicrosoftTeamsAuthenticationMechanismTypeSourceMicrosoftTeamsAuthenticateViaMicrosoft

	return SourceMicrosoftTeamsAuthenticationMechanism{
		SourceMicrosoftTeamsAuthenticateViaMicrosoft: &sourceMicrosoftTeamsAuthenticateViaMicrosoft,
		Type: typ,
	}
}

func (u *SourceMicrosoftTeamsAuthenticationMechanism) UnmarshalJSON(data []byte) error {

	var sourceMicrosoftTeamsAuthenticateViaMicrosoft SourceMicrosoftTeamsAuthenticateViaMicrosoft = SourceMicrosoftTeamsAuthenticateViaMicrosoft{}
	if err := utils.UnmarshalJSON(data, &sourceMicrosoftTeamsAuthenticateViaMicrosoft, "", true, true); err == nil {
		u.SourceMicrosoftTeamsAuthenticateViaMicrosoft = &sourceMicrosoftTeamsAuthenticateViaMicrosoft
		u.Type = SourceMicrosoftTeamsAuthenticationMechanismTypeSourceMicrosoftTeamsAuthenticateViaMicrosoft
		return nil
	}

	var sourceMicrosoftTeamsAuthenticateViaMicrosoftOAuth20 SourceMicrosoftTeamsAuthenticateViaMicrosoftOAuth20 = SourceMicrosoftTeamsAuthenticateViaMicrosoftOAuth20{}
	if err := utils.UnmarshalJSON(data, &sourceMicrosoftTeamsAuthenticateViaMicrosoftOAuth20, "", true, true); err == nil {
		u.SourceMicrosoftTeamsAuthenticateViaMicrosoftOAuth20 = &sourceMicrosoftTeamsAuthenticateViaMicrosoftOAuth20
		u.Type = SourceMicrosoftTeamsAuthenticationMechanismTypeSourceMicrosoftTeamsAuthenticateViaMicrosoftOAuth20
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for SourceMicrosoftTeamsAuthenticationMechanism", string(data))
}

func (u SourceMicrosoftTeamsAuthenticationMechanism) MarshalJSON() ([]byte, error) {
	if u.SourceMicrosoftTeamsAuthenticateViaMicrosoftOAuth20 != nil {
		return utils.MarshalJSON(u.SourceMicrosoftTeamsAuthenticateViaMicrosoftOAuth20, "", true)
	}

	if u.SourceMicrosoftTeamsAuthenticateViaMicrosoft != nil {
		return utils.MarshalJSON(u.SourceMicrosoftTeamsAuthenticateViaMicrosoft, "", true)
	}

	return nil, errors.New("could not marshal union type SourceMicrosoftTeamsAuthenticationMechanism: all fields are null")
}

type MicrosoftTeams string

const (
	MicrosoftTeamsMicrosoftTeams MicrosoftTeams = "microsoft-teams"
)

func (e MicrosoftTeams) ToPointer() *MicrosoftTeams {
	return &e
}
func (e *MicrosoftTeams) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "microsoft-teams":
		*e = MicrosoftTeams(v)
		return nil
	default:
		return fmt.Errorf("invalid value for MicrosoftTeams: %v", v)
	}
}

type SourceMicrosoftTeams struct {
	// Choose how to authenticate to Microsoft
	Credentials *SourceMicrosoftTeamsAuthenticationMechanism `json:"credentials,omitempty"`
	// Specifies the length of time over which the Team Device Report stream is aggregated. The supported values are: D7, D30, D90, and D180.
	Period     string         `json:"period"`
	sourceType MicrosoftTeams `const:"microsoft-teams" json:"sourceType"`
}

func (s SourceMicrosoftTeams) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceMicrosoftTeams) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SourceMicrosoftTeams) GetCredentials() *SourceMicrosoftTeamsAuthenticationMechanism {
	if o == nil {
		return nil
	}
	return o.Credentials
}

func (o *SourceMicrosoftTeams) GetPeriod() string {
	if o == nil {
		return ""
	}
	return o.Period
}

func (o *SourceMicrosoftTeams) GetSourceType() MicrosoftTeams {
	return MicrosoftTeamsMicrosoftTeams
}
