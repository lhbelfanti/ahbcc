package criteria_test

import (
	"context"
	"testing"

	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"ahbcc/cmd/api/search/criteria"
	"ahbcc/internal/database"
)

func TestInsertSingle_success(t *testing.T) {
	tests := []struct {
		forced bool
	}{
		{forced: false},
		{forced: true},
	}

	for _, tt := range tests {
		searchCriteriaExecutionID := 1
		mockPostgresConnection := new(database.MockPostgresConnection)
		mockPgxRow := new(database.MockPgxRow)
		database.MockScan[int](mockPgxRow, searchCriteriaExecutionID, t)
		mockPostgresConnection.On("QueryRow", mock.Anything, mock.Anything, mock.Anything).Return(mockPgxRow)

		insertSearchCriteriaExecution := criteria.MakeInsertExecution(mockPostgresConnection)

		want := searchCriteriaExecutionID
		got, err := insertSearchCriteriaExecution(context.Background(), 5, tt.forced)

		assert.Equal(t, want, got)
		assert.Nil(t, err)
		mockPostgresConnection.AssertExpectations(t)
	}
}

func TestInsertSingle_failsWhenInsertOperationThrowsError(t *testing.T) {
	mockPostgresConnection := new(database.MockPostgresConnection)
	mockPgxRow := new(database.MockPgxRow)
	mockPgxRow.On("Scan", mock.Anything).Return(pgx.ErrNoRows)
	mockPostgresConnection.On("QueryRow", mock.Anything, mock.Anything, mock.Anything).Return(mockPgxRow)

	insertSearchCriteriaExecution := criteria.MakeInsertExecution(mockPostgresConnection)

	want := criteria.FailedToInsertSearchExecutionCriteria
	_, got := insertSearchCriteriaExecution(context.Background(), 5, false)

	assert.Equal(t, want, got)
	mockPostgresConnection.AssertExpectations(t)
	mockPgxRow.AssertExpectations(t)
}
