package tweets

import "ahbcc/cmd/api/tweets/quotes"

// MockInsert mocks Insert function
func MockInsert(err error) Insert {
	return func(tweet []TweetDTO) error {
		return err
	}
}

// MockTweetDTO mocks a TweetDTO
func MockTweetDTO() TweetDTO {
	hash := "02bd92faa38aaa6cc0ea75e59937a1ef8d6ad3a9f75f3ac4166fef23da9f209b"
	textContent := "test"
	quote := quotes.MockQuoteDTO()
	searchCriteriaID := 1

	return TweetDTO{
		Hash:             &hash,
		IsAReply:         true,
		TextContent:      &textContent,
		Images:           []string{"test1", "test2"},
		Quote:            &quote,
		SearchCriteriaID: &searchCriteriaID,
	}
}

// MockTweets mocks a slice of TweetDTO
func MockTweets() []TweetDTO {
	return []TweetDTO{
		MockTweetDTO(),
		MockTweetDTO(),
	}
}
