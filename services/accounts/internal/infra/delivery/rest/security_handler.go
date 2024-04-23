package rest

import (
	"context"

	"github.com/indigowar/food_out/services/accounts/internal/infra/delivery/rest/api"
)

// JwtSecurityHandler implements api.SecurityHandler,
// It should be passed to the constructor of the API Server as SecurityHandler
type JwtSecurityHandler struct {
	key []byte
}

var _ api.SecurityHandler = &JwtSecurityHandler{}

// HandleJWTAuth implements api.SecurityHandler.
func (j *JwtSecurityHandler) HandleJWTAuth(ctx context.Context, operationName string, t api.JWTAuth) (context.Context, error) {
	// todo: implement JwtSecurityHandler
	panic("unimplemented")
}
