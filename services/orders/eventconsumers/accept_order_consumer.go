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

func NewAcceptOrderConsumer(
	logger *slog.Logger,
	command *commands.AcceptOrder,
	brokers []string,
	group string,
	topic string,
	partition int,
) (Consumer, error) {
	id := fmt.Sprintf("broker %d of %s for %s", rand.Intn(100), group, topic)

	return newConsumer(
		id,
		acceptOrderUnpacker,
		acceptOrderValidator,
		acceptOrderExecutioner(command),
		logger,
		brokers,
		group,
		topic,
		partition,
	)
}

type acceptOrderData struct {
	Order     uuid.UUID
	Manager   uuid.UUID
	Timestmap time.Time
}

func acceptOrderValidator(request *events.RestaurantAcceptedOrder) (acceptOrderData, error) {
	order, err := uuid.Parse(request.Order)
	if err != nil {
		return acceptOrderData{}, errors.New("order: is invalid")
	}

	manager, err := uuid.Parse(request.Manager)
	if err != nil {
		return acceptOrderData{}, errors.New("manager: is invalid")
	}

	if request.Timestamp.AsTime().After(time.Now()) {
		return acceptOrderData{}, errors.New("acceptedAt: is invalid(haven't happened yet)")
	}

	return acceptOrderData{
		Order:     order,
		Manager:   manager,
		Timestmap: request.Timestamp.AsTime(),
	}, nil
}

func acceptOrderUnpacker(msg kafka.Message) (*events.RestaurantAcceptedOrder, error) {
	var value events.RestaurantAcceptedOrder
	if err := proto.Unmarshal(msg.Value, &value); err != nil {
		return nil, err
	}
	return &value, nil
}

func acceptOrderExecutioner(cmd *commands.AcceptOrder) executioner[acceptOrderData] {
	return func(ctx context.Context, data acceptOrderData) error {
		return cmd.AcceptOrder(ctx, data.Order, data.Manager, data.Timestmap)
	}
}
