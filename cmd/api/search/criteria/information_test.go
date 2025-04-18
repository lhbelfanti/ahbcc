package criteria_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"ahbcc/cmd/api/search/criteria"
	"ahbcc/cmd/api/search/criteria/executions/summary"
	"ahbcc/cmd/api/tweets/categorized"
)

func TestInformation_success(t *testing.T) {
	mockExecutionsSummaryDAOSlice := summary.MockExecutionsSummaryDAOSlice()
	mockSelectAllCriteriaExecutionsSummaries := summary.MockSelectAll(mockExecutionsSummaryDAOSlice, nil)
	mockCriteriaDAOSlice := criteria.MockCriteriaDAOSlice()
	mockSelectAllSearchCriteria := criteria.MockSelectAll(mockCriteriaDAOSlice, nil)
	mockCategorizedTweetsDAOSlice := categorized.MockCategorizedTweetsDAOSlice()
	mockSelectAllByUserID := categorized.MockSelectAllByUserID(mockCategorizedTweetsDAOSlice, nil)

	information := criteria.MakeInformation(mockSelectAllCriteriaExecutionsSummaries, mockSelectAllSearchCriteria, mockSelectAllByUserID)

	want := criteria.MockInformationDTOs()
	got, err := information(context.Background(), 1)

	assert.Nil(t, err)
	assert.Equal(t, want, got)
}

func TestInformation_failsWhenSelectAllCriteriaExecutionsSummariesThrowsError(t *testing.T) {
	mockExecutionsSummaryDAOSlice := summary.MockExecutionsSummaryDAOSlice()
	mockSelectAllCriteriaExecutionsSummaries := summary.MockSelectAll(mockExecutionsSummaryDAOSlice, errors.New("failed to execute select all criteria executions summaries"))
	mockCriteriaDAOSlice := criteria.MockCriteriaDAOSlice()
	mockSelectAllSearchCriteria := criteria.MockSelectAll(mockCriteriaDAOSlice, nil)
	mockCategorizedTweetsDAOSlice := categorized.MockCategorizedTweetsDAOSlice()
	mockSelectAllByUserID := categorized.MockSelectAllByUserID(mockCategorizedTweetsDAOSlice, nil)

	information := criteria.MakeInformation(mockSelectAllCriteriaExecutionsSummaries, mockSelectAllSearchCriteria, mockSelectAllByUserID)

	want := criteria.FailedToRetrieveSearchCriteriaExecutionsSummaries
	_, got := information(context.Background(), 1)

	assert.Equal(t, want, got)
}

func TestInformation_failsWhenSelectAllSearchCriteriaThrowsError(t *testing.T) {
	mockExecutionsSummaryDAOSlice := summary.MockExecutionsSummaryDAOSlice()
	mockSelectAllCriteriaExecutionsSummaries := summary.MockSelectAll(mockExecutionsSummaryDAOSlice, nil)
	mockCriteriaDAOSlice := criteria.MockCriteriaDAOSlice()
	mockSelectAllSearchCriteria := criteria.MockSelectAll(mockCriteriaDAOSlice, errors.New("failed to execute select all search criteria"))
	mockCategorizedTweetsDAOSlice := categorized.MockCategorizedTweetsDAOSlice()
	mockSelectAllByUserID := categorized.MockSelectAllByUserID(mockCategorizedTweetsDAOSlice, nil)

	information := criteria.MakeInformation(mockSelectAllCriteriaExecutionsSummaries, mockSelectAllSearchCriteria, mockSelectAllByUserID)

	want := criteria.FailedToRetrieveSearchCriteria
	_, got := information(context.Background(), 1)

	assert.Equal(t, want, got)
}

func TestInformation_failsWhenSelectAllCategorizedTweetsThrowsError(t *testing.T) {
	mockExecutionsSummaryDAOSlice := summary.MockExecutionsSummaryDAOSlice()
	mockSelectAllCriteriaExecutionsSummaries := summary.MockSelectAll(mockExecutionsSummaryDAOSlice, nil)
	mockCriteriaDAOSlice := criteria.MockCriteriaDAOSlice()
	mockSelectAllSearchCriteria := criteria.MockSelectAll(mockCriteriaDAOSlice, nil)
	mockCategorizedTweetsDAOSlice := categorized.MockCategorizedTweetsDAOSlice()
	mockSelectAllByUserID := categorized.MockSelectAllByUserID(mockCategorizedTweetsDAOSlice, errors.New("failed to execute select all categorized tweets"))

	information := criteria.MakeInformation(mockSelectAllCriteriaExecutionsSummaries, mockSelectAllSearchCriteria, mockSelectAllByUserID)

	want := criteria.FailedToRetrieveCategorizedTweetsByUserID
	_, got := information(context.Background(), 1)

	assert.Equal(t, want, got)
}
