package commands

import (
	"context"
	"log/slog"
	"os"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"

	"github.com/indigowar/food_out/services/menu/domain"
)

type deleteDishTestSuite struct {
	suite.Suite

	dishStorage       DishStoragePortMock
	restaurantStorage RestaurantStoragePortMock
	menuStorage       MenuStoragePortMock

	cmd *DeleteDishCommand

	restaurantId uuid.UUID
	dishId       uuid.UUID
	menu         *domain.Menu
}

func (suite *deleteDishTestSuite) SetupTest() {
	suite.cmd = NewDeleteDishCommand(
		&suite.dishStorage,
		&suite.restaurantStorage,
		&suite.menuStorage,
		slog.New(slog.NewTextHandler(os.Stdout, nil)),
	)

	suite.restaurantId = uuid.New()
	suite.dishId = uuid.New()
	suite.menu, _ = domain.NewMenu("Menu", suite.restaurantId, validUrl)
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

func (suite *deleteDishTestSuite) TestDeleteDish_NoMenuWithGivenDish() {
	suite.restaurantStorage.RestaurantExistsFunc = mockRestaurantExists(true, nil)
	suite.menuStorage.GetMenuByRestaurantFunc = mockGetMenuByRestaurant([]*domain.Menu{suite.menu}, nil)

	err := suite.cmd.DeleteDish(context.Background(), uuid.New(), suite.dishId)

	suite.NotNil(err, "DeleteDish should return an error, when dish is not found")
	suite.ErrorIs(err, ErrDishNotFound, "DeleteDish should return ErrDishNotFound")

}

func (suite *deleteDishTestSuite) TestDeleteDish_DishNotFound() {
	suite.restaurantStorage.RestaurantExistsFunc = mockRestaurantExists(true, nil)
	suite.menuStorage.GetMenuByRestaurantFunc = mockGetMenuByRestaurant([]*domain.Menu{suite.menu}, nil)
	suite.dishStorage.DeleteDishFunc = mockDeleteDish(ErrNotFound)

	suite.menu.AddDish(suite.dishId)

	err := suite.cmd.DeleteDish(context.Background(), uuid.New(), suite.dishId)

	suite.NotNil(err, "DeleteDish should return an error, when dish is not found")
	suite.ErrorIs(err, ErrDishNotFound, "DeleteDish should return ErrDishNotFound")
}

func (suite *deleteDishTestSuite) TestDeleteDish_Valid() {
	suite.restaurantStorage.RestaurantExistsFunc = mockRestaurantExists(true, nil)
	suite.menuStorage.GetMenuByRestaurantFunc = mockGetMenuByRestaurant([]*domain.Menu{suite.menu}, nil)
	suite.dishStorage.DeleteDishFunc = mockDeleteDish(nil)

	suite.menu.AddDish(suite.dishId)

	err := suite.cmd.DeleteDish(context.Background(), suite.restaurantId, suite.dishId)

	suite.Nil(err, "DeleteDish should not return an error, when dish is not found")
}

func TestDeleteDish(t *testing.T) {
	suite.Run(t, new(deleteDishTestSuite))
}
