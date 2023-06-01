// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *SourceRecreationResourceModel) ToCreateSDKType() *shared.SourceRecreationCreateRequest {
	apikey := r.Configuration.Apikey.ValueString()
	queryCampsites := new(string)
	if !r.Configuration.QueryCampsites.IsUnknown() && !r.Configuration.QueryCampsites.IsNull() {
		*queryCampsites = r.Configuration.QueryCampsites.ValueString()
	} else {
		queryCampsites = nil
	}
	sourceType := shared.SourceRecreationRecreation(r.Configuration.SourceType.ValueString())
	configuration := shared.SourceRecreation{
		Apikey:         apikey,
		QueryCampsites: queryCampsites,
		SourceType:     sourceType,
	}
	name := r.Name.ValueString()
	secretID := new(string)
	if !r.SecretID.IsUnknown() && !r.SecretID.IsNull() {
		*secretID = r.SecretID.ValueString()
	} else {
		secretID = nil
	}
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceRecreationCreateRequest{
		Configuration: configuration,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceRecreationResourceModel) ToUpdateSDKType() *shared.SourceRecreationPutRequest {
	apikey := r.Configuration.Apikey.ValueString()
	queryCampsites := new(string)
	if !r.Configuration.QueryCampsites.IsUnknown() && !r.Configuration.QueryCampsites.IsNull() {
		*queryCampsites = r.Configuration.QueryCampsites.ValueString()
	} else {
		queryCampsites = nil
	}
	configuration := shared.SourceRecreationUpdate{
		Apikey:         apikey,
		QueryCampsites: queryCampsites,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceRecreationPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceRecreationResourceModel) ToDeleteSDKType() *shared.SourceRecreationCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceRecreationResourceModel) RefreshFromCreateResponse(resp *shared.SourceResponse) {
	r.Name = types.StringValue(resp.Name)
	r.SourceID = types.StringValue(resp.SourceID)
	r.SourceType = types.StringValue(resp.SourceType)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}