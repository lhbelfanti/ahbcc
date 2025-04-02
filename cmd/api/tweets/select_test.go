package tweets_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"ahbcc/cmd/api/tweets"
	"ahbcc/cmd/api/tweets/counts"
	"ahbcc/internal/database"
)

func TestSelectMonthlyTweetsCountsByYearByCriteriaID_success(t *testing.T) {
	mockPostgresConnection := new(database.MockPostgresConnection)
	mockPgxRows := new(database.MockPgxRows)
	mockPostgresConnection.On("Query", mock.Anything, mock.Anything, mock.Anything).Return(mockPgxRows, nil)
	mockTweetsCountsDAOSlice := counts.MockTweetsCountsDAOSlice()
	mockCollectRows := database.MockCollectRows[counts.DAO](mockTweetsCountsDAOSlice, nil)

	selectMonthlyTweetsCountsByYearByCriteriaID := tweets.MakeSelectMonthlyTweetsCountsByYearByCriteriaID(mockPostgresConnection, mockCollectRows)

	want := mockTweetsCountsDAOSlice
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
	mockTweetsCountsDAOSlice := counts.MockTweetsCountsDAOSlice()
	mockCollectRows := database.MockCollectRows[counts.DAO](mockTweetsCountsDAOSlice, nil)

	selectMonthlyTweetsCountsByYearByCriteriaID := tweets.MakeSelectMonthlyTweetsCountsByYearByCriteriaID(mockPostgresConnection, mockCollectRows)

	want := tweets.FailedToRetrieveMonthlyTweetsCountsByYear
	_, got := selectMonthlyTweetsCountsByYearByCriteriaID(context.Background(), 1)

	assert.Equal(t, want, got)
	mockPostgresConnection.AssertExpectations(t)
	mockPgxRows.AssertExpectations(t)
}

func TestSelectMonthlyTweetsCountsByYearByCriteriaID_failsWhenCollectRowsThrowsError(t *testing.T) {
	mockPostgresConnection := new(database.MockPostgresConnection)
	mockPgxRows := new(database.MockPgxRows)
	mockPostgresConnection.On("Query", mock.Anything, mock.Anything, mock.Anything).Return(mockPgxRows, nil)
	mockTweetsCountsDAOSlice := counts.MockTweetsCountsDAOSlice()
	mockCollectRows := database.MockCollectRows[counts.DAO](mockTweetsCountsDAOSlice, errors.New("failed to collect rows"))

	selectMonthlyTweetsCountsByYearByCriteriaID := tweets.MakeSelectMonthlyTweetsCountsByYearByCriteriaID(mockPostgresConnection, mockCollectRows)

	want := tweets.FailedToExecuteCollectRowsInSelectMonthlyTweetsCountsByYear
	_, got := selectMonthlyTweetsCountsByYearByCriteriaID(context.Background(), 1)

	assert.Equal(t, want, got)
	mockPostgresConnection.AssertExpectations(t)
	mockPgxRows.AssertExpectations(t)
}
