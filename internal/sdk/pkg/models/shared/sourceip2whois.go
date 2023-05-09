// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type SourceIp2whoisIp2whoisEnum string

const (
	SourceIp2whoisIp2whoisEnumIp2whois SourceIp2whoisIp2whoisEnum = "ip2whois"
)

func (e SourceIp2whoisIp2whoisEnum) ToPointer() *SourceIp2whoisIp2whoisEnum {
	return &e
}

func (e *SourceIp2whoisIp2whoisEnum) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "ip2whois":
		*e = SourceIp2whoisIp2whoisEnum(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceIp2whoisIp2whoisEnum: %v", v)
	}
}

// SourceIp2whois - The values required to configure the source.
type SourceIp2whois struct {
	// Your API Key. See <a href="https://www.ip2whois.com/developers-api">here</a>.
	APIKey *string `json:"api_key,omitempty"`
	// Domain name. See <a href="https://www.ip2whois.com/developers-api">here</a>.
	Domain     *string                    `json:"domain,omitempty"`
	SourceType SourceIp2whoisIp2whoisEnum `json:"sourceType"`
}
