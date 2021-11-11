package service

import (
	"github.com/dgrijalva/jwt-go"
)

// Jwt Jwt related operations.
type Jwt interface {
	ValidateToken(tokenString string) (bool, *jwt.Token)
}
