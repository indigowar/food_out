package rest

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/indigowar/food_out/services/menu/domain"
	"github.com/indigowar/food_out/services/menu/ports/delivery/rest/gen"
	"github.com/indigowar/food_out/services/menu/service/queries"
)

type handler struct {
	dishById         *queries.GetDishByIDQuery
	menuById         *queries.GetMenuByIDQuery
	menuByRestaurant *queries.GetRestaurantsMenusQuery
	restaurantList   *queries.GetRestaurantsQuery
	dishValidation   *queries.ValidateDishOwnershipQuery
}

var _ gen.Handler = &handler{}

// RetrieveDishByID implements api.Handler.
func (h *handler) RetrieveDishByID(ctx context.Context, params gen.RetrieveDishByIDParams) (gen.RetrieveDishByIDRes, error) {
	dish, err := h.dishById.GetDishByID(ctx, params.ID)
	if err != nil {
		if errors.Is(err, queries.ErrDishNotFound) {
			return &gen.RetrieveDishByIDNotFound{
				Code:    http.StatusNotFound,
				Message: fmt.Sprintf("dish with id %s is not found", params.ID.String()),
			}, nil
		}

		return &gen.RetrieveDishByIDInternalServerError{
			Code:    http.StatusInternalServerError,
			Message: "internal service error occurred",
		}, nil
	}

	res := convertDish(*dish)
	return &res, nil
}

// RetrieveListOfRestaurants implements api.Handler.
func (h *handler) RetrieveListOfRestaurants(ctx context.Context) (gen.RetrieveListOfRestaurantsRes, error) {
	id, err := h.restaurantList.GetRestaurants(ctx)
	if err != nil {
		return &gen.Error{
			Code:    http.StatusInternalServerError,
			Message: "internal service errror occurred",
		}, nil
	}

	res := gen.RetrieveListOfRestaurantsOKApplicationJSON(id)
	return &res, nil
}

// RetrieveMenuByID implements api.Handler.
func (h *handler) RetrieveMenuByID(ctx context.Context, params gen.RetrieveMenuByIDParams) (gen.RetrieveMenuByIDRes, error) {
	menu, err := h.menuById.GetMenuByID(ctx, params.ID)
	if err != nil {
		if errors.Is(err, queries.ErrMenuNotFound) {
			return &gen.RetrieveMenuByIDNotFound{
				Code:    http.StatusNotFound,
				Message: fmt.Sprintf("menu with id %s is not found", params.ID.String()),
			}, nil
		}

		return &gen.RetrieveMenuByIDInternalServerError{
			Code:    http.StatusInternalServerError,
			Message: "internal service error occurred",
		}, nil
	}

	res := convertMenu(*menu)
	return &res, nil
}

// RetrieveMenusByRestaurant implements api.Handler.
func (h *handler) RetrieveMenusByRestaurant(ctx context.Context, params gen.RetrieveMenusByRestaurantParams) (gen.RetrieveMenusByRestaurantRes, error) {
	menu, err := h.menuByRestaurant.GetRestaurantsMenus(ctx, params.ID)
	if err != nil {
		if errors.Is(err, queries.ErrRestaurantNotFound) {
			return &gen.RetrieveMenusByRestaurantNotFound{
				Code:    http.StatusNotFound,
				Message: fmt.Sprintf("restaurant with id %s is not found", params.ID),
			}, nil
		}

		return &gen.RetrieveMenusByRestaurantInternalServerError{
			Code:    http.StatusInternalServerError,
			Message: "internal service error occurred",
		}, nil
	}

	res := make(gen.RetrieveMenusByRestaurantOKApplicationJSON, len(menu))
	for i, v := range menu {
		res[i] = convertMenu(*v)
	}
	return &res, nil
}

// ValidateRestaurantDishes implements api.Handler.
func (h *handler) ValidateRestaurantDishes(ctx context.Context, req *gen.ValidationList) (gen.ValidateRestaurantDishesRes, error) {
	result, err := h.dishValidation.ValidateDishOwnership(ctx, req.Restaurant, req.Dishes)
	if err != nil {
		if errors.Is(err, queries.ErrRestaurantNotFound) {
			return &gen.ValidateRestaurantDishesNotFound{
				Code:    http.StatusNotFound,
				Message: fmt.Sprintf("restaurant with id %s is not found", req.Restaurant.String()),
			}, nil
		}

		return &gen.ValidateRestaurantDishesInternalServerError{
			Code:    http.StatusInternalServerError,
			Message: "internal service error occurred",
		}, nil
	}

	if !result {
		return &gen.ValidateRestaurantDishesBadRequest{
			Code:    http.StatusBadRequest,
			Message: fmt.Sprintf("failed ownership validation for %s and a list of IDs", req.Restaurant.String()),
		}, nil
	}

	return &gen.ValidateRestaurantDishesOK{}, nil
}

func convertDish(dish domain.Dish) gen.Dish {
	return gen.Dish{
		ID:    dish.ID(),
		Name:  dish.Name(),
		Image: *dish.Image(),
		Price: dish.Price(),
	}
}

func convertMenu(menu domain.Menu) gen.Menu {
	return gen.Menu{
		ID:         menu.ID(),
		Name:       menu.Name(),
		Restaurant: menu.Restaurant(),
		Image:      *menu.Image(),
		Dishes:     menu.Dishes(),
	}
}
