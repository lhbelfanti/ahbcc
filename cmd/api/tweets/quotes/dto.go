package quotes

// QuoteDTO represents a quote of a tweet that will be inserted in the 'tweets_quotes' table
type QuoteDTO struct {
	IsAReply    bool     `json:"is_a_reply"`
	TextContent *string  `json:"text_content,omitempty"`
	Images      []string `json:"images,omitempty"`
}
