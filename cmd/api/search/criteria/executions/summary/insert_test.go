package summary_test

import (
	"context"
	"testing"

	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/lhbelfanti/corpus-creator/cmd/api/search/criteria/executions/summary"
	"github.com/lhbelfanti/corpus-creator/internal/database"
)

func TestInsert_success(t *testing.T) {
	executionSummaryID := 1
	mockPostgresConnection := new(database.MockPostgresConnection)
	mockPgxRow := new(database.MockPgxRow)
	database.MockScan(mockPgxRow, []any{executionSummaryID}, t)
	mockPostgresConnection.On("QueryRow", mock.Anything, mock.Anything, mock.Anything).Return(mockPgxRow)
	mockTweetsCountsDAO := summary.MockExecutionSummaryDAO(1, 2025, 1, 1000)

	insertExecutionSummary := summary.MakeInsert(mockPostgresConnection)

	want := executionSummaryID
	got, err := insertExecutionSummary(nil, context.Background(), mockTweetsCountsDAO)

	assert.Equal(t, want, got)
	assert.Nil(t, err)
	mockPostgresConnection.AssertExpectations(t)
}

func TestInsert_successWithATransaction(t *testing.T) {
	executionSummaryID := 1
	mockPostgresConnection := new(database.MockPostgresConnection)
	mockPostgresTx := new(database.MockPgxTx)
	mockPostgresConnection.On("Begin", mock.Anything).Return(mockPostgresTx, nil)
	mockPgxRow := new(database.MockPgxRow)
	database.MockScan(mockPgxRow, []any{executionSummaryID}, t)
	mockPostgresTx.On("QueryRow", mock.Anything, mock.Anything, mock.Anything).Return(mockPgxRow)
	mockTweetsCountsDAO := summary.MockExecutionSummaryDAO(1, 2025, 1, 1000)

	ctx := context.Background()
	tx, _ := mockPostgresConnection.Begin(ctx)
	insertExecutionSummary := summary.MakeInsert(mockPostgresConnection)

	want := executionSummaryID
	got, err := insertExecutionSummary(tx, ctx, mockTweetsCountsDAO)

	assert.Equal(t, want, got)
	assert.Nil(t, err)
	mockPostgresConnection.AssertExpectations(t)
}

func TestInsert_failsWhenInsertOperationThrowsError(t *testing.T) {
	mockPostgresConnection := new(database.MockPostgresConnection)
	mockPgxRow := new(database.MockPgxRow)
	mockPgxRow.On("Scan", mock.Anything).Return(pgx.ErrNoRows)
	mockPostgresConnection.On("QueryRow", mock.Anything, mock.Anything, mock.Anything).Return(mockPgxRow)
	mockTweetsCountsDAO := summary.MockExecutionSummaryDAO(1, 2025, 1, 1000)

	insertExecutionSummary := summary.MakeInsert(mockPostgresConnection)

	want := summary.FailedToInsertExecutionSummary
	_, got := insertExecutionSummary(nil, context.Background(), mockTweetsCountsDAO)

	assert.Equal(t, want, got)
	mockPostgresConnection.AssertExpectations(t)
	mockPgxRow.AssertExpectations(t)
}
