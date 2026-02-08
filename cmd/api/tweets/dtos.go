package tweets

import (
	"time"

	"github.com/lhbelfanti/corpus-creator/cmd/api/tweets/quotes"
)

type (
	// TweetDTO represents a tweet to be inserted into the 'tweets' table
	TweetDTO struct {
		StatusID         string           `json:"status_id"`
		Author           string           `json:"author"`
		Avatar           *string          `json:"avatar,omitempty"`
		PostedAt         string           `json:"posted_at"`
		IsAReply         bool             `json:"is_a_reply"`
		TextContent      *string          `json:"text_content,omitempty"`
		Images           []string         `json:"images,omitempty"`
		SearchCriteriaID *int             `json:"search_criteria_id,omitempty"`
		Quote            *quotes.QuoteDTO `json:"quote,omitempty"`
	}

	// CustomTweetDTO represents a tweet obtained from the database with the extra property *quotes.CustomQuoteDTO
	CustomTweetDTO struct {
		ID               int                    `json:"id"`
		StatusID         string                 `json:"status_id"`
		Author           string                 `json:"author"`
		Avatar           *string                `json:"avatar,omitempty"`
		PostedAt         time.Time              `json:"posted_at"`
		IsAReply         bool                   `json:"is_a_reply"`
		TextContent      *string                `json:"text_content,omitempty"`
		Images           []string               `json:"images,omitempty"`
		QuoteID          *int                   `json:"quote_id,omitempty"`
		SearchCriteriaID *int                   `json:"search_criteria_id,omitempty"`
		Quote            *quotes.CustomQuoteDTO `json:"quote,omitempty"`
	}
)
