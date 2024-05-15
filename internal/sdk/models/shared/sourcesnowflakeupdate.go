// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/internal/utils"
)

type SourceSnowflakeUpdateSchemasAuthType string

const (
	SourceSnowflakeUpdateSchemasAuthTypeUsernamePassword SourceSnowflakeUpdateSchemasAuthType = "username/password"
)

func (e SourceSnowflakeUpdateSchemasAuthType) ToPointer() *SourceSnowflakeUpdateSchemasAuthType {
	return &e
}

func (e *SourceSnowflakeUpdateSchemasAuthType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "username/password":
		*e = SourceSnowflakeUpdateSchemasAuthType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceSnowflakeUpdateSchemasAuthType: %v", v)
	}
}

type SourceSnowflakeUpdateUsernameAndPassword struct {
	authType SourceSnowflakeUpdateSchemasAuthType `const:"username/password" json:"auth_type"`
	// The password associated with the username.
	Password string `json:"password"`
	// The username you created to allow Airbyte to access the database.
	Username string `json:"username"`
}

func (s SourceSnowflakeUpdateUsernameAndPassword) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceSnowflakeUpdateUsernameAndPassword) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *SourceSnowflakeUpdateUsernameAndPassword) GetAuthType() SourceSnowflakeUpdateSchemasAuthType {
	return SourceSnowflakeUpdateSchemasAuthTypeUsernamePassword
}

func (o *SourceSnowflakeUpdateUsernameAndPassword) GetPassword() string {
	if o == nil {
		return ""
	}
	return o.Password
}

func (o *SourceSnowflakeUpdateUsernameAndPassword) GetUsername() string {
	if o == nil {
		return ""
	}
	return o.Username
}

type SourceSnowflakeUpdateAuthType string

const (
	SourceSnowflakeUpdateAuthTypeOAuth SourceSnowflakeUpdateAuthType = "OAuth"
)

func (e SourceSnowflakeUpdateAuthType) ToPointer() *SourceSnowflakeUpdateAuthType {
	return &e
}

func (e *SourceSnowflakeUpdateAuthType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "OAuth":
		*e = SourceSnowflakeUpdateAuthType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceSnowflakeUpdateAuthType: %v", v)
	}
}

type SourceSnowflakeUpdateOAuth20 struct {
	// Access Token for making authenticated requests.
	AccessToken *string                       `json:"access_token,omitempty"`
	authType    SourceSnowflakeUpdateAuthType `const:"OAuth" json:"auth_type"`
	// The Client ID of your Snowflake developer application.
	ClientID string `json:"client_id"`
	// The Client Secret of your Snowflake developer application.
	ClientSecret string `json:"client_secret"`
	// Refresh Token for making authenticated requests.
	RefreshToken *string `json:"refresh_token,omitempty"`
}

func (s SourceSnowflakeUpdateOAuth20) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceSnowflakeUpdateOAuth20) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *SourceSnowflakeUpdateOAuth20) GetAccessToken() *string {
	if o == nil {
		return nil
	}
	return o.AccessToken
}

func (o *SourceSnowflakeUpdateOAuth20) GetAuthType() SourceSnowflakeUpdateAuthType {
	return SourceSnowflakeUpdateAuthTypeOAuth
}

func (o *SourceSnowflakeUpdateOAuth20) GetClientID() string {
	if o == nil {
		return ""
	}
	return o.ClientID
}

func (o *SourceSnowflakeUpdateOAuth20) GetClientSecret() string {
	if o == nil {
		return ""
	}
	return o.ClientSecret
}

func (o *SourceSnowflakeUpdateOAuth20) GetRefreshToken() *string {
	if o == nil {
		return nil
	}
	return o.RefreshToken
}

type SourceSnowflakeUpdateAuthorizationMethodType string

const (
	SourceSnowflakeUpdateAuthorizationMethodTypeSourceSnowflakeUpdateOAuth20             SourceSnowflakeUpdateAuthorizationMethodType = "source-snowflake-update_OAuth2.0"
	SourceSnowflakeUpdateAuthorizationMethodTypeSourceSnowflakeUpdateUsernameAndPassword SourceSnowflakeUpdateAuthorizationMethodType = "source-snowflake-update_Username and Password"
)

type SourceSnowflakeUpdateAuthorizationMethod struct {
	SourceSnowflakeUpdateOAuth20             *SourceSnowflakeUpdateOAuth20
	SourceSnowflakeUpdateUsernameAndPassword *SourceSnowflakeUpdateUsernameAndPassword

	Type SourceSnowflakeUpdateAuthorizationMethodType
}

func CreateSourceSnowflakeUpdateAuthorizationMethodSourceSnowflakeUpdateOAuth20(sourceSnowflakeUpdateOAuth20 SourceSnowflakeUpdateOAuth20) SourceSnowflakeUpdateAuthorizationMethod {
	typ := SourceSnowflakeUpdateAuthorizationMethodTypeSourceSnowflakeUpdateOAuth20

	return SourceSnowflakeUpdateAuthorizationMethod{
		SourceSnowflakeUpdateOAuth20: &sourceSnowflakeUpdateOAuth20,
		Type:                         typ,
	}
}

func CreateSourceSnowflakeUpdateAuthorizationMethodSourceSnowflakeUpdateUsernameAndPassword(sourceSnowflakeUpdateUsernameAndPassword SourceSnowflakeUpdateUsernameAndPassword) SourceSnowflakeUpdateAuthorizationMethod {
	typ := SourceSnowflakeUpdateAuthorizationMethodTypeSourceSnowflakeUpdateUsernameAndPassword

	return SourceSnowflakeUpdateAuthorizationMethod{
		SourceSnowflakeUpdateUsernameAndPassword: &sourceSnowflakeUpdateUsernameAndPassword,
		Type:                                     typ,
	}
}

func (u *SourceSnowflakeUpdateAuthorizationMethod) UnmarshalJSON(data []byte) error {

	var sourceSnowflakeUpdateUsernameAndPassword SourceSnowflakeUpdateUsernameAndPassword = SourceSnowflakeUpdateUsernameAndPassword{}
	if err := utils.UnmarshalJSON(data, &sourceSnowflakeUpdateUsernameAndPassword, "", true, true); err == nil {
		u.SourceSnowflakeUpdateUsernameAndPassword = &sourceSnowflakeUpdateUsernameAndPassword
		u.Type = SourceSnowflakeUpdateAuthorizationMethodTypeSourceSnowflakeUpdateUsernameAndPassword
		return nil
	}

	var sourceSnowflakeUpdateOAuth20 SourceSnowflakeUpdateOAuth20 = SourceSnowflakeUpdateOAuth20{}
	if err := utils.UnmarshalJSON(data, &sourceSnowflakeUpdateOAuth20, "", true, true); err == nil {
		u.SourceSnowflakeUpdateOAuth20 = &sourceSnowflakeUpdateOAuth20
		u.Type = SourceSnowflakeUpdateAuthorizationMethodTypeSourceSnowflakeUpdateOAuth20
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u SourceSnowflakeUpdateAuthorizationMethod) MarshalJSON() ([]byte, error) {
	if u.SourceSnowflakeUpdateOAuth20 != nil {
		return utils.MarshalJSON(u.SourceSnowflakeUpdateOAuth20, "", true)
	}

	if u.SourceSnowflakeUpdateUsernameAndPassword != nil {
		return utils.MarshalJSON(u.SourceSnowflakeUpdateUsernameAndPassword, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type SourceSnowflakeUpdate struct {
	Credentials *SourceSnowflakeUpdateAuthorizationMethod `json:"credentials,omitempty"`
	// The database you created for Airbyte to access data.
	Database string `json:"database"`
	// The host domain of the snowflake instance (must include the account, region, cloud environment, and end with snowflakecomputing.com).
	Host string `json:"host"`
	// Additional properties to pass to the JDBC URL string when connecting to the database formatted as 'key=value' pairs separated by the symbol '&'. (example: key1=value1&key2=value2&key3=value3).
	JdbcURLParams *string `json:"jdbc_url_params,omitempty"`
	// The role you created for Airbyte to access Snowflake.
	Role string `json:"role"`
	// The source Snowflake schema tables. Leave empty to access tables from multiple schemas.
	Schema *string `json:"schema,omitempty"`
	// The warehouse you created for Airbyte to access data.
	Warehouse string `json:"warehouse"`
}

func (o *SourceSnowflakeUpdate) GetCredentials() *SourceSnowflakeUpdateAuthorizationMethod {
	if o == nil {
		return nil
	}
	return o.Credentials
}

func (o *SourceSnowflakeUpdate) GetDatabase() string {
	if o == nil {
		return ""
	}
	return o.Database
}

func (o *SourceSnowflakeUpdate) GetHost() string {
	if o == nil {
		return ""
	}
	return o.Host
}

func (o *SourceSnowflakeUpdate) GetJdbcURLParams() *string {
	if o == nil {
		return nil
	}
	return o.JdbcURLParams
}

func (o *SourceSnowflakeUpdate) GetRole() string {
	if o == nil {
		return ""
	}
	return o.Role
}

func (o *SourceSnowflakeUpdate) GetSchema() *string {
	if o == nil {
		return nil
	}
	return o.Schema
}

func (o *SourceSnowflakeUpdate) GetWarehouse() string {
	if o == nil {
		return ""
	}
	return o.Warehouse
}