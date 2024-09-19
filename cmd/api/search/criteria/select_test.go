package criteria_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"ahbcc/cmd/api/search/criteria"
	"ahbcc/internal/database"
)

func TestSelectByID_success(t *testing.T) {
	mockPostgresConnection := new(database.MockPostgresConnection)
	mockPgxRow := new(database.MockPgxRow)
	mockCriteria := criteria.MockCriteriaDAO()
	database.MockScan[criteria.DAO](mockPgxRow, mockCriteria, t)
	mockPostgresConnection.On("QueryRow", mock.Anything, mock.Anything, mock.Anything).Return(mockPgxRow)

	selectCriteriaByID := criteria.MakeSelectByID(mockPostgresConnection)

	want := mockCriteria
	got, err := selectCriteriaByID(context.Background(), 1)

	assert.Nil(t, err)
	assert.Equal(t, want, got)
	mockPostgresConnection.AssertExpectations(t)
	mockPgxRow.AssertExpectations(t)
}

func TestSelectByID_failsWhenSelectOperationFails(t *testing.T) {
	tests := []struct {
		err      error
		expected error
	}{
		{err: pgx.ErrNoRows, expected: criteria.NoCriteriaDataFoundForTheGivenCriteriaID},
		{err: errors.New("failed to execute select operation"), expected: criteria.FailedExecuteQueryToRetrieveCriteriaData},
	}

	for _, tt := range tests {
		mockPostgresConnection := new(database.MockPostgresConnection)
		mockPgxRow := new(database.MockPgxRow)
		mockPgxRow.On("Scan", mock.Anything).Return(tt.err)
		mockPostgresConnection.On("QueryRow", mock.Anything, mock.Anything, mock.Anything).Return(mockPgxRow)

		selectCriteriaByID := criteria.MakeSelectByID(mockPostgresConnection)

		want := tt.expected
		_, got := selectCriteriaByID(context.Background(), 1)

		assert.Equal(t, want, got)
		mockPostgresConnection.AssertExpectations(t)
		mockPgxRow.AssertExpectations(t)
	}
}

func TestSelectAll_success(t *testing.T) {
	mockPostgresConnection := new(database.MockPostgresConnection)
	mockPgxRows := new(database.MockPgxRows)
	mockPgxRows.On("Close").Return()
	mockPostgresConnection.On("Query", mock.Anything, mock.Anything, mock.Anything).Return(mockPgxRows, nil)
	mockExecutionDAOSlice := criteria.MockCriteriaDAOSlice()
	mockCollectRows := database.MockCollectRows[criteria.DAO](mockExecutionDAOSlice, nil)

	selectAllCriteria := criteria.MakeSelectAll(mockPostgresConnection, mockCollectRows)

	want := mockExecutionDAOSlice
	got, err := selectAllCriteria(context.Background())

	assert.Nil(t, err)
	assert.Equal(t, want, got)
	mockPostgresConnection.AssertExpectations(t)
	mockPgxRows.AssertExpectations(t)
}

func TestSelectAll_failsWhenSelectOperationThrowsError(t *testing.T) {
	mockPostgresConnection := new(database.MockPostgresConnection)
	mockPgxRows := new(database.MockPgxRows)
	mockPostgresConnection.On("Query", mock.Anything, mock.Anything, mock.Anything).Return(mockPgxRows, errors.New("failed to select all criteria"))
	mockExecutionDAOSlice := criteria.MockCriteriaDAOSlice()
	mockCollectRows := database.MockCollectRows[criteria.DAO](mockExecutionDAOSlice, nil)

	selectAllCriteria := criteria.MakeSelectAll(mockPostgresConnection, mockCollectRows)

	want := criteria.FailedToRetrieveAllCriteriaData
	_, got := selectAllCriteria(context.Background())

	assert.Equal(t, want, got)
	mockPostgresConnection.AssertExpectations(t)
	mockPgxRows.AssertExpectations(t)
}

func TestSelectAll_failsWhenCollectRowsThrowsError(t *testing.T) {
	mockPostgresConnection := new(database.MockPostgresConnection)
	mockPgxRows := new(database.MockPgxRows)
	mockPgxRows.On("Close").Return()
	mockPostgresConnection.On("Query", mock.Anything, mock.Anything, mock.Anything).Return(mockPgxRows, nil)
	mockExecutionDAOSlice := criteria.MockCriteriaDAOSlice()
	mockCollectRows := database.MockCollectRows[criteria.DAO](mockExecutionDAOSlice, errors.New("failed to collect rows"))

	selectAllCriteria := criteria.MakeSelectAll(mockPostgresConnection, mockCollectRows)

	want := criteria.FailedToExecuteSelectCollectRowsInSelectAll
	_, got := selectAllCriteria(context.Background())

	assert.Equal(t, want, got)
	mockPostgresConnection.AssertExpectations(t)
	mockPgxRows.AssertExpectations(t)
}

func TestSelectExecutionsByState_success(t *testing.T) {
	mockPostgresConnection := new(database.MockPostgresConnection)
	mockPgxRows := new(database.MockPgxRows)
	mockPgxRows.On("Close").Return()
	mockPostgresConnection.On("Query", mock.Anything, mock.Anything, mock.Anything).Return(mockPgxRows, nil)
	mockExecutionDAOSlice := criteria.MockExecutionsDAO()
	mockCollectRows := database.MockCollectRows[criteria.ExecutionDAO](mockExecutionDAOSlice, nil)

	selectExecutionsByState := criteria.MakeSelectExecutionsByStatuses(mockPostgresConnection, mockCollectRows)

	want := mockExecutionDAOSlice
	got, err := selectExecutionsByState(context.Background(), []string{criteria.PendingStatus, criteria.InProgressStatus, criteria.DoneStatus})

	assert.Nil(t, err)
	assert.Equal(t, want, got)
	mockPostgresConnection.AssertExpectations(t)
	mockPgxRows.AssertExpectations(t)
}

func TestSelectExecutionsByStatuses_failsWhenSelectOperationThrowsError(t *testing.T) {
	mockPostgresConnection := new(database.MockPostgresConnection)
	mockPgxRows := new(database.MockPgxRows)
	mockPostgresConnection.On("Query", mock.Anything, mock.Anything, mock.Anything).Return(mockPgxRows, errors.New("failed to select executions by state"))
	mockExecutionDAOSlice := criteria.MockExecutionsDAO()
	mockCollectRows := database.MockCollectRows[criteria.ExecutionDAO](mockExecutionDAOSlice, nil)

	selectExecutionsByState := criteria.MakeSelectExecutionsByStatuses(mockPostgresConnection, mockCollectRows)

	want := criteria.FailedToExecuteSelectSearchCriteriaExecutionByState
	_, got := selectExecutionsByState(context.Background(), []string{criteria.PendingStatus, criteria.InProgressStatus, criteria.DoneStatus})

	assert.Equal(t, want, got)
	mockPostgresConnection.AssertExpectations(t)
	mockPgxRows.AssertExpectations(t)
}

func TestSelectExecutionsByStatuses_failsWhenCollectRowsThrowsError(t *testing.T) {
	mockPostgresConnection := new(database.MockPostgresConnection)
	mockPgxRows := new(database.MockPgxRows)
	mockPgxRows.On("Close").Return()
	mockPostgresConnection.On("Query", mock.Anything, mock.Anything, mock.Anything).Return(mockPgxRows, nil)
	mockCollectRows := database.MockCollectRows[criteria.ExecutionDAO](nil, errors.New("failed to collect rows"))

	selectExecutionsByState := criteria.MakeSelectExecutionsByStatuses(mockPostgresConnection, mockCollectRows)

	want := criteria.FailedToExecuteSelectCollectRowsInSelectExecutionByState
	_, got := selectExecutionsByState(context.Background(), []string{criteria.PendingStatus, criteria.InProgressStatus, criteria.DoneStatus})

	assert.Equal(t, want, got)
	mockPostgresConnection.AssertExpectations(t)
	mockPgxRows.AssertExpectations(t)
}

func TestSelectLastDayExecutedByCriteriaID_success(t *testing.T) {
	mockPostgresConnection := new(database.MockPostgresConnection)
	mockPgxRow := new(database.MockPgxRow)
	mockDate := time.Date(2024, 9, 19, 0, 0, 0, 0, time.UTC)
	database.MockScan[time.Time](mockPgxRow, mockDate, t)
	mockPostgresConnection.On("QueryRow", mock.Anything, mock.Anything, mock.Anything).Return(mockPgxRow)

	selectLastDayExecutedByCriteriaID := criteria.MakeSelectLastDayExecutedByCriteriaID(mockPostgresConnection)

	want := "2024-09-19"
	got, err := selectLastDayExecutedByCriteriaID(context.Background(), 1)

	assert.Nil(t, err)
	assert.Equal(t, want, got)
	mockPostgresConnection.AssertExpectations(t)
	mockPgxRow.AssertExpectations(t)
}

func TestSelectLastDayExecutedByCriteriaID_failsWhenSelectOperationFails(t *testing.T) {
	tests := []struct {
		err      error
		expected error
	}{
		{err: pgx.ErrNoRows, expected: criteria.NoExecutionDaysFoundForTheGivenCriteriaID},
		{err: errors.New("failed to execute select operation"), expected: criteria.FailedToRetrieveLastDayExecutedDate},
	}

	for _, tt := range tests {
		mockPostgresConnection := new(database.MockPostgresConnection)
		mockPgxRow := new(database.MockPgxRow)
		mockPgxRow.On("Scan", mock.Anything).Return(tt.err)
		mockPostgresConnection.On("QueryRow", mock.Anything, mock.Anything, mock.Anything).Return(mockPgxRow)

		selectLastDayExecutedByCriteriaID := criteria.MakeSelectLastDayExecutedByCriteriaID(mockPostgresConnection)

		want := tt.expected
		_, got := selectLastDayExecutedByCriteriaID(context.Background(), 1)

		assert.Equal(t, want, got)
		mockPostgresConnection.AssertExpectations(t)
		mockPgxRow.AssertExpectations(t)
	}
}
