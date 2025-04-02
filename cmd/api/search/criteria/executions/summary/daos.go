package summary

// DAO represents a row of the search_criteria_executions_summary table
type DAO struct {
	SearchCriteriaID int `json:"search_criteria_id"`
	Year             int `json:"tweets_year"`
	Month            int `json:"tweets_month"`
	Total            int `json:"total_tweets"`
}
