package tweets_test

import (
	"errors"
	"testing"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"ahbcc/cmd/api/tweets"
	"ahbcc/cmd/api/tweets/quotes"
	"ahbcc/internal/database"
)

func TestMakeInsert_success(t *testing.T) {
	mockPostgresConnection := new(database.MockPostgresConnection)
	mockPostgresConnection.On("Exec", mock.Anything, mock.Anything, mock.Anything).Return(pgconn.CommandTag{}, nil)
	mockInsertSingleQuote := quotes.MockInsertSingle(1, nil)
	mockTweetDTO := tweets.MockTweets()

	insertTweet := tweets.MakeInsert(mockPostgresConnection, mockInsertSingleQuote)

	got := insertTweet(mockTweetDTO)

	assert.Nil(t, got)
	mockPostgresConnection.AssertExpectations(t)
}

func TestMakeInsert_successWithTextContentImagesAndQuoteNil(t *testing.T) {
	mockPostgresConnection := new(database.MockPostgresConnection)
	mockPostgresConnection.On("Exec", mock.Anything, mock.Anything, mock.Anything).Return(pgconn.CommandTag{}, nil)
	mockInsertSingleQuote := quotes.MockInsertSingle(1, nil)
	mockTweetDTO := tweets.MockTweets()
	mockTweetDTO[0].TextContent = nil
	mockTweetDTO[0].Images = nil
	mockTweetDTO[0].Quote = nil
	mockTweetDTO[1].TextContent = nil
	mockTweetDTO[1].Images = nil
	mockTweetDTO[1].Quote = nil

	insertTweet := tweets.MakeInsert(mockPostgresConnection, mockInsertSingleQuote)

	got := insertTweet(mockTweetDTO)

	assert.Nil(t, got)
	mockPostgresConnection.AssertExpectations(t)
}

func TestMakeInsert_successEvenWhenTheQuoteInsertFailsInsertingNilQuoteInTweetsTable(t *testing.T) {
	mockPostgresConnection := new(database.MockPostgresConnection)
	mockPostgresConnection.On("Exec", mock.Anything, mock.Anything, mock.Anything).Return(pgconn.CommandTag{}, nil)
	mockInsertSingleQuote := quotes.MockInsertSingle(-1, errors.New("failed to insert single quote"))
	mockTweetDTO := tweets.MockTweets()

	insertTweet := tweets.MakeInsert(mockPostgresConnection, mockInsertSingleQuote)

	got := insertTweet(mockTweetDTO)

	assert.Nil(t, got)
	mockPostgresConnection.AssertExpectations(t)
}

func TestMakeInsert_failsWhenInsertOperationThrowsError(t *testing.T) {
	mockPostgresConnection := new(database.MockPostgresConnection)
	mockPostgresConnection.On("Exec", mock.Anything, mock.Anything, mock.Anything).Return(pgconn.CommandTag{}, errors.New("failed to insert tweets"))
	mockInsertSingleQuote := quotes.MockInsertSingle(1, nil)
	mockTweetDTO := tweets.MockTweets()

	insertTweet := tweets.MakeInsert(mockPostgresConnection, mockInsertSingleQuote)

	want := tweets.FailedToInsertTweets
	got := insertTweet(mockTweetDTO)

	assert.Equal(t, want, got)
	mockPostgresConnection.AssertExpectations(t)
}
