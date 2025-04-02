package counts

// DAO represents a row of the tweets_counts table
type DAO struct {
	SearchCriteriaID int `json:"search_criteria_id"`
	Year             int `json:"tweets_year"`
	Month            int `json:"tweets_month"`
	Total            int `json:"total_tweets"`
}
