package services

import (
	"encoding/json"
)

type KafkaMessage struct {
	ID        string          `json:"id"`
	EventType string          `json:"event_type"` // type of event
	Payload   json.RawMessage `json:"payload"`    // Stores raw JSON data, so that it is general and can be unmarshalled in any consumer according to topic
}

