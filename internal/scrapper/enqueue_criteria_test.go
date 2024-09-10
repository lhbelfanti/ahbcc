package scrapper_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"ahbcc/cmd/api/search/criteria"
	"ahbcc/internal/http"
	"ahbcc/internal/scrapper"
)

func TestEnqueueCriteria_success(t *testing.T) {
	mockHTTPClient := new(http.MockHTTPClient)
	resp := http.Response{
		Status: "200 OK",
		Body:   `{"test": "body"}`,
	}
	mockHTTPClient.On("NewRequest", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(resp, nil)
	mockSelectCriteriaByID := criteria.MockSelectByID(criteria.MockCriteriaDAO(), nil)
	mockCtx := context.Background()
	enqueueCriteria := scrapper.MakeEnqueueCriteria(mockHTTPClient, "http://example.com", mockSelectCriteriaByID)

	got := enqueueCriteria(mockCtx, 1)

	assert.Nil(t, got)
	mockHTTPClient.AssertExpectations(t)
}

func TestEnqueueCriteria_failsWhenSelectCriteriaByIDThrowsError(t *testing.T) {
	mockHTTPClient := new(http.MockHTTPClient)
	resp := http.Response{
		Status: "200 OK",
		Body:   `{"test": "body"}`,
	}
	mockHTTPClient.On("NewRequest", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(resp, nil)
	mockSelectCriteriaByID := criteria.MockSelectByID(criteria.DAO{}, errors.New("select criteria by id failed"))
	mockCtx := context.Background()
	enqueueCriteria := scrapper.MakeEnqueueCriteria(mockHTTPClient, "http://example.com", mockSelectCriteriaByID)

	want := scrapper.FailedToExecuteSelectCriteriaByID
	got := enqueueCriteria(mockCtx, 1)

	assert.Equal(t, want, got)
}

func TestEnqueueCriteria_failsWhenNewRequestThrowsError(t *testing.T) {
	mockHTTPClient := new(http.MockHTTPClient)
	mockHTTPClient.On("NewRequest", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(http.Response{}, errors.New("failed to execute NewRequest"))
	mockSelectCriteriaByID := criteria.MockSelectByID(criteria.MockCriteriaDAO(), nil)
	mockCtx := context.Background()
	enqueueCriteria := scrapper.MakeEnqueueCriteria(mockHTTPClient, "http://example.com", mockSelectCriteriaByID)

	want := scrapper.FailedToExecuteRequest
	got := enqueueCriteria(mockCtx, 1)

	assert.Equal(t, want, got)
	mockHTTPClient.AssertExpectations(t)
}
