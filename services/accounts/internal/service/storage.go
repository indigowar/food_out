package service

import (
	"context"
	"errors"

	"github.com/google/uuid"

	"github.com/indigowar/food_out/services/accounts/internal/domain"
)

var (
	ErrAccountNotFoundInStorage = errors.New("account was not found in the storage")
	ErrAccountAlreadyInStorage  = errors.New("account already exists in the storage")
)

//go:generate go run github.com/matryer/moq -out storage_moq_test.go . Storage

type Storage interface {
	GetAll(ctx context.Context) ([]domain.Account, error)
	GetByID(ctx context.Context, id uuid.UUID) (domain.Account, error)
	GetByPhone(ctx context.Context, phone string) (domain.Account, error)
	Add(ctx context.Context, account domain.Account) error
	Remove(ctx context.Context, id uuid.UUID) error
	Update(ctx context.Context, account domain.Account) error
}
