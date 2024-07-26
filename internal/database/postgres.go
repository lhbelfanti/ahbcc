package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

// InitPostgres creates a new
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
