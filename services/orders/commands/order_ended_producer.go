package commands

import (
	"context"

	"github.com/indigowar/services/orders/models"
)

//go:generate go run github.com/matryer/moq -out order_ended_producer_moq.go . OrderEndedProducer

type OrderEndedProducer interface {
	Produce(ctx context.Context, order models.Order) error
}
