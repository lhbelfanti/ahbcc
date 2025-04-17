package summary_test

import (
	"context"
	"errors"
	"testing"

	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"ahbcc/cmd/api/search/criteria/executions/summary"
	"ahbcc/internal/database"
)

func TestSelectIDBySearchCriteriaIDYearAndMonth_success(t *testing.T) {
	mockPostgresConnection := new(database.MockPostgresConnection)
	mockPgxRow := new(database.MockPgxRow)
	database.MockScan(mockPgxRow, []any{1}, t)
	mockPostgresConnection.On("QueryRow", mock.Anything, mock.Anything, mock.Anything).Return(mockPgxRow)

	selectIDBySearchCriteriaIDYearAndMonth := summary.MakeSelectIDBySearchCriteriaIDYearAndMonth(mockPostgresConnection)

	want := 1
	got, err := selectIDBySearchCriteriaIDYearAndMonth(context.Background(), 1, 2025, 1)

	assert.Nil(t, err)
	assert.Equal(t, want, got)
	mockPostgresConnection.AssertExpectations(t)
	mockPgxRow.AssertExpectations(t)
}

func TestSelectIDBySearchCriteriaIDYearAndMonth_failsWhenSelectOperationThrowsError(t *testing.T) {
	tests := []struct {
		err      error
		expected error
	}{
		{err: pgx.ErrNoRows, expected: summary.NoExecutionSummaryFoundForTheGivenCriteria},
		{err: errors.New("failed to execute select operation"), expected: summary.FailedToExecuteQueryToRetrieveExecutionsSummary},
	}

	for _, tt := range tests {
		mockPostgresConnection := new(database.MockPostgresConnection)
		mockPgxRow := new(database.MockPgxRow)
		mockPgxRow.On("Scan", mock.Anything).Return(tt.err)
		mockPostgresConnection.On("QueryRow", mock.Anything, mock.Anything, mock.Anything).Return(mockPgxRow)

		selectIDBySearchCriteriaIDYearAndMonth := summary.MakeSelectIDBySearchCriteriaIDYearAndMonth(mockPostgresConnection)

		want := tt.expected
		_, got := selectIDBySearchCriteriaIDYearAndMonth(context.Background(), 1, 2025, 1)

		assert.Equal(t, want, got)
		mockPostgresConnection.AssertExpectations(t)
		mockPgxRow.AssertExpectations(t)
	}
}

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
