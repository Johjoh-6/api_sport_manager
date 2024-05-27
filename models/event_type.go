package models

// EventType represents an event type
type EventType struct {
	ID            string `json:"id,omitempty"`
	EventTypeName string `json:"event_type_name"`
}
