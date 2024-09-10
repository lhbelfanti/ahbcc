package criteria_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"ahbcc/cmd/api/search/criteria"
	"ahbcc/internal/scrapper"
)

func TestEnqueue_success(t *testing.T) {
	mockSelectCriteriaByID := criteria.MockSelectByID(criteria.MockCriteriaDAO(), nil)
	mockEnqueueCriteria := scrapper.MockEnqueueCriteria(nil)
	enqueueCriteria := criteria.MakeEnqueue(mockSelectCriteriaByID, mockEnqueueCriteria)

	got := enqueueCriteria(context.Background(), 1)

	assert.Nil(t, got)
}

func TestEnqueue_failsWhenSelectCriteriaByIDThrowsError(t *testing.T) {
	mockSelectCriteriaByID := criteria.MockSelectByID(criteria.MockCriteriaDAO(), errors.New("failed to execute select criteria by id"))
	mockEnqueueCriteria := scrapper.MockEnqueueCriteria(nil)
	enqueueCriteria := criteria.MakeEnqueue(mockSelectCriteriaByID, mockEnqueueCriteria)

	want := criteria.FailedToExecuteSelectCriteriaByID
	got := enqueueCriteria(context.Background(), 1)

	assert.Equal(t, want, got)
}

func TestEnqueue_failsWhenEnqueueCriteriaThrowsError(t *testing.T) {
	mockSelectCriteriaByID := criteria.MockSelectByID(criteria.MockCriteriaDAO(), nil)
	mockEnqueueCriteria := scrapper.MockEnqueueCriteria(errors.New("failed to execute enqueue criteria"))
	enqueueCriteria := criteria.MakeEnqueue(mockSelectCriteriaByID, mockEnqueueCriteria)

	want := criteria.FailedToExecuteEnqueueCriteria
	got := enqueueCriteria(context.Background(), 1)

	assert.Equal(t, want, got)
}
