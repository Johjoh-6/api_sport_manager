package models

import (
	"time"
)

// Event represents an event
type Event struct {
    BaseModel
    TeamID      string       `json:"team_id"`
    EventTypeID string       `json:"event_type_id"`
    Name        string    `json:"name"`
    Location    string    `json:"location,omitempty"`
    DateStart   time.Time `json:"date_start"`
    DateEnd     *time.Time `json:"date_end,omitempty"`
}
