package kafka

import (
	"context"
	"log/slog"
	"sync"

	"github.com/segmentio/kafka-go"
	"google.golang.org/protobuf/proto"

	"github.com/indigowar/food_out/services/order_history/config"
	"github.com/indigowar/food_out/services/order_history/eventconsumers/kafka/events"
	"github.com/indigowar/food_out/services/order_history/service"
)

//go:generate mkdir -p ./events
//go:generate protoc --go_out=./events --go_opt=paths=source_relative ./events.proto

type OrderEndedConsumer struct {
	reader    *kafka.Reader
	waitGroup sync.WaitGroup

	logger  *slog.Logger
	service *service.OrderHistory
}

func (consumer *OrderEndedConsumer) Consume(ctx context.Context) error {
	for {
		msg, err := consumer.reader.ReadMessage(ctx)
		if err != nil {
			if err == context.Canceled {
				return nil
			}
			consumer.logger.Error("Failed to ReadMessage from Kafka", "err", err)
			continue
		}

		var event events.OrderEnded
		if err := proto.Unmarshal(msg.Value, &event); err != nil {
			consumer.logger.Error("Mailformed Kafka message", "stage", "unmarshal", "err", err)
			continue
		}

		order, err := orderEndedToOrder(&event)
		if err != nil {
			consumer.logger.Error("Mailformed Kafka message", "stage", "convertion", "err", err)
		}

		if err := consumer.service.AddOrder(ctx, order); err != nil {
			if err == context.Canceled {
				return nil
			}

			// TODO: handle possible service errors
		}
	}
}

func NewOrderEndedConsumer(
	logger *slog.Logger,
	service *service.OrderHistory,
	config *config.Kafka,
) *OrderEndedConsumer {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   config.Brokers,
		GroupID:   config.Group,
		Topic:     config.OrderEndedTopicName,
		Partition: 0,
	})

	return &OrderEndedConsumer{
		reader:    reader,
		waitGroup: sync.WaitGroup{},
		logger:    logger,
		service:   service,
	}
}
