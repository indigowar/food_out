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

type cancelOrderTestSuite struct {
	suite.Suite

	orderEndedProducer *OrderEndedProducerMock
	storage            *OrderStorageMock
	storageStubs       orderStorageStubFuncs
	cmd                *CancelOrder
}

func (suite *cancelOrderTestSuite) SetupTest() {
	suite.orderEndedProducer = &OrderEndedProducerMock{}
	suite.storage = &OrderStorageMock{}

	suite.cmd = NewCancelOrder(
		slog.New(slog.NewTextHandler(os.Stdout, nil)),
		suite.storage,
		suite.orderEndedProducer,
	)
}

func (suite *cancelOrderTestSuite) TestCancelOrder_InvalidTimestamp() {
	orderId := uuid.New()
	canceller := uuid.New()
	timestamp := time.Now().Add(15 * time.Second)

	err := suite.cmd.CancelOrder(context.Background(), orderId, canceller, timestamp)

	suite.NotNil(err)
	suite.ErrorIs(err, ErrInvalidRequest)
}

func (suite *cancelOrderTestSuite) TestCancelOrder_OrderNotFoundInStorage() {
	suite.storage.GetFunc = suite.storageStubs.Get(models.Order{}, &StorageError{
		ErrorType: StorageErrorTypeNotFound,
		Object:    "order",
		Field:     "id",
		Message:   errors.New("not rows"),
	})

	err := suite.cmd.CancelOrder(context.Background(), uuid.New(), uuid.New(), time.Now())
	suite.NotNil(err)
	suite.ErrorIs(err, ErrOrderNotFound)
}

func TestCancelOrder(t *testing.T) {
	suite.Run(t, new(cancelOrderTestSuite))
}
