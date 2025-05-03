package categorized_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"ahbcc/cmd/api/tweets/categorized"
	"ahbcc/cmd/api/user/session"
)

func TestInsertCategorizedTweet_success(t *testing.T) {
	mockSelectUserIDByToken := session.MockSelectUserIDByToken(789, nil)
	mockInsertSingle := categorized.MockInsertSingle(1, nil)
	mockDTO := categorized.MockDTO()
	mockDTO.UserID = 0 // This should be overwritten by the service

	insertCategorizedTweet := categorized.MakeInsertCategorizedTweet(mockSelectUserIDByToken, mockInsertSingle)

	id, err := insertCategorizedTweet(context.Background(), "valid-token", mockDTO)

	assert.NoError(t, err)
	assert.Equal(t, 1, id)
}

func TestInsertCategorizedTweet_failsWhenSelectUserIDByTokenFails(t *testing.T) {
	mockSelectUserIDByToken := session.MockSelectUserIDByToken(789, errors.New("failed to select user id by token"))
	mockInsertSingle := categorized.MockInsertSingle(1, nil)
	mockDTO := categorized.MockDTO()

	insertCategorizedTweet := categorized.MakeInsertCategorizedTweet(mockSelectUserIDByToken, mockInsertSingle)

	want := categorized.FailedToRetrieveUserID
	_, got := insertCategorizedTweet(context.Background(), "invalid-token", mockDTO)

	assert.Equal(t, want, got)
}

func TestInsertCategorizedTweet_failsWhenInsertSingleFails(t *testing.T) {
	mockSelectUserIDByToken := func(ctx context.Context, token string) (int, error) {
		return 789, nil
	}
	mockInsertSingle := categorized.MockInsertSingle(-1, errors.New("failed to insert categorized tweet"))
	mockDTO := categorized.MockDTO()

	insertCategorizedTweet := categorized.MakeInsertCategorizedTweet(mockSelectUserIDByToken, mockInsertSingle)

	want := categorized.FailedToInsertSingleCategorizedTweet
	_, got := insertCategorizedTweet(context.Background(), "valid-token", mockDTO)

	assert.Equal(t, want, got)
}
