// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type SourceSmartengageSmartengageEnum string

const (
	SourceSmartengageSmartengageEnumSmartengage SourceSmartengageSmartengageEnum = "smartengage"
)

func (e SourceSmartengageSmartengageEnum) ToPointer() *SourceSmartengageSmartengageEnum {
	return &e
}

func (e *SourceSmartengageSmartengageEnum) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "smartengage":
		*e = SourceSmartengageSmartengageEnum(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceSmartengageSmartengageEnum: %v", v)
	}
}

// SourceSmartengage - The values required to configure the source.
type SourceSmartengage struct {
	// API Key
	APIKey     string                           `json:"api_key"`
	SourceType SourceSmartengageSmartengageEnum `json:"sourceType"`
}
