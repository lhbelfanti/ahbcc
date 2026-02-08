package migrations_test

import (
	"context"
	"errors"
	"testing"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/lhbelfanti/corpus-creator/cmd/api/migrations"
	"github.com/lhbelfanti/corpus-creator/internal/database"
)

func TestCreateMigrationsTable_success(t *testing.T) {
	mockPostgresConnection := new(database.MockPostgresConnection)
	mockPostgresConnection.On("Exec", mock.Anything, mock.Anything, mock.Anything).Return(pgconn.CommandTag{}, nil)

	createMigrationsTable := migrations.MakeCreateMigrationsTable(mockPostgresConnection)

	got := createMigrationsTable(context.Background())

	assert.NoError(t, got)
	mockPostgresConnection.AssertExpectations(t)
}

func TestCreateMigrationsTable_failsWhenTableCreationThrowsError(t *testing.T) {
	mockPostgresConnection := new(database.MockPostgresConnection)
	mockPostgresConnection.On("Exec", mock.Anything, mock.Anything, mock.Anything).Return(pgconn.CommandTag{}, errors.New("failed to create migrations table"))

	createMigrationsTable := migrations.MakeCreateMigrationsTable(mockPostgresConnection)

	want := migrations.FailedToCreateMigrationsTable
	got := createMigrationsTable(context.Background())

	assert.Equal(t, want, got)
	mockPostgresConnection.AssertExpectations(t)
}

func TestIsMigrationApplied_success(t *testing.T) {
	mockPostgresConnection := new(database.MockPostgresConnection)
	mockPgxRow := new(database.MockPgxRow)
	database.MockScan(mockPgxRow, []any{true}, t)
	mockPostgresConnection.On("QueryRow", mock.Anything, mock.Anything, mock.Anything).Return(mockPgxRow)

	isMigrationApplied := migrations.MakeIsMigrationApplied(mockPostgresConnection)

	got, err := isMigrationApplied(context.Background(), "test")

	assert.Nil(t, err)
	assert.True(t, got)
	mockPostgresConnection.AssertExpectations(t)
	mockPgxRow.AssertExpectations(t)
}

func TestIsMigrationApplied_failsWhenSelectOperationFails(t *testing.T) {
	mockPostgresConnection := new(database.MockPostgresConnection)
	mockPgxRow := new(database.MockPgxRow)
	mockPgxRow.On("Scan", mock.Anything).Return(pgx.ErrNoRows)
	mockPostgresConnection.On("QueryRow", mock.Anything, mock.Anything, mock.Anything).Return(mockPgxRow)

	isMigrationApplied := migrations.MakeIsMigrationApplied(mockPostgresConnection)

	want := migrations.FailedToRetrieveIfMigrationWasApplied
	_, got := isMigrationApplied(context.Background(), "test")

	assert.Equal(t, want, got)
	mockPostgresConnection.AssertExpectations(t)
	mockPgxRow.AssertExpectations(t)
}

func TestInsertAppliedMigration_success(t *testing.T) {
	mockPostgresConnection := new(database.MockPostgresConnection)
	mockPostgresConnection.On("Exec", mock.Anything, mock.Anything, mock.Anything).Return(pgconn.CommandTag{}, nil)

	insertAppliedMigration := migrations.MakeInsertAppliedMigration(mockPostgresConnection)

	got := insertAppliedMigration(context.Background(), "test")

	assert.NoError(t, got)
	mockPostgresConnection.AssertExpectations(t)
}

func TestInsertAppliedMigration_failsWhenInsertOperationThrowsError(t *testing.T) {
	mockPostgresConnection := new(database.MockPostgresConnection)
	mockPostgresConnection.On("Exec", mock.Anything, mock.Anything, mock.Anything).Return(pgconn.CommandTag{}, errors.New("failed to insert applied migration"))

	insertAppliedMigration := migrations.MakeInsertAppliedMigration(mockPostgresConnection)

	want := migrations.FailedToInsertAppliedMigration
	got := insertAppliedMigration(context.Background(), "test")

	assert.Equal(t, want, got)
	mockPostgresConnection.AssertExpectations(t)
}
