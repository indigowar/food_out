package kafka

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/segmentio/kafka-go"

	"github.com/indigowar/food_out/services/accounts/internal/service"
)

type AccountDeletedPublisher struct {
	writer *kafka.Writer
}

var _ service.AccountDeletedPublisher = &AccountDeletedPublisher{}

// PublishAccountDeleted implements service.AccountDeletedPublisher.
func (a *AccountDeletedPublisher) PublishAccountDeleted(ctx context.Context, id uuid.UUID) error {
	panic("unimplemented")
}

func NewAccountDeletedPublisher(host string, port int, topic string, partition int) *AccountDeletedPublisher {
	return &AccountDeletedPublisher{
		writer: kafka.NewWriter(kafka.WriterConfig{
			Brokers: []string{fmt.Sprintf("%s:%d", host, port)},
			Topic:   topic,
		}),
	}
}
