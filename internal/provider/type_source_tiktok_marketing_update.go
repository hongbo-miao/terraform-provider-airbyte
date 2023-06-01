// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type SourceTiktokMarketingUpdate struct {
	AttributionWindow types.Int64                                      `tfsdk:"attribution_window"`
	Credentials       *SourceTiktokMarketingUpdateAuthenticationMethod `tfsdk:"credentials"`
	EndDate           types.String                                     `tfsdk:"end_date"`
	StartDate         types.String                                     `tfsdk:"start_date"`
	SourceType        types.String                                     `tfsdk:"source_type"`
}