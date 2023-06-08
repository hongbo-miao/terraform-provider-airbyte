// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type SourceFileSecure struct {
	DatasetName   types.String                    `tfsdk:"dataset_name"`
	Format        types.String                    `tfsdk:"format"`
	Provider      SourceFileSecureStorageProvider `tfsdk:"provider"`
	ReaderOptions types.String                    `tfsdk:"reader_options"`
	SourceType    types.String                    `tfsdk:"source_type"`
	URL           types.String                    `tfsdk:"url"`
}