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
		mockSelectExecutionsByStatuses := criteria.MockSelectExecutionsByStatuses(criteria.MockExecutionsDAO(), nil)

		enqueueCriteria := criteria.MakeEnqueue(mockSelectCriteriaByID, mockSelectExecutionsByStatuses, mockEnqueueCriteria)

		got := enqueueCriteria(context.Background(), 1, tt.forced)

		assert.Nil(t, got)
	}
}

func TestEnqueue_failsWhenSelectCriteriaByIDThrowsError(t *testing.T) {
	mockSelectCriteriaByID := criteria.MockSelectByID(criteria.MockCriteriaDAO(), errors.New("failed to execute select criteria by id"))
	mockSelectExecutionsByStatuses := criteria.MockSelectExecutionsByStatuses(criteria.MockExecutionsDAO(), nil)
	mockEnqueueCriteria := scrapper.MockEnqueueCriteria(nil)

	enqueueCriteria := criteria.MakeEnqueue(mockSelectCriteriaByID, mockSelectExecutionsByStatuses, mockEnqueueCriteria)

	want := criteria.FailedToExecuteSelectCriteriaByID
	got := enqueueCriteria(context.Background(), 1, false)

	assert.Equal(t, want, got)
}

func TestEnqueue_failsWhenSelectExecutionsByStatusesThrowsError(t *testing.T) {
	mockSelectCriteriaByID := criteria.MockSelectByID(criteria.MockCriteriaDAO(), nil)
	mockSelectExecutionsByStatuses := criteria.MockSelectExecutionsByStatuses(criteria.MockExecutionsDAO(), errors.New("failed to execute select executions by statuses"))
	mockEnqueueCriteria := scrapper.MockEnqueueCriteria(nil)

	enqueueCriteria := criteria.MakeEnqueue(mockSelectCriteriaByID, mockSelectExecutionsByStatuses, mockEnqueueCriteria)

	want := criteria.FailedToExecuteSelectExecutionsByStatuses
	got := enqueueCriteria(context.Background(), 1, false)

	assert.Equal(t, want, got)
}

func TestEnqueue_failsWhenThereIsAlreadyAnExecutionWithTheSameCriteriaIDEnqueued(t *testing.T) {
	mockSelectCriteriaByID := criteria.MockSelectByID(criteria.MockCriteriaDAO(), nil)
	mockSelectExecutionsByStatuses := criteria.MockSelectExecutionsByStatuses(criteria.MockExecutionsDAO(), nil)
	mockEnqueueCriteria := scrapper.MockEnqueueCriteria(nil)

	enqueueCriteria := criteria.MakeEnqueue(mockSelectCriteriaByID, mockSelectExecutionsByStatuses, mockEnqueueCriteria)

	want := criteria.AnExecutionOfThisCriteriaIDIsAlreadyEnqueued
	got := enqueueCriteria(context.Background(), 4, false)

	assert.Equal(t, want, got)
}

func TestEnqueue_failsWhenEnqueueCriteriaThrowsError(t *testing.T) {
	mockSelectCriteriaByID := criteria.MockSelectByID(criteria.MockCriteriaDAO(), nil)
	mockSelectExecutionsByStatuses := criteria.MockSelectExecutionsByStatuses(criteria.MockExecutionsDAO(), nil)
	mockEnqueueCriteria := scrapper.MockEnqueueCriteria(errors.New("failed to execute enqueue criteria"))

	enqueueCriteria := criteria.MakeEnqueue(mockSelectCriteriaByID, mockSelectExecutionsByStatuses, mockEnqueueCriteria)

	want := criteria.FailedToExecuteEnqueueCriteria
	got := enqueueCriteria(context.Background(), 1, false)

	assert.Equal(t, want, got)
}

func TestResume_successWhenSelectLastDayExecutedReturnsAnExecutionDay(t *testing.T) {
	mockSelectCriteriaByID := criteria.MockSelectByID(criteria.MockCriteriaDAO(), nil)
	mockDate := time.Date(2024, time.September, 19, 0, 0, 0, 0, time.Local)
	mockExecutionDayDAO := criteria.ExecutionDayDAO{ExecutionDate: mockDate, SearchCriteriaExecutionID: 1}
	mockSelectLastDayExecutedByCriteriaID := criteria.MockSelectLastDayExecutedByCriteriaID(mockExecutionDayDAO, nil)
	mockSelectExecutionsByStatuses := criteria.MockSelectExecutionsByStatuses(criteria.MockExecutionsDAO(), nil)
	mockEnqueueCriteria := scrapper.MockEnqueueCriteria(nil)

	resumeCriteria := criteria.MakeResume(mockSelectCriteriaByID, mockSelectLastDayExecutedByCriteriaID, mockSelectExecutionsByStatuses, mockEnqueueCriteria)

	got := resumeCriteria(context.Background(), 2)

	assert.Nil(t, got)
}

func TestResume_successWhenSelectLastDayExecutedDoesntReturnAnExecutionDay(t *testing.T) {
	mockSelectCriteriaByID := criteria.MockSelectByID(criteria.MockCriteriaDAO(), nil)
	mockSelectLastDayExecutedByCriteriaID := criteria.MockSelectLastDayExecutedByCriteriaID(criteria.ExecutionDayDAO{}, criteria.NoExecutionDaysFoundForTheGivenCriteriaID)
	mockSelectExecutionsByStatuses := criteria.MockSelectExecutionsByStatuses(criteria.MockExecutionsDAO(), nil)
	mockEnqueueCriteria := scrapper.MockEnqueueCriteria(nil)

	resumeCriteria := criteria.MakeResume(mockSelectCriteriaByID, mockSelectLastDayExecutedByCriteriaID, mockSelectExecutionsByStatuses, mockEnqueueCriteria)

	got := resumeCriteria(context.Background(), 2)

	assert.Nil(t, got)
}

func TestResume_failsWhenSelectCriteriaByIDThrowsError(t *testing.T) {
	mockSelectCriteriaByID := criteria.MockSelectByID(criteria.MockCriteriaDAO(), errors.New("failed to execute select criteria by id"))
	mockDate := time.Date(2024, time.September, 19, 0, 0, 0, 0, time.Local)
	mockExecutionDayDAO := criteria.ExecutionDayDAO{ExecutionDate: mockDate, SearchCriteriaExecutionID: 1}
	mockSelectLastDayExecutedByCriteriaID := criteria.MockSelectLastDayExecutedByCriteriaID(mockExecutionDayDAO, nil)
	mockSelectExecutionsByStatuses := criteria.MockSelectExecutionsByStatuses(criteria.MockExecutionsDAO(), nil)
	mockEnqueueCriteria := scrapper.MockEnqueueCriteria(nil)

	resumeCriteria := criteria.MakeResume(mockSelectCriteriaByID, mockSelectLastDayExecutedByCriteriaID, mockSelectExecutionsByStatuses, mockEnqueueCriteria)

	want := criteria.FailedToExecuteSelectCriteriaByID
	got := resumeCriteria(context.Background(), 2)

	assert.Equal(t, want, got)
}

func TestResume_failsWhenSelectLastDayExecutedByCriteriaThrowsError(t *testing.T) {
	mockSelectCriteriaByID := criteria.MockSelectByID(criteria.MockCriteriaDAO(), nil)
	mockSelectLastDayExecutedByCriteriaID := criteria.MockSelectLastDayExecutedByCriteriaID(criteria.ExecutionDayDAO{}, errors.New("failed to execute select last day executed by criteria id"))
	mockSelectExecutionsByStatuses := criteria.MockSelectExecutionsByStatuses(criteria.MockExecutionsDAO(), nil)
	mockEnqueueCriteria := scrapper.MockEnqueueCriteria(nil)

	resumeCriteria := criteria.MakeResume(mockSelectCriteriaByID, mockSelectLastDayExecutedByCriteriaID, mockSelectExecutionsByStatuses, mockEnqueueCriteria)

	want := criteria.FailedToExecuteSelectLastDayExecutedByCriteriaID
	got := resumeCriteria(context.Background(), 2)

	assert.Equal(t, want, got)
}

func TestResume_failsWhenSelectExecutionsByStatusesThrowsError(t *testing.T) {
	mockSelectCriteriaByID := criteria.MockSelectByID(criteria.MockCriteriaDAO(), nil)
	mockSelectLastDayExecutedByCriteriaID := criteria.MockSelectLastDayExecutedByCriteriaID(criteria.ExecutionDayDAO{}, criteria.NoExecutionDaysFoundForTheGivenCriteriaID)
	mockSelectExecutionsByStatuses := criteria.MockSelectExecutionsByStatuses(criteria.MockExecutionsDAO(), errors.New("failed to execute select executions by statuses"))
	mockEnqueueCriteria := scrapper.MockEnqueueCriteria(nil)

	resumeCriteria := criteria.MakeResume(mockSelectCriteriaByID, mockSelectLastDayExecutedByCriteriaID, mockSelectExecutionsByStatuses, mockEnqueueCriteria)

	want := criteria.FailedToExecuteSelectExecutionsByStatuses
	got := resumeCriteria(context.Background(), 2)

	assert.Equal(t, want, got)
}

func TestResume_failsWhenEnqueueCriteriaThrowsError(t *testing.T) {
	mockSelectCriteriaByID := criteria.MockSelectByID(criteria.MockCriteriaDAO(), nil)
	mockDate := time.Date(2024, time.September, 19, 0, 0, 0, 0, time.Local)
	mockExecutionDayDAO := criteria.ExecutionDayDAO{ExecutionDate: mockDate, SearchCriteriaExecutionID: 1}
	mockSelectLastDayExecutedByCriteriaID := criteria.MockSelectLastDayExecutedByCriteriaID(mockExecutionDayDAO, nil)
	mockSelectExecutionsByStatuses := criteria.MockSelectExecutionsByStatuses(criteria.MockExecutionsDAO(), nil)
	mockEnqueueCriteria := scrapper.MockEnqueueCriteria(errors.New("failed to execute enqueue criteria"))

	resumeCriteria := criteria.MakeResume(mockSelectCriteriaByID, mockSelectLastDayExecutedByCriteriaID, mockSelectExecutionsByStatuses, mockEnqueueCriteria)

	want := criteria.FailedToExecuteEnqueueCriteria
	got := resumeCriteria(context.Background(), 2)

	assert.Equal(t, want, got)
}

func TestResume_failsWhenSelectLastDayExecutedDoesntReturnAnExecutionDayAndTheExecutionsInTheDBDoesntBelongToTheCriteria(t *testing.T) {
	mockSelectCriteriaByID := criteria.MockSelectByID(criteria.MockCriteriaDAO(), nil)
	mockSelectLastDayExecutedByCriteriaID := criteria.MockSelectLastDayExecutedByCriteriaID(criteria.ExecutionDayDAO{}, criteria.NoExecutionDaysFoundForTheGivenCriteriaID)
	mockSelectExecutionsByStatuses := criteria.MockSelectExecutionsByStatuses(criteria.MockExecutionsDAO(), nil)
	mockEnqueueCriteria := scrapper.MockEnqueueCriteria(nil)

	resumeCriteria := criteria.MakeResume(mockSelectCriteriaByID, mockSelectLastDayExecutedByCriteriaID, mockSelectExecutionsByStatuses, mockEnqueueCriteria)

	want := criteria.FailedToRetrieveSearchCriteriaExecutionID
	got := resumeCriteria(context.Background(), 9999) // some random number for a criteria that is not present in the DB

	assert.Equal(t, got, want)
}
