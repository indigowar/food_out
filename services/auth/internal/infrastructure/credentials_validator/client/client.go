package client

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"

	"github.com/indigowar/food_out/services/auth/internal/infrastructure/credentials_validator/client/api"
	"github.com/indigowar/food_out/services/auth/internal/service"
)

type CredentialsValidator struct {
	api *api.Client
}

var _ service.CredentialsValidator = &CredentialsValidator{}

// Validate implements service.CredentialsValidator.
func (c *CredentialsValidator) Validate(ctx context.Context, phone string, password string) (uuid.UUID, error) {
	res, err := c.makeRequest(ctx, phone, password)
	if err != nil {
		return uuid.UUID{}, err
	}

	idRes, ok := res.(*api.AccountId)
	if !ok {
		return uuid.UUID{}, c.handleErrorResponse(res)
	}

	id, err := uuid.Parse(idRes.ID)
	if err != nil {
		return uuid.UUID{}, errors.New("invalid ID received")
	}

	return id, nil
}

func (c *CredentialsValidator) makeRequest(ctx context.Context, phone string, password string) (api.ValidateCredentialsRes, error) {
	res, err := c.api.ValidateCredentials(ctx, &api.AccountCredentials{
		Phone:    phone,
		Password: password,
	})

	if err != nil {
		return nil, fmt.Errorf("http error: %w", err)
	}

	return res, nil
}

func (c *CredentialsValidator) handleErrorResponse(res api.ValidateCredentialsRes) error {
	if _, ok := res.(*api.ValidateCredentialsBadRequest); ok {
		return service.ErrCredentialsValidatorInvalid
	}

	if _, ok := res.(*api.ValidateCredentialsInternalServerError); ok {
		return errors.New("internal server error")
	}

	return errors.New("unexpected response received")
}

func NewCredentialsValidator(accountUrl string) (*CredentialsValidator, error) {
	api, err := api.NewClient(accountUrl)
	if err != nil {
		return nil, err
	}
	return &CredentialsValidator{api: api}, nil
}
