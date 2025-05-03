package tweets

// DAO represents a tweet from the 'tweets' table
type DAO struct {
	UUID             int      `json:"uuid"`
	ID               string   `json:"id"`
	Author           string   `json:"author_id"`
	Avatar           string   `json:"avatar"`
	PostedAt         string   `json:"posted_at"`
	IsAReply         bool     `json:"is_a_reply"`
	TextContent      string   `json:"text_content"`
	Images           []string `json:"images"`
	QuoteID          int      `json:"quote_id"`
	SearchCriteriaID int      `json:"search_criteria_id"`
}
