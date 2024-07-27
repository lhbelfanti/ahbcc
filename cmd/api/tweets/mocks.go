package tweets

// MockTweetDTO mocks a TweetDTO
func MockTweetDTO() TweetDTO {
	quoteID := 12345

	return TweetDTO{
		Hash:             "02bd92faa38aaa6cc0ea75e59937a1ef8d6ad3a9f75f3ac4166fef23da9f209b",
		IsAReply:         true,
		HasText:          true,
		HasImages:        true,
		TextContent:      "test",
		Images:           []string{"test1", "test2"},
		HasQuote:         true,
		QuoteID:          &quoteID,
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
