// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type SourcePaypalTransactionPaypalTransactionEnum string

const (
	SourcePaypalTransactionPaypalTransactionEnumPaypalTransaction SourcePaypalTransactionPaypalTransactionEnum = "paypal-transaction"
)

func (e SourcePaypalTransactionPaypalTransactionEnum) ToPointer() *SourcePaypalTransactionPaypalTransactionEnum {
	return &e
}

func (e *SourcePaypalTransactionPaypalTransactionEnum) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "paypal-transaction":
		*e = SourcePaypalTransactionPaypalTransactionEnum(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourcePaypalTransactionPaypalTransactionEnum: %v", v)
	}
}

// SourcePaypalTransaction - The values required to configure the source.
type SourcePaypalTransaction struct {
	// The Client ID of your Paypal developer application.
	ClientID *string `json:"client_id,omitempty"`
	// The Client Secret of your Paypal developer application.
	ClientSecret *string `json:"client_secret,omitempty"`
	// Determines whether to use the sandbox or production environment.
	IsSandbox bool `json:"is_sandbox"`
	// The key to refresh the expired access token.
	RefreshToken *string                                      `json:"refresh_token,omitempty"`
	SourceType   SourcePaypalTransactionPaypalTransactionEnum `json:"sourceType"`
	// Start Date for data extraction in <a href="https://datatracker.ietf.org/doc/html/rfc3339#section-5.6">ISO format</a>. Date must be in range from 3 years till 12 hrs before present time.
	StartDate string `json:"start_date"`
}
