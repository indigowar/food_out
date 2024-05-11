package rest

import (
	"context"
	"errors"
	"net/http"

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
				Code:    api.NewOptInt(http.StatusBadRequest),
				Message: api.NewOptString(err.Error()),
			}, nil
		}

		if errors.Is(err, service.ErrSessionAlreadyExists) {
			return &api.LogInBadRequest{
				Code:    api.NewOptInt(http.StatusBadRequest),
				Message: api.NewOptString(err.Error()),
			}, nil
		}

		return &api.LogInInternalServerError{
			Code:    api.NewOptInt(http.StatusInternalServerError),
			Message: api.NewOptString(err.Error()),
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
				Code:    api.NewOptInt(http.StatusNotFound),
				Message: api.NewOptString(err.Error()),
			}, nil
		}

		return &api.LogoutFromSessionInternalServerError{
			Code:    api.NewOptInt(http.StatusInternalServerError),
			Message: api.NewOptString(err.Error()),
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
				Code:    api.NewOptInt(http.StatusBadRequest),
				Message: api.NewOptString(err.Error()),
			}, nil
		}

		return &api.RefreshAccessTokenInternalServerError{
			Code:    api.NewOptInt(http.StatusInternalServerError),
			Message: api.NewOptString(err.Error()),
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
				Code:    api.NewOptInt(http.StatusNotFound),
				Message: api.NewOptString(err.Error()),
			}, nil
		}
		return &api.RefreshPairInternalServerError{
			Code:    api.NewOptInt(http.StatusInternalServerError),
			Message: api.NewOptString(err.Error()),
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
