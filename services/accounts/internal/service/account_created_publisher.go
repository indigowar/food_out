package service

import (
	"context"

	"github.com/indigowar/food_out/services/accounts/internal/domain"
)

//go:generate go run github.com/matryer/moq -out account_created_publisher_moq_test.go . AccountCreatedPublisher

type AccountCreatedPublisher interface {
	PublishAccountCreated(ctx context.Context, account domain.Account) error
}
