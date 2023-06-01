// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
)

// SourceMssqlUpdateReplicationMethodLogicalReplicationCDCDataToSync - What data should be synced under the CDC. "Existing and New" will read existing data as a snapshot, and sync new changes through CDC. "New Changes Only" will skip the initial snapshot, and only sync new changes through CDC.
type SourceMssqlUpdateReplicationMethodLogicalReplicationCDCDataToSync string

const (
	SourceMssqlUpdateReplicationMethodLogicalReplicationCDCDataToSyncExistingAndNew SourceMssqlUpdateReplicationMethodLogicalReplicationCDCDataToSync = "Existing and New"
	SourceMssqlUpdateReplicationMethodLogicalReplicationCDCDataToSyncNewChangesOnly SourceMssqlUpdateReplicationMethodLogicalReplicationCDCDataToSync = "New Changes Only"
)

func (e SourceMssqlUpdateReplicationMethodLogicalReplicationCDCDataToSync) ToPointer() *SourceMssqlUpdateReplicationMethodLogicalReplicationCDCDataToSync {
	return &e
}

func (e *SourceMssqlUpdateReplicationMethodLogicalReplicationCDCDataToSync) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Existing and New":
		fallthrough
	case "New Changes Only":
		*e = SourceMssqlUpdateReplicationMethodLogicalReplicationCDCDataToSync(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceMssqlUpdateReplicationMethodLogicalReplicationCDCDataToSync: %v", v)
	}
}

type SourceMssqlUpdateReplicationMethodLogicalReplicationCDCMethod string

const (
	SourceMssqlUpdateReplicationMethodLogicalReplicationCDCMethodCdc SourceMssqlUpdateReplicationMethodLogicalReplicationCDCMethod = "CDC"
)

func (e SourceMssqlUpdateReplicationMethodLogicalReplicationCDCMethod) ToPointer() *SourceMssqlUpdateReplicationMethodLogicalReplicationCDCMethod {
	return &e
}

func (e *SourceMssqlUpdateReplicationMethodLogicalReplicationCDCMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "CDC":
		*e = SourceMssqlUpdateReplicationMethodLogicalReplicationCDCMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceMssqlUpdateReplicationMethodLogicalReplicationCDCMethod: %v", v)
	}
}

// SourceMssqlUpdateReplicationMethodLogicalReplicationCDCInitialSnapshotIsolationLevel - Existing data in the database are synced through an initial snapshot. This parameter controls the isolation level that will be used during the initial snapshotting. If you choose the "Snapshot" level, you must enable the <a href="https://docs.microsoft.com/en-us/dotnet/framework/data/adonet/sql/snapshot-isolation-in-sql-server">snapshot isolation mode</a> on the database.
type SourceMssqlUpdateReplicationMethodLogicalReplicationCDCInitialSnapshotIsolationLevel string

const (
	SourceMssqlUpdateReplicationMethodLogicalReplicationCDCInitialSnapshotIsolationLevelSnapshot      SourceMssqlUpdateReplicationMethodLogicalReplicationCDCInitialSnapshotIsolationLevel = "Snapshot"
	SourceMssqlUpdateReplicationMethodLogicalReplicationCDCInitialSnapshotIsolationLevelReadCommitted SourceMssqlUpdateReplicationMethodLogicalReplicationCDCInitialSnapshotIsolationLevel = "Read Committed"
)

func (e SourceMssqlUpdateReplicationMethodLogicalReplicationCDCInitialSnapshotIsolationLevel) ToPointer() *SourceMssqlUpdateReplicationMethodLogicalReplicationCDCInitialSnapshotIsolationLevel {
	return &e
}

func (e *SourceMssqlUpdateReplicationMethodLogicalReplicationCDCInitialSnapshotIsolationLevel) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Snapshot":
		fallthrough
	case "Read Committed":
		*e = SourceMssqlUpdateReplicationMethodLogicalReplicationCDCInitialSnapshotIsolationLevel(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceMssqlUpdateReplicationMethodLogicalReplicationCDCInitialSnapshotIsolationLevel: %v", v)
	}
}

// SourceMssqlUpdateReplicationMethodLogicalReplicationCDC - CDC uses {TBC} to detect inserts, updates, and deletes. This needs to be configured on the source database itself.
type SourceMssqlUpdateReplicationMethodLogicalReplicationCDC struct {
	// What data should be synced under the CDC. "Existing and New" will read existing data as a snapshot, and sync new changes through CDC. "New Changes Only" will skip the initial snapshot, and only sync new changes through CDC.
	DataToSync *SourceMssqlUpdateReplicationMethodLogicalReplicationCDCDataToSync `json:"data_to_sync,omitempty"`
	// The amount of time the connector will wait when it launches to determine if there is new data to sync or not. Defaults to 300 seconds. Valid range: 120 seconds to 1200 seconds. Read about <a href="https://docs.airbyte.com/integrations/sources/mysql/#change-data-capture-cdc">initial waiting time</a>.
	InitialWaitingSeconds *int64                                                        `json:"initial_waiting_seconds,omitempty"`
	Method                SourceMssqlUpdateReplicationMethodLogicalReplicationCDCMethod `json:"method"`
	// Existing data in the database are synced through an initial snapshot. This parameter controls the isolation level that will be used during the initial snapshotting. If you choose the "Snapshot" level, you must enable the <a href="https://docs.microsoft.com/en-us/dotnet/framework/data/adonet/sql/snapshot-isolation-in-sql-server">snapshot isolation mode</a> on the database.
	SnapshotIsolation *SourceMssqlUpdateReplicationMethodLogicalReplicationCDCInitialSnapshotIsolationLevel `json:"snapshot_isolation,omitempty"`
}

type SourceMssqlUpdateReplicationMethodStandardMethod string

const (
	SourceMssqlUpdateReplicationMethodStandardMethodStandard SourceMssqlUpdateReplicationMethodStandardMethod = "STANDARD"
)

func (e SourceMssqlUpdateReplicationMethodStandardMethod) ToPointer() *SourceMssqlUpdateReplicationMethodStandardMethod {
	return &e
}

func (e *SourceMssqlUpdateReplicationMethodStandardMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "STANDARD":
		*e = SourceMssqlUpdateReplicationMethodStandardMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceMssqlUpdateReplicationMethodStandardMethod: %v", v)
	}
}

// SourceMssqlUpdateReplicationMethodStandard - Standard replication requires no setup on the DB side but will not be able to represent deletions incrementally.
type SourceMssqlUpdateReplicationMethodStandard struct {
	Method SourceMssqlUpdateReplicationMethodStandardMethod `json:"method"`
}

type SourceMssqlUpdateReplicationMethodType string

const (
	SourceMssqlUpdateReplicationMethodTypeSourceMssqlUpdateReplicationMethodStandard              SourceMssqlUpdateReplicationMethodType = "source-mssql-update_Replication Method_Standard"
	SourceMssqlUpdateReplicationMethodTypeSourceMssqlUpdateReplicationMethodLogicalReplicationCDC SourceMssqlUpdateReplicationMethodType = "source-mssql-update_Replication Method_Logical Replication (CDC)"
)

type SourceMssqlUpdateReplicationMethod struct {
	SourceMssqlUpdateReplicationMethodStandard              *SourceMssqlUpdateReplicationMethodStandard
	SourceMssqlUpdateReplicationMethodLogicalReplicationCDC *SourceMssqlUpdateReplicationMethodLogicalReplicationCDC

	Type SourceMssqlUpdateReplicationMethodType
}

func CreateSourceMssqlUpdateReplicationMethodSourceMssqlUpdateReplicationMethodStandard(sourceMssqlUpdateReplicationMethodStandard SourceMssqlUpdateReplicationMethodStandard) SourceMssqlUpdateReplicationMethod {
	typ := SourceMssqlUpdateReplicationMethodTypeSourceMssqlUpdateReplicationMethodStandard

	return SourceMssqlUpdateReplicationMethod{
		SourceMssqlUpdateReplicationMethodStandard: &sourceMssqlUpdateReplicationMethodStandard,
		Type: typ,
	}
}

func CreateSourceMssqlUpdateReplicationMethodSourceMssqlUpdateReplicationMethodLogicalReplicationCDC(sourceMssqlUpdateReplicationMethodLogicalReplicationCDC SourceMssqlUpdateReplicationMethodLogicalReplicationCDC) SourceMssqlUpdateReplicationMethod {
	typ := SourceMssqlUpdateReplicationMethodTypeSourceMssqlUpdateReplicationMethodLogicalReplicationCDC

	return SourceMssqlUpdateReplicationMethod{
		SourceMssqlUpdateReplicationMethodLogicalReplicationCDC: &sourceMssqlUpdateReplicationMethodLogicalReplicationCDC,
		Type: typ,
	}
}

func (u *SourceMssqlUpdateReplicationMethod) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	sourceMssqlUpdateReplicationMethodStandard := new(SourceMssqlUpdateReplicationMethodStandard)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&sourceMssqlUpdateReplicationMethodStandard); err == nil {
		u.SourceMssqlUpdateReplicationMethodStandard = sourceMssqlUpdateReplicationMethodStandard
		u.Type = SourceMssqlUpdateReplicationMethodTypeSourceMssqlUpdateReplicationMethodStandard
		return nil
	}

	sourceMssqlUpdateReplicationMethodLogicalReplicationCDC := new(SourceMssqlUpdateReplicationMethodLogicalReplicationCDC)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&sourceMssqlUpdateReplicationMethodLogicalReplicationCDC); err == nil {
		u.SourceMssqlUpdateReplicationMethodLogicalReplicationCDC = sourceMssqlUpdateReplicationMethodLogicalReplicationCDC
		u.Type = SourceMssqlUpdateReplicationMethodTypeSourceMssqlUpdateReplicationMethodLogicalReplicationCDC
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u SourceMssqlUpdateReplicationMethod) MarshalJSON() ([]byte, error) {
	if u.SourceMssqlUpdateReplicationMethodStandard != nil {
		return json.Marshal(u.SourceMssqlUpdateReplicationMethodStandard)
	}

	if u.SourceMssqlUpdateReplicationMethodLogicalReplicationCDC != nil {
		return json.Marshal(u.SourceMssqlUpdateReplicationMethodLogicalReplicationCDC)
	}

	return nil, nil
}

type SourceMssqlUpdateSSLMethodEncryptedVerifyCertificateSSLMethod string

const (
	SourceMssqlUpdateSSLMethodEncryptedVerifyCertificateSSLMethodEncryptedVerifyCertificate SourceMssqlUpdateSSLMethodEncryptedVerifyCertificateSSLMethod = "encrypted_verify_certificate"
)

func (e SourceMssqlUpdateSSLMethodEncryptedVerifyCertificateSSLMethod) ToPointer() *SourceMssqlUpdateSSLMethodEncryptedVerifyCertificateSSLMethod {
	return &e
}

func (e *SourceMssqlUpdateSSLMethodEncryptedVerifyCertificateSSLMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "encrypted_verify_certificate":
		*e = SourceMssqlUpdateSSLMethodEncryptedVerifyCertificateSSLMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceMssqlUpdateSSLMethodEncryptedVerifyCertificateSSLMethod: %v", v)
	}
}

// SourceMssqlUpdateSSLMethodEncryptedVerifyCertificate - Verify and use the certificate provided by the server.
type SourceMssqlUpdateSSLMethodEncryptedVerifyCertificate struct {
	// Specifies the host name of the server. The value of this property must match the subject property of the certificate.
	HostNameInCertificate *string                                                       `json:"hostNameInCertificate,omitempty"`
	SslMethod             SourceMssqlUpdateSSLMethodEncryptedVerifyCertificateSSLMethod `json:"ssl_method"`
}

type SourceMssqlUpdateSSLMethodEncryptedTrustServerCertificateSSLMethod string

const (
	SourceMssqlUpdateSSLMethodEncryptedTrustServerCertificateSSLMethodEncryptedTrustServerCertificate SourceMssqlUpdateSSLMethodEncryptedTrustServerCertificateSSLMethod = "encrypted_trust_server_certificate"
)

func (e SourceMssqlUpdateSSLMethodEncryptedTrustServerCertificateSSLMethod) ToPointer() *SourceMssqlUpdateSSLMethodEncryptedTrustServerCertificateSSLMethod {
	return &e
}

func (e *SourceMssqlUpdateSSLMethodEncryptedTrustServerCertificateSSLMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "encrypted_trust_server_certificate":
		*e = SourceMssqlUpdateSSLMethodEncryptedTrustServerCertificateSSLMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceMssqlUpdateSSLMethodEncryptedTrustServerCertificateSSLMethod: %v", v)
	}
}

// SourceMssqlUpdateSSLMethodEncryptedTrustServerCertificate - Use the certificate provided by the server without verification. (For testing purposes only!)
type SourceMssqlUpdateSSLMethodEncryptedTrustServerCertificate struct {
	SslMethod SourceMssqlUpdateSSLMethodEncryptedTrustServerCertificateSSLMethod `json:"ssl_method"`
}

type SourceMssqlUpdateSSLMethodType string

const (
	SourceMssqlUpdateSSLMethodTypeSourceMssqlUpdateSSLMethodEncryptedTrustServerCertificate SourceMssqlUpdateSSLMethodType = "source-mssql-update_SSL Method_Encrypted (trust server certificate)"
	SourceMssqlUpdateSSLMethodTypeSourceMssqlUpdateSSLMethodEncryptedVerifyCertificate      SourceMssqlUpdateSSLMethodType = "source-mssql-update_SSL Method_Encrypted (verify certificate)"
)

type SourceMssqlUpdateSSLMethod struct {
	SourceMssqlUpdateSSLMethodEncryptedTrustServerCertificate *SourceMssqlUpdateSSLMethodEncryptedTrustServerCertificate
	SourceMssqlUpdateSSLMethodEncryptedVerifyCertificate      *SourceMssqlUpdateSSLMethodEncryptedVerifyCertificate

	Type SourceMssqlUpdateSSLMethodType
}

func CreateSourceMssqlUpdateSSLMethodSourceMssqlUpdateSSLMethodEncryptedTrustServerCertificate(sourceMssqlUpdateSSLMethodEncryptedTrustServerCertificate SourceMssqlUpdateSSLMethodEncryptedTrustServerCertificate) SourceMssqlUpdateSSLMethod {
	typ := SourceMssqlUpdateSSLMethodTypeSourceMssqlUpdateSSLMethodEncryptedTrustServerCertificate

	return SourceMssqlUpdateSSLMethod{
		SourceMssqlUpdateSSLMethodEncryptedTrustServerCertificate: &sourceMssqlUpdateSSLMethodEncryptedTrustServerCertificate,
		Type: typ,
	}
}

func CreateSourceMssqlUpdateSSLMethodSourceMssqlUpdateSSLMethodEncryptedVerifyCertificate(sourceMssqlUpdateSSLMethodEncryptedVerifyCertificate SourceMssqlUpdateSSLMethodEncryptedVerifyCertificate) SourceMssqlUpdateSSLMethod {
	typ := SourceMssqlUpdateSSLMethodTypeSourceMssqlUpdateSSLMethodEncryptedVerifyCertificate

	return SourceMssqlUpdateSSLMethod{
		SourceMssqlUpdateSSLMethodEncryptedVerifyCertificate: &sourceMssqlUpdateSSLMethodEncryptedVerifyCertificate,
		Type: typ,
	}
}

func (u *SourceMssqlUpdateSSLMethod) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	sourceMssqlUpdateSSLMethodEncryptedTrustServerCertificate := new(SourceMssqlUpdateSSLMethodEncryptedTrustServerCertificate)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&sourceMssqlUpdateSSLMethodEncryptedTrustServerCertificate); err == nil {
		u.SourceMssqlUpdateSSLMethodEncryptedTrustServerCertificate = sourceMssqlUpdateSSLMethodEncryptedTrustServerCertificate
		u.Type = SourceMssqlUpdateSSLMethodTypeSourceMssqlUpdateSSLMethodEncryptedTrustServerCertificate
		return nil
	}

	sourceMssqlUpdateSSLMethodEncryptedVerifyCertificate := new(SourceMssqlUpdateSSLMethodEncryptedVerifyCertificate)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&sourceMssqlUpdateSSLMethodEncryptedVerifyCertificate); err == nil {
		u.SourceMssqlUpdateSSLMethodEncryptedVerifyCertificate = sourceMssqlUpdateSSLMethodEncryptedVerifyCertificate
		u.Type = SourceMssqlUpdateSSLMethodTypeSourceMssqlUpdateSSLMethodEncryptedVerifyCertificate
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u SourceMssqlUpdateSSLMethod) MarshalJSON() ([]byte, error) {
	if u.SourceMssqlUpdateSSLMethodEncryptedTrustServerCertificate != nil {
		return json.Marshal(u.SourceMssqlUpdateSSLMethodEncryptedTrustServerCertificate)
	}

	if u.SourceMssqlUpdateSSLMethodEncryptedVerifyCertificate != nil {
		return json.Marshal(u.SourceMssqlUpdateSSLMethodEncryptedVerifyCertificate)
	}

	return nil, nil
}

// SourceMssqlUpdateSSHTunnelMethodPasswordAuthenticationTunnelMethod - Connect through a jump server tunnel host using username and password authentication
type SourceMssqlUpdateSSHTunnelMethodPasswordAuthenticationTunnelMethod string

const (
	SourceMssqlUpdateSSHTunnelMethodPasswordAuthenticationTunnelMethodSSHPasswordAuth SourceMssqlUpdateSSHTunnelMethodPasswordAuthenticationTunnelMethod = "SSH_PASSWORD_AUTH"
)

func (e SourceMssqlUpdateSSHTunnelMethodPasswordAuthenticationTunnelMethod) ToPointer() *SourceMssqlUpdateSSHTunnelMethodPasswordAuthenticationTunnelMethod {
	return &e
}

func (e *SourceMssqlUpdateSSHTunnelMethodPasswordAuthenticationTunnelMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "SSH_PASSWORD_AUTH":
		*e = SourceMssqlUpdateSSHTunnelMethodPasswordAuthenticationTunnelMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceMssqlUpdateSSHTunnelMethodPasswordAuthenticationTunnelMethod: %v", v)
	}
}

// SourceMssqlUpdateSSHTunnelMethodPasswordAuthentication - Whether to initiate an SSH tunnel before connecting to the database, and if so, which kind of authentication to use.
type SourceMssqlUpdateSSHTunnelMethodPasswordAuthentication struct {
	// Hostname of the jump server host that allows inbound ssh tunnel.
	TunnelHost string `json:"tunnel_host"`
	// Connect through a jump server tunnel host using username and password authentication
	TunnelMethod SourceMssqlUpdateSSHTunnelMethodPasswordAuthenticationTunnelMethod `json:"tunnel_method"`
	// Port on the proxy/jump server that accepts inbound ssh connections.
	TunnelPort int64 `json:"tunnel_port"`
	// OS-level username for logging into the jump server host
	TunnelUser string `json:"tunnel_user"`
	// OS-level password for logging into the jump server host
	TunnelUserPassword string `json:"tunnel_user_password"`
}

// SourceMssqlUpdateSSHTunnelMethodSSHKeyAuthenticationTunnelMethod - Connect through a jump server tunnel host using username and ssh key
type SourceMssqlUpdateSSHTunnelMethodSSHKeyAuthenticationTunnelMethod string

const (
	SourceMssqlUpdateSSHTunnelMethodSSHKeyAuthenticationTunnelMethodSSHKeyAuth SourceMssqlUpdateSSHTunnelMethodSSHKeyAuthenticationTunnelMethod = "SSH_KEY_AUTH"
)

func (e SourceMssqlUpdateSSHTunnelMethodSSHKeyAuthenticationTunnelMethod) ToPointer() *SourceMssqlUpdateSSHTunnelMethodSSHKeyAuthenticationTunnelMethod {
	return &e
}

func (e *SourceMssqlUpdateSSHTunnelMethodSSHKeyAuthenticationTunnelMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "SSH_KEY_AUTH":
		*e = SourceMssqlUpdateSSHTunnelMethodSSHKeyAuthenticationTunnelMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceMssqlUpdateSSHTunnelMethodSSHKeyAuthenticationTunnelMethod: %v", v)
	}
}

// SourceMssqlUpdateSSHTunnelMethodSSHKeyAuthentication - Whether to initiate an SSH tunnel before connecting to the database, and if so, which kind of authentication to use.
type SourceMssqlUpdateSSHTunnelMethodSSHKeyAuthentication struct {
	// OS-level user account ssh key credentials in RSA PEM format ( created with ssh-keygen -t rsa -m PEM -f myuser_rsa )
	SSHKey string `json:"ssh_key"`
	// Hostname of the jump server host that allows inbound ssh tunnel.
	TunnelHost string `json:"tunnel_host"`
	// Connect through a jump server tunnel host using username and ssh key
	TunnelMethod SourceMssqlUpdateSSHTunnelMethodSSHKeyAuthenticationTunnelMethod `json:"tunnel_method"`
	// Port on the proxy/jump server that accepts inbound ssh connections.
	TunnelPort int64 `json:"tunnel_port"`
	// OS-level username for logging into the jump server host.
	TunnelUser string `json:"tunnel_user"`
}

// SourceMssqlUpdateSSHTunnelMethodNoTunnelTunnelMethod - No ssh tunnel needed to connect to database
type SourceMssqlUpdateSSHTunnelMethodNoTunnelTunnelMethod string

const (
	SourceMssqlUpdateSSHTunnelMethodNoTunnelTunnelMethodNoTunnel SourceMssqlUpdateSSHTunnelMethodNoTunnelTunnelMethod = "NO_TUNNEL"
)

func (e SourceMssqlUpdateSSHTunnelMethodNoTunnelTunnelMethod) ToPointer() *SourceMssqlUpdateSSHTunnelMethodNoTunnelTunnelMethod {
	return &e
}

func (e *SourceMssqlUpdateSSHTunnelMethodNoTunnelTunnelMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "NO_TUNNEL":
		*e = SourceMssqlUpdateSSHTunnelMethodNoTunnelTunnelMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceMssqlUpdateSSHTunnelMethodNoTunnelTunnelMethod: %v", v)
	}
}

// SourceMssqlUpdateSSHTunnelMethodNoTunnel - Whether to initiate an SSH tunnel before connecting to the database, and if so, which kind of authentication to use.
type SourceMssqlUpdateSSHTunnelMethodNoTunnel struct {
	// No ssh tunnel needed to connect to database
	TunnelMethod SourceMssqlUpdateSSHTunnelMethodNoTunnelTunnelMethod `json:"tunnel_method"`
}

type SourceMssqlUpdateSSHTunnelMethodType string

const (
	SourceMssqlUpdateSSHTunnelMethodTypeSourceMssqlUpdateSSHTunnelMethodNoTunnel               SourceMssqlUpdateSSHTunnelMethodType = "source-mssql-update_SSH Tunnel Method_No Tunnel"
	SourceMssqlUpdateSSHTunnelMethodTypeSourceMssqlUpdateSSHTunnelMethodSSHKeyAuthentication   SourceMssqlUpdateSSHTunnelMethodType = "source-mssql-update_SSH Tunnel Method_SSH Key Authentication"
	SourceMssqlUpdateSSHTunnelMethodTypeSourceMssqlUpdateSSHTunnelMethodPasswordAuthentication SourceMssqlUpdateSSHTunnelMethodType = "source-mssql-update_SSH Tunnel Method_Password Authentication"
)

type SourceMssqlUpdateSSHTunnelMethod struct {
	SourceMssqlUpdateSSHTunnelMethodNoTunnel               *SourceMssqlUpdateSSHTunnelMethodNoTunnel
	SourceMssqlUpdateSSHTunnelMethodSSHKeyAuthentication   *SourceMssqlUpdateSSHTunnelMethodSSHKeyAuthentication
	SourceMssqlUpdateSSHTunnelMethodPasswordAuthentication *SourceMssqlUpdateSSHTunnelMethodPasswordAuthentication

	Type SourceMssqlUpdateSSHTunnelMethodType
}

func CreateSourceMssqlUpdateSSHTunnelMethodSourceMssqlUpdateSSHTunnelMethodNoTunnel(sourceMssqlUpdateSSHTunnelMethodNoTunnel SourceMssqlUpdateSSHTunnelMethodNoTunnel) SourceMssqlUpdateSSHTunnelMethod {
	typ := SourceMssqlUpdateSSHTunnelMethodTypeSourceMssqlUpdateSSHTunnelMethodNoTunnel

	return SourceMssqlUpdateSSHTunnelMethod{
		SourceMssqlUpdateSSHTunnelMethodNoTunnel: &sourceMssqlUpdateSSHTunnelMethodNoTunnel,
		Type:                                     typ,
	}
}

func CreateSourceMssqlUpdateSSHTunnelMethodSourceMssqlUpdateSSHTunnelMethodSSHKeyAuthentication(sourceMssqlUpdateSSHTunnelMethodSSHKeyAuthentication SourceMssqlUpdateSSHTunnelMethodSSHKeyAuthentication) SourceMssqlUpdateSSHTunnelMethod {
	typ := SourceMssqlUpdateSSHTunnelMethodTypeSourceMssqlUpdateSSHTunnelMethodSSHKeyAuthentication

	return SourceMssqlUpdateSSHTunnelMethod{
		SourceMssqlUpdateSSHTunnelMethodSSHKeyAuthentication: &sourceMssqlUpdateSSHTunnelMethodSSHKeyAuthentication,
		Type: typ,
	}
}

func CreateSourceMssqlUpdateSSHTunnelMethodSourceMssqlUpdateSSHTunnelMethodPasswordAuthentication(sourceMssqlUpdateSSHTunnelMethodPasswordAuthentication SourceMssqlUpdateSSHTunnelMethodPasswordAuthentication) SourceMssqlUpdateSSHTunnelMethod {
	typ := SourceMssqlUpdateSSHTunnelMethodTypeSourceMssqlUpdateSSHTunnelMethodPasswordAuthentication

	return SourceMssqlUpdateSSHTunnelMethod{
		SourceMssqlUpdateSSHTunnelMethodPasswordAuthentication: &sourceMssqlUpdateSSHTunnelMethodPasswordAuthentication,
		Type: typ,
	}
}

func (u *SourceMssqlUpdateSSHTunnelMethod) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	sourceMssqlUpdateSSHTunnelMethodNoTunnel := new(SourceMssqlUpdateSSHTunnelMethodNoTunnel)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&sourceMssqlUpdateSSHTunnelMethodNoTunnel); err == nil {
		u.SourceMssqlUpdateSSHTunnelMethodNoTunnel = sourceMssqlUpdateSSHTunnelMethodNoTunnel
		u.Type = SourceMssqlUpdateSSHTunnelMethodTypeSourceMssqlUpdateSSHTunnelMethodNoTunnel
		return nil
	}

	sourceMssqlUpdateSSHTunnelMethodSSHKeyAuthentication := new(SourceMssqlUpdateSSHTunnelMethodSSHKeyAuthentication)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&sourceMssqlUpdateSSHTunnelMethodSSHKeyAuthentication); err == nil {
		u.SourceMssqlUpdateSSHTunnelMethodSSHKeyAuthentication = sourceMssqlUpdateSSHTunnelMethodSSHKeyAuthentication
		u.Type = SourceMssqlUpdateSSHTunnelMethodTypeSourceMssqlUpdateSSHTunnelMethodSSHKeyAuthentication
		return nil
	}

	sourceMssqlUpdateSSHTunnelMethodPasswordAuthentication := new(SourceMssqlUpdateSSHTunnelMethodPasswordAuthentication)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&sourceMssqlUpdateSSHTunnelMethodPasswordAuthentication); err == nil {
		u.SourceMssqlUpdateSSHTunnelMethodPasswordAuthentication = sourceMssqlUpdateSSHTunnelMethodPasswordAuthentication
		u.Type = SourceMssqlUpdateSSHTunnelMethodTypeSourceMssqlUpdateSSHTunnelMethodPasswordAuthentication
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u SourceMssqlUpdateSSHTunnelMethod) MarshalJSON() ([]byte, error) {
	if u.SourceMssqlUpdateSSHTunnelMethodNoTunnel != nil {
		return json.Marshal(u.SourceMssqlUpdateSSHTunnelMethodNoTunnel)
	}

	if u.SourceMssqlUpdateSSHTunnelMethodSSHKeyAuthentication != nil {
		return json.Marshal(u.SourceMssqlUpdateSSHTunnelMethodSSHKeyAuthentication)
	}

	if u.SourceMssqlUpdateSSHTunnelMethodPasswordAuthentication != nil {
		return json.Marshal(u.SourceMssqlUpdateSSHTunnelMethodPasswordAuthentication)
	}

	return nil, nil
}

type SourceMssqlUpdate struct {
	// The name of the database.
	Database string `json:"database"`
	// The hostname of the database.
	Host string `json:"host"`
	// Additional properties to pass to the JDBC URL string when connecting to the database formatted as 'key=value' pairs separated by the symbol '&'. (example: key1=value1&key2=value2&key3=value3).
	JdbcURLParams *string `json:"jdbc_url_params,omitempty"`
	// The password associated with the username.
	Password *string `json:"password,omitempty"`
	// The port of the database.
	Port int64 `json:"port"`
	// The replication method used for extracting data from the database. STANDARD replication requires no setup on the DB side but will not be able to represent deletions incrementally. CDC uses {TBC} to detect inserts, updates, and deletes. This needs to be configured on the source database itself.
	ReplicationMethod *SourceMssqlUpdateReplicationMethod `json:"replication_method,omitempty"`
	// The list of schemas to sync from. Defaults to user. Case sensitive.
	Schemas []string `json:"schemas,omitempty"`
	// The encryption method which is used when communicating with the database.
	SslMethod *SourceMssqlUpdateSSLMethod `json:"ssl_method,omitempty"`
	// Whether to initiate an SSH tunnel before connecting to the database, and if so, which kind of authentication to use.
	TunnelMethod *SourceMssqlUpdateSSHTunnelMethod `json:"tunnel_method,omitempty"`
	// The username which is used to access the database.
	Username string `json:"username"`
}