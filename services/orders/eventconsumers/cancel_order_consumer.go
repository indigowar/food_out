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

func NewCancelOrderConsumer(
	logger *slog.Logger,
	cmd *commands.CancelOrder,
	brokers []string,
	group string,
	topic string,
	partititon int,
) (Consumer, error) {
	id := fmt.Sprintf("broker %d of %s for %s", rand.Intn(100), group, topic)

	return newConsumer(
		id,
		cancelOrderUnpacker,
		cancelOrderValidator,
		cancelOrderExecutioner(cmd),
		logger,
		brokers,
		group,
		topic,
		partititon,
	)
}

type cancelOrderData struct {
	Order     uuid.UUID
	Canceller uuid.UUID
	Timestamp time.Time
}

func cancelOrderValidator(request *events.CancellOrder) (cancelOrderData, error) {
	order, err := uuid.Parse(request.Order)
	if err != nil {
		return cancelOrderData{}, errors.New("cancel order: order id is invalid")
	}

	canceller, err := uuid.Parse(request.Canceller)
	if err != nil {
		return cancelOrderData{}, errors.New("cancel order: canceller id is invalid")
	}

	return cancelOrderData{
		Order:     order,
		Canceller: canceller,
		Timestamp: request.Timestamp.AsTime(),
	}, nil
}

func cancelOrderUnpacker(msg kafka.Message) (*events.CancellOrder, error) {
	var value events.CancellOrder
	if err := proto.Unmarshal(msg.Value, &value); err != nil {
		return nil, err
	}
	return &value, nil
}

func cancelOrderExecutioner(cmd *commands.CancelOrder) executioner[cancelOrderData] {
	return func(ctx context.Context, data cancelOrderData) error {
		return cmd.CancelOrder(ctx, data.Order, data.Canceller, data.Timestamp)
	}
}
