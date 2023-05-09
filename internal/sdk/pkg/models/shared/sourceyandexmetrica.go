// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"airbyte/internal/sdk/pkg/types"
	"encoding/json"
	"fmt"
)

type SourceYandexMetricaYandexMetricaEnum string

const (
	SourceYandexMetricaYandexMetricaEnumYandexMetrica SourceYandexMetricaYandexMetricaEnum = "yandex-metrica"
)

func (e SourceYandexMetricaYandexMetricaEnum) ToPointer() *SourceYandexMetricaYandexMetricaEnum {
	return &e
}

func (e *SourceYandexMetricaYandexMetricaEnum) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "yandex-metrica":
		*e = SourceYandexMetricaYandexMetricaEnum(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceYandexMetricaYandexMetricaEnum: %v", v)
	}
}

// SourceYandexMetrica - The values required to configure the source.
type SourceYandexMetrica struct {
	// Your Yandex Metrica API access token
	AuthToken string `json:"auth_token"`
	// Counter ID
	CounterID string `json:"counter_id"`
	// Starting point for your data replication, in format of "YYYY-MM-DD". If not provided will sync till most recent date.
	EndDate    *types.Date                          `json:"end_date,omitempty"`
	SourceType SourceYandexMetricaYandexMetricaEnum `json:"sourceType"`
	// Starting point for your data replication, in format of "YYYY-MM-DD".
	StartDate types.Date `json:"start_date"`
}
