package rest

import (
	"context"
	"errors"
	"log/slog"

	"github.com/google/uuid"

	"github.com/indigowar/food_out/services/accounts/internal/domain"
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
				Message: err.Error(),
			}, nil
		}

		return &api.CreateAccountInternalServerError{
			Message: err.Error(),
		}, nil
	}
	return &api.AccountId{ID: account.ID().String()}, nil
}

// DeleteAccount implements api.Handler.
func (w *Wrapper) DeleteAccount(ctx context.Context, params api.DeleteAccountParams) (api.DeleteAccountRes, error) {
	id, err := uuid.Parse(params.ID)
	if err != nil {
		return &api.DeleteAccountBadRequest{
			Message: "invalid account ID is provided",
		}, nil
	}

	if err := w.svc.DeleteAccount(ctx, id); err != nil {
		if errors.Is(err, service.ErrNotFound) {
			return &api.DeleteAccountNotFound{
				Message: err.Error(),
			}, nil
		}

		return &api.DeleteAccountInternalServerError{
			Message: err.Error(),
		}, nil
	}

	return &api.DeleteAccountAccepted{}, err
}

// GetAccountInfo implements api.Handler.
func (w *Wrapper) GetAccountInfo(ctx context.Context, params api.GetAccountInfoParams) (api.GetAccountInfoRes, error) {
	id, err := uuid.Parse(params.ID)
	if err != nil {
		return &api.GetAccountInfoNotFound{
			Message: "invalid account ID is provided",
		}, nil
	}

	account, err := w.svc.GetAccountByID(ctx, id)
	if err != nil {
		if errors.Is(err, service.ErrNotFound) {
			return &api.GetAccountInfoNotFound{
				Message: err.Error(),
			}, nil
		}

		return &api.GetAccountInfoInternalServerError{
			Message: err.Error(),
		}, nil
	}

	return w.toAccountInfo(account), nil
}

// GetOwnInfo implements api.Handler.
func (w *Wrapper) GetOwnInfo(ctx context.Context) (api.GetOwnInfoRes, error) {
	id, err := w.getId(ctx)
	if err != nil {
		return &api.GetOwnInfoForbidden{
			Message: "account id is not provided",
		}, nil
	}

	account, err := w.svc.GetAccountByID(ctx, id)
	if err != nil {
		if errors.Is(err, service.ErrNotFound) {
			return &api.GetOwnInfoNotFound{
				Message: err.Error(),
			}, nil
		}

		return &api.GetOwnInfoInternalServerError{
			Message: err.Error(),
		}, nil
	}

	return w.toAccountInfo(account), nil
}

// UpdatePassword implements api.Handler.
func (w *Wrapper) UpdatePassword(ctx context.Context, req *api.PasswordUpdateInfo) (api.UpdatePasswordRes, error) {
	id, err := uuid.Parse(req.ID)
	if err != nil {
		return &api.UpdatePasswordBadRequest{
			Message: "account id is not provided",
		}, nil
	}

	if err := w.svc.UpdatePassword(ctx, id, req.OldPassword, req.NewPassword); err != nil {
		if errors.Is(err, service.ErrNotFound) {
			return &api.UpdatePasswordNotFound{
				Message: err.Error(),
			}, err
		}

		if errors.Is(err, service.ErrInvalidOldPassword) {
			return &api.UpdatePasswordForbidden{
				Message: err.Error(),
			}, nil
		}

		if errors.Is(err, service.ErrInvalidValue) {
			return &api.UpdatePasswordBadRequest{
				Message: err.Error(),
			}, err
		}

		return &api.UpdatePasswordInternalServerError{
			Message: err.Error(),
		}, err
	}

	return &api.UpdatePasswordOK{}, err
}

// ValidateCredentials implements api.Handler.
func (w *Wrapper) ValidateCredentials(ctx context.Context, req api.OptAccountCredentials) (api.ValidateCredentialsRes, error) {
	id, err := w.svc.GetUserIDByCredentials(ctx, req.Value.Phone, req.Value.Password)
	if err != nil {
		if errors.Is(err, service.ErrInvalidCredentials) {
			return &api.ValidateCredentialsBadRequest{
				Message: err.Error(),
			}, nil
		}

		return &api.ValidateCredentialsInternalServerError{
			Message: err.Error(),
		}, nil
	}

	return &api.AccountId{ID: id.String()}, nil
}

func (w *Wrapper) getId(ctx context.Context) (uuid.UUID, error) {
	userId, ok := ctx.Value(securityHandlerKeyID).(string)
	if !ok {
		return uuid.UUID{}, errors.New("auth token is invalid")
	}

	id, err := uuid.Parse(userId)
	if err != nil {
		return uuid.UUID{}, errors.New("auth token is invalid")
	}

	return id, nil
}

func (w *Wrapper) toAccountInfo(account *domain.Account) *api.AccountInfo {
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

	return info
}

func NewWrapper(s *service.Service, logger *slog.Logger) *Wrapper {
	return &Wrapper{
		svc:    s,
		logger: logger,
	}
}
