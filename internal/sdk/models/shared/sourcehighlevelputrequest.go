// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type SourceHighLevelPutRequest struct {
	Configuration SourceHighLevelUpdate `json:"configuration"`
	Name          string                `json:"name"`
	WorkspaceID   string                `json:"workspaceId"`
}

func (o *SourceHighLevelPutRequest) GetConfiguration() SourceHighLevelUpdate {
	if o == nil {
		return SourceHighLevelUpdate{}
	}
	return o.Configuration
}

func (o *SourceHighLevelPutRequest) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *SourceHighLevelPutRequest) GetWorkspaceID() string {
	if o == nil {
		return ""
	}
	return o.WorkspaceID
}