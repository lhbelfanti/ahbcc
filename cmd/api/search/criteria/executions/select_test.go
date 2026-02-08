package executions_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/lhbelfanti/corpus-creator/cmd/api/search/criteria/executions"
	"github.com/lhbelfanti/corpus-creator/internal/database"
)

func TestSelectExecutionByID_success(t *testing.T) {
	mockPostgresConnection := new(database.MockPostgresConnection)
	mockPgxRow := new(database.MockPgxRow)
	mockExecution := executions.MockExecutionDAO()
	mockScanCriteriaDAOValues := executions.MockExecutionDAOValues(mockExecution)
	database.MockScan(mockPgxRow, mockScanCriteriaDAOValues, t)
	mockPostgresConnection.On("QueryRow", mock.Anything, mock.Anything, mock.Anything).Return(mockPgxRow)

	selectExecutionByID := executions.MakeSelectExecutionByID(mockPostgresConnection)

	want := mockExecution
	got, err := selectExecutionByID(context.Background(), 1)

	assert.Nil(t, err)
	assert.Equal(t, want, got)
	mockPostgresConnection.AssertExpectations(t)
	mockPgxRow.AssertExpectations(t)
}

func TestSelectExecutionByID_failsWhenSelectOperationFails(t *testing.T) {
	tests := []struct {
		err      error
		expected error
	}{
		{err: pgx.ErrNoRows, expected: executions.NoExecutionFoundForTheGivenID},
		{err: errors.New("failed to execute select operation"), expected: executions.FailedToExecuteQueryToRetrieveExecutionData},
	}

	for _, tt := range tests {
		mockPostgresConnection := new(database.MockPostgresConnection)
		mockPgxRow := new(database.MockPgxRow)
		mockPgxRow.On("Scan", mock.Anything).Return(tt.err)
		mockPostgresConnection.On("QueryRow", mock.Anything, mock.Anything, mock.Anything).Return(mockPgxRow)

		selectExecutionByID := executions.MakeSelectExecutionByID(mockPostgresConnection)

		want := tt.expected
		_, got := selectExecutionByID(context.Background(), 1)

		assert.Equal(t, want, got)
		mockPostgresConnection.AssertExpectations(t)
		mockPgxRow.AssertExpectations(t)
	}
}

func TestSelectExecutionsByState_success(t *testing.T) {
	mockPostgresConnection := new(database.MockPostgresConnection)
	mockPgxRows := new(database.MockPgxRows)
	mockPostgresConnection.On("Query", mock.Anything, mock.Anything, mock.Anything).Return(mockPgxRows, nil)
	mockExecutionDAOSlice := executions.MockExecutionsDAO()
	mockCollectRows := database.MockCollectRows[executions.ExecutionDAO](mockExecutionDAOSlice, nil)

	selectExecutionsByState := executions.MakeSelectExecutionsByStatuses(mockPostgresConnection, mockCollectRows)

	want := mockExecutionDAOSlice
	got, err := selectExecutionsByState(context.Background(), []string{executions.PendingStatus, executions.InProgressStatus, executions.DoneStatus})

	assert.Nil(t, err)
	assert.Equal(t, want, got)
	mockPostgresConnection.AssertExpectations(t)
	mockPgxRows.AssertExpectations(t)
}

func TestSelectExecutionsByStatuses_failsWhenSelectOperationThrowsError(t *testing.T) {
	mockPostgresConnection := new(database.MockPostgresConnection)
	mockPgxRows := new(database.MockPgxRows)
	mockPostgresConnection.On("Query", mock.Anything, mock.Anything, mock.Anything).Return(mockPgxRows, errors.New("failed to select executions by state"))
	mockExecutionDAOSlice := executions.MockExecutionsDAO()
	mockCollectRows := database.MockCollectRows[executions.ExecutionDAO](mockExecutionDAOSlice, nil)

	selectExecutionsByState := executions.MakeSelectExecutionsByStatuses(mockPostgresConnection, mockCollectRows)

	want := executions.FailedToExecuteSelectSearchCriteriaExecutionByState
	_, got := selectExecutionsByState(context.Background(), []string{executions.PendingStatus, executions.InProgressStatus, executions.DoneStatus})

	assert.Equal(t, want, got)
	mockPostgresConnection.AssertExpectations(t)
	mockPgxRows.AssertExpectations(t)
}

func TestSelectExecutionsByStatuses_failsWhenCollectRowsThrowsError(t *testing.T) {
	mockPostgresConnection := new(database.MockPostgresConnection)
	mockPgxRows := new(database.MockPgxRows)
	mockPostgresConnection.On("Query", mock.Anything, mock.Anything, mock.Anything).Return(mockPgxRows, nil)
	mockCollectRows := database.MockCollectRows[executions.ExecutionDAO](nil, errors.New("failed to collect rows"))

	selectExecutionsByState := executions.MakeSelectExecutionsByStatuses(mockPostgresConnection, mockCollectRows)

	want := executions.FailedToExecuteCollectRowsInSelectExecutionByState
	_, got := selectExecutionsByState(context.Background(), []string{executions.PendingStatus, executions.InProgressStatus, executions.DoneStatus})

	assert.Equal(t, want, got)
	mockPostgresConnection.AssertExpectations(t)
	mockPgxRows.AssertExpectations(t)
}

func TestSelectLastExecutionDayExecutedByCriteriaID_success(t *testing.T) {
	mockPostgresConnection := new(database.MockPostgresConnection)
	mockPgxRow := new(database.MockPgxRow)
	mockDate := time.Date(2024, 9, 19, 0, 0, 0, 0, time.Local)
	mockExecutionID := 1
	database.MockScan(mockPgxRow, []any{mockDate, mockExecutionID}, t)
	mockPostgresConnection.On("QueryRow", mock.Anything, mock.Anything, mock.Anything).Return(mockPgxRow)

	selectLastDayExecutedByCriteriaID := executions.MakeSelectLastDayExecutedByCriteriaID(mockPostgresConnection)

	want := executions.ExecutionDayDAO{ExecutionDate: mockDate, SearchCriteriaExecutionID: mockExecutionID}
	got, err := selectLastDayExecutedByCriteriaID(context.Background(), 1)

	assert.Nil(t, err)
	assert.Equal(t, want, got)
	mockPostgresConnection.AssertExpectations(t)
	mockPgxRow.AssertExpectations(t)
}

func TestSelectLastExecutionDayExecutedByCriteriaID_failsWhenSelectOperationFails(t *testing.T) {
	tests := []struct {
		err      error
		expected error
	}{
		{err: pgx.ErrNoRows, expected: executions.NoExecutionDaysFoundForTheGivenCriteriaID},
		{err: errors.New("failed to execute select operation"), expected: executions.FailedToRetrieveLastDayExecutedDate},
	}

	for _, tt := range tests {
		mockPostgresConnection := new(database.MockPostgresConnection)
		mockPgxRow := new(database.MockPgxRow)
		mockPgxRow.On("Scan", mock.Anything).Return(tt.err)
		mockPostgresConnection.On("QueryRow", mock.Anything, mock.Anything, mock.Anything).Return(mockPgxRow)

		selectLastDayExecutedByCriteriaID := executions.MakeSelectLastDayExecutedByCriteriaID(mockPostgresConnection)

		want := tt.expected
		_, got := selectLastDayExecutedByCriteriaID(context.Background(), 1)

		assert.Equal(t, want, got)
		mockPostgresConnection.AssertExpectations(t)
		mockPgxRow.AssertExpectations(t)
	}
}
