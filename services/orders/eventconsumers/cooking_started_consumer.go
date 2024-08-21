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

func NewCookingStartedConsumer(
	logger *slog.Logger,
	command *commands.MarkCookingStarted,
	brokers []string,
	group string,
	topic string,
	partition int,
) (Consumer, error) {
	id := fmt.Sprintf("broker %d of %s for %s", rand.Intn(100), group, topic)

	return newConsumer(
		id,
		cookingStartedUnpacker,
		cookingStartedValidator,
		cookingStartedExecutioner(command),
		logger,
		brokers,
		group,
		topic,
		partition,
	)
}

type cookingStartedData struct {
	Order uuid.UUID
	Time  time.Time
}

func cookingStartedUnpacker(msg kafka.Message) (*events.CookingStarted, error) {
	var event events.CookingStarted
	if err := proto.Unmarshal(msg.Value, &event); err != nil {
		return nil, err
	}
	return &event, nil
}

func cookingStartedValidator(event *events.CookingStarted) (cookingStartedData, error) {
	order, err := uuid.Parse(event.Order)
	if err != nil {
		return cookingStartedData{}, errors.New("order: is invalid")
	}

	if event.Timestamp.AsTime().After(time.Now()) {
		return cookingStartedData{}, errors.New("cookingStartedAt: is invalid")
	}

	return cookingStartedData{
		Order: order,
		Time:  event.Timestamp.AsTime(),
	}, nil
}

func cookingStartedExecutioner(cmd *commands.MarkCookingStarted) executioner[cookingStartedData] {
	return func(ctx context.Context, data cookingStartedData) error {
		return cmd.MarkCookingStarted(ctx, data.Order, data.Time)
	}
}
