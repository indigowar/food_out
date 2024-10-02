package kafka

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/indigowar/food_out/services/order_history/eventconsumers/kafka/events"
	"github.com/indigowar/food_out/services/order_history/service"
)

func orderEndedToOrder(event *events.OrderEnded) (service.Order, error) {
	id, err := parseID(event.Id)
	if err != nil {
		return service.Order{}, fmt.Errorf("Order ID is invalid: %w", err)
	}

	restaurant, err := parseID(event.Restaurant)
	if err != nil {
		return service.Order{}, fmt.Errorf("Restaurant ID is invalid: %w", err)
	}

	customer, err := parseID(event.CustomerId)
	if err != nil {
		return service.Order{}, fmt.Errorf("Customer ID is invalid: %w", err)
	}

	order := service.Order{
		ID:         id,
		Restaurant: restaurant,
		// Products:   []service.Product{},
		Customer: struct {
			ID      uuid.UUID
			Address string
		}{
			ID:      customer,
			Address: event.CustomerAddress,
		},
		CreatedAt:         event.CreatedAt.AsTime(),
		CookingStartedAt:  convertOptionalTimestamp(event.CookingStartedAt),
		DeliveryStartedAt: convertOptionalTimestamp(event.DeliveryStartedAt),
		DeliveryCompleted: convertOptionalTimestamp(event.DeliveryCompletedAt),
	}

	if event.Acceptance != nil {
		id, err := parseID(event.Acceptance.Manager)
		if err != nil {
			return service.Order{}, fmt.Errorf("Order Acceptance Manager ID is invalid: %w", err)
		}
		order.Acceptance = &struct {
			Manager    uuid.UUID
			AcceptedAt time.Time
		}{
			Manager:    id,
			AcceptedAt: event.Acceptance.AcceptedAt.AsTime(),
		}
	}

	if event.Courier != nil {
		id, err := parseID(event.Courier.Courier)
		if err != nil {
			return service.Order{}, fmt.Errorf("Order Courier ID is invalid: %w", err)
		}
		order.Courier = &struct {
			ID      uuid.UUID
			TakenAt time.Time
		}{
			ID:      id,
			TakenAt: event.Courier.TakenAt.AsTime(),
		}
	}

	if event.Payment != nil {
		order.Payment = &struct {
			Transaction string
			PayedAt     time.Time
		}{
			Transaction: event.Payment.Transaction,
			PayedAt:     event.Payment.PayedAt.AsTime(),
		}
	}

	if event.Cancellation != nil {
		id, err := parseID(event.Cancellation.Canceller)
		if err != nil {
			return service.Order{}, fmt.Errorf("Order Canceller ID is invalid: %w", err)
		}
		order.Cancellation = &struct {
			Canceller   uuid.UUID
			CancelledAt time.Time
		}{
			Canceller:   id,
			CancelledAt: event.Cancellation.CancelledAt.AsTime(),
		}
	}

	order.Products = make([]service.Product, len(event.Products))
	for i, v := range event.Products {
		order.Products[i], err = parseProduct(v)
		if err != nil {
			return service.Order{}, fmt.Errorf("Product is invalid:%w", err)
		}
	}

	return order, nil
}

func parseProduct(value *events.Product) (service.Product, error) {
	id, err := parseID(value.Id)
	if err != nil {
		return service.Product{}, fmt.Errorf("ID is invalid: %w", err)
	}

	restaurant, err := parseID(value.Restaurant)
	if err != nil {
		return service.Product{}, fmt.Errorf("Restaurant is invalid: %w", err)
	}

	original, err := parseID(value.OriginalId)
	if err != nil {
		return service.Product{}, fmt.Errorf("Original is invalid: %w", err)
	}

	return service.Product{
		ID:          id,
		Original:    original,
		Restaurant:  restaurant,
		Name:        value.Name,
		Picture:     value.Picture,
		Price:       float64(value.Price),
		Description: value.Description,
	}, nil
}

func parseID(value string) (uuid.UUID, error) {
	result, err := uuid.Parse(value)
	if err != nil {
		return uuid.UUID{}, fmt.Errorf("ID is invalid: %w", err)
	}
	return result, nil
}

func convertOptionalTimestamp(value *timestamppb.Timestamp) *time.Time {
	if value != nil {
		v := value.AsTime()
		return &v
	}
	return nil
}
