package model

import "time"

// EntityType represents entities that can be existing in database
type EntityType string

const (
	// UserEntityType represents user entity type
	UserEntityType EntityType = "user"
)

// Audit is an entity for logging
type Audit struct {
	Entity    EntityType
	Action    string
	CreatedAt time.Time
}
