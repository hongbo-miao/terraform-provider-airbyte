// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type DestinationSnowflakeUpdate struct {
	Credentials     *DestinationSnowflakeUpdateAuthorizationMethod `tfsdk:"credentials"`
	Database        types.String                                   `tfsdk:"database"`
	FileBufferCount types.Int64                                    `tfsdk:"file_buffer_count"`
	Host            types.String                                   `tfsdk:"host"`
	JdbcURLParams   types.String                                   `tfsdk:"jdbc_url_params"`
	LoadingMethod   *DestinationSnowflakeUpdateDataStagingMethod   `tfsdk:"loading_method"`
	Role            types.String                                   `tfsdk:"role"`
	Schema          types.String                                   `tfsdk:"schema"`
	Username        types.String                                   `tfsdk:"username"`
	Warehouse       types.String                                   `tfsdk:"warehouse"`
	DestinationType types.String                                   `tfsdk:"destination_type"`
}