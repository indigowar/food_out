package commands

import (
	"context"
	"log/slog"
	"os"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
)

type deleteDishTestSuite struct {
	suite.Suite

	dishStorage       DishStoragePortMock
	restaurantStorage RestaurantStoragePortMock
	cmd               *DeleteDishCommand
}

func (suite *deleteDishTestSuite) SetupTest() {
	suite.cmd = NewDeleteDishCommand(
		&suite.dishStorage,
		&suite.restaurantStorage,
		slog.New(slog.NewTextHandler(os.Stdout, nil)),
	)
}

func (suite *deleteDishTestSuite) TestDeleteDish_InvalidUUID() {
	args := [][2]uuid.UUID{
		{uuid.New(), uuid.UUID{}},
		{uuid.UUID{}, uuid.New()},
		{uuid.UUID{}, uuid.UUID{}},
	}

	for _, arg := range args {
		err := suite.cmd.DeleteDish(context.Background(), arg[0], arg[1])

		suite.NotNil(err, "DeleteDish should return an error, when input data is default uuid(uuid.UUID{})")
		suite.ErrorIs(err, ErrInvalidData, "DeleteDish should return an ErrInvalidData")
	}
}

func (suite *deleteDishTestSuite) TestDeleteDish_RestaurantNotFound() {
	suite.restaurantStorage.RestaurantExistsFunc = mockRestaurantExists(false, nil)
	suite.dishStorage.DeleteDishFunc = mockDeleteDish(ErrNotFound)

	err := suite.cmd.DeleteDish(context.Background(), uuid.New(), uuid.New())

	suite.NotNil(err, "DeleteDish should return an error, when dish is not found")
	suite.ErrorIs(err, ErrRestaurantNotFound, "DeleteDish should return ErrRestaurantNotFound")
}

func (suite *deleteDishTestSuite) TestDeleteDish_DishNotFound() {
	suite.restaurantStorage.RestaurantExistsFunc = mockRestaurantExists(true, nil)
	suite.dishStorage.DeleteDishFunc = mockDeleteDish(ErrNotFound)

	err := suite.cmd.DeleteDish(context.Background(), uuid.New(), uuid.New())

	suite.NotNil(err, "DeleteDish should return an error, when dish is not found")
	suite.ErrorIs(err, ErrDishNotFound, "DeleteDish should return ErrDishNotFound")
}

func (suite *deleteDishTestSuite) TestDeleteDish_Valid() {
	suite.restaurantStorage.RestaurantExistsFunc = mockRestaurantExists(true, nil)
	suite.dishStorage.DeleteDishFunc = mockDeleteDish(nil)

	err := suite.cmd.DeleteDish(context.Background(), uuid.New(), uuid.New())

	suite.Nil(err, "DeleteDish should not return an error, when dish is not found")
}

func TestDeleteDish(t *testing.T) {
	suite.Run(t, new(deleteDishTestSuite))
}
