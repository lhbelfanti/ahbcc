package tweets

import (
	"context"

	"ahbcc/cmd/api/tweets/quotes"
)

// MockInsert mocks Insert function
func MockInsert(err error) Insert {
	return func(ctx context.Context, tweet []TweetDTO) error {
		return err
	}
}

// MockSelectBySearchCriteriaIDYearAndMonth mocks SelectBySearchCriteriaIDYearAndMonth function
func MockSelectBySearchCriteriaIDYearAndMonth(tweets []TweetDTO, err error) SelectBySearchCriteriaIDYearAndMonth {
	return func(ctx context.Context, searchCriteriaID, year, month, limit int, token string) ([]TweetDTO, error) {
		return tweets, err
	}
}

// MockTweetDTO mocks a TweetDTO
func MockTweetDTO() TweetDTO {
	avatar := "https://testuseravatar.com"

	textContent := "test"
	quote := quotes.MockQuoteDTO()
	searchCriteriaID := 1

	return TweetDTO{
		ID:               "1234567890987654321",
		IsAReply:         true,
		Author:           "TestAuthor",
		Avatar:           &avatar,
		PostedAt:         "2024-11-18T15:04:05Z",
		TextContent:      &textContent,
		Images:           []string{"test1", "test2"},
		Quote:            &quote,
		SearchCriteriaID: &searchCriteriaID,
	}
}

// MockTweetsDTOs mocks a slice of TweetDTO
func MockTweetsDTOs() []TweetDTO {
	return []TweetDTO{
		MockTweetDTO(),
		MockTweetDTO(),
	}
}
