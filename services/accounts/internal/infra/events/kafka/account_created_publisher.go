package kafka

import (
	"context"
	"fmt"

	"github.com/segmentio/kafka-go"
	"google.golang.org/protobuf/proto"

	"github.com/indigowar/food_out/services/accounts/internal/domain"
	"github.com/indigowar/food_out/services/accounts/internal/service"
)

type AccountCreatedPublisher struct {
	writer *kafka.Writer
}

var _ service.AccountCreatedPublisher = &AccountCreatedPublisher{}

// PublishAccountCreated implements service.AccountCreatedPublisher.
func (a *AccountCreatedPublisher) PublishAccountCreated(ctx context.Context, account *domain.Account) error {
	data, err := proto.Marshal(&AccountCreated{
		Id:    account.ID().String(),
		Phone: account.Phone(),
	})
	if err != nil {
		return err
	}

	return a.writer.WriteMessages(ctx, kafka.Message{Value: data})
}

func NewAccountCreatedPublisher(host string, port int, topic string, partition int) *AccountCreatedPublisher {
	return &AccountCreatedPublisher{
		writer: kafka.NewWriter(kafka.WriterConfig{
			Brokers: []string{fmt.Sprintf("%s:%d", host, port)},
			Topic:   topic,
		}),
	}
}
