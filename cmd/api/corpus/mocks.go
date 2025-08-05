package corpus

// MockDTO mocks a corpus.DTO
func MockDTO() DTO {
	tweetAvatar := "test_avatar"
	tweetText := "test_text"
	quoteAuthor := "quote_author"
	quoteAvatar := "quote_avatar"
	quoteText := "quote_text"
	isQuoteAReply := true

	return DTO{
		TweetAuthor:   "test_author",
		TweetAvatar:   &tweetAvatar,
		TweetText:     &tweetText,
		TweetImages:   []string{"image1.jpg", "image2.jpg"},
		IsTweetAReply: false,
		QuoteAuthor:   &quoteAuthor,
		QuoteAvatar:   &quoteAvatar,
		QuoteText:     &quoteText,
		QuoteImages:   []string{"quote_image1.jpg"},
		IsQuoteAReply: &isQuoteAReply,
	}
}

// MockDAO mocks a corpus.DAO
func MockDAO() DAO {
	tweetAvatar := "test_avatar"
	tweetText := "test_text"
	quoteAuthor := "quote_author"
	quoteAvatar := "quote_avatar"
	quoteText := "quote_text"
	isQuoteAReply := true

	return DAO{
		ID:            1,
		TweetAuthor:   "test_author",
		TweetAvatar:   &tweetAvatar,
		TweetText:     &tweetText,
		TweetImages:   []string{"image1.jpg"},
		IsTweetAReply: false,
		QuoteAuthor:   &quoteAuthor,
		QuoteAvatar:   &quoteAvatar,
		QuoteText:     &quoteText,
		QuoteImages:   []string{"quote_image1.jpg"},
		IsQuoteAReply: &isQuoteAReply,
	}
}
