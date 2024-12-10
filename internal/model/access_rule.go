package model

// AccessRule represents access to endpoint for rules
type AccessRule struct {
	Endpoint string
	Role     map[string]struct{}
}
