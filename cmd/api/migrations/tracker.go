package migrations

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
	"log/slog"

	"ahbcc/internal/database"
)

type (
	// CreateMigrationsTable creates a migrations table to keep tracking of which migrations were applied
	CreateMigrationsTable func(ctx context.Context) error

	// IsMigrationApplied checks if a migration file was already applied
	IsMigrationApplied func(ctx context.Context, migrationName string) (bool, error)

	// InsertAppliedMigration inserts a new migration into the migrations table to track that it was already applied
	InsertAppliedMigration func(ctx context.Context, migrationName string) error
)

// MakeCreateMigrationsTable creates a new CreateMigrationsTable
func MakeCreateMigrationsTable(db database.Connection) CreateMigrationsTable {
	const query string = `
		CREATE TABLE IF NOT EXISTS migrations (
			id SERIAL PRIMARY KEY,
			name TEXT UNIQUE NOT NULL,
			applied_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
		);
	`

	return func(ctx context.Context) error {
		_, err := db.Exec(ctx, query)
		if err != nil {
			slog.Error(err.Error())
			return FailedToCreateMigrationsTable
		}

		return nil
	}
}

// MakeIsMigrationApplied creates a new IsMigrationApplied
func MakeIsMigrationApplied(db database.Connection) IsMigrationApplied {
	const query string = `
		SELECT EXISTS (
			SELECT 1 
			FROM migrations 
			WHERE name = $1
		)
	`

	return func(ctx context.Context, migrationName string) (bool, error) {
		var applied bool

		err := db.QueryRow(ctx, query, migrationName).Scan(&applied)
		if errors.Is(err, pgx.ErrNoRows) {
			slog.Error(err.Error())
			return false, FailedToRetrieveIfMigrationWasApplied
		}

		return applied, nil
	}
}

// MakeInsertAppliedMigration creates a new InsertAppliedMigration
func MakeInsertAppliedMigration(db database.Connection) InsertAppliedMigration {
	const query string = `
		INSERT INTO migrations (name) 
		VALUES ($1)
	`
	return func(ctx context.Context, migrationName string) error {
		_, err := db.Exec(ctx, query, migrationName)
		if err != nil {
			return FailedToInsertAppliedMigration
		}

		return nil
	}
}
