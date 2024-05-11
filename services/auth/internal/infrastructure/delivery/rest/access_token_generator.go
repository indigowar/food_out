package rest

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type AccessTokenGenerator interface {
	Generate(id uuid.UUID) string
}

type JwtTokenGenerator struct {
	key      []byte
	duration time.Duration
}

var _ AccessTokenGenerator = &JwtTokenGenerator{}

// Generate implements AccessTokenGenerator.
func (gen *JwtTokenGenerator) Generate(id uuid.UUID) string {
	claims := &jwt.RegisteredClaims{
		Issuer:    "auth-service",
		Subject:   id.String(),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(gen.duration)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	}

	t := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	token, _ := t.SignedString(gen.key)

	return token
}

func NewJwtTokenGenerator(key []byte, duration time.Duration) *JwtTokenGenerator {
	return &JwtTokenGenerator{
		key:      key,
		duration: duration,
	}
}
