package criteria

type (
	// SearchCriteriaInformation represents the information of a search criteria
	SearchCriteriaInformation struct {
		Name  string     `json:"name"`
		ID    int        `json:"id"`
		Years []YearData `json:"years"`
	}

	// YearData represent a year of the search criteria retrieved data
	YearData struct {
		Year   int         `json:"year"`
		Months []MonthData `json:"months"`
	}

	// MonthData represent a month of a year of the search criteria retrieved data
	MonthData struct {
		Month          int `json:"month"`
		AnalyzedTweets int `json:"analyzed_tweets"`
		TotalTweets    int `json:"total_tweets"`
	}
)
