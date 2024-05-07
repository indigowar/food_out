package domain

import (
	"time"

	"github.com/google/uuid"
)

type Session struct {
	id         uuid.UUID
	token      SessionToken
	expiration time.Time
}

func (s Session) ID() uuid.UUID {
	return s.id
}

func (s Session) Token() SessionToken {
	return s.token
}

func (s Session) Expiration() time.Time {
	return s.expiration
}

func NewSession(id uuid.UUID, token SessionToken, expiration time.Time) Session {
	return Session{
		id:         id,
		token:      token,
		expiration: expiration,
	}
}
