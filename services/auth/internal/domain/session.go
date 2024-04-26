package domain

import (
	"github.com/google/uuid"
)

type Session struct {
	id    uuid.UUID
	token SessionToken
}

func (s Session) ID() uuid.UUID {
	return s.id
}

func (s Session) Token() SessionToken {
	return s.token
}

func NewSession(id uuid.UUID, token SessionToken) Session {
	return Session{
		id:    id,
		token: token,
	}
}
