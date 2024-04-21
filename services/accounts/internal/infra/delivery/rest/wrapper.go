package rest

import (
	"context"
	"errors"
	"log/slog"
	"net/http"

	"github.com/google/uuid"
	"github.com/indigowar/food_out/services/accounts/internal/infra/delivery/rest/api"
	"github.com/indigowar/food_out/services/accounts/internal/service"
)

type Wrapper struct {
	svc    *service.Service
	logger *slog.Logger
}

var _ api.Handler = &Wrapper{}

// CreateAccount implements api.Handler.
func (w *Wrapper) CreateAccount(ctx context.Context, req *api.AccountCreationInfo) (api.CreateAccountRes, error) {
	account, err := w.svc.CreateAccount(ctx, req.Phone, req.Password)
	if err != nil {
		if errors.Is(err, service.ErrInvalidCredentials) || errors.Is(err, service.ErrPhoneNumberAlreadyInUse) {
			return &api.CreateAccountBadRequest{
				Message: api.NewOptString(err.Error()),
			}, nil
		}

		return &api.CreateAccountInternalServerError{
			Message: api.NewOptString(err.Error()),
		}, nil
	}
	return &api.AccountId{ID: account.ID().String()}, nil
}

// DeleteAccount implements api.Handler.
func (w *Wrapper) DeleteAccount(ctx context.Context, params api.DeleteAccountParams) (api.DeleteAccountRes, error) {
	panic("unimplemented")
}

// GetAccountInfo implements api.Handler.
func (w *Wrapper) GetAccountInfo(ctx context.Context, params api.GetAccountInfoParams) (api.GetAccountInfoRes, error) {
	id, err := uuid.Parse(params.ID)
	if err != nil {
		return &api.GetAccountInfoNotFound{
			Code:    api.NewOptInt(http.StatusBadRequest),
			Message: api.NewOptString("invalid id"),
		}, nil
	}

	account, err := w.svc.GetAccountByID(ctx, id)
	if err != nil {
		if errors.Is(err, service.ErrNotFound) {
			return &api.GetAccountInfoNotFound{
				Code:    api.NewOptInt(http.StatusNotFound),
				Message: api.NewOptString("account not found"),
			}, nil
		}

		return &api.GetAccountInfoInternalServerError{
			Code:    api.NewOptInt(http.StatusInternalServerError),
			Message: api.NewOptString(err.Error()),
		}, nil
	}

	info := &api.AccountInfo{
		ID:    account.ID().String(),
		Phone: account.Phone(),
	}

	if account.HasName() {
		info.Name = api.NewOptString(account.Name())
	}

	if account.HasProfilePicture() {
		info.Profile = api.NewOptString(account.ProfilePicture().String())
	}

	return info, nil
}

// GetOwnInfo implements api.Handler.
func (w *Wrapper) GetOwnInfo(ctx context.Context) (api.GetOwnInfoRes, error) {
	// todo: implement GetOwnInfo
	panic("unimplemented")
}

// UpdatePassword implements api.Handler.
func (w *Wrapper) UpdatePassword(ctx context.Context, req *api.PasswordUpdateInfo) (api.UpdatePasswordRes, error) {
	// todo: implement UpdatePassword
	panic("unimplemented")
}

// ValidateCredentials implements api.Handler.
func (w *Wrapper) ValidateCredentials(ctx context.Context, req api.OptAccountCredentials) (api.ValidateCredentialsRes, error) {
	id, err := w.svc.GetUserIDByCredentials(ctx, req.Value.Phone, req.Value.Password)
	if err != nil {
		if errors.Is(err, service.ErrInvalidCredentials) {
			return &api.ValidateCredentialsBadRequest{
				Message: api.NewOptString(err.Error()),
			}, nil
		}

		return &api.ValidateCredentialsInternalServerError{
			Code:    api.OptInt{},
			Message: api.NewOptString(err.Error()),
		}, nil
	}

	return &api.AccountId{ID: id.String()}, nil
}

func NewWrapper(s *service.Service, logger *slog.Logger) *Wrapper {
	return &Wrapper{
		svc:    s,
		logger: logger,
	}
}
