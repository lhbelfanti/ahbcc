package session_test

import (
	"context"
	"errors"
	"testing"

	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/lhbelfanti/corpus-creator/cmd/api/user/session"
	"github.com/lhbelfanti/corpus-creator/internal/database"
)

func TestSelectUserIDByToken_success(t *testing.T) {
	mockPostgresConnection := new(database.MockPostgresConnection)
	mockPgxRow := new(database.MockPgxRow)
	database.MockScan(mockPgxRow, []any{1}, t)
	mockPostgresConnection.On("QueryRow", mock.Anything, mock.Anything, mock.Anything).Return(mockPgxRow)

	selectUserIDByToken := session.MakeSelectUserIDByToken(mockPostgresConnection)

	want := 1
	got, err := selectUserIDByToken(context.Background(), "token")

	assert.Nil(t, err)
	assert.Equal(t, want, got)
	mockPostgresConnection.AssertExpectations(t)
	mockPgxRow.AssertExpectations(t)
}

func TestSelectUserIDByToken_failsWhenSelectOperationThrowsError(t *testing.T) {
	tests := []struct {
		err      error
		expected error
	}{
		{err: pgx.ErrNoRows, expected: session.NoUserIDFoundForTheGivenToken},
		{err: errors.New("failed to execute select operation"), expected: session.FailedToExecuteQueryToRetrieveUserID},
	}

	for _, tt := range tests {
		mockPostgresConnection := new(database.MockPostgresConnection)
		mockPgxRow := new(database.MockPgxRow)
		mockPgxRow.On("Scan", mock.Anything).Return(tt.err)
		mockPostgresConnection.On("QueryRow", mock.Anything, mock.Anything, mock.Anything).Return(mockPgxRow)

		selectUserIDByToken := session.MakeSelectUserIDByToken(mockPostgresConnection)

		want := tt.expected
		_, got := selectUserIDByToken(context.Background(), "token")

		assert.Equal(t, want, got)
		mockPostgresConnection.AssertExpectations(t)
		mockPgxRow.AssertExpectations(t)
	}
}
