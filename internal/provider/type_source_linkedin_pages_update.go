// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type SourceLinkedinPagesUpdate struct {
	Credentials *SourceLinkedinPagesUpdateAuthentication `tfsdk:"credentials"`
	OrgID       types.String                             `tfsdk:"org_id"`
	SourceType  types.String                             `tfsdk:"source_type"`
}