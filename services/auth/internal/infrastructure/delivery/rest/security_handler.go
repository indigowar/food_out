package rest

import (
	"context"

	"github.com/indigowar/food_out/services/auth/internal/domain"
	"github.com/indigowar/food_out/services/auth/internal/infrastructure/delivery/rest/gen"
	"github.com/indigowar/food_out/services/auth/internal/service"
)

type securityKey string

const (
	securityKeyID    securityKey = "security_key_id"
	securityKeyToken securityKey = "security_key_token"
)

type RefreshSecurityHandler struct {
	sv *service.Service
}

// HandleRefreshAuth implements api.SecurityHandler.
func (h *RefreshSecurityHandler) HandleRefreshAuth(ctx context.Context, operationName string, t gen.RefreshAuth) (context.Context, error) {
	token := domain.SessionToken(t.GetAPIKey())
	session, err := h.sv.GetExistingSession(ctx, token)
	if err != nil {
		return ctx, err
	}

	return context.WithValue(
		context.WithValue(ctx, securityKeyID, session.ID()),
		securityKeyToken, token,
	), nil
}

var _ gen.SecurityHandler = &RefreshSecurityHandler{}

func NewRefreshSecurityHandler(sv *service.Service) *RefreshSecurityHandler {
	return &RefreshSecurityHandler{
		sv: sv,
	}
}
