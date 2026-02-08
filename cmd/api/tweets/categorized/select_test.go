package categorized_test

import (
	"context"
	"errors"
	"testing"

	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/lhbelfanti/corpus-creator/cmd/api/tweets/categorized"
	"github.com/lhbelfanti/corpus-creator/internal/database"
)

func TestSelectAllByUserID_success(t *testing.T) {
	mockPostgresConnection := new(database.MockPostgresConnection)
	mockPgxRows := new(database.MockPgxRows)
	mockPostgresConnection.On("Query", mock.Anything, mock.Anything, mock.Anything).Return(mockPgxRows, nil)
	mockCategorizedTweetsDAOSlice := categorized.MockCategorizedTweetsDAOSlice()
	mockCollectRows := database.MockCollectRows[categorized.AnalyzedTweetsDTO](mockCategorizedTweetsDAOSlice, nil)

	selectAllByUserID := categorized.MakeSelectAllByUserID(mockPostgresConnection, mockCollectRows)

	want := mockCategorizedTweetsDAOSlice
	got, err := selectAllByUserID(context.Background(), 1)

	assert.Nil(t, err)
	assert.Equal(t, want, got)
	mockPostgresConnection.AssertExpectations(t)
	mockPgxRows.AssertExpectations(t)
}

func TestSelectAllByUserID_failsWhenSelectOperationThrowsError(t *testing.T) {
	mockPostgresConnection := new(database.MockPostgresConnection)
	mockPgxRows := new(database.MockPgxRows)
	mockPostgresConnection.On("Query", mock.Anything, mock.Anything, mock.Anything).Return(mockPgxRows, errors.New("failed to select all by user id"))
	mockCategorizedTweetsDAOSlice := categorized.MockCategorizedTweetsDAOSlice()
	mockCollectRows := database.MockCollectRows[categorized.AnalyzedTweetsDTO](mockCategorizedTweetsDAOSlice, nil)

	selectAllByUserID := categorized.MakeSelectAllByUserID(mockPostgresConnection, mockCollectRows)

	want := categorized.FailedToExecuteSelectAllCategorizedTweetsByUserID
	_, got := selectAllByUserID(context.Background(), 1)

	assert.Equal(t, want, got)
	mockPostgresConnection.AssertExpectations(t)
	mockPgxRows.AssertExpectations(t)
}

func TestSelectAllByUserID_failsWhenCollectRowsThrowsError(t *testing.T) {
	mockPostgresConnection := new(database.MockPostgresConnection)
	mockPgxRows := new(database.MockPgxRows)
	mockPostgresConnection.On("Query", mock.Anything, mock.Anything, mock.Anything).Return(mockPgxRows, nil)
	mockCollectRows := database.MockCollectRows[categorized.AnalyzedTweetsDTO](nil, errors.New("failed to collect rows"))

	selectAllByUserID := categorized.MakeSelectAllByUserID(mockPostgresConnection, mockCollectRows)

	want := categorized.FailedToExecuteCollectRowsInSelectAllCategorizedTweetsByUserID
	_, got := selectAllByUserID(context.Background(), 1)

	assert.Equal(t, want, got)
	mockPostgresConnection.AssertExpectations(t)
	mockPgxRows.AssertExpectations(t)
}

func TestSelectByUserIDTweetIDAndSearchCriteriaID_success(t *testing.T) {
	mockPostgresConnection := new(database.MockPostgresConnection)

	mockPgxRow := new(database.MockPgxRow)
	mockCategorizedTweetDAO := categorized.MockCategorizedTweetDAO()
	mockScanTweetDAOValues := categorized.MockScanCategorizedTweetsDAOValues(mockCategorizedTweetDAO)
	database.MockScan(mockPgxRow, mockScanTweetDAOValues, t)
	mockPostgresConnection.On("QueryRow", mock.Anything, mock.Anything, mock.Anything).Return(mockPgxRow)

	selectByUserIDTweetIDAndSearchCriteriaID := categorized.MakeSelectByUserIDTweetIDAndSearchCriteriaID(mockPostgresConnection)

	want := mockCategorizedTweetDAO
	got, err := selectByUserIDTweetIDAndSearchCriteriaID(context.Background(), 456, 123, 2)

	assert.Nil(t, err)
	assert.Equal(t, want, got)
	mockPostgresConnection.AssertExpectations(t)
	mockPgxRow.AssertExpectations(t)
}

func TestSelectByUserIDTweetIDAndSearchCriteriaID_failsWhenSelectOperationFails(t *testing.T) {
	tests := []struct {
		err      error
		expected error
	}{
		{err: pgx.ErrNoRows, expected: categorized.NoCategorizedTweetFound},
		{err: errors.New("failed to execute select operation"), expected: categorized.FailedExecuteQueryToRetrieveCategorizedTweetData},
	}

	for _, tt := range tests {
		mockPostgresConnection := new(database.MockPostgresConnection)
		mockPgxRow := new(database.MockPgxRow)
		mockPgxRow.On("Scan", mock.Anything).Return(tt.err)
		mockPostgresConnection.On("QueryRow", mock.Anything, mock.Anything, mock.Anything).Return(mockPgxRow)

		selectByUserIDTweetIDAndSearchCriteriaID := categorized.MakeSelectByUserIDTweetIDAndSearchCriteriaID(mockPostgresConnection)

		want := tt.expected
		_, got := selectByUserIDTweetIDAndSearchCriteriaID(context.Background(), 456, 123, 2)

		assert.Equal(t, want, got)
	}
}

func TestSelectByCategorizations_success(t *testing.T) {
	mockPostgresConnection := new(database.MockPostgresConnection)
	mockPgxRows := new(database.MockPgxRows)
	mockPostgresConnection.On("Query", mock.Anything, mock.Anything, mock.Anything).Return(mockPgxRows, nil)
	mockCategorizedTweetsDAOSlice := []categorized.DAO{categorized.MockCategorizedTweetDAO()}
	mockCollectRows := database.MockCollectRows[categorized.DAO](mockCategorizedTweetsDAOSlice, nil)

	selectByCategorizations := categorized.MakeSelectByCategorizations(mockPostgresConnection, mockCollectRows)

	want := mockCategorizedTweetsDAOSlice
	got, err := selectByCategorizations(context.Background(), []string{categorized.VerdictPositive, categorized.VerdictNegative})

	assert.Nil(t, err)
	assert.Equal(t, want, got)
	mockPostgresConnection.AssertExpectations(t)
	mockPgxRows.AssertExpectations(t)
}

func TestSelectByCategorizations_failsWhenSelectOperationThrowsError(t *testing.T) {
	mockPostgresConnection := new(database.MockPostgresConnection)
	mockPgxRows := new(database.MockPgxRows)
	mockPostgresConnection.On("Query", mock.Anything, mock.Anything, mock.Anything).Return(mockPgxRows, errors.New("failed to select by categorizations"))
	mockCategorizedTweetsDAOSlice := []categorized.DAO{categorized.MockCategorizedTweetDAO()}
	mockCollectRows := database.MockCollectRows[categorized.DAO](mockCategorizedTweetsDAOSlice, nil)

	selectByCategorizations := categorized.MakeSelectByCategorizations(mockPostgresConnection, mockCollectRows)

	want := categorized.FailedToExecuteSelectByCategorizations
	_, got := selectByCategorizations(context.Background(), []string{categorized.VerdictPositive, categorized.VerdictNegative})

	assert.Equal(t, want, got)
	mockPostgresConnection.AssertExpectations(t)
	mockPgxRows.AssertExpectations(t)
}

func TestSelectByCategorizations_failsWhenCollectRowsThrowsError(t *testing.T) {
	mockPostgresConnection := new(database.MockPostgresConnection)
	mockPgxRows := new(database.MockPgxRows)
	mockPostgresConnection.On("Query", mock.Anything, mock.Anything, mock.Anything).Return(mockPgxRows, nil)
	mockCollectRows := database.MockCollectRows[categorized.DAO](nil, errors.New("failed to collect rows"))

	selectByCategorizations := categorized.MakeSelectByCategorizations(mockPostgresConnection, mockCollectRows)

	want := categorized.FailedToExecuteCollectRowsInSelectByCategorizations
	_, got := selectByCategorizations(context.Background(), []string{categorized.VerdictPositive, categorized.VerdictNegative})

	assert.Equal(t, want, got)
	mockPostgresConnection.AssertExpectations(t)
	mockPgxRows.AssertExpectations(t)
}
