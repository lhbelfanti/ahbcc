package criteria_test

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
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

func TestEnqueueHandlerV1_failsWhenTheURLParamIsEmpty(t *testing.T) {
	mockEnqueueCriteria := criteria.MockEnqueue(nil)
	mockResponseWriter := httptest.NewRecorder()
	mockRequest, _ := http.NewRequestWithContext(context.Background(), http.MethodPost, "/criteria/{criteria_id}/enqueue/v1?forced=false", http.NoBody)

	handlerV1 := criteria.EnqueueHandlerV1(mockEnqueueCriteria)

	handlerV1(mockResponseWriter, mockRequest)

	want := http.StatusBadRequest
	got := mockResponseWriter.Result().StatusCode

	assert.Equal(t, want, got)
}

func TestEnqueueHandlerV1_failsWhenTheURLParamCannotBeParsed(t *testing.T) {
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

func TestEnqueueHandlerV1_failsWhenTheQueryParamCantBeParsed(t *testing.T) {
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

func TestGetExecutionByIDHandlerV1_success(t *testing.T) {
	mockExecutionDAO := criteria.MockExecutionDAO()
	mockSelectExecutionByID := criteria.MockSelectExecutionByID(mockExecutionDAO, nil)
	mockResponseWriter := httptest.NewRecorder()
	mockRequest, _ := http.NewRequestWithContext(context.Background(), http.MethodGet, "/criteria/executions/{execution_id}/v1", http.NoBody)
	mockRequest.SetPathValue("execution_id", "1")

	handlerV1 := criteria.GetExecutionByIDHandlerV1(mockSelectExecutionByID)

	handlerV1(mockResponseWriter, mockRequest)

	body, err := io.ReadAll(mockResponseWriter.Result().Body)
	if err != nil {
		t.Fatalf("Failed to read response body: %v", err)
	}

	want := mockExecutionDAO
	var got criteria.ExecutionDAO
	err = json.Unmarshal(body, &got)
	if err != nil {
		t.Fatalf("Failed to parse response body as JSON: %v", err)
	}

	assert.Equal(t, want, got)
	assert.Equal(t, "application/json", mockResponseWriter.Header().Get("Content-Type"))
	assert.Equal(t, http.StatusOK, mockResponseWriter.Result().StatusCode)
}

func TestGetExecutionByIDHandlerV1_failsWhenTheURLParamIsEmpty(t *testing.T) {
	mockExecutionDAO := criteria.MockExecutionDAO()
	mockSelectExecutionByID := criteria.MockSelectExecutionByID(mockExecutionDAO, nil)
	mockResponseWriter := httptest.NewRecorder()
	mockRequest, _ := http.NewRequestWithContext(context.Background(), http.MethodGet, "/criteria/executions/{execution_id}/v1", http.NoBody)

	handlerV1 := criteria.GetExecutionByIDHandlerV1(mockSelectExecutionByID)

	handlerV1(mockResponseWriter, mockRequest)

	want := http.StatusBadRequest
	got := mockResponseWriter.Result().StatusCode

	assert.Equal(t, want, got)
}

func TestGetExecutionByIDHandlerV1_failsWhenSelectExecutionByIDThrowsError(t *testing.T) {
	mockSelectExecutionByID := criteria.MockSelectExecutionByID(criteria.ExecutionDAO{}, errors.New("failed to select execution by id"))
	mockResponseWriter := httptest.NewRecorder()
	mockRequest, _ := http.NewRequestWithContext(context.Background(), http.MethodGet, "/criteria/executions/{execution_id}/v1", http.NoBody)
	mockRequest.SetPathValue("execution_id", "1")

	handlerV1 := criteria.GetExecutionByIDHandlerV1(mockSelectExecutionByID)

	handlerV1(mockResponseWriter, mockRequest)

	want := http.StatusInternalServerError
	got := mockResponseWriter.Result().StatusCode

	assert.Equal(t, want, got)
}

func TestUpdateExecutionHandlerV1_success(t *testing.T) {
	mockUpdateExecution := criteria.MockUpdateExecution(nil)
	mockResponseWriter := httptest.NewRecorder()
	mockExecution := criteria.MockExecutionDTO()
	mockBody, _ := json.Marshal(mockExecution)
	mockRequest, _ := http.NewRequestWithContext(context.Background(), http.MethodPut, "/criteria/executions/{execution_id}/v1", bytes.NewReader(mockBody))
	mockRequest.SetPathValue("execution_id", "1")

	handlerV1 := criteria.UpdateExecutionHandlerV1(mockUpdateExecution)

	handlerV1(mockResponseWriter, mockRequest)

	want := http.StatusOK
	got := mockResponseWriter.Result().StatusCode

	assert.Equal(t, want, got)
}

func TestUpdateExecutionHandlerV1_failsWhenTheURLParamIsEmpty(t *testing.T) {
	mockUpdateExecution := criteria.MockUpdateExecution(nil)
	mockResponseWriter := httptest.NewRecorder()
	mockExecution := criteria.MockExecutionDTO()
	mockBody, _ := json.Marshal(mockExecution)
	mockRequest, _ := http.NewRequestWithContext(context.Background(), http.MethodPut, "/criteria/executions/{execution_id}/v1", bytes.NewReader(mockBody))

	handlerV1 := criteria.UpdateExecutionHandlerV1(mockUpdateExecution)

	handlerV1(mockResponseWriter, mockRequest)

	want := http.StatusBadRequest
	got := mockResponseWriter.Result().StatusCode

	assert.Equal(t, want, got)
}

func TestUpdateExecutionHandlerV1_failsWhenTheBodyCantBeParsed(t *testing.T) {
	mockUpdateExecution := criteria.MockUpdateExecution(nil)
	mockResponseWriter := httptest.NewRecorder()
	mockBody, _ := json.Marshal(`{"wrong": "body"}`)
	mockRequest, _ := http.NewRequestWithContext(context.Background(), http.MethodPut, "/criteria/executions/{execution_id}/v1", bytes.NewReader(mockBody))
	mockRequest.SetPathValue("execution_id", "1")

	handlerV1 := criteria.UpdateExecutionHandlerV1(mockUpdateExecution)

	handlerV1(mockResponseWriter, mockRequest)

	want := http.StatusBadRequest
	got := mockResponseWriter.Result().StatusCode

	assert.Equal(t, want, got)
}

func TestUpdateExecutionHandlerV1_failsWhenUpdateExecutionThrowsError(t *testing.T) {
	mockUpdateExecution := criteria.MockUpdateExecution(errors.New("failed to update execution"))
	mockResponseWriter := httptest.NewRecorder()
	mockExecution := criteria.MockExecutionDTO()
	mockBody, _ := json.Marshal(mockExecution)
	mockRequest, _ := http.NewRequestWithContext(context.Background(), http.MethodPut, "/criteria/executions/{execution_id}/v1", bytes.NewReader(mockBody))
	mockRequest.SetPathValue("execution_id", "1")

	handlerV1 := criteria.UpdateExecutionHandlerV1(mockUpdateExecution)

	handlerV1(mockResponseWriter, mockRequest)

	want := http.StatusInternalServerError
	got := mockResponseWriter.Result().StatusCode

	assert.Equal(t, want, got)
}

func TestCreateExecutionDayHandlerV1_success(t *testing.T) {
	mockInsertExecutionDay := criteria.MockInsertExecutionDay(nil)
	mockResponseWriter := httptest.NewRecorder()
	mockExecutionDay := criteria.MockExecutionDayDTO(nil)
	mockBody, _ := json.Marshal(mockExecutionDay)
	mockRequest, _ := http.NewRequestWithContext(context.Background(), http.MethodPost, "/criteria/executions/{execution_id}/day/v1", bytes.NewReader(mockBody))
	mockRequest.SetPathValue("execution_id", "1")

	handlerV1 := criteria.CreateExecutionDayHandlerV1(mockInsertExecutionDay)

	handlerV1(mockResponseWriter, mockRequest)

	want := http.StatusOK
	got := mockResponseWriter.Result().StatusCode

	assert.Equal(t, want, got)
}

func TestCreateExecutionDayHandlerV1_failsWhenTheURLParamIsEmpty(t *testing.T) {
	mockInsertExecutionDay := criteria.MockInsertExecutionDay(nil)
	mockResponseWriter := httptest.NewRecorder()
	mockExecutionDay := criteria.MockExecutionDayDTO(nil)
	mockBody, _ := json.Marshal(mockExecutionDay)
	mockRequest, _ := http.NewRequestWithContext(context.Background(), http.MethodPost, "/criteria/executions/{execution_id}/day/v1", bytes.NewReader(mockBody))

	handlerV1 := criteria.CreateExecutionDayHandlerV1(mockInsertExecutionDay)

	handlerV1(mockResponseWriter, mockRequest)

	want := http.StatusBadRequest
	got := mockResponseWriter.Result().StatusCode

	assert.Equal(t, want, got)
}

func TestCreateExecutionDayHandlerV1_failsWhenTheBodyCantBeParsed(t *testing.T) {
	mockInsertExecutionDay := criteria.MockInsertExecutionDay(nil)
	mockResponseWriter := httptest.NewRecorder()
	mockBody, _ := json.Marshal(`{"wrong": "body"}`)
	mockRequest, _ := http.NewRequestWithContext(context.Background(), http.MethodPost, "/criteria/executions/{execution_id}/day/v1", bytes.NewReader(mockBody))
	mockRequest.SetPathValue("execution_id", "1")

	handlerV1 := criteria.CreateExecutionDayHandlerV1(mockInsertExecutionDay)

	handlerV1(mockResponseWriter, mockRequest)

	want := http.StatusBadRequest
	got := mockResponseWriter.Result().StatusCode

	assert.Equal(t, want, got)
}

func TestCreateExecutionDayHandlerV1_failsWhenInsertExecutionDayThrowsError(t *testing.T) {
	mockInsertExecutionDay := criteria.MockInsertExecutionDay(errors.New("failed to insert execution day"))
	mockResponseWriter := httptest.NewRecorder()
	mockExecutionDay := criteria.MockExecutionDayDTO(nil)
	mockBody, _ := json.Marshal(mockExecutionDay)
	mockRequest, _ := http.NewRequestWithContext(context.Background(), http.MethodPost, "/criteria/executions/{execution_id}/day/v1", bytes.NewReader(mockBody))
	mockRequest.SetPathValue("execution_id", "1")

	handlerV1 := criteria.CreateExecutionDayHandlerV1(mockInsertExecutionDay)

	handlerV1(mockResponseWriter, mockRequest)

	want := http.StatusInternalServerError
	got := mockResponseWriter.Result().StatusCode

	assert.Equal(t, want, got)
}
