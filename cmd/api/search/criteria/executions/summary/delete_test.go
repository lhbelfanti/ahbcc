package summary_test

import (
	"context"
	"errors"
	"testing"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/lhbelfanti/corpus-creator/cmd/api/search/criteria/executions/summary"
	"github.com/lhbelfanti/corpus-creator/internal/database"
)

func TestDelete_success(t *testing.T) {
	mockPostgresConnection := new(database.MockPostgresConnection)
	mockPostgresConnection.On("Exec", mock.Anything, mock.Anything, mock.Anything).Return(pgconn.CommandTag{}, nil)

	deleteAllExecutionsSummary := summary.MakeDeleteAll(mockPostgresConnection)

	got := deleteAllExecutionsSummary(context.Background())

	assert.Nil(t, got)
	mockPostgresConnection.AssertExpectations(t)
}

func TestDelete_failsWhenDeleteOperationThrowsError(t *testing.T) {
	mockPostgresConnection := new(database.MockPostgresConnection)
	mockPostgresConnection.On("Exec", mock.Anything, mock.Anything, mock.Anything).Return(pgconn.CommandTag{}, errors.New("failed to delete executions summary"))

	deleteAllExecutionsSummary := summary.MakeDeleteAll(mockPostgresConnection)

	want := summary.FailedToDeleteAllSearchCriteriaExecutionsSummary
	got := deleteAllExecutionsSummary(context.Background())

	assert.Equal(t, want, got)
	mockPostgresConnection.AssertExpectations(t)
}
