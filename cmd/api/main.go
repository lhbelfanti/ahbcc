package main

import (
	"log"
	"net/http"

	"ahbcc/cmd/api/migrations"
	"ahbcc/cmd/api/ping"
	"ahbcc/internal/database"
	"ahbcc/internal/setup"
)

func main() {
	/* --- Dependencies --- */
	// Database
	pg := setup.Init(database.InitPostgres())
	defer pg.Close()

	// Services
	runMigrations := migrations.MakeRun(pg.Database())

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
