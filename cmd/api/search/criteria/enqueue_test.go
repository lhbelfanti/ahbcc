package criteria_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"ahbcc/cmd/api/search/criteria"
	"ahbcc/internal/scrapper"
)

func TestEnqueue_success(t *testing.T) {
	tests := []struct {
		forced bool
	}{
		{true},
		{false},
	}

	for _, tt := range tests {
		mockSelectCriteriaByID := criteria.MockSelectByID(criteria.MockCriteriaDAO(), nil)
		mockEnqueueCriteria := scrapper.MockEnqueueCriteria(nil)
		mockDate := time.Date(2024, time.September, 19, 0, 0, 0, 0, time.Local)
		mockSelectLastDayExecutedByCriteriaID := criteria.MockSelectLastDayExecutedByCriteriaID(mockDate, nil)
		mockSelectExecutionsByStatuses := criteria.MockSelectExecutionsByStatuses(criteria.MockExecutionsDAO(), nil)
		enqueueCriteria := criteria.MakeEnqueue(mockSelectCriteriaByID, mockSelectLastDayExecutedByCriteriaID, mockSelectExecutionsByStatuses, mockEnqueueCriteria)

		got := enqueueCriteria(context.Background(), 1, tt.forced)

		assert.Nil(t, got)
	}
}

func TestEnqueue_failsWhenSelectLastDayExecutedByCriteriaThrowsError(t *testing.T) {
	mockSelectCriteriaByID := criteria.MockSelectByID(criteria.MockCriteriaDAO(), nil)
	mockEnqueueCriteria := scrapper.MockEnqueueCriteria(nil)
	mockSelectLastDayExecutedByCriteriaID := criteria.MockSelectLastDayExecutedByCriteriaID(time.Time{}, errors.New("failed to execute select last day executed by criteria id"))
	mockSelectExecutionsByStatuses := criteria.MockSelectExecutionsByStatuses(criteria.MockExecutionsDAO(), nil)
	enqueueCriteria := criteria.MakeEnqueue(mockSelectCriteriaByID, mockSelectLastDayExecutedByCriteriaID, mockSelectExecutionsByStatuses, mockEnqueueCriteria)

	want := criteria.FailedToExecuteSelectLastDayExecutedByCriteriaID
	got := enqueueCriteria(context.Background(), 1, true)

	assert.Equal(t, want, got)
}

func TestEnqueue_failsWhenSelectCriteriaByIDThrowsError(t *testing.T) {
	mockSelectCriteriaByID := criteria.MockSelectByID(criteria.MockCriteriaDAO(), errors.New("failed to execute select criteria by id"))
	mockEnqueueCriteria := scrapper.MockEnqueueCriteria(nil)
	mockDate := time.Date(2024, time.September, 19, 0, 0, 0, 0, time.Local)
	mockSelectLastDayExecutedByCriteriaID := criteria.MockSelectLastDayExecutedByCriteriaID(mockDate, nil)
	mockSelectExecutionsByStatuses := criteria.MockSelectExecutionsByStatuses(criteria.MockExecutionsDAO(), nil)
	enqueueCriteria := criteria.MakeEnqueue(mockSelectCriteriaByID, mockSelectLastDayExecutedByCriteriaID, mockSelectExecutionsByStatuses, mockEnqueueCriteria)

	want := criteria.FailedToExecuteSelectCriteriaByID
	got := enqueueCriteria(context.Background(), 1, false)

	assert.Equal(t, want, got)
}

func TestEnqueue_failsWhenSelectExecutionsByStatusesThrowsError(t *testing.T) {
	mockSelectCriteriaByID := criteria.MockSelectByID(criteria.MockCriteriaDAO(), nil)
	mockEnqueueCriteria := scrapper.MockEnqueueCriteria(nil)
	mockDate := time.Date(2024, time.September, 19, 0, 0, 0, 0, time.Local)
	mockSelectLastDayExecutedByCriteriaID := criteria.MockSelectLastDayExecutedByCriteriaID(mockDate, nil)
	mockSelectExecutionsByStatuses := criteria.MockSelectExecutionsByStatuses(criteria.MockExecutionsDAO(), errors.New("failed to execute select executions by statuses"))
	enqueueCriteria := criteria.MakeEnqueue(mockSelectCriteriaByID, mockSelectLastDayExecutedByCriteriaID, mockSelectExecutionsByStatuses, mockEnqueueCriteria)

	want := criteria.FailedToExecuteSelectExecutionsByStatuses
	got := enqueueCriteria(context.Background(), 1, false)

	assert.Equal(t, want, got)
}

func TestEnqueue_failsWhenThereIsAlreadyAnExecutionWithTheSameCriteriaIDEnqueued(t *testing.T) {
	mockSelectCriteriaByID := criteria.MockSelectByID(criteria.MockCriteriaDAO(), nil)
	mockEnqueueCriteria := scrapper.MockEnqueueCriteria(nil)
	mockDate := time.Date(2024, time.September, 19, 0, 0, 0, 0, time.Local)
	mockSelectLastDayExecutedByCriteriaID := criteria.MockSelectLastDayExecutedByCriteriaID(mockDate, nil)
	mockSelectExecutionsByStatuses := criteria.MockSelectExecutionsByStatuses(criteria.MockExecutionsDAO(), nil)
	enqueueCriteria := criteria.MakeEnqueue(mockSelectCriteriaByID, mockSelectLastDayExecutedByCriteriaID, mockSelectExecutionsByStatuses, mockEnqueueCriteria)

	want := criteria.AnExecutionOfThisCriteriaIDIsAlreadyEnqueued
	got := enqueueCriteria(context.Background(), 4, false)

	assert.Equal(t, want, got)
}

func TestEnqueue_failsWhenEnqueueCriteriaThrowsError(t *testing.T) {
	mockSelectCriteriaByID := criteria.MockSelectByID(criteria.MockCriteriaDAO(), nil)
	mockEnqueueCriteria := scrapper.MockEnqueueCriteria(errors.New("failed to execute enqueue criteria"))
	mockDate := time.Date(2024, time.September, 19, 0, 0, 0, 0, time.Local)
	mockSelectLastDayExecutedByCriteriaID := criteria.MockSelectLastDayExecutedByCriteriaID(mockDate, nil)
	mockSelectExecutionsByStatuses := criteria.MockSelectExecutionsByStatuses(criteria.MockExecutionsDAO(), nil)
	enqueueCriteria := criteria.MakeEnqueue(mockSelectCriteriaByID, mockSelectLastDayExecutedByCriteriaID, mockSelectExecutionsByStatuses, mockEnqueueCriteria)

	want := criteria.FailedToExecuteEnqueueCriteria
	got := enqueueCriteria(context.Background(), 1, false)

	assert.Equal(t, want, got)
}
