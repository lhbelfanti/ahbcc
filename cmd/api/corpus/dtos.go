package corpus

// DTO represents a corpus entry to be inserted into the 'corpus' table
type DTO struct {
	TweetAuthor   string   `json:"tweet_author"`
	TweetAvatar   *string  `json:"tweet_avatar,omitempty"`
	TweetText     *string  `json:"tweet_text,omitempty"`
	TweetImages   []string `json:"tweet_images,omitempty"`
	IsTweetAReply bool     `json:"is_tweet_a_reply"`
	QuoteAuthor   *string  `json:"quote_author,omitempty"`
	QuoteAvatar   *string  `json:"quote_avatar,omitempty"`
	QuoteText     *string  `json:"quote_text,omitempty"`
	QuoteImages   []string `json:"quote_images,omitempty"`
	IsQuoteAReply *bool    `json:"is_quote_a_reply,omitempty"`
}
