package rest

import (
	"context"
	"net/http"

	"github.com/indigowar/food_out/services/accounts/internal/infra/delivery/rest/api"
	"github.com/indigowar/food_out/services/accounts/internal/service"
)

type Delivery struct {
	service *serviceWrapper
	server  *http.Server
}

func (d *Delivery) Run() error {
	return d.server.ListenAndServe()
}

func (d *Delivery) Shutdown(ctx context.Context) error {
	return d.server.Shutdown(ctx)
}

func NewDelivery(service *service.Service) *Delivery {
	server := &http.Server{}
	return &Delivery{service: &serviceWrapper{service: service}, server: server}
}

type serviceWrapper struct {
	service *service.Service
}

var _ api.Handler = (*serviceWrapper)(nil)

// CreateAccount implements api.Handler.
func (s *serviceWrapper) CreateAccount(ctx context.Context, req *api.AccountCreationInfo) (api.CreateAccountRes, error) {
	panic("unimplemented")
}

// GetAccountInfo implements api.Handler.
func (s *serviceWrapper) GetAccountInfo(ctx context.Context, params api.GetAccountInfoParams) (api.GetAccountInfoRes, error) {
	panic("unimplemented")
}

// GetOwnInfo implements api.Handler.
func (s *serviceWrapper) GetOwnInfo(ctx context.Context) (api.GetOwnInfoRes, error) {
	panic("unimplemented")
}

// UpdatePassword implements api.Handler.
func (s *serviceWrapper) UpdatePassword(ctx context.Context, req *api.PasswordUpdateInfo) (api.UpdatePasswordRes, error) {
	panic("unimplemented")
}
