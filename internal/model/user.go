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

// UserInfo model
type UserInfo struct {
	Name     string
	Email    string
	Password string
	Role     string
}

// IsEmptyName check if username is empty
func (u *UserInfo) IsEmptyName() bool {
	return u.Name == ""
}

// IsEmptyEmail check if user email is empty
func (u *UserInfo) IsEmptyEmail() bool {
	return u.Email == ""
}

// IsEmptyRole check if user role is empty
func (u *UserInfo) IsEmptyRole() bool {
	return u.Role == ""
}
