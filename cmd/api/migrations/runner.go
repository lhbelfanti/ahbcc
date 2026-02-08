package migrations

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/lhbelfanti/corpus-creator/internal/database"
	"github.com/lhbelfanti/corpus-creator/internal/log"
)

// Run executes the migrations after the database is initialized
type Run func(ctx context.Context, migrationsDir string) error

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
				log.Info(ctx, fmt.Sprintf("Executing %s...\n", file))
				err = executeSQLFromFile(ctx, db, file)
				if err != nil {
					log.Error(ctx, err.Error())
					return FailedToExecuteMigration
				}
				log.Info(ctx, fmt.Sprintf("Executed %s successfully\n", file))

				err = insertAppliedMigration(ctx, file)
				if err != nil {
					return err
				}
			} else {
				log.Info(ctx, fmt.Sprintf("Migration file %s already applied\n", file))
			}
		}

		return nil
	}
}

// executeSQLFromFile reads and executes an SQL file
func executeSQLFromFile(ctx context.Context, db database.Connection, filename string) error {
	content, err := os.ReadFile(filename)
	if err != nil {
		log.Error(ctx, err.Error())
		return UnableToReadFile
	}

	_, err = db.Exec(ctx, string(content))
	if err != nil {
		log.Error(ctx, err.Error())
		return UnableToExecuteSQL
	}

	return nil
}
