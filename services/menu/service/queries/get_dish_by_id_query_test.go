package queries

import (
	"context"
	"errors"
	"log/slog"
	"net/url"
	"os"
	"testing"

	"github.com/google/uuid"
	"github.com/indigowar/food_out/services/menu/domain"
	"github.com/stretchr/testify/suite"
)

type getDishByIDQueryTestSuite struct {
	suite.Suite

	retriever DishRetrieverMock
	query     *GetDishByIDQuery
}

func (suite *getDishByIDQueryTestSuite) SetupTest() {
	suite.query = NewGetDishByIDQuery(
		&suite.retriever,
		slog.New(slog.NewTextHandler(os.Stdout, nil)),
	)
}

func (suite *getDishByIDQueryTestSuite) TestGetDishByID_NotFound() {
	suite.retriever.RetrieveByIDFunc = mockRetrieveByID(nil, ErrNotFound)

	dish, err := suite.query.GetDishByID(context.Background(), uuid.New())
	suite.NotNil(err, "GetDishByID should return an error")
	suite.Nil(dish, "GetDishByID should return a nil instead of dish")
	suite.ErrorIs(err, ErrDishNotFound, "GetDishByID should return a ErrDishNotFound error")
}

func (suite *getDishByIDQueryTestSuite) TestGetDishByID_UnexpectedError() {
	suite.retriever.RetrieveByIDFunc = mockRetrieveByID(nil, errors.New("unexpected error occurred"))

	dish, err := suite.query.GetDishByID(context.Background(), uuid.New())
	suite.NotNil(err, "GetDishByID should return an error")
	suite.Nil(dish, "GetDishByID should return a nil instead of dish")
	suite.ErrorIs(err, ErrInternal, "GetDishByID should return a ErrInternal error")
}

func (suite *getDishByIDQueryTestSuite) TestGetDishByID_Valid() {
	validurl, _ := url.Parse("https://google.com")
	dish, _ := domain.NewDish("DishName", validurl, 120.5)

	suite.retriever.RetrieveByIDFunc = mockRetrieveByID(dish, nil)

	result, err := suite.query.GetDishByID(context.Background(), dish.ID())

	suite.Nil(err, "GetDishByID should not return an error")
	suite.Equal(dish, result)
}

func TestGetDishByIDQuery(t *testing.T) {
	suite.Run(t, new(getDishByIDQueryTestSuite))
}
