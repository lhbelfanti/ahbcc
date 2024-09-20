package criteria_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"ahbcc/cmd/api/search/criteria"
)

func TestInit_success(t *testing.T) {
	mockExecutionsDAO := criteria.MockExecutionsDAO()
	mockSelectExecutionsByStatuses := criteria.MockSelectExecutionsByStatuses(mockExecutionsDAO, nil)
	mockEnqueue := criteria.MockEnqueue(nil)

	init := criteria.MakeInit(mockSelectExecutionsByStatuses, mockEnqueue)

	got := init(context.Background())

	assert.Nil(t, got)
}

func TestInit_failsWhenSelectExecutionsByStatusesThrowsError(t *testing.T) {
	mockSelectExecutionsByStatuses := criteria.MockSelectExecutionsByStatuses(nil, errors.New("failed while executing select executions by statuses"))
	mockEnqueue := criteria.MockEnqueue(nil)

	init := criteria.MakeInit(mockSelectExecutionsByStatuses, mockEnqueue)

	want := criteria.FailedToExecuteSelectExecutionsByStatuses
	got := init(context.Background())

	assert.Equal(t, want, got)
}

func TestInit_failsWhenEnqueueThrowsError(t *testing.T) {
	mockExecutionsDAO := criteria.MockExecutionsDAO()
	mockSelectExecutionsByStatuses := criteria.MockSelectExecutionsByStatuses(mockExecutionsDAO, nil)
	mockEnqueue := criteria.MockEnqueue(errors.New("failed while executing enqueue"))

	init := criteria.MakeInit(mockSelectExecutionsByStatuses, mockEnqueue)

	want := criteria.FailedToExecuteEnqueueCriteria
	got := init(context.Background())

	assert.Equal(t, want, got)
}
