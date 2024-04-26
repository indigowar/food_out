package service

import "log/slog"

type Service struct {
	logger               *slog.Logger
	storage              Storage
	credentialsValidator CredentialsValidator
}

func NewService(logger *slog.Logger, storage Storage, credentialsValidator CredentialsValidator) *Service {
	return &Service{
		logger:               logger,
		storage:              storage,
		credentialsValidator: credentialsValidator,
	}
}
