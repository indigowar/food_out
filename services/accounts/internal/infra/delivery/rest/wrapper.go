package rest

import (
	"context"
	"errors"
	"log/slog"
	"net/http"

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

	return w.toAccountInfo(account), nil
}

// GetOwnInfo implements api.Handler.
func (w *Wrapper) GetOwnInfo(ctx context.Context) (api.GetOwnInfoRes, error) {
	id, err := w.getId(ctx)
	if err != nil {
		return &api.GetOwnInfoForbidden{
			Code:    api.NewOptInt(http.StatusForbidden),
			Message: api.NewOptString("user is not authenticated"),
		}, nil
	}

	account, err := w.svc.GetAccountByID(ctx, id)
	if err != nil {
		if errors.Is(err, service.ErrNotFound) {
			return &api.GetOwnInfoNotFound{
				Code:    api.NewOptInt(http.StatusNotFound),
				Message: api.NewOptString("account not found"),
			}, nil
		}

		return &api.GetOwnInfoInternalServerError{
			Code:    api.NewOptInt(http.StatusInternalServerError),
			Message: api.NewOptString(err.Error()),
		}, nil
	}

	return w.toAccountInfo(account), nil
}

// UpdatePassword implements api.Handler.
func (w *Wrapper) UpdatePassword(ctx context.Context, req *api.PasswordUpdateInfo) (api.UpdatePasswordRes, error) {
	id, err := uuid.Parse(req.ID)
	if err != nil {
		return &api.UpdatePasswordBadRequest{
			Code:    api.NewOptInt(http.StatusBadRequest),
			Message: api.NewOptString("invalid id"),
		}, nil
	}

	if err := w.svc.UpdatePassword(ctx, id, req.OldPassword, req.NewPassword); err != nil {
		if errors.Is(err, service.ErrNotFound) {
			return &api.UpdatePasswordNotFound{
				Code:    api.NewOptInt(http.StatusNotFound),
				Message: api.NewOptString(err.Error()),
			}, err
		}

		if errors.Is(err, service.ErrInvalidOldPassword) {
			return &api.UpdatePasswordForbidden{
				Code:    api.NewOptInt(http.StatusForbidden),
				Message: api.NewOptString(err.Error()),
			}, nil
		}

		if errors.Is(err, service.ErrInvalidValue) {
			return &api.UpdatePasswordBadRequest{
				Code:    api.NewOptInt(http.StatusBadRequest),
				Message: api.NewOptString(err.Error()),
			}, err
		}

		return &api.UpdatePasswordInternalServerError{
			Code:    api.NewOptInt(http.StatusInternalServerError),
			Message: api.NewOptString(err.Error()),
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
