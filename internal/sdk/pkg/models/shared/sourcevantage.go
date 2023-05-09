// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type SourceVantageVantageEnum string

const (
	SourceVantageVantageEnumVantage SourceVantageVantageEnum = "vantage"
)

func (e SourceVantageVantageEnum) ToPointer() *SourceVantageVantageEnum {
	return &e
}

func (e *SourceVantageVantageEnum) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "vantage":
		*e = SourceVantageVantageEnum(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceVantageVantageEnum: %v", v)
	}
}

// SourceVantage - The values required to configure the source.
type SourceVantage struct {
	// Your API Access token. See <a href="https://vantage.readme.io/reference/authentication">here</a>.
	AccessToken string                   `json:"access_token"`
	SourceType  SourceVantageVantageEnum `json:"sourceType"`
}
