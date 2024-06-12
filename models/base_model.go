package models

import (
	"time"
)

type BaseModel struct {
	Id      string    `json:"id,omitempty"`
	Created time.Time `json:"created_at,omitempty"`
	Updated time.Time `json:"updated_at,omitempty"`
}

// GetId returns the model id.
func (m *BaseModel) GetId() string {
	return m.Id
}

// SetId sets the model id to the provided string value.
func (m *BaseModel) SetId(id string) {
	m.Id = id
}

func (m *BaseModel) CreatedIsZero() bool {
	return m.Created.IsZero()
}

func (m *BaseModel) UpdatedIsZero() bool {
	return m.Updated.IsZero()
}
