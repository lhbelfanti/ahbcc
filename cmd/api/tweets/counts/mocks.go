package counts

// MockTweetsCountsDAOSlice mocks a []counts.DAO
func MockTweetsCountsDAOSlice() []DAO {
	return []DAO{
		{
			SearchCriteriaID: 1,
			Year:             2024,
			Month:            9,
			Total:            350,
		},
		{
			SearchCriteriaID: 1,
			Year:             2025,
			Month:            1,
			Total:            1000,
		},
		{
			SearchCriteriaID: 1,
			Year:             2025,
			Month:            2,
			Total:            500,
		},
	}
}
