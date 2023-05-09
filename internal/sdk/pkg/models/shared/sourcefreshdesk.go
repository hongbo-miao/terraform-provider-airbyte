// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"time"
)

type SourceFreshdeskFreshdeskEnum string

const (
	SourceFreshdeskFreshdeskEnumFreshdesk SourceFreshdeskFreshdeskEnum = "freshdesk"
)

func (e SourceFreshdeskFreshdeskEnum) ToPointer() *SourceFreshdeskFreshdeskEnum {
	return &e
}

func (e *SourceFreshdeskFreshdeskEnum) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "freshdesk":
		*e = SourceFreshdeskFreshdeskEnum(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceFreshdeskFreshdeskEnum: %v", v)
	}
}

// SourceFreshdesk - The values required to configure the source.
type SourceFreshdesk struct {
	// Freshdesk API Key. See the <a href="https://docs.airbyte.com/integrations/sources/freshdesk">docs</a> for more information on how to obtain this key.
	APIKey string `json:"api_key"`
	// Freshdesk domain
	Domain string `json:"domain"`
	// The number of requests per minute that this source allowed to use. There is a rate limit of 50 requests per minute per app per account.
	RequestsPerMinute *int64                       `json:"requests_per_minute,omitempty"`
	SourceType        SourceFreshdeskFreshdeskEnum `json:"sourceType"`
	// UTC date and time. Any data created after this date will be replicated. If this parameter is not set, all data will be replicated.
	StartDate *time.Time `json:"start_date,omitempty"`
}
