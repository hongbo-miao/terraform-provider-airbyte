// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// DestinationPulsarCompressionTypeEnum - Compression type for the producer.
type DestinationPulsarCompressionTypeEnum string

const (
	DestinationPulsarCompressionTypeEnumNone   DestinationPulsarCompressionTypeEnum = "NONE"
	DestinationPulsarCompressionTypeEnumLz4    DestinationPulsarCompressionTypeEnum = "LZ4"
	DestinationPulsarCompressionTypeEnumZlib   DestinationPulsarCompressionTypeEnum = "ZLIB"
	DestinationPulsarCompressionTypeEnumZstd   DestinationPulsarCompressionTypeEnum = "ZSTD"
	DestinationPulsarCompressionTypeEnumSnappy DestinationPulsarCompressionTypeEnum = "SNAPPY"
)

func (e DestinationPulsarCompressionTypeEnum) ToPointer() *DestinationPulsarCompressionTypeEnum {
	return &e
}

func (e *DestinationPulsarCompressionTypeEnum) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "NONE":
		fallthrough
	case "LZ4":
		fallthrough
	case "ZLIB":
		fallthrough
	case "ZSTD":
		fallthrough
	case "SNAPPY":
		*e = DestinationPulsarCompressionTypeEnum(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationPulsarCompressionTypeEnum: %v", v)
	}
}

type DestinationPulsarPulsarEnum string

const (
	DestinationPulsarPulsarEnumPulsar DestinationPulsarPulsarEnum = "pulsar"
)

func (e DestinationPulsarPulsarEnum) ToPointer() *DestinationPulsarPulsarEnum {
	return &e
}

func (e *DestinationPulsarPulsarEnum) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "pulsar":
		*e = DestinationPulsarPulsarEnum(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationPulsarPulsarEnum: %v", v)
	}
}

// DestinationPulsarTopicTypeEnum - It identifies type of topic. Pulsar supports two kind of topics: persistent and non-persistent. In persistent topic, all messages are durably persisted on disk (that means on multiple disks unless the broker is standalone), whereas non-persistent topic does not persist message into storage disk.
type DestinationPulsarTopicTypeEnum string

const (
	DestinationPulsarTopicTypeEnumPersistent    DestinationPulsarTopicTypeEnum = "persistent"
	DestinationPulsarTopicTypeEnumNonPersistent DestinationPulsarTopicTypeEnum = "non-persistent"
)

func (e DestinationPulsarTopicTypeEnum) ToPointer() *DestinationPulsarTopicTypeEnum {
	return &e
}

func (e *DestinationPulsarTopicTypeEnum) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "persistent":
		fallthrough
	case "non-persistent":
		*e = DestinationPulsarTopicTypeEnum(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationPulsarTopicTypeEnum: %v", v)
	}
}

// DestinationPulsar - The values required to configure the destination.
type DestinationPulsar struct {
	// Control whether automatic batching of messages is enabled for the producer.
	BatchingEnabled bool `json:"batching_enabled"`
	// Maximum number of messages permitted in a batch.
	BatchingMaxMessages int64 `json:"batching_max_messages"`
	//  Time period in milliseconds within which the messages sent will be batched.
	BatchingMaxPublishDelay int64 `json:"batching_max_publish_delay"`
	// If the send operation should block when the outgoing message queue is full.
	BlockIfQueueFull bool `json:"block_if_queue_full"`
	// A list of host/port pairs to use for establishing the initial connection to the Pulsar cluster.
	Brokers string `json:"brokers"`
	// Compression type for the producer.
	CompressionType DestinationPulsarCompressionTypeEnum `json:"compression_type"`
	DestinationType DestinationPulsarPulsarEnum          `json:"destinationType"`
	// The maximum size of a queue holding pending messages.
	MaxPendingMessages int64 `json:"max_pending_messages"`
	// The maximum number of pending messages across partitions.
	MaxPendingMessagesAcrossPartitions int64 `json:"max_pending_messages_across_partitions"`
	// Name for the producer. If not filled, the system will generate a globally unique name which can be accessed with.
	ProducerName *string `json:"producer_name,omitempty"`
	// Wait synchronously until the record has been sent to Pulsar.
	ProducerSync *bool `json:"producer_sync,omitempty"`
	// If a message is not acknowledged by a server before the send-timeout expires, an error occurs (in ms).
	SendTimeoutMs int64 `json:"send_timeout_ms"`
	// The administrative unit of the topic, which acts as a grouping mechanism for related topics. Most topic configuration is performed at the namespace level. Each tenant has one or multiple namespaces.
	TopicNamespace string `json:"topic_namespace"`
	// Topic pattern in which the records will be sent. You can use patterns like '{namespace}' and/or '{stream}' to send the message to a specific topic based on these values. Notice that the topic name will be transformed to a standard naming convention.
	TopicPattern string `json:"topic_pattern"`
	// The topic tenant within the instance. Tenants are essential to multi-tenancy in Pulsar, and spread across clusters.
	TopicTenant string `json:"topic_tenant"`
	// Topic to test if Airbyte can produce messages.
	TopicTest *string `json:"topic_test,omitempty"`
	// It identifies type of topic. Pulsar supports two kind of topics: persistent and non-persistent. In persistent topic, all messages are durably persisted on disk (that means on multiple disks unless the broker is standalone), whereas non-persistent topic does not persist message into storage disk.
	TopicType DestinationPulsarTopicTypeEnum `json:"topic_type"`
	// Whether to use TLS encryption on the connection.
	UseTLS bool `json:"use_tls"`
}
