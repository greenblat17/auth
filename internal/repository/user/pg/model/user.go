package model

import (
	"database/sql"
	"time"
)

// User represents a db user
type User struct {
	ID        int64        `db:"id"`
	Info      Info         `db:""`
	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt sql.NullTime `db:"updated_at"`
}

// Info represents a user's information
type Info struct {
	Name     string `db:"name"`
	Email    string `db:"email"`
	Role     string `db:"role"`
	Password string `db:"password"`
}
