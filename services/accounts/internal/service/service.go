package service

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/url"

	"github.com/google/uuid"

	"github.com/indigowar/food_out/services/accounts/internal/domain"
)

var (
	ErrNotFound                = errors.New("account is not found")
	ErrPhoneNumberAlreadyInUse = errors.New("phone number is already in use")
	ErrInvalidCredentials      = errors.New("invalid credentials")
	ErrInvalidOldPassword      = errors.New("provided old password is invalid")
	ErrInvalidValue            = errors.New("provided value is invalid")
	ErrInternal                = errors.New("internal service error")
)

type Service struct {
	logger *slog.Logger

	storage Storage

	accountCreatedPublisher AccountCreatedPublisher
	accountDeletedPublisher AccountDeletedPublisher
}

func (svc *Service) GetAccountByID(ctx context.Context, id uuid.UUID) (*domain.Account, error) {
	account, err := svc.storage.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, ErrAccountNotFoundInStorage) {
			return nil, ErrNotFound
		}

		svc.logger.Warn(
			"Internal Server Error",
			"service", "accounts",
			"action", "get_by_id",
			"error", err.Error(),
		)

		return nil, ErrInternal
	}

	return account, nil
}

func (svc *Service) GetAccountByPhone(ctx context.Context, phone string) (*domain.Account, error) {
	if err := domain.ValidatePhoneNumber(phone); err != nil {
		return nil, fmt.Errorf("%w:%s", ErrInvalidValue, err)
	}

	account, err := svc.storage.GetByPhone(ctx, phone)
	if err != nil {
		if errors.Is(err, ErrAccountNotFoundInStorage) {
			return nil, ErrNotFound
		}

		svc.logger.Warn(
			"Internal Server Error",
			"service", "accounts",
			"action", "get_by_phone",
			"error", err.Error(),
		)

		return nil, ErrInternal
	}

	return account, nil
}

func (svc *Service) CreateAccount(ctx context.Context, phone string, password string) (*domain.Account, error) {
	if err := svc.validateCredentials(phone, password); err != nil {
		return nil, fmt.Errorf("%w:%s", ErrInvalidCredentials, err)
	}

	account, err := domain.NewAccount(phone, password)
	if err != nil {
		return nil, fmt.Errorf("%w:%s", ErrInvalidCredentials, err)
	}

	if err := svc.storage.Add(ctx, account); err != nil {
		if errors.Is(err, ErrAccountAlreadyInStorage) {
			return nil, ErrPhoneNumberAlreadyInUse
		}

		svc.logger.Warn(
			"Internal Server Error",
			"service", "accounts",
			"action", "create_account",
			"error", err.Error(),
		)

		return nil, ErrInternal
	}

	svc.accountCreatedPublisher.PublishAccountCreated(ctx, account)

	return account, nil
}

func (svc *Service) DeleteAccount(ctx context.Context, id uuid.UUID) error {
	if err := svc.storage.Remove(ctx, id); err != nil {
		if errors.Is(err, ErrAccountNotFoundInStorage) {
			return ErrNotFound
		}

		svc.logger.Warn(
			"Internal Server Error",
			"service", "accounts",
			"action", "delete_account",
			"error", err.Error(),
		)

		return ErrInternal
	}

	svc.accountDeletedPublisher.PublishAccountDeleted(ctx, id)

	return nil
}

func (svc *Service) UpdatePassword(ctx context.Context, id uuid.UUID, oldPassword string, password string) error {
	account, err := svc.GetAccountByID(ctx, id)
	if err != nil {
		svc.logger.Info(
			"Action stopped",
			"service", "accounts",
			"action", "update_password",
			"stopped by", "get_account_by_id",
			"error", err.Error(),
		)
		return err
	}

	if !account.IsPasswordEqual(oldPassword) {
		return ErrInvalidOldPassword
	}

	if err := account.SetPassword(password); err != nil {
		return fmt.Errorf("%w:%s", ErrInvalidValue, err)
	}

	return svc.updateAccount(ctx, account, "update_password")
}

func (svc *Service) UpdateName(ctx context.Context, id uuid.UUID, name string) error {
	account, err := svc.storage.GetByID(ctx, id)
	if err != nil {
		svc.logger.Info(
			"Action stopped",
			"service", "accounts",
			"action", "update_password",
			"stopped by", "get_account_by_id",
			"error", err.Error(),
		)
		return err
	}

	if err := account.SetName(name); err != nil {
		return fmt.Errorf("%w:%s", ErrInvalidValue, err)
	}

	return svc.updateAccount(ctx, account, "update_name")
}

func (svc *Service) UpdateProfilePicture(ctx context.Context, id uuid.UUID, profilePicture *url.URL) error {
	account, err := svc.storage.GetByID(ctx, id)
	if err != nil {
		svc.logger.Info(
			"Action stopped",
			"service", "accounts",
			"action", "update_profile_picture",
			"stopped by", "get_account_by_id",
			"error", err.Error(),
		)
		return err
	}

	account.SetProfilePicture(profilePicture)

	return svc.updateAccount(ctx, account, "update_profile_picture")
}

func (svc *Service) GetUserIDByCredentials(ctx context.Context, phone string, password string) (uuid.UUID, error) {
	if err := svc.validateCredentials(phone, password); err != nil {
		return uuid.UUID{}, fmt.Errorf("%w:%s", ErrInvalidCredentials, err)
	}

	account, err := svc.storage.GetByPhone(ctx, phone)
	if err != nil {
		if errors.Is(err, ErrAccountNotFoundInStorage) {
			return uuid.UUID{}, ErrInvalidCredentials
		}

		svc.logger.Warn(
			"Internal Server Error",
			"service", "accounts",
			"action", "get_user_id_by_credentials",
			"error", err.Error(),
		)

		return uuid.UUID{}, ErrInternal
	}

	if !account.IsPasswordEqual(password) {
		return uuid.UUID{}, ErrInvalidCredentials
	}

	return account.ID(), nil
}

func (svc *Service) validateCredentials(phone string, password string) error {
	phoneErr := domain.ValidatePhoneNumber(phone)
	passwordErr := domain.ValidatePassword(password)
	return errors.Join(phoneErr, passwordErr)
}

func (svc *Service) updateAccount(ctx context.Context, account *domain.Account, masterAction string) error {
	if err := svc.storage.Update(ctx, account); err != nil {
		svc.logger.Info(
			"Action stopped",
			"service", "accounts",
			"action", masterAction,
			"stopped by", "update_account",
			"error", err,
		)

		if errors.Is(err, ErrAccountNotFoundInStorage) {
			return ErrNotFound
		}

		if errors.Is(err, ErrAccountAlreadyInStorage) {
			return ErrPhoneNumberAlreadyInUse
		}

		svc.logger.Warn(
			"Internal Server Error",
			"service", "accounts",
			"action", "update_account",
			"error", err.Error(),
		)

		return ErrInternal
	}

	return nil
}

func NewService(
	logger *slog.Logger,
	storage Storage,
	accountCreatedPublisher AccountCreatedPublisher,
	accountDeletedPublisher AccountDeletedPublisher,
) *Service {
	return &Service{
		logger:                  logger,
		storage:                 storage,
		accountCreatedPublisher: accountCreatedPublisher,
		accountDeletedPublisher: accountDeletedPublisher,
	}
}
