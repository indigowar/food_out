package redis

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type txAction func(context.Context, redis.Pipeliner) error
type txExecutor func(context.Context) error

func makeTx(cmd redis.Cmdable, actions ...txAction) txExecutor {
	return func(ctx context.Context) error {
		tx := cmd.TxPipeline()

		for _, action := range actions {
			if err := action(ctx, tx); err != nil {
				tx.Discard()
				return err
			}
		}

		if _, err := tx.Exec(ctx); err != nil {
			return err
		}
		return nil
	}
}
