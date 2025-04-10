package summary_test

import (
	"context"
	"errors"
	"testing"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"ahbcc/cmd/api/search/criteria/executions/summary"
	"ahbcc/internal/database"
)

func TestUpdateTotalTweets_success(t *testing.T) {
	mockPostgresConnection := new(database.MockPostgresConnection)
	mockPostgresConnection.On("Exec", mock.Anything, mock.Anything, mock.Anything).Return(pgconn.CommandTag{}, nil)

	updateTotalTweets := summary.MakeUpdateTotalTweets(mockPostgresConnection)

	got := updateTotalTweets(nil, context.Background(), 1, 1234567)

	assert.Nil(t, got)
	mockPostgresConnection.AssertExpectations(t)
}

func TestUpdateTotalTweets_successWithATransaction(t *testing.T) {
	mockPostgresConnection := new(database.MockPostgresConnection)
	mockPostgresTx := new(database.MockPgxTx)
	mockPostgresConnection.On("Begin", mock.Anything).Return(mockPostgresTx, nil)
	mockPostgresTx.On("Exec", mock.Anything, mock.Anything, mock.Anything).Return(pgconn.CommandTag{}, nil)

	ctx := context.Background()
	tx, _ := mockPostgresConnection.Begin(ctx)
	updateTotalTweets := summary.MakeUpdateTotalTweets(mockPostgresConnection)

	got := updateTotalTweets(tx, ctx, 1, 1234567)

	assert.Nil(t, got)
	mockPostgresConnection.AssertExpectations(t)
}

func TestUpdateTotalTweets_failsWhenUpdateOperationThrowsError(t *testing.T) {
	mockPostgresConnection := new(database.MockPostgresConnection)
	mockPostgresConnection.On("Exec", mock.Anything, mock.Anything, mock.Anything).Return(pgconn.CommandTag{}, errors.New("failed to update execution"))

	updateTotalTweets := summary.MakeUpdateTotalTweets(mockPostgresConnection)

	want := summary.FailedToUpdateTotalTweets
	got := updateTotalTweets(nil, context.Background(), 1, 1234567)

	assert.Equal(t, want, got)
	mockPostgresConnection.AssertExpectations(t)
}
