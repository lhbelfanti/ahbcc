package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/rs/zerolog"

	"ahbcc/cmd/api/auth"
	"ahbcc/cmd/api/migrations"
	"ahbcc/cmd/api/ping"
	"ahbcc/cmd/api/search/criteria"
	"ahbcc/cmd/api/tweets"
	"ahbcc/cmd/api/tweets/quotes"
	"ahbcc/cmd/api/users"
	"ahbcc/internal/database"
	_http "ahbcc/internal/http"
	"ahbcc/internal/log"
	"ahbcc/internal/scrapper"
	"ahbcc/internal/setup"
)

var prodEnv bool

func init() {
	flag.BoolVar(&prodEnv, "prod", false, "Run in production environment")
	flag.Parse()
}

func main() {
	/* --- Dependencies --- */
	ctx := context.Background()

	logLevel := zerolog.DebugLevel
	if prodEnv {
		logLevel = zerolog.InfoLevel
	}

	log.NewCustomLogger(os.Stdout, logLevel)

	httpClient := _http.NewClient()

	// Database
	pg := setup.Init(database.InitPostgres())
	defer pg.Close()
	db := pg.Database()

	// Services
	createMigrationsTable := migrations.MakeCreateMigrationsTable(db)
	isMigrationApplied := migrations.MakeIsMigrationApplied(db)
	insertAppliedMigration := migrations.MakeInsertAppliedMigration(db)
	runMigrations := migrations.MakeRun(db, createMigrationsTable, isMigrationApplied, insertAppliedMigration)

	userExists := users.MakeUserExists(db)
	insertUser := users.MakeInsert(db)
	signUp := auth.MakeSignUp(userExists, insertUser)

	insertSingleQuote := quotes.MakeInsertSingle(db)
	deleteOrphanQuotes := quotes.MakeDeleteOrphans(db)
	insertTweets := tweets.MakeInsert(db, insertSingleQuote, deleteOrphanQuotes)

	selectCriteriaByID := criteria.MakeSelectByID(db)
	collectExecutionDAORows := database.MakeCollectRows[criteria.ExecutionDAO]()
	selectExecutionsByStatuses := criteria.MakeSelectExecutionsByStatuses(db, collectExecutionDAORows)
	insertCriteriaExecution := criteria.MakeInsertExecution(db)
	scrapperEnqueueCriteria := scrapper.MakeEnqueueCriteria(httpClient, os.Getenv("ENQUEUE_CRITERIA_API_URL"))
	enqueueCriteria := criteria.MakeEnqueue(selectCriteriaByID, selectExecutionsByStatuses, insertCriteriaExecution, scrapperEnqueueCriteria)

	selectLastDayExecutedByCriteriaID := criteria.MakeSelectLastDayExecutedByCriteriaID(db)
	resumeCriteria := criteria.MakeResume(selectCriteriaByID, selectLastDayExecutedByCriteriaID, selectExecutionsByStatuses, scrapperEnqueueCriteria)
	initCriteria := criteria.MakeInit(selectExecutionsByStatuses, resumeCriteria)

	selectExecutionByID := criteria.MakeSelectExecutionByID(db)

	updateCriteriaExecution := criteria.MakeUpdateExecution(db)

	insertCriteriaExecutionDay := criteria.MakeInsertExecutionDay(db)

	/* --- Router --- */
	log.Info(ctx, "Initializing router...")
	router := http.NewServeMux()
	router.HandleFunc("GET /ping/v1", ping.HandlerV1())
	router.HandleFunc("POST /migrations/run/v1", migrations.RunHandlerV1(runMigrations))
	router.HandleFunc("POST /auth/signup/v1", auth.SignUpHandlerV1(signUp))
	router.HandleFunc("POST /tweets/v1", tweets.InsertHandlerV1(insertTweets))
	router.HandleFunc("POST /criteria/{criteria_id}/enqueue/v1", criteria.EnqueueHandlerV1(enqueueCriteria))
	router.HandleFunc("POST /criteria/init/v1", criteria.InitHandlerV1(initCriteria))
	router.HandleFunc("GET /criteria/executions/{execution_id}/v1", criteria.GetExecutionByIDHandlerV1(selectExecutionByID))
	router.HandleFunc("PUT /criteria/executions/{execution_id}/v1", criteria.UpdateExecutionHandlerV1(updateCriteriaExecution))
	router.HandleFunc("POST /criteria/executions/{execution_id}/day/v1", criteria.CreateExecutionDayHandlerV1(insertCriteriaExecutionDay))
	log.Info(ctx, "Router initialized!")

	/* --- Server --- */
	port := fmt.Sprintf(":%s", os.Getenv("API_PORT"))
	log.Info(ctx, fmt.Sprintf("AHBCC server is ready to receive request on port %s", port))
	setup.Must(http.ListenAndServe(port, router))
}
