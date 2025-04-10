package summary_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"ahbcc/cmd/api/search/criteria/executions/summary"
	"ahbcc/internal/database"
)

func TestUpsert_successCallingToInsert(t *testing.T) {
	mockSelectIDBySearchCriteriaIDYearAndMonth := summary.MockSelectIDBySearchCriteriaIDYearAndMonth(1, summary.NoExecutionSummaryFoundForTheGivenCriteria)
	mockInsert := summary.MockInsert(1, nil)
	mockUpdateTotalTweets := summary.MockUpdateTotalTweets(nil)
	mockPostgresTx := new(database.MockPgxTx)
	mockExecutionSummaryDAO := summary.MockExecutionSummaryDAO(1, 2025, 04, 11)

	upsertSummary := summary.MakeUpsert(mockSelectIDBySearchCriteriaIDYearAndMonth, mockInsert, mockUpdateTotalTweets)

	got := upsertSummary(context.Background(), mockPostgresTx, mockExecutionSummaryDAO)

	assert.Nil(t, got)
}

func TestUpsert_successCallingToUpdate(t *testing.T) {
	mockSelectIDBySearchCriteriaIDYearAndMonth := summary.MockSelectIDBySearchCriteriaIDYearAndMonth(1, nil)
	mockInsert := summary.MockInsert(1, nil)
	mockUpdateTotalTweets := summary.MockUpdateTotalTweets(nil)
	mockPostgresTx := new(database.MockPgxTx)
	mockExecutionSummaryDAO := summary.MockExecutionSummaryDAO(1, 2025, 04, 11)

	upsertSummary := summary.MakeUpsert(mockSelectIDBySearchCriteriaIDYearAndMonth, mockInsert, mockUpdateTotalTweets)

	got := upsertSummary(context.Background(), mockPostgresTx, mockExecutionSummaryDAO)

	assert.Nil(t, got)
}

func TestUpsert_failsWhenSelectIDBySearchCriteriaIDYearAndMonthThrowsError(t *testing.T) {
	mockSelectIDBySearchCriteriaIDYearAndMonth := summary.MockSelectIDBySearchCriteriaIDYearAndMonth(1, errors.New("failed to execute select ID by search criteria, year and month"))
	mockInsert := summary.MockInsert(1, nil)
	mockUpdateTotalTweets := summary.MockUpdateTotalTweets(nil)
	mockPostgresTx := new(database.MockPgxTx)
	mockExecutionSummaryDAO := summary.MockExecutionSummaryDAO(1, 2025, 04, 11)

	upsertSummary := summary.MakeUpsert(mockSelectIDBySearchCriteriaIDYearAndMonth, mockInsert, mockUpdateTotalTweets)

	want := summary.FailedToRetrieveExecutionSummaryID
	got := upsertSummary(context.Background(), mockPostgresTx, mockExecutionSummaryDAO)

	assert.Equal(t, want, got)
}

func TestUpsert_failsWhenInsertThrowsError(t *testing.T) {
	mockSelectIDBySearchCriteriaIDYearAndMonth := summary.MockSelectIDBySearchCriteriaIDYearAndMonth(1, summary.NoExecutionSummaryFoundForTheGivenCriteria)
	mockInsert := summary.MockInsert(1, errors.New("failed to execute insert"))
	mockUpdateTotalTweets := summary.MockUpdateTotalTweets(nil)
	mockPostgresTx := new(database.MockPgxTx)
	mockExecutionSummaryDAO := summary.MockExecutionSummaryDAO(1, 2025, 04, 11)

	upsertSummary := summary.MakeUpsert(mockSelectIDBySearchCriteriaIDYearAndMonth, mockInsert, mockUpdateTotalTweets)

	want := summary.FailedToInsertExecutionSummary
	got := upsertSummary(context.Background(), mockPostgresTx, mockExecutionSummaryDAO)

	assert.Equal(t, want, got)
}

func TestUpsert_failsWhenUpdateTotalTweetsThrowsError(t *testing.T) {
	mockSelectIDBySearchCriteriaIDYearAndMonth := summary.MockSelectIDBySearchCriteriaIDYearAndMonth(1, nil)
	mockInsert := summary.MockInsert(1, nil)
	mockUpdateTotalTweets := summary.MockUpdateTotalTweets(errors.New("failed to execute update total tweets"))
	mockPostgresTx := new(database.MockPgxTx)
	mockExecutionSummaryDAO := summary.MockExecutionSummaryDAO(1, 2025, 04, 11)

	upsertSummary := summary.MakeUpsert(mockSelectIDBySearchCriteriaIDYearAndMonth, mockInsert, mockUpdateTotalTweets)

	want := summary.FailedToUpdateTotalTweets
	got := upsertSummary(context.Background(), mockPostgresTx, mockExecutionSummaryDAO)

	assert.Equal(t, want, got)
}
