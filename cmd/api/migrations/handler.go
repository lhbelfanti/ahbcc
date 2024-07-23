package migrations

import (
	"log/slog"
	"net/http"
)

const migrationsDir string = "./migrations"

// RunHandlerV1 HTTP Handler of the endpoint /migrations/run/v1
func RunHandlerV1(runMigrations Run) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		err := runMigrations(ctx, migrationsDir)
		if err != nil {
			slog.Error(err.Error())
			http.Error(w, FailedToRunMigrations, http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}
