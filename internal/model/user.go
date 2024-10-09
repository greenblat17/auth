package model

import "time"

// Role user
const (
	_         = "UNKNOWN"
	RoleAdmin = "ADMIN"
	RoleUser  = "USER"
)

// User model
type User struct {
	ID        int64
	Name      string
	Email     string
	Password  string
	Role      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
