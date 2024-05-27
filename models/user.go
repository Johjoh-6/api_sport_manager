package models

// User represents a user
type User struct {
	BaseModel
	FirstName    string `json:"first_name,omitempty"`
	LastName     string `json:"last_name,omitempty"`
	Email        string `json:"email,omitempty"`
	PasswordHash string `json:"password_hash,omitempty"`
	Role         string `json:"role,omitempty"`
}
