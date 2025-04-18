package criteria_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"ahbcc/cmd/api/search/criteria"
	"ahbcc/cmd/api/search/criteria/executions/summary"
	"ahbcc/cmd/api/tweets/categorized"
	"ahbcc/cmd/api/user/session"
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
