package storage

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func runInTxWithReturn[Return any](
	ctx context.Context,
	conn *pgx.Conn,
	action func(ctx context.Context, tx pgx.Tx) (Return, error),
) (Return, error) {
	var empty Return

	tx, err := conn.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return empty, fmt.Errorf("failed to start a transaction: %w", err)
	}

	res, err := action(ctx, tx)
	if err != nil {
		_ = tx.Rollback(ctx)
	} else {
		_ = tx.Commit(ctx)
	}
	return res, err
}

func runInTx(
	ctx context.Context,
	conn *pgx.Conn,
	action func(context.Context, pgx.Tx) error,
) error {
	tx, err := conn.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return fmt.Errorf("failed to start a transaction: %w", err)
	}

	err = action(ctx, tx)
	if err != nil {
		_ = tx.Rollback(ctx)
		return err
	}

	return tx.Commit(ctx)
}
