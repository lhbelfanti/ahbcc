package tweets_test

import (
	"errors"
	"testing"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"ahbcc/cmd/api/tweets"
	"ahbcc/internal/database"
)

func TestMakeInsertTweet_success(t *testing.T) {
	mockPostgresConnection := new(database.MockPostgresConnection)
	mockPostgresConnection.On("Exec", mock.Anything, mock.Anything, mock.Anything).Return(pgconn.CommandTag{}, nil)
	mockTweetDTO := tweets.MockTweetDTO()

	insertTweet := tweets.MakeInsertTweet(mockPostgresConnection)

	got := insertTweet(mockTweetDTO)

	assert.Nil(t, got)
}

func TestMakeInsertTweet_failsWhenInsertOperationThrowsError(t *testing.T) {
	mockPostgresConnection := new(database.MockPostgresConnection)
	mockPostgresConnection.On("Exec", mock.Anything, mock.Anything, mock.Anything).Return(pgconn.CommandTag{}, errors.New("failed to insert tweet"))
	mockTweetDTO := tweets.MockTweetDTO()

	insertTweet := tweets.MakeInsertTweet(mockPostgresConnection)

	want := tweets.FailedToInsertTweet
	got := insertTweet(mockTweetDTO)

	assert.Equal(t, want, got)
}
