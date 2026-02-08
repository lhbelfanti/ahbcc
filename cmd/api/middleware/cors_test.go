package middleware_test

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/lhbelfanti/corpus-creator/cmd/api/middleware"
)

func TestMain(m *testing.M) {
	_ = os.Setenv("CORS_ALLOWED_ORIGINS", "http://localhost:3000")
	os.Exit(m.Run())
}

func TestCORS_successWithAValidOrigin(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set("Origin", "http://localhost:3000")

	rr := httptest.NewRecorder()

	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	handler := middleware.CORS(next)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, "http://localhost:3000", rr.Header().Get("Access-Control-Allow-Origin"))
	assert.Equal(t, "Origin", rr.Header().Get("Vary"))
	assert.Contains(t, rr.Header().Get("Access-Control-Allow-Methods"), "GET")
	assert.Contains(t, rr.Header().Get("Access-Control-Allow-Headers"), "Authorization")
}

func TestCORS_successHandlingOptionsRequest(t *testing.T) {
	req := httptest.NewRequest(http.MethodOptions, "/", nil)
	req.Header.Set("Origin", "http://localhost:3000")

	rr := httptest.NewRecorder()

	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Fatal("Next handler should not be called on OPTIONS request")
	})

	handler := middleware.CORS(next)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusNoContent, rr.Code)
	assert.Equal(t, "http://localhost:3000", rr.Header().Get("Access-Control-Allow-Origin"))
	assert.Equal(t, "Origin", rr.Header().Get("Vary"))
	assert.Contains(t, rr.Header().Get("Access-Control-Allow-Methods"), "OPTIONS")
}

func TestCORS_failsWithAnInvalidOrigin(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set("Origin", "https://notallowed.com")

	rr := httptest.NewRecorder()

	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	handler := middleware.CORS(next)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Empty(t, rr.Header().Get("Access-Control-Allow-Origin"))
	assert.Equal(t, "", rr.Header().Get("Vary")) // Not set because origin not allowed
	assert.Contains(t, rr.Header().Get("Access-Control-Allow-Methods"), "GET")
	assert.Contains(t, rr.Header().Get("Access-Control-Allow-Headers"), "Content-Type")
}
