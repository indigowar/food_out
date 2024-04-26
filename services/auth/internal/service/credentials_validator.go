package service

import (
	"context"
	"errors"

	"github.com/google/uuid"
)

//go:generate moq -out credentials_validator_moq.go . CredentialsValidator

var (
	ErrCredentialsValidatorNotFound = errors.New("credentials validator error: account not found")
	ErrCredentialsValidatorInvalid  = errors.New("credentials validator error: invalid input")
)

type CredentialsValidator interface {
	Validate(ctx context.Context, phone string, password string) (uuid.UUID, error)
}
