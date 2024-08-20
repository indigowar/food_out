package commands

import (
	"context"
	"fmt"

	"github.com/google/uuid"

	"github.com/indigowar/services/orders/models"
)

//go:generate go run github.com/matryer/moq -out storage_moq_test.go . OrderStorage

type StorageErrorType uint

const (
	StorageErrorTypeNotFound StorageErrorType = iota
	StorageErrorTypeAlreadyExists
)

type OrderStorage interface {
	Get(ctx context.Context, id uuid.UUID) (models.Order, error)
	Add(ctx context.Context, order models.Order) error
	Delete(ctx context.Context, id uuid.UUID) error
	Update(ctx context.Context, order models.Order) error
}

type StorageError struct {
	ErrorType StorageErrorType
	Object    string
	Field     string
	Message   error
}

func (e *StorageError) Error() string {
	return fmt.Sprintf("storage error: %s for %s: %s", e.Object, e.Field, e.Message)
}

type orderStorageStubFuncs struct{}

func (s *orderStorageStubFuncs) Get(o models.Order, e error) func(context.Context, uuid.UUID) (models.Order, error) {
	return func(context.Context, uuid.UUID) (models.Order, error) { return o, e }
}

func (s *orderStorageStubFuncs) Add(e error) func(context.Context, models.Order) error {
	return func(context.Context, models.Order) error { return e }
}

func (s *orderStorageStubFuncs) Delete(e error) func(context.Context, uuid.UUID) error {
	return func(context.Context, uuid.UUID) error { return e }
}

func (s *orderStorageStubFuncs) Update(e error) func(context.Context, models.Order) error {
	return func(context.Context, models.Order) error { return e }
}
