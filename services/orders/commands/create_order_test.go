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

type createOrderTestSuite struct {
	suite.Suite

	stub    orderStorageStubFuncs
	storage *OrderStorageMock
	cmd     *CreateOrder

	restaurantID uuid.UUID
	products     []models.Product
}

func (suite *createOrderTestSuite) SetupSuite() {
	suite.restaurantID = uuid.New()
	suite.products = make([]models.Product, 10)
	for i := 0; i < 10; i++ {
		suite.products[i] = models.Product{
			ID:          uuid.New(),
			Restaurant:  suite.restaurantID,
			Name:        "Product Name X",
			Picture:     "ksfjdldval;fd",
			Price:       40.0,
			Description: "cool product",
		}
	}
}

func (suite *createOrderTestSuite) SetupTest() {
	suite.storage = &OrderStorageMock{}
	suite.cmd = NewCreateOrder(
		slog.New(slog.NewTextHandler(os.Stdout, nil)),
		suite.storage,
	)
}

func (suite *createOrderTestSuite) TestCreateOrder_InvalidProducts() {
	restaurant := uuid.New()
	products := make([]models.Product, 10)
	for i := 0; i < 10; i++ {
		products[i] = models.Product{
			ID:          uuid.New(),
			Restaurant:  uuid.New(), // Product's restaurant is different than restaurant variable
			Name:        "Product Name X",
			Picture:     "ksfjdldval;fd",
			Price:       40.0,
			Description: "cool product",
		}
	}

	timestamp := time.Now().Add(-15 * time.Hour)
	err := suite.cmd.CreateOrder(
		context.TODO(),
		uuid.New(),
		"customer address",
		restaurant,
		products,
		timestamp,
	)

	suite.NotNil(err)
	suite.ErrorIs(err, ErrInvalidRequest)
}

func (suite *createOrderTestSuite) TestCreateOrder_InvalidTime() {
	timestamp := time.Now().Add(1 * time.Hour)
	err := suite.cmd.CreateOrder(
		context.TODO(),
		uuid.New(),
		"customer address",
		suite.restaurantID,
		suite.products,
		timestamp,
	)

	suite.NotNil(err)
	suite.ErrorIs(err, ErrInvalidRequest)
}

func (suite *createOrderTestSuite) TestCreateOrder_AlreadyInStorage() {
	suite.storage.AddFunc = suite.stub.Add(&StorageError{
		ErrorType: StorageErrorTypeAlreadyExists,
		Object:    "order",
		Field:     "customer, restaurant, timestamp",
		Message:   errors.New("required unique set already exists"),
	})

	err := suite.cmd.CreateOrder(
		context.TODO(),
		uuid.New(),
		"customer address",
		suite.restaurantID,
		suite.products,
		time.Now().Add(-1*time.Minute),
	)

	suite.NotNil(err)
	suite.ErrorIs(err, ErrOrderDuplicated)
}

func (suite *createOrderTestSuite) TestCreateOrder_StorageUnexpectedError() {
	suite.storage.AddFunc = suite.stub.Add(errors.New("unexpected bug error"))

	err := suite.cmd.CreateOrder(
		context.TODO(),
		uuid.New(),
		"customer address",
		suite.restaurantID,
		suite.products,
		time.Now().Add(-1*time.Minute),
	)

	suite.NotNil(err)
	suite.ErrorIs(err, ErrUnexpected)
}

func TestCreateOrder(t *testing.T) {
	suite.Run(t, new(createOrderTestSuite))
}
