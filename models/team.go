package models

// Team represents a team
type Team struct {
	BaseModel
	Name        string  `json:"name"`
	SportID     string  `json:"sport_id"`
	Location    string  `json:"location,omitempty"`
	ContactInfo string  `json:"contact_info,omitempty"`
	LogoURL     string  `json:"logo_url,omitempty"`
	Description string  `json:"description,omitempty"`
	ManagerID   *string `json:"manager_id"`
}
