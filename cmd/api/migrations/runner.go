package migrations

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"

	"github.com/jackc/pgx/v5/pgconn"
)

type (
	// PgConnExecutor PostgresSQL database connection interface
	PgConnExecutor interface {
		Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
	}

	// Run executes the migrations after the database is initialized
	Run func(ctx context.Context, migrationsDir string) error
)

// MakeRun creates a new Run
func MakeRun(conn PgConnExecutor) Run {
	return func(ctx context.Context, migrationsDir string) error {
		files, _ := filepath.Glob(filepath.Join(migrationsDir, "*.sql"))

		var err error
		for _, file := range files {
			fmt.Printf("Executing %s...\n", file)
			err = executeSQLFromFile(ctx, conn, file)
			if err != nil {
				slog.Error(err.Error())
				return FailedToExecuteMigration
			}
			fmt.Printf("Executed %s successfully\n", file)
		}

		return nil
	}
}

// executeSQLFromFile reads and executes an SQL file
func executeSQLFromFile(ctx context.Context, conn PgConnExecutor, filename string) error {
	content, err := os.ReadFile(filename)
	if err != nil {
		slog.Error(err.Error())
		return UnableToReadFile
	}

	_, err = conn.Exec(ctx, string(content))
	if err != nil {
		slog.Error(err.Error())
		return UnableToExecuteSQL
	}

	return nil
}
