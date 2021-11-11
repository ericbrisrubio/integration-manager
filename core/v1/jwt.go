package v1

import (
	"crypto/rsa"
)

// Jwt contains jwt keys
type Jwt struct {
	PublicKey *rsa.PublicKey
}
