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

func TestMakeInsert_success(t *testing.T) {
	mockPostgresConnection := new(database.MockPostgresConnection)
	mockPostgresConnection.On("Exec", mock.Anything, mock.Anything, mock.Anything).Return(pgconn.CommandTag{}, nil)
	mockTweetDTO := tweets.MockTweetDTOSlice()

	insertTweet := tweets.MakeInsert(mockPostgresConnection)

	got := insertTweet(mockTweetDTO)

	assert.Nil(t, got)
}

func TestMakeInsert_failsWhenInsertOperationThrowsError(t *testing.T) {
	mockPostgresConnection := new(database.MockPostgresConnection)
	mockPostgresConnection.On("Exec", mock.Anything, mock.Anything, mock.Anything).Return(pgconn.CommandTag{}, errors.New("failed to insert tweets"))
	mockTweetDTO := tweets.MockTweetDTOSlice()

	insertTweet := tweets.MakeInsert(mockPostgresConnection)

	want := tweets.FailedToInsertTweets
	got := insertTweet(mockTweetDTO)

	assert.Equal(t, want, got)
}
