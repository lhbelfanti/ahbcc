package main

import (
	"ahbcc/cmd/api/search/criteria/executions"
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
	"ahbcc/cmd/api/user"
	"ahbcc/cmd/api/user/session"
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

	userExists := user.MakeExists(db)
	insertUser := user.MakeInsert(db)
	signUp := auth.MakeSignUp(userExists, insertUser)

	selectUserByUsername := user.MakeSelectByUsername(db)
	deleteExpiredUserSessions := session.MakeDeleteExpiredSessions(db)
	insertUserSession := session.MakeInsert(db)
	createSessionToken := session.MakeCreateToken(insertUserSession)
	logIn := auth.MakeLogIn(selectUserByUsername, deleteExpiredUserSessions, createSessionToken)

	deleteUserSession := session.MakeDelete(db)
	logOut := auth.MakeLogOut(deleteUserSession)

	insertSingleQuote := quotes.MakeInsertSingle(db)
	deleteOrphanQuotes := quotes.MakeDeleteOrphans(db)
	insertTweets := tweets.MakeInsert(db, insertSingleQuote, deleteOrphanQuotes)

	selectCriteriaByID := criteria.MakeSelectByID(db)
	collectExecutionDAORows := database.MakeCollectRows[executions.ExecutionDAO]()
	selectExecutionsByStatuses := executions.MakeSelectExecutionsByStatuses(db, collectExecutionDAORows)
	insertCriteriaExecution := executions.MakeInsertExecution(db)
	scrapperEnqueueCriteria := scrapper.MakeEnqueueCriteria(httpClient, os.Getenv("ENQUEUE_CRITERIA_API_URL"))
	enqueueCriteria := criteria.MakeEnqueue(selectCriteriaByID, selectExecutionsByStatuses, insertCriteriaExecution, scrapperEnqueueCriteria)

	selectLastDayExecutedByCriteriaID := executions.MakeSelectLastDayExecutedByCriteriaID(db)
	resumeCriteria := criteria.MakeResume(selectCriteriaByID, selectLastDayExecutedByCriteriaID, selectExecutionsByStatuses, scrapperEnqueueCriteria)
	initCriteria := criteria.MakeInit(selectExecutionsByStatuses, resumeCriteria)

	selectExecutionByID := executions.MakeSelectExecutionByID(db)

	updateCriteriaExecution := executions.MakeUpdateExecution(db)

	insertCriteriaExecutionDay := executions.MakeInsertExecutionDay(db)

	/* --- Router --- */
	log.Info(ctx, "Initializing router...")
	router := http.NewServeMux()
	router.HandleFunc("GET /ping/v1", ping.HandlerV1())
	router.HandleFunc("POST /migrations/run/v1", migrations.RunHandlerV1(runMigrations))
	router.HandleFunc("POST /auth/signup/v1", auth.SignUpHandlerV1(signUp))
	router.HandleFunc("POST /auth/login/v1", auth.LogInHandlerV1(logIn))
	router.HandleFunc("POST /auth/logout/v1", auth.LogOutHandlerV1(logOut))
	router.HandleFunc("POST /tweets/v1", tweets.InsertHandlerV1(insertTweets))
	router.HandleFunc("POST /criteria/{criteria_id}/enqueue/v1", criteria.EnqueueHandlerV1(enqueueCriteria))
	router.HandleFunc("POST /criteria/init/v1", criteria.InitHandlerV1(initCriteria))
	router.HandleFunc("GET /criteria/executions/{execution_id}/v1", executions.GetExecutionByIDHandlerV1(selectExecutionByID))
	router.HandleFunc("PUT /criteria/executions/{execution_id}/v1", executions.UpdateExecutionHandlerV1(updateCriteriaExecution))
	router.HandleFunc("POST /criteria/executions/{execution_id}/day/v1", executions.CreateExecutionDayHandlerV1(insertCriteriaExecutionDay))
	log.Info(ctx, "Router initialized!")

	/* --- Server --- */
	port := fmt.Sprintf(":%s", os.Getenv("API_PORT"))
	log.Info(ctx, fmt.Sprintf("AHBCC server is ready to receive request on port %s", port))
	setup.Must(http.ListenAndServe(port, router))
}
