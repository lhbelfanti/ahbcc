package categorized

import "context"

// MockSelectAllByUserID mocks a SelectAllByUserID function
func MockSelectAllByUserID(daos []DAO, err error) SelectAllByUserID {
	return func(ctx context.Context, userID int) ([]DAO, error) {
		return daos, err
	}
}

// MockCategorizedTweetsDAO mocks a DAO
func MockCategorizedTweetsDAO(searchCriteriaID, year, month, analyzed int) DAO {
	return DAO{
		SearchCriteriaID: searchCriteriaID,
		Year:             year,
		Month:            month,
		Analyzed:         analyzed,
	}
}

// MockCategorizedTweetsDAOSlice mocks a []categorized.DAO
func MockCategorizedTweetsDAOSlice() []DAO {
	return []DAO{
		MockCategorizedTweetsDAO(1, 2024, 9, 15),
		MockCategorizedTweetsDAO(1, 2025, 1, 10),
		MockCategorizedTweetsDAO(2, 2025, 2, 33),
	}
}
