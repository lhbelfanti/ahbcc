package categorized

type (
	// AnalyzedTweetsDTO represents the result of the count of all the analyzed tweets by search criteria id, year and month
	AnalyzedTweetsDTO struct {
		SearchCriteriaID int `json:"search_criteria_id"`
		Year             int `json:"tweet_year"`
		Month            int `json:"tweet_month"`
		Analyzed         int `json:"analyzed"`
	}

	// DTO represents a categorized tweet with its properties such as search criteria, tweet details and categorization status
	DTO struct {
		SearchCriteriaID int    `json:"search_criteria_id"`
		TweetID          int    `json:"tweet_id"`
		TweetYear        int    `json:"tweet_year"`
		TweetMonth       int    `json:"tweet_month"`
		UserID           int    `json:"user_id"`
		Categorization   string `json:"categorization"`
	}

	// InsertSingleResponseDTO is the response of the /tweets/categorized/v1 endpoint
	InsertSingleResponseDTO struct {
		ID int `json:"id"`
	}
)

const (
	VerdictPositive      string = "POSITIVE"
	VerdictIndeterminate string = "INDETERMINATE"
	VerdictNegative      string = "NEGATIVE"
)
