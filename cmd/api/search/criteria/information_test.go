package criteria_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/lhbelfanti/corpus-creator/cmd/api/search/criteria"
	"github.com/lhbelfanti/corpus-creator/cmd/api/search/criteria/executions/summary"
	"github.com/lhbelfanti/corpus-creator/cmd/api/tweets/categorized"
	"github.com/lhbelfanti/corpus-creator/cmd/api/user/session"
)

func TestInformation_success(t *testing.T) {
	mockSelectUserIDByToken := session.MockSelectUserIDByToken(1, nil)
	mockExecutionsSummaryDAOSlice := summary.MockExecutionsSummaryDAOSlice()
	mockSelectAllCriteriaExecutionsSummaries := summary.MockSelectAll(mockExecutionsSummaryDAOSlice, nil)
	mockCriteriaDAOSlice := criteria.MockCriteriaDAOSlice()
	mockSelectAllSearchCriteria := criteria.MockSelectAll(mockCriteriaDAOSlice, nil)
	mockCategorizedTweetsDAOSlice := categorized.MockCategorizedTweetsDAOSlice()
	mockSelectAllByUserID := categorized.MockSelectAllByUserID(mockCategorizedTweetsDAOSlice, nil)

	information := criteria.MakeInformation(mockSelectUserIDByToken, mockSelectAllCriteriaExecutionsSummaries, mockSelectAllSearchCriteria, mockSelectAllByUserID)

	want := criteria.MockInformationDTOs()
	got, err := information(context.Background(), "token")

	assert.Nil(t, err)
	assert.Equal(t, want, got)
}

func TestInformation_failsWhenSelectUserIDByTokenThrowsError(t *testing.T) {
	mockSelectUserIDByToken := session.MockSelectUserIDByToken(1, errors.New("failed to execute select user ID by token"))
	mockExecutionsSummaryDAOSlice := summary.MockExecutionsSummaryDAOSlice()
	mockSelectAllCriteriaExecutionsSummaries := summary.MockSelectAll(mockExecutionsSummaryDAOSlice, nil)
	mockCriteriaDAOSlice := criteria.MockCriteriaDAOSlice()
	mockSelectAllSearchCriteria := criteria.MockSelectAll(mockCriteriaDAOSlice, nil)
	mockCategorizedTweetsDAOSlice := categorized.MockCategorizedTweetsDAOSlice()
	mockSelectAllByUserID := categorized.MockSelectAllByUserID(mockCategorizedTweetsDAOSlice, nil)

	information := criteria.MakeInformation(mockSelectUserIDByToken, mockSelectAllCriteriaExecutionsSummaries, mockSelectAllSearchCriteria, mockSelectAllByUserID)

	want := criteria.FailedToRetrieveUserID
	_, got := information(context.Background(), "token")

	assert.Equal(t, want, got)
}

func TestInformation_failsWhenSelectAllCriteriaExecutionsSummariesThrowsError(t *testing.T) {
	mockSelectUserIDByToken := session.MockSelectUserIDByToken(1, nil)
	mockExecutionsSummaryDAOSlice := summary.MockExecutionsSummaryDAOSlice()
	mockSelectAllCriteriaExecutionsSummaries := summary.MockSelectAll(mockExecutionsSummaryDAOSlice, errors.New("failed to execute select all criteria executions summaries"))
	mockCriteriaDAOSlice := criteria.MockCriteriaDAOSlice()
	mockSelectAllSearchCriteria := criteria.MockSelectAll(mockCriteriaDAOSlice, nil)
	mockCategorizedTweetsDAOSlice := categorized.MockCategorizedTweetsDAOSlice()
	mockSelectAllByUserID := categorized.MockSelectAllByUserID(mockCategorizedTweetsDAOSlice, nil)

	information := criteria.MakeInformation(mockSelectUserIDByToken, mockSelectAllCriteriaExecutionsSummaries, mockSelectAllSearchCriteria, mockSelectAllByUserID)

	want := criteria.FailedToRetrieveSearchCriteriaExecutionsSummaries
	_, got := information(context.Background(), "token")

	assert.Equal(t, want, got)
}

func TestInformation_failsWhenSelectAllSearchCriteriaThrowsError(t *testing.T) {
	mockSelectUserIDByToken := session.MockSelectUserIDByToken(1, nil)
	mockExecutionsSummaryDAOSlice := summary.MockExecutionsSummaryDAOSlice()
	mockSelectAllCriteriaExecutionsSummaries := summary.MockSelectAll(mockExecutionsSummaryDAOSlice, nil)
	mockCriteriaDAOSlice := criteria.MockCriteriaDAOSlice()
	mockSelectAllSearchCriteria := criteria.MockSelectAll(mockCriteriaDAOSlice, errors.New("failed to execute select all search criteria"))
	mockCategorizedTweetsDAOSlice := categorized.MockCategorizedTweetsDAOSlice()
	mockSelectAllByUserID := categorized.MockSelectAllByUserID(mockCategorizedTweetsDAOSlice, nil)

	information := criteria.MakeInformation(mockSelectUserIDByToken, mockSelectAllCriteriaExecutionsSummaries, mockSelectAllSearchCriteria, mockSelectAllByUserID)

	want := criteria.FailedToRetrieveSearchCriteria
	_, got := information(context.Background(), "token")

	assert.Equal(t, want, got)
}

func TestInformation_failsWhenSelectAllCategorizedTweetsThrowsError(t *testing.T) {
	mockSelectUserIDByToken := session.MockSelectUserIDByToken(1, nil)
	mockExecutionsSummaryDAOSlice := summary.MockExecutionsSummaryDAOSlice()
	mockSelectAllCriteriaExecutionsSummaries := summary.MockSelectAll(mockExecutionsSummaryDAOSlice, nil)
	mockCriteriaDAOSlice := criteria.MockCriteriaDAOSlice()
	mockSelectAllSearchCriteria := criteria.MockSelectAll(mockCriteriaDAOSlice, nil)
	mockCategorizedTweetsDAOSlice := categorized.MockCategorizedTweetsDAOSlice()
	mockSelectAllByUserID := categorized.MockSelectAllByUserID(mockCategorizedTweetsDAOSlice, errors.New("failed to execute select all categorized tweets"))

	information := criteria.MakeInformation(mockSelectUserIDByToken, mockSelectAllCriteriaExecutionsSummaries, mockSelectAllSearchCriteria, mockSelectAllByUserID)

	want := criteria.FailedToRetrieveCategorizedTweetsByUserID
	_, got := information(context.Background(), "token")

	assert.Equal(t, want, got)
}

func TestSummarizedInformation_success(t *testing.T) {
	mockSelectUserIDByToken := session.MockSelectUserIDByToken(1, nil)
	mockCriteriaDAO := criteria.MockCriteriaDAO()
	mockSelectCriteriaByID := criteria.MockSelectByID(mockCriteriaDAO, nil)
	mockExecutionsSummaryDAOSlice := summary.MockExecutionsSummaryDAOSlice()
	mockSelectAllCriteriaExecutionsSummaries := summary.MockSelectAll(mockExecutionsSummaryDAOSlice, nil)
	mockCategorizedTweetsDAOSlice := categorized.MockCategorizedTweetsDAOSlice()
	mockSelectAllByUserID := categorized.MockSelectAllByUserID(mockCategorizedTweetsDAOSlice, nil)

	summarizedInformation := criteria.MakeSummarizedInformation(mockSelectUserIDByToken, mockSelectCriteriaByID, mockSelectAllCriteriaExecutionsSummaries, mockSelectAllByUserID)

	for _, test := range []struct {
		year             int
		month            int
		expectedAnalyzed int
		expectedTotal    int
	}{
		{year: 2024, month: 9, expectedAnalyzed: 15, expectedTotal: 350},
		{year: 0, month: 9, expectedAnalyzed: 15, expectedTotal: 350},
		{year: 2024, month: 0, expectedAnalyzed: 15, expectedTotal: 350},
		{year: 0, month: 0, expectedAnalyzed: 25, expectedTotal: 1850},
	} {
		want := criteria.MockSummarizedInformationDTO(test.year, test.month, test.expectedAnalyzed, test.expectedTotal)
		got, err := summarizedInformation(context.Background(), "token", 1, test.year, test.month)

		assert.Nil(t, err)
		assert.Equal(t, want, got)
	}
}

func TestSummarizedInformation_failsWhenSelectUserIDByTokenThrowsError(t *testing.T) {
	mockSelectUserIDByToken := session.MockSelectUserIDByToken(1, errors.New("failed to execute select user ID by token"))
	mockCriteriaDAO := criteria.MockCriteriaDAO()
	mockSelectCriteriaByID := criteria.MockSelectByID(mockCriteriaDAO, nil)
	mockExecutionsSummaryDAOSlice := summary.MockExecutionsSummaryDAOSlice()
	mockSelectAllCriteriaExecutionsSummaries := summary.MockSelectAll(mockExecutionsSummaryDAOSlice, nil)
	mockCategorizedTweetsDAOSlice := categorized.MockCategorizedTweetsDAOSlice()
	mockSelectAllByUserID := categorized.MockSelectAllByUserID(mockCategorizedTweetsDAOSlice, nil)

	summarizedInformation := criteria.MakeSummarizedInformation(mockSelectUserIDByToken, mockSelectCriteriaByID, mockSelectAllCriteriaExecutionsSummaries, mockSelectAllByUserID)

	want := criteria.FailedToRetrieveUserID
	_, got := summarizedInformation(context.Background(), "token", 1, 2024, 9)

	assert.Equal(t, want, got)
}

func TestSummarizedInformation_failsWhenSelectCriteriaByIDThrowsError(t *testing.T) {
	mockSelectUserIDByToken := session.MockSelectUserIDByToken(1, nil)
	mockCriteriaDAO := criteria.MockCriteriaDAO()
	mockSelectCriteriaByID := criteria.MockSelectByID(mockCriteriaDAO, errors.New("failed to execute select criteria by ID"))
	mockExecutionsSummaryDAOSlice := summary.MockExecutionsSummaryDAOSlice()
	mockSelectAllCriteriaExecutionsSummaries := summary.MockSelectAll(mockExecutionsSummaryDAOSlice, nil)
	mockCategorizedTweetsDAOSlice := categorized.MockCategorizedTweetsDAOSlice()
	mockSelectAllByUserID := categorized.MockSelectAllByUserID(mockCategorizedTweetsDAOSlice, nil)

	summarizedInformation := criteria.MakeSummarizedInformation(mockSelectUserIDByToken, mockSelectCriteriaByID, mockSelectAllCriteriaExecutionsSummaries, mockSelectAllByUserID)

	want := criteria.FailedToRetrieveSearchCriteria
	_, got := summarizedInformation(context.Background(), "token", 1, 2024, 9)

	assert.Equal(t, want, got)
}

func TestSummarizedInformation_failsWhenSelectAllCriteriaExecutionsSummariesThrowsError(t *testing.T) {
	mockSelectUserIDByToken := session.MockSelectUserIDByToken(1, nil)
	mockCriteriaDAO := criteria.MockCriteriaDAO()
	mockSelectCriteriaByID := criteria.MockSelectByID(mockCriteriaDAO, nil)
	mockExecutionsSummaryDAOSlice := summary.MockExecutionsSummaryDAOSlice()
	mockSelectAllCriteriaExecutionsSummaries := summary.MockSelectAll(mockExecutionsSummaryDAOSlice, errors.New("failed to execute select all criteria executions summaries"))
	mockCategorizedTweetsDAOSlice := categorized.MockCategorizedTweetsDAOSlice()
	mockSelectAllByUserID := categorized.MockSelectAllByUserID(mockCategorizedTweetsDAOSlice, nil)

	summarizedInformation := criteria.MakeSummarizedInformation(mockSelectUserIDByToken, mockSelectCriteriaByID, mockSelectAllCriteriaExecutionsSummaries, mockSelectAllByUserID)

	want := criteria.FailedToRetrieveSearchCriteriaExecutionsSummaries
	_, got := summarizedInformation(context.Background(), "token", 1, 2024, 9)

	assert.Equal(t, want, got)
}

func TestSummarizedInformation_failsWhenSelectAllCategorizedTweetsThrowsError(t *testing.T) {
	mockSelectUserIDByToken := session.MockSelectUserIDByToken(1, nil)
	mockCriteriaDAO := criteria.MockCriteriaDAO()
	mockSelectCriteriaByID := criteria.MockSelectByID(mockCriteriaDAO, nil)
	mockExecutionsSummaryDAOSlice := summary.MockExecutionsSummaryDAOSlice()
	mockSelectAllCriteriaExecutionsSummaries := summary.MockSelectAll(mockExecutionsSummaryDAOSlice, nil)
	mockCategorizedTweetsDAOSlice := categorized.MockCategorizedTweetsDAOSlice()
	mockSelectAllByUserID := categorized.MockSelectAllByUserID(mockCategorizedTweetsDAOSlice, errors.New("failed to execute select all categorized tweets"))

	summarizedInformation := criteria.MakeSummarizedInformation(mockSelectUserIDByToken, mockSelectCriteriaByID, mockSelectAllCriteriaExecutionsSummaries, mockSelectAllByUserID)

	want := criteria.FailedToRetrieveCategorizedTweetsByUserID
	_, got := summarizedInformation(context.Background(), "token", 1, 2024, 9)

	assert.Equal(t, want, got)
}
