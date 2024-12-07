package auth

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/greenblat17/auth/internal/model"
	"golang.org/x/crypto/bcrypt"
)

// VerifyPassword compare two passwords
func VerifyPassword(hashedPassword, candidatePassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(candidatePassword))
	if err != nil {
		return false
	}

	return true
}

// GenerateToken returns new token
func GenerateToken(user *model.User, secretKey []byte, duration time.Duration) (string, error) {
	claims := &model.UserClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(duration).Unix(),
		},
		Username: user.Info.Name,
		Role:     user.Info.Role,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(secretKey)
}

// VerifyToken check if token is valid
func VerifyToken(token string, secretKey []byte) (*model.UserClaims, error) {
	jwtToken, err := jwt.ParseWithClaims(token, &model.UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("unexpected token signing method")
		}

		return secretKey, nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := jwtToken.Claims.(*model.UserClaims)
	if !ok {
		return nil, errors.New("invalid token claims")
	}

	return claims, nil
}
