// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type SourceAzureBlobStoragePutRequest struct {
	// NOTE: When this Spec is changed, legacy_config_transformer.py must also be modified to uptake the changes
	// because it is responsible for converting legacy Azure Blob Storage v0 configs into v1 configs using the File-Based CDK.
	Configuration SourceAzureBlobStorageUpdate `json:"configuration"`
	Name          string                       `json:"name"`
	WorkspaceID   string                       `json:"workspaceId"`
}

func (o *SourceAzureBlobStoragePutRequest) GetConfiguration() SourceAzureBlobStorageUpdate {
	if o == nil {
		return SourceAzureBlobStorageUpdate{}
	}
	return o.Configuration
}

func (o *SourceAzureBlobStoragePutRequest) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *SourceAzureBlobStoragePutRequest) GetWorkspaceID() string {
	if o == nil {
		return ""
	}
	return o.WorkspaceID
}