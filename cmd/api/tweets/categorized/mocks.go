package categorized

import "context"

// MockSelectAllByUserID mocks a SelectAllByUserID function
func MockSelectAllByUserID(dtos []AnalyzedTweetsDTO, err error) SelectAllByUserID {
	return func(ctx context.Context, userID int) ([]AnalyzedTweetsDTO, error) {
		return dtos, err
	}
}

// MockCategorizedTweetsDAO mocks an AnalyzedTweetsDTO
func MockCategorizedTweetsDAO(searchCriteriaID, year, month, analyzed int) AnalyzedTweetsDTO {
	return AnalyzedTweetsDTO{
		SearchCriteriaID: searchCriteriaID,
		Year:             year,
		Month:            month,
		Analyzed:         analyzed,
	}
}

// MockCategorizedTweetsDAOSlice mocks a []categorized.AnalyzedTweetsDTO
func MockCategorizedTweetsDAOSlice() []AnalyzedTweetsDTO {
	return []AnalyzedTweetsDTO{
		MockCategorizedTweetsDAO(1, 2024, 9, 15),
		MockCategorizedTweetsDAO(1, 2025, 1, 10),
		MockCategorizedTweetsDAO(2, 2025, 2, 33),
	}
}

// MockInsertSingle mocks an InsertSingle function
func MockInsertSingle(id int, err error) InsertSingle {
	return func(ctx context.Context, dto DAO) (int, error) {
		return id, err
	}
}

// MockDAO mocks a DAO
func MockDAO() DAO {
	return DAO{
		SearchCriteriaID: 1,
		TweetID:          123456,
		TweetYear:        2024,
		TweetMonth:       5,
		UserID:           789,
		Categorization:   VerdictPositive,
	}
}
