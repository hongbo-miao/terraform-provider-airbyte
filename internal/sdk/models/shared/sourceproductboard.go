// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/internal/utils"
	"time"
)

type Productboard string

const (
	ProductboardProductboard Productboard = "productboard"
)

func (e Productboard) ToPointer() *Productboard {
	return &e
}
func (e *Productboard) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "productboard":
		*e = Productboard(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Productboard: %v", v)
	}
}

type SourceProductboard struct {
	// Your Productboard access token. See https://developer.productboard.com/reference/authentication for steps to generate one.
	AccessToken string       `json:"access_token"`
	sourceType  Productboard `const:"productboard" json:"sourceType"`
	StartDate   time.Time    `json:"start_date"`
}

func (s SourceProductboard) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceProductboard) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SourceProductboard) GetAccessToken() string {
	if o == nil {
		return ""
	}
	return o.AccessToken
}

func (o *SourceProductboard) GetSourceType() Productboard {
	return ProductboardProductboard
}

func (o *SourceProductboard) GetStartDate() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.StartDate
}