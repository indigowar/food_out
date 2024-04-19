package service

import (
	"context"
	"net/url"

	"github.com/google/uuid"

	"github.com/indigowar/food_out/services/accounts/internal/domain"
)

type Service struct {
	storage Storage
}

func (svc *Service) GetAccountByID(ctx context.Context, id uuid.UUID) (*domain.Account, error) {
	panic("not implemented")
}

func (svc *Service) GetAccountByPhone(ctx context.Context, phone string) (*domain.Account, error) {
	panic("not implemented")
}

func (svc *Service) CreateAccount(ctx context.Context, phone string, password string) (*domain.Account, error) {
	panic("not implemented")
}

func (svc *Service) DeleteAccount(ctx context.Context, id uuid.UUID) error {
	panic("not implemented")
}

func (svc *Service) UpdatePassword(ctx context.Context, id uuid.UUID, password string) error {
	panic("not implemented")
}

func (svc *Service) UpdateName(ctx context.Context, id uuid.UUID, name string) error {
	panic("not implemented")
}

func (svc *Service) UpdateProfilePicture(ctx context.Context, id uuid.UUID, profilePicture *url.URL) error {
	panic("not implemented")
}

func NewService(storage Storage) *Service {
	return &Service{
		storage: storage,
	}
}
