package categorized

type DAO struct {
	SearchCriteriaID int    `json:"search_criteria_id"`
	TweetID          int    `json:"tweet_id"`
	TweetYear        int    `json:"tweet_year"`
	TweetMonth       int    `json:"tweet_month"`
	UserID           int    `json:"user_id"`
	Categorization   string `json:"categorization"`
}

const (
	VerdictPositive      string = "POSITIVE"
	VerdictIndeterminate string = "INDETERMINATE"
	VerdictNegative      string = "NEGATIVE"
)
