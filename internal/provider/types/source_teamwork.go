// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type SourceTeamwork struct {
	Password  types.String `tfsdk:"password"`
	SiteName  types.String `tfsdk:"site_name"`
	StartDate types.String `tfsdk:"start_date"`
	Username  types.String `tfsdk:"username"`
}