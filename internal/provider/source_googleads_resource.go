// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	speakeasy_objectplanmodifier "github.com/airbytehq/terraform-provider-airbyte/internal/planmodifiers/objectplanmodifier"
	speakeasy_stringplanmodifier "github.com/airbytehq/terraform-provider-airbyte/internal/planmodifiers/stringplanmodifier"
	tfTypes "github.com/airbytehq/terraform-provider-airbyte/internal/provider/types"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/models/operations"
	"github.com/airbytehq/terraform-provider-airbyte/internal/validators"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &SourceGoogleAdsResource{}
var _ resource.ResourceWithImportState = &SourceGoogleAdsResource{}

func NewSourceGoogleAdsResource() resource.Resource {
	return &SourceGoogleAdsResource{}
}

// SourceGoogleAdsResource defines the resource implementation.
type SourceGoogleAdsResource struct {
	client *sdk.SDK
}

// SourceGoogleAdsResourceModel describes the resource data model.
type SourceGoogleAdsResourceModel struct {
	Configuration tfTypes.SourceGoogleAds `tfsdk:"configuration"`
	DefinitionID  types.String            `tfsdk:"definition_id"`
	Name          types.String            `tfsdk:"name"`
	SecretID      types.String            `tfsdk:"secret_id"`
	SourceID      types.String            `tfsdk:"source_id"`
	SourceType    types.String            `tfsdk:"source_type"`
	WorkspaceID   types.String            `tfsdk:"workspace_id"`
}

func (r *SourceGoogleAdsResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_source_google_ads"
}

func (r *SourceGoogleAdsResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "SourceGoogleAds Resource",
		Attributes: map[string]schema.Attribute{
			"configuration": schema.SingleNestedAttribute{
				PlanModifiers: []planmodifier.Object{
					speakeasy_objectplanmodifier.SuppressDiff(speakeasy_objectplanmodifier.ExplicitSuppress),
				},
				Required: true,
				Attributes: map[string]schema.Attribute{
					"conversion_window_days": schema.Int64Attribute{
						Computed:    true,
						Optional:    true,
						Default:     int64default.StaticInt64(14),
						Description: `A conversion window is the number of days after an ad interaction (such as an ad click or video view) during which a conversion, such as a purchase, is recorded in Google Ads. For more information, see <a href="https://support.google.com/google-ads/answer/3123169?hl=en">Google's documentation</a>. Default: 14`,
					},
					"credentials": schema.SingleNestedAttribute{
						Required: true,
						Attributes: map[string]schema.Attribute{
							"access_token": schema.StringAttribute{
								Optional:    true,
								Sensitive:   true,
								Description: `The Access Token for making authenticated requests. For detailed instructions on finding this value, refer to our <a href="https://docs.airbyte.com/integrations/sources/google-ads#setup-guide">documentation</a>.`,
							},
							"client_id": schema.StringAttribute{
								Required:    true,
								Description: `The Client ID of your Google Ads developer application. For detailed instructions on finding this value, refer to our <a href="https://docs.airbyte.com/integrations/sources/google-ads#setup-guide">documentation</a>.`,
							},
							"client_secret": schema.StringAttribute{
								Required:    true,
								Description: `The Client Secret of your Google Ads developer application. For detailed instructions on finding this value, refer to our <a href="https://docs.airbyte.com/integrations/sources/google-ads#setup-guide">documentation</a>.`,
							},
							"developer_token": schema.StringAttribute{
								Required:    true,
								Sensitive:   true,
								Description: `The Developer Token granted by Google to use their APIs. For detailed instructions on finding this value, refer to our <a href="https://docs.airbyte.com/integrations/sources/google-ads#setup-guide">documentation</a>.`,
							},
							"refresh_token": schema.StringAttribute{
								Required:    true,
								Sensitive:   true,
								Description: `The token used to obtain a new Access Token. For detailed instructions on finding this value, refer to our <a href="https://docs.airbyte.com/integrations/sources/google-ads#setup-guide">documentation</a>.`,
							},
						},
					},
					"custom_queries_array": schema.ListNestedAttribute{
						Optional: true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"query": schema.StringAttribute{
									Required:    true,
									Description: `A custom defined GAQL query for building the report. Avoid including the segments.date field; wherever possible, Airbyte will automatically include it for incremental syncs. For more information, refer to <a href="https://developers.google.com/google-ads/api/fields/v11/overview_query_builder">Google's documentation</a>.`,
								},
								"table_name": schema.StringAttribute{
									Required:    true,
									Description: `The table name in your destination database for the chosen query.`,
								},
							},
						},
					},
					"customer_id": schema.StringAttribute{
						Optional:    true,
						Description: `Comma-separated list of (client) customer IDs. Each customer ID must be specified as a 10-digit number without dashes. For detailed instructions on finding this value, refer to our <a href="https://docs.airbyte.com/integrations/sources/google-ads#setup-guide">documentation</a>.`,
					},
					"customer_status_filter": schema.ListAttribute{
						Optional:    true,
						ElementType: types.StringType,
						Description: `A list of customer statuses to filter on. For detailed info about what each status mean refer to Google Ads <a href="https://developers.google.com/google-ads/api/reference/rpc/v15/CustomerStatusEnum.CustomerStatus">documentation</a>.`,
					},
					"end_date": schema.StringAttribute{
						Optional:    true,
						Description: `UTC date in the format YYYY-MM-DD. Any data after this date will not be replicated. (Default value of today is used if not set)`,
						Validators: []validator.String{
							validators.IsValidDate(),
						},
					},
					"start_date": schema.StringAttribute{
						Optional:    true,
						Description: `UTC date in the format YYYY-MM-DD. Any data before this date will not be replicated. (Default value of two years ago is used if not set)`,
						Validators: []validator.String{
							validators.IsValidDate(),
						},
					},
				},
			},
			"definition_id": schema.StringAttribute{
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplaceIfConfigured(),
				},
				Optional:    true,
				Description: `The UUID of the connector definition. One of configuration.sourceType or definitionId must be provided. Requires replacement if changed. `,
			},
			"name": schema.StringAttribute{
				PlanModifiers: []planmodifier.String{
					speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
				},
				Required:    true,
				Description: `Name of the source e.g. dev-mysql-instance.`,
			},
			"secret_id": schema.StringAttribute{
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplaceIfConfigured(),
				},
				Optional:    true,
				Description: `Optional secretID obtained through the public API OAuth redirect flow. Requires replacement if changed. `,
			},
			"source_id": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
				},
			},
			"source_type": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
				},
			},
			"workspace_id": schema.StringAttribute{
				PlanModifiers: []planmodifier.String{
					speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
				},
				Required: true,
			},
		},
	}
}

func (r *SourceGoogleAdsResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*sdk.SDK)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *sdk.SDK, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *SourceGoogleAdsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *SourceGoogleAdsResourceModel
	var plan types.Object

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(plan.As(ctx, &data, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	if resp.Diagnostics.HasError() {
		return
	}

	request := data.ToSharedSourceGoogleAdsCreateRequest()
	res, err := r.client.Sources.CreateSourceGoogleAds(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		if res != nil && res.RawResponse != nil {
			resp.Diagnostics.AddError("unexpected http request/response", debugResponse(res.RawResponse))
		}
		return
	}
	if res == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res))
		return
	}
	if res.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}
	if !(res.SourceResponse != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedSourceResponse(res.SourceResponse)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)
	sourceID := data.SourceID.ValueString()
	request1 := operations.GetSourceGoogleAdsRequest{
		SourceID: sourceID,
	}
	res1, err := r.client.Sources.GetSourceGoogleAds(ctx, request1)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		if res1 != nil && res1.RawResponse != nil {
			resp.Diagnostics.AddError("unexpected http request/response", debugResponse(res1.RawResponse))
		}
		return
	}
	if res1 == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res1))
		return
	}
	if res1.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res1.StatusCode), debugResponse(res1.RawResponse))
		return
	}
	if !(res1.SourceResponse != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res1.RawResponse))
		return
	}
	data.RefreshFromSharedSourceResponse(res1.SourceResponse)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SourceGoogleAdsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *SourceGoogleAdsResourceModel
	var item types.Object

	resp.Diagnostics.Append(req.State.Get(ctx, &item)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(item.As(ctx, &data, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	if resp.Diagnostics.HasError() {
		return
	}

	sourceID := data.SourceID.ValueString()
	request := operations.GetSourceGoogleAdsRequest{
		SourceID: sourceID,
	}
	res, err := r.client.Sources.GetSourceGoogleAds(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		if res != nil && res.RawResponse != nil {
			resp.Diagnostics.AddError("unexpected http request/response", debugResponse(res.RawResponse))
		}
		return
	}
	if res == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res))
		return
	}
	if res.StatusCode == 404 {
		resp.State.RemoveResource(ctx)
		return
	}
	if res.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}
	if !(res.SourceResponse != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedSourceResponse(res.SourceResponse)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SourceGoogleAdsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *SourceGoogleAdsResourceModel
	var plan types.Object

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	merge(ctx, req, resp, &data)
	if resp.Diagnostics.HasError() {
		return
	}

	sourceGoogleAdsPutRequest := data.ToSharedSourceGoogleAdsPutRequest()
	sourceID := data.SourceID.ValueString()
	request := operations.PutSourceGoogleAdsRequest{
		SourceGoogleAdsPutRequest: sourceGoogleAdsPutRequest,
		SourceID:                  sourceID,
	}
	res, err := r.client.Sources.PutSourceGoogleAds(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		if res != nil && res.RawResponse != nil {
			resp.Diagnostics.AddError("unexpected http request/response", debugResponse(res.RawResponse))
		}
		return
	}
	if res == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res))
		return
	}
	if fmt.Sprintf("%v", res.StatusCode)[0] != '2' {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}
	refreshPlan(ctx, plan, &data, resp.Diagnostics)
	sourceId1 := data.SourceID.ValueString()
	request1 := operations.GetSourceGoogleAdsRequest{
		SourceID: sourceId1,
	}
	res1, err := r.client.Sources.GetSourceGoogleAds(ctx, request1)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		if res1 != nil && res1.RawResponse != nil {
			resp.Diagnostics.AddError("unexpected http request/response", debugResponse(res1.RawResponse))
		}
		return
	}
	if res1 == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res1))
		return
	}
	if res1.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res1.StatusCode), debugResponse(res1.RawResponse))
		return
	}
	if !(res1.SourceResponse != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res1.RawResponse))
		return
	}
	data.RefreshFromSharedSourceResponse(res1.SourceResponse)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SourceGoogleAdsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *SourceGoogleAdsResourceModel
	var item types.Object

	resp.Diagnostics.Append(req.State.Get(ctx, &item)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(item.As(ctx, &data, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	if resp.Diagnostics.HasError() {
		return
	}

	sourceID := data.SourceID.ValueString()
	request := operations.DeleteSourceGoogleAdsRequest{
		SourceID: sourceID,
	}
	res, err := r.client.Sources.DeleteSourceGoogleAds(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		if res != nil && res.RawResponse != nil {
			resp.Diagnostics.AddError("unexpected http request/response", debugResponse(res.RawResponse))
		}
		return
	}
	if res == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res))
		return
	}
	if fmt.Sprintf("%v", res.StatusCode)[0] != '2' {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}

}

func (r *SourceGoogleAdsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("source_id"), req.ID)...)
}
