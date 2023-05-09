// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
)

type DestinationClickhouseClickhouseEnum string

const (
	DestinationClickhouseClickhouseEnumClickhouse DestinationClickhouseClickhouseEnum = "clickhouse"
)

func (e DestinationClickhouseClickhouseEnum) ToPointer() *DestinationClickhouseClickhouseEnum {
	return &e
}

func (e *DestinationClickhouseClickhouseEnum) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "clickhouse":
		*e = DestinationClickhouseClickhouseEnum(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationClickhouseClickhouseEnum: %v", v)
	}
}

// DestinationClickhouseSSHTunnelMethodPasswordAuthenticationTunnelMethodEnum - Connect through a jump server tunnel host using username and password authentication
type DestinationClickhouseSSHTunnelMethodPasswordAuthenticationTunnelMethodEnum string

const (
	DestinationClickhouseSSHTunnelMethodPasswordAuthenticationTunnelMethodEnumSSHPasswordAuth DestinationClickhouseSSHTunnelMethodPasswordAuthenticationTunnelMethodEnum = "SSH_PASSWORD_AUTH"
)

func (e DestinationClickhouseSSHTunnelMethodPasswordAuthenticationTunnelMethodEnum) ToPointer() *DestinationClickhouseSSHTunnelMethodPasswordAuthenticationTunnelMethodEnum {
	return &e
}

func (e *DestinationClickhouseSSHTunnelMethodPasswordAuthenticationTunnelMethodEnum) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "SSH_PASSWORD_AUTH":
		*e = DestinationClickhouseSSHTunnelMethodPasswordAuthenticationTunnelMethodEnum(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationClickhouseSSHTunnelMethodPasswordAuthenticationTunnelMethodEnum: %v", v)
	}
}

// DestinationClickhouseSSHTunnelMethodPasswordAuthentication - Whether to initiate an SSH tunnel before connecting to the database, and if so, which kind of authentication to use.
type DestinationClickhouseSSHTunnelMethodPasswordAuthentication struct {
	// Hostname of the jump server host that allows inbound ssh tunnel.
	TunnelHost string `json:"tunnel_host"`
	// Connect through a jump server tunnel host using username and password authentication
	TunnelMethod DestinationClickhouseSSHTunnelMethodPasswordAuthenticationTunnelMethodEnum `json:"tunnel_method"`
	// Port on the proxy/jump server that accepts inbound ssh connections.
	TunnelPort int64 `json:"tunnel_port"`
	// OS-level username for logging into the jump server host
	TunnelUser string `json:"tunnel_user"`
	// OS-level password for logging into the jump server host
	TunnelUserPassword string `json:"tunnel_user_password"`
}

// DestinationClickhouseSSHTunnelMethodSSHKeyAuthenticationTunnelMethodEnum - Connect through a jump server tunnel host using username and ssh key
type DestinationClickhouseSSHTunnelMethodSSHKeyAuthenticationTunnelMethodEnum string

const (
	DestinationClickhouseSSHTunnelMethodSSHKeyAuthenticationTunnelMethodEnumSSHKeyAuth DestinationClickhouseSSHTunnelMethodSSHKeyAuthenticationTunnelMethodEnum = "SSH_KEY_AUTH"
)

func (e DestinationClickhouseSSHTunnelMethodSSHKeyAuthenticationTunnelMethodEnum) ToPointer() *DestinationClickhouseSSHTunnelMethodSSHKeyAuthenticationTunnelMethodEnum {
	return &e
}

func (e *DestinationClickhouseSSHTunnelMethodSSHKeyAuthenticationTunnelMethodEnum) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "SSH_KEY_AUTH":
		*e = DestinationClickhouseSSHTunnelMethodSSHKeyAuthenticationTunnelMethodEnum(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationClickhouseSSHTunnelMethodSSHKeyAuthenticationTunnelMethodEnum: %v", v)
	}
}

// DestinationClickhouseSSHTunnelMethodSSHKeyAuthentication - Whether to initiate an SSH tunnel before connecting to the database, and if so, which kind of authentication to use.
type DestinationClickhouseSSHTunnelMethodSSHKeyAuthentication struct {
	// OS-level user account ssh key credentials in RSA PEM format ( created with ssh-keygen -t rsa -m PEM -f myuser_rsa )
	SSHKey string `json:"ssh_key"`
	// Hostname of the jump server host that allows inbound ssh tunnel.
	TunnelHost string `json:"tunnel_host"`
	// Connect through a jump server tunnel host using username and ssh key
	TunnelMethod DestinationClickhouseSSHTunnelMethodSSHKeyAuthenticationTunnelMethodEnum `json:"tunnel_method"`
	// Port on the proxy/jump server that accepts inbound ssh connections.
	TunnelPort int64 `json:"tunnel_port"`
	// OS-level username for logging into the jump server host.
	TunnelUser string `json:"tunnel_user"`
}

// DestinationClickhouseSSHTunnelMethodNoTunnelTunnelMethodEnum - No ssh tunnel needed to connect to database
type DestinationClickhouseSSHTunnelMethodNoTunnelTunnelMethodEnum string

const (
	DestinationClickhouseSSHTunnelMethodNoTunnelTunnelMethodEnumNoTunnel DestinationClickhouseSSHTunnelMethodNoTunnelTunnelMethodEnum = "NO_TUNNEL"
)

func (e DestinationClickhouseSSHTunnelMethodNoTunnelTunnelMethodEnum) ToPointer() *DestinationClickhouseSSHTunnelMethodNoTunnelTunnelMethodEnum {
	return &e
}

func (e *DestinationClickhouseSSHTunnelMethodNoTunnelTunnelMethodEnum) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "NO_TUNNEL":
		*e = DestinationClickhouseSSHTunnelMethodNoTunnelTunnelMethodEnum(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationClickhouseSSHTunnelMethodNoTunnelTunnelMethodEnum: %v", v)
	}
}

// DestinationClickhouseSSHTunnelMethodNoTunnel - Whether to initiate an SSH tunnel before connecting to the database, and if so, which kind of authentication to use.
type DestinationClickhouseSSHTunnelMethodNoTunnel struct {
	// No ssh tunnel needed to connect to database
	TunnelMethod DestinationClickhouseSSHTunnelMethodNoTunnelTunnelMethodEnum `json:"tunnel_method"`
}

type DestinationClickhouseSSHTunnelMethodType string

const (
	DestinationClickhouseSSHTunnelMethodTypeDestinationClickhouseSSHTunnelMethodNoTunnel               DestinationClickhouseSSHTunnelMethodType = "destination-clickhouse_SSH Tunnel Method_No Tunnel"
	DestinationClickhouseSSHTunnelMethodTypeDestinationClickhouseSSHTunnelMethodSSHKeyAuthentication   DestinationClickhouseSSHTunnelMethodType = "destination-clickhouse_SSH Tunnel Method_SSH Key Authentication"
	DestinationClickhouseSSHTunnelMethodTypeDestinationClickhouseSSHTunnelMethodPasswordAuthentication DestinationClickhouseSSHTunnelMethodType = "destination-clickhouse_SSH Tunnel Method_Password Authentication"
)

type DestinationClickhouseSSHTunnelMethod struct {
	DestinationClickhouseSSHTunnelMethodNoTunnel               *DestinationClickhouseSSHTunnelMethodNoTunnel
	DestinationClickhouseSSHTunnelMethodSSHKeyAuthentication   *DestinationClickhouseSSHTunnelMethodSSHKeyAuthentication
	DestinationClickhouseSSHTunnelMethodPasswordAuthentication *DestinationClickhouseSSHTunnelMethodPasswordAuthentication

	Type DestinationClickhouseSSHTunnelMethodType
}

func CreateDestinationClickhouseSSHTunnelMethodDestinationClickhouseSSHTunnelMethodNoTunnel(destinationClickhouseSSHTunnelMethodNoTunnel DestinationClickhouseSSHTunnelMethodNoTunnel) DestinationClickhouseSSHTunnelMethod {
	typ := DestinationClickhouseSSHTunnelMethodTypeDestinationClickhouseSSHTunnelMethodNoTunnel

	return DestinationClickhouseSSHTunnelMethod{
		DestinationClickhouseSSHTunnelMethodNoTunnel: &destinationClickhouseSSHTunnelMethodNoTunnel,
		Type: typ,
	}
}

func CreateDestinationClickhouseSSHTunnelMethodDestinationClickhouseSSHTunnelMethodSSHKeyAuthentication(destinationClickhouseSSHTunnelMethodSSHKeyAuthentication DestinationClickhouseSSHTunnelMethodSSHKeyAuthentication) DestinationClickhouseSSHTunnelMethod {
	typ := DestinationClickhouseSSHTunnelMethodTypeDestinationClickhouseSSHTunnelMethodSSHKeyAuthentication

	return DestinationClickhouseSSHTunnelMethod{
		DestinationClickhouseSSHTunnelMethodSSHKeyAuthentication: &destinationClickhouseSSHTunnelMethodSSHKeyAuthentication,
		Type: typ,
	}
}

func CreateDestinationClickhouseSSHTunnelMethodDestinationClickhouseSSHTunnelMethodPasswordAuthentication(destinationClickhouseSSHTunnelMethodPasswordAuthentication DestinationClickhouseSSHTunnelMethodPasswordAuthentication) DestinationClickhouseSSHTunnelMethod {
	typ := DestinationClickhouseSSHTunnelMethodTypeDestinationClickhouseSSHTunnelMethodPasswordAuthentication

	return DestinationClickhouseSSHTunnelMethod{
		DestinationClickhouseSSHTunnelMethodPasswordAuthentication: &destinationClickhouseSSHTunnelMethodPasswordAuthentication,
		Type: typ,
	}
}

func (u *DestinationClickhouseSSHTunnelMethod) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	destinationClickhouseSSHTunnelMethodNoTunnel := new(DestinationClickhouseSSHTunnelMethodNoTunnel)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&destinationClickhouseSSHTunnelMethodNoTunnel); err == nil {
		u.DestinationClickhouseSSHTunnelMethodNoTunnel = destinationClickhouseSSHTunnelMethodNoTunnel
		u.Type = DestinationClickhouseSSHTunnelMethodTypeDestinationClickhouseSSHTunnelMethodNoTunnel
		return nil
	}

	destinationClickhouseSSHTunnelMethodSSHKeyAuthentication := new(DestinationClickhouseSSHTunnelMethodSSHKeyAuthentication)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&destinationClickhouseSSHTunnelMethodSSHKeyAuthentication); err == nil {
		u.DestinationClickhouseSSHTunnelMethodSSHKeyAuthentication = destinationClickhouseSSHTunnelMethodSSHKeyAuthentication
		u.Type = DestinationClickhouseSSHTunnelMethodTypeDestinationClickhouseSSHTunnelMethodSSHKeyAuthentication
		return nil
	}

	destinationClickhouseSSHTunnelMethodPasswordAuthentication := new(DestinationClickhouseSSHTunnelMethodPasswordAuthentication)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&destinationClickhouseSSHTunnelMethodPasswordAuthentication); err == nil {
		u.DestinationClickhouseSSHTunnelMethodPasswordAuthentication = destinationClickhouseSSHTunnelMethodPasswordAuthentication
		u.Type = DestinationClickhouseSSHTunnelMethodTypeDestinationClickhouseSSHTunnelMethodPasswordAuthentication
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u DestinationClickhouseSSHTunnelMethod) MarshalJSON() ([]byte, error) {
	if u.DestinationClickhouseSSHTunnelMethodNoTunnel != nil {
		return json.Marshal(u.DestinationClickhouseSSHTunnelMethodNoTunnel)
	}

	if u.DestinationClickhouseSSHTunnelMethodSSHKeyAuthentication != nil {
		return json.Marshal(u.DestinationClickhouseSSHTunnelMethodSSHKeyAuthentication)
	}

	if u.DestinationClickhouseSSHTunnelMethodPasswordAuthentication != nil {
		return json.Marshal(u.DestinationClickhouseSSHTunnelMethodPasswordAuthentication)
	}

	return nil, nil
}

// DestinationClickhouse - The values required to configure the destination.
type DestinationClickhouse struct {
	// Name of the database.
	Database        string                              `json:"database"`
	DestinationType DestinationClickhouseClickhouseEnum `json:"destinationType"`
	// Hostname of the database.
	Host string `json:"host"`
	// Additional properties to pass to the JDBC URL string when connecting to the database formatted as 'key=value' pairs separated by the symbol '&'. (example: key1=value1&key2=value2&key3=value3).
	JdbcURLParams *string `json:"jdbc_url_params,omitempty"`
	// Password associated with the username.
	Password *string `json:"password,omitempty"`
	// HTTP port of the database.
	Port int64 `json:"port"`
	// Whether to initiate an SSH tunnel before connecting to the database, and if so, which kind of authentication to use.
	TunnelMethod *DestinationClickhouseSSHTunnelMethod `json:"tunnel_method,omitempty"`
	// Username to use to access the database.
	Username string `json:"username"`
}
