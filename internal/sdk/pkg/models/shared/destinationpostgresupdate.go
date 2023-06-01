// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
)

type DestinationPostgresUpdateSSLModesVerifyFullMode string

const (
	DestinationPostgresUpdateSSLModesVerifyFullModeVerifyFull DestinationPostgresUpdateSSLModesVerifyFullMode = "verify-full"
)

func (e DestinationPostgresUpdateSSLModesVerifyFullMode) ToPointer() *DestinationPostgresUpdateSSLModesVerifyFullMode {
	return &e
}

func (e *DestinationPostgresUpdateSSLModesVerifyFullMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "verify-full":
		*e = DestinationPostgresUpdateSSLModesVerifyFullMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationPostgresUpdateSSLModesVerifyFullMode: %v", v)
	}
}

// DestinationPostgresUpdateSSLModesVerifyFull - Verify-full SSL mode.
type DestinationPostgresUpdateSSLModesVerifyFull struct {
	// CA certificate
	CaCertificate string `json:"ca_certificate"`
	// Client certificate
	ClientCertificate string `json:"client_certificate"`
	// Client key
	ClientKey string `json:"client_key"`
	// Password for keystorage. This field is optional. If you do not add it - the password will be generated automatically.
	ClientKeyPassword *string                                         `json:"client_key_password,omitempty"`
	Mode              DestinationPostgresUpdateSSLModesVerifyFullMode `json:"mode"`
}

type DestinationPostgresUpdateSSLModesVerifyCaMode string

const (
	DestinationPostgresUpdateSSLModesVerifyCaModeVerifyCa DestinationPostgresUpdateSSLModesVerifyCaMode = "verify-ca"
)

func (e DestinationPostgresUpdateSSLModesVerifyCaMode) ToPointer() *DestinationPostgresUpdateSSLModesVerifyCaMode {
	return &e
}

func (e *DestinationPostgresUpdateSSLModesVerifyCaMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "verify-ca":
		*e = DestinationPostgresUpdateSSLModesVerifyCaMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationPostgresUpdateSSLModesVerifyCaMode: %v", v)
	}
}

// DestinationPostgresUpdateSSLModesVerifyCa - Verify-ca SSL mode.
type DestinationPostgresUpdateSSLModesVerifyCa struct {
	// CA certificate
	CaCertificate string `json:"ca_certificate"`
	// Password for keystorage. This field is optional. If you do not add it - the password will be generated automatically.
	ClientKeyPassword *string                                       `json:"client_key_password,omitempty"`
	Mode              DestinationPostgresUpdateSSLModesVerifyCaMode `json:"mode"`
}

type DestinationPostgresUpdateSSLModesRequireMode string

const (
	DestinationPostgresUpdateSSLModesRequireModeRequire DestinationPostgresUpdateSSLModesRequireMode = "require"
)

func (e DestinationPostgresUpdateSSLModesRequireMode) ToPointer() *DestinationPostgresUpdateSSLModesRequireMode {
	return &e
}

func (e *DestinationPostgresUpdateSSLModesRequireMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "require":
		*e = DestinationPostgresUpdateSSLModesRequireMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationPostgresUpdateSSLModesRequireMode: %v", v)
	}
}

// DestinationPostgresUpdateSSLModesRequire - Require SSL mode.
type DestinationPostgresUpdateSSLModesRequire struct {
	Mode DestinationPostgresUpdateSSLModesRequireMode `json:"mode"`
}

type DestinationPostgresUpdateSSLModesPreferMode string

const (
	DestinationPostgresUpdateSSLModesPreferModePrefer DestinationPostgresUpdateSSLModesPreferMode = "prefer"
)

func (e DestinationPostgresUpdateSSLModesPreferMode) ToPointer() *DestinationPostgresUpdateSSLModesPreferMode {
	return &e
}

func (e *DestinationPostgresUpdateSSLModesPreferMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "prefer":
		*e = DestinationPostgresUpdateSSLModesPreferMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationPostgresUpdateSSLModesPreferMode: %v", v)
	}
}

// DestinationPostgresUpdateSSLModesPrefer - Prefer SSL mode.
type DestinationPostgresUpdateSSLModesPrefer struct {
	Mode DestinationPostgresUpdateSSLModesPreferMode `json:"mode"`
}

type DestinationPostgresUpdateSSLModesAllowMode string

const (
	DestinationPostgresUpdateSSLModesAllowModeAllow DestinationPostgresUpdateSSLModesAllowMode = "allow"
)

func (e DestinationPostgresUpdateSSLModesAllowMode) ToPointer() *DestinationPostgresUpdateSSLModesAllowMode {
	return &e
}

func (e *DestinationPostgresUpdateSSLModesAllowMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "allow":
		*e = DestinationPostgresUpdateSSLModesAllowMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationPostgresUpdateSSLModesAllowMode: %v", v)
	}
}

// DestinationPostgresUpdateSSLModesAllow - Allow SSL mode.
type DestinationPostgresUpdateSSLModesAllow struct {
	Mode DestinationPostgresUpdateSSLModesAllowMode `json:"mode"`
}

type DestinationPostgresUpdateSSLModesDisableMode string

const (
	DestinationPostgresUpdateSSLModesDisableModeDisable DestinationPostgresUpdateSSLModesDisableMode = "disable"
)

func (e DestinationPostgresUpdateSSLModesDisableMode) ToPointer() *DestinationPostgresUpdateSSLModesDisableMode {
	return &e
}

func (e *DestinationPostgresUpdateSSLModesDisableMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "disable":
		*e = DestinationPostgresUpdateSSLModesDisableMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationPostgresUpdateSSLModesDisableMode: %v", v)
	}
}

// DestinationPostgresUpdateSSLModesDisable - Disable SSL.
type DestinationPostgresUpdateSSLModesDisable struct {
	Mode DestinationPostgresUpdateSSLModesDisableMode `json:"mode"`
}

type DestinationPostgresUpdateSSLModesType string

const (
	DestinationPostgresUpdateSSLModesTypeDestinationPostgresUpdateSSLModesDisable    DestinationPostgresUpdateSSLModesType = "destination-postgres-update_SSL modes_disable"
	DestinationPostgresUpdateSSLModesTypeDestinationPostgresUpdateSSLModesAllow      DestinationPostgresUpdateSSLModesType = "destination-postgres-update_SSL modes_allow"
	DestinationPostgresUpdateSSLModesTypeDestinationPostgresUpdateSSLModesPrefer     DestinationPostgresUpdateSSLModesType = "destination-postgres-update_SSL modes_prefer"
	DestinationPostgresUpdateSSLModesTypeDestinationPostgresUpdateSSLModesRequire    DestinationPostgresUpdateSSLModesType = "destination-postgres-update_SSL modes_require"
	DestinationPostgresUpdateSSLModesTypeDestinationPostgresUpdateSSLModesVerifyCa   DestinationPostgresUpdateSSLModesType = "destination-postgres-update_SSL modes_verify-ca"
	DestinationPostgresUpdateSSLModesTypeDestinationPostgresUpdateSSLModesVerifyFull DestinationPostgresUpdateSSLModesType = "destination-postgres-update_SSL modes_verify-full"
)

type DestinationPostgresUpdateSSLModes struct {
	DestinationPostgresUpdateSSLModesDisable    *DestinationPostgresUpdateSSLModesDisable
	DestinationPostgresUpdateSSLModesAllow      *DestinationPostgresUpdateSSLModesAllow
	DestinationPostgresUpdateSSLModesPrefer     *DestinationPostgresUpdateSSLModesPrefer
	DestinationPostgresUpdateSSLModesRequire    *DestinationPostgresUpdateSSLModesRequire
	DestinationPostgresUpdateSSLModesVerifyCa   *DestinationPostgresUpdateSSLModesVerifyCa
	DestinationPostgresUpdateSSLModesVerifyFull *DestinationPostgresUpdateSSLModesVerifyFull

	Type DestinationPostgresUpdateSSLModesType
}

func CreateDestinationPostgresUpdateSSLModesDestinationPostgresUpdateSSLModesDisable(destinationPostgresUpdateSSLModesDisable DestinationPostgresUpdateSSLModesDisable) DestinationPostgresUpdateSSLModes {
	typ := DestinationPostgresUpdateSSLModesTypeDestinationPostgresUpdateSSLModesDisable

	return DestinationPostgresUpdateSSLModes{
		DestinationPostgresUpdateSSLModesDisable: &destinationPostgresUpdateSSLModesDisable,
		Type:                                     typ,
	}
}

func CreateDestinationPostgresUpdateSSLModesDestinationPostgresUpdateSSLModesAllow(destinationPostgresUpdateSSLModesAllow DestinationPostgresUpdateSSLModesAllow) DestinationPostgresUpdateSSLModes {
	typ := DestinationPostgresUpdateSSLModesTypeDestinationPostgresUpdateSSLModesAllow

	return DestinationPostgresUpdateSSLModes{
		DestinationPostgresUpdateSSLModesAllow: &destinationPostgresUpdateSSLModesAllow,
		Type:                                   typ,
	}
}

func CreateDestinationPostgresUpdateSSLModesDestinationPostgresUpdateSSLModesPrefer(destinationPostgresUpdateSSLModesPrefer DestinationPostgresUpdateSSLModesPrefer) DestinationPostgresUpdateSSLModes {
	typ := DestinationPostgresUpdateSSLModesTypeDestinationPostgresUpdateSSLModesPrefer

	return DestinationPostgresUpdateSSLModes{
		DestinationPostgresUpdateSSLModesPrefer: &destinationPostgresUpdateSSLModesPrefer,
		Type:                                    typ,
	}
}

func CreateDestinationPostgresUpdateSSLModesDestinationPostgresUpdateSSLModesRequire(destinationPostgresUpdateSSLModesRequire DestinationPostgresUpdateSSLModesRequire) DestinationPostgresUpdateSSLModes {
	typ := DestinationPostgresUpdateSSLModesTypeDestinationPostgresUpdateSSLModesRequire

	return DestinationPostgresUpdateSSLModes{
		DestinationPostgresUpdateSSLModesRequire: &destinationPostgresUpdateSSLModesRequire,
		Type:                                     typ,
	}
}

func CreateDestinationPostgresUpdateSSLModesDestinationPostgresUpdateSSLModesVerifyCa(destinationPostgresUpdateSSLModesVerifyCa DestinationPostgresUpdateSSLModesVerifyCa) DestinationPostgresUpdateSSLModes {
	typ := DestinationPostgresUpdateSSLModesTypeDestinationPostgresUpdateSSLModesVerifyCa

	return DestinationPostgresUpdateSSLModes{
		DestinationPostgresUpdateSSLModesVerifyCa: &destinationPostgresUpdateSSLModesVerifyCa,
		Type: typ,
	}
}

func CreateDestinationPostgresUpdateSSLModesDestinationPostgresUpdateSSLModesVerifyFull(destinationPostgresUpdateSSLModesVerifyFull DestinationPostgresUpdateSSLModesVerifyFull) DestinationPostgresUpdateSSLModes {
	typ := DestinationPostgresUpdateSSLModesTypeDestinationPostgresUpdateSSLModesVerifyFull

	return DestinationPostgresUpdateSSLModes{
		DestinationPostgresUpdateSSLModesVerifyFull: &destinationPostgresUpdateSSLModesVerifyFull,
		Type: typ,
	}
}

func (u *DestinationPostgresUpdateSSLModes) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	destinationPostgresUpdateSSLModesDisable := new(DestinationPostgresUpdateSSLModesDisable)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&destinationPostgresUpdateSSLModesDisable); err == nil {
		u.DestinationPostgresUpdateSSLModesDisable = destinationPostgresUpdateSSLModesDisable
		u.Type = DestinationPostgresUpdateSSLModesTypeDestinationPostgresUpdateSSLModesDisable
		return nil
	}

	destinationPostgresUpdateSSLModesAllow := new(DestinationPostgresUpdateSSLModesAllow)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&destinationPostgresUpdateSSLModesAllow); err == nil {
		u.DestinationPostgresUpdateSSLModesAllow = destinationPostgresUpdateSSLModesAllow
		u.Type = DestinationPostgresUpdateSSLModesTypeDestinationPostgresUpdateSSLModesAllow
		return nil
	}

	destinationPostgresUpdateSSLModesPrefer := new(DestinationPostgresUpdateSSLModesPrefer)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&destinationPostgresUpdateSSLModesPrefer); err == nil {
		u.DestinationPostgresUpdateSSLModesPrefer = destinationPostgresUpdateSSLModesPrefer
		u.Type = DestinationPostgresUpdateSSLModesTypeDestinationPostgresUpdateSSLModesPrefer
		return nil
	}

	destinationPostgresUpdateSSLModesRequire := new(DestinationPostgresUpdateSSLModesRequire)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&destinationPostgresUpdateSSLModesRequire); err == nil {
		u.DestinationPostgresUpdateSSLModesRequire = destinationPostgresUpdateSSLModesRequire
		u.Type = DestinationPostgresUpdateSSLModesTypeDestinationPostgresUpdateSSLModesRequire
		return nil
	}

	destinationPostgresUpdateSSLModesVerifyCa := new(DestinationPostgresUpdateSSLModesVerifyCa)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&destinationPostgresUpdateSSLModesVerifyCa); err == nil {
		u.DestinationPostgresUpdateSSLModesVerifyCa = destinationPostgresUpdateSSLModesVerifyCa
		u.Type = DestinationPostgresUpdateSSLModesTypeDestinationPostgresUpdateSSLModesVerifyCa
		return nil
	}

	destinationPostgresUpdateSSLModesVerifyFull := new(DestinationPostgresUpdateSSLModesVerifyFull)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&destinationPostgresUpdateSSLModesVerifyFull); err == nil {
		u.DestinationPostgresUpdateSSLModesVerifyFull = destinationPostgresUpdateSSLModesVerifyFull
		u.Type = DestinationPostgresUpdateSSLModesTypeDestinationPostgresUpdateSSLModesVerifyFull
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u DestinationPostgresUpdateSSLModes) MarshalJSON() ([]byte, error) {
	if u.DestinationPostgresUpdateSSLModesDisable != nil {
		return json.Marshal(u.DestinationPostgresUpdateSSLModesDisable)
	}

	if u.DestinationPostgresUpdateSSLModesAllow != nil {
		return json.Marshal(u.DestinationPostgresUpdateSSLModesAllow)
	}

	if u.DestinationPostgresUpdateSSLModesPrefer != nil {
		return json.Marshal(u.DestinationPostgresUpdateSSLModesPrefer)
	}

	if u.DestinationPostgresUpdateSSLModesRequire != nil {
		return json.Marshal(u.DestinationPostgresUpdateSSLModesRequire)
	}

	if u.DestinationPostgresUpdateSSLModesVerifyCa != nil {
		return json.Marshal(u.DestinationPostgresUpdateSSLModesVerifyCa)
	}

	if u.DestinationPostgresUpdateSSLModesVerifyFull != nil {
		return json.Marshal(u.DestinationPostgresUpdateSSLModesVerifyFull)
	}

	return nil, nil
}

// DestinationPostgresUpdateSSHTunnelMethodPasswordAuthenticationTunnelMethod - Connect through a jump server tunnel host using username and password authentication
type DestinationPostgresUpdateSSHTunnelMethodPasswordAuthenticationTunnelMethod string

const (
	DestinationPostgresUpdateSSHTunnelMethodPasswordAuthenticationTunnelMethodSSHPasswordAuth DestinationPostgresUpdateSSHTunnelMethodPasswordAuthenticationTunnelMethod = "SSH_PASSWORD_AUTH"
)

func (e DestinationPostgresUpdateSSHTunnelMethodPasswordAuthenticationTunnelMethod) ToPointer() *DestinationPostgresUpdateSSHTunnelMethodPasswordAuthenticationTunnelMethod {
	return &e
}

func (e *DestinationPostgresUpdateSSHTunnelMethodPasswordAuthenticationTunnelMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "SSH_PASSWORD_AUTH":
		*e = DestinationPostgresUpdateSSHTunnelMethodPasswordAuthenticationTunnelMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationPostgresUpdateSSHTunnelMethodPasswordAuthenticationTunnelMethod: %v", v)
	}
}

// DestinationPostgresUpdateSSHTunnelMethodPasswordAuthentication - Whether to initiate an SSH tunnel before connecting to the database, and if so, which kind of authentication to use.
type DestinationPostgresUpdateSSHTunnelMethodPasswordAuthentication struct {
	// Hostname of the jump server host that allows inbound ssh tunnel.
	TunnelHost string `json:"tunnel_host"`
	// Connect through a jump server tunnel host using username and password authentication
	TunnelMethod DestinationPostgresUpdateSSHTunnelMethodPasswordAuthenticationTunnelMethod `json:"tunnel_method"`
	// Port on the proxy/jump server that accepts inbound ssh connections.
	TunnelPort int64 `json:"tunnel_port"`
	// OS-level username for logging into the jump server host
	TunnelUser string `json:"tunnel_user"`
	// OS-level password for logging into the jump server host
	TunnelUserPassword string `json:"tunnel_user_password"`
}

// DestinationPostgresUpdateSSHTunnelMethodSSHKeyAuthenticationTunnelMethod - Connect through a jump server tunnel host using username and ssh key
type DestinationPostgresUpdateSSHTunnelMethodSSHKeyAuthenticationTunnelMethod string

const (
	DestinationPostgresUpdateSSHTunnelMethodSSHKeyAuthenticationTunnelMethodSSHKeyAuth DestinationPostgresUpdateSSHTunnelMethodSSHKeyAuthenticationTunnelMethod = "SSH_KEY_AUTH"
)

func (e DestinationPostgresUpdateSSHTunnelMethodSSHKeyAuthenticationTunnelMethod) ToPointer() *DestinationPostgresUpdateSSHTunnelMethodSSHKeyAuthenticationTunnelMethod {
	return &e
}

func (e *DestinationPostgresUpdateSSHTunnelMethodSSHKeyAuthenticationTunnelMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "SSH_KEY_AUTH":
		*e = DestinationPostgresUpdateSSHTunnelMethodSSHKeyAuthenticationTunnelMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationPostgresUpdateSSHTunnelMethodSSHKeyAuthenticationTunnelMethod: %v", v)
	}
}

// DestinationPostgresUpdateSSHTunnelMethodSSHKeyAuthentication - Whether to initiate an SSH tunnel before connecting to the database, and if so, which kind of authentication to use.
type DestinationPostgresUpdateSSHTunnelMethodSSHKeyAuthentication struct {
	// OS-level user account ssh key credentials in RSA PEM format ( created with ssh-keygen -t rsa -m PEM -f myuser_rsa )
	SSHKey string `json:"ssh_key"`
	// Hostname of the jump server host that allows inbound ssh tunnel.
	TunnelHost string `json:"tunnel_host"`
	// Connect through a jump server tunnel host using username and ssh key
	TunnelMethod DestinationPostgresUpdateSSHTunnelMethodSSHKeyAuthenticationTunnelMethod `json:"tunnel_method"`
	// Port on the proxy/jump server that accepts inbound ssh connections.
	TunnelPort int64 `json:"tunnel_port"`
	// OS-level username for logging into the jump server host.
	TunnelUser string `json:"tunnel_user"`
}

// DestinationPostgresUpdateSSHTunnelMethodNoTunnelTunnelMethod - No ssh tunnel needed to connect to database
type DestinationPostgresUpdateSSHTunnelMethodNoTunnelTunnelMethod string

const (
	DestinationPostgresUpdateSSHTunnelMethodNoTunnelTunnelMethodNoTunnel DestinationPostgresUpdateSSHTunnelMethodNoTunnelTunnelMethod = "NO_TUNNEL"
)

func (e DestinationPostgresUpdateSSHTunnelMethodNoTunnelTunnelMethod) ToPointer() *DestinationPostgresUpdateSSHTunnelMethodNoTunnelTunnelMethod {
	return &e
}

func (e *DestinationPostgresUpdateSSHTunnelMethodNoTunnelTunnelMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "NO_TUNNEL":
		*e = DestinationPostgresUpdateSSHTunnelMethodNoTunnelTunnelMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationPostgresUpdateSSHTunnelMethodNoTunnelTunnelMethod: %v", v)
	}
}

// DestinationPostgresUpdateSSHTunnelMethodNoTunnel - Whether to initiate an SSH tunnel before connecting to the database, and if so, which kind of authentication to use.
type DestinationPostgresUpdateSSHTunnelMethodNoTunnel struct {
	// No ssh tunnel needed to connect to database
	TunnelMethod DestinationPostgresUpdateSSHTunnelMethodNoTunnelTunnelMethod `json:"tunnel_method"`
}

type DestinationPostgresUpdateSSHTunnelMethodType string

const (
	DestinationPostgresUpdateSSHTunnelMethodTypeDestinationPostgresUpdateSSHTunnelMethodNoTunnel               DestinationPostgresUpdateSSHTunnelMethodType = "destination-postgres-update_SSH Tunnel Method_No Tunnel"
	DestinationPostgresUpdateSSHTunnelMethodTypeDestinationPostgresUpdateSSHTunnelMethodSSHKeyAuthentication   DestinationPostgresUpdateSSHTunnelMethodType = "destination-postgres-update_SSH Tunnel Method_SSH Key Authentication"
	DestinationPostgresUpdateSSHTunnelMethodTypeDestinationPostgresUpdateSSHTunnelMethodPasswordAuthentication DestinationPostgresUpdateSSHTunnelMethodType = "destination-postgres-update_SSH Tunnel Method_Password Authentication"
)

type DestinationPostgresUpdateSSHTunnelMethod struct {
	DestinationPostgresUpdateSSHTunnelMethodNoTunnel               *DestinationPostgresUpdateSSHTunnelMethodNoTunnel
	DestinationPostgresUpdateSSHTunnelMethodSSHKeyAuthentication   *DestinationPostgresUpdateSSHTunnelMethodSSHKeyAuthentication
	DestinationPostgresUpdateSSHTunnelMethodPasswordAuthentication *DestinationPostgresUpdateSSHTunnelMethodPasswordAuthentication

	Type DestinationPostgresUpdateSSHTunnelMethodType
}

func CreateDestinationPostgresUpdateSSHTunnelMethodDestinationPostgresUpdateSSHTunnelMethodNoTunnel(destinationPostgresUpdateSSHTunnelMethodNoTunnel DestinationPostgresUpdateSSHTunnelMethodNoTunnel) DestinationPostgresUpdateSSHTunnelMethod {
	typ := DestinationPostgresUpdateSSHTunnelMethodTypeDestinationPostgresUpdateSSHTunnelMethodNoTunnel

	return DestinationPostgresUpdateSSHTunnelMethod{
		DestinationPostgresUpdateSSHTunnelMethodNoTunnel: &destinationPostgresUpdateSSHTunnelMethodNoTunnel,
		Type: typ,
	}
}

func CreateDestinationPostgresUpdateSSHTunnelMethodDestinationPostgresUpdateSSHTunnelMethodSSHKeyAuthentication(destinationPostgresUpdateSSHTunnelMethodSSHKeyAuthentication DestinationPostgresUpdateSSHTunnelMethodSSHKeyAuthentication) DestinationPostgresUpdateSSHTunnelMethod {
	typ := DestinationPostgresUpdateSSHTunnelMethodTypeDestinationPostgresUpdateSSHTunnelMethodSSHKeyAuthentication

	return DestinationPostgresUpdateSSHTunnelMethod{
		DestinationPostgresUpdateSSHTunnelMethodSSHKeyAuthentication: &destinationPostgresUpdateSSHTunnelMethodSSHKeyAuthentication,
		Type: typ,
	}
}

func CreateDestinationPostgresUpdateSSHTunnelMethodDestinationPostgresUpdateSSHTunnelMethodPasswordAuthentication(destinationPostgresUpdateSSHTunnelMethodPasswordAuthentication DestinationPostgresUpdateSSHTunnelMethodPasswordAuthentication) DestinationPostgresUpdateSSHTunnelMethod {
	typ := DestinationPostgresUpdateSSHTunnelMethodTypeDestinationPostgresUpdateSSHTunnelMethodPasswordAuthentication

	return DestinationPostgresUpdateSSHTunnelMethod{
		DestinationPostgresUpdateSSHTunnelMethodPasswordAuthentication: &destinationPostgresUpdateSSHTunnelMethodPasswordAuthentication,
		Type: typ,
	}
}

func (u *DestinationPostgresUpdateSSHTunnelMethod) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	destinationPostgresUpdateSSHTunnelMethodNoTunnel := new(DestinationPostgresUpdateSSHTunnelMethodNoTunnel)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&destinationPostgresUpdateSSHTunnelMethodNoTunnel); err == nil {
		u.DestinationPostgresUpdateSSHTunnelMethodNoTunnel = destinationPostgresUpdateSSHTunnelMethodNoTunnel
		u.Type = DestinationPostgresUpdateSSHTunnelMethodTypeDestinationPostgresUpdateSSHTunnelMethodNoTunnel
		return nil
	}

	destinationPostgresUpdateSSHTunnelMethodSSHKeyAuthentication := new(DestinationPostgresUpdateSSHTunnelMethodSSHKeyAuthentication)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&destinationPostgresUpdateSSHTunnelMethodSSHKeyAuthentication); err == nil {
		u.DestinationPostgresUpdateSSHTunnelMethodSSHKeyAuthentication = destinationPostgresUpdateSSHTunnelMethodSSHKeyAuthentication
		u.Type = DestinationPostgresUpdateSSHTunnelMethodTypeDestinationPostgresUpdateSSHTunnelMethodSSHKeyAuthentication
		return nil
	}

	destinationPostgresUpdateSSHTunnelMethodPasswordAuthentication := new(DestinationPostgresUpdateSSHTunnelMethodPasswordAuthentication)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&destinationPostgresUpdateSSHTunnelMethodPasswordAuthentication); err == nil {
		u.DestinationPostgresUpdateSSHTunnelMethodPasswordAuthentication = destinationPostgresUpdateSSHTunnelMethodPasswordAuthentication
		u.Type = DestinationPostgresUpdateSSHTunnelMethodTypeDestinationPostgresUpdateSSHTunnelMethodPasswordAuthentication
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u DestinationPostgresUpdateSSHTunnelMethod) MarshalJSON() ([]byte, error) {
	if u.DestinationPostgresUpdateSSHTunnelMethodNoTunnel != nil {
		return json.Marshal(u.DestinationPostgresUpdateSSHTunnelMethodNoTunnel)
	}

	if u.DestinationPostgresUpdateSSHTunnelMethodSSHKeyAuthentication != nil {
		return json.Marshal(u.DestinationPostgresUpdateSSHTunnelMethodSSHKeyAuthentication)
	}

	if u.DestinationPostgresUpdateSSHTunnelMethodPasswordAuthentication != nil {
		return json.Marshal(u.DestinationPostgresUpdateSSHTunnelMethodPasswordAuthentication)
	}

	return nil, nil
}

type DestinationPostgresUpdate struct {
	// Name of the database.
	Database string `json:"database"`
	// Hostname of the database.
	Host string `json:"host"`
	// Additional properties to pass to the JDBC URL string when connecting to the database formatted as 'key=value' pairs separated by the symbol '&'. (example: key1=value1&key2=value2&key3=value3).
	JdbcURLParams *string `json:"jdbc_url_params,omitempty"`
	// Password associated with the username.
	Password *string `json:"password,omitempty"`
	// Port of the database.
	Port int64 `json:"port"`
	// The default schema tables are written to if the source does not specify a namespace. The usual value for this field is "public".
	Schema string `json:"schema"`
	// SSL connection modes.
	//  <b>disable</b> - Chose this mode to disable encryption of communication between Airbyte and destination database
	//  <b>allow</b> - Chose this mode to enable encryption only when required by the source database
	//  <b>prefer</b> - Chose this mode to allow unencrypted connection only if the source database does not support encryption
	//  <b>require</b> - Chose this mode to always require encryption. If the source database server does not support encryption, connection will fail
	//   <b>verify-ca</b> - Chose this mode to always require encryption and to verify that the source database server has a valid SSL certificate
	//   <b>verify-full</b> - This is the most secure mode. Chose this mode to always require encryption and to verify the identity of the source database server
	//  See more information - <a href="https://jdbc.postgresql.org/documentation/head/ssl-client.html"> in the docs</a>.
	SslMode *DestinationPostgresUpdateSSLModes `json:"ssl_mode,omitempty"`
	// Whether to initiate an SSH tunnel before connecting to the database, and if so, which kind of authentication to use.
	TunnelMethod *DestinationPostgresUpdateSSHTunnelMethod `json:"tunnel_method,omitempty"`
	// Username to use to access the database.
	Username string `json:"username"`
}