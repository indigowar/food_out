package service

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"time"

	"github.com/indigowar/food_out/services/auth/internal/domain"
)

var (
	ErrCredentialsAreInvalid = errors.New("credentials are invalid")
	ErrSessionAlreadyExists  = errors.New("session already exists")
	ErrSessionDoesNotExist   = errors.New("session does not exist")
	ErrInternal              = errors.New("internal error occurred")
)

type Service struct {
	logger               *slog.Logger
	storage              Storage
	credentialsValidator CredentialsValidator

	sessionDuration time.Duration
}

func (s *Service) Login(ctx context.Context, phone string, password string) (domain.Session, error) {
	id, err := s.credentialsValidator.Validate(ctx, phone, password)
	if err != nil {
		if errors.Is(err, ErrCredentialsValidatorInvalid) || errors.Is(err, ErrCredentialsValidatorNotFound) {
			return domain.Session{}, ErrCredentialsAreInvalid
		}
		return domain.Session{}, s.handleInternalError("Login", "CredentialsValidator", "Validate", err)
	}

	session := domain.NewSession(id, domain.GenerateSessionToken(), time.Now().Add(s.sessionDuration))

	if err := s.storage.Add(ctx, session); err != nil {
		if errors.Is(err, ErrStorageAlreadyExists) {
			return domain.Session{}, ErrSessionAlreadyExists
		}
		return domain.Session{}, s.handleInternalError("Login", "Storage", "Add", err)
	}

	return session, nil
}

func (s *Service) Logout(ctx context.Context, token domain.SessionToken) error {
	if _, err := s.storage.RemoveByToken(ctx, token); err != nil {
		if errors.Is(err, ErrStorageNotFound) {
			return ErrSessionDoesNotExist
		}
		return s.handleInternalError("Logout", "Storage", "RemoveByToken", err)
	}

	return nil
}

func (s *Service) RenewSession(ctx context.Context, token domain.SessionToken) (domain.Session, error) {
	session, err := s.storage.RemoveByToken(ctx, token)
	if err != nil {
		if errors.Is(err, ErrStorageNotFound) {
			return domain.Session{}, ErrSessionDoesNotExist
		}
		return domain.Session{}, s.handleInternalError("RenewSession", "Storage", "RemoveByToken", err)
	}

	newSession := domain.NewSession(session.ID(), token, time.Now().Add(s.sessionDuration))

	if err := s.storage.Add(ctx, newSession); err != nil {
		if errors.Is(err, ErrStorageAlreadyExists) {
			return domain.Session{}, s.handleInternalError("RenewSession", "Storage", "Add", fmt.Errorf("%w: after removal", err))
		}
		return domain.Session{}, s.handleInternalError("RenewSession", "Storage", "Add", err)
	}

	return newSession, nil
}

func (s *Service) handleInternalError(action string, reason string, subAction string, err error) error {
	s.logger.Warn(
		"Service Operation FAILED",
		"action", action,
		"reason", reason,
		"sub-action", subAction,
		"error", err.Error(),
	)

	return ErrInternal
}

func NewService(logger *slog.Logger, storage Storage, credentialsValidator CredentialsValidator) *Service {
	return &Service{
		logger:               logger,
		storage:              storage,
		credentialsValidator: credentialsValidator,
	}
}
