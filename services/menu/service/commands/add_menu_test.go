package commands

import (
	"context"
	"log/slog"
	"os"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
)

type addMenuTestSuite struct {
	suite.Suite

	menuStorage       MenuStoragePortMock
	restaurantStorage RestaurantStoragePortMock

	cmd *AddMenuCommand
}

func (suite *addMenuTestSuite) SetupTest() {
	suite.cmd = NewAddMenuCommand(
		&suite.restaurantStorage,
		&suite.menuStorage,
		slog.New(slog.NewTextHandler(os.Stdout, nil)),
	)
}

// TestAddMenu_IncompleteArgument tests AddMenu with incomplete argument set.
// Mocks are not specified because the input validation should be before interaction with ports
func (suite *addMenuTestSuite) TestAddMenu_IncompleteArgument() {
	args := []AddMenuArgument{
		{"", uuid.New(), validUrl},                 // no name
		{"Valid Menu Name", uuid.UUID{}, validUrl}, // no restaurant id
		{"Valid Menu Name", uuid.New(), nil},       // no image url
	}

	for _, arg := range args {
		menu, err := suite.cmd.AddMenu(context.Background(), arg)

		suite.NotNil(err, "AddMenu shouldn't return nil error, when argument set is incomplete")
		suite.Equal(uuid.UUID{}, menu, "AddMenu should return an empty id, when argument set is incomplete")
		suite.ErrorIs(err, ErrInvalidData, "AddMenu should return ErrInvalidData, when argument set is incomplete")
	}
}

// TestAddMenu_RestaurantNotInStorage test AddMenu when restaurant is not in the storage
func (suite *addMenuTestSuite) TestAddMenu_RestaurantNotInStorage() {
	suite.restaurantStorage.RestaurantExistsFunc = mockRestaurantExists(false, nil)

	input := AddMenuArgument{"Menu", uuid.New(), validUrl}
	menu, err := suite.cmd.AddMenu(context.Background(), input)

	suite.NotNil(err, "AddMenu should return an error, restaurant is not found")
	suite.Equal(uuid.UUID{}, menu, "AddMenu should return an empty id, restaurant is not found")
	suite.ErrorIs(err, ErrRestaurantNotFound, "AddMenu should return ErrRestaurantNotFound")
}

// TestAddMenu_MenuAlreadyExists tests AddMenu when the Menu with given name already exists for this restaurant
func (suite *addMenuTestSuite) TestAddMenu_MenuAlreadyExists() {
	suite.restaurantStorage.RestaurantExistsFunc = mockRestaurantExists(true, nil)
	suite.menuStorage.AddMenuFunc = mockAddMenu(ErrAlreadyExists)

	input := AddMenuArgument{"Menu", uuid.New(), validUrl}
	menu, err := suite.cmd.AddMenu(context.Background(), input)

	suite.NotNil(err, "AddMenu should return an error, menu is duplicated")
	suite.Equal(uuid.UUID{}, menu, "AddMenu should return an empty id, menu is duplicated")
	suite.ErrorIs(err, ErrMenuAlreadyExists, "AddMenu should return ErrMenuAlreadyExists")
}

func (suite *addMenuTestSuite) TestAddMenu_Valid() {
	suite.restaurantStorage.RestaurantExistsFunc = mockRestaurantExists(true, nil)
	suite.menuStorage.AddMenuFunc = mockAddMenu(nil)

	input := AddMenuArgument{"Menu", uuid.New(), validUrl}

	menu, err := suite.cmd.AddMenu(context.Background(), input)
	empty := uuid.UUID{}

	suite.Nil(err, "AddMenu should not return an error")
	suite.NotEqual(menu, empty, "AddMenu should not return an empty id")
}

func TestAddMenuCommand(t *testing.T) {
	suite.Run(t, new(addMenuTestSuite))
}
