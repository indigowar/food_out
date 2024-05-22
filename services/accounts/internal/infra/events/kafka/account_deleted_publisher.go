package kafka

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/segmentio/kafka-go"
	"google.golang.org/protobuf/proto"

	"github.com/indigowar/food_out/services/accounts/internal/service"
)

type AccountDeletedPublisher struct {
	writer *kafka.Writer
}

var _ service.AccountDeletedPublisher = &AccountDeletedPublisher{}

// PublishAccountDeleted implements service.AccountDeletedPublisher.
func (a *AccountDeletedPublisher) PublishAccountDeleted(ctx context.Context, id uuid.UUID) error {
	data, err := proto.Marshal(&AccountDeleted{
		Id: id.String(),
	})
	if err != nil {
		return err
	}

	return a.writer.WriteMessages(ctx, kafka.Message{Value: data})
}

func NewAccountDeletedPublisher(host string, port int, topic string, partition int) *AccountDeletedPublisher {
	return &AccountDeletedPublisher{
		writer: &kafka.Writer{
			Addr:                   kafka.TCP(fmt.Sprintf("%s:%d", host, port)),
			Topic:                  topic,
			AllowAutoTopicCreation: true,
		},
	}
}
