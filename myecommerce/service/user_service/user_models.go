package model

// User defines the structure of a user
type User struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Mobile string `json:"mobile"`
	Email  string `json:"email,omitempty"`
}

