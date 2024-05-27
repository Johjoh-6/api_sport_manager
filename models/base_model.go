package models

import (
	"time"
)

type BaseModel struct {
	Id      string     `json:"id,omitempty"`
	Created *time.Time `json:"created_at,omitempty"`
	Updated *time.Time `json:"updated_at,omitempty"`
}

// GetId returns the model id.
func (m *BaseModel) GetId() string {
	return m.Id
}

// SetId sets the model id to the provided string value.
func (m *BaseModel) SetId(id string) {
	m.Id = id
}

// GetCreated returns the model Created datetime.
func (m *BaseModel) GetCreated() *time.Time {
	return m.Created
}

// GetUpdated returns the model Updated datetime.
func (m *BaseModel) GetUpdated() *time.Time {
	return m.Updated
}

// RefreshUpdated updates the model Updated field with the current datetime.
func (m *BaseModel) RefreshUpdated() {
	now := time.Now()
	m.Updated = &now
}
