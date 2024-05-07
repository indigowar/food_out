package redis

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"

	"github.com/indigowar/food_out/services/auth/internal/domain"
)

// todo: add proper error management

const (
	fieldID         = "ID"
	fieldToken      = "Token"
	fieldExpiration = "Expiration"
)

const timeFormat = time.RFC3339

func addSessionData(session domain.Session) txAction {
	return func(ctx context.Context, p redis.Pipeliner) error {
		err := p.HSet(
			ctx, makeSessionKey(session),
			fieldID, session.ID().String(),
			fieldToken, session.Token(),
			fieldExpiration, session.Expiration().Format(timeFormat),
		).Err()
		if err != nil {
			return err
		}

		if err := p.ExpireAt(ctx, makeSessionKey(session), session.Expiration()).Err(); err != nil {
			return err
		}

		return nil
	}
}

func addTokenIndex(session domain.Session) txAction {
	return func(ctx context.Context, p redis.Pipeliner) error {
		duration := time.Until(session.Expiration())
		if err := p.Set(ctx, makeTokenKey(session.Token()), makeSessionKey(session), duration).Err(); err != nil {
			return err
		}

		return nil
	}
}

func removeSession(id uuid.UUID) txAction {
	return func(ctx context.Context, p redis.Pipeliner) error {
		return p.Del(ctx, makeIDKey(id)).Err()
	}
}

func removeTokenIndex(token domain.SessionToken) txAction {
	return func(ctx context.Context, p redis.Pipeliner) error {
		return p.Del(ctx, makeTokenKey(token)).Err()
	}
}
