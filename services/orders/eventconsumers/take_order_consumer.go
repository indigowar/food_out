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

func NewTakeOrderConsumer(
	logger *slog.Logger,
	command *commands.TakeOrder,
	brokers []string,
	group string,
	topic string,
	partition int,
) (Consumer, error) {
	id := fmt.Sprintf("broker %d of %s for %s", rand.Intn(100), group, topic)

	return newConsumer(
		id,
		takeOrderUnpacker,
		takeOrderValidator,
		takeOrderExecutioner(command),
		logger,
		brokers,
		group,
		topic,
		partition,
	)
}

type takeOrderData struct {
	Order     uuid.UUID
	Courier   uuid.UUID
	Timestmap time.Time
}

func takeOrderUnpacker(msg kafka.Message) (*events.CourierTookOrder, error) {
	var event events.CourierTookOrder
	if err := proto.Unmarshal(msg.Value, &event); err != nil {
		return nil, err
	}
	return &event, nil
}

func takeOrderValidator(event *events.CourierTookOrder) (takeOrderData, error) {
	order, err := uuid.Parse(event.Order)
	if err != nil {
		return takeOrderData{}, errors.New("order: is invalid")
	}

	courier, err := uuid.Parse(event.Courier)
	if err != nil {
		return takeOrderData{}, errors.New("courier: is invalid")
	}

	if event.Timestamp.AsTime().After(time.Now()) {
		return takeOrderData{}, errors.New("taken_at: is invalid")
	}

	return takeOrderData{
		Order:     order,
		Courier:   courier,
		Timestmap: event.Timestamp.AsTime(),
	}, nil
}

func takeOrderExecutioner(cmd *commands.TakeOrder) executioner[takeOrderData] {
	return func(ctx context.Context, data takeOrderData) error {
		return cmd.TakeOrder(ctx, data.Order, data.Courier, data.Timestmap)
	}
}
