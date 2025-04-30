package tweets

import "ahbcc/cmd/api/tweets/quotes"

// TweetDTO represents a tweet to be inserted into the 'tweets' table
type TweetDTO struct {
	ID               string           `json:"id"`
	Author           string           `json:"author"`
	Avatar           *string          `json:"avatar,omitempty"`
	PostedAt         string           `json:"posted_at"`
	IsAReply         bool             `json:"is_a_reply"`
	TextContent      *string          `json:"text_content,omitempty"`
	Images           []string         `json:"images,omitempty"`
	QuoteID          *int             `json:"quote_id,omitempty"`
	SearchCriteriaID *int             `json:"search_criteria_id,omitempty"`
	Quote            *quotes.QuoteDTO `json:"quote,omitempty"`
}
