package database

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

// InitPostgres creates a new postgres instance
func InitPostgres() (*Postgres, error) {
	var initErr error
	pgOnce.Do(func() {
		db, err := pgxpool.New(context.Background(), resolveDatabaseURL())
		initErr = err
		pgInstance = &Postgres{db}
	})

	return pgInstance, initErr
}

// Database returns the Postgres connection pool
func (pg *Postgres) Database() *pgxpool.Pool {
	return pg.db
}

// Close closes the database connection
func (pg *Postgres) Close() {
	pg.db.Close()
}

// MakeCollectRows creates a new CollectRows
func MakeCollectRows[T any]() CollectRows[T] {
	return func(rows pgx.Rows) ([]T, error) {
		return pgx.CollectRows(rows, pgx.RowToStructByPos[T])
	}
}
