package models

// Position represents a player position
type Position struct {
	ID           string `json:"id,omitempty"`
	PositionName string `json:"position_name"`
}
