package migrations

import (
	"net/http"

	"github.com/lhbelfanti/corpus-creator/internal/http/response"
)

const migrationsDir string = "./migrations"

// RunHandlerV1 HTTP Handler of the endpoint /migrations/run/v1
func RunHandlerV1(runMigrations Run) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		err := runMigrations(ctx, migrationsDir)
		if err != nil {
			response.Send(ctx, w, http.StatusInternalServerError, FailedToRunMigrations, nil, err)
			return
		}

		response.Send(ctx, w, http.StatusOK, "Migrations successfully run", nil, nil)
	}
}
