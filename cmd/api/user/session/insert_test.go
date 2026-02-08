package session_test

import (
	"context"
	"errors"
	"testing"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/lhbelfanti/corpus-creator/cmd/api/user/session"
	"github.com/lhbelfanti/corpus-creator/internal/database"
)

func TestInsert_success(t *testing.T) {
	mockPostgresConnection := new(database.MockPostgresConnection)
	mockPostgresConnection.On("Exec", mock.Anything, mock.Anything, mock.Anything).Return(pgconn.CommandTag{}, nil)
	mockUserSessionDAO := session.MockUserSessionDAO()

	insertUserSession := session.MakeInsert(mockPostgresConnection)

	got := insertUserSession(context.Background(), mockUserSessionDAO)

	assert.Nil(t, got)
	mockPostgresConnection.AssertExpectations(t)
}

func TestInsert_failsWhenInsertOperationThrowsError(t *testing.T) {
	mockPostgresConnection := new(database.MockPostgresConnection)
	mockPostgresConnection.On("Exec", mock.Anything, mock.Anything, mock.Anything).Return(pgconn.CommandTag{}, errors.New("failed to insert user session"))
	mockUserSessionDAO := session.MockUserSessionDAO()

	insertUserSession := session.MakeInsert(mockPostgresConnection)

	want := session.FailedToInsertUserSession
	got := insertUserSession(context.Background(), mockUserSessionDAO)

	assert.Equal(t, want, got)
	mockPostgresConnection.AssertExpectations(t)
}
