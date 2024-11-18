package quotes

// QuoteDTO represents a quote of a tweet that will be inserted in the 'tweets_quotes' table
type QuoteDTO struct {
	Author      string   `json:"author"`
	Avatar      *string  `json:"avatar,omitempty"`
	PostedAt    string   `json:"posted_at"`
	IsAReply    bool     `json:"is_a_reply"`
	TextContent *string  `json:"text_content,omitempty"`
	Images      []string `json:"images,omitempty"`
}
