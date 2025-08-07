package corpus

import "context"

// MockInsert mocks Insert function
func MockInsert(err error) Insert {
	return func(ctx context.Context, entry DTO) (int, error) {
		return 1, err
	}
}

// MockSelectAll mocks SelectAll function
func MockSelectAll(entries []DAO, err error) SelectAll {
	return func(ctx context.Context) ([]DAO, error) {
		return entries, err
	}
}

// MockDeleteAll mocks DeleteAll function
func MockDeleteAll(err error) DeleteAll {
	return func(ctx context.Context) error {
		return err
	}
}

// MockCreate mocks Create function
func MockCreate(err error) Create {
	return func(ctx context.Context) error {
		return err
	}
}

// MockDTO mocks a corpus.DTO
func MockDTO() DTO {
	tweetAvatar := "test_avatar"
	tweetText := "test_text"
	quoteAuthor := "quote_author"
	quoteAvatar := "quote_avatar"
	quoteText := "quote_text"
	isQuoteAReply := true

	return DTO{
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

// MockDAO mocks a corpus.DAO
func MockDAO() DAO {
	tweetAvatar := "test_avatar"
	tweetText := "test_text"
	quoteAuthor := "quote_author"
	quoteAvatar := "quote_avatar"
	quoteText := "quote_text"
	isQuoteAReply := true

	return DAO{
		ID:             1,
		TweetAuthor:    "test_author",
		TweetAvatar:    &tweetAvatar,
		TweetText:      &tweetText,
		TweetImages:    []string{"image1.jpg"},
		IsTweetAReply:  false,
		QuoteAuthor:    &quoteAuthor,
		QuoteAvatar:    &quoteAvatar,
		QuoteText:      &quoteText,
		QuoteImages:    []string{"quote_image1.jpg"},
		IsQuoteAReply:  &isQuoteAReply,
		Categorization: "POSITIVE",
	}
}
