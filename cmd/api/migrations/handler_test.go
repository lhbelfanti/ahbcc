package migrations_test

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"ahbcc/cmd/api/migrations"
)

func TestRunHandlerV1_success(t *testing.T) {
	mockRun := migrations.MockRun(nil)
	mockResponseWriter := httptest.NewRecorder()
	mockRequest, _ := http.NewRequestWithContext(context.Background(), http.MethodPost, "/run-migrations/v1", strings.NewReader(""))

	runHandlerV1 := migrations.RunHandlerV1(mockRun)

	runHandlerV1(mockResponseWriter, mockRequest)

	want := http.StatusOK
	got := mockResponseWriter.Code

	assert.Equal(t, want, got)
}

func TestRunHandlerV1_failsWhenMigrationsRunThrowsError(t *testing.T) {
	mockRun := migrations.MockRun(errors.New("migrations run failed"))
	mockResponseWriter := httptest.NewRecorder()
	mockRequest, _ := http.NewRequestWithContext(context.Background(), http.MethodPost, "/run-migrations/v1", strings.NewReader(""))

	runHandlerV1 := migrations.RunHandlerV1(mockRun)

	runHandlerV1(mockResponseWriter, mockRequest)

	want := http.StatusInternalServerError
	got := mockResponseWriter.Code

	assert.Equal(t, want, got)
}
