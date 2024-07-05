package rest

import (
	"context"

	"github.com/indigowar/food_out/services/menu/ports/delivery/rest/api"
)

type handler struct {
}

var _ api.Handler = &handler{}

// RetrieveDishByID implements api.Handler.
func (h *handler) RetrieveDishByID(ctx context.Context, params api.RetrieveDishByIDParams) (api.RetrieveDishByIDRes, error) {
	panic("unimplemented")
}

// RetrieveDishesByMenuId implements api.Handler.
func (h *handler) RetrieveDishesByMenuId(ctx context.Context, params api.RetrieveDishesByMenuIdParams) (api.RetrieveDishesByMenuIdRes, error) {
	panic("unimplemented")
}

// RetrieveListOfRestaurants implements api.Handler.
func (h *handler) RetrieveListOfRestaurants(ctx context.Context) (api.RetrieveListOfRestaurantsRes, error) {
	panic("unimplemented")
}

// RetrieveMenuByID implements api.Handler.
func (h *handler) RetrieveMenuByID(ctx context.Context, params api.RetrieveMenuByIDParams) (api.RetrieveMenuByIDRes, error) {
	panic("unimplemented")
}

// RetrieveMenusByRestaurant implements api.Handler.
func (h *handler) RetrieveMenusByRestaurant(ctx context.Context, params api.RetrieveMenusByRestaurantParams) (api.RetrieveMenusByRestaurantRes, error) {
	panic("unimplemented")
}

// ValidateRestaurantDishes implements api.Handler.
func (h *handler) ValidateRestaurantDishes(ctx context.Context, req *api.ValidationList) (api.ValidateRestaurantDishesRes, error) {
	panic("unimplemented")
}
