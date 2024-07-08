package rest

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/indigowar/food_out/services/menu/domain"
	"github.com/indigowar/food_out/services/menu/ports/delivery/rest/api"
	"github.com/indigowar/food_out/services/menu/service/queries"
)

type handler struct {
	dishById         queries.GetDishByIDQuery
	menuById         queries.GetMenuByIDQuery
	menuByRestaurant queries.GetRestaurantsMenusQuery
	restaurantList   queries.GetRestaurantsQuery
	dishValidation   queries.ValidateDishOwnershipQuery
}

var _ api.Handler = &handler{}

// RetrieveDishByID implements api.Handler.
func (h *handler) RetrieveDishByID(ctx context.Context, params api.RetrieveDishByIDParams) (api.RetrieveDishByIDRes, error) {
	dish, err := h.dishById.GetDishByID(ctx, params.ID)
	if err != nil {
		if errors.Is(err, queries.ErrDishNotFound) {
			return &api.RetrieveDishByIDNotFound{
				Code:    http.StatusNotFound,
				Message: fmt.Sprintf("dish with id %s is not found", params.ID.String()),
			}, nil
		}

		return &api.RetrieveDishByIDInternalServerError{
			Code:    http.StatusInternalServerError,
			Message: "internal service error occurred",
		}, nil
	}

	res := convertDish(*dish)
	return &res, nil
}

// RetrieveListOfRestaurants implements api.Handler.
func (h *handler) RetrieveListOfRestaurants(ctx context.Context) (api.RetrieveListOfRestaurantsRes, error) {
	id, err := h.restaurantList.GetRestaurants(ctx)
	if err != nil {
		return &api.Error{
			Code:    http.StatusInternalServerError,
			Message: "internal service errror occurred",
		}, nil
	}

	res := api.RetrieveListOfRestaurantsOKApplicationJSON(id)
	return &res, nil
}

// RetrieveMenuByID implements api.Handler.
func (h *handler) RetrieveMenuByID(ctx context.Context, params api.RetrieveMenuByIDParams) (api.RetrieveMenuByIDRes, error) {
	menu, err := h.menuById.GetMenuByID(ctx, params.ID)
	if err != nil {
		if errors.Is(err, queries.ErrMenuNotFound) {
			return &api.RetrieveMenuByIDNotFound{
				Code:    http.StatusNotFound,
				Message: fmt.Sprintf("menu with id %s is not found", params.ID.String()),
			}, nil
		}

		return &api.RetrieveMenuByIDInternalServerError{
			Code:    http.StatusInternalServerError,
			Message: "internal service error occurred",
		}, nil
	}

	res := convertMenu(*menu)
	return &res, nil
}

// RetrieveMenusByRestaurant implements api.Handler.
func (h *handler) RetrieveMenusByRestaurant(ctx context.Context, params api.RetrieveMenusByRestaurantParams) (api.RetrieveMenusByRestaurantRes, error) {
	menu, err := h.menuByRestaurant.GetRestaurantsMenus(ctx, params.ID)
	if err != nil {
		if errors.Is(err, queries.ErrRestaurantNotFound) {
			return &api.RetrieveMenusByRestaurantNotFound{
				Code:    http.StatusNotFound,
				Message: fmt.Sprintf("restaurant with id %s is not found", params.ID),
			}, nil
		}

		return &api.RetrieveMenusByRestaurantInternalServerError{
			Code:    http.StatusInternalServerError,
			Message: "internal service error occurred",
		}, nil
	}

	res := make(api.RetrieveMenusByRestaurantOKApplicationJSON, len(menu))
	for i, v := range menu {
		res[i] = convertMenu(*v)
	}
	return &res, nil
}

// ValidateRestaurantDishes implements api.Handler.
func (h *handler) ValidateRestaurantDishes(ctx context.Context, req *api.ValidationList) (api.ValidateRestaurantDishesRes, error) {
	result, err := h.dishValidation.ValidateDishOwnership(ctx, req.Restaurant, req.Dishes)
	if err != nil {
		if errors.Is(err, queries.ErrRestaurantNotFound) {
			return &api.ValidateRestaurantDishesNotFound{
				Code:    http.StatusNotFound,
				Message: fmt.Sprintf("restaurant with id %s is not found", req.Restaurant.String()),
			}, nil
		}

		return &api.ValidateRestaurantDishesInternalServerError{
			Code:    http.StatusInternalServerError,
			Message: "internal service error occurred",
		}, nil
	}

	if !result {
		return &api.ValidateRestaurantDishesBadRequest{
			Code:    http.StatusBadRequest,
			Message: fmt.Sprintf("failed ownership validation for %s and a list of IDs", req.Restaurant.String()),
		}, nil
	}

	return &api.ValidateRestaurantDishesOK{}, nil
}

func convertDish(dish domain.Dish) api.Dish {
	return api.Dish{
		ID:    dish.ID(),
		Name:  dish.Name(),
		Image: *dish.Image(),
		Price: dish.Price(),
	}
}

func convertMenu(menu domain.Menu) api.Menu {
	return api.Menu{
		ID:         menu.ID(),
		Name:       menu.Name(),
		Restaurant: menu.Restaurant(),
		Image:      *menu.Image(),
		Dishes:     menu.Dishes(),
	}
}
