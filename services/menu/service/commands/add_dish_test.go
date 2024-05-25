package commands

import (
	"context"
	"github.com/indigowar/food_out/services/menu/domain"
	"net/url"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
)

// just a valid Url for tests
var validUrl, _ = url.Parse("https://google.com")

type addDishCommandSuite struct {
	suite.Suite

	dishStorage       DishStoragePortMock
	restaurantStorage RestaurantStoragePortMock
	menuStorage       MenuStoragePortMock

	cmd *AddDishCommand
}

func (suite *addDishCommandSuite) SetupTest() {
	suite.cmd = NewAddDishCommand(
		&suite.dishStorage,
		&suite.menuStorage,
		&suite.restaurantStorage,
		nil,
	)
}

// TestAddDish_IncompleteArgument tests the AddDish command on Invalid data for creation of dish.
// The mock functions are not specified, which will inflict an error, because interaction with ports should be after
// validation of the input data.
func (suite *addDishCommandSuite) TestAddDish_IncompleteArgument() {
	args := []AddDishArgument{
		{uuid.UUID{}, uuid.New(), "name", validUrl, 64.0}, // No restaurant
		{uuid.New(), uuid.UUID{}, "name", validUrl, 64.0}, // No menu
		{uuid.New(), uuid.New(), "", validUrl, 64.0},      // No name
		{uuid.New(), uuid.New(), "Name", nil, 64.0},       // No image
		{uuid.New(), uuid.New(), "Name", validUrl, 0},     // No price
	}

	for _, arg := range args {
		dish, err := suite.cmd.AddDish(context.Background(), arg)

		suite.NotNil(err, "AddDish shouldn't return nil error, when argument set is incomplete")
		suite.Equal(uuid.UUID{}, dish, "AddDish should return an empty id, when argument set is incomplete")
		suite.ErrorIs(err, ErrInvalidData, "AddDish should return ErrInvalidData, when argument set is incomplete")
	}
}

// TestAddDish_InvalidPrice tests the AddDish command on invalid price value(zero or negative), when creating a dish.
// The mock functions are not specified, which will inflict a valid error, because interaction with ports should be
// after validation of input data.
func (suite *addDishCommandSuite) TestAddDish_InvalidPrice() {
	args := []AddDishArgument{
		{uuid.New(), uuid.New(), "Name", validUrl, -500.0}, // negative
		{uuid.New(), uuid.New(), "Name", validUrl, 0},      // equals zero
	}

	for _, arg := range args {
		dish, err := suite.cmd.AddDish(context.Background(), arg)

		suite.NotNil(err, "AddDish shouldn't return nil error, when price is invalid")
		suite.Equal(uuid.UUID{}, dish, "AddDish should return an empty id, when price is invalid")
		suite.ErrorIs(err, ErrInvalidData, "AddDish should return ErrInvalidData, when price is invalid")
	}
}

// TestAddDish_RestaurantNotInStorage tests the AddDish command on absence of restaurant in restaurant storage.
func (suite *addDishCommandSuite) TestAddDish_RestaurantNotInStorage() {
	suite.restaurantStorage.RestaurantExistsFunc = mockRestaurantExists(false, nil)

	input := AddDishArgument{uuid.New(), uuid.New(), "Name", validUrl, 100.0}
	dish, err := suite.cmd.AddDish(context.Background(), input)

	suite.NotNil(err, "AddDish should return an error, restaurant is not found")
	suite.Equal(uuid.UUID{}, dish, "AddDish should return an empty id, restaurant is not found")
	suite.ErrorIs(err, ErrRestaurantNotFound, "AddDish should return ErrRestaurantNotFound")
}

// TestAddDish_MenuNotInStorage tests the AddDish command on absence of menu in menu restaurant.
func (suite *addDishCommandSuite) TestAddDish_MenuNotInStorage() {
	suite.restaurantStorage.RestaurantExistsFunc = mockRestaurantExists(true, nil)
	suite.menuStorage.GetMenuFunc = mockGetMenu(nil, ErrNotFound)

	input := AddDishArgument{uuid.New(), uuid.New(), "Name", validUrl, 100.0}
	dish, err := suite.cmd.AddDish(context.Background(), input)

	suite.NotNil(err, "AddDish should return an error, restaurant is not found")
	suite.Equal(uuid.UUID{}, dish, "AddDish should return an empty id, restaurant is not found")
	suite.ErrorIs(err, ErrMenuNotFound, "AddDish should return ErrMenuNotFound")
}

func (suite *addDishCommandSuite) TestAddDish_Valid() {
	suite.restaurantStorage.RestaurantExistsFunc = mockRestaurantExists(true, nil)
	suite.menuStorage.GetMenuFunc = mockGetMenu(&domain.Menu{}, nil)
	suite.dishStorage.AddDishFunc = mockAddDish(nil)
	suite.menuStorage.UpdateMenuFunc = mockUpdateMenu(nil)

	input := AddDishArgument{uuid.New(), uuid.New(), "Name", validUrl, 100.0}
	dish, err := suite.cmd.AddDish(context.Background(), input)

	suite.Nil(err, "AddDish should not return an error: %s", err)
	suite.NotEqual(uuid.UUID{}, dish, "AddDish should return a valid id of created dish")
}

func TestAddDishCommand(t *testing.T) {
	suite.Run(t, new(addDishCommandSuite))
}
