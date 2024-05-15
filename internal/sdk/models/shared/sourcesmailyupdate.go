// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type SourceSmailyUpdate struct {
	// API user password. See https://smaily.com/help/api/general/create-api-user/
	APIPassword string `json:"api_password"`
	// API Subdomain. See https://smaily.com/help/api/general/create-api-user/
	APISubdomain string `json:"api_subdomain"`
	// API user username. See https://smaily.com/help/api/general/create-api-user/
	APIUsername string `json:"api_username"`
}

func (o *SourceSmailyUpdate) GetAPIPassword() string {
	if o == nil {
		return ""
	}
	return o.APIPassword
}

func (o *SourceSmailyUpdate) GetAPISubdomain() string {
	if o == nil {
		return ""
	}
	return o.APISubdomain
}

func (o *SourceSmailyUpdate) GetAPIUsername() string {
	if o == nil {
		return ""
	}
	return o.APIUsername
}