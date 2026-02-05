package cleaner

import "context"

// MockCleanTweets mocks CleanTweets function
func MockCleanTweets(err error) CleanTweets {
	return func(ctx context.Context, tweets []TweetToClean) error {
		return err
	}
}

// MockTweetToClean mocks a TweetToClean
func MockTweetToClean(tweetText string) TweetToClean {
	tweetAvatar := "test_avatar"
	quoteAuthor := "quote_author"
	quoteAvatar := "quote_avatar"
	quoteText := "quote_text"
	isQuoteAReply := true

	return TweetToClean{
		TweetAuthor:    "test_author",
		TweetAvatar:    &tweetAvatar,
		TweetText:      &tweetText,
		TweetImages:    []string{"image1.jpg", "image2.jpg"},
		IsTweetAReply:  false,
		QuoteAuthor:    &quoteAuthor,
		QuoteAvatar:    &quoteAvatar,
		QuoteText:      &quoteText,
		QuoteImages:    []string{"quote_image1.jpg"},
		IsQuoteAReply:  &isQuoteAReply,
		Categorization: "POSITIVE",
	}
}
