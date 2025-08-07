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
