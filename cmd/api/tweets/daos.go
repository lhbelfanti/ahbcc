package tweets

import "time"

// DAO represents a tweet from the 'tweets' table
type DAO struct {
	ID               int       `json:"id"`
	StatusID         string    `json:"status_id"`
	Author           string    `json:"author_id"`
	Avatar           *string   `json:"avatar,omitempty"`
	PostedAt         time.Time `json:"posted_at"`
	IsAReply         bool      `json:"is_a_reply"`
	TextContent      *string   `json:"text_content,omitempty"`
	Images           []string  `json:"images"`
	QuoteID          *int      `json:"quote_id,omitempty"`
	SearchCriteriaID int       `json:"search_criteria_id"`
}
