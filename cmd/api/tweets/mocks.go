package tweets

import (
	"context"
	"time"

	"ahbcc/cmd/api/tweets/quotes"
)

// MockInsert mocks Insert function
func MockInsert(err error) Insert {
	return func(ctx context.Context, tweet []TweetDTO) error {
		return err
	}
}

// MockSelectBySearchCriteriaIDYearAndMonth mocks SelectBySearchCriteriaIDYearAndMonth function
func MockSelectBySearchCriteriaIDYearAndMonth(tweets []CustomTweetDTO, err error) SelectBySearchCriteriaIDYearAndMonth {
	return func(ctx context.Context, searchCriteriaID, year, month, limit int, token string) ([]CustomTweetDTO, error) {
		return tweets, err
	}
}

// MockTweetDTO mocks a TweetDTO
func MockTweetDTO() TweetDTO {
	avatar := "https://testuseravatar.com"

	textContent := "test"
	searchCriteriaID := 1
	quote := quotes.MockQuoteDTO()

	return TweetDTO{
		ID:               "1234567890987654321",
		IsAReply:         true,
		Author:           "TestAuthor",
		Avatar:           &avatar,
		PostedAt:         "2024-11-18T15:04:05Z",
		TextContent:      &textContent,
		Images:           []string{"test1", "test2"},
		SearchCriteriaID: &searchCriteriaID,
		Quote:            &quote,
	}
}

// MockTweetsDTOs mocks a slice of TweetDTO
func MockTweetsDTOs() []TweetDTO {
	return []TweetDTO{
		MockTweetDTO(),
		MockTweetDTO(),
	}
}

// MockCustomTweetDTO mocks a CustomTweetDTO
func MockCustomTweetDTO() CustomTweetDTO {
	avatar := "https://testuseravatar.com"

	textContent := "test"
	quoteID := 3
	searchCriteriaID := 1
	quote := quotes.MockCustomQuoteDTO()

	return CustomTweetDTO{
		ID:               "1234567890987654321",
		IsAReply:         true,
		Author:           "TestAuthor",
		Avatar:           &avatar,
		PostedAt:         time.Now(),
		TextContent:      &textContent,
		Images:           []string{"test1", "test2"},
		QuoteID:          &quoteID,
		SearchCriteriaID: &searchCriteriaID,
		Quote:            &quote,
	}
}

// MockCustomTweetDTOs mocks a slice of CustomTweetDTO
func MockCustomTweetDTOs() []CustomTweetDTO {
	return []CustomTweetDTO{
		MockCustomTweetDTO(),
		MockCustomTweetDTO(),
	}
}

// MockTweetCollectedRow mocks a row with the Tweet, and its Quote, information, obtained from a select
func MockTweetCollectedRow(tweet CustomTweetDTO) []any {
	row := []any{
		tweet.ID,
		tweet.Author,
		tweet.Avatar,
		tweet.PostedAt,
		tweet.IsAReply,
		tweet.TextContent,
		tweet.Images,
		tweet.QuoteID,
		tweet.SearchCriteriaID,
	}

	if tweet.Quote != nil {
		row = append(row,
			tweet.Quote.Author,
			tweet.Quote.Avatar,
			tweet.Quote.PostedAt,
			tweet.Quote.IsAReply,
			tweet.Quote.TextContent,
			tweet.Quote.Images,
		)
	} else {
		row = append(row, nil, nil, nil, nil, nil, nil)
	}

	return row
}
