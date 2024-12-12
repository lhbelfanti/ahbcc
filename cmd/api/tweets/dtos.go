package tweets

import "ahbcc/cmd/api/tweets/quotes"

// TweetDTO represents a tweet to be inserted into the 'tweets' table
type TweetDTO struct {
	UUID             string           `json:"uuid"`
	Author           string           `json:"author"`
	Avatar           *string          `json:"avatar,omitempty"`
	PostedAt         string           `json:"posted_at"`
	IsAReply         bool             `json:"is_a_reply"`
	TextContent      *string          `json:"text_content,omitempty"`
	Images           []string         `json:"images,omitempty"`
	Quote            *quotes.QuoteDTO `json:"quote,omitempty"`
	SearchCriteriaID *int             `json:"search_criteria_id,omitempty"`
}
