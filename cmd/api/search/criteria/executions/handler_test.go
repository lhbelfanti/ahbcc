package executions_test

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

	"github.com/lhbelfanti/corpus-creator/cmd/api/search/criteria/executions"
	"github.com/lhbelfanti/corpus-creator/internal/http/response"
)

func TestGetExecutionByIDHandlerV1_success(t *testing.T) {
	mockExecutionDAO := executions.MockExecutionDAO()
	mockSelectExecutionByID := executions.MockSelectExecutionByID(mockExecutionDAO, nil)
	mockResponseWriter := httptest.NewRecorder()
	mockRequest, _ := http.NewRequestWithContext(context.Background(), http.MethodGet, "/criteria-executions/{execution_id}/v1", http.NoBody)
	mockRequest.SetPathValue("execution_id", "1")

	handlerV1 := executions.GetExecutionByIDHandlerV1(mockSelectExecutionByID)

	handlerV1(mockResponseWriter, mockRequest)

	body, err := io.ReadAll(mockResponseWriter.Result().Body)
	if err != nil {
		t.Fatalf("Failed to read response body: %v", err)
	}

	want := mockExecutionDAO

	var responseDTO response.DTO
	err = json.Unmarshal(body, &responseDTO)
	if err != nil {
		t.Fatalf("Failed to parse response body as JSON: %v", err)
	}

	var got executions.ExecutionDAO
	dataBytes, err := json.Marshal(responseDTO.Data)
	if err != nil {
		t.Fatalf("Failed to marshal Data field: %v", err)
	}

	err = json.Unmarshal(dataBytes, &got)
	if err != nil {
		t.Fatalf("Failed to unmarshal Data field into ExecutionDAO: %v", err)
	}

	assert.Equal(t, want, got)
	assert.Equal(t, "application/json", mockResponseWriter.Header().Get("Content-Type"))
	assert.Equal(t, http.StatusOK, mockResponseWriter.Result().StatusCode)
}

func TestGetExecutionByIDHandlerV1_failsWhenTheURLParamIsEmpty(t *testing.T) {
	mockExecutionDAO := executions.MockExecutionDAO()
	mockSelectExecutionByID := executions.MockSelectExecutionByID(mockExecutionDAO, nil)
	mockResponseWriter := httptest.NewRecorder()
	mockRequest, _ := http.NewRequestWithContext(context.Background(), http.MethodGet, "/criteria-executions/{execution_id}/v1", http.NoBody)

	handlerV1 := executions.GetExecutionByIDHandlerV1(mockSelectExecutionByID)

	handlerV1(mockResponseWriter, mockRequest)

	want := http.StatusBadRequest
	got := mockResponseWriter.Result().StatusCode

	assert.Equal(t, want, got)
}

func TestGetExecutionByIDHandlerV1_failsWhenSelectExecutionByIDThrowsError(t *testing.T) {
	mockSelectExecutionByID := executions.MockSelectExecutionByID(executions.ExecutionDAO{}, errors.New("failed to select execution by id"))
	mockResponseWriter := httptest.NewRecorder()
	mockRequest, _ := http.NewRequestWithContext(context.Background(), http.MethodGet, "/criteria-executions/{execution_id}/v1", http.NoBody)
	mockRequest.SetPathValue("execution_id", "1")

	handlerV1 := executions.GetExecutionByIDHandlerV1(mockSelectExecutionByID)

	handlerV1(mockResponseWriter, mockRequest)

	want := http.StatusInternalServerError
	got := mockResponseWriter.Result().StatusCode

	assert.Equal(t, want, got)
}

func TestUpdateExecutionHandlerV1_success(t *testing.T) {
	mockUpdateExecution := executions.MockUpdateExecution(nil)
	mockResponseWriter := httptest.NewRecorder()
	mockExecution := executions.MockExecutionDTO()
	mockBody, _ := json.Marshal(mockExecution)
	mockRequest, _ := http.NewRequestWithContext(context.Background(), http.MethodPut, "/criteria-executions/{execution_id}/v1", bytes.NewReader(mockBody))
	mockRequest.SetPathValue("execution_id", "1")

	handlerV1 := executions.UpdateExecutionHandlerV1(mockUpdateExecution)

	handlerV1(mockResponseWriter, mockRequest)

	want := http.StatusOK
	got := mockResponseWriter.Result().StatusCode

	assert.Equal(t, want, got)
}

func TestUpdateExecutionHandlerV1_failsWhenTheURLParamIsEmpty(t *testing.T) {
	mockUpdateExecution := executions.MockUpdateExecution(nil)
	mockResponseWriter := httptest.NewRecorder()
	mockExecution := executions.MockExecutionDTO()
	mockBody, _ := json.Marshal(mockExecution)
	mockRequest, _ := http.NewRequestWithContext(context.Background(), http.MethodPut, "/criteria-executions/{execution_id}/v1", bytes.NewReader(mockBody))

	handlerV1 := executions.UpdateExecutionHandlerV1(mockUpdateExecution)

	handlerV1(mockResponseWriter, mockRequest)

	want := http.StatusBadRequest
	got := mockResponseWriter.Result().StatusCode

	assert.Equal(t, want, got)
}

func TestUpdateExecutionHandlerV1_failsWhenTheBodyCannotBeParsed(t *testing.T) {
	mockUpdateExecution := executions.MockUpdateExecution(nil)
	mockResponseWriter := httptest.NewRecorder()
	mockBody, _ := json.Marshal(`{"wrong": "body"}`)
	mockRequest, _ := http.NewRequestWithContext(context.Background(), http.MethodPut, "/criteria-executions/{execution_id}/v1", bytes.NewReader(mockBody))
	mockRequest.SetPathValue("execution_id", "1")

	handlerV1 := executions.UpdateExecutionHandlerV1(mockUpdateExecution)

	handlerV1(mockResponseWriter, mockRequest)

	want := http.StatusBadRequest
	got := mockResponseWriter.Result().StatusCode

	assert.Equal(t, want, got)
}

func TestUpdateExecutionHandlerV1_failsWhenUpdateExecutionThrowsError(t *testing.T) {
	mockUpdateExecution := executions.MockUpdateExecution(errors.New("failed to update execution"))
	mockResponseWriter := httptest.NewRecorder()
	mockExecution := executions.MockExecutionDTO()
	mockBody, _ := json.Marshal(mockExecution)
	mockRequest, _ := http.NewRequestWithContext(context.Background(), http.MethodPut, "/criteria-executions/{execution_id}/v1", bytes.NewReader(mockBody))
	mockRequest.SetPathValue("execution_id", "1")

	handlerV1 := executions.UpdateExecutionHandlerV1(mockUpdateExecution)

	handlerV1(mockResponseWriter, mockRequest)

	want := http.StatusInternalServerError
	got := mockResponseWriter.Result().StatusCode

	assert.Equal(t, want, got)
}

func TestCreateExecutionDayHandlerV1_success(t *testing.T) {
	mockInsertExecutionDay := executions.MockInsertExecutionDay(nil)
	mockResponseWriter := httptest.NewRecorder()
	mockExecutionDay := executions.MockExecutionDayDTO(nil)
	mockBody, _ := json.Marshal(mockExecutionDay)
	mockRequest, _ := http.NewRequestWithContext(context.Background(), http.MethodPost, "/criteria-executions/{execution_id}/day/v1", bytes.NewReader(mockBody))
	mockRequest.SetPathValue("execution_id", "1")

	handlerV1 := executions.CreateExecutionDayHandlerV1(mockInsertExecutionDay)

	handlerV1(mockResponseWriter, mockRequest)

	want := http.StatusOK
	got := mockResponseWriter.Result().StatusCode

	assert.Equal(t, want, got)
}

func TestCreateExecutionDayHandlerV1_failsWhenTheURLParamIsEmpty(t *testing.T) {
	mockInsertExecutionDay := executions.MockInsertExecutionDay(nil)
	mockResponseWriter := httptest.NewRecorder()
	mockExecutionDay := executions.MockExecutionDayDTO(nil)
	mockBody, _ := json.Marshal(mockExecutionDay)
	mockRequest, _ := http.NewRequestWithContext(context.Background(), http.MethodPost, "/criteria-executions/{execution_id}/day/v1", bytes.NewReader(mockBody))

	handlerV1 := executions.CreateExecutionDayHandlerV1(mockInsertExecutionDay)

	handlerV1(mockResponseWriter, mockRequest)

	want := http.StatusBadRequest
	got := mockResponseWriter.Result().StatusCode

	assert.Equal(t, want, got)
}

func TestCreateExecutionDayHandlerV1_failsWhenTheBodyCannotBeParsed(t *testing.T) {
	mockInsertExecutionDay := executions.MockInsertExecutionDay(nil)
	mockResponseWriter := httptest.NewRecorder()
	mockBody, _ := json.Marshal(`{"wrong": "body"}`)
	mockRequest, _ := http.NewRequestWithContext(context.Background(), http.MethodPost, "/criteria-executions/{execution_id}/day/v1", bytes.NewReader(mockBody))
	mockRequest.SetPathValue("execution_id", "1")

	handlerV1 := executions.CreateExecutionDayHandlerV1(mockInsertExecutionDay)

	handlerV1(mockResponseWriter, mockRequest)

	want := http.StatusBadRequest
	got := mockResponseWriter.Result().StatusCode

	assert.Equal(t, want, got)
}

func TestCreateExecutionDayHandlerV1_failsWhenInsertExecutionDayThrowsError(t *testing.T) {
	mockInsertExecutionDay := executions.MockInsertExecutionDay(errors.New("failed to insert execution day"))
	mockResponseWriter := httptest.NewRecorder()
	mockExecutionDay := executions.MockExecutionDayDTO(nil)
	mockBody, _ := json.Marshal(mockExecutionDay)
	mockRequest, _ := http.NewRequestWithContext(context.Background(), http.MethodPost, "/criteria-executions/{execution_id}/day/v1", bytes.NewReader(mockBody))
	mockRequest.SetPathValue("execution_id", "1")

	handlerV1 := executions.CreateExecutionDayHandlerV1(mockInsertExecutionDay)

	handlerV1(mockResponseWriter, mockRequest)

	want := http.StatusInternalServerError
	got := mockResponseWriter.Result().StatusCode

	assert.Equal(t, want, got)
}

func TestSummarizeHandlerV1_success(t *testing.T) {
	mockSummarize := executions.MockSummarize(nil)
	mockResponseWriter := httptest.NewRecorder()
	mockRequest, _ := http.NewRequestWithContext(context.Background(), http.MethodPost, "/criteria-executions/summarize/v1", nil)

	handlerV1 := executions.SummarizeHandlerV1(mockSummarize)

	handlerV1(mockResponseWriter, mockRequest)

	want := http.StatusOK
	got := mockResponseWriter.Result().StatusCode

	assert.Equal(t, want, got)
}

func TestSummarizeHandlerV1_failsWhenSummarizeThrowsError(t *testing.T) {
	mockSummarize := executions.MockSummarize(errors.New("failed to execute summarize"))
	mockResponseWriter := httptest.NewRecorder()
	mockRequest, _ := http.NewRequestWithContext(context.Background(), http.MethodPost, "/criteria-executions/summarize/v1", nil)

	handlerV1 := executions.SummarizeHandlerV1(mockSummarize)

	handlerV1(mockResponseWriter, mockRequest)

	want := http.StatusInternalServerError
	got := mockResponseWriter.Result().StatusCode

	assert.Equal(t, want, got)
}
