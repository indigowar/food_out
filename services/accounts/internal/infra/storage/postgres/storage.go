package postgres

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"

	"github.com/indigowar/food_out/services/accounts/internal/domain"
	"github.com/indigowar/food_out/services/accounts/internal/infra/storage/postgres/gen"
	"github.com/indigowar/food_out/services/accounts/internal/service"
)

type Storage struct {
	queries *gen.Queries
}

var _ service.Storage = &Storage{}

// Add implements service.Storage.
func (s *Storage) Add(ctx context.Context, account *domain.Account) error {
	params := createInsertParams(account)

	if err := s.queries.InsertAccount(ctx, params); err != nil {
		// todo: add proper error handling for storage
		return err
	}

	return nil
}

// GetAll implements service.Storage.
func (s *Storage) GetAll(ctx context.Context) ([]*domain.Account, error) {
	data, err := s.queries.GetAllAccounts(ctx)
	if err != nil {
		// todo: add proper error handling for storage
		return nil, err
	}

	accounts := make([]*domain.Account, len(data))
	for i, v := range data {
		a := toModel(v)
		accounts[i] = &a
	}
	return accounts, nil
}

// GetByID implements service.Storage.
func (s *Storage) GetByID(ctx context.Context, id uuid.UUID) (*domain.Account, error) {
	a, err := s.queries.GetAccountByID(ctx, id)
	if err != nil {
		// todo: add proper error handling for storage
		return nil, err
	}
	account := toModel(a)
	return &account, nil
}

// GetByPhone implements service.Storage.
func (s *Storage) GetByPhone(ctx context.Context, phone string) (*domain.Account, error) {
	a, err := s.queries.GetAccountByPhone(ctx, phone)
	if err != nil {
		// todo: add proper error handling for storage
		return nil, err
	}
	account := toModel(a)
	return &account, nil
}

// Remove implements service.Storage.
func (s *Storage) Remove(ctx context.Context, id uuid.UUID) error {
	// todo: add proper error handling for storage
	return s.queries.DeleteAccount(ctx, id)
}

// Update implements service.Storage.
func (s *Storage) Update(ctx context.Context, account *domain.Account) error {
	params := gen.UpdateAccountParams{
		Phone:    account.Phone(),
		Password: account.Password(),
		Name:     pgtype.Text{},
		Profile:  pgtype.Text{},
		ID:       account.ID(),
	}

	if account.HasName() {
		params.Name = pgtype.Text{String: account.Name(), Valid: true}
	}

	if account.HasProfilePicture() {
		params.Profile = pgtype.Text{String: account.ProfilePicture().String(), Valid: true}
	}

	// todo: add proper error handling for storage
	return s.queries.UpdateAccount(ctx, params)
}

func NewStorage(conn *pgx.Conn) *Storage {
	return &Storage{
		queries: gen.New(conn),
	}
}
