// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"time"
)

type SourceRechargeRechargeEnum string

const (
	SourceRechargeRechargeEnumRecharge SourceRechargeRechargeEnum = "recharge"
)

func (e SourceRechargeRechargeEnum) ToPointer() *SourceRechargeRechargeEnum {
	return &e
}

func (e *SourceRechargeRechargeEnum) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "recharge":
		*e = SourceRechargeRechargeEnum(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceRechargeRechargeEnum: %v", v)
	}
}

// SourceRecharge - The values required to configure the source.
type SourceRecharge struct {
	// The value of the Access Token generated. See the <a href="https://docs.airbyte.com/integrations/sources/recharge">docs</a> for more information.
	AccessToken string                     `json:"access_token"`
	SourceType  SourceRechargeRechargeEnum `json:"sourceType"`
	// The date from which you'd like to replicate data for Recharge API, in the format YYYY-MM-DDT00:00:00Z. Any data before this date will not be replicated.
	StartDate time.Time `json:"start_date"`
}
