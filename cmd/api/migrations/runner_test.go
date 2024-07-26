package migrations_test

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"ahbcc/cmd/api/migrations"
	"ahbcc/internal/database"
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

	run := migrations.MakeRun(mockPostgresConnection)

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
	mockPostgresConnection.On("Exec", mock.Anything, mock.Anything, mock.Anything).Return(pgconn.CommandTag{}, nil)

	run := migrations.MakeRun(mockPostgresConnection)

	want := migrations.FailedToExecuteMigration
	got := run(context.Background(), migrationsTestDir)

	assert.Equal(t, want, got)
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

	run := migrations.MakeRun(mockPostgresConnection)

	want := migrations.FailedToExecuteMigration
	got := run(context.Background(), migrationsTestDir)

	assert.Equal(t, want, got)
}
