package kafka

import (
	"context"
	"fmt"

	"github.com/segmentio/kafka-go"

	"github.com/indigowar/food_out/services/accounts/internal/domain"
	"github.com/indigowar/food_out/services/accounts/internal/service"
)

type AccountCreatedPublisher struct {
	writer *kafka.Writer
}

var _ service.AccountCreatedPublisher = &AccountCreatedPublisher{}

// PublishAccountCreated implements service.AccountCreatedPublisher.
func (a *AccountCreatedPublisher) PublishAccountCreated(ctx context.Context, account *domain.Account) error {
	panic("unimplemented")
}

func NewAccountCreatedPublisher(host string, port int, topic string, partition int) *AccountCreatedPublisher {
	return &AccountCreatedPublisher{
		writer: kafka.NewWriter(kafka.WriterConfig{
			Brokers: []string{fmt.Sprintf("%s:%d", host, port)},
			Topic:   topic,
		}),
	}
}
