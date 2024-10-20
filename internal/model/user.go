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
	Info      UserInfo
	CreatedAt time.Time
	UpdatedAt *time.Time
}

type UserInfo struct {
	Name     string
	Email    string
	Password string
	Role     string
}
