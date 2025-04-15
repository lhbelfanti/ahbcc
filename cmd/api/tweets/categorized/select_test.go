package categorized_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"ahbcc/cmd/api/tweets/categorized"
	"ahbcc/internal/database"
)

func TestSelectAllByUserID_success(t *testing.T) {
	mockPostgresConnection := new(database.MockPostgresConnection)
	mockPgxRows := new(database.MockPgxRows)
	mockPostgresConnection.On("Query", mock.Anything, mock.Anything, mock.Anything).Return(mockPgxRows, nil)
	mockAnalyzedTweetsDAOSlice := categorized.MockAnalyzedTweetsDAOSlice()
	mockCollectRows := database.MockCollectRows[categorized.AnalyzedTweetsDAO](mockAnalyzedTweetsDAOSlice, nil)

	selectAllByUserID := categorized.MakeSelectAllByUserID(mockPostgresConnection, mockCollectRows)

	want := mockAnalyzedTweetsDAOSlice
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
	mockAnalyzedTweetsDAOSlice := categorized.MockAnalyzedTweetsDAOSlice()
	mockCollectRows := database.MockCollectRows[categorized.AnalyzedTweetsDAO](mockAnalyzedTweetsDAOSlice, nil)

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
	mockCollectRows := database.MockCollectRows[categorized.AnalyzedTweetsDAO](nil, errors.New("failed to collect rows"))

	selectAllByUserID := categorized.MakeSelectAllByUserID(mockPostgresConnection, mockCollectRows)

	want := categorized.FailedToExecuteCollectRowsInSelectAllCategorizedTweetsByUserID
	_, got := selectAllByUserID(context.Background(), 1)

	assert.Equal(t, want, got)
	mockPostgresConnection.AssertExpectations(t)
	mockPgxRows.AssertExpectations(t)
}
