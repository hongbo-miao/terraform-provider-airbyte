// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type SourceZendeskTalkAuthenticationAPIToken struct {
	APIToken             types.String            `tfsdk:"api_token"`
	AuthType             types.String            `tfsdk:"auth_type"`
	Email                types.String            `tfsdk:"email"`
	AdditionalProperties map[string]types.String `tfsdk:"additional_properties"`
}
