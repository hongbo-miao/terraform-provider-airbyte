// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type SourceFileSecureUpdateStorageProviderSCPSecureCopyProtocol struct {
	Host     types.String `tfsdk:"host"`
	Password types.String `tfsdk:"password"`
	Port     types.String `tfsdk:"port"`
	Storage  types.String `tfsdk:"storage"`
	User     types.String `tfsdk:"user"`
}