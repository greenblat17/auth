package model

import "github.com/golang-jwt/jwt"

// UserClaims represents user claims for token
type UserClaims struct {
	jwt.StandardClaims
	Username string `json:"username"`
	Role     string `json:"role"`
}
