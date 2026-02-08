package tweets_test

import (
	"context"
	"errors"
	"testing"

	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/lhbelfanti/corpus-creator/cmd/api/tweets"
	"github.com/lhbelfanti/corpus-creator/cmd/api/user/session"
	"github.com/lhbelfanti/corpus-creator/internal/database"
)

func TestSelectBySearchCriteriaIDYearAndMonth_success(t *testing.T) {
	mockPostgresConnection := new(database.MockPostgresConnection)
	mockPgxRows := new(database.MockPgxRows)
	mockPostgresConnection.On("Query", mock.Anything, mock.Anything, mock.Anything).Return(mockPgxRows, nil)
	mockTweetsDTOs := tweets.MockCustomTweetDTOs()
	mockCollectRows := database.MockCollectRows[tweets.CustomTweetDTO](mockTweetsDTOs, nil)
	mockSelectUserIDByToken := session.MockSelectUserIDByToken(1, nil)

	selectBySearchCriteriaIDYearAndMonth := tweets.MakeSelectBySearchCriteriaIDYearAndMonth(mockPostgresConnection, mockCollectRows, mockSelectUserIDByToken)

	want := mockTweetsDTOs
	got, err := selectBySearchCriteriaIDYearAndMonth(context.Background(), 1, 2025, 04, 10, "token")

	assert.Nil(t, err)
	assert.Equal(t, want, got)
	mockPostgresConnection.AssertExpectations(t)
	mockPgxRows.AssertExpectations(t)
}

func TestSelectBySearchCriteriaIDYearAndMonth_successWithMonthZero(t *testing.T) {
	mockPostgresConnection := new(database.MockPostgresConnection)
	mockPgxRows := new(database.MockPgxRows)
	mockPostgresConnection.On("Query", mock.Anything, mock.Anything, mock.Anything).Return(mockPgxRows, nil)
	mockTweetsDTOs := tweets.MockCustomTweetDTOs()
	mockCollectRows := database.MockCollectRows[tweets.CustomTweetDTO](mockTweetsDTOs, nil)
	mockSelectUserIDByToken := session.MockSelectUserIDByToken(1, nil)

	selectBySearchCriteriaIDYearAndMonth := tweets.MakeSelectBySearchCriteriaIDYearAndMonth(mockPostgresConnection, mockCollectRows, mockSelectUserIDByToken)

	want := mockTweetsDTOs
	got, err := selectBySearchCriteriaIDYearAndMonth(context.Background(), 1, 2025, 0, 10, "token")

	assert.Nil(t, err)
	assert.Equal(t, want, got)
	mockPostgresConnection.AssertExpectations(t)
	mockPgxRows.AssertExpectations(t)
}

func TestSelectBySearchCriteriaIDYearAndMonth_successWithYearZeroAndMonthZero(t *testing.T) {
	mockPostgresConnection := new(database.MockPostgresConnection)
	mockPgxRows := new(database.MockPgxRows)
	mockPostgresConnection.On("Query", mock.Anything, mock.Anything, mock.Anything).Return(mockPgxRows, nil)
	mockTweetsDTOs := tweets.MockCustomTweetDTOs()
	mockCollectRows := database.MockCollectRows[tweets.CustomTweetDTO](mockTweetsDTOs, nil)
	mockSelectUserIDByToken := session.MockSelectUserIDByToken(1, nil)

	selectBySearchCriteriaIDYearAndMonth := tweets.MakeSelectBySearchCriteriaIDYearAndMonth(mockPostgresConnection, mockCollectRows, mockSelectUserIDByToken)

	want := mockTweetsDTOs
	got, err := selectBySearchCriteriaIDYearAndMonth(context.Background(), 1, 0, 0, 10, "token")

	assert.Nil(t, err)
	assert.Equal(t, want, got)
	mockPostgresConnection.AssertExpectations(t)
	mockPgxRows.AssertExpectations(t)
}

func TestSelectBySearchCriteriaIDYearAndMonth_failsWhenSelectUserIDByTokenThrowsError(t *testing.T) {
	mockPostgresConnection := new(database.MockPostgresConnection)
	mockPgxRows := new(database.MockPgxRows)
	mockPostgresConnection.On("Query", mock.Anything, mock.Anything, mock.Anything).Return(mockPgxRows, nil)
	mockTweetsDTOs := tweets.MockCustomTweetDTOs()
	mockCollectRows := database.MockCollectRows[tweets.CustomTweetDTO](mockTweetsDTOs, nil)
	mockSelectUserIDByToken := session.MockSelectUserIDByToken(1, errors.New("failed to select user id by token"))

	selectBySearchCriteriaIDYearAndMonth := tweets.MakeSelectBySearchCriteriaIDYearAndMonth(mockPostgresConnection, mockCollectRows, mockSelectUserIDByToken)

	want := tweets.FailedToRetrieveUserID
	_, got := selectBySearchCriteriaIDYearAndMonth(context.Background(), 1, 2025, 04, 10, "token")

	assert.Equal(t, want, got)
	mockPgxRows.AssertExpectations(t)
}

func TestSelectBySearchCriteriaIDYearAndMonth_failsWhenSelectOperationThrowsError(t *testing.T) {
	mockPostgresConnection := new(database.MockPostgresConnection)
	mockPgxRows := new(database.MockPgxRows)
	mockPostgresConnection.On("Query", mock.Anything, mock.Anything, mock.Anything).Return(mockPgxRows, errors.New("failed to select by search criteria id year and month"))
	mockTweetsDTOs := tweets.MockCustomTweetDTOs()
	mockCollectRows := database.MockCollectRows[tweets.CustomTweetDTO](mockTweetsDTOs, nil)
	mockSelectUserIDByToken := session.MockSelectUserIDByToken(1, nil)

	selectBySearchCriteriaIDYearAndMonth := tweets.MakeSelectBySearchCriteriaIDYearAndMonth(mockPostgresConnection, mockCollectRows, mockSelectUserIDByToken)

	want := tweets.FailedToRetrieveUserUncategorizedTweets
	_, got := selectBySearchCriteriaIDYearAndMonth(context.Background(), 1, 2025, 04, 10, "token")

	assert.Equal(t, want, got)
	mockPostgresConnection.AssertExpectations(t)
	mockPgxRows.AssertExpectations(t)
}

func TestSelectBySearchCriteriaIDYearAndMonth_failsWhenCollectRowsThrowsError(t *testing.T) {
	mockPostgresConnection := new(database.MockPostgresConnection)
	mockPgxRows := new(database.MockPgxRows)
	mockPostgresConnection.On("Query", mock.Anything, mock.Anything, mock.Anything).Return(mockPgxRows, nil)
	mockTweetsDTOs := tweets.MockCustomTweetDTOs()
	mockCollectRows := database.MockCollectRows[tweets.CustomTweetDTO](mockTweetsDTOs, errors.New("failed to collect rows"))
	mockSelectUserIDByToken := session.MockSelectUserIDByToken(1, nil)

	selectBySearchCriteriaIDYearAndMonth := tweets.MakeSelectBySearchCriteriaIDYearAndMonth(mockPostgresConnection, mockCollectRows, mockSelectUserIDByToken)

	want := tweets.FailedToExecuteCollectRowsInSelectUserUncategorizedTweets
	_, got := selectBySearchCriteriaIDYearAndMonth(context.Background(), 1, 2025, 04, 10, "token")

	assert.Equal(t, want, got)
	mockPostgresConnection.AssertExpectations(t)
	mockPgxRows.AssertExpectations(t)
}

func TestSelectByID_success(t *testing.T) {
	mockPostgresConnection := new(database.MockPostgresConnection)

	mockPgxRow := new(database.MockPgxRow)
	mockTweetDAO := tweets.MockTweetDAO()
	mockScanTweetDAOValues := tweets.MockScanTweetDAOValues(mockTweetDAO)
	database.MockScan(mockPgxRow, mockScanTweetDAOValues, t)
	mockPostgresConnection.On("QueryRow", mock.Anything, mock.Anything, mock.Anything).Return(mockPgxRow)

	selectByID := tweets.MakeSelectByID(mockPostgresConnection)

	want := mockTweetDAO
	got, err := selectByID(context.Background(), 1)

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
		{err: pgx.ErrNoRows, expected: tweets.NoTweetFoundForTheGivenTweetID},
		{err: errors.New("failed to execute select operation"), expected: tweets.FailedExecuteQueryToRetrieveTweetData},
	}

	for _, tt := range tests {
		mockPostgresConnection := new(database.MockPostgresConnection)
		mockPgxRow := new(database.MockPgxRow)
		mockPgxRow.On("Scan", mock.Anything).Return(tt.err)
		mockPostgresConnection.On("QueryRow", mock.Anything, mock.Anything, mock.Anything).Return(mockPgxRow)

		selectByID := tweets.MakeSelectByID(mockPostgresConnection)

		want := tt.expected
		_, got := selectByID(context.Background(), 1)

		assert.Equal(t, want, got)
	}
}
