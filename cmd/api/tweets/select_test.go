package tweets_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"ahbcc/cmd/api/tweets"
	"ahbcc/cmd/api/user/session"
	"ahbcc/internal/database"
)

func TestSelectBySearchCriteriaIDYearAndMonth_success(t *testing.T) {
	mockPostgresConnection := new(database.MockPostgresConnection)
	mockPgxRows := new(database.MockPgxRows)
	mockPostgresConnection.On("Query", mock.Anything, mock.Anything, mock.Anything).Return(mockPgxRows, nil)
	mockTweetsDTOs := tweets.MockTweetsDTOs()
	mockCollectRows := database.MockCollectRows[tweets.TweetDTO](mockTweetsDTOs, nil)
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
	mockTweetsDTOs := tweets.MockTweetsDTOs()
	mockCollectRows := database.MockCollectRows[tweets.TweetDTO](mockTweetsDTOs, nil)
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
	mockTweetsDTOs := tweets.MockTweetsDTOs()
	mockCollectRows := database.MockCollectRows[tweets.TweetDTO](mockTweetsDTOs, nil)
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
	mockTweetsDTOs := tweets.MockTweetsDTOs()
	mockCollectRows := database.MockCollectRows[tweets.TweetDTO](mockTweetsDTOs, nil)
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
	mockTweetsDTOs := tweets.MockTweetsDTOs()
	mockCollectRows := database.MockCollectRows[tweets.TweetDTO](mockTweetsDTOs, nil)
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
	mockTweetsDTOs := tweets.MockTweetsDTOs()
	mockCollectRows := database.MockCollectRows[tweets.TweetDTO](mockTweetsDTOs, errors.New("failed to collect rows"))
	mockSelectUserIDByToken := session.MockSelectUserIDByToken(1, nil)

	selectBySearchCriteriaIDYearAndMonth := tweets.MakeSelectBySearchCriteriaIDYearAndMonth(mockPostgresConnection, mockCollectRows, mockSelectUserIDByToken)

	want := tweets.FailedToExecuteCollectRowsInSelectUserUncategorizedTweets
	_, got := selectBySearchCriteriaIDYearAndMonth(context.Background(), 1, 2025, 04, 10, "token")

	assert.Equal(t, want, got)
	mockPostgresConnection.AssertExpectations(t)
	mockPgxRows.AssertExpectations(t)
}
