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

func NewDeliveryStartedConsumer(
	logger *slog.Logger,
	command *commands.MarkDeliveryStarted,
	brokers []string,
	group string,
	topic string,
	partition int,
) (Consumer, error) {
	id := fmt.Sprintf("broker %d of %s for %s", rand.Intn(100), group, topic)

	return newConsumer(
		id,
		deliveryStartedUnpacker,
		deliveryStartedValidator,
		deliveryStartedExecutioner(command),
		logger,
		brokers,
		group,
		topic,
		partition,
	)
}

type deliveryStartedData struct {
	Order uuid.UUID
	Time  time.Time
}

func deliveryStartedUnpacker(msg kafka.Message) (*events.DeliveryStarted, error) {
	var event events.DeliveryStarted
	if err := proto.Unmarshal(msg.Value, &event); err != nil {
		return nil, err
	}
	return &event, nil
}

func deliveryStartedValidator(event *events.DeliveryStarted) (deliveryStartedData, error) {
	order, err := uuid.Parse(event.Order)
	if err != nil {
		return deliveryStartedData{}, errors.New("order: is invalid")
	}

	if event.Timestamp.AsTime().After(time.Now()) {
		return deliveryStartedData{}, errors.New("deliveryStartedAt: is invalid")
	}

	return deliveryStartedData{
		Order: order,
		Time:  event.Timestamp.AsTime(),
	}, nil
}

func deliveryStartedExecutioner(cmd *commands.MarkDeliveryStarted) executioner[deliveryStartedData] {
	return func(ctx context.Context, data deliveryStartedData) error {
		return cmd.MarkDeliveryStarted(ctx, data.Order, data.Time)
	}
}
