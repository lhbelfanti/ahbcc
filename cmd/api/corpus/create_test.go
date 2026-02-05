package corpus_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"ahbcc/cmd/api/corpus"
	"ahbcc/cmd/api/corpus/cleaner"
	"ahbcc/cmd/api/tweets"
	"ahbcc/cmd/api/tweets/categorized"
	"ahbcc/cmd/api/tweets/quotes"
)

func TestCreate_successWithoutPerfectlyBalancedCorpus(t *testing.T) {
	mockSelectByCategorizations := categorized.MockSelectByCategorizations([]categorized.DAO{categorized.MockCategorizedTweetDAO()}, nil)
	mockSelectTweetByID := tweets.MockSelectByID(tweets.MockTweetDAO(), nil)
	mockSelectQuoteByID := quotes.MockSelectByID(quotes.MockTweetQuoteDAO(), nil)
	mockDeleteAll := corpus.MockDeleteAll(nil)
	mockCleanTweets := cleaner.MockCleanTweets(nil)
	mockInsert := corpus.MockInsert(nil)

	create := corpus.MakeCreate(mockSelectByCategorizations, mockSelectTweetByID, mockSelectQuoteByID, mockDeleteAll, mockCleanTweets, mockInsert)

	got := create(context.Background(), false)

	assert.Nil(t, got)
}

func TestCreate_successWithPerfectlyBalancedCorpus(t *testing.T) {
	mockSelectByCategorizations := categorized.MockSelectByCategorizations([]categorized.DAO{categorized.MockCategorizedTweetDAO()}, nil)
	mockSelectTweetByID := tweets.MockSelectByID(tweets.MockTweetDAO(), nil)
	mockSelectQuoteByID := quotes.MockSelectByID(quotes.MockTweetQuoteDAO(), nil)
	mockDeleteAll := corpus.MockDeleteAll(nil)
	mockCleanTweets := cleaner.MockCleanTweets(nil)
	mockInsert := corpus.MockInsert(nil)

	create := corpus.MakeCreate(mockSelectByCategorizations, mockSelectTweetByID, mockSelectQuoteByID, mockDeleteAll, mockCleanTweets, mockInsert)

	got := create(context.Background(), true)

	assert.Nil(t, got)
}

func TestCreate_successEvenWhenSelectTweetByIDThrowsError(t *testing.T) {
	mockSelectByCategorizations := categorized.MockSelectByCategorizations([]categorized.DAO{categorized.MockCategorizedTweetDAO()}, nil)
	mockSelectTweetByID := tweets.MockSelectByID(tweets.MockTweetDAO(), errors.New("failed to select tweet by id"))
	mockSelectQuoteByID := quotes.MockSelectByID(quotes.MockTweetQuoteDAO(), nil)
	mockDeleteAll := corpus.MockDeleteAll(nil)
	mockCleanTweets := cleaner.MockCleanTweets(nil)
	mockInsert := corpus.MockInsert(nil)

	create := corpus.MakeCreate(mockSelectByCategorizations, mockSelectTweetByID, mockSelectQuoteByID, mockDeleteAll, mockCleanTweets, mockInsert)

	got := create(context.Background(), false)

	assert.Nil(t, got)
}

func TestCreate_successEvenWhenSelectTweetQuoteByIDThrowsError(t *testing.T) {
	mockSelectByCategorizations := categorized.MockSelectByCategorizations([]categorized.DAO{categorized.MockCategorizedTweetDAO()}, nil)
	mockSelectTweetByID := tweets.MockSelectByID(tweets.MockTweetDAO(), nil)
	mockSelectQuoteByID := quotes.MockSelectByID(quotes.MockTweetQuoteDAO(), errors.New("failed to select quote by id"))
	mockDeleteAll := corpus.MockDeleteAll(nil)
	mockCleanTweets := cleaner.MockCleanTweets(nil)
	mockInsert := corpus.MockInsert(nil)

	create := corpus.MakeCreate(mockSelectByCategorizations, mockSelectTweetByID, mockSelectQuoteByID, mockDeleteAll, mockCleanTweets, mockInsert)

	got := create(context.Background(), false)

	assert.Nil(t, got)
}

func TestCreate_successEvenWhenInsertThrowsError(t *testing.T) {
	mockSelectByCategorizations := categorized.MockSelectByCategorizations([]categorized.DAO{categorized.MockCategorizedTweetDAO()}, nil)
	mockSelectTweetByID := tweets.MockSelectByID(tweets.MockTweetDAO(), nil)
	mockSelectQuoteByID := quotes.MockSelectByID(quotes.MockTweetQuoteDAO(), nil)
	mockDeleteAll := corpus.MockDeleteAll(nil)
	mockCleanTweets := cleaner.MockCleanTweets(nil)
	mockInsert := corpus.MockInsert(errors.New("failed to insert"))

	create := corpus.MakeCreate(mockSelectByCategorizations, mockSelectTweetByID, mockSelectQuoteByID, mockDeleteAll, mockCleanTweets, mockInsert)

	got := create(context.Background(), false)

	assert.Nil(t, got)
}

func TestCreate_failsWhenSelectByCategorizationsThrowsError(t *testing.T) {
	mockSelectByCategorizations := categorized.MockSelectByCategorizations(nil, errors.New("failed to select by categorizations"))
	mockSelectTweetByID := tweets.MockSelectByID(tweets.MockTweetDAO(), nil)
	mockSelectQuoteByID := quotes.MockSelectByID(quotes.MockTweetQuoteDAO(), nil)
	mockDeleteAll := corpus.MockDeleteAll(nil)
	mockCleanTweets := cleaner.MockCleanTweets(nil)
	mockInsert := corpus.MockInsert(nil)

	create := corpus.MakeCreate(mockSelectByCategorizations, mockSelectTweetByID, mockSelectQuoteByID, mockDeleteAll, mockCleanTweets, mockInsert)

	want := corpus.FailedToRetrieveCategorizedTweets
	got := create(context.Background(), false)

	assert.Equal(t, want, got)
}

func TestCreate_failsWhenDeleteAllThrowsError(t *testing.T) {
	mockSelectByCategorizations := categorized.MockSelectByCategorizations([]categorized.DAO{categorized.MockCategorizedTweetDAO()}, nil)
	mockSelectTweetByID := tweets.MockSelectByID(tweets.MockTweetDAO(), nil)
	mockSelectQuoteByID := quotes.MockSelectByID(quotes.MockTweetQuoteDAO(), nil)
	mockDeleteAll := corpus.MockDeleteAll(errors.New("failed to delete all"))
	mockCleanTweets := cleaner.MockCleanTweets(nil)
	mockInsert := corpus.MockInsert(nil)

	create := corpus.MakeCreate(mockSelectByCategorizations, mockSelectTweetByID, mockSelectQuoteByID, mockDeleteAll, mockCleanTweets, mockInsert)

	want := corpus.FailedToCleanUpCorpusTable
	got := create(context.Background(), false)

	assert.Equal(t, want, got)
}

func TestCreate_failsWhenCleanTweetsThrowsError(t *testing.T) {
	mockSelectByCategorizations := categorized.MockSelectByCategorizations([]categorized.DAO{categorized.MockCategorizedTweetDAO()}, nil)
	mockSelectTweetByID := tweets.MockSelectByID(tweets.MockTweetDAO(), nil)
	mockSelectQuoteByID := quotes.MockSelectByID(quotes.MockTweetQuoteDAO(), nil)
	mockDeleteAll := corpus.MockDeleteAll(nil)
	mockCleanTweets := cleaner.MockCleanTweets(errors.New("failed to clean tweets"))
	mockInsert := corpus.MockInsert(nil)

	create := corpus.MakeCreate(mockSelectByCategorizations, mockSelectTweetByID, mockSelectQuoteByID, mockDeleteAll, mockCleanTweets, mockInsert)

	want := corpus.FailedToCleanTweets
	got := create(context.Background(), false)

	assert.Equal(t, want, got)
}
