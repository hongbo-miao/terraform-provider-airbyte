// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type DestinationCassandraCassandraEnum string

const (
	DestinationCassandraCassandraEnumCassandra DestinationCassandraCassandraEnum = "cassandra"
)

func (e DestinationCassandraCassandraEnum) ToPointer() *DestinationCassandraCassandraEnum {
	return &e
}

func (e *DestinationCassandraCassandraEnum) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "cassandra":
		*e = DestinationCassandraCassandraEnum(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationCassandraCassandraEnum: %v", v)
	}
}

// DestinationCassandra - The values required to configure the destination.
type DestinationCassandra struct {
	// Address to connect to.
	Address string `json:"address"`
	// Datacenter of the cassandra cluster.
	Datacenter      *string                           `json:"datacenter,omitempty"`
	DestinationType DestinationCassandraCassandraEnum `json:"destinationType"`
	// Default Cassandra keyspace to create data in.
	Keyspace string `json:"keyspace"`
	// Password associated with Cassandra.
	Password string `json:"password"`
	// Port of Cassandra.
	Port int64 `json:"port"`
	// Indicates to how many nodes the data should be replicated to.
	Replication *int64 `json:"replication,omitempty"`
	// Username to use to access Cassandra.
	Username string `json:"username"`
}
