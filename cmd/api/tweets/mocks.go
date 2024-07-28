package tweets

import "ahbcc/cmd/api/tweets/quotes"

// MockTweetDTO mocks a TweetDTO
func MockTweetDTO() TweetDTO {
	textContent := "test"
	quote := quotes.MockQuoteDTO()

	return TweetDTO{
		Hash:             "02bd92faa38aaa6cc0ea75e59937a1ef8d6ad3a9f75f3ac4166fef23da9f209b",
		IsAReply:         true,
		TextContent:      &textContent,
		Images:           []string{"test1", "test2"},
		Quote:            &quote,
		SearchCriteriaID: 1,
	}
}

// MockTweetDTOSlice mocks a slice of TweetDTO
func MockTweetDTOSlice() []TweetDTO {
	return []TweetDTO{
		MockTweetDTO(),
		MockTweetDTO(),
	}
}
