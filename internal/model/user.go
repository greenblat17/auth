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

func (u *UserInfo) IsEmptyName() bool {
	return u.Name == ""
}

func (u *UserInfo) IsEmptyEmail() bool {
	return u.Email == ""
}

func (u *UserInfo) IsEmptyRole() bool {
	return u.Role == ""
}
