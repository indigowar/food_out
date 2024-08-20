package commands

import (
	"context"

	"github.com/indigowar/services/orders/models"
)

type OrderEndedProducer interface {
	Produce(ctx context.Context, order models.Order) error
}
