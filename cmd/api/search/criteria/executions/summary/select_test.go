package summary_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/lhbelfanti/corpus-creator/cmd/api/search/criteria/executions/summary"
	"github.com/lhbelfanti/corpus-creator/internal/database"
)

func TestSelectMonthlyTweetsCountsByYearByCriteriaID_success(t *testing.T) {
	mockPostgresConnection := new(database.MockPostgresConnection)
	mockPgxRows := new(database.MockPgxRows)
	mockPostgresConnection.On("Query", mock.Anything, mock.Anything, mock.Anything).Return(mockPgxRows, nil)
	mockExecutionsSummaryDAOSlice := summary.MockExecutionsSummaryDAOSlice()
	mockCollectRows := database.MockCollectRows[summary.DAO](mockExecutionsSummaryDAOSlice, nil)

	selectMonthlyTweetsCountsByYearByCriteriaID := summary.MakeSelectMonthlyTweetsCountsByYearByCriteriaID(mockPostgresConnection, mockCollectRows)

	want := mockExecutionsSummaryDAOSlice
	got, err := selectMonthlyTweetsCountsByYearByCriteriaID(context.Background(), 1)

	assert.Nil(t, err)
	assert.Equal(t, want, got)
	mockPostgresConnection.AssertExpectations(t)
	mockPgxRows.AssertExpectations(t)
}

func TestSelectMonthlyTweetsCountsByYearByCriteriaID_failsWhenSelectOperationThrowsError(t *testing.T) {
	mockPostgresConnection := new(database.MockPostgresConnection)
	mockPgxRows := new(database.MockPgxRows)
	mockPostgresConnection.On("Query", mock.Anything, mock.Anything, mock.Anything).Return(mockPgxRows, errors.New("failed to select monthly tweets counts by year"))
	mockExecutionsSummaryDAOSlice := summary.MockExecutionsSummaryDAOSlice()
	mockCollectRows := database.MockCollectRows[summary.DAO](mockExecutionsSummaryDAOSlice, nil)

	selectMonthlyTweetsCountsByYearByCriteriaID := summary.MakeSelectMonthlyTweetsCountsByYearByCriteriaID(mockPostgresConnection, mockCollectRows)

	want := summary.FailedToRetrieveMonthlyTweetsCountsByYear
	_, got := selectMonthlyTweetsCountsByYearByCriteriaID(context.Background(), 1)

	assert.Equal(t, want, got)
	mockPostgresConnection.AssertExpectations(t)
	mockPgxRows.AssertExpectations(t)
}

func TestSelectMonthlyTweetsCountsByYearByCriteriaID_failsWhenCollectRowsThrowsError(t *testing.T) {
	mockPostgresConnection := new(database.MockPostgresConnection)
	mockPgxRows := new(database.MockPgxRows)
	mockPostgresConnection.On("Query", mock.Anything, mock.Anything, mock.Anything).Return(mockPgxRows, nil)
	mockExecutionsSummaryDAOSlice := summary.MockExecutionsSummaryDAOSlice()
	mockCollectRows := database.MockCollectRows[summary.DAO](mockExecutionsSummaryDAOSlice, errors.New("failed to collect rows"))

	selectMonthlyTweetsCountsByYearByCriteriaID := summary.MakeSelectMonthlyTweetsCountsByYearByCriteriaID(mockPostgresConnection, mockCollectRows)

	want := summary.FailedToExecuteCollectRowsInSelectMonthlyTweetsCountsByYear
	_, got := selectMonthlyTweetsCountsByYearByCriteriaID(context.Background(), 1)

	assert.Equal(t, want, got)
	mockPostgresConnection.AssertExpectations(t)
	mockPgxRows.AssertExpectations(t)
}

func TestSelectAll_success(t *testing.T) {
	mockPostgresConnection := new(database.MockPostgresConnection)
	mockPgxRows := new(database.MockPgxRows)
	mockPostgresConnection.On("Query", mock.Anything, mock.Anything, mock.Anything).Return(mockPgxRows, nil)
	mockExecutionsSummaryDAOSlice := summary.MockExecutionsSummaryDAOSlice()
	mockCollectRows := database.MockCollectRows[summary.DAO](mockExecutionsSummaryDAOSlice, nil)

	selectAll := summary.MakeSelectAll(mockPostgresConnection, mockCollectRows)

	want := mockExecutionsSummaryDAOSlice
	got, err := selectAll(context.Background())

	assert.Nil(t, err)
	assert.Equal(t, want, got)
	mockPostgresConnection.AssertExpectations(t)
	mockPgxRows.AssertExpectations(t)
}

func TestSelectAll_failsWhenSelectOperationThrowsError(t *testing.T) {
	mockPostgresConnection := new(database.MockPostgresConnection)
	mockPgxRows := new(database.MockPgxRows)
	mockPostgresConnection.On("Query", mock.Anything, mock.Anything, mock.Anything).Return(mockPgxRows, errors.New("failed to select all the executions summary by criteria id"))
	mockExecutionsSummaryDAOSlice := summary.MockExecutionsSummaryDAOSlice()
	mockCollectRows := database.MockCollectRows[summary.DAO](mockExecutionsSummaryDAOSlice, nil)

	selectAll := summary.MakeSelectAll(mockPostgresConnection, mockCollectRows)

	want := summary.FailedToRetrieveExecutionsSummary
	_, got := selectAll(context.Background())

	assert.Equal(t, want, got)
	mockPostgresConnection.AssertExpectations(t)
	mockPgxRows.AssertExpectations(t)
}

func TestSelectAll_failsWhenCollectRowsThrowsError(t *testing.T) {
	mockPostgresConnection := new(database.MockPostgresConnection)
	mockPgxRows := new(database.MockPgxRows)
	mockPostgresConnection.On("Query", mock.Anything, mock.Anything, mock.Anything).Return(mockPgxRows, nil)
	mockExecutionsSummaryDAOSlice := summary.MockExecutionsSummaryDAOSlice()
	mockCollectRows := database.MockCollectRows[summary.DAO](mockExecutionsSummaryDAOSlice, errors.New("failed to collect rows"))

	selectAll := summary.MakeSelectAll(mockPostgresConnection, mockCollectRows)

	want := summary.FailedToExecuteCollectRowsInSelectAll
	_, got := selectAll(context.Background())

	assert.Equal(t, want, got)
	mockPostgresConnection.AssertExpectations(t)
	mockPgxRows.AssertExpectations(t)
}
