// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type SourceYouniumUpdate struct {
	LegalEntity types.String `tfsdk:"legal_entity"`
	Password    types.String `tfsdk:"password"`
	Playground  types.Bool   `tfsdk:"playground"`
	Username    types.String `tfsdk:"username"`
	SourceType  types.String `tfsdk:"source_type"`
}