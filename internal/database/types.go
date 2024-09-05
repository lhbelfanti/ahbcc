package database

import (
	"context"
	"fmt"
	"os"
	"sync"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type (
	// Connection is an interface created as an abstraction of pgxpool.Pool to be able to mock it
	Connection interface {
		Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
		Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
		QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
	}

	// Postgres is the representation of a postgres database connection
	Postgres struct {
		db *pgxpool.Pool
	}
)

var (
	pgInstance *Postgres
	pgOnce     sync.Once
)

const databaseURL string = "postgresql://%s:%s@postgres_db:%s/%s?sslmode=disable"

func resolveDatabaseURL() string {
	dbUser := os.Getenv("POSTGRES_DB_USER")
	dbPass := os.Getenv("POSTGRES_DB_PASS")
	dbName := os.Getenv("POSTGRES_DB_NAME")
	dbPort := os.Getenv("POSTGRES_DB_PORT")

	return fmt.Sprintf(databaseURL, dbUser, dbPass, dbPort, dbName)
}
