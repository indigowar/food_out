package commands

import (
	"context"
	"errors"
	"log/slog"
	"os"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"

	"github.com/indigowar/services/orders/models"
)

type acceptOrderTestSuite struct {
	suite.Suite

	storage *OrderStorageMock
	stub    orderStorageStubFuncs
	cmd     *AcceptOrder

	order models.Order
}

func (suite *acceptOrderTestSuite) SetupTest() {
	suite.stub = orderStorageStubFuncs{}

	suite.storage = &OrderStorageMock{}

	suite.cmd = NewAcceptOrder(
		slog.New(slog.NewTextHandler(os.Stdout, nil)),
		suite.storage,
	)

	restaurant := uuid.New()

	products := make([]models.Product, 10)
	for i := 0; i < 10; i++ {
		products[i] = models.Product{
			ID:          uuid.New(),
			Restaurant:  restaurant,
			Name:        "Product Name X",
			Picture:     "ksfjdldval;fd",
			Price:       40.0,
			Description: "cool product",
			Categories:  []uuid.UUID{uuid.New(), uuid.New()},
		}
	}

	suite.order = models.Order{
		ID:         uuid.New(),
		Restaurant: restaurant,
		Products:   products,
		Customer: struct {
			ID      uuid.UUID
			Address string
		}{
			ID:      uuid.New(),
			Address: "customer address",
		},
		CreatedAt: time.Now().Add(-2 * time.Hour),
	}
}

func (suite *acceptOrderTestSuite) TestAcceptOrder_OrderDoesNotExist() {
	suite.storage.GetFunc = suite.stub.Get(models.Order{}, &StorageError{
		ErrorType: StorageErrorTypeNotFound,
		Object:    "order",
		Field:     "id",
		Message:   errors.New("order not found with given id"),
	})
	suite.storage.UpdateFunc = suite.stub.Update(&StorageError{
		ErrorType: StorageErrorTypeNotFound,
		Object:    "order",
		Field:     "id",
		Message:   errors.New("order not found with given id"),
	})

	err := suite.cmd.AcceptOrder(
		context.TODO(),
		suite.order.ID,
		uuid.New(),
		time.Now().Add(-1*time.Minute),
	)

	suite.NotNil(err)
	suite.ErrorIs(err, ErrOrderNotFound)

	suite.storage.GetFunc = suite.stub.Get(suite.order, nil)

	err = suite.cmd.AcceptOrder(
		context.TODO(),
		suite.order.ID,
		uuid.New(),
		time.Now().Add(-1*time.Minute),
	)

	suite.NotNil(err)
	suite.ErrorIs(err, ErrOrderNotFound)
}

func (suite *acceptOrderTestSuite) TestAcceptOrder_InvalidTimestamp() {
	err := suite.cmd.AcceptOrder(
		context.TODO(),
		suite.order.ID,
		uuid.New(),
		time.Now().Add(2*time.Minute),
	)

	suite.NotNil(err)
	suite.ErrorIs(err, ErrInvalidRequest)
}

func (suite *acceptOrderTestSuite) TestAcceptOrder_OrderAlreadyAccepted() {
	suite.storage.GetFunc = suite.stub.Get(suite.order, nil)
	suite.storage.UpdateFunc = suite.stub.Update(nil)

	suite.order.Acceptance = &struct {
		Manager    uuid.UUID
		AcceptedAt time.Time
	}{
		Manager:    uuid.New(),
		AcceptedAt: time.Now().Add(-1 * time.Minute),
	}

	err := suite.cmd.AcceptOrder(
		context.TODO(),
		suite.order.ID,
		uuid.New(),
		time.Now().Add(2*time.Minute),
	)

	suite.NotNil(err)
	suite.ErrorIs(err, ErrInvalidRequest)
}

func TestAcceptOrder(t *testing.T) {
	suite.Run(t, new(acceptOrderTestSuite))
}
