package migrations

import "errors"

var (
	FailedToExecuteMigration = errors.New("failed to execute migrations")
	UnableToReadFile         = errors.New("unable to read file")
	UnableToExecuteSQL       = errors.New("unable to execute SQL")

	FailedToCreateMigrationsTable         = errors.New("failed to create migrations table")
	FailedToRetrieveIfMigrationWasApplied = errors.New("failed to retrieve the applied migration")
	FailedToInsertAppliedMigration        = errors.New("failed to insert applied migration")
)

const FailedToRunMigrations string = "Failed to run migrations"
