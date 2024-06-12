package models

// User represents a user
type User struct {
	BaseModel
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Email     string `json:"email,omitempty"`
	Password  string `json:"password,omitempty"`
	Role      string `json:"role,omitempty"`
}
