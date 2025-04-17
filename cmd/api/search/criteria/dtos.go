package criteria

type (
	// InformationDTO represents the information of a search criteria
	InformationDTO struct {
		Name  string        `json:"name"`
		ID    int           `json:"id"`
		Years []YearDataDTO `json:"years"`
	}

	// InformationDTOs represents a slice of InformationDTO
	InformationDTOs []InformationDTO

	// YearDataDTO represent a year of the search criteria retrieved data
	YearDataDTO struct {
		Year   int            `json:"year"`
		Months []MonthDataDTO `json:"months"`
	}

	// YearDataDTOs represents a slice of YearDataDTO
	YearDataDTOs []YearDataDTO

	// MonthDataDTO represent a month of a year of the search criteria retrieved data
	MonthDataDTO struct {
		Month          int `json:"month"`
		AnalyzedTweets int `json:"analyzed_tweets"`
		TotalTweets    int `json:"total_tweets"`
	}

	// MonthDataDTOs represents a slice of MonthDataDTO
	MonthDataDTOs []MonthDataDTO
)

func (informationDTOs InformationDTOs) Len() int {
	return len(informationDTOs)
}

func (informationDTOs InformationDTOs) Swap(i, j int) {
	informationDTOs[i], informationDTOs[j] = informationDTOs[j], informationDTOs[i]
}

func (informationDTOs InformationDTOs) Less(i, j int) bool {
	return informationDTOs[i].ID < informationDTOs[j].ID
}

func (yearDataDTOs YearDataDTOs) Len() int {
	return len(yearDataDTOs)
}

func (yearDataDTOs YearDataDTOs) Swap(i, j int) {
	yearDataDTOs[i], yearDataDTOs[j] = yearDataDTOs[j], yearDataDTOs[i]
}

func (yearDataDTOs YearDataDTOs) Less(i, j int) bool {
	return yearDataDTOs[i].Year < yearDataDTOs[j].Year
}

func (monthDataDTOs MonthDataDTOs) Len() int {
	return len(monthDataDTOs)
}

func (monthDataDTOs MonthDataDTOs) Swap(i, j int) {
	monthDataDTOs[i], monthDataDTOs[j] = monthDataDTOs[j], monthDataDTOs[i]
}

func (monthDataDTOs MonthDataDTOs) Less(i, j int) bool {
	return monthDataDTOs[i].Month < monthDataDTOs[j].Month
}
