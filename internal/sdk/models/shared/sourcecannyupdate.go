// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type SourceCannyUpdate struct {
	// You can find your secret API key in Your Canny Subdomain > Settings > API
	APIKey string `json:"api_key"`
}

func (o *SourceCannyUpdate) GetAPIKey() string {
	if o == nil {
		return ""
	}
	return o.APIKey
}