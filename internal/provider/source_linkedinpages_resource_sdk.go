// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"time"
)

func (r *SourceLinkedinPagesResourceModel) ToSharedSourceLinkedinPagesCreateRequest() *shared.SourceLinkedinPagesCreateRequest {
	var credentials *shared.SourceLinkedinPagesAuthentication
	if r.Configuration.Credentials != nil {
		var sourceLinkedinPagesOAuth20 *shared.SourceLinkedinPagesOAuth20
		if r.Configuration.Credentials.OAuth20 != nil {
			clientID := r.Configuration.Credentials.OAuth20.ClientID.ValueString()
			clientSecret := r.Configuration.Credentials.OAuth20.ClientSecret.ValueString()
			refreshToken := r.Configuration.Credentials.OAuth20.RefreshToken.ValueString()
			sourceLinkedinPagesOAuth20 = &shared.SourceLinkedinPagesOAuth20{
				ClientID:     clientID,
				ClientSecret: clientSecret,
				RefreshToken: refreshToken,
			}
		}
		if sourceLinkedinPagesOAuth20 != nil {
			credentials = &shared.SourceLinkedinPagesAuthentication{
				SourceLinkedinPagesOAuth20: sourceLinkedinPagesOAuth20,
			}
		}
		var sourceLinkedinPagesAccessToken *shared.SourceLinkedinPagesAccessToken
		if r.Configuration.Credentials.AccessToken != nil {
			accessToken := r.Configuration.Credentials.AccessToken.AccessToken.ValueString()
			sourceLinkedinPagesAccessToken = &shared.SourceLinkedinPagesAccessToken{
				AccessToken: accessToken,
			}
		}
		if sourceLinkedinPagesAccessToken != nil {
			credentials = &shared.SourceLinkedinPagesAuthentication{
				SourceLinkedinPagesAccessToken: sourceLinkedinPagesAccessToken,
			}
		}
	}
	orgID := r.Configuration.OrgID.ValueString()
	startDate := new(time.Time)
	if !r.Configuration.StartDate.IsUnknown() && !r.Configuration.StartDate.IsNull() {
		*startDate, _ = time.Parse(time.RFC3339Nano, r.Configuration.StartDate.ValueString())
	} else {
		startDate = nil
	}
	timeGranularityType := new(shared.SourceLinkedinPagesTimeGranularityType)
	if !r.Configuration.TimeGranularityType.IsUnknown() && !r.Configuration.TimeGranularityType.IsNull() {
		*timeGranularityType = shared.SourceLinkedinPagesTimeGranularityType(r.Configuration.TimeGranularityType.ValueString())
	} else {
		timeGranularityType = nil
	}
	configuration := shared.SourceLinkedinPages{
		Credentials:         credentials,
		OrgID:               orgID,
		StartDate:           startDate,
		TimeGranularityType: timeGranularityType,
	}
	definitionID := new(string)
	if !r.DefinitionID.IsUnknown() && !r.DefinitionID.IsNull() {
		*definitionID = r.DefinitionID.ValueString()
	} else {
		definitionID = nil
	}
	name := r.Name.ValueString()
	secretID := new(string)
	if !r.SecretID.IsUnknown() && !r.SecretID.IsNull() {
		*secretID = r.SecretID.ValueString()
	} else {
		secretID = nil
	}
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceLinkedinPagesCreateRequest{
		Configuration: configuration,
		DefinitionID:  definitionID,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceLinkedinPagesResourceModel) RefreshFromSharedSourceResponse(resp *shared.SourceResponse) {
	if resp != nil {
		r.Name = types.StringValue(resp.Name)
		r.SourceID = types.StringValue(resp.SourceID)
		r.SourceType = types.StringValue(resp.SourceType)
		r.WorkspaceID = types.StringValue(resp.WorkspaceID)
	}
}

func (r *SourceLinkedinPagesResourceModel) ToSharedSourceLinkedinPagesPutRequest() *shared.SourceLinkedinPagesPutRequest {
	var credentials *shared.SourceLinkedinPagesUpdateAuthentication
	if r.Configuration.Credentials != nil {
		var sourceLinkedinPagesUpdateOAuth20 *shared.SourceLinkedinPagesUpdateOAuth20
		if r.Configuration.Credentials.OAuth20 != nil {
			clientID := r.Configuration.Credentials.OAuth20.ClientID.ValueString()
			clientSecret := r.Configuration.Credentials.OAuth20.ClientSecret.ValueString()
			refreshToken := r.Configuration.Credentials.OAuth20.RefreshToken.ValueString()
			sourceLinkedinPagesUpdateOAuth20 = &shared.SourceLinkedinPagesUpdateOAuth20{
				ClientID:     clientID,
				ClientSecret: clientSecret,
				RefreshToken: refreshToken,
			}
		}
		if sourceLinkedinPagesUpdateOAuth20 != nil {
			credentials = &shared.SourceLinkedinPagesUpdateAuthentication{
				SourceLinkedinPagesUpdateOAuth20: sourceLinkedinPagesUpdateOAuth20,
			}
		}
		var sourceLinkedinPagesUpdateAccessToken *shared.SourceLinkedinPagesUpdateAccessToken
		if r.Configuration.Credentials.AccessToken != nil {
			accessToken := r.Configuration.Credentials.AccessToken.AccessToken.ValueString()
			sourceLinkedinPagesUpdateAccessToken = &shared.SourceLinkedinPagesUpdateAccessToken{
				AccessToken: accessToken,
			}
		}
		if sourceLinkedinPagesUpdateAccessToken != nil {
			credentials = &shared.SourceLinkedinPagesUpdateAuthentication{
				SourceLinkedinPagesUpdateAccessToken: sourceLinkedinPagesUpdateAccessToken,
			}
		}
	}
	orgID := r.Configuration.OrgID.ValueString()
	startDate := new(time.Time)
	if !r.Configuration.StartDate.IsUnknown() && !r.Configuration.StartDate.IsNull() {
		*startDate, _ = time.Parse(time.RFC3339Nano, r.Configuration.StartDate.ValueString())
	} else {
		startDate = nil
	}
	timeGranularityType := new(shared.TimeGranularityType)
	if !r.Configuration.TimeGranularityType.IsUnknown() && !r.Configuration.TimeGranularityType.IsNull() {
		*timeGranularityType = shared.TimeGranularityType(r.Configuration.TimeGranularityType.ValueString())
	} else {
		timeGranularityType = nil
	}
	configuration := shared.SourceLinkedinPagesUpdate{
		Credentials:         credentials,
		OrgID:               orgID,
		StartDate:           startDate,
		TimeGranularityType: timeGranularityType,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceLinkedinPagesPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}
