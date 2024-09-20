package criteria_test

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"ahbcc/cmd/api/search/criteria"
)

func TestEnqueueHandlerV1_success(t *testing.T) {
	mockEnqueueCriteria := criteria.MockEnqueue(nil)

	mockResponseWriter := httptest.NewRecorder()
	mockRequest, _ := http.NewRequestWithContext(context.Background(), http.MethodPost, "/criteria/{criteria_id}/enqueue/v1?forced=false", http.NoBody)
	mockRequest.SetPathValue("criteria_id", "1")

	handlerV1 := criteria.EnqueueHandlerV1(mockEnqueueCriteria)

	handlerV1(mockResponseWriter, mockRequest)

	want := http.StatusOK
	got := mockResponseWriter.Result().StatusCode

	assert.Equal(t, want, got)
}

func TestEnqueueHandlerV1_failedWhenTheURLParamIsEmpty(t *testing.T) {
	mockEnqueueCriteria := criteria.MockEnqueue(nil)

	mockResponseWriter := httptest.NewRecorder()
	mockRequest, _ := http.NewRequestWithContext(context.Background(), http.MethodPost, "/criteria/{criteria_id}/enqueue/v1?forced=false", http.NoBody)

	handlerV1 := criteria.EnqueueHandlerV1(mockEnqueueCriteria)

	handlerV1(mockResponseWriter, mockRequest)

	want := http.StatusBadRequest
	got := mockResponseWriter.Result().StatusCode

	assert.Equal(t, want, got)
}

func TestEnqueueHandlerV1_failedWhenTheURLParamCannotBeParsed(t *testing.T) {
	mockEnqueueCriteria := criteria.MockEnqueue(nil)

	mockResponseWriter := httptest.NewRecorder()
	mockRequest, _ := http.NewRequestWithContext(context.Background(), http.MethodPost, "/criteria/{criteria_id}/enqueue/v1?forced=false", http.NoBody)
	mockRequest.SetPathValue("criteria_id", "error")

	handlerV1 := criteria.EnqueueHandlerV1(mockEnqueueCriteria)

	handlerV1(mockResponseWriter, mockRequest)

	want := http.StatusBadRequest
	got := mockResponseWriter.Result().StatusCode

	assert.Equal(t, want, got)
}

func TestEnqueueHandlerV1_failedWhenTheQueryParamCantBeParsed(t *testing.T) {
	mockEnqueueCriteria := criteria.MockEnqueue(nil)

	mockResponseWriter := httptest.NewRecorder()
	mockRequest, _ := http.NewRequestWithContext(context.Background(), http.MethodPost, "/criteria/{criteria_id}/enqueue/v1", http.NoBody)
	mockRequest.SetPathValue("criteria_id", "1")

	handlerV1 := criteria.EnqueueHandlerV1(mockEnqueueCriteria)

	handlerV1(mockResponseWriter, mockRequest)

	want := http.StatusBadRequest
	got := mockResponseWriter.Result().StatusCode

	assert.Equal(t, want, got)
}

func TestEnqueueHandlerV1_failsWhenEnqueueCriteriaThrowsError(t *testing.T) {
	tests := []struct {
		enqueueError error
		expectedCode int
	}{
		{enqueueError: errors.New("failed while executing enqueue criteria"), expectedCode: http.StatusInternalServerError},
		{enqueueError: criteria.AnExecutionOfThisCriteriaIDIsAlreadyEnqueued, expectedCode: http.StatusConflict},
	}

	for _, tt := range tests {
		mockEnqueueCriteria := criteria.MockEnqueue(tt.enqueueError)

		mockResponseWriter := httptest.NewRecorder()
		mockRequest, _ := http.NewRequestWithContext(context.Background(), http.MethodPost, "/criteria/{criteria_id}/enqueue/v1?forced=false", http.NoBody)
		mockRequest.SetPathValue("criteria_id", "1")

		handlerV1 := criteria.EnqueueHandlerV1(mockEnqueueCriteria)

		handlerV1(mockResponseWriter, mockRequest)

		want := tt.expectedCode
		got := mockResponseWriter.Result().StatusCode

		assert.Equal(t, want, got)
	}
}

func TestInitHandlerV1_success(t *testing.T) {
	mockInit := criteria.MockInit(nil)

	mockResponseWriter := httptest.NewRecorder()
	mockRequest, _ := http.NewRequestWithContext(context.Background(), http.MethodPost, "/criteria/init", http.NoBody)

	handlerV1 := criteria.InitHandlerV1(mockInit)

	handlerV1(mockResponseWriter, mockRequest)

	want := http.StatusOK
	got := mockResponseWriter.Result().StatusCode

	assert.Equal(t, want, got)
}

func TestInitHandlerV1_failsWhenInitThrowsError(t *testing.T) {
	mockInit := criteria.MockInit(errors.New("failed while executing init criteria"))

	mockResponseWriter := httptest.NewRecorder()
	mockRequest, _ := http.NewRequestWithContext(context.Background(), http.MethodPost, "/criteria/init", http.NoBody)

	handlerV1 := criteria.InitHandlerV1(mockInit)

	handlerV1(mockResponseWriter, mockRequest)

	want := http.StatusInternalServerError
	got := mockResponseWriter.Result().StatusCode

	assert.Equal(t, want, got)
}
