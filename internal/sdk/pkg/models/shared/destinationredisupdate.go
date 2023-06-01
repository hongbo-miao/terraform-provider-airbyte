// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
)

// DestinationRedisUpdateCacheType - Redis cache type to store data in.
type DestinationRedisUpdateCacheType string

const (
	DestinationRedisUpdateCacheTypeHash DestinationRedisUpdateCacheType = "hash"
)

func (e DestinationRedisUpdateCacheType) ToPointer() *DestinationRedisUpdateCacheType {
	return &e
}

func (e *DestinationRedisUpdateCacheType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "hash":
		*e = DestinationRedisUpdateCacheType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationRedisUpdateCacheType: %v", v)
	}
}

type DestinationRedisUpdateSSLModesVerifyFullMode string

const (
	DestinationRedisUpdateSSLModesVerifyFullModeVerifyFull DestinationRedisUpdateSSLModesVerifyFullMode = "verify-full"
)

func (e DestinationRedisUpdateSSLModesVerifyFullMode) ToPointer() *DestinationRedisUpdateSSLModesVerifyFullMode {
	return &e
}

func (e *DestinationRedisUpdateSSLModesVerifyFullMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "verify-full":
		*e = DestinationRedisUpdateSSLModesVerifyFullMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationRedisUpdateSSLModesVerifyFullMode: %v", v)
	}
}

// DestinationRedisUpdateSSLModesVerifyFull - Verify-full SSL mode.
type DestinationRedisUpdateSSLModesVerifyFull struct {
	// CA certificate
	CaCertificate string `json:"ca_certificate"`
	// Client certificate
	ClientCertificate string `json:"client_certificate"`
	// Client key
	ClientKey string `json:"client_key"`
	// Password for keystorage. If you do not add it - the password will be generated automatically.
	ClientKeyPassword *string                                      `json:"client_key_password,omitempty"`
	Mode              DestinationRedisUpdateSSLModesVerifyFullMode `json:"mode"`
}

type DestinationRedisUpdateSSLModesDisableMode string

const (
	DestinationRedisUpdateSSLModesDisableModeDisable DestinationRedisUpdateSSLModesDisableMode = "disable"
)

func (e DestinationRedisUpdateSSLModesDisableMode) ToPointer() *DestinationRedisUpdateSSLModesDisableMode {
	return &e
}

func (e *DestinationRedisUpdateSSLModesDisableMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "disable":
		*e = DestinationRedisUpdateSSLModesDisableMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationRedisUpdateSSLModesDisableMode: %v", v)
	}
}

// DestinationRedisUpdateSSLModesDisable - Disable SSL.
type DestinationRedisUpdateSSLModesDisable struct {
	Mode DestinationRedisUpdateSSLModesDisableMode `json:"mode"`
}

type DestinationRedisUpdateSSLModesType string

const (
	DestinationRedisUpdateSSLModesTypeDestinationRedisUpdateSSLModesDisable    DestinationRedisUpdateSSLModesType = "destination-redis-update_SSL Modes_disable"
	DestinationRedisUpdateSSLModesTypeDestinationRedisUpdateSSLModesVerifyFull DestinationRedisUpdateSSLModesType = "destination-redis-update_SSL Modes_verify-full"
)

type DestinationRedisUpdateSSLModes struct {
	DestinationRedisUpdateSSLModesDisable    *DestinationRedisUpdateSSLModesDisable
	DestinationRedisUpdateSSLModesVerifyFull *DestinationRedisUpdateSSLModesVerifyFull

	Type DestinationRedisUpdateSSLModesType
}

func CreateDestinationRedisUpdateSSLModesDestinationRedisUpdateSSLModesDisable(destinationRedisUpdateSSLModesDisable DestinationRedisUpdateSSLModesDisable) DestinationRedisUpdateSSLModes {
	typ := DestinationRedisUpdateSSLModesTypeDestinationRedisUpdateSSLModesDisable

	return DestinationRedisUpdateSSLModes{
		DestinationRedisUpdateSSLModesDisable: &destinationRedisUpdateSSLModesDisable,
		Type:                                  typ,
	}
}

func CreateDestinationRedisUpdateSSLModesDestinationRedisUpdateSSLModesVerifyFull(destinationRedisUpdateSSLModesVerifyFull DestinationRedisUpdateSSLModesVerifyFull) DestinationRedisUpdateSSLModes {
	typ := DestinationRedisUpdateSSLModesTypeDestinationRedisUpdateSSLModesVerifyFull

	return DestinationRedisUpdateSSLModes{
		DestinationRedisUpdateSSLModesVerifyFull: &destinationRedisUpdateSSLModesVerifyFull,
		Type:                                     typ,
	}
}

func (u *DestinationRedisUpdateSSLModes) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	destinationRedisUpdateSSLModesDisable := new(DestinationRedisUpdateSSLModesDisable)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&destinationRedisUpdateSSLModesDisable); err == nil {
		u.DestinationRedisUpdateSSLModesDisable = destinationRedisUpdateSSLModesDisable
		u.Type = DestinationRedisUpdateSSLModesTypeDestinationRedisUpdateSSLModesDisable
		return nil
	}

	destinationRedisUpdateSSLModesVerifyFull := new(DestinationRedisUpdateSSLModesVerifyFull)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&destinationRedisUpdateSSLModesVerifyFull); err == nil {
		u.DestinationRedisUpdateSSLModesVerifyFull = destinationRedisUpdateSSLModesVerifyFull
		u.Type = DestinationRedisUpdateSSLModesTypeDestinationRedisUpdateSSLModesVerifyFull
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u DestinationRedisUpdateSSLModes) MarshalJSON() ([]byte, error) {
	if u.DestinationRedisUpdateSSLModesDisable != nil {
		return json.Marshal(u.DestinationRedisUpdateSSLModesDisable)
	}

	if u.DestinationRedisUpdateSSLModesVerifyFull != nil {
		return json.Marshal(u.DestinationRedisUpdateSSLModesVerifyFull)
	}

	return nil, nil
}

// DestinationRedisUpdateSSHTunnelMethodPasswordAuthenticationTunnelMethod - Connect through a jump server tunnel host using username and password authentication
type DestinationRedisUpdateSSHTunnelMethodPasswordAuthenticationTunnelMethod string

const (
	DestinationRedisUpdateSSHTunnelMethodPasswordAuthenticationTunnelMethodSSHPasswordAuth DestinationRedisUpdateSSHTunnelMethodPasswordAuthenticationTunnelMethod = "SSH_PASSWORD_AUTH"
)

func (e DestinationRedisUpdateSSHTunnelMethodPasswordAuthenticationTunnelMethod) ToPointer() *DestinationRedisUpdateSSHTunnelMethodPasswordAuthenticationTunnelMethod {
	return &e
}

func (e *DestinationRedisUpdateSSHTunnelMethodPasswordAuthenticationTunnelMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "SSH_PASSWORD_AUTH":
		*e = DestinationRedisUpdateSSHTunnelMethodPasswordAuthenticationTunnelMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationRedisUpdateSSHTunnelMethodPasswordAuthenticationTunnelMethod: %v", v)
	}
}

// DestinationRedisUpdateSSHTunnelMethodPasswordAuthentication - Whether to initiate an SSH tunnel before connecting to the database, and if so, which kind of authentication to use.
type DestinationRedisUpdateSSHTunnelMethodPasswordAuthentication struct {
	// Hostname of the jump server host that allows inbound ssh tunnel.
	TunnelHost string `json:"tunnel_host"`
	// Connect through a jump server tunnel host using username and password authentication
	TunnelMethod DestinationRedisUpdateSSHTunnelMethodPasswordAuthenticationTunnelMethod `json:"tunnel_method"`
	// Port on the proxy/jump server that accepts inbound ssh connections.
	TunnelPort int64 `json:"tunnel_port"`
	// OS-level username for logging into the jump server host
	TunnelUser string `json:"tunnel_user"`
	// OS-level password for logging into the jump server host
	TunnelUserPassword string `json:"tunnel_user_password"`
}

// DestinationRedisUpdateSSHTunnelMethodSSHKeyAuthenticationTunnelMethod - Connect through a jump server tunnel host using username and ssh key
type DestinationRedisUpdateSSHTunnelMethodSSHKeyAuthenticationTunnelMethod string

const (
	DestinationRedisUpdateSSHTunnelMethodSSHKeyAuthenticationTunnelMethodSSHKeyAuth DestinationRedisUpdateSSHTunnelMethodSSHKeyAuthenticationTunnelMethod = "SSH_KEY_AUTH"
)

func (e DestinationRedisUpdateSSHTunnelMethodSSHKeyAuthenticationTunnelMethod) ToPointer() *DestinationRedisUpdateSSHTunnelMethodSSHKeyAuthenticationTunnelMethod {
	return &e
}

func (e *DestinationRedisUpdateSSHTunnelMethodSSHKeyAuthenticationTunnelMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "SSH_KEY_AUTH":
		*e = DestinationRedisUpdateSSHTunnelMethodSSHKeyAuthenticationTunnelMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationRedisUpdateSSHTunnelMethodSSHKeyAuthenticationTunnelMethod: %v", v)
	}
}

// DestinationRedisUpdateSSHTunnelMethodSSHKeyAuthentication - Whether to initiate an SSH tunnel before connecting to the database, and if so, which kind of authentication to use.
type DestinationRedisUpdateSSHTunnelMethodSSHKeyAuthentication struct {
	// OS-level user account ssh key credentials in RSA PEM format ( created with ssh-keygen -t rsa -m PEM -f myuser_rsa )
	SSHKey string `json:"ssh_key"`
	// Hostname of the jump server host that allows inbound ssh tunnel.
	TunnelHost string `json:"tunnel_host"`
	// Connect through a jump server tunnel host using username and ssh key
	TunnelMethod DestinationRedisUpdateSSHTunnelMethodSSHKeyAuthenticationTunnelMethod `json:"tunnel_method"`
	// Port on the proxy/jump server that accepts inbound ssh connections.
	TunnelPort int64 `json:"tunnel_port"`
	// OS-level username for logging into the jump server host.
	TunnelUser string `json:"tunnel_user"`
}

// DestinationRedisUpdateSSHTunnelMethodNoTunnelTunnelMethod - No ssh tunnel needed to connect to database
type DestinationRedisUpdateSSHTunnelMethodNoTunnelTunnelMethod string

const (
	DestinationRedisUpdateSSHTunnelMethodNoTunnelTunnelMethodNoTunnel DestinationRedisUpdateSSHTunnelMethodNoTunnelTunnelMethod = "NO_TUNNEL"
)

func (e DestinationRedisUpdateSSHTunnelMethodNoTunnelTunnelMethod) ToPointer() *DestinationRedisUpdateSSHTunnelMethodNoTunnelTunnelMethod {
	return &e
}

func (e *DestinationRedisUpdateSSHTunnelMethodNoTunnelTunnelMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "NO_TUNNEL":
		*e = DestinationRedisUpdateSSHTunnelMethodNoTunnelTunnelMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationRedisUpdateSSHTunnelMethodNoTunnelTunnelMethod: %v", v)
	}
}

// DestinationRedisUpdateSSHTunnelMethodNoTunnel - Whether to initiate an SSH tunnel before connecting to the database, and if so, which kind of authentication to use.
type DestinationRedisUpdateSSHTunnelMethodNoTunnel struct {
	// No ssh tunnel needed to connect to database
	TunnelMethod DestinationRedisUpdateSSHTunnelMethodNoTunnelTunnelMethod `json:"tunnel_method"`
}

type DestinationRedisUpdateSSHTunnelMethodType string

const (
	DestinationRedisUpdateSSHTunnelMethodTypeDestinationRedisUpdateSSHTunnelMethodNoTunnel               DestinationRedisUpdateSSHTunnelMethodType = "destination-redis-update_SSH Tunnel Method_No Tunnel"
	DestinationRedisUpdateSSHTunnelMethodTypeDestinationRedisUpdateSSHTunnelMethodSSHKeyAuthentication   DestinationRedisUpdateSSHTunnelMethodType = "destination-redis-update_SSH Tunnel Method_SSH Key Authentication"
	DestinationRedisUpdateSSHTunnelMethodTypeDestinationRedisUpdateSSHTunnelMethodPasswordAuthentication DestinationRedisUpdateSSHTunnelMethodType = "destination-redis-update_SSH Tunnel Method_Password Authentication"
)

type DestinationRedisUpdateSSHTunnelMethod struct {
	DestinationRedisUpdateSSHTunnelMethodNoTunnel               *DestinationRedisUpdateSSHTunnelMethodNoTunnel
	DestinationRedisUpdateSSHTunnelMethodSSHKeyAuthentication   *DestinationRedisUpdateSSHTunnelMethodSSHKeyAuthentication
	DestinationRedisUpdateSSHTunnelMethodPasswordAuthentication *DestinationRedisUpdateSSHTunnelMethodPasswordAuthentication

	Type DestinationRedisUpdateSSHTunnelMethodType
}

func CreateDestinationRedisUpdateSSHTunnelMethodDestinationRedisUpdateSSHTunnelMethodNoTunnel(destinationRedisUpdateSSHTunnelMethodNoTunnel DestinationRedisUpdateSSHTunnelMethodNoTunnel) DestinationRedisUpdateSSHTunnelMethod {
	typ := DestinationRedisUpdateSSHTunnelMethodTypeDestinationRedisUpdateSSHTunnelMethodNoTunnel

	return DestinationRedisUpdateSSHTunnelMethod{
		DestinationRedisUpdateSSHTunnelMethodNoTunnel: &destinationRedisUpdateSSHTunnelMethodNoTunnel,
		Type: typ,
	}
}

func CreateDestinationRedisUpdateSSHTunnelMethodDestinationRedisUpdateSSHTunnelMethodSSHKeyAuthentication(destinationRedisUpdateSSHTunnelMethodSSHKeyAuthentication DestinationRedisUpdateSSHTunnelMethodSSHKeyAuthentication) DestinationRedisUpdateSSHTunnelMethod {
	typ := DestinationRedisUpdateSSHTunnelMethodTypeDestinationRedisUpdateSSHTunnelMethodSSHKeyAuthentication

	return DestinationRedisUpdateSSHTunnelMethod{
		DestinationRedisUpdateSSHTunnelMethodSSHKeyAuthentication: &destinationRedisUpdateSSHTunnelMethodSSHKeyAuthentication,
		Type: typ,
	}
}

func CreateDestinationRedisUpdateSSHTunnelMethodDestinationRedisUpdateSSHTunnelMethodPasswordAuthentication(destinationRedisUpdateSSHTunnelMethodPasswordAuthentication DestinationRedisUpdateSSHTunnelMethodPasswordAuthentication) DestinationRedisUpdateSSHTunnelMethod {
	typ := DestinationRedisUpdateSSHTunnelMethodTypeDestinationRedisUpdateSSHTunnelMethodPasswordAuthentication

	return DestinationRedisUpdateSSHTunnelMethod{
		DestinationRedisUpdateSSHTunnelMethodPasswordAuthentication: &destinationRedisUpdateSSHTunnelMethodPasswordAuthentication,
		Type: typ,
	}
}

func (u *DestinationRedisUpdateSSHTunnelMethod) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	destinationRedisUpdateSSHTunnelMethodNoTunnel := new(DestinationRedisUpdateSSHTunnelMethodNoTunnel)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&destinationRedisUpdateSSHTunnelMethodNoTunnel); err == nil {
		u.DestinationRedisUpdateSSHTunnelMethodNoTunnel = destinationRedisUpdateSSHTunnelMethodNoTunnel
		u.Type = DestinationRedisUpdateSSHTunnelMethodTypeDestinationRedisUpdateSSHTunnelMethodNoTunnel
		return nil
	}

	destinationRedisUpdateSSHTunnelMethodSSHKeyAuthentication := new(DestinationRedisUpdateSSHTunnelMethodSSHKeyAuthentication)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&destinationRedisUpdateSSHTunnelMethodSSHKeyAuthentication); err == nil {
		u.DestinationRedisUpdateSSHTunnelMethodSSHKeyAuthentication = destinationRedisUpdateSSHTunnelMethodSSHKeyAuthentication
		u.Type = DestinationRedisUpdateSSHTunnelMethodTypeDestinationRedisUpdateSSHTunnelMethodSSHKeyAuthentication
		return nil
	}

	destinationRedisUpdateSSHTunnelMethodPasswordAuthentication := new(DestinationRedisUpdateSSHTunnelMethodPasswordAuthentication)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&destinationRedisUpdateSSHTunnelMethodPasswordAuthentication); err == nil {
		u.DestinationRedisUpdateSSHTunnelMethodPasswordAuthentication = destinationRedisUpdateSSHTunnelMethodPasswordAuthentication
		u.Type = DestinationRedisUpdateSSHTunnelMethodTypeDestinationRedisUpdateSSHTunnelMethodPasswordAuthentication
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u DestinationRedisUpdateSSHTunnelMethod) MarshalJSON() ([]byte, error) {
	if u.DestinationRedisUpdateSSHTunnelMethodNoTunnel != nil {
		return json.Marshal(u.DestinationRedisUpdateSSHTunnelMethodNoTunnel)
	}

	if u.DestinationRedisUpdateSSHTunnelMethodSSHKeyAuthentication != nil {
		return json.Marshal(u.DestinationRedisUpdateSSHTunnelMethodSSHKeyAuthentication)
	}

	if u.DestinationRedisUpdateSSHTunnelMethodPasswordAuthentication != nil {
		return json.Marshal(u.DestinationRedisUpdateSSHTunnelMethodPasswordAuthentication)
	}

	return nil, nil
}

type DestinationRedisUpdate struct {
	// Redis cache type to store data in.
	CacheType DestinationRedisUpdateCacheType `json:"cache_type"`
	// Redis host to connect to.
	Host string `json:"host"`
	// Password associated with Redis.
	Password *string `json:"password,omitempty"`
	// Port of Redis.
	Port int64 `json:"port"`
	// Indicates whether SSL encryption protocol will be used to connect to Redis. It is recommended to use SSL connection if possible.
	Ssl *bool `json:"ssl,omitempty"`
	// SSL connection modes.
	//   <li><b>verify-full</b> - This is the most secure mode. Always require encryption and verifies the identity of the source database server
	SslMode *DestinationRedisUpdateSSLModes `json:"ssl_mode,omitempty"`
	// Whether to initiate an SSH tunnel before connecting to the database, and if so, which kind of authentication to use.
	TunnelMethod *DestinationRedisUpdateSSHTunnelMethod `json:"tunnel_method,omitempty"`
	// Username associated with Redis.
	Username string `json:"username"`
}