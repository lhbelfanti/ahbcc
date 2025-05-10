package categorized

// DAO represents a row from the 'categorized_tweets' table
type DAO struct {
	ID               int    `json:"id"`
	SearchCriteriaID int    `json:"search_criteria_id"`
	TweetID          int    `json:"tweet_id"`
	TweetYear        int    `json:"tweet_year"`
	TweetMonth       int    `json:"tweet_month"`
	UserID           int    `json:"user_id"`
	Categorization   string `json:"categorization"`
}
