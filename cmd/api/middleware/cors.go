package middleware

import (
	"net/http"
	"os"
	"strings"
	"sync"
)

var (
	allowedOrigins     map[string]bool
	loadAllowedOrigins sync.Once
)

// CORS is an HTTP middleware that adds CORS headers to support requests
// from allowed origins (e.g., local frontend or production domains). It handles
// preflight (OPTIONS) requests and ensures proper caching behavior by setting
// the 'Vary: Origin' header to prevent CORS misconfigurations due to shared caches.
//
// It reads the list of allowed origins from the ALLOWED_ORIGINS environment
// variable (once, safely), and allows credentials and custom headers like
// X-Session-Token if the request's origin is allowed.
func CORS(next http.Handler) http.Handler {
	loadAllowedOrigins.Do(loadOrigins)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		if allowedOrigins[origin] {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Vary", "Origin")
		}

		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Session-Token")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}

// loadOrigins loads the list of allowed origins from the CORS_ALLOWED_ORIGINS environment variable and stores them
// in a map for fast lookup.
//
// Example .env value:
//
//	CORS_ALLOWED_ORIGINS=http://localhost:3000,http://example.com
func loadOrigins() {
	allowedOrigins = make(map[string]bool)
	origins := os.Getenv("CORS_ALLOWED_ORIGINS")
	for _, origin := range strings.Split(origins, ",") {
		origin = strings.TrimSpace(origin)
		if origin != "" {
			allowedOrigins[origin] = true
		}
	}
}
