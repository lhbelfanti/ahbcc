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
	return func(ctx context.Context, dto DTO) (int, error) {
		return id, err
	}
}

// MockInsertCategorizedTweet mocks an InsertCategorizedTweet function
func MockInsertCategorizedTweet(id int, err error) InsertCategorizedTweet {
	return func(ctx context.Context, token string, dto DTO) (int, error) {
		return id, err
	}
}

// MockDTO mocks a DTO
func MockDTO() DTO {
	return DTO{
		SearchCriteriaID: 1,
		TweetID:          123456,
		TweetYear:        2024,
		TweetMonth:       5,
		UserID:           789,
		Categorization:   VerdictPositive,
	}
}
