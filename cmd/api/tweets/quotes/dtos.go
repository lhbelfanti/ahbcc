package quotes

import "time"

type (
	// QuoteDTO represents a quote of a tweet that will be inserted in the 'tweets_quotes' table
	QuoteDTO struct {
		Author      string   `json:"author"`
		Avatar      *string  `json:"avatar,omitempty"`
		PostedAt    string   `json:"posted_at"`
		IsAReply    bool     `json:"is_a_reply"`
		TextContent *string  `json:"text_content,omitempty"`
		Images      []string `json:"images,omitempty"`
	}

	// CustomQuoteDTO represents a quote of a tweet obtained from the database. Note that the PostedAt param type
	// is different from the QuoteDTO.PostedAt param type
	CustomQuoteDTO struct {
		Author      string    `json:"author"`
		Avatar      *string   `json:"avatar,omitempty"`
		PostedAt    time.Time `json:"posted_at"`
		IsAReply    bool      `json:"is_a_reply"`
		TextContent *string   `json:"text_content,omitempty"`
		Images      []string  `json:"images,omitempty"`
	}
)
