package quotes

// DAO represents a quote from the 'tweets_quotes' table
type DAO struct {
	ID          int      `json:"id"`
	Author      string   `json:"author"`
	Avatar      *string  `json:"avatar,omitempty"`
	PostedAt    string   `json:"posted_at"`
	IsAReply    bool     `json:"is_a_reply"`
	TextContent *string  `json:"text_content,omitempty"`
	Images      []string `json:"images,omitempty"`
}
