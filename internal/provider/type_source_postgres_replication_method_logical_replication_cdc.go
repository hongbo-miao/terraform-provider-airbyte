// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type SourcePostgresReplicationMethodLogicalReplicationCDC struct {
	InitialWaitingSeconds types.Int64             `tfsdk:"initial_waiting_seconds"`
	LsnCommitBehaviour    types.String            `tfsdk:"lsn_commit_behaviour"`
	Method                types.String            `tfsdk:"method"`
	Plugin                types.String            `tfsdk:"plugin"`
	Publication           types.String            `tfsdk:"publication"`
	ReplicationSlot       types.String            `tfsdk:"replication_slot"`
	AdditionalProperties  map[string]types.String `tfsdk:"additional_properties"`
}
