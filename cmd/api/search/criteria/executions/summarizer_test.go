package executions_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"ahbcc/cmd/api/search/criteria/executions"
	"ahbcc/cmd/api/search/criteria/executions/summary"
	"ahbcc/internal/database"
)

func TestSummarizeExecutions_success(t *testing.T) {
	mockPostgresConnection := new(database.MockPostgresConnection)
	mockPostgresTx := new(database.MockPgxTx)
	mockPostgresTx.On("Rollback", mock.Anything).Return(nil)
	mockPostgresTx.On("Commit", mock.Anything).Return(nil)
	mockPostgresConnection.On("Begin", mock.Anything).Return(mockPostgresTx, nil)
	mockSelectExecutionsByStatuses := executions.MockSelectExecutionsByStatuses(executions.MockExecutionsDAO(), nil)
	mockSelectMonthlyTweetsCountsByYear := summary.MockSelectMonthlyTweetsCountsByYearByCriteriaID(summary.MockExecutionsSummaryDAOSlice(), nil)
	mockUpsert := summary.MockUpsert(nil)

	summarizeExecutions := executions.MakeSummarize(mockPostgresConnection, mockSelectExecutionsByStatuses, mockSelectMonthlyTweetsCountsByYear, mockUpsert)

	got := summarizeExecutions(context.Background())

	assert.Nil(t, got)
	mockPostgresConnection.AssertExpectations(t)
	mockPostgresTx.AssertExpectations(t)
}

func TestSummarizeExecutions_failsWhenSelectExecutionsByStatusesThrowsError(t *testing.T) {
	mockPostgresConnection := new(database.MockPostgresConnection)
	mockSelectExecutionsByStatuses := executions.MockSelectExecutionsByStatuses([]executions.ExecutionDAO{}, errors.New("failed to execute select executions by statuses"))
	mockSelectMonthlyTweetsCountsByYear := summary.MockSelectMonthlyTweetsCountsByYearByCriteriaID(summary.MockExecutionsSummaryDAOSlice(), nil)
	mockUpsert := summary.MockUpsert(nil)

	summarizeExecutions := executions.MakeSummarize(mockPostgresConnection, mockSelectExecutionsByStatuses, mockSelectMonthlyTweetsCountsByYear, mockUpsert)

	want := executions.FailedToExecuteSelectSearchCriteriaExecutionByState
	got := summarizeExecutions(context.Background())

	assert.Equal(t, want, got)
	mockPostgresConnection.AssertExpectations(t)
}

func TestSummarizeExecutions_failsWhenBeginTransactionThrowsError(t *testing.T) {
	mockPostgresConnection := new(database.MockPostgresConnection)
	mockPostgresTx := new(database.MockPgxTx)
	mockPostgresConnection.On("Begin", mock.Anything).Return(mockPostgresTx, errors.New("failed to begin transaction"))
	mockSelectExecutionsByStatuses := executions.MockSelectExecutionsByStatuses([]executions.ExecutionDAO{}, nil)
	mockSelectMonthlyTweetsCountsByYear := summary.MockSelectMonthlyTweetsCountsByYearByCriteriaID(summary.MockExecutionsSummaryDAOSlice(), nil)
	mockUpsert := summary.MockUpsert(nil)

	summarizeExecutions := executions.MakeSummarize(mockPostgresConnection, mockSelectExecutionsByStatuses, mockSelectMonthlyTweetsCountsByYear, mockUpsert)

	want := executions.FailedToBeginTransaction
	got := summarizeExecutions(context.Background())

	assert.Equal(t, want, got)
	mockPostgresConnection.AssertExpectations(t)
}

func TestSummarizeExecutions_failsWhenSelectMonthlyTweetsCountsByYearThrowsError(t *testing.T) {
	mockPostgresConnection := new(database.MockPostgresConnection)
	mockPostgresTx := new(database.MockPgxTx)
	mockPostgresTx.On("Rollback", mock.Anything).Return(nil)
	mockPostgresConnection.On("Begin", mock.Anything).Return(mockPostgresTx, nil)
	mockSelectExecutionsByStatuses := executions.MockSelectExecutionsByStatuses(executions.MockExecutionsDAO(), nil)
	mockSelectMonthlyTweetsCountsByYear := summary.MockSelectMonthlyTweetsCountsByYearByCriteriaID([]summary.DAO{}, errors.New("failed to execute select monthly tweets count by year"))
	mockUpsert := summary.MockUpsert(nil)

	summarizeExecutions := executions.MakeSummarize(mockPostgresConnection, mockSelectExecutionsByStatuses, mockSelectMonthlyTweetsCountsByYear, mockUpsert)

	want := executions.FailedToExecuteSelectMonthlyTweetsCountsByYear
	got := summarizeExecutions(context.Background())

	assert.Equal(t, want, got)
	mockPostgresConnection.AssertExpectations(t)
	mockPostgresTx.AssertExpectations(t)
}

func TestSummarizeExecutions_failsWhenUpsertExecutionSummaryThrowsError(t *testing.T) {
	mockPostgresConnection := new(database.MockPostgresConnection)
	mockPostgresTx := new(database.MockPgxTx)
	mockPostgresTx.On("Rollback", mock.Anything).Return(nil)
	mockPostgresConnection.On("Begin", mock.Anything).Return(mockPostgresTx, nil)
	mockSelectExecutionsByStatuses := executions.MockSelectExecutionsByStatuses(executions.MockExecutionsDAO(), nil)
	mockSelectMonthlyTweetsCountsByYear := summary.MockSelectMonthlyTweetsCountsByYearByCriteriaID(summary.MockExecutionsSummaryDAOSlice(), nil)
	mockUpsert := summary.MockUpsert(errors.New("failed to execute upsert execution summary"))

	summarizeExecutions := executions.MakeSummarize(mockPostgresConnection, mockSelectExecutionsByStatuses, mockSelectMonthlyTweetsCountsByYear, mockUpsert)

	want := executions.FailedToExecuteUpsertExecutionSummary
	got := summarizeExecutions(context.Background())

	assert.Equal(t, want, got)
	mockPostgresConnection.AssertExpectations(t)
	mockPostgresTx.AssertExpectations(t)
}

func TestSummarizeExecutions_failsWhenCommitTransactionThrowsError(t *testing.T) {
	mockPostgresConnection := new(database.MockPostgresConnection)
	mockPostgresTx := new(database.MockPgxTx)
	mockPostgresTx.On("Rollback", mock.Anything).Return(nil)
	mockPostgresTx.On("Commit", mock.Anything).Return(errors.New("failed to commit transaction"))
	mockPostgresConnection.On("Begin", mock.Anything).Return(mockPostgresTx, nil)
	mockSelectExecutionsByStatuses := executions.MockSelectExecutionsByStatuses(executions.MockExecutionsDAO(), nil)
	mockSelectMonthlyTweetsCountsByYear := summary.MockSelectMonthlyTweetsCountsByYearByCriteriaID(summary.MockExecutionsSummaryDAOSlice(), nil)
	mockUpsert := summary.MockUpsert(nil)

	summarizeExecutions := executions.MakeSummarize(mockPostgresConnection, mockSelectExecutionsByStatuses, mockSelectMonthlyTweetsCountsByYear, mockUpsert)

	want := executions.FailedToCommitTransaction
	got := summarizeExecutions(context.Background())

	assert.Equal(t, want, got)
	mockPostgresConnection.AssertExpectations(t)
	mockPostgresTx.AssertExpectations(t)
}
