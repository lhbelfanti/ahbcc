package scrapper_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/lhbelfanti/corpus-creator/internal/http"
	"github.com/lhbelfanti/corpus-creator/internal/scrapper"
)

func TestEnqueueCriteria_success(t *testing.T) {
	mockHTTPClient := new(http.MockHTTPClient)
	resp := http.Response{
		Status: "200 OK",
		Body:   `{"test": "body"}`,
	}
	mockHTTPClient.On("NewRequest", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(resp, nil)
	enqueueCriteria := scrapper.MakeEnqueueCriteria(mockHTTPClient, "http://example.com")

	got := enqueueCriteria(context.Background(), scrapper.MockCriteriaDTO(), 1)

	assert.Nil(t, got)
	mockHTTPClient.AssertExpectations(t)
}

func TestEnqueueCriteria_failsWhenNewRequestThrowsError(t *testing.T) {
	mockHTTPClient := new(http.MockHTTPClient)
	mockHTTPClient.On("NewRequest", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(http.Response{}, errors.New("failed to execute NewRequest"))
	enqueueCriteria := scrapper.MakeEnqueueCriteria(mockHTTPClient, "http://example.com")

	want := scrapper.FailedToExecuteRequest
	got := enqueueCriteria(context.Background(), scrapper.MockCriteriaDTO(), 1)

	assert.Equal(t, want, got)
	mockHTTPClient.AssertExpectations(t)
}
