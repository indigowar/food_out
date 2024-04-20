package postgres

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"

	"github.com/indigowar/food_out/services/accounts/internal/domain"
	"github.com/indigowar/food_out/services/accounts/internal/infra/storage/postgres/queries"
	"github.com/indigowar/food_out/services/accounts/internal/service"
)

type Storage struct {
	queries *queries.Queries
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
	a, err := s.queries.GetAccountByID(ctx, pgtype.UUID{Bytes: id, Valid: true})
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
	return s.queries.DeleteAccount(ctx, pgtype.UUID{Bytes: id, Valid: true})
}

// Update implements service.Storage.
func (s *Storage) Update(ctx context.Context, account *domain.Account) error {
	params := queries.UpdateAccountParams{
		Phone:          account.Phone(),
		Password:       account.Password(),
		Name:           pgtype.Text{},
		ProfilePicture: pgtype.Text{},
		ID:             pgtype.UUID{Bytes: account.ID(), Valid: true},
	}

	if account.HasName() {
		params.Name = pgtype.Text{String: account.Name(), Valid: true}
	}

	if account.HasProfilePicture() {
		params.ProfilePicture = pgtype.Text{String: account.ProfilePicture().String(), Valid: true}
	}

	// todo: add proper error handling for storage
	return s.queries.UpdateAccount(ctx, params)
}

func NewStorage(conn *pgx.Conn) *Storage {
	migration := `
		CREATE TABLE IF NOT EXISTS accounts(
			id UUID PRIMARY KEY,
			phone VARCHAR(32) UNIQUE NOT NULL,
			password VARCHAR(1024) NOT NULL,
			name VARCHAR(64),
			profile_picture VARCHAR(2048)
		);
	`

	conn.Exec(context.Background(), migration)

	return &Storage{
		queries: queries.New(conn),
	}
}
