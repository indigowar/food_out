package commands

import (
	"context"
	"errors"
	"log/slog"
	"os"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
)

type addRestaurantTestSuite struct {
	suite.Suite

	restaurants RestaurantStoragePortMock

	cmd *AddRestaurantCommand
}

func (suite *addRestaurantTestSuite) SetupTest() {
	suite.cmd = NewAddRestaurantCommand(&suite.restaurants, slog.New(slog.NewTextHandler(os.Stdout, nil)))
}

func (suite *addRestaurantTestSuite) TestAddRestaurant_InvalidUUID() {
	id := uuid.UUID{}

	result, err := suite.cmd.AddRestaurant(context.Background(), id)

	suite.NotNil(err, "AddRestaurant should return an error, when input data is default uuid(uuid.UUID{})")
	suite.Equal(result, uuid.UUID{}, "AddRestaurant should return an empty id, when an error occurred")
	suite.ErrorIs(err, ErrInvalidData, "AddRestaurant should return an ErrInvalidData")
}

func (suite *addRestaurantTestSuite) TestAddRestaurant_AlreadyExists() {
	suite.restaurants.AddRestaurantFunc = mockAddRestaurant(ErrAlreadyExists)

	id := uuid.New()

	result, err := suite.cmd.AddRestaurant(context.Background(), id)

	suite.NotNil(err, "AddRestaurant should return an error, when restaurant is already in storage")
	suite.Equal(result, uuid.UUID{}, "AddRestaurant should return an empty id, when an error occurred")
	suite.ErrorIs(err, ErrRestaurantAlreadyExists, "AddRestaurant should return an ErrRestaurantAlreadyExists")
}

func (suite *addRestaurantTestSuite) TestAddRestaurant_InternalError() {
	suite.restaurants.AddRestaurantFunc = mockAddRestaurant(errors.New("unexpected error"))

	id := uuid.New()

	result, err := suite.cmd.AddRestaurant(context.Background(), id)

	suite.NotNil(err, "AddRestaurant should return an error, when restaurant is already in storage")
	suite.Equal(result, uuid.UUID{}, "AddRestaurant should return an empty id, when an error occurred")
	suite.ErrorIs(err, ErrInternal, "AddRestaurant should return an ErrRestaurantAlreadyExists")
}

func (suite *addRestaurantTestSuite) TestAddRestaurant_Valid() {
	suite.restaurants.AddRestaurantFunc = mockAddRestaurant(nil)

	id := uuid.New()

	result, err := suite.cmd.AddRestaurant(context.Background(), id)

	suite.Nil(err, "AddRestaurant should not return an error, when restaurant is added")
	suite.Equal(id, result, "AddRestaurant should return the same ID as input data")
}

func TestAddRestaurant(t *testing.T) {
	suite.Run(t, new(addRestaurantTestSuite))
}
