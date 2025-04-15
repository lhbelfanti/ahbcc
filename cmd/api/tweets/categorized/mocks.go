package categorized

// MockAnalyzedTweetsDAOSlice mocks an []AnalyzedTweetsDAO
func MockAnalyzedTweetsDAOSlice() []AnalyzedTweetsDAO {
	return []AnalyzedTweetsDAO{
		{
			SearchCriteriaID: 1,
			Year:             2025,
			Month:            4,
			Analyzed:         15,
		}, {
			SearchCriteriaID: 1,
			Year:             2020,
			Month:            3,
			Analyzed:         10,
		},
		{
			SearchCriteriaID: 2,
			Year:             2010,
			Month:            12,
			Analyzed:         1,
		},
	}
}
