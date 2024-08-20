package migrations

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"

	"ahbcc/internal/database"
)

type (

	// Run executes the migrations after the database is initialized
	Run func(ctx context.Context, migrationsDir string) error
)

// MakeRun creates a new Run
func MakeRun(db database.Connection, createMigrationsTable CreateMigrationsTable, isMigrationApplied IsMigrationApplied, insertAppliedMigration InsertAppliedMigration) Run {
	return func(ctx context.Context, migrationsDir string) error {
		err := createMigrationsTable(ctx)
		if err != nil {
			return err
		}

		files, _ := filepath.Glob(filepath.Join(migrationsDir, "*.sql"))

		for _, file := range files {
			applied, err := isMigrationApplied(ctx, file)
			if err != nil {
				return err
			}

			if !applied {
				fmt.Printf("Executing %s...\n", file)
				err = executeSQLFromFile(ctx, db, file)
				if err != nil {
					slog.Error(err.Error())
					return FailedToExecuteMigration
				}
				fmt.Printf("Executed %s successfully\n", file)

				err = insertAppliedMigration(ctx, file)
				if err != nil {
					return err
				}
			} else {
				fmt.Printf("Migration file %s already applied\n", file)
			}
		}

		return nil
	}
}

// executeSQLFromFile reads and executes an SQL file
func executeSQLFromFile(ctx context.Context, db database.Connection, filename string) error {
	content, err := os.ReadFile(filename)
	if err != nil {
		slog.Error(err.Error())
		return UnableToReadFile
	}

	_, err = db.Exec(ctx, string(content))
	if err != nil {
		slog.Error(err.Error())
		return UnableToExecuteSQL
	}

	return nil
}
