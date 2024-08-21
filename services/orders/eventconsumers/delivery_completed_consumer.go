package eventconsumers

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"math/rand"
	"time"

	"github.com/google/uuid"
	"github.com/segmentio/kafka-go"
	"google.golang.org/protobuf/proto"

	"github.com/indigowar/services/orders/commands"
	"github.com/indigowar/services/orders/eventconsumers/events"
)

func NewDeliveryCompletedConsumer(
	logger *slog.Logger,
	command *commands.MarkDeliveryCompleted,
	brokers []string,
	group string,
	topic string,
	partition int,
) (Consumer, error) {
	id := fmt.Sprintf("broker %d of %s for %s", rand.Intn(100), group, topic)

	return newConsumer(
		id,
		deliveryCompletedUnpacker,
		deliveryCompletedValidator,
		deliveryCompletedExecutioner(command),
		logger,
		brokers,
		group,
		topic,
		partition,
	)
}

type deliveryCompletedData struct {
	Order uuid.UUID
	Time  time.Time
}

func deliveryCompletedUnpacker(msg kafka.Message) (*events.DeliveryCompleted, error) {
	var event events.DeliveryCompleted
	if err := proto.Unmarshal(msg.Value, &event); err != nil {
		return nil, err
	}
	return &event, nil
}

func deliveryCompletedValidator(event *events.DeliveryCompleted) (deliveryCompletedData, error) {
	order, err := uuid.Parse(event.Order)
	if err != nil {
		return deliveryCompletedData{}, errors.New("order: is invalid")
	}

	if event.Timestamp.AsTime().After(time.Now()) {
		return deliveryCompletedData{}, errors.New("deliveryCompletedAt: is invalid")
	}

	return deliveryCompletedData{
		Order: order,
		Time:  event.Timestamp.AsTime(),
	}, nil
}

func deliveryCompletedExecutioner(cmd *commands.MarkDeliveryCompleted) executioner[deliveryCompletedData] {
	return func(ctx context.Context, data deliveryCompletedData) error {
		return cmd.MarkDeliveryCompleted(ctx, data.Order, data.Time)
	}
}
