// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type SourcePypiUpdate struct {
	ProjectName types.String `tfsdk:"project_name"`
	Version     types.String `tfsdk:"version"`
	SourceType  types.String `tfsdk:"source_type"`
}