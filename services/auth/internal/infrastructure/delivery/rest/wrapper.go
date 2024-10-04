package rest

import (
	"context"
	"errors"

	"github.com/indigowar/food_out/services/auth/internal/domain"
	"github.com/indigowar/food_out/services/auth/internal/infrastructure/delivery/rest/gen"
	"github.com/indigowar/food_out/services/auth/internal/service"
)

// Wrapper wraps the service.Service for the generated OpenAPI specification
type Wrapper struct {
	sv *service.Service
	tk AccessTokenGenerator
}

var _ gen.Handler = &Wrapper{}

// LogIn implements api.Handler.
func (w *Wrapper) LogIn(ctx context.Context, req *gen.Credentials) (gen.LogInRes, error) {
	session, err := w.sv.Login(ctx, req.GetPhone(), req.GetPassword())
	if err != nil {
		if errors.Is(err, service.ErrCredentialsAreInvalid) {
			return &gen.LogInBadRequest{
				Message: err.Error(),
			}, nil
		}

		if errors.Is(err, service.ErrSessionAlreadyExists) {
			return &gen.LogInBadRequest{
				Message: err.Error(),
			}, nil
		}

		return &gen.LogInInternalServerError{
			Message: err.Error(),
		}, nil
	}

	return &gen.TokenPair{
		AccessToken:  w.tk.Generate(session.ID()),
		RefreshToken: string(session.Token()),
	}, nil
}

// LogoutFromSession implements api.Handler.
func (w *Wrapper) LogoutFromSession(ctx context.Context) (gen.LogoutFromSessionRes, error) {
	token := ctx.Value(securityKeyToken).(domain.SessionToken)

	if err := w.sv.Logout(ctx, token); err != nil {
		if errors.Is(err, service.ErrSessionDoesNotExist) {
			return &gen.LogoutFromSessionBadRequest{
				Message: err.Error(),
			}, nil
		}

		return &gen.LogoutFromSessionInternalServerError{
			Message: err.Error(),
		}, nil
	}

	return &gen.LogoutFromSessionOK{}, nil
}

// RefreshAccessToken implements api.Handler.
func (w *Wrapper) RefreshAccessToken(ctx context.Context) (gen.RefreshAccessTokenRes, error) {
	token := ctx.Value(securityKeyToken).(domain.SessionToken)

	session, err := w.sv.GetExistingSession(ctx, token)
	if err != nil {
		if errors.Is(err, service.ErrSessionDoesNotExist) {
			return &gen.RefreshAccessTokenBadRequest{
				Message: err.Error(),
			}, nil
		}

		return &gen.RefreshAccessTokenInternalServerError{
			Message: err.Error(),
		}, nil
	}

	return &gen.AccessToken{
		AccessToken: w.tk.Generate(session.ID()),
	}, nil
}

// RefreshPair implements api.Handler.
func (w *Wrapper) RefreshPair(ctx context.Context) (gen.RefreshPairRes, error) {
	token := ctx.Value(securityKeyToken).(domain.SessionToken)

	session, err := w.sv.RenewSession(ctx, token)
	if err != nil {
		if errors.Is(err, service.ErrSessionDoesNotExist) {
			return &gen.RefreshPairBadRequest{
				Message: err.Error(),
			}, nil
		}
		return &gen.RefreshPairInternalServerError{
			Message: err.Error(),
		}, nil
	}

	return &gen.TokenPair{
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
