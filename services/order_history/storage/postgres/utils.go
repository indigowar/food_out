package postgres

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

func executeInTransaction(
	ctx context.Context,
	conn *pgx.Conn,
	runner func(context.Context, pgx.Tx) error,
) error {
	tx, err := conn.Begin(ctx)
	if err != nil {
		return fmt.Errorf("failed to start a transaction")
	}

	if err := runner(ctx, tx); err != nil {
		rollbackErr := tx.Rollback(ctx)
		return errors.Join(err, rollbackErr)
	}
	return tx.Commit(ctx)
}

func optionalTimeToTimestamp(value *time.Time) pgtype.Timestamp {
	if value == nil {
		return pgtype.Timestamp{Valid: false}
	}
	return pgtype.Timestamp{Time: *value, Valid: true}
}
