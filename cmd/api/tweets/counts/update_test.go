package counts_test

import (
	"context"
	"errors"
	"testing"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"ahbcc/cmd/api/tweets/counts"
	"ahbcc/internal/database"
)

func TestUpdateTotalTweets_success(t *testing.T) {
	mockPostgresConnection := new(database.MockPostgresConnection)
	mockPostgresConnection.On("Exec", mock.Anything, mock.Anything, mock.Anything).Return(pgconn.CommandTag{}, nil)

	updateTotalTweets := counts.MakeUpdateTotalTweets(mockPostgresConnection)

	got := updateTotalTweets(context.Background(), 1, 1234567)

	assert.Nil(t, got)
	mockPostgresConnection.AssertExpectations(t)
}

func TestUpdateTotalTweets_failsWhenUpdateOperationThrowsError(t *testing.T) {
	mockPostgresConnection := new(database.MockPostgresConnection)
	mockPostgresConnection.On("Exec", mock.Anything, mock.Anything, mock.Anything).Return(pgconn.CommandTag{}, errors.New("failed to update execution"))

	updateTotalTweets := counts.MakeUpdateTotalTweets(mockPostgresConnection)

	want := counts.FailedToUpdateTotalTweets
	got := updateTotalTweets(context.Background(), 1, 1234567)

	assert.Equal(t, want, got)
	mockPostgresConnection.AssertExpectations(t)
}
