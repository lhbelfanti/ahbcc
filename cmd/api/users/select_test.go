package users_test

import (
	"context"
	"testing"

	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"ahbcc/cmd/api/users"
	"ahbcc/internal/database"
)

func TestUserExists_success(t *testing.T) {
	mockPostgresConnection := new(database.MockPostgresConnection)
	mockPgxRow := new(database.MockPgxRow)
	database.MockScan(mockPgxRow, []any{true}, t)
	mockPostgresConnection.On("QueryRow", mock.Anything, mock.Anything, mock.Anything).Return(mockPgxRow)

	userExists := users.MakeUserExists(mockPostgresConnection)

	got, err := userExists(context.Background(), "user")

	assert.Nil(t, err)
	assert.True(t, got)
	mockPostgresConnection.AssertExpectations(t)
	mockPgxRow.AssertExpectations(t)
}

func TestUserExists_failsWhenSelectOperationFails(t *testing.T) {
	mockPostgresConnection := new(database.MockPostgresConnection)
	mockPgxRow := new(database.MockPgxRow)
	mockPgxRow.On("Scan", mock.Anything).Return(pgx.ErrNoRows)
	mockPostgresConnection.On("QueryRow", mock.Anything, mock.Anything, mock.Anything).Return(mockPgxRow)

	userExists := users.MakeUserExists(mockPostgresConnection)

	want := users.FailedToRetrieveIfUserAlreadyExists
	_, got := userExists(context.Background(), "user")

	assert.Equal(t, want, got)
	mockPostgresConnection.AssertExpectations(t)
	mockPgxRow.AssertExpectations(t)
}
