package corpus_test

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"ahbcc/cmd/api/corpus"
)

func TestCreateCorpusHandlerV1_success(t *testing.T) {
	mockCreateCorpus := corpus.MockCreate(nil)
	mockResponseWriter := httptest.NewRecorder()
	mockRequest, _ := http.NewRequestWithContext(context.Background(), http.MethodPost, "/corpus/v1", nil)

	createCorpusHandlerV1 := corpus.CreateCorpusHandlerV1(mockCreateCorpus)

	createCorpusHandlerV1(mockResponseWriter, mockRequest)

	want := http.StatusOK
	got := mockResponseWriter.Result().StatusCode

	assert.Equal(t, want, got)
}

func TestCreateCorpusHandlerV1_failsWhenCreateCorpusThrowsError(t *testing.T) {
	mockCreateCorpus := corpus.MockCreate(errors.New("failed to create corpus"))
	mockResponseWriter := httptest.NewRecorder()
	mockRequest, _ := http.NewRequestWithContext(context.Background(), http.MethodPost, "/corpus/v1", nil)

	createCorpusHandlerV1 := corpus.CreateCorpusHandlerV1(mockCreateCorpus)

	createCorpusHandlerV1(mockResponseWriter, mockRequest)

	want := http.StatusInternalServerError
	got := mockResponseWriter.Result().StatusCode

	assert.Equal(t, want, got)
}

func TestExportCorpusHandlerV1_successWithJSONExport(t *testing.T) {
	mockJSONExportResult := corpus.MockJSONExportResult()
	mockExportCorpus := corpus.MockExportCorpus(mockJSONExportResult, nil)
	mockResponseWriter := httptest.NewRecorder()
	mockRequest, _ := http.NewRequestWithContext(context.Background(), http.MethodGet, "/corpus/v1?format=json", nil)

	exportCorpusHandlerV1 := corpus.ExportCorpusHandlerV1(mockExportCorpus)

	exportCorpusHandlerV1(mockResponseWriter, mockRequest)

	assert.Equal(t, http.StatusOK, mockResponseWriter.Code)
	assert.Equal(t, "application/json", mockResponseWriter.Header().Get("Content-Type"))
	assert.Equal(t, "attachment; filename=corpus.json", mockResponseWriter.Header().Get("Content-Disposition"))
	assert.Equal(t, mockResponseWriter.Body.String(), string(mockJSONExportResult.Data))
}

func TestExportCorpusHandlerV1_successWithCSVExport(t *testing.T) {
	mockCSVExportResult := corpus.MockCSVExportResult()
	mockExportCorpus := corpus.MockExportCorpus(mockCSVExportResult, nil)
	mockResponseWriter := httptest.NewRecorder()
	mockRequest, _ := http.NewRequestWithContext(context.Background(), http.MethodGet, "/corpus/v1?format=json", nil)

	exportCorpusHandlerV1 := corpus.ExportCorpusHandlerV1(mockExportCorpus)

	exportCorpusHandlerV1(mockResponseWriter, mockRequest)

	assert.Equal(t, http.StatusOK, mockResponseWriter.Code)
	assert.Equal(t, "text/csv", mockResponseWriter.Header().Get("Content-Type"))
	assert.Equal(t, "attachment; filename=corpus.csv", mockResponseWriter.Header().Get("Content-Disposition"))
	assert.Equal(t, mockResponseWriter.Body.String(), string(mockCSVExportResult.Data))
}

func TestExportCorpusHandlerV1_Error(t *testing.T) {
	mockExportCorpus := corpus.MockExportCorpus(nil, errors.New("failed to export corpus"))
	mockResponseWriter := httptest.NewRecorder()
	mockRequest, _ := http.NewRequestWithContext(context.Background(), http.MethodGet, "/corpus/v1", nil)

	exportCorpusHandlerV1 := corpus.ExportCorpusHandlerV1(mockExportCorpus)

	exportCorpusHandlerV1(mockResponseWriter, mockRequest)

	want := http.StatusInternalServerError
	got := mockResponseWriter.Result().StatusCode

	assert.Equal(t, want, got)
}
