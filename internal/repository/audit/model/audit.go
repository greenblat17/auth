package model

import "time"

// Audit represents entity in database
type Audit struct {
	ID        int64     `db:"id"`
	Entity    string    `db:"entity"`
	Action    string    `db:"action"`
	CreatedAt time.Time `db:"created_at"`
}
