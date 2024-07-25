package service

import (
	"context"
	"errors"

	"github.com/google/uuid"

	"github.com/indigowar/food_out/services/auth/internal/domain"
)

//go:generate go run github.com/matryer/moq -out storage_moq_test.go . Storage

var (
	ErrStorageNotFound      = errors.New("session is not found in the storage")
	ErrStorageAlreadyExists = errors.New("session already exists in the storage")
)

type Storage interface {
	GetByID(ctx context.Context, id uuid.UUID) (domain.Session, error)
	GetByToken(ctx context.Context, token domain.SessionToken) (domain.Session, error)
	Add(ctx context.Context, session domain.Session) error
	RemoveByID(ctx context.Context, id uuid.UUID) (domain.Session, error)
	RemoveByToken(ctx context.Context, token domain.SessionToken) (domain.Session, error)
}
