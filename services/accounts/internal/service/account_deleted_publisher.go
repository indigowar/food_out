package service

import (
	"context"

	"github.com/google/uuid"
)

//go:generate go run github.com/matryer/moq -out account_deleted_publisher_moq_test.go . AccountDeletedPublisher

type AccountDeletedPublisher interface {
	PublishAccountDeleted(ctx context.Context, id uuid.UUID) error
}
