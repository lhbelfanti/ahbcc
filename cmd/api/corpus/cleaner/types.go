package cleaner

// TweetToClean represents a tweet that needs to be scrubbed with the cleaning rules
type TweetToClean struct {
	TweetAuthor    string
	TweetAvatar    *string
	TweetText      *string
	TweetImages    []string
	IsTweetAReply  bool
	QuoteAuthor    *string
	QuoteAvatar    *string
	QuoteText      *string
	QuoteImages    []string
	IsQuoteAReply  *bool
	Categorization string
}
