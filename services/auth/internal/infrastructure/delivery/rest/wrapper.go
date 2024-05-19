package rest

import (
	"context"
	"errors"

	"github.com/indigowar/food_out/services/auth/internal/domain"
	"github.com/indigowar/food_out/services/auth/internal/infrastructure/delivery/rest/api"
	"github.com/indigowar/food_out/services/auth/internal/service"
)

// Wrapper wraps the service.Service for the generated OpenAPI specification
type Wrapper struct {
	sv *service.Service
	tk AccessTokenGenerator
}

var _ api.Handler = &Wrapper{}

// LogIn implements api.Handler.
func (w *Wrapper) LogIn(ctx context.Context, req *api.Credentials) (api.LogInRes, error) {
	session, err := w.sv.Login(ctx, req.GetPhone(), req.GetPassword())
	if err != nil {
		if errors.Is(err, service.ErrCredentialsAreInvalid) {
			return &api.LogInBadRequest{
				Message: err.Error(),
			}, nil
		}

		if errors.Is(err, service.ErrSessionAlreadyExists) {
			return &api.LogInBadRequest{
				Message: err.Error(),
			}, nil
		}

		return &api.LogInInternalServerError{
			Message: err.Error(),
		}, nil
	}

	return &api.TokenPair{
		AccessToken:  w.tk.Generate(session.ID()),
		RefreshToken: string(session.Token()),
	}, nil
}

// LogoutFromSession implements api.Handler.
func (w *Wrapper) LogoutFromSession(ctx context.Context) (api.LogoutFromSessionRes, error) {
	token := ctx.Value(securityKeyToken).(domain.SessionToken)

	if err := w.sv.Logout(ctx, token); err != nil {
		if errors.Is(err, service.ErrSessionDoesNotExist) {
			return &api.LogoutFromSessionBadRequest{
				Message: err.Error(),
			}, nil
		}

		return &api.LogoutFromSessionInternalServerError{
			Message: err.Error(),
		}, nil
	}

	return &api.LogoutFromSessionOK{}, nil
}

// RefreshAccessToken implements api.Handler.
func (w *Wrapper) RefreshAccessToken(ctx context.Context) (api.RefreshAccessTokenRes, error) {
	token := ctx.Value(securityKeyToken).(domain.SessionToken)

	session, err := w.sv.GetExistingSession(ctx, token)
	if err != nil {
		if errors.Is(err, service.ErrSessionDoesNotExist) {
			return &api.RefreshAccessTokenBadRequest{
				Message: err.Error(),
			}, nil
		}

		return &api.RefreshAccessTokenInternalServerError{
			Message: err.Error(),
		}, nil
	}

	return &api.AccessToken{
		AccessToken: w.tk.Generate(session.ID()),
	}, nil
}

// RefreshPair implements api.Handler.
func (w *Wrapper) RefreshPair(ctx context.Context) (api.RefreshPairRes, error) {
	token := ctx.Value(securityKeyToken).(domain.SessionToken)

	session, err := w.sv.RenewSession(ctx, token)
	if err != nil {
		if errors.Is(err, service.ErrSessionDoesNotExist) {
			return &api.RefreshPairBadRequest{
				Message: err.Error(),
			}, nil
		}
		return &api.RefreshPairInternalServerError{
			Message: err.Error(),
		}, nil
	}

	return &api.TokenPair{
		AccessToken:  w.tk.Generate(session.ID()),
		RefreshToken: string(session.Token()),
	}, nil
}

func NewWrapper(sv *service.Service, tk AccessTokenGenerator) *Wrapper {
	return &Wrapper{
		sv: sv,
		tk: tk,
	}
}
