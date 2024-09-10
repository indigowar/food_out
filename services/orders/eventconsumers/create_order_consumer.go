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
	"github.com/indigowar/services/orders/models"
)

type createOrderData struct {
	Customer        uuid.UUID
	CustomerAddress string
	Restaurant      uuid.UUID
	Products        []models.Product
	CreatedAt       time.Time
}

// NewCreateOrderConsumer creates a Consumer for creating new orders,
// uses [[consumer]]
func NewCreateOrderConsumer(
	logger *slog.Logger,
	command *commands.CreateOrder,
	brokers []string,
	group string,
	topic string,
	partition int,
) (Consumer, error) {
	id := fmt.Sprintf("broker %d of %s for %s", rand.Intn(100), group, topic)

	return newConsumer(
		id,
		createOrderUnpacker,
		createOrderValidator,
		createOrderExecutor(command),
		logger,
		brokers,
		group,
		topic,
		partition,
	)
}

// createOrderValidator validates CreateOrder using [[validator]] interface
func createOrderValidator(request *events.CreateOrderRequest) (createOrderData, error) {
	customer, err := uuid.Parse(request.Customer)
	if err != nil {
		return createOrderData{}, errors.New("customer: is invalid")
	}

	restaurant, err := uuid.Parse(request.Restaurant)
	if err != nil {
		return createOrderData{}, errors.New("restaurant: is invalid")
	}

	products := make([]models.Product, len(request.Products))

	for i, v := range request.Products {
		p, err := productToModel(v)
		if err != nil {
			return createOrderData{}, fmt.Errorf("products: %w", err)
		}
		products[i] = p
	}

	if request.Timestamp.AsTime().After(time.Now()) {
		return createOrderData{}, errors.New("createdAt: timestamp is invalid")
	}

	return createOrderData{
		Customer:        customer,
		CustomerAddress: request.CustomerAddress,
		Restaurant:      restaurant,
		Products:        products,
		CreatedAt:       request.Timestamp.AsTime(),
	}, nil
}

// createOrderUnpacker unpacks the kafka.Message into [[events.CreateOrderRequest]]
func createOrderUnpacker(msg kafka.Message) (*events.CreateOrderRequest, error) {
	var value events.CreateOrderRequest
	if err := proto.Unmarshal(msg.Value, &value); err != nil {
		return nil, err
	}
	return &value, nil
}

// createOrderExecutor - executes command with [[createOrderData]]
func createOrderExecutor(cmd *commands.CreateOrder) executioner[createOrderData] {
	return func(ctx context.Context, data createOrderData) error {
		return cmd.CreateOrder(
			ctx,
			data.Customer,
			data.CustomerAddress,
			data.Restaurant,
			data.Products,
			data.CreatedAt,
		)
	}
}

func productToModel(product *events.Product) (models.Product, error) {
	id, err := uuid.Parse(product.Id)
	if err != nil {
		return models.Product{}, errors.New("id: is invalid")
	}

	restaurant, err := uuid.Parse(product.Restaurant)
	if err != nil {
		return models.Product{}, errors.New("restaurant: is invalid")
	}

	if product.Price < 0 {
		return models.Product{}, errors.New("price: is not positive")
	}

	categories := make([]uuid.UUID, len(product.Categories))
	for i, v := range product.Categories {
		id, err := uuid.Parse(v)
		if err != nil {
			return models.Product{}, fmt.Errorf("categories: %d is invalid", i)
		}
		categories[i] = id
	}

	return models.Product{
		ID:          id,
		Restaurant:  restaurant,
		Name:        product.Name,
		Picture:     product.Picture,
		Price:       float64(product.Price),
		Description: product.Description,
	}, nil
}
