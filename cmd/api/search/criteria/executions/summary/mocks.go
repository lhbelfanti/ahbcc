package summary

// MockExecutionSummaryDAO mocks a summary.DAO
func MockExecutionSummaryDAO(searchCriteriaID, year, month, total int) DAO {
	return DAO{
		SearchCriteriaID: searchCriteriaID,
		Year:             year,
		Month:            month,
		Total:            total,
	}
}

// MockExecutionsSummaryDAOSlice mocks a []summary.DAO
func MockExecutionsSummaryDAOSlice() []DAO {
	return []DAO{
		MockExecutionSummaryDAO(1, 2024, 9, 350),
		MockExecutionSummaryDAO(1, 2025, 1, 1000),
		MockExecutionSummaryDAO(1, 2025, 2, 500),
	}
}
