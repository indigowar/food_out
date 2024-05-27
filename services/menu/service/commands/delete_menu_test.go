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

type deleteMenuTestSuite struct {
	suite.Suite

	restaurantStorage RestaurantStoragePortMock
	menuStorage       MenuStoragePortMock

	cmd *DeleteMenuCommand
}

func (suite *deleteMenuTestSuite) SetupTest() {
	suite.cmd = NewDeleteMenuCommand(
		&suite.restaurantStorage,
		&suite.menuStorage,
		slog.New(slog.NewTextHandler(os.Stdout, nil)),
	)
}

func (suite *deleteMenuTestSuite) TestDeleteMenu_InvalidUUID() {
	args := [][2]uuid.UUID{
		{uuid.New(), uuid.UUID{}},
		{uuid.UUID{}, uuid.New()},
		{uuid.UUID{}, uuid.UUID{}},
	}

	for _, arg := range args {
		err := suite.cmd.DeleteMenu(context.Background(), arg[0], arg[1])

		suite.NotNil(err, "DeleteMenu should return an error, when input data is default uuid(uuid.UUID{})")
		suite.ErrorIs(err, ErrInvalidData, "DeleteMenu should return an ErrInvalidData")
	}
}

func (suite *deleteMenuTestSuite) TestDeleteMenu_RestaurantNotFound() {
	suite.restaurantStorage.RestaurantExistsFunc = mockRestaurantExists(false, nil)

	err := suite.cmd.DeleteMenu(context.Background(), uuid.New(), uuid.New())

	suite.NotNil(err, "DeleteMenu should return an error, because restaurant is not found")
	suite.ErrorIs(err, ErrRestaurantNotFound, "DeleteMenu should return ErrRestaurantNotFound")
}

func (suite *deleteMenuTestSuite) TestDeleteMenu_RestaurantPortError() {
	suite.restaurantStorage.RestaurantExistsFunc = mockRestaurantExists(false, errors.New("unexpected errror"))

	err := suite.cmd.DeleteMenu(context.Background(), uuid.New(), uuid.New())

	suite.NotNil(err, "Error occurred in RestaurantStoragePort, DeleteMenu should return an error")
	suite.ErrorIs(err, ErrInternal, "DeleteMenu should return ErrInternal")
}

func (suite *deleteMenuTestSuite) TestDeleteMenu_MenuNotFound() {
	suite.restaurantStorage.RestaurantExistsFunc = mockRestaurantExists(true, nil)
	suite.menuStorage.GetMenuFunc = mockGetMenu(nil, ErrNotFound)

	err := suite.cmd.DeleteMenu(context.Background(), uuid.New(), uuid.New())

	suite.NotNil(err, "DeleteMenu should return an error, because menu is not found in MenuStoragePort")
	suite.ErrorIs(err, ErrMenuNotFound, "DeleteMenu should return ErrMenuNotFound")
}

func TestDeleteMenu(t *testing.T) {
	suite.Run(t, new(deleteMenuTestSuite))
}
