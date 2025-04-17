package categorized

// DAO represents the result of the count of all the analyzed tweets by search criteria id, year and month
type DAO struct {
	SearchCriteriaID int `json:"search_criteria_id"`
	Year             int `json:"tweet_year"`
	Month            int `json:"tweet_month"`
	Analyzed         int `json:"analyzed"`
}
