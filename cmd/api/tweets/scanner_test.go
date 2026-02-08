package tweets_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/lhbelfanti/corpus-creator/cmd/api/tweets"
	"github.com/lhbelfanti/corpus-creator/internal/database"
)

func TestCustomScanner_success(t *testing.T) {
	mockPgxCollectableRow := new(database.MockPgxCollectableRow)
	mockTweetDTO := tweets.MockCustomTweetDTO()
	mockTweetQuoteCollectedRow := tweets.MockTweetCollectedRow(mockTweetDTO)
	database.MockPgxCollectableRowMethods(mockPgxCollectableRow, mockTweetQuoteCollectedRow, t)

	customScanner := tweets.CustomScanner()

	want := mockTweetDTO
	got, err := customScanner(mockPgxCollectableRow)

	assert.Nil(t, err)
	assert.Equal(t, want, got)

	mockPgxCollectableRow.AssertExpectations(t)
}

func TestCustomScanner_successWithNilQuote(t *testing.T) {
	mockPgxCollectableRow := new(database.MockPgxCollectableRow)
	mockTweetDTO := tweets.MockCustomTweetDTO()
	mockTweetDTO.QuoteID = nil
	mockTweetDTO.Quote = nil
	mockTweetQuoteCollectedRow := tweets.MockTweetCollectedRow(mockTweetDTO)
	database.MockPgxCollectableRowMethods(mockPgxCollectableRow, mockTweetQuoteCollectedRow, t)

	customScanner := tweets.CustomScanner()

	want := mockTweetDTO
	got, err := customScanner(mockPgxCollectableRow)

	assert.Nil(t, err)
	assert.Equal(t, want, got)
	mockPgxCollectableRow.AssertExpectations(t)
}

func TestCustomScanner_successWithNilOrEmptyQuoteValues(t *testing.T) {
	mockPgxCollectableRow := new(database.MockPgxCollectableRow)
	mockTweetDTO := tweets.MockCustomTweetDTO()
	mockTweetDTO.Quote.Avatar = nil
	emptyString := ""
	mockTweetDTO.Quote.TextContent = &emptyString
	mockTweetQuoteCollectedRow := tweets.MockTweetCollectedRow(mockTweetDTO)
	database.MockPgxCollectableRowMethods(mockPgxCollectableRow, mockTweetQuoteCollectedRow, t)
	mockTweetDTOWant := tweets.MockCustomTweetDTO()
	mockTweetDTOWant.PostedAt = mockTweetDTO.PostedAt
	mockTweetDTOWant.Quote.Avatar = nil
	mockTweetDTOWant.Quote.TextContent = nil
	mockTweetDTOWant.Quote.PostedAt = mockTweetDTO.Quote.PostedAt

	customScanner := tweets.CustomScanner()

	want := mockTweetDTOWant
	got, err := customScanner(mockPgxCollectableRow)

	assert.Nil(t, err)
	assert.Equal(t, want, got)
	mockPgxCollectableRow.AssertExpectations(t)
}

func TestCustomScanner_failsWhenScanThrowsError(t *testing.T) {
	mockPgxCollectableRow := new(database.MockPgxCollectableRow)
	want := errors.New("scan error")
	mockPgxCollectableRow.On("Scan", mock.Anything).Return(want)

	customScanner := tweets.CustomScanner()

	_, got := customScanner(mockPgxCollectableRow)

	assert.Equal(t, want, got)
	mockPgxCollectableRow.AssertExpectations(t)
}
