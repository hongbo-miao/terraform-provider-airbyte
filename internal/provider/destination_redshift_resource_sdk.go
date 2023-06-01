// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *DestinationRedshiftResourceModel) ToCreateSDKType() *shared.DestinationRedshiftCreateRequest {
	database := r.Configuration.Database.ValueString()
	destinationType := shared.DestinationRedshiftRedshift(r.Configuration.DestinationType.ValueString())
	host := r.Configuration.Host.ValueString()
	jdbcURLParams := new(string)
	if !r.Configuration.JdbcURLParams.IsUnknown() && !r.Configuration.JdbcURLParams.IsNull() {
		*jdbcURLParams = r.Configuration.JdbcURLParams.ValueString()
	} else {
		jdbcURLParams = nil
	}
	password := r.Configuration.Password.ValueString()
	port := r.Configuration.Port.ValueInt64()
	schema := r.Configuration.Schema.ValueString()
	var tunnelMethod *shared.DestinationRedshiftSSHTunnelMethod
	var destinationRedshiftSSHTunnelMethodNoTunnel *shared.DestinationRedshiftSSHTunnelMethodNoTunnel
	if r.Configuration.TunnelMethod.DestinationRedshiftSSHTunnelMethodNoTunnel != nil {
		tunnelMethod1 := shared.DestinationRedshiftSSHTunnelMethodNoTunnelTunnelMethod(r.Configuration.TunnelMethod.DestinationRedshiftSSHTunnelMethodNoTunnel.TunnelMethod.ValueString())
		destinationRedshiftSSHTunnelMethodNoTunnel = &shared.DestinationRedshiftSSHTunnelMethodNoTunnel{
			TunnelMethod: tunnelMethod1,
		}
	}
	if destinationRedshiftSSHTunnelMethodNoTunnel != nil {
		tunnelMethod = &shared.DestinationRedshiftSSHTunnelMethod{
			DestinationRedshiftSSHTunnelMethodNoTunnel: destinationRedshiftSSHTunnelMethodNoTunnel,
		}
	}
	var destinationRedshiftSSHTunnelMethodSSHKeyAuthentication *shared.DestinationRedshiftSSHTunnelMethodSSHKeyAuthentication
	if r.Configuration.TunnelMethod.DestinationRedshiftSSHTunnelMethodSSHKeyAuthentication != nil {
		sshKey := r.Configuration.TunnelMethod.DestinationRedshiftSSHTunnelMethodSSHKeyAuthentication.SSHKey.ValueString()
		tunnelHost := r.Configuration.TunnelMethod.DestinationRedshiftSSHTunnelMethodSSHKeyAuthentication.TunnelHost.ValueString()
		tunnelMethod2 := shared.DestinationRedshiftSSHTunnelMethodSSHKeyAuthenticationTunnelMethod(r.Configuration.TunnelMethod.DestinationRedshiftSSHTunnelMethodSSHKeyAuthentication.TunnelMethod.ValueString())
		tunnelPort := r.Configuration.TunnelMethod.DestinationRedshiftSSHTunnelMethodSSHKeyAuthentication.TunnelPort.ValueInt64()
		tunnelUser := r.Configuration.TunnelMethod.DestinationRedshiftSSHTunnelMethodSSHKeyAuthentication.TunnelUser.ValueString()
		destinationRedshiftSSHTunnelMethodSSHKeyAuthentication = &shared.DestinationRedshiftSSHTunnelMethodSSHKeyAuthentication{
			SSHKey:       sshKey,
			TunnelHost:   tunnelHost,
			TunnelMethod: tunnelMethod2,
			TunnelPort:   tunnelPort,
			TunnelUser:   tunnelUser,
		}
	}
	if destinationRedshiftSSHTunnelMethodSSHKeyAuthentication != nil {
		tunnelMethod = &shared.DestinationRedshiftSSHTunnelMethod{
			DestinationRedshiftSSHTunnelMethodSSHKeyAuthentication: destinationRedshiftSSHTunnelMethodSSHKeyAuthentication,
		}
	}
	var destinationRedshiftSSHTunnelMethodPasswordAuthentication *shared.DestinationRedshiftSSHTunnelMethodPasswordAuthentication
	if r.Configuration.TunnelMethod.DestinationRedshiftSSHTunnelMethodPasswordAuthentication != nil {
		tunnelHost1 := r.Configuration.TunnelMethod.DestinationRedshiftSSHTunnelMethodPasswordAuthentication.TunnelHost.ValueString()
		tunnelMethod3 := shared.DestinationRedshiftSSHTunnelMethodPasswordAuthenticationTunnelMethod(r.Configuration.TunnelMethod.DestinationRedshiftSSHTunnelMethodPasswordAuthentication.TunnelMethod.ValueString())
		tunnelPort1 := r.Configuration.TunnelMethod.DestinationRedshiftSSHTunnelMethodPasswordAuthentication.TunnelPort.ValueInt64()
		tunnelUser1 := r.Configuration.TunnelMethod.DestinationRedshiftSSHTunnelMethodPasswordAuthentication.TunnelUser.ValueString()
		tunnelUserPassword := r.Configuration.TunnelMethod.DestinationRedshiftSSHTunnelMethodPasswordAuthentication.TunnelUserPassword.ValueString()
		destinationRedshiftSSHTunnelMethodPasswordAuthentication = &shared.DestinationRedshiftSSHTunnelMethodPasswordAuthentication{
			TunnelHost:         tunnelHost1,
			TunnelMethod:       tunnelMethod3,
			TunnelPort:         tunnelPort1,
			TunnelUser:         tunnelUser1,
			TunnelUserPassword: tunnelUserPassword,
		}
	}
	if destinationRedshiftSSHTunnelMethodPasswordAuthentication != nil {
		tunnelMethod = &shared.DestinationRedshiftSSHTunnelMethod{
			DestinationRedshiftSSHTunnelMethodPasswordAuthentication: destinationRedshiftSSHTunnelMethodPasswordAuthentication,
		}
	}
	var uploadingMethod *shared.DestinationRedshiftUploadingMethod
	var destinationRedshiftUploadingMethodStandard *shared.DestinationRedshiftUploadingMethodStandard
	if r.Configuration.UploadingMethod.DestinationRedshiftUploadingMethodStandard != nil {
		method := shared.DestinationRedshiftUploadingMethodStandardMethod(r.Configuration.UploadingMethod.DestinationRedshiftUploadingMethodStandard.Method.ValueString())
		destinationRedshiftUploadingMethodStandard = &shared.DestinationRedshiftUploadingMethodStandard{
			Method: method,
		}
	}
	if destinationRedshiftUploadingMethodStandard != nil {
		uploadingMethod = &shared.DestinationRedshiftUploadingMethod{
			DestinationRedshiftUploadingMethodStandard: destinationRedshiftUploadingMethodStandard,
		}
	}
	var destinationRedshiftUploadingMethodS3Staging *shared.DestinationRedshiftUploadingMethodS3Staging
	if r.Configuration.UploadingMethod.DestinationRedshiftUploadingMethodS3Staging != nil {
		accessKeyID := r.Configuration.UploadingMethod.DestinationRedshiftUploadingMethodS3Staging.AccessKeyID.ValueString()
		var encryption *shared.DestinationRedshiftUploadingMethodS3StagingEncryption
		var destinationRedshiftUploadingMethodS3StagingEncryptionNoEncryption *shared.DestinationRedshiftUploadingMethodS3StagingEncryptionNoEncryption
		if r.Configuration.UploadingMethod.DestinationRedshiftUploadingMethodS3Staging.Encryption.DestinationRedshiftUploadingMethodS3StagingEncryptionNoEncryption != nil {
			encryptionType := shared.DestinationRedshiftUploadingMethodS3StagingEncryptionNoEncryptionEncryptionType(r.Configuration.UploadingMethod.DestinationRedshiftUploadingMethodS3Staging.Encryption.DestinationRedshiftUploadingMethodS3StagingEncryptionNoEncryption.EncryptionType.ValueString())
			destinationRedshiftUploadingMethodS3StagingEncryptionNoEncryption = &shared.DestinationRedshiftUploadingMethodS3StagingEncryptionNoEncryption{
				EncryptionType: encryptionType,
			}
		}
		if destinationRedshiftUploadingMethodS3StagingEncryptionNoEncryption != nil {
			encryption = &shared.DestinationRedshiftUploadingMethodS3StagingEncryption{
				DestinationRedshiftUploadingMethodS3StagingEncryptionNoEncryption: destinationRedshiftUploadingMethodS3StagingEncryptionNoEncryption,
			}
		}
		var destinationRedshiftUploadingMethodS3StagingEncryptionAESCBCEnvelopeEncryption *shared.DestinationRedshiftUploadingMethodS3StagingEncryptionAESCBCEnvelopeEncryption
		if r.Configuration.UploadingMethod.DestinationRedshiftUploadingMethodS3Staging.Encryption.DestinationRedshiftUploadingMethodS3StagingEncryptionAESCBCEnvelopeEncryption != nil {
			encryptionType1 := shared.DestinationRedshiftUploadingMethodS3StagingEncryptionAESCBCEnvelopeEncryptionEncryptionType(r.Configuration.UploadingMethod.DestinationRedshiftUploadingMethodS3Staging.Encryption.DestinationRedshiftUploadingMethodS3StagingEncryptionAESCBCEnvelopeEncryption.EncryptionType.ValueString())
			keyEncryptingKey := new(string)
			if !r.Configuration.UploadingMethod.DestinationRedshiftUploadingMethodS3Staging.Encryption.DestinationRedshiftUploadingMethodS3StagingEncryptionAESCBCEnvelopeEncryption.KeyEncryptingKey.IsUnknown() && !r.Configuration.UploadingMethod.DestinationRedshiftUploadingMethodS3Staging.Encryption.DestinationRedshiftUploadingMethodS3StagingEncryptionAESCBCEnvelopeEncryption.KeyEncryptingKey.IsNull() {
				*keyEncryptingKey = r.Configuration.UploadingMethod.DestinationRedshiftUploadingMethodS3Staging.Encryption.DestinationRedshiftUploadingMethodS3StagingEncryptionAESCBCEnvelopeEncryption.KeyEncryptingKey.ValueString()
			} else {
				keyEncryptingKey = nil
			}
			destinationRedshiftUploadingMethodS3StagingEncryptionAESCBCEnvelopeEncryption = &shared.DestinationRedshiftUploadingMethodS3StagingEncryptionAESCBCEnvelopeEncryption{
				EncryptionType:   encryptionType1,
				KeyEncryptingKey: keyEncryptingKey,
			}
		}
		if destinationRedshiftUploadingMethodS3StagingEncryptionAESCBCEnvelopeEncryption != nil {
			encryption = &shared.DestinationRedshiftUploadingMethodS3StagingEncryption{
				DestinationRedshiftUploadingMethodS3StagingEncryptionAESCBCEnvelopeEncryption: destinationRedshiftUploadingMethodS3StagingEncryptionAESCBCEnvelopeEncryption,
			}
		}
		fileBufferCount := new(int64)
		if !r.Configuration.UploadingMethod.DestinationRedshiftUploadingMethodS3Staging.FileBufferCount.IsUnknown() && !r.Configuration.UploadingMethod.DestinationRedshiftUploadingMethodS3Staging.FileBufferCount.IsNull() {
			*fileBufferCount = r.Configuration.UploadingMethod.DestinationRedshiftUploadingMethodS3Staging.FileBufferCount.ValueInt64()
		} else {
			fileBufferCount = nil
		}
		fileNamePattern := new(string)
		if !r.Configuration.UploadingMethod.DestinationRedshiftUploadingMethodS3Staging.FileNamePattern.IsUnknown() && !r.Configuration.UploadingMethod.DestinationRedshiftUploadingMethodS3Staging.FileNamePattern.IsNull() {
			*fileNamePattern = r.Configuration.UploadingMethod.DestinationRedshiftUploadingMethodS3Staging.FileNamePattern.ValueString()
		} else {
			fileNamePattern = nil
		}
		method1 := shared.DestinationRedshiftUploadingMethodS3StagingMethod(r.Configuration.UploadingMethod.DestinationRedshiftUploadingMethodS3Staging.Method.ValueString())
		purgeStagingData := new(bool)
		if !r.Configuration.UploadingMethod.DestinationRedshiftUploadingMethodS3Staging.PurgeStagingData.IsUnknown() && !r.Configuration.UploadingMethod.DestinationRedshiftUploadingMethodS3Staging.PurgeStagingData.IsNull() {
			*purgeStagingData = r.Configuration.UploadingMethod.DestinationRedshiftUploadingMethodS3Staging.PurgeStagingData.ValueBool()
		} else {
			purgeStagingData = nil
		}
		s3BucketName := r.Configuration.UploadingMethod.DestinationRedshiftUploadingMethodS3Staging.S3BucketName.ValueString()
		s3BucketPath := new(string)
		if !r.Configuration.UploadingMethod.DestinationRedshiftUploadingMethodS3Staging.S3BucketPath.IsUnknown() && !r.Configuration.UploadingMethod.DestinationRedshiftUploadingMethodS3Staging.S3BucketPath.IsNull() {
			*s3BucketPath = r.Configuration.UploadingMethod.DestinationRedshiftUploadingMethodS3Staging.S3BucketPath.ValueString()
		} else {
			s3BucketPath = nil
		}
		s3BucketRegion := shared.DestinationRedshiftUploadingMethodS3StagingS3BucketRegion(r.Configuration.UploadingMethod.DestinationRedshiftUploadingMethodS3Staging.S3BucketRegion.ValueString())
		secretAccessKey := r.Configuration.UploadingMethod.DestinationRedshiftUploadingMethodS3Staging.SecretAccessKey.ValueString()
		destinationRedshiftUploadingMethodS3Staging = &shared.DestinationRedshiftUploadingMethodS3Staging{
			AccessKeyID:      accessKeyID,
			Encryption:       encryption,
			FileBufferCount:  fileBufferCount,
			FileNamePattern:  fileNamePattern,
			Method:           method1,
			PurgeStagingData: purgeStagingData,
			S3BucketName:     s3BucketName,
			S3BucketPath:     s3BucketPath,
			S3BucketRegion:   s3BucketRegion,
			SecretAccessKey:  secretAccessKey,
		}
	}
	if destinationRedshiftUploadingMethodS3Staging != nil {
		uploadingMethod = &shared.DestinationRedshiftUploadingMethod{
			DestinationRedshiftUploadingMethodS3Staging: destinationRedshiftUploadingMethodS3Staging,
		}
	}
	username := r.Configuration.Username.ValueString()
	configuration := shared.DestinationRedshift{
		Database:        database,
		DestinationType: destinationType,
		Host:            host,
		JdbcURLParams:   jdbcURLParams,
		Password:        password,
		Port:            port,
		Schema:          schema,
		TunnelMethod:    tunnelMethod,
		UploadingMethod: uploadingMethod,
		Username:        username,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.DestinationRedshiftCreateRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *DestinationRedshiftResourceModel) ToUpdateSDKType() *shared.DestinationRedshiftPutRequest {
	database := r.Configuration.Database.ValueString()
	host := r.Configuration.Host.ValueString()
	jdbcURLParams := new(string)
	if !r.Configuration.JdbcURLParams.IsUnknown() && !r.Configuration.JdbcURLParams.IsNull() {
		*jdbcURLParams = r.Configuration.JdbcURLParams.ValueString()
	} else {
		jdbcURLParams = nil
	}
	password := r.Configuration.Password.ValueString()
	port := r.Configuration.Port.ValueInt64()
	schema := r.Configuration.Schema.ValueString()
	var tunnelMethod *shared.DestinationRedshiftUpdateSSHTunnelMethod
	var destinationRedshiftUpdateSSHTunnelMethodNoTunnel *shared.DestinationRedshiftUpdateSSHTunnelMethodNoTunnel
	if r.Configuration.TunnelMethod.DestinationRedshiftSSHTunnelMethodNoTunnel != nil {
		tunnelMethod1 := shared.DestinationRedshiftUpdateSSHTunnelMethodNoTunnelTunnelMethod(r.Configuration.TunnelMethod.DestinationRedshiftSSHTunnelMethodNoTunnel.TunnelMethod.ValueString())
		destinationRedshiftUpdateSSHTunnelMethodNoTunnel = &shared.DestinationRedshiftUpdateSSHTunnelMethodNoTunnel{
			TunnelMethod: tunnelMethod1,
		}
	}
	if destinationRedshiftUpdateSSHTunnelMethodNoTunnel != nil {
		tunnelMethod = &shared.DestinationRedshiftUpdateSSHTunnelMethod{
			DestinationRedshiftUpdateSSHTunnelMethodNoTunnel: destinationRedshiftUpdateSSHTunnelMethodNoTunnel,
		}
	}
	var destinationRedshiftUpdateSSHTunnelMethodSSHKeyAuthentication *shared.DestinationRedshiftUpdateSSHTunnelMethodSSHKeyAuthentication
	if r.Configuration.TunnelMethod.DestinationRedshiftSSHTunnelMethodSSHKeyAuthentication != nil {
		sshKey := r.Configuration.TunnelMethod.DestinationRedshiftSSHTunnelMethodSSHKeyAuthentication.SSHKey.ValueString()
		tunnelHost := r.Configuration.TunnelMethod.DestinationRedshiftSSHTunnelMethodSSHKeyAuthentication.TunnelHost.ValueString()
		tunnelMethod2 := shared.DestinationRedshiftUpdateSSHTunnelMethodSSHKeyAuthenticationTunnelMethod(r.Configuration.TunnelMethod.DestinationRedshiftSSHTunnelMethodSSHKeyAuthentication.TunnelMethod.ValueString())
		tunnelPort := r.Configuration.TunnelMethod.DestinationRedshiftSSHTunnelMethodSSHKeyAuthentication.TunnelPort.ValueInt64()
		tunnelUser := r.Configuration.TunnelMethod.DestinationRedshiftSSHTunnelMethodSSHKeyAuthentication.TunnelUser.ValueString()
		destinationRedshiftUpdateSSHTunnelMethodSSHKeyAuthentication = &shared.DestinationRedshiftUpdateSSHTunnelMethodSSHKeyAuthentication{
			SSHKey:       sshKey,
			TunnelHost:   tunnelHost,
			TunnelMethod: tunnelMethod2,
			TunnelPort:   tunnelPort,
			TunnelUser:   tunnelUser,
		}
	}
	if destinationRedshiftUpdateSSHTunnelMethodSSHKeyAuthentication != nil {
		tunnelMethod = &shared.DestinationRedshiftUpdateSSHTunnelMethod{
			DestinationRedshiftUpdateSSHTunnelMethodSSHKeyAuthentication: destinationRedshiftUpdateSSHTunnelMethodSSHKeyAuthentication,
		}
	}
	var destinationRedshiftUpdateSSHTunnelMethodPasswordAuthentication *shared.DestinationRedshiftUpdateSSHTunnelMethodPasswordAuthentication
	if r.Configuration.TunnelMethod.DestinationRedshiftSSHTunnelMethodPasswordAuthentication != nil {
		tunnelHost1 := r.Configuration.TunnelMethod.DestinationRedshiftSSHTunnelMethodPasswordAuthentication.TunnelHost.ValueString()
		tunnelMethod3 := shared.DestinationRedshiftUpdateSSHTunnelMethodPasswordAuthenticationTunnelMethod(r.Configuration.TunnelMethod.DestinationRedshiftSSHTunnelMethodPasswordAuthentication.TunnelMethod.ValueString())
		tunnelPort1 := r.Configuration.TunnelMethod.DestinationRedshiftSSHTunnelMethodPasswordAuthentication.TunnelPort.ValueInt64()
		tunnelUser1 := r.Configuration.TunnelMethod.DestinationRedshiftSSHTunnelMethodPasswordAuthentication.TunnelUser.ValueString()
		tunnelUserPassword := r.Configuration.TunnelMethod.DestinationRedshiftSSHTunnelMethodPasswordAuthentication.TunnelUserPassword.ValueString()
		destinationRedshiftUpdateSSHTunnelMethodPasswordAuthentication = &shared.DestinationRedshiftUpdateSSHTunnelMethodPasswordAuthentication{
			TunnelHost:         tunnelHost1,
			TunnelMethod:       tunnelMethod3,
			TunnelPort:         tunnelPort1,
			TunnelUser:         tunnelUser1,
			TunnelUserPassword: tunnelUserPassword,
		}
	}
	if destinationRedshiftUpdateSSHTunnelMethodPasswordAuthentication != nil {
		tunnelMethod = &shared.DestinationRedshiftUpdateSSHTunnelMethod{
			DestinationRedshiftUpdateSSHTunnelMethodPasswordAuthentication: destinationRedshiftUpdateSSHTunnelMethodPasswordAuthentication,
		}
	}
	var uploadingMethod *shared.DestinationRedshiftUpdateUploadingMethod
	var destinationRedshiftUpdateUploadingMethodStandard *shared.DestinationRedshiftUpdateUploadingMethodStandard
	if r.Configuration.UploadingMethod.DestinationRedshiftUploadingMethodStandard != nil {
		method := shared.DestinationRedshiftUpdateUploadingMethodStandardMethod(r.Configuration.UploadingMethod.DestinationRedshiftUploadingMethodStandard.Method.ValueString())
		destinationRedshiftUpdateUploadingMethodStandard = &shared.DestinationRedshiftUpdateUploadingMethodStandard{
			Method: method,
		}
	}
	if destinationRedshiftUpdateUploadingMethodStandard != nil {
		uploadingMethod = &shared.DestinationRedshiftUpdateUploadingMethod{
			DestinationRedshiftUpdateUploadingMethodStandard: destinationRedshiftUpdateUploadingMethodStandard,
		}
	}
	var destinationRedshiftUpdateUploadingMethodS3Staging *shared.DestinationRedshiftUpdateUploadingMethodS3Staging
	if r.Configuration.UploadingMethod.DestinationRedshiftUploadingMethodS3Staging != nil {
		accessKeyID := r.Configuration.UploadingMethod.DestinationRedshiftUploadingMethodS3Staging.AccessKeyID.ValueString()
		var encryption *shared.DestinationRedshiftUpdateUploadingMethodS3StagingEncryption
		var destinationRedshiftUpdateUploadingMethodS3StagingEncryptionNoEncryption *shared.DestinationRedshiftUpdateUploadingMethodS3StagingEncryptionNoEncryption
		if r.Configuration.UploadingMethod.DestinationRedshiftUploadingMethodS3Staging.Encryption.DestinationRedshiftUploadingMethodS3StagingEncryptionNoEncryption != nil {
			encryptionType := shared.DestinationRedshiftUpdateUploadingMethodS3StagingEncryptionNoEncryptionEncryptionType(r.Configuration.UploadingMethod.DestinationRedshiftUploadingMethodS3Staging.Encryption.DestinationRedshiftUploadingMethodS3StagingEncryptionNoEncryption.EncryptionType.ValueString())
			destinationRedshiftUpdateUploadingMethodS3StagingEncryptionNoEncryption = &shared.DestinationRedshiftUpdateUploadingMethodS3StagingEncryptionNoEncryption{
				EncryptionType: encryptionType,
			}
		}
		if destinationRedshiftUpdateUploadingMethodS3StagingEncryptionNoEncryption != nil {
			encryption = &shared.DestinationRedshiftUpdateUploadingMethodS3StagingEncryption{
				DestinationRedshiftUpdateUploadingMethodS3StagingEncryptionNoEncryption: destinationRedshiftUpdateUploadingMethodS3StagingEncryptionNoEncryption,
			}
		}
		var destinationRedshiftUpdateUploadingMethodS3StagingEncryptionAESCBCEnvelopeEncryption *shared.DestinationRedshiftUpdateUploadingMethodS3StagingEncryptionAESCBCEnvelopeEncryption
		if r.Configuration.UploadingMethod.DestinationRedshiftUploadingMethodS3Staging.Encryption.DestinationRedshiftUploadingMethodS3StagingEncryptionAESCBCEnvelopeEncryption != nil {
			encryptionType1 := shared.DestinationRedshiftUpdateUploadingMethodS3StagingEncryptionAESCBCEnvelopeEncryptionEncryptionType(r.Configuration.UploadingMethod.DestinationRedshiftUploadingMethodS3Staging.Encryption.DestinationRedshiftUploadingMethodS3StagingEncryptionAESCBCEnvelopeEncryption.EncryptionType.ValueString())
			keyEncryptingKey := new(string)
			if !r.Configuration.UploadingMethod.DestinationRedshiftUploadingMethodS3Staging.Encryption.DestinationRedshiftUploadingMethodS3StagingEncryptionAESCBCEnvelopeEncryption.KeyEncryptingKey.IsUnknown() && !r.Configuration.UploadingMethod.DestinationRedshiftUploadingMethodS3Staging.Encryption.DestinationRedshiftUploadingMethodS3StagingEncryptionAESCBCEnvelopeEncryption.KeyEncryptingKey.IsNull() {
				*keyEncryptingKey = r.Configuration.UploadingMethod.DestinationRedshiftUploadingMethodS3Staging.Encryption.DestinationRedshiftUploadingMethodS3StagingEncryptionAESCBCEnvelopeEncryption.KeyEncryptingKey.ValueString()
			} else {
				keyEncryptingKey = nil
			}
			destinationRedshiftUpdateUploadingMethodS3StagingEncryptionAESCBCEnvelopeEncryption = &shared.DestinationRedshiftUpdateUploadingMethodS3StagingEncryptionAESCBCEnvelopeEncryption{
				EncryptionType:   encryptionType1,
				KeyEncryptingKey: keyEncryptingKey,
			}
		}
		if destinationRedshiftUpdateUploadingMethodS3StagingEncryptionAESCBCEnvelopeEncryption != nil {
			encryption = &shared.DestinationRedshiftUpdateUploadingMethodS3StagingEncryption{
				DestinationRedshiftUpdateUploadingMethodS3StagingEncryptionAESCBCEnvelopeEncryption: destinationRedshiftUpdateUploadingMethodS3StagingEncryptionAESCBCEnvelopeEncryption,
			}
		}
		fileBufferCount := new(int64)
		if !r.Configuration.UploadingMethod.DestinationRedshiftUploadingMethodS3Staging.FileBufferCount.IsUnknown() && !r.Configuration.UploadingMethod.DestinationRedshiftUploadingMethodS3Staging.FileBufferCount.IsNull() {
			*fileBufferCount = r.Configuration.UploadingMethod.DestinationRedshiftUploadingMethodS3Staging.FileBufferCount.ValueInt64()
		} else {
			fileBufferCount = nil
		}
		fileNamePattern := new(string)
		if !r.Configuration.UploadingMethod.DestinationRedshiftUploadingMethodS3Staging.FileNamePattern.IsUnknown() && !r.Configuration.UploadingMethod.DestinationRedshiftUploadingMethodS3Staging.FileNamePattern.IsNull() {
			*fileNamePattern = r.Configuration.UploadingMethod.DestinationRedshiftUploadingMethodS3Staging.FileNamePattern.ValueString()
		} else {
			fileNamePattern = nil
		}
		method1 := shared.DestinationRedshiftUpdateUploadingMethodS3StagingMethod(r.Configuration.UploadingMethod.DestinationRedshiftUploadingMethodS3Staging.Method.ValueString())
		purgeStagingData := new(bool)
		if !r.Configuration.UploadingMethod.DestinationRedshiftUploadingMethodS3Staging.PurgeStagingData.IsUnknown() && !r.Configuration.UploadingMethod.DestinationRedshiftUploadingMethodS3Staging.PurgeStagingData.IsNull() {
			*purgeStagingData = r.Configuration.UploadingMethod.DestinationRedshiftUploadingMethodS3Staging.PurgeStagingData.ValueBool()
		} else {
			purgeStagingData = nil
		}
		s3BucketName := r.Configuration.UploadingMethod.DestinationRedshiftUploadingMethodS3Staging.S3BucketName.ValueString()
		s3BucketPath := new(string)
		if !r.Configuration.UploadingMethod.DestinationRedshiftUploadingMethodS3Staging.S3BucketPath.IsUnknown() && !r.Configuration.UploadingMethod.DestinationRedshiftUploadingMethodS3Staging.S3BucketPath.IsNull() {
			*s3BucketPath = r.Configuration.UploadingMethod.DestinationRedshiftUploadingMethodS3Staging.S3BucketPath.ValueString()
		} else {
			s3BucketPath = nil
		}
		s3BucketRegion := shared.DestinationRedshiftUpdateUploadingMethodS3StagingS3BucketRegion(r.Configuration.UploadingMethod.DestinationRedshiftUploadingMethodS3Staging.S3BucketRegion.ValueString())
		secretAccessKey := r.Configuration.UploadingMethod.DestinationRedshiftUploadingMethodS3Staging.SecretAccessKey.ValueString()
		destinationRedshiftUpdateUploadingMethodS3Staging = &shared.DestinationRedshiftUpdateUploadingMethodS3Staging{
			AccessKeyID:      accessKeyID,
			Encryption:       encryption,
			FileBufferCount:  fileBufferCount,
			FileNamePattern:  fileNamePattern,
			Method:           method1,
			PurgeStagingData: purgeStagingData,
			S3BucketName:     s3BucketName,
			S3BucketPath:     s3BucketPath,
			S3BucketRegion:   s3BucketRegion,
			SecretAccessKey:  secretAccessKey,
		}
	}
	if destinationRedshiftUpdateUploadingMethodS3Staging != nil {
		uploadingMethod = &shared.DestinationRedshiftUpdateUploadingMethod{
			DestinationRedshiftUpdateUploadingMethodS3Staging: destinationRedshiftUpdateUploadingMethodS3Staging,
		}
	}
	username := r.Configuration.Username.ValueString()
	configuration := shared.DestinationRedshiftUpdate{
		Database:        database,
		Host:            host,
		JdbcURLParams:   jdbcURLParams,
		Password:        password,
		Port:            port,
		Schema:          schema,
		TunnelMethod:    tunnelMethod,
		UploadingMethod: uploadingMethod,
		Username:        username,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.DestinationRedshiftPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *DestinationRedshiftResourceModel) ToDeleteSDKType() *shared.DestinationRedshiftCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *DestinationRedshiftResourceModel) RefreshFromCreateResponse(resp *shared.DestinationResponse) {
	r.DestinationID = types.StringValue(resp.DestinationID)
	r.DestinationType = types.StringValue(resp.DestinationType)
	r.Name = types.StringValue(resp.Name)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}