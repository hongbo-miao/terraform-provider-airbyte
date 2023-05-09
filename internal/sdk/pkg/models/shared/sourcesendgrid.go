// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"time"
)

type SourceSendgridSendgridEnum string

const (
	SourceSendgridSendgridEnumSendgrid SourceSendgridSendgridEnum = "sendgrid"
)

func (e SourceSendgridSendgridEnum) ToPointer() *SourceSendgridSendgridEnum {
	return &e
}

func (e *SourceSendgridSendgridEnum) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "sendgrid":
		*e = SourceSendgridSendgridEnum(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceSendgridSendgridEnum: %v", v)
	}
}

// SourceSendgrid - The values required to configure the source.
type SourceSendgrid struct {
	// API Key, use <a href="https://app.sendgrid.com/settings/api_keys/">admin</a> to generate this key.
	Apikey     string                     `json:"apikey"`
	SourceType SourceSendgridSendgridEnum `json:"sourceType"`
	// Start time in ISO8601 format. Any data before this time point will not be replicated.
	StartTime *time.Time `json:"start_time,omitempty"`
}
