// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type SourceWhenIWorkUpdate struct {
	// Email of your when-i-work account
	Email string `json:"email"`
	// Password for your when-i-work account
	Password string `json:"password"`
}

func (o *SourceWhenIWorkUpdate) GetEmail() string {
	if o == nil {
		return ""
	}
	return o.Email
}

func (o *SourceWhenIWorkUpdate) GetPassword() string {
	if o == nil {
		return ""
	}
	return o.Password
}