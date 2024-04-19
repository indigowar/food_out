package service

import (
	"context"
	"errors"
	"log/slog"
	"net/url"

	"github.com/google/uuid"

	"github.com/indigowar/food_out/services/accounts/internal/domain"
)

var (
	ErrNotFound                = errors.New("account is not found")
	ErrPhoneNumberAlreadyInUse = errors.New("phone number is already in use")
	ErrInvalidCredentials      = errors.New("invalid credentials")
	ErrInvalidValue            = errors.New("provided value is invalid")
	ErrInternal                = errors.New("internal service error")
)

type Service struct {
	storage Storage
	logger  *slog.Logger
}

func (svc *Service) GetAccountByID(ctx context.Context, id uuid.UUID) (*domain.Account, error) {
	panic("not implemented")
}

func (svc *Service) GetAccountByPhone(ctx context.Context, phone string) (*domain.Account, error) {
	panic("not implemented")
}

func (svc *Service) CreateAccount(ctx context.Context, phone string, password string) (*domain.Account, error) {
	panic("not implemented")
}

func (svc *Service) DeleteAccount(ctx context.Context, id uuid.UUID) error {
	panic("not implemented")
}

func (svc *Service) UpdatePassword(ctx context.Context, id uuid.UUID, password string) error {
	panic("not implemented")
}

func (svc *Service) UpdateName(ctx context.Context, id uuid.UUID, name string) error {
	panic("not implemented")
}

func (svc *Service) UpdateProfilePicture(ctx context.Context, id uuid.UUID, profilePicture *url.URL) error {
	panic("not implemented")
}

func (svc *Service) GetUserIDByCredentials(ctx context.Context, phone string, password string) (uuid.UUID, error) {
	// account, err := svc.storage.GetByPhone(ctx, phone)
	// if err != nil {
	// 	if errors.Is(err, ErrAccountNotFoundInStorage) {
	// 		svc.logger.Info("Searched user account by phone does not exists")
	// 		return uuid.UUID{}, ErrNotFound
	// 	}
	// 	svc.logger.Warn("Failed to search user account by id")
	//
	// 	return uuid.UUID{}, ErrInternal
	// }
	//
	// if account.Password() != password {
	// 	svc.logger.Info("Password check failed, while attempting to access user id by credentials")
	// 	return uuid.UUID{}, ErrNotFound
	// }
	//
	// return account.ID(), nil
	panic("not implemented")
}

func NewService(storage Storage, logger *slog.Logger) *Service {
	return &Service{
		storage: storage,
		logger:  logger,
	}
}
