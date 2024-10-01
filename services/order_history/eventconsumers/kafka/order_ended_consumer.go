package kafka

import (
	"context"

	"github.com/indigowar/food_out/services/order_history/config"
	"github.com/indigowar/food_out/services/order_history/service"
)

type OrderEndedConsumer struct{}

func (consumer *OrderEndedConsumer) Consume(_ context.Context) error {
	// TODO: Implement OrderEndedConsumer.Consume.
	panic("unimplemented")
}

func NewOrderEndedConsumer(service *service.OrderHistory, config *config.Kafka) *OrderEndedConsumer {
	// TODO: Implement NewOrderEndedConsumer.
	panic("unimplemented")
}
