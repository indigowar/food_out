package rest

import (
	"context"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"

	"github.com/indigowar/food_out/services/accounts/internal/infra/delivery/rest/api"
)

type securityHandlerKey uint

const (
	securityHandlerKeyID securityHandlerKey = iota
)

var (
	errSecurityHandlerInvalidToken = errors.New("provided token is invalid")
)

// JwtSecurityHandler implements api.SecurityHandler,
// It should be passed to the constructor of the API Server as SecurityHandler
type JwtSecurityHandler struct {
	key []byte
}

var _ api.SecurityHandler = &JwtSecurityHandler{}

// HandleJWTAuth implements api.SecurityHandler.
func (h *JwtSecurityHandler) HandleJWTAuth(ctx context.Context, operationName string, request api.JWTAuth) (context.Context, error) {
	token, err := jwt.ParseWithClaims(request.Token, &jwt.MapClaims{}, func(t *jwt.Token) (interface{}, error) {
		return h.key, nil
	})
	if err != nil {
		return ctx, errSecurityHandlerInvalidToken
	}

	expirationTime, err := token.Claims.GetExpirationTime()
	if err != nil {
		return ctx, errSecurityHandlerInvalidToken
	}

	if time.Now().After(expirationTime.Time) {
		return ctx, errors.New("token is expired")
	}

	userId, err := token.Claims.GetSubject()
	if err != nil {
		return ctx, errSecurityHandlerInvalidToken
	}

	id, err := uuid.Parse(userId)
	if err != nil {
		return ctx, errSecurityHandlerInvalidToken
	}

	return context.WithValue(ctx, securityHandlerKeyID, id.String()), nil
}

func NewJwtSecurityHandler(key []byte) *JwtSecurityHandler {
	return &JwtSecurityHandler{
		key: key,
	}
}
