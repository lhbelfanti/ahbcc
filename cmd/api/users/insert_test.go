package users_test

import (
	"context"
	"errors"
	"testing"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"ahbcc/cmd/api/users"
	"ahbcc/internal/database"
)

func TestInsert_success(t *testing.T) {
	mockPostgresConnection := new(database.MockPostgresConnection)
	mockPostgresConnection.On("Exec", mock.Anything, mock.Anything, mock.Anything).Return(pgconn.CommandTag{}, nil)
	mockUserDTO := users.MockUserDTO()

	insertUser := users.MakeInsert(mockPostgresConnection)

	got := insertUser(context.Background(), mockUserDTO)

	assert.Nil(t, got)
	mockPostgresConnection.AssertExpectations(t)
}

func TestInsert_failsWhenInsertOperationThrowsError(t *testing.T) {
	mockPostgresConnection := new(database.MockPostgresConnection)
	mockPostgresConnection.On("Exec", mock.Anything, mock.Anything, mock.Anything).Return(pgconn.CommandTag{}, errors.New("failed to insert user"))
	mockUserDTO := users.MockUserDTO()

	insertUser := users.MakeInsert(mockPostgresConnection)

	want := users.FailedToInsertUser
	got := insertUser(context.Background(), mockUserDTO)

	assert.Equal(t, want, got)
	mockPostgresConnection.AssertExpectations(t)
}
