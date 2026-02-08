package executions_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/lhbelfanti/corpus-creator/cmd/api/search/criteria/executions"
	"github.com/lhbelfanti/corpus-creator/cmd/api/search/criteria/executions/summary"
	"github.com/lhbelfanti/corpus-creator/internal/database"
)

func TestSummarizeExecutions_success(t *testing.T) {
	mockPostgresConnection := new(database.MockPostgresConnection)
	mockPostgresTx := new(database.MockPgxTx)
	mockPostgresTx.On("Rollback", mock.Anything).Return(nil)
	mockPostgresTx.On("Commit", mock.Anything).Return(nil)
	mockPostgresConnection.On("Begin", mock.Anything).Return(mockPostgresTx, nil)
	mockSelectExecutionsByStatuses := executions.MockSelectExecutionsByStatuses(executions.MockExecutionsDAO(), nil)
	mockDeleteAllExecutionsSummary := summary.MockDeleteAll(nil)
	mockSelectMonthlyTweetsCountsByYear := summary.MockSelectMonthlyTweetsCountsByYearByCriteriaID(summary.MockExecutionsSummaryDAOSlice(), nil)
	mockInsertExecutionSummary := summary.MockInsert(nil)

	summarizeExecutions := executions.MakeSummarize(mockPostgresConnection, mockSelectExecutionsByStatuses, mockDeleteAllExecutionsSummary, mockSelectMonthlyTweetsCountsByYear, mockInsertExecutionSummary)

	got := summarizeExecutions(context.Background())

	assert.Nil(t, got)
	mockPostgresConnection.AssertExpectations(t)
	mockPostgresTx.AssertExpectations(t)
}

func TestSummarizeExecutions_failsWhenSelectExecutionsByStatusesThrowsError(t *testing.T) {
	mockPostgresConnection := new(database.MockPostgresConnection)
	mockSelectExecutionsByStatuses := executions.MockSelectExecutionsByStatuses([]executions.ExecutionDAO{}, errors.New("failed to execute select executions by statuses"))
	mockDeleteAllExecutionsSummary := summary.MockDeleteAll(nil)
	mockSelectMonthlyTweetsCountsByYear := summary.MockSelectMonthlyTweetsCountsByYearByCriteriaID(summary.MockExecutionsSummaryDAOSlice(), nil)
	mockInsertExecutionSummary := summary.MockInsert(nil)

	summarizeExecutions := executions.MakeSummarize(mockPostgresConnection, mockSelectExecutionsByStatuses, mockDeleteAllExecutionsSummary, mockSelectMonthlyTweetsCountsByYear, mockInsertExecutionSummary)

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
	mockDeleteAllExecutionsSummary := summary.MockDeleteAll(nil)
	mockSelectMonthlyTweetsCountsByYear := summary.MockSelectMonthlyTweetsCountsByYearByCriteriaID(summary.MockExecutionsSummaryDAOSlice(), nil)
	mockInsertExecutionSummary := summary.MockInsert(nil)

	summarizeExecutions := executions.MakeSummarize(mockPostgresConnection, mockSelectExecutionsByStatuses, mockDeleteAllExecutionsSummary, mockSelectMonthlyTweetsCountsByYear, mockInsertExecutionSummary)

	want := executions.FailedToBeginTransaction
	got := summarizeExecutions(context.Background())

	assert.Equal(t, want, got)
	mockPostgresConnection.AssertExpectations(t)
}

func TestSummarizeExecutions_failsWhenDeleteAllExecutionsSummaryThrowsError(t *testing.T) {
	mockPostgresConnection := new(database.MockPostgresConnection)
	mockPostgresTx := new(database.MockPgxTx)
	mockPostgresTx.On("Rollback", mock.Anything).Return(nil)
	mockPostgresConnection.On("Begin", mock.Anything).Return(mockPostgresTx, nil)
	mockSelectExecutionsByStatuses := executions.MockSelectExecutionsByStatuses(executions.MockExecutionsDAO(), nil)
	mockDeleteAllExecutionsSummary := summary.MockDeleteAll(errors.New("failed to execute delete all executions summary"))
	mockSelectMonthlyTweetsCountsByYear := summary.MockSelectMonthlyTweetsCountsByYearByCriteriaID(summary.MockExecutionsSummaryDAOSlice(), nil)
	mockInsertExecutionSummary := summary.MockInsert(nil)

	summarizeExecutions := executions.MakeSummarize(mockPostgresConnection, mockSelectExecutionsByStatuses, mockDeleteAllExecutionsSummary, mockSelectMonthlyTweetsCountsByYear, mockInsertExecutionSummary)

	want := executions.FailedToClearOldSummary
	got := summarizeExecutions(context.Background())

	assert.Equal(t, want, got)
	mockPostgresConnection.AssertExpectations(t)
	mockPostgresTx.AssertExpectations(t)
}

func TestSummarizeExecutions_failsWhenSelectMonthlyTweetsCountsByYearThrowsError(t *testing.T) {
	mockPostgresConnection := new(database.MockPostgresConnection)
	mockPostgresTx := new(database.MockPgxTx)
	mockPostgresTx.On("Rollback", mock.Anything).Return(nil)
	mockPostgresConnection.On("Begin", mock.Anything).Return(mockPostgresTx, nil)
	mockSelectExecutionsByStatuses := executions.MockSelectExecutionsByStatuses(executions.MockExecutionsDAO(), nil)
	mockDeleteAllExecutionsSummary := summary.MockDeleteAll(nil)
	mockSelectMonthlyTweetsCountsByYear := summary.MockSelectMonthlyTweetsCountsByYearByCriteriaID([]summary.DAO{}, errors.New("failed to execute select monthly tweets count by year"))
	mockInsertExecutionSummary := summary.MockInsert(nil)

	summarizeExecutions := executions.MakeSummarize(mockPostgresConnection, mockSelectExecutionsByStatuses, mockDeleteAllExecutionsSummary, mockSelectMonthlyTweetsCountsByYear, mockInsertExecutionSummary)

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
	mockDeleteAllExecutionsSummary := summary.MockDeleteAll(nil)
	mockSelectMonthlyTweetsCountsByYear := summary.MockSelectMonthlyTweetsCountsByYearByCriteriaID(summary.MockExecutionsSummaryDAOSlice(), nil)
	mockInsertExecutionSummary := summary.MockInsert(errors.New("failed to execute insert execution summary"))

	summarizeExecutions := executions.MakeSummarize(mockPostgresConnection, mockSelectExecutionsByStatuses, mockDeleteAllExecutionsSummary, mockSelectMonthlyTweetsCountsByYear, mockInsertExecutionSummary)

	want := executions.FailedToExecuteInsertExecutionSummary
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
	mockDeleteAllExecutionsSummary := summary.MockDeleteAll(nil)
	mockSelectMonthlyTweetsCountsByYear := summary.MockSelectMonthlyTweetsCountsByYearByCriteriaID(summary.MockExecutionsSummaryDAOSlice(), nil)
	mockInsertExecutionSummary := summary.MockInsert(nil)

	summarizeExecutions := executions.MakeSummarize(mockPostgresConnection, mockSelectExecutionsByStatuses, mockDeleteAllExecutionsSummary, mockSelectMonthlyTweetsCountsByYear, mockInsertExecutionSummary)

	want := executions.FailedToCommitTransaction
	got := summarizeExecutions(context.Background())

	assert.Equal(t, want, got)
	mockPostgresConnection.AssertExpectations(t)
	mockPostgresTx.AssertExpectations(t)
}
