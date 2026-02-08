package user_test

import (
	"context"
	"errors"
	"testing"

	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/lhbelfanti/corpus-creator/cmd/api/user"
	"github.com/lhbelfanti/corpus-creator/internal/database"
)

func TestExists_success(t *testing.T) {
	mockPostgresConnection := new(database.MockPostgresConnection)
	mockPgxRow := new(database.MockPgxRow)
	database.MockScan(mockPgxRow, []any{true}, t)
	mockPostgresConnection.On("QueryRow", mock.Anything, mock.Anything, mock.Anything).Return(mockPgxRow)

	exists := user.MakeExists(mockPostgresConnection)

	got, err := exists(context.Background(), "user")

	assert.Nil(t, err)
	assert.True(t, got)
	mockPostgresConnection.AssertExpectations(t)
	mockPgxRow.AssertExpectations(t)
}

func TestExists_failsWhenSelectOperationFails(t *testing.T) {
	mockPostgresConnection := new(database.MockPostgresConnection)
	mockPgxRow := new(database.MockPgxRow)
	mockPgxRow.On("Scan", mock.Anything).Return(pgx.ErrNoRows)
	mockPostgresConnection.On("QueryRow", mock.Anything, mock.Anything, mock.Anything).Return(mockPgxRow)

	exists := user.MakeExists(mockPostgresConnection)

	want := user.FailedToRetrieveIfUserAlreadyExists
	_, got := exists(context.Background(), "user")

	assert.Equal(t, want, got)
	mockPostgresConnection.AssertExpectations(t)
	mockPgxRow.AssertExpectations(t)
}

func TestSelectByUsername_success(t *testing.T) {
	mockPostgresConnection := new(database.MockPostgresConnection)
	mockPgxRow := new(database.MockPgxRow)
	mockUser := user.MockDAO()
	mockScanCriteriaDAOValues := user.MockScanUserDAOValues(mockUser)
	database.MockScan(mockPgxRow, mockScanCriteriaDAOValues, t)
	mockPostgresConnection.On("QueryRow", mock.Anything, mock.Anything, mock.Anything).Return(mockPgxRow)

	selectUserByUsername := user.MakeSelectByUsername(mockPostgresConnection)

	want := mockUser
	got, err := selectUserByUsername(context.Background(), "user")

	assert.Nil(t, err)
	assert.Equal(t, want, got)
	mockPostgresConnection.AssertExpectations(t)
	mockPgxRow.AssertExpectations(t)
}

func TestSelectByUsername_failsWhenSelectOperationFails(t *testing.T) {
	tests := []struct {
		err      error
		expected error
	}{
		{err: pgx.ErrNoRows, expected: user.NoUserFoundForTheGivenUsername},
		{err: errors.New("failed to execute select operation"), expected: user.FailedExecuteQueryToRetrieveUser},
	}

	for _, tt := range tests {
		mockPostgresConnection := new(database.MockPostgresConnection)
		mockPgxRow := new(database.MockPgxRow)
		mockPgxRow.On("Scan", mock.Anything).Return(tt.err)
		mockPostgresConnection.On("QueryRow", mock.Anything, mock.Anything, mock.Anything).Return(mockPgxRow)

		selectUserByUsername := user.MakeSelectByUsername(mockPostgresConnection)

		want := tt.expected
		_, got := selectUserByUsername(context.Background(), "user")

		assert.Equal(t, want, got)
		mockPostgresConnection.AssertExpectations(t)
		mockPgxRow.AssertExpectations(t)
	}
}
