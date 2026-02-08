package tweets_test

import (
	"context"
	"errors"
	"testing"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/lhbelfanti/corpus-creator/cmd/api/tweets"
	"github.com/lhbelfanti/corpus-creator/cmd/api/tweets/quotes"
	"github.com/lhbelfanti/corpus-creator/internal/database"
)

func TestInsert_success(t *testing.T) {
	mockPostgresConnection := new(database.MockPostgresConnection)
	mockPostgresConnection.On("Exec", mock.Anything, mock.Anything, mock.Anything).Return(pgconn.CommandTag{}, nil)
	mockInsertSingleQuote := quotes.MockInsertSingle(1, nil)
	mockDeleteOrphanQuotes := quotes.MockDeleteOrphans(nil)
	mockTweetDTO := tweets.MockTweetsDTOs()

	insertTweet := tweets.MakeInsert(mockPostgresConnection, mockInsertSingleQuote, mockDeleteOrphanQuotes)

	got := insertTweet(context.Background(), mockTweetDTO)

	assert.Nil(t, got)
	mockPostgresConnection.AssertExpectations(t)
}

func TestInsert_successWithTextContentImagesAndQuoteNil(t *testing.T) {
	mockPostgresConnection := new(database.MockPostgresConnection)
	mockPostgresConnection.On("Exec", mock.Anything, mock.Anything, mock.Anything).Return(pgconn.CommandTag{}, nil)
	mockInsertSingleQuote := quotes.MockInsertSingle(1, nil)
	mockDeleteOrphanQuotes := quotes.MockDeleteOrphans(nil)
	mockTweetDTO := tweets.MockTweetsDTOs()
	mockTweetDTO[0].TextContent = nil
	mockTweetDTO[0].Images = nil
	mockTweetDTO[0].Quote = nil
	mockTweetDTO[1].TextContent = nil
	mockTweetDTO[1].Images = nil
	mockTweetDTO[1].Quote = nil

	insertTweet := tweets.MakeInsert(mockPostgresConnection, mockInsertSingleQuote, mockDeleteOrphanQuotes)

	got := insertTweet(context.Background(), mockTweetDTO)

	assert.Nil(t, got)
	mockPostgresConnection.AssertExpectations(t)
}

func TestInsert_successEvenWhenTheTimestampParseFails(t *testing.T) {
	mockPostgresConnection := new(database.MockPostgresConnection)
	mockPostgresConnection.On("Exec", mock.Anything, mock.Anything, mock.Anything).Return(pgconn.CommandTag{}, nil)
	mockInsertSingleQuote := quotes.MockInsertSingle(1, nil)
	mockDeleteOrphanQuotes := quotes.MockDeleteOrphans(nil)
	mockTweetDTO := tweets.MockTweetsDTOs()
	mockTweetDTO[0].PostedAt = "wrong"

	insertTweet := tweets.MakeInsert(mockPostgresConnection, mockInsertSingleQuote, mockDeleteOrphanQuotes)

	got := insertTweet(context.Background(), mockTweetDTO)

	assert.Nil(t, got)
	mockPostgresConnection.AssertExpectations(t)
}

func TestInsert_successEvenWhenTheQuoteInsertFailsInsertingNilQuoteInTweetsTable(t *testing.T) {
	mockPostgresConnection := new(database.MockPostgresConnection)
	mockPostgresConnection.On("Exec", mock.Anything, mock.Anything, mock.Anything).Return(pgconn.CommandTag{}, nil)
	mockInsertSingleQuote := quotes.MockInsertSingle(-1, errors.New("failed to insert single quote"))
	mockDeleteOrphanQuotes := quotes.MockDeleteOrphans(nil)
	mockTweetDTO := tweets.MockTweetsDTOs()

	insertTweet := tweets.MakeInsert(mockPostgresConnection, mockInsertSingleQuote, mockDeleteOrphanQuotes)

	got := insertTweet(context.Background(), mockTweetDTO)

	assert.Nil(t, got)
	mockPostgresConnection.AssertExpectations(t)
}

func TestInsert_successEvenWhenTheDeleteOrphanQuotesThrowsError(t *testing.T) {
	mockPostgresConnection := new(database.MockPostgresConnection)
	mockPostgresConnection.On("Exec", mock.Anything, mock.Anything, mock.Anything).Return(pgconn.CommandTag{}, nil)
	mockInsertSingleQuote := quotes.MockInsertSingle(1, nil)
	mockDeleteOrphanQuotes := quotes.MockDeleteOrphans(errors.New("failed to delete orphan quotes"))
	mockTweetDTO := tweets.MockTweetsDTOs()

	insertTweet := tweets.MakeInsert(mockPostgresConnection, mockInsertSingleQuote, mockDeleteOrphanQuotes)

	got := insertTweet(context.Background(), mockTweetDTO)

	assert.Nil(t, got)
	mockPostgresConnection.AssertExpectations(t)
}

func TestInsert_failsWhenInsertOperationThrowsError(t *testing.T) {
	mockPostgresConnection := new(database.MockPostgresConnection)
	mockPostgresConnection.On("Exec", mock.Anything, mock.Anything, mock.Anything).Return(pgconn.CommandTag{}, errors.New("failed to insert tweets"))
	mockInsertSingleQuote := quotes.MockInsertSingle(1, nil)
	mockDeleteOrphanQuotes := quotes.MockDeleteOrphans(nil)
	mockTweetDTO := tweets.MockTweetsDTOs()

	insertTweet := tweets.MakeInsert(mockPostgresConnection, mockInsertSingleQuote, mockDeleteOrphanQuotes)

	want := tweets.FailedToInsertTweets
	got := insertTweet(context.Background(), mockTweetDTO)

	assert.Equal(t, want, got)
	mockPostgresConnection.AssertExpectations(t)
}
