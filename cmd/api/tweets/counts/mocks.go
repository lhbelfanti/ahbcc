package counts

// MockTweetsCountsDAO mocks a counts.DAO
func MockTweetsCountsDAO(searchCriteriaID, year, month, total int) DAO {
	return DAO{
		SearchCriteriaID: searchCriteriaID,
		Year:             year,
		Month:            month,
		Total:            total,
	}
}

// MockTweetsCountsDAOSlice mocks a []counts.DAO
func MockTweetsCountsDAOSlice() []DAO {
	return []DAO{
		MockTweetsCountsDAO(1, 2024, 9, 350),
		MockTweetsCountsDAO(1, 2025, 1, 1000),
		MockTweetsCountsDAO(1, 2025, 2, 500),
	}
}
