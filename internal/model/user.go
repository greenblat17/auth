package model

import "time"

// Role user
const (
	RoleUnknown = "UNKNOWN"
	RoleAdmin   = "ADMIN"
	RoleUser    = "USER"
)

// User model
type User struct {
	ID        int64
	Name      string
	Email     string
	Password  string
	Role      string
	CreatedAt time.Time
	UpdatedAt *time.Time
}

// GetUpdateAt get update time value
func (u *User) GetUpdateAt() time.Time {
	if u.UpdatedAt != nil {
		return *u.UpdatedAt
	}

	return time.Time{}
}
