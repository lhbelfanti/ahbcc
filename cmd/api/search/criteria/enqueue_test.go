package criteria_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/lhbelfanti/corpus-creator/cmd/api/search/criteria"
	"github.com/lhbelfanti/corpus-creator/cmd/api/search/criteria/executions"
	"github.com/lhbelfanti/corpus-creator/internal/scrapper"
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
		mockSelectExecutionsByStatuses := executions.MockSelectExecutionsByStatuses(executions.MockExecutionsDAO(), nil)
		mockInsertExecution := executions.MockInsertExecution(1, nil)
		mockEnqueueCriteria := scrapper.MockEnqueueCriteria(nil)

		enqueueCriteria := criteria.MakeEnqueue(mockSelectCriteriaByID, mockSelectExecutionsByStatuses, mockInsertExecution, mockEnqueueCriteria)

		got := enqueueCriteria(context.Background(), 1, tt.forced)

		assert.Nil(t, got)
	}
}

func TestEnqueue_failsWhenSelectCriteriaByIDThrowsError(t *testing.T) {
	mockSelectCriteriaByID := criteria.MockSelectByID(criteria.MockCriteriaDAO(), errors.New("failed to execute select criteria by id"))
	mockSelectExecutionsByStatuses := executions.MockSelectExecutionsByStatuses(executions.MockExecutionsDAO(), nil)
	mockInsertExecution := executions.MockInsertExecution(1, nil)
	mockEnqueueCriteria := scrapper.MockEnqueueCriteria(nil)

	enqueueCriteria := criteria.MakeEnqueue(mockSelectCriteriaByID, mockSelectExecutionsByStatuses, mockInsertExecution, mockEnqueueCriteria)

	want := criteria.FailedToExecuteSelectCriteriaByID
	got := enqueueCriteria(context.Background(), 1, false)

	assert.Equal(t, want, got)
}

func TestEnqueue_failsWhenSelectExecutionsByStatusesThrowsError(t *testing.T) {
	mockSelectCriteriaByID := criteria.MockSelectByID(criteria.MockCriteriaDAO(), nil)
	mockSelectExecutionsByStatuses := executions.MockSelectExecutionsByStatuses(executions.MockExecutionsDAO(), errors.New("failed to execute select executions by statuses"))
	mockInsertExecution := executions.MockInsertExecution(1, nil)
	mockEnqueueCriteria := scrapper.MockEnqueueCriteria(nil)

	enqueueCriteria := criteria.MakeEnqueue(mockSelectCriteriaByID, mockSelectExecutionsByStatuses, mockInsertExecution, mockEnqueueCriteria)

	want := criteria.FailedToExecuteSelectExecutionsByStatuses
	got := enqueueCriteria(context.Background(), 1, false)

	assert.Equal(t, want, got)
}

func TestEnqueue_failsWhenThereIsAlreadyAnExecutionWithTheSameCriteriaIDEnqueued(t *testing.T) {
	mockSelectCriteriaByID := criteria.MockSelectByID(criteria.MockCriteriaDAO(), nil)
	mockSelectExecutionsByStatuses := executions.MockSelectExecutionsByStatuses(executions.MockExecutionsDAO(), nil)
	mockInsertExecution := executions.MockInsertExecution(1, nil)
	mockEnqueueCriteria := scrapper.MockEnqueueCriteria(nil)

	enqueueCriteria := criteria.MakeEnqueue(mockSelectCriteriaByID, mockSelectExecutionsByStatuses, mockInsertExecution, mockEnqueueCriteria)

	want := criteria.AnExecutionOfThisCriteriaIDIsAlreadyEnqueued
	got := enqueueCriteria(context.Background(), 2, false)

	assert.Equal(t, want, got)
}

func TestEnqueue_failsWhenInsertExecutionThrowsError(t *testing.T) {
	mockSelectCriteriaByID := criteria.MockSelectByID(criteria.MockCriteriaDAO(), nil)
	mockSelectExecutionsByStatuses := executions.MockSelectExecutionsByStatuses(executions.MockExecutionsDAO(), nil)
	mockInsertExecution := executions.MockInsertExecution(-1, errors.New("failed to insert execution"))
	mockEnqueueCriteria := scrapper.MockEnqueueCriteria(nil)

	enqueueCriteria := criteria.MakeEnqueue(mockSelectCriteriaByID, mockSelectExecutionsByStatuses, mockInsertExecution, mockEnqueueCriteria)

	want := criteria.FailedToInsertSearchCriteriaExecution
	got := enqueueCriteria(context.Background(), 1, false)

	assert.Equal(t, want, got)
}

func TestEnqueue_failsWhenEnqueueCriteriaThrowsError(t *testing.T) {
	mockSelectCriteriaByID := criteria.MockSelectByID(criteria.MockCriteriaDAO(), nil)
	mockSelectExecutionsByStatuses := executions.MockSelectExecutionsByStatuses(executions.MockExecutionsDAO(), nil)
	mockInsertExecution := executions.MockInsertExecution(1, nil)
	mockEnqueueCriteria := scrapper.MockEnqueueCriteria(errors.New("failed to execute enqueue criteria"))

	enqueueCriteria := criteria.MakeEnqueue(mockSelectCriteriaByID, mockSelectExecutionsByStatuses, mockInsertExecution, mockEnqueueCriteria)

	want := criteria.FailedToExecuteEnqueueCriteria
	got := enqueueCriteria(context.Background(), 1, false)

	assert.Equal(t, want, got)
}

func TestResume_successWhenSelectLastDayExecutedReturnsAnExecutionDay(t *testing.T) {
	mockSelectCriteriaByID := criteria.MockSelectByID(criteria.MockCriteriaDAO(), nil)
	mockDate := time.Date(2024, time.September, 19, 0, 0, 0, 0, time.Local)
	mockExecutionDayDAO := executions.ExecutionDayDAO{ExecutionDate: mockDate, SearchCriteriaExecutionID: 1}
	mockSelectLastDayExecutedByCriteriaID := executions.MockSelectLastDayExecutedByCriteriaID(mockExecutionDayDAO, nil)
	mockSelectExecutionsByStatuses := executions.MockSelectExecutionsByStatuses(executions.MockExecutionsDAO(), nil)
	mockEnqueueCriteria := scrapper.MockEnqueueCriteria(nil)

	resumeCriteria := criteria.MakeResume(mockSelectCriteriaByID, mockSelectLastDayExecutedByCriteriaID, mockSelectExecutionsByStatuses, mockEnqueueCriteria)

	got := resumeCriteria(context.Background(), 2)

	assert.Nil(t, got)
}

func TestResume_successWhenSelectLastDayExecutedDoesntReturnAnExecutionDay(t *testing.T) {
	mockSelectCriteriaByID := criteria.MockSelectByID(criteria.MockCriteriaDAO(), nil)
	mockSelectLastDayExecutedByCriteriaID := executions.MockSelectLastDayExecutedByCriteriaID(executions.ExecutionDayDAO{}, criteria.NoExecutionDaysFoundForTheGivenCriteriaID)
	mockSelectExecutionsByStatuses := executions.MockSelectExecutionsByStatuses(executions.MockExecutionsDAO(), nil)
	mockEnqueueCriteria := scrapper.MockEnqueueCriteria(nil)

	resumeCriteria := criteria.MakeResume(mockSelectCriteriaByID, mockSelectLastDayExecutedByCriteriaID, mockSelectExecutionsByStatuses, mockEnqueueCriteria)

	got := resumeCriteria(context.Background(), 2)

	assert.Nil(t, got)
}

func TestResume_failsWhenSelectCriteriaByIDThrowsError(t *testing.T) {
	mockSelectCriteriaByID := criteria.MockSelectByID(criteria.MockCriteriaDAO(), errors.New("failed to execute select criteria by id"))
	mockDate := time.Date(2024, time.September, 19, 0, 0, 0, 0, time.Local)
	mockExecutionDayDAO := executions.ExecutionDayDAO{ExecutionDate: mockDate, SearchCriteriaExecutionID: 1}
	mockSelectLastDayExecutedByCriteriaID := executions.MockSelectLastDayExecutedByCriteriaID(mockExecutionDayDAO, nil)
	mockSelectExecutionsByStatuses := executions.MockSelectExecutionsByStatuses(executions.MockExecutionsDAO(), nil)
	mockEnqueueCriteria := scrapper.MockEnqueueCriteria(nil)

	resumeCriteria := criteria.MakeResume(mockSelectCriteriaByID, mockSelectLastDayExecutedByCriteriaID, mockSelectExecutionsByStatuses, mockEnqueueCriteria)

	want := criteria.FailedToExecuteSelectCriteriaByID
	got := resumeCriteria(context.Background(), 2)

	assert.Equal(t, want, got)
}

func TestResume_failsWhenSelectLastDayExecutedByCriteriaThrowsError(t *testing.T) {
	mockSelectCriteriaByID := criteria.MockSelectByID(criteria.MockCriteriaDAO(), nil)
	mockSelectLastDayExecutedByCriteriaID := executions.MockSelectLastDayExecutedByCriteriaID(executions.ExecutionDayDAO{}, errors.New("failed to execute select last day executed by criteria id"))
	mockSelectExecutionsByStatuses := executions.MockSelectExecutionsByStatuses(executions.MockExecutionsDAO(), nil)
	mockEnqueueCriteria := scrapper.MockEnqueueCriteria(nil)

	resumeCriteria := criteria.MakeResume(mockSelectCriteriaByID, mockSelectLastDayExecutedByCriteriaID, mockSelectExecutionsByStatuses, mockEnqueueCriteria)

	want := criteria.FailedToExecuteSelectLastDayExecutedByCriteriaID
	got := resumeCriteria(context.Background(), 2)

	assert.Equal(t, want, got)
}

func TestResume_failsWhenSelectExecutionsByStatusesThrowsError(t *testing.T) {
	mockSelectCriteriaByID := criteria.MockSelectByID(criteria.MockCriteriaDAO(), nil)
	mockSelectLastDayExecutedByCriteriaID := executions.MockSelectLastDayExecutedByCriteriaID(executions.ExecutionDayDAO{}, criteria.NoExecutionDaysFoundForTheGivenCriteriaID)
	mockSelectExecutionsByStatuses := executions.MockSelectExecutionsByStatuses(executions.MockExecutionsDAO(), errors.New("failed to execute select executions by statuses"))
	mockEnqueueCriteria := scrapper.MockEnqueueCriteria(nil)

	resumeCriteria := criteria.MakeResume(mockSelectCriteriaByID, mockSelectLastDayExecutedByCriteriaID, mockSelectExecutionsByStatuses, mockEnqueueCriteria)

	want := criteria.FailedToExecuteSelectExecutionsByStatuses
	got := resumeCriteria(context.Background(), 2)

	assert.Equal(t, want, got)
}

func TestResume_failsWhenEnqueueCriteriaThrowsError(t *testing.T) {
	mockSelectCriteriaByID := criteria.MockSelectByID(criteria.MockCriteriaDAO(), nil)
	mockDate := time.Date(2024, time.September, 19, 0, 0, 0, 0, time.Local)
	mockExecutionDayDAO := executions.ExecutionDayDAO{ExecutionDate: mockDate, SearchCriteriaExecutionID: 1}
	mockSelectLastDayExecutedByCriteriaID := executions.MockSelectLastDayExecutedByCriteriaID(mockExecutionDayDAO, nil)
	mockSelectExecutionsByStatuses := executions.MockSelectExecutionsByStatuses(executions.MockExecutionsDAO(), nil)
	mockEnqueueCriteria := scrapper.MockEnqueueCriteria(errors.New("failed to execute enqueue criteria"))

	resumeCriteria := criteria.MakeResume(mockSelectCriteriaByID, mockSelectLastDayExecutedByCriteriaID, mockSelectExecutionsByStatuses, mockEnqueueCriteria)

	want := criteria.FailedToExecuteEnqueueCriteria
	got := resumeCriteria(context.Background(), 2)

	assert.Equal(t, want, got)
}

func TestResume_failsWhenSelectLastDayExecutedDoesntReturnAnExecutionDayAndTheExecutionsInTheDBDoesntBelongToTheCriteria(t *testing.T) {
	mockSelectCriteriaByID := criteria.MockSelectByID(criteria.MockCriteriaDAO(), nil)
	mockSelectLastDayExecutedByCriteriaID := executions.MockSelectLastDayExecutedByCriteriaID(executions.ExecutionDayDAO{}, criteria.NoExecutionDaysFoundForTheGivenCriteriaID)
	mockSelectExecutionsByStatuses := executions.MockSelectExecutionsByStatuses(executions.MockExecutionsDAO(), nil)
	mockEnqueueCriteria := scrapper.MockEnqueueCriteria(nil)

	resumeCriteria := criteria.MakeResume(mockSelectCriteriaByID, mockSelectLastDayExecutedByCriteriaID, mockSelectExecutionsByStatuses, mockEnqueueCriteria)

	want := criteria.FailedToRetrieveSearchCriteriaExecutionID
	got := resumeCriteria(context.Background(), 9999) // some random number for a criteria that is not present in the DB

	assert.Equal(t, got, want)
}
