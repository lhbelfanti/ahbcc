package user_test

import (
	"context"
	"errors"
	"testing"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/lhbelfanti/corpus-creator/cmd/api/user"
	"github.com/lhbelfanti/corpus-creator/internal/database"
)

func TestInsert_success(t *testing.T) {
	mockPostgresConnection := new(database.MockPostgresConnection)
	mockPostgresConnection.On("Exec", mock.Anything, mock.Anything, mock.Anything).Return(pgconn.CommandTag{}, nil)
	mockUserDTO := user.MockDTO()

	insertUser := user.MakeInsert(mockPostgresConnection)

	got := insertUser(context.Background(), mockUserDTO)

	assert.Nil(t, got)
	mockPostgresConnection.AssertExpectations(t)
}

func TestInsert_failsWhenInsertOperationThrowsError(t *testing.T) {
	mockPostgresConnection := new(database.MockPostgresConnection)
	mockPostgresConnection.On("Exec", mock.Anything, mock.Anything, mock.Anything).Return(pgconn.CommandTag{}, errors.New("failed to insert user"))
	mockUserDTO := user.MockDTO()

	insertUser := user.MakeInsert(mockPostgresConnection)

	want := user.FailedToInsertUser
	got := insertUser(context.Background(), mockUserDTO)

	assert.Equal(t, want, got)
	mockPostgresConnection.AssertExpectations(t)
}
