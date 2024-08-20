package migrations

import "context"

// MockCreateMigrationsTable mocks CreateMigrationsTable function
func MockCreateMigrationsTable(err error) CreateMigrationsTable {
	return func(ctx context.Context) error {
		return err
	}
}

// MockIsMigrationApplied mocks IsMigrationApplied function
func MockIsMigrationApplied(applied bool, err error) IsMigrationApplied {
	return func(ctx context.Context, migrationName string) (bool, error) {
		return applied, err
	}
}

// MockInsertAppliedMigration mocks InsertAppliedMigration function
func MockInsertAppliedMigration(err error) InsertAppliedMigration {
	return func(ctx context.Context, migrationName string) error {
		return err
	}
}

// MockRun mocks Run function
func MockRun(err error) Run {
	return func(ctx context.Context, migrationsDir string) error {
		return err
	}
}
