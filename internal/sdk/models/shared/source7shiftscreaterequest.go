// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type Source7shiftsCreateRequest struct {
	Configuration Source7shifts `json:"configuration"`
	// The UUID of the connector definition. One of configuration.sourceType or definitionId must be provided.
	DefinitionID *string `json:"definitionId,omitempty"`
	// Name of the source e.g. dev-mysql-instance.
	Name string `json:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretID    *string `json:"secretId,omitempty"`
	WorkspaceID string  `json:"workspaceId"`
}

func (o *Source7shiftsCreateRequest) GetConfiguration() Source7shifts {
	if o == nil {
		return Source7shifts{}
	}
	return o.Configuration
}

func (o *Source7shiftsCreateRequest) GetDefinitionID() *string {
	if o == nil {
		return nil
	}
	return o.DefinitionID
}

func (o *Source7shiftsCreateRequest) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *Source7shiftsCreateRequest) GetSecretID() *string {
	if o == nil {
		return nil
	}
	return o.SecretID
}

func (o *Source7shiftsCreateRequest) GetWorkspaceID() string {
	if o == nil {
		return ""
	}
	return o.WorkspaceID
}