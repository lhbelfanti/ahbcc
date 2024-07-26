package database

import (
	"context"
	"fmt"
	"os"
	"sync"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type (
	Connection interface {
		Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
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

const databaseURL string = "postgresql://%s:%s@postgres_db:5432/%s?sslmode=disable"

func resolveDatabaseURL() string {
	dbUser := os.Getenv("POSTGRES_DB_USER")
	dbPass := os.Getenv("POSTGRES_DB_PASS")
	dbName := os.Getenv("POSTGRES_DB_NAME")

	return fmt.Sprintf(databaseURL, dbUser, dbPass, dbName)
}
