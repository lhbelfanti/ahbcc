package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"

	"ahbcc/cmd/api/migrations"
	"ahbcc/cmd/api/ping"
	"ahbcc/internal/setup"
)

const databaseURL string = "postgresql://%s:%s@postgres_db:5432/%s?sslmode=disable"

func main() {
	/* --- Dependencies --- */
	// Database
	pool := setup.Init(pgxpool.New(context.Background(), resolveDatabaseURL()))
	defer pool.Close()

	// Services
	runMigrations := migrations.MakeRun(pool)

	/* --- Router --- */
	router := http.NewServeMux()
	router.HandleFunc("GET /ping/v1", ping.HandlerV1())
	router.HandleFunc("POST /migrations/run/v1", migrations.RunHandlerV1(runMigrations))

	/* --- Server --- */
	log.Println("Starting AHBCC server on :8090")
	err := http.ListenAndServe(":8090", router)
	if err != nil {
		log.Fatalf("Could not start server: %s\n", err.Error())
	}
}

func resolveDatabaseURL() string {
	dbUser := os.Getenv("POSTGRES_DB_USER")
	dbPass := os.Getenv("POSTGRES_DB_PASS")
	dbName := os.Getenv("POSTGRES_DB_NAME")

	return fmt.Sprintf(databaseURL, dbUser, dbPass, dbName)
}
