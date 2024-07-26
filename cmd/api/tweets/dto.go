package tweets

// TweetDTO represents a tweet to be inserted into the database
type TweetDTO struct {
	Hash             string   `json:"hash"`
	IsAReply         bool     `json:"is_a_reply"`
	HasText          bool     `json:"has_text"`
	HasImages        bool     `json:"has_images"`
	TextContent      string   `json:"text_content"`
	Images           []string `json:"images"`
	HasQuote         bool     `json:"has_quote"`
	QuoteID          *int     `json:"quote_id"`
	SearchCriteriaID int      `json:"search_criteria_id"`
}
