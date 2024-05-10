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

// todo: add proper error management

// Those constants are the keys for fields of Session struct in Redis.
const (
	fieldID         = "ID"
	fieldToken      = "Token"
	fieldExpiration = "Expiration"
)

// This time format is used to store a timestamp as a string in Redis
const timeFormat = time.RFC3339

func addSessionData(session domain.Session) txAction {
	return func(ctx context.Context, p redis.Pipeliner) error {
		err := p.HSet(
			ctx, makeIDKey(session.ID()),
			fieldID, session.ID().String(),
			fieldToken, session.Token(),
			fieldExpiration, session.Expiration().Format(timeFormat),
		).Err()
		if err != nil {
			return err
		}

		if err := p.ExpireAt(ctx, makeIDKey(session.ID()), session.Expiration()).Err(); err != nil {
			return err
		}

		return nil
	}
}

func addTokenIndex(session domain.Session) txAction {
	return func(ctx context.Context, p redis.Pipeliner) error {
		duration := time.Until(session.Expiration())
		if err := p.Set(ctx, makeTokenKey(session.Token()), makeIDKey(session.ID()), duration).Err(); err != nil {
			return err
		}

		return nil
	}
}

func removeSession(id uuid.UUID) txAction {
	return func(ctx context.Context, p redis.Pipeliner) error {
		err := p.Del(ctx, makeIDKey(id)).Err()
		if errors.Is(err, redis.Nil) {
			return service.ErrStorageNotFound
		}
		return err
	}
}

func removeTokenIndex(token domain.SessionToken) txAction {
	return func(ctx context.Context, p redis.Pipeliner) error {
		err := p.Del(ctx, makeTokenKey(token)).Err()
		if errors.Is(err, redis.Nil) {
			return service.ErrStorageNotFound
		}
		return err
	}
}
