// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

type DestinationTeradataSSLModes struct {
	Allow      *Fake                        `tfsdk:"allow"`
	Disable    *Fake                        `tfsdk:"disable"`
	Prefer     *Fake                        `tfsdk:"prefer"`
	Require    *Fake                        `tfsdk:"require"`
	VerifyCa   *DestinationTeradataVerifyCa `tfsdk:"verify_ca"`
	VerifyFull *DestinationTeradataVerifyCa `tfsdk:"verify_full"`
}