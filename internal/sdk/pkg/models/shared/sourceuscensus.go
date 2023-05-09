// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type SourceUsCensusUsCensusEnum string

const (
	SourceUsCensusUsCensusEnumUsCensus SourceUsCensusUsCensusEnum = "us-census"
)

func (e SourceUsCensusUsCensusEnum) ToPointer() *SourceUsCensusUsCensusEnum {
	return &e
}

func (e *SourceUsCensusUsCensusEnum) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "us-census":
		*e = SourceUsCensusUsCensusEnum(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceUsCensusUsCensusEnum: %v", v)
	}
}

// SourceUsCensus - The values required to configure the source.
type SourceUsCensus struct {
	// Your API Key. Get your key <a href="https://api.census.gov/data/key_signup.html">here</a>.
	APIKey string `json:"api_key"`
	// The query parameters portion of the GET request, without the api key
	QueryParams *string `json:"query_params,omitempty"`
	// The path portion of the GET request
	QueryPath  string                     `json:"query_path"`
	SourceType SourceUsCensusUsCensusEnum `json:"sourceType"`
}
