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

// MockTweetDTO mocks a TweetDTO
func MockTweetDTO() TweetDTO {
	hash := "02bd92faa38aaa6cc0ea75e59937a1ef8d6ad3a9f75f3ac4166fef23da9f209b"
	avatar := "https://testuseravatar.com"

	textContent := "test"
	quote := quotes.MockQuoteDTO()
	searchCriteriaID := 1

	return TweetDTO{
		UUID:             "1234567890987654321",
		Hash:             &hash,
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
