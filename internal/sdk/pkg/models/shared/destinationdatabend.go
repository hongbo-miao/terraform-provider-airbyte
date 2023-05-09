// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type DestinationDatabendDatabendEnum string

const (
	DestinationDatabendDatabendEnumDatabend DestinationDatabendDatabendEnum = "databend"
)

func (e DestinationDatabendDatabendEnum) ToPointer() *DestinationDatabendDatabendEnum {
	return &e
}

func (e *DestinationDatabendDatabendEnum) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "databend":
		*e = DestinationDatabendDatabendEnum(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationDatabendDatabendEnum: %v", v)
	}
}

// DestinationDatabend - The values required to configure the destination.
type DestinationDatabend struct {
	// Name of the database.
	Database        string                          `json:"database"`
	DestinationType DestinationDatabendDatabendEnum `json:"destinationType"`
	// Hostname of the database.
	Host string `json:"host"`
	// Password associated with the username.
	Password *string `json:"password,omitempty"`
	// Port of the database.
	Port *int64 `json:"port,omitempty"`
	// The default  table was written to.
	Table *string `json:"table,omitempty"`
	// Username to use to access the database.
	Username string `json:"username"`
}
