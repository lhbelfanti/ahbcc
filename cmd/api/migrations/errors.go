package migrations

import "errors"

var (
	FailedToExecuteMigration = errors.New("failed to execute migrations")
	UnableToReadFile         = errors.New("unable to read file")
	UnableToExecuteSQL       = errors.New("unable to execute SQL")
)

const FailedToRunMigrations string = "Failed to run migrations"
