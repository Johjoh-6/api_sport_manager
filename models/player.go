package models

// Player represents a player
type Player struct {
	BaseModel
	UserLink     *string `json:"user_link,omitempty"`
	TeamID       *string `json:"team_id,omitempty"`
	FirstName    string  `json:"first_name"`
	LastName     string  `json:"last_name"`
	PositionID   *int    `json:"position_id,omitempty"`
	Biography    string  `json:"biography,omitempty"`
	Weight       *int    `json:"weight,omitempty"`
	Height       *int    `json:"height,omitempty"`
	PlayerNumber *int    `json:"player_number,omitempty"`
	Claimed      bool    `json:"claimed"`
	// to check for later
	// BirthDate    time.Time `json:"birth_date"`
	// ImageUrl	 string    `json:"image_url,omitempty"`
	// HandycapComment      string    `json:"handy_comment,omitempty"`
	// HandycapPoints       int       `json:"handy_points"`
}
