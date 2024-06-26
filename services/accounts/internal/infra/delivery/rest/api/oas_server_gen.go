// Code generated by ogen, DO NOT EDIT.

package api

import (
	"context"
)

// Handler handles operations described by OpenAPI v3 specification.
type Handler interface {
	// CreateAccount implements CreateAccount operation.
	//
	// Create a new account.
	//
	// POST /account
	CreateAccount(ctx context.Context, req *AccountCreationInfo) (CreateAccountRes, error)
	// DeleteAccount implements DeleteAccount operation.
	//
	// Delete user with provided ID.
	//
	// DELETE /account/{id}
	DeleteAccount(ctx context.Context, params DeleteAccountParams) (DeleteAccountRes, error)
	// GetAccountInfo implements GetAccountInfo operation.
	//
	// Get Account info by its ID.
	//
	// GET /account/{id}
	GetAccountInfo(ctx context.Context, params GetAccountInfoParams) (GetAccountInfoRes, error)
	// GetOwnInfo implements GetOwnInfo operation.
	//
	// Retrieve info about own account using jwt token.
	//
	// GET /account
	GetOwnInfo(ctx context.Context) (GetOwnInfoRes, error)
	// UpdatePassword implements UpdatePassword operation.
	//
	// Updates Account's password.
	//
	// PUT /account/password
	UpdatePassword(ctx context.Context, req *PasswordUpdateInfo) (UpdatePasswordRes, error)
	// ValidateCredentials implements ValidateCredentials operation.
	//
	// GET /account/credentials
	ValidateCredentials(ctx context.Context, req OptAccountCredentials) (ValidateCredentialsRes, error)
}

// Server implements http server based on OpenAPI v3 specification and
// calls Handler to handle requests.
type Server struct {
	h   Handler
	sec SecurityHandler
	baseServer
}

// NewServer creates new Server.
func NewServer(h Handler, sec SecurityHandler, opts ...ServerOption) (*Server, error) {
	s, err := newServerConfig(opts...).baseServer()
	if err != nil {
		return nil, err
	}
	return &Server{
		h:          h,
		sec:        sec,
		baseServer: s,
	}, nil
}
