package categorized_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"ahbcc/cmd/api/tweets"
	"ahbcc/cmd/api/tweets/categorized"
	"ahbcc/cmd/api/user/session"
)

func TestInsertCategorizedTweet_success(t *testing.T) {
	mockSelectUserIDByToken := session.MockSelectUserIDByToken(789, nil)
	mockTweetDAO := tweets.MockTweetDAO()
	mockSelectTweetByID := tweets.MockSelectByID(mockTweetDAO, nil)
	mockInsertSingle := categorized.MockInsertSingle(1, nil)
	mockBody := categorized.MockInsertSingleBodyDTO(categorized.VerdictPositive)

	insertCategorizedTweet := categorized.MakeInsertCategorizedTweet(mockSelectUserIDByToken, mockSelectTweetByID, mockInsertSingle)

	want := 1
	got, err := insertCategorizedTweet(context.Background(), "token", mockTweetDAO.ID, mockBody)

	assert.NoError(t, err)
	assert.Equal(t, want, got)
}

func TestInsertCategorizedTweet_failsWhenSelectUserIDByTokenThrowsError(t *testing.T) {
	mockSelectUserIDByToken := session.MockSelectUserIDByToken(789, errors.New("failed to select user id by token"))
	mockTweetDAO := tweets.MockTweetDAO()
	mockSelectTweetByID := tweets.MockSelectByID(mockTweetDAO, nil)
	mockInsertSingle := categorized.MockInsertSingle(1, nil)
	mockBody := categorized.MockInsertSingleBodyDTO(categorized.VerdictPositive)

	insertCategorizedTweet := categorized.MakeInsertCategorizedTweet(mockSelectUserIDByToken, mockSelectTweetByID, mockInsertSingle)

	want := categorized.FailedToRetrieveUserID
	_, got := insertCategorizedTweet(context.Background(), "token", mockTweetDAO.ID, mockBody)

	assert.Equal(t, want, got)
}

func TestInsertCategorizedTweet_failsWhenSelectTweetByIDThrowsError(t *testing.T) {
	mockSelectUserIDByToken := session.MockSelectUserIDByToken(789, nil)
	mockTweetDAO := tweets.MockTweetDAO()
	mockSelectTweetByID := tweets.MockSelectByID(mockTweetDAO, errors.New("failed to select tweet by id"))
	mockInsertSingle := categorized.MockInsertSingle(1, nil)
	mockBody := categorized.MockInsertSingleBodyDTO(categorized.VerdictPositive)

	insertCategorizedTweet := categorized.MakeInsertCategorizedTweet(mockSelectUserIDByToken, mockSelectTweetByID, mockInsertSingle)

	want := categorized.FailedToRetrieveTweetByID
	_, got := insertCategorizedTweet(context.Background(), "token", mockTweetDAO.ID, mockBody)

	assert.Equal(t, want, got)
}

func TestInsertCategorizedTweet_failsWhenInsertSingleThrowsError(t *testing.T) {
	mockSelectUserIDByToken := session.MockSelectUserIDByToken(789, nil)
	mockTweetDAO := tweets.MockTweetDAO()
	mockSelectTweetByID := tweets.MockSelectByID(mockTweetDAO, nil)
	mockInsertSingle := categorized.MockInsertSingle(-1, errors.New("failed to insert categorized tweet"))
	mockBody := categorized.MockInsertSingleBodyDTO(categorized.VerdictPositive)

	insertCategorizedTweet := categorized.MakeInsertCategorizedTweet(mockSelectUserIDByToken, mockSelectTweetByID, mockInsertSingle)

	want := categorized.FailedToInsertSingleCategorizedTweet
	_, got := insertCategorizedTweet(context.Background(), "token", mockTweetDAO.ID, mockBody)

	assert.Equal(t, want, got)
}
