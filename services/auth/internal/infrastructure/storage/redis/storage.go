package redis

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"

	"github.com/indigowar/food_out/services/auth/internal/domain"
	"github.com/indigowar/food_out/services/auth/internal/service"
)

type Storage struct {
	client *redis.Client
}

var _ service.Storage = &Storage{}

func NewStorage(client *redis.Client) *Storage {
	return &Storage{
		client: client,
	}
}

// Add implements service.Storage.
func (s *Storage) Add(ctx context.Context, session domain.Session) error {
	return makeTx(
		s.client,
		addSessionData(session),
		addTokenIndex(session),
	)(ctx)
}

// GetByID implements service.Storage.
func (s *Storage) GetByID(ctx context.Context, id uuid.UUID) (domain.Session, error) {
	data, err := s.client.HGetAll(ctx, makeIDKey(id)).Result()
	if err != nil {
		return domain.Session{}, err
	}

	return sessionFromData(data)
}

// GetByToken implements service.Storage.
func (s *Storage) GetByToken(ctx context.Context, token domain.SessionToken) (domain.Session, error) {
	sessionKey, err := s.client.Get(ctx, makeTokenKey(token)).Result()
	if err != nil {
		return domain.Session{}, err
	}

	data, err := s.client.HGetAll(ctx, sessionKey).Result()
	if err != nil {
		return domain.Session{}, err
	}

	return sessionFromData(data)
}

// RemoveByID implements service.Storage.
func (s *Storage) RemoveByID(ctx context.Context, id uuid.UUID) (domain.Session, error) {
	session, err := s.GetByID(ctx, id)
	if err != nil {
		return domain.Session{}, err
	}

	err = makeTx(
		s.client,
		removeSession(session.ID()),
		removeTokenIndex(session.Token()),
	)(ctx)

	if err != nil {
		return domain.Session{}, err
	}

	return session, nil
}

// RemoveByToken implements service.Storage.
func (s *Storage) RemoveByToken(ctx context.Context, token domain.SessionToken) (domain.Session, error) {
	session, err := s.GetByToken(ctx, token)
	if err != nil {
		return domain.Session{}, err
	}

	err = makeTx(
		s.client,
		removeSession(session.ID()),
		removeTokenIndex(session.Token()),
	)(ctx)

	if err != nil {
		return domain.Session{}, err
	}

	return session, nil
}

func sessionFromData(data map[string]string) (domain.Session, error) {
	var id uuid.UUID
	var token domain.SessionToken
	var expDate time.Time

	var err error
	for key, value := range data {
		switch key {
		case fieldID:
			id, err = uuid.Parse(value)
			if err != nil {
				return domain.Session{}, err
			}
		case fieldToken:
			token = domain.SessionToken(value)
		case fieldExpiration:
			expDate, err = time.Parse(timeFormat, value)
			if err != nil {
				return domain.Session{}, err
			}
		default:
			return domain.Session{}, errors.New("invalid key")
		}
	}

	return domain.NewSession(id, token, expDate), nil
}
