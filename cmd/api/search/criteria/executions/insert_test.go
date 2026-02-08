package executions_test

import (
	"context"
	"errors"
	"testing"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/lhbelfanti/corpus-creator/cmd/api/search/criteria/executions"
	"github.com/lhbelfanti/corpus-creator/internal/database"
)

func TestInsertExecution_success(t *testing.T) {
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
		database.MockScan(mockPgxRow, []any{searchCriteriaExecutionID}, t)
		mockPostgresConnection.On("QueryRow", mock.Anything, mock.Anything, mock.Anything).Return(mockPgxRow)

		insertExecution := executions.MakeInsertExecution(mockPostgresConnection)

		want := searchCriteriaExecutionID
		got, err := insertExecution(context.Background(), 5, tt.forced)

		assert.Equal(t, want, got)
		assert.Nil(t, err)
		mockPostgresConnection.AssertExpectations(t)
	}
}

func TestInsertExecution_failsWhenInsertOperationThrowsError(t *testing.T) {
	mockPostgresConnection := new(database.MockPostgresConnection)
	mockPgxRow := new(database.MockPgxRow)
	mockPgxRow.On("Scan", mock.Anything).Return(pgx.ErrNoRows)
	mockPostgresConnection.On("QueryRow", mock.Anything, mock.Anything, mock.Anything).Return(mockPgxRow)

	insertExecution := executions.MakeInsertExecution(mockPostgresConnection)

	want := executions.FailedToInsertSearchCriteriaExecution
	_, got := insertExecution(context.Background(), 5, false)

	assert.Equal(t, want, got)
	mockPostgresConnection.AssertExpectations(t)
	mockPgxRow.AssertExpectations(t)
}

func TestInsertExecutionDay_success(t *testing.T) {
	errorReason := "error reason"
	tests := []struct {
		errorReason *string
	}{
		{errorReason: nil},
		{errorReason: &errorReason},
	}

	for _, tt := range tests {
		mockPostgresConnection := new(database.MockPostgresConnection)
		mockPostgresConnection.On("Exec", mock.Anything, mock.Anything, mock.Anything).Return(pgconn.CommandTag{}, nil)
		mockExecutionDayDTO := executions.MockExecutionDayDTO(tt.errorReason)

		insertExecutionDay := executions.MakeInsertExecutionDay(mockPostgresConnection)

		got := insertExecutionDay(context.Background(), mockExecutionDayDTO)

		assert.Nil(t, got)
		mockPostgresConnection.AssertExpectations(t)
	}
}

func TestInsertExecutionDay_failsWhenInsertOperationThrowsError(t *testing.T) {
	mockPostgresConnection := new(database.MockPostgresConnection)
	mockPostgresConnection.On("Exec", mock.Anything, mock.Anything, mock.Anything).Return(pgconn.CommandTag{}, errors.New("failed to insert execution day"))
	mockExecutionDayDTO := executions.MockExecutionDayDTO(nil)

	insertExecutionDay := executions.MakeInsertExecutionDay(mockPostgresConnection)

	want := executions.FailedToInsertSearchCriteriaExecutionDay
	got := insertExecutionDay(context.Background(), mockExecutionDayDTO)

	assert.Equal(t, want, got)
	mockPostgresConnection.AssertExpectations(t)
}
