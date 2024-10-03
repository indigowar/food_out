package service

import (
	"context"
)

//go:generate go run github.com/matryer/moq -out storage_moq_test.go . Storage

// TODO: Define the Storage interface.

type Storage interface {
	Add(ctx context.Context, order Order) error
}
