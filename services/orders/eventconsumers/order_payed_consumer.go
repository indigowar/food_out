package eventconsumers

import (
	"context"
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

func NewOrderPayedConsumer(
	logger *slog.Logger,
	command *commands.PayForOrder,
	brokers []string,
	group string,
	topic string,
	partition int,
) (Consumer, error) {
	id := fmt.Sprintf("broker %d of %s for %s", rand.Intn(100), group, topic)

	return newConsumer(
		id,
		orderPayedUnpacker,
		orderPayedValidator,
		orderPayedExecutioner(command),
		logger,
		brokers,
		group,
		topic,
		partition,
	)
}

type orderPayedData struct {
	Order       uuid.UUID
	Transaction string
	PayedAt     time.Time
}

func orderPayedUnpacker(msg kafka.Message) (*events.OrderHasBeenPayed, error) {
	var event events.OrderHasBeenPayed
	if err := proto.Unmarshal(msg.Value, &event); err != nil {
		return nil, err
	}
	return &event, nil
}

func orderPayedValidator(event *events.OrderHasBeenPayed) (orderPayedData, error) {
	panic("not implemented")
}

func orderPayedExecutioner(cmd *commands.PayForOrder) executioner[orderPayedData] {
	return func(ctx context.Context, data orderPayedData) error {
		return cmd.PayForOrder(ctx, data.Order, data.Transaction, data.PayedAt)
	}
}
