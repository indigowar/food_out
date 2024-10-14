package postgres

import (
	"net/url"

	"github.com/jackc/pgx/v5/pgtype"

	"github.com/indigowar/food_out/services/accounts/internal/domain"
	"github.com/indigowar/food_out/services/accounts/internal/infra/storage/postgres/gen"
)

const (
// todo: add pgx errors codes
)

func toModel(a gen.Account) domain.Account {
	var name *string = nil
	if a.Name.Valid {
		name = &a.Name.String
	}
	var url *url.URL = nil
	if a.ProfilePicture.Valid {
		url, _ = url.Parse(a.ProfilePicture.String)
	}
	return domain.NewAccountRaw(a.ID, a.Phone, a.Password, name, url)
}

func createInsertParams(account *domain.Account) gen.InsertAccountParams {
	params := gen.InsertAccountParams{
		ID:       account.ID(),
		Phone:    account.Phone(),
		Password: account.Password(),
		Name:     pgtype.Text{},
		Profile:  pgtype.Text{},
	}

	if account.HasName() {
		params.Name = pgtype.Text{String: account.Name(), Valid: true}
	}

	if account.HasProfilePicture() {
		params.Profile = pgtype.Text{String: account.ProfilePicture().String(), Valid: true}
	}

	return params
}
