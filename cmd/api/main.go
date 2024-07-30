package main

import (
	"log"
	"log/slog"
	"net/http"

	"ahbcc/cmd/api/migrations"
	"ahbcc/cmd/api/ping"
	"ahbcc/cmd/api/tweets"
	"ahbcc/cmd/api/tweets/quotes"
	"ahbcc/internal/database"
	"ahbcc/internal/setup"
)

func main() {
	/* --- Dependencies --- */
	// Database
	pg := setup.Init(database.InitPostgres())
	defer pg.Close()
	db := pg.Database()

	// Services
	runMigrations := migrations.MakeRun(db)
	insertSingleQuote := quotes.MakeInsertSingle(db)
	deleteOrphanQuotes := quotes.MakeDeleteOrphans(db)
	insertTweets := tweets.MakeInsert(db, insertSingleQuote, deleteOrphanQuotes)

	/* --- Router --- */
	router := http.NewServeMux()
	router.HandleFunc("GET /ping/v1", ping.HandlerV1())
	router.HandleFunc("POST /migrations/run/v1", migrations.RunHandlerV1(runMigrations))
	router.HandleFunc("POST /tweets/v1", tweets.InsertHandlerV1(insertTweets))

	/* --- Server --- */
	slog.Info("AHBCC server is ready to receive request on port :8090")
	err := http.ListenAndServe(":8090", router)
	if err != nil {
		log.Fatalf("Could not start server: %s\n", err.Error())
	}
}
