// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
)

type DestinationMssqlMssqlEnum string

const (
	DestinationMssqlMssqlEnumMssql DestinationMssqlMssqlEnum = "mssql"
)

func (e DestinationMssqlMssqlEnum) ToPointer() *DestinationMssqlMssqlEnum {
	return &e
}

func (e *DestinationMssqlMssqlEnum) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "mssql":
		*e = DestinationMssqlMssqlEnum(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationMssqlMssqlEnum: %v", v)
	}
}

type DestinationMssqlSSLMethodEncryptedVerifyCertificateSSLMethodEnum string

const (
	DestinationMssqlSSLMethodEncryptedVerifyCertificateSSLMethodEnumEncryptedVerifyCertificate DestinationMssqlSSLMethodEncryptedVerifyCertificateSSLMethodEnum = "encrypted_verify_certificate"
)

func (e DestinationMssqlSSLMethodEncryptedVerifyCertificateSSLMethodEnum) ToPointer() *DestinationMssqlSSLMethodEncryptedVerifyCertificateSSLMethodEnum {
	return &e
}

func (e *DestinationMssqlSSLMethodEncryptedVerifyCertificateSSLMethodEnum) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "encrypted_verify_certificate":
		*e = DestinationMssqlSSLMethodEncryptedVerifyCertificateSSLMethodEnum(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationMssqlSSLMethodEncryptedVerifyCertificateSSLMethodEnum: %v", v)
	}
}

// DestinationMssqlSSLMethodEncryptedVerifyCertificate - Verify and use the certificate provided by the server.
type DestinationMssqlSSLMethodEncryptedVerifyCertificate struct {
	// Specifies the host name of the server. The value of this property must match the subject property of the certificate.
	HostNameInCertificate *string                                                          `json:"hostNameInCertificate,omitempty"`
	SslMethod             DestinationMssqlSSLMethodEncryptedVerifyCertificateSSLMethodEnum `json:"ssl_method"`
}

type DestinationMssqlSSLMethodEncryptedTrustServerCertificateSSLMethodEnum string

const (
	DestinationMssqlSSLMethodEncryptedTrustServerCertificateSSLMethodEnumEncryptedTrustServerCertificate DestinationMssqlSSLMethodEncryptedTrustServerCertificateSSLMethodEnum = "encrypted_trust_server_certificate"
)

func (e DestinationMssqlSSLMethodEncryptedTrustServerCertificateSSLMethodEnum) ToPointer() *DestinationMssqlSSLMethodEncryptedTrustServerCertificateSSLMethodEnum {
	return &e
}

func (e *DestinationMssqlSSLMethodEncryptedTrustServerCertificateSSLMethodEnum) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "encrypted_trust_server_certificate":
		*e = DestinationMssqlSSLMethodEncryptedTrustServerCertificateSSLMethodEnum(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationMssqlSSLMethodEncryptedTrustServerCertificateSSLMethodEnum: %v", v)
	}
}

// DestinationMssqlSSLMethodEncryptedTrustServerCertificate - Use the certificate provided by the server without verification. (For testing purposes only!)
type DestinationMssqlSSLMethodEncryptedTrustServerCertificate struct {
	SslMethod DestinationMssqlSSLMethodEncryptedTrustServerCertificateSSLMethodEnum `json:"ssl_method"`
}

type DestinationMssqlSSLMethodType string

const (
	DestinationMssqlSSLMethodTypeDestinationMssqlSSLMethodEncryptedTrustServerCertificate DestinationMssqlSSLMethodType = "destination-mssql_SSL Method_Encrypted (trust server certificate)"
	DestinationMssqlSSLMethodTypeDestinationMssqlSSLMethodEncryptedVerifyCertificate      DestinationMssqlSSLMethodType = "destination-mssql_SSL Method_Encrypted (verify certificate)"
)

type DestinationMssqlSSLMethod struct {
	DestinationMssqlSSLMethodEncryptedTrustServerCertificate *DestinationMssqlSSLMethodEncryptedTrustServerCertificate
	DestinationMssqlSSLMethodEncryptedVerifyCertificate      *DestinationMssqlSSLMethodEncryptedVerifyCertificate

	Type DestinationMssqlSSLMethodType
}

func CreateDestinationMssqlSSLMethodDestinationMssqlSSLMethodEncryptedTrustServerCertificate(destinationMssqlSSLMethodEncryptedTrustServerCertificate DestinationMssqlSSLMethodEncryptedTrustServerCertificate) DestinationMssqlSSLMethod {
	typ := DestinationMssqlSSLMethodTypeDestinationMssqlSSLMethodEncryptedTrustServerCertificate

	return DestinationMssqlSSLMethod{
		DestinationMssqlSSLMethodEncryptedTrustServerCertificate: &destinationMssqlSSLMethodEncryptedTrustServerCertificate,
		Type: typ,
	}
}

func CreateDestinationMssqlSSLMethodDestinationMssqlSSLMethodEncryptedVerifyCertificate(destinationMssqlSSLMethodEncryptedVerifyCertificate DestinationMssqlSSLMethodEncryptedVerifyCertificate) DestinationMssqlSSLMethod {
	typ := DestinationMssqlSSLMethodTypeDestinationMssqlSSLMethodEncryptedVerifyCertificate

	return DestinationMssqlSSLMethod{
		DestinationMssqlSSLMethodEncryptedVerifyCertificate: &destinationMssqlSSLMethodEncryptedVerifyCertificate,
		Type: typ,
	}
}

func (u *DestinationMssqlSSLMethod) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	destinationMssqlSSLMethodEncryptedTrustServerCertificate := new(DestinationMssqlSSLMethodEncryptedTrustServerCertificate)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&destinationMssqlSSLMethodEncryptedTrustServerCertificate); err == nil {
		u.DestinationMssqlSSLMethodEncryptedTrustServerCertificate = destinationMssqlSSLMethodEncryptedTrustServerCertificate
		u.Type = DestinationMssqlSSLMethodTypeDestinationMssqlSSLMethodEncryptedTrustServerCertificate
		return nil
	}

	destinationMssqlSSLMethodEncryptedVerifyCertificate := new(DestinationMssqlSSLMethodEncryptedVerifyCertificate)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&destinationMssqlSSLMethodEncryptedVerifyCertificate); err == nil {
		u.DestinationMssqlSSLMethodEncryptedVerifyCertificate = destinationMssqlSSLMethodEncryptedVerifyCertificate
		u.Type = DestinationMssqlSSLMethodTypeDestinationMssqlSSLMethodEncryptedVerifyCertificate
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u DestinationMssqlSSLMethod) MarshalJSON() ([]byte, error) {
	if u.DestinationMssqlSSLMethodEncryptedTrustServerCertificate != nil {
		return json.Marshal(u.DestinationMssqlSSLMethodEncryptedTrustServerCertificate)
	}

	if u.DestinationMssqlSSLMethodEncryptedVerifyCertificate != nil {
		return json.Marshal(u.DestinationMssqlSSLMethodEncryptedVerifyCertificate)
	}

	return nil, nil
}

// DestinationMssqlSSHTunnelMethodPasswordAuthenticationTunnelMethodEnum - Connect through a jump server tunnel host using username and password authentication
type DestinationMssqlSSHTunnelMethodPasswordAuthenticationTunnelMethodEnum string

const (
	DestinationMssqlSSHTunnelMethodPasswordAuthenticationTunnelMethodEnumSSHPasswordAuth DestinationMssqlSSHTunnelMethodPasswordAuthenticationTunnelMethodEnum = "SSH_PASSWORD_AUTH"
)

func (e DestinationMssqlSSHTunnelMethodPasswordAuthenticationTunnelMethodEnum) ToPointer() *DestinationMssqlSSHTunnelMethodPasswordAuthenticationTunnelMethodEnum {
	return &e
}

func (e *DestinationMssqlSSHTunnelMethodPasswordAuthenticationTunnelMethodEnum) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "SSH_PASSWORD_AUTH":
		*e = DestinationMssqlSSHTunnelMethodPasswordAuthenticationTunnelMethodEnum(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationMssqlSSHTunnelMethodPasswordAuthenticationTunnelMethodEnum: %v", v)
	}
}

// DestinationMssqlSSHTunnelMethodPasswordAuthentication - Whether to initiate an SSH tunnel before connecting to the database, and if so, which kind of authentication to use.
type DestinationMssqlSSHTunnelMethodPasswordAuthentication struct {
	// Hostname of the jump server host that allows inbound ssh tunnel.
	TunnelHost string `json:"tunnel_host"`
	// Connect through a jump server tunnel host using username and password authentication
	TunnelMethod DestinationMssqlSSHTunnelMethodPasswordAuthenticationTunnelMethodEnum `json:"tunnel_method"`
	// Port on the proxy/jump server that accepts inbound ssh connections.
	TunnelPort int64 `json:"tunnel_port"`
	// OS-level username for logging into the jump server host
	TunnelUser string `json:"tunnel_user"`
	// OS-level password for logging into the jump server host
	TunnelUserPassword string `json:"tunnel_user_password"`
}

// DestinationMssqlSSHTunnelMethodSSHKeyAuthenticationTunnelMethodEnum - Connect through a jump server tunnel host using username and ssh key
type DestinationMssqlSSHTunnelMethodSSHKeyAuthenticationTunnelMethodEnum string

const (
	DestinationMssqlSSHTunnelMethodSSHKeyAuthenticationTunnelMethodEnumSSHKeyAuth DestinationMssqlSSHTunnelMethodSSHKeyAuthenticationTunnelMethodEnum = "SSH_KEY_AUTH"
)

func (e DestinationMssqlSSHTunnelMethodSSHKeyAuthenticationTunnelMethodEnum) ToPointer() *DestinationMssqlSSHTunnelMethodSSHKeyAuthenticationTunnelMethodEnum {
	return &e
}

func (e *DestinationMssqlSSHTunnelMethodSSHKeyAuthenticationTunnelMethodEnum) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "SSH_KEY_AUTH":
		*e = DestinationMssqlSSHTunnelMethodSSHKeyAuthenticationTunnelMethodEnum(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationMssqlSSHTunnelMethodSSHKeyAuthenticationTunnelMethodEnum: %v", v)
	}
}

// DestinationMssqlSSHTunnelMethodSSHKeyAuthentication - Whether to initiate an SSH tunnel before connecting to the database, and if so, which kind of authentication to use.
type DestinationMssqlSSHTunnelMethodSSHKeyAuthentication struct {
	// OS-level user account ssh key credentials in RSA PEM format ( created with ssh-keygen -t rsa -m PEM -f myuser_rsa )
	SSHKey string `json:"ssh_key"`
	// Hostname of the jump server host that allows inbound ssh tunnel.
	TunnelHost string `json:"tunnel_host"`
	// Connect through a jump server tunnel host using username and ssh key
	TunnelMethod DestinationMssqlSSHTunnelMethodSSHKeyAuthenticationTunnelMethodEnum `json:"tunnel_method"`
	// Port on the proxy/jump server that accepts inbound ssh connections.
	TunnelPort int64 `json:"tunnel_port"`
	// OS-level username for logging into the jump server host.
	TunnelUser string `json:"tunnel_user"`
}

// DestinationMssqlSSHTunnelMethodNoTunnelTunnelMethodEnum - No ssh tunnel needed to connect to database
type DestinationMssqlSSHTunnelMethodNoTunnelTunnelMethodEnum string

const (
	DestinationMssqlSSHTunnelMethodNoTunnelTunnelMethodEnumNoTunnel DestinationMssqlSSHTunnelMethodNoTunnelTunnelMethodEnum = "NO_TUNNEL"
)

func (e DestinationMssqlSSHTunnelMethodNoTunnelTunnelMethodEnum) ToPointer() *DestinationMssqlSSHTunnelMethodNoTunnelTunnelMethodEnum {
	return &e
}

func (e *DestinationMssqlSSHTunnelMethodNoTunnelTunnelMethodEnum) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "NO_TUNNEL":
		*e = DestinationMssqlSSHTunnelMethodNoTunnelTunnelMethodEnum(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationMssqlSSHTunnelMethodNoTunnelTunnelMethodEnum: %v", v)
	}
}

// DestinationMssqlSSHTunnelMethodNoTunnel - Whether to initiate an SSH tunnel before connecting to the database, and if so, which kind of authentication to use.
type DestinationMssqlSSHTunnelMethodNoTunnel struct {
	// No ssh tunnel needed to connect to database
	TunnelMethod DestinationMssqlSSHTunnelMethodNoTunnelTunnelMethodEnum `json:"tunnel_method"`
}

type DestinationMssqlSSHTunnelMethodType string

const (
	DestinationMssqlSSHTunnelMethodTypeDestinationMssqlSSHTunnelMethodNoTunnel               DestinationMssqlSSHTunnelMethodType = "destination-mssql_SSH Tunnel Method_No Tunnel"
	DestinationMssqlSSHTunnelMethodTypeDestinationMssqlSSHTunnelMethodSSHKeyAuthentication   DestinationMssqlSSHTunnelMethodType = "destination-mssql_SSH Tunnel Method_SSH Key Authentication"
	DestinationMssqlSSHTunnelMethodTypeDestinationMssqlSSHTunnelMethodPasswordAuthentication DestinationMssqlSSHTunnelMethodType = "destination-mssql_SSH Tunnel Method_Password Authentication"
)

type DestinationMssqlSSHTunnelMethod struct {
	DestinationMssqlSSHTunnelMethodNoTunnel               *DestinationMssqlSSHTunnelMethodNoTunnel
	DestinationMssqlSSHTunnelMethodSSHKeyAuthentication   *DestinationMssqlSSHTunnelMethodSSHKeyAuthentication
	DestinationMssqlSSHTunnelMethodPasswordAuthentication *DestinationMssqlSSHTunnelMethodPasswordAuthentication

	Type DestinationMssqlSSHTunnelMethodType
}

func CreateDestinationMssqlSSHTunnelMethodDestinationMssqlSSHTunnelMethodNoTunnel(destinationMssqlSSHTunnelMethodNoTunnel DestinationMssqlSSHTunnelMethodNoTunnel) DestinationMssqlSSHTunnelMethod {
	typ := DestinationMssqlSSHTunnelMethodTypeDestinationMssqlSSHTunnelMethodNoTunnel

	return DestinationMssqlSSHTunnelMethod{
		DestinationMssqlSSHTunnelMethodNoTunnel: &destinationMssqlSSHTunnelMethodNoTunnel,
		Type:                                    typ,
	}
}

func CreateDestinationMssqlSSHTunnelMethodDestinationMssqlSSHTunnelMethodSSHKeyAuthentication(destinationMssqlSSHTunnelMethodSSHKeyAuthentication DestinationMssqlSSHTunnelMethodSSHKeyAuthentication) DestinationMssqlSSHTunnelMethod {
	typ := DestinationMssqlSSHTunnelMethodTypeDestinationMssqlSSHTunnelMethodSSHKeyAuthentication

	return DestinationMssqlSSHTunnelMethod{
		DestinationMssqlSSHTunnelMethodSSHKeyAuthentication: &destinationMssqlSSHTunnelMethodSSHKeyAuthentication,
		Type: typ,
	}
}

func CreateDestinationMssqlSSHTunnelMethodDestinationMssqlSSHTunnelMethodPasswordAuthentication(destinationMssqlSSHTunnelMethodPasswordAuthentication DestinationMssqlSSHTunnelMethodPasswordAuthentication) DestinationMssqlSSHTunnelMethod {
	typ := DestinationMssqlSSHTunnelMethodTypeDestinationMssqlSSHTunnelMethodPasswordAuthentication

	return DestinationMssqlSSHTunnelMethod{
		DestinationMssqlSSHTunnelMethodPasswordAuthentication: &destinationMssqlSSHTunnelMethodPasswordAuthentication,
		Type: typ,
	}
}

func (u *DestinationMssqlSSHTunnelMethod) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	destinationMssqlSSHTunnelMethodNoTunnel := new(DestinationMssqlSSHTunnelMethodNoTunnel)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&destinationMssqlSSHTunnelMethodNoTunnel); err == nil {
		u.DestinationMssqlSSHTunnelMethodNoTunnel = destinationMssqlSSHTunnelMethodNoTunnel
		u.Type = DestinationMssqlSSHTunnelMethodTypeDestinationMssqlSSHTunnelMethodNoTunnel
		return nil
	}

	destinationMssqlSSHTunnelMethodSSHKeyAuthentication := new(DestinationMssqlSSHTunnelMethodSSHKeyAuthentication)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&destinationMssqlSSHTunnelMethodSSHKeyAuthentication); err == nil {
		u.DestinationMssqlSSHTunnelMethodSSHKeyAuthentication = destinationMssqlSSHTunnelMethodSSHKeyAuthentication
		u.Type = DestinationMssqlSSHTunnelMethodTypeDestinationMssqlSSHTunnelMethodSSHKeyAuthentication
		return nil
	}

	destinationMssqlSSHTunnelMethodPasswordAuthentication := new(DestinationMssqlSSHTunnelMethodPasswordAuthentication)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&destinationMssqlSSHTunnelMethodPasswordAuthentication); err == nil {
		u.DestinationMssqlSSHTunnelMethodPasswordAuthentication = destinationMssqlSSHTunnelMethodPasswordAuthentication
		u.Type = DestinationMssqlSSHTunnelMethodTypeDestinationMssqlSSHTunnelMethodPasswordAuthentication
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u DestinationMssqlSSHTunnelMethod) MarshalJSON() ([]byte, error) {
	if u.DestinationMssqlSSHTunnelMethodNoTunnel != nil {
		return json.Marshal(u.DestinationMssqlSSHTunnelMethodNoTunnel)
	}

	if u.DestinationMssqlSSHTunnelMethodSSHKeyAuthentication != nil {
		return json.Marshal(u.DestinationMssqlSSHTunnelMethodSSHKeyAuthentication)
	}

	if u.DestinationMssqlSSHTunnelMethodPasswordAuthentication != nil {
		return json.Marshal(u.DestinationMssqlSSHTunnelMethodPasswordAuthentication)
	}

	return nil, nil
}

// DestinationMssql - The values required to configure the destination.
type DestinationMssql struct {
	// The name of the MSSQL database.
	Database        string                    `json:"database"`
	DestinationType DestinationMssqlMssqlEnum `json:"destinationType"`
	// The host name of the MSSQL database.
	Host string `json:"host"`
	// Additional properties to pass to the JDBC URL string when connecting to the database formatted as 'key=value' pairs separated by the symbol '&'. (example: key1=value1&key2=value2&key3=value3).
	JdbcURLParams *string `json:"jdbc_url_params,omitempty"`
	// The password associated with this username.
	Password *string `json:"password,omitempty"`
	// The port of the MSSQL database.
	Port int64 `json:"port"`
	// The default schema tables are written to if the source does not specify a namespace. The usual value for this field is "public".
	Schema string `json:"schema"`
	// The encryption method which is used to communicate with the database.
	SslMethod *DestinationMssqlSSLMethod `json:"ssl_method,omitempty"`
	// Whether to initiate an SSH tunnel before connecting to the database, and if so, which kind of authentication to use.
	TunnelMethod *DestinationMssqlSSHTunnelMethod `json:"tunnel_method,omitempty"`
	// The username which is used to access the database.
	Username string `json:"username"`
}
