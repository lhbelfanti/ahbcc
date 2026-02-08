package migrations_test

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/lhbelfanti/corpus-creator/cmd/api/migrations"
)

func TestRunHandlerV1_success(t *testing.T) {
	mockRun := migrations.MockRun(nil)
	mockResponseWriter := httptest.NewRecorder()
	mockRequest, _ := http.NewRequestWithContext(context.Background(), http.MethodPost, "/migrations/run/v1", strings.NewReader(""))

	runHandlerV1 := migrations.RunHandlerV1(mockRun)

	runHandlerV1(mockResponseWriter, mockRequest)

	want := http.StatusOK
	got := mockResponseWriter.Result().StatusCode

	assert.Equal(t, want, got)
}

func TestRunHandlerV1_failsWhenMigrationsRunThrowsError(t *testing.T) {
	mockRun := migrations.MockRun(errors.New("migrations run failed"))
	mockResponseWriter := httptest.NewRecorder()
	mockRequest, _ := http.NewRequestWithContext(context.Background(), http.MethodPost, "/migrations/run/v1", strings.NewReader(""))

	runHandlerV1 := migrations.RunHandlerV1(mockRun)

	runHandlerV1(mockResponseWriter, mockRequest)

	want := http.StatusInternalServerError
	got := mockResponseWriter.Result().StatusCode

	assert.Equal(t, want, got)
}
