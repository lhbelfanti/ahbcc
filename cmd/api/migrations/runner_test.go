package migrations_test

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/lhbelfanti/corpus-creator/cmd/api/migrations"
	"github.com/lhbelfanti/corpus-creator/internal/database"
)

const migrationsTestDir string = "./migrations_test"

func TestMain(m *testing.M) {
	setupTestEnvironment()
	code := m.Run()
	teardownTestEnvironment()
	os.Exit(code)
}

func setupTestEnvironment() {
	_ = os.Mkdir(migrationsTestDir, 0766)
}

func teardownTestEnvironment() {
	_ = os.RemoveAll(migrationsTestDir)
}

func TestRun_success(t *testing.T) {
	migrationFile := filepath.Join(migrationsTestDir, "001_init.sql")
	err := os.WriteFile(migrationFile, []byte("CREATE TABLE test (id INT);"), 0644)
	assert.NoError(t, err)
	defer func(name string) {
		err := os.Remove(name)
		if err != nil {
			t.Errorf("%v", err)
		}
	}(migrationFile)
	mockPostgresConnection := new(database.MockPostgresConnection)
	mockPostgresConnection.On("Exec", mock.Anything, mock.Anything, mock.Anything).Return(pgconn.CommandTag{}, nil)
	mockCreateMigrationsTable := migrations.MockCreateMigrationsTable(nil)
	mockIsMigrationApplied := migrations.MockIsMigrationApplied(false, nil)
	mockInsertAppliedMigration := migrations.MockInsertAppliedMigration(nil)

	run := migrations.MakeRun(mockPostgresConnection, mockCreateMigrationsTable, mockIsMigrationApplied, mockInsertAppliedMigration)

	got := run(context.Background(), migrationsTestDir)

	assert.NoError(t, got)
	mockPostgresConnection.AssertExpectations(t)
}

func TestRun_successWhenTheMigrationsAreAlreadyApplied(t *testing.T) {
	migrationFile := filepath.Join(migrationsTestDir, "001_init.sql")
	err := os.WriteFile(migrationFile, []byte("CREATE TABLE test (id INT);"), 0644)
	assert.NoError(t, err)
	defer func(name string) {
		err := os.Remove(name)
		if err != nil {
			t.Errorf("%v", err)
		}
	}(migrationFile)
	mockPostgresConnection := new(database.MockPostgresConnection)
	mockPostgresConnection.On("Exec", mock.Anything, mock.Anything, mock.Anything).Return(pgconn.CommandTag{}, nil)
	mockCreateMigrationsTable := migrations.MockCreateMigrationsTable(nil)
	mockIsMigrationApplied := migrations.MockIsMigrationApplied(true, nil)
	mockInsertAppliedMigration := migrations.MockInsertAppliedMigration(nil)

	run := migrations.MakeRun(mockPostgresConnection, mockCreateMigrationsTable, mockIsMigrationApplied, mockInsertAppliedMigration)

	got := run(context.Background(), migrationsTestDir)

	assert.NoError(t, got)
}

func TestRun_failsWhenExecuteSQLFromFileThrowsUnableToReadFileError(t *testing.T) {
	invalidFile := filepath.Join(migrationsTestDir, "invalid.sql")
	err := os.WriteFile(invalidFile, []byte("CREATE TABLE test (id INT);"), 0000)
	assert.NoError(t, err)
	defer func(name string) {
		err := os.Remove(name)
		if err != nil {
			t.Errorf("%v", err)
		}
	}(invalidFile)
	mockPostgresConnection := new(database.MockPostgresConnection)
	mockCreateMigrationsTable := migrations.MockCreateMigrationsTable(nil)
	mockIsMigrationApplied := migrations.MockIsMigrationApplied(false, nil)
	mockInsertAppliedMigration := migrations.MockInsertAppliedMigration(nil)

	run := migrations.MakeRun(mockPostgresConnection, mockCreateMigrationsTable, mockIsMigrationApplied, mockInsertAppliedMigration)

	want := migrations.FailedToExecuteMigration
	got := run(context.Background(), migrationsTestDir)

	assert.Equal(t, want, got)
	mockPostgresConnection.AssertExpectations(t)
}

func TestRun_failsWhenExecuteSQLFromFileThrowsUnableToExecuteSQLError(t *testing.T) {
	migrationFile := filepath.Join(migrationsTestDir, "001_init.sql")
	err := os.WriteFile(migrationFile, []byte("CREATE TABLE test (id INT);"), 0644)
	assert.NoError(t, err)
	defer func(name string) {
		err := os.Remove(name)
		if err != nil {
			t.Errorf("%v", err)
		}
	}(migrationFile)
	mockPostgresConnection := new(database.MockPostgresConnection)
	mockPostgresConnection.On("Exec", mock.Anything, mock.Anything, mock.Anything).Return(pgconn.CommandTag{}, migrations.UnableToExecuteSQL)
	mockCreateMigrationsTable := migrations.MockCreateMigrationsTable(nil)
	mockIsMigrationApplied := migrations.MockIsMigrationApplied(false, nil)
	mockInsertAppliedMigration := migrations.MockInsertAppliedMigration(nil)

	run := migrations.MakeRun(mockPostgresConnection, mockCreateMigrationsTable, mockIsMigrationApplied, mockInsertAppliedMigration)

	want := migrations.FailedToExecuteMigration
	got := run(context.Background(), migrationsTestDir)

	assert.Equal(t, want, got)
	mockPostgresConnection.AssertExpectations(t)
}

func TestRun_failsWhenCreateMigrationsTableThrowsError(t *testing.T) {
	migrationFile := filepath.Join(migrationsTestDir, "001_init.sql")
	err := os.WriteFile(migrationFile, []byte("CREATE TABLE test (id INT);"), 0644)
	assert.NoError(t, err)
	defer func(name string) {
		err := os.Remove(name)
		if err != nil {
			t.Errorf("%v", err)
		}
	}(migrationFile)
	mockPostgresConnection := new(database.MockPostgresConnection)
	mockCreateMigrationsTable := migrations.MockCreateMigrationsTable(migrations.FailedToCreateMigrationsTable)
	mockIsMigrationApplied := migrations.MockIsMigrationApplied(false, nil)
	mockInsertAppliedMigration := migrations.MockInsertAppliedMigration(nil)

	run := migrations.MakeRun(mockPostgresConnection, mockCreateMigrationsTable, mockIsMigrationApplied, mockInsertAppliedMigration)

	want := migrations.FailedToCreateMigrationsTable
	got := run(context.Background(), migrationsTestDir)

	assert.Equal(t, want, got)
	mockPostgresConnection.AssertExpectations(t)
}

func TestRun_failsWhenIsMigrationAppliedThrowsError(t *testing.T) {
	migrationFile := filepath.Join(migrationsTestDir, "001_init.sql")
	err := os.WriteFile(migrationFile, []byte("CREATE TABLE test (id INT);"), 0644)
	assert.NoError(t, err)
	defer func(name string) {
		err := os.Remove(name)
		if err != nil {
			t.Errorf("%v", err)
		}
	}(migrationFile)
	mockPostgresConnection := new(database.MockPostgresConnection)
	mockCreateMigrationsTable := migrations.MockCreateMigrationsTable(nil)
	mockIsMigrationApplied := migrations.MockIsMigrationApplied(false, migrations.FailedToRetrieveIfMigrationWasApplied)
	mockInsertAppliedMigration := migrations.MockInsertAppliedMigration(nil)

	run := migrations.MakeRun(mockPostgresConnection, mockCreateMigrationsTable, mockIsMigrationApplied, mockInsertAppliedMigration)

	want := migrations.FailedToRetrieveIfMigrationWasApplied
	got := run(context.Background(), migrationsTestDir)

	assert.Equal(t, want, got)
	mockPostgresConnection.AssertExpectations(t)
}

func TestRun_failsWhenInsertAppliedMigrationThrowsError(t *testing.T) {
	migrationFile := filepath.Join(migrationsTestDir, "001_init.sql")
	err := os.WriteFile(migrationFile, []byte("CREATE TABLE test (id INT);"), 0644)
	assert.NoError(t, err)
	defer func(name string) {
		err := os.Remove(name)
		if err != nil {
			t.Errorf("%v", err)
		}
	}(migrationFile)
	mockPostgresConnection := new(database.MockPostgresConnection)
	mockPostgresConnection.On("Exec", mock.Anything, mock.Anything, mock.Anything).Return(pgconn.CommandTag{}, nil)
	mockCreateMigrationsTable := migrations.MockCreateMigrationsTable(nil)
	mockIsMigrationApplied := migrations.MockIsMigrationApplied(false, nil)
	mockInsertAppliedMigration := migrations.MockInsertAppliedMigration(migrations.FailedToInsertAppliedMigration)

	run := migrations.MakeRun(mockPostgresConnection, mockCreateMigrationsTable, mockIsMigrationApplied, mockInsertAppliedMigration)

	want := migrations.FailedToInsertAppliedMigration
	got := run(context.Background(), migrationsTestDir)

	assert.Equal(t, want, got)
	mockPostgresConnection.AssertExpectations(t)
}
