package criteria_test

import (
	"context"
	"errors"
	"testing"

	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/lhbelfanti/corpus-creator/cmd/api/search/criteria"
	"github.com/lhbelfanti/corpus-creator/internal/database"
)

func TestSelectByID_success(t *testing.T) {
	mockPostgresConnection := new(database.MockPostgresConnection)
	mockPgxRow := new(database.MockPgxRow)
	mockCriteria := criteria.MockCriteriaDAO()
	mockScanCriteriaDAOValues := criteria.MockScanCriteriaDAOValues(mockCriteria)
	database.MockScan(mockPgxRow, mockScanCriteriaDAOValues, t)
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
	mockPostgresConnection.On("Query", mock.Anything, mock.Anything, mock.Anything).Return(mockPgxRows, nil)
	mockExecutionDAOSlice := criteria.MockCriteriaDAOSlice()
	mockCollectRows := database.MockCollectRows[criteria.DAO](mockExecutionDAOSlice, errors.New("failed to collect rows"))

	selectAllCriteria := criteria.MakeSelectAll(mockPostgresConnection, mockCollectRows)

	want := criteria.FailedToExecuteCollectRowsInSelectAll
	_, got := selectAllCriteria(context.Background())

	assert.Equal(t, want, got)
	mockPostgresConnection.AssertExpectations(t)
	mockPgxRows.AssertExpectations(t)
}
