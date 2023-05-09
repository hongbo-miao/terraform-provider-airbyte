// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"time"
)

type SourceTwilioTwilioEnum string

const (
	SourceTwilioTwilioEnumTwilio SourceTwilioTwilioEnum = "twilio"
)

func (e SourceTwilioTwilioEnum) ToPointer() *SourceTwilioTwilioEnum {
	return &e
}

func (e *SourceTwilioTwilioEnum) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "twilio":
		*e = SourceTwilioTwilioEnum(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceTwilioTwilioEnum: %v", v)
	}
}

// SourceTwilio - The values required to configure the source.
type SourceTwilio struct {
	// Twilio account SID
	AccountSid string `json:"account_sid"`
	// Twilio Auth Token.
	AuthToken string `json:"auth_token"`
	// How far into the past to look for records. (in minutes)
	LookbackWindow *int64                 `json:"lookback_window,omitempty"`
	SourceType     SourceTwilioTwilioEnum `json:"sourceType"`
	// UTC date and time in the format 2020-10-01T00:00:00Z. Any data before this date will not be replicated.
	StartDate time.Time `json:"start_date"`
}
