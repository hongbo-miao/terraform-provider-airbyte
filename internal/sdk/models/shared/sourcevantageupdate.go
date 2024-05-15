// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type SourceVantageUpdate struct {
	// Your API Access token. See <a href="https://vantage.readme.io/reference/authentication">here</a>.
	AccessToken string `json:"access_token"`
}

func (o *SourceVantageUpdate) GetAccessToken() string {
	if o == nil {
		return ""
	}
	return o.AccessToken
}