package eventproducers

import (
	"context"
	"fmt"

	"github.com/segmentio/kafka-go"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/indigowar/services/orders/commands"
	"github.com/indigowar/services/orders/eventproducers/events"
	"github.com/indigowar/services/orders/models"
)

//go:generate mkdir -p ./events
//go:generate protoc --go_out=./events --go_opt=paths=source_relative ./produced_events.proto

type OrderEndedProducer struct {
	writer *kafka.Writer
}

var _ commands.OrderEndedProducer = &OrderEndedProducer{}

// Produce implements commands.OrderEndedProducer.
func (producer *OrderEndedProducer) Produce(ctx context.Context, order models.Order) error {
	message := orderToEvent(order)
	marshalized, err := proto.Marshal(message)
	if err != nil {
		return fmt.Errorf("orderEndedProducer: failed to serialize: %w", err)
	}

	err = producer.writer.WriteMessages(ctx, kafka.Message{Key: nil, Value: marshalized})
	if err != nil {
		return fmt.Errorf("orderEndedProducer: failed to produce an event: %w", err)
	}

	return nil
}

func NewOrderEndedProducer(topic string, addresses []string) *OrderEndedProducer {
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  addresses,
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	})

	return &OrderEndedProducer{
		writer: writer,
	}
}

func orderToEvent(order models.Order) *events.OrderEnded {
	products := make([]*events.Product, 0, len(order.Products))
	for _, v := range order.Products {
		products = append(products, productToEventProduct(v))
	}

	message := events.OrderEnded{
		Id:              order.ID.String(),
		CustomerId:      order.Customer.ID.String(),
		CustomerAddress: order.Customer.Address,
		CreatedAt:       timestamppb.New(order.CreatedAt),
		Restaurant:      order.Restaurant.String(),
		Products:        products,
	}

	if order.Acceptance != nil {
		message.Acceptance = &events.OrderEnded_Acceptance{
			Manager:    order.Acceptance.Manager.String(),
			AcceptedAt: timestamppb.New(order.Acceptance.AcceptedAt),
		}
	}

	if order.Courier != nil {
		message.Courier = &events.OrderEnded_Courier{
			Courier: order.Courier.ID.String(),
			TakenAt: timestamppb.New(order.Courier.TakenAt),
		}
	}

	if order.Payment != nil {
		message.Payment = &events.OrderEnded_Payment{
			Transaction: order.Payment.Transaction,
			PayedAt:     timestamppb.New(order.Payment.PayedAt),
		}
	}

	if order.Cancellation != nil {
		message.Cancellation = &events.OrderEnded_Cancellation{
			Canceller:   order.Cancellation.Canceller.String(),
			CancelledAt: timestamppb.New(order.Cancellation.CancelledAt),
		}
	}

	if order.CookingStartedAt != nil {
		message.CookingStartedAt = timestamppb.New(*order.CookingStartedAt)
	}

	if order.DeliveryStartedAt != nil {
		message.DeliveryStartedAt = timestamppb.New(*order.DeliveryStartedAt)
	}

	if order.DeliveryCompleted != nil {
		message.DeliveryCompletedAt = timestamppb.New(*order.DeliveryCompleted)
	}

	return &message
}

func productToEventProduct(product models.Product) *events.Product {
	return &events.Product{
		Id:          product.ID.String(),
		OriginalId:  product.Original.String(),
		Restaurant:  product.Restaurant.String(),
		Name:        product.Name,
		Picture:     product.Picture,
		Price:       float32(product.Price),
		Description: product.Description,
	}
}
