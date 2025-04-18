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
	"ahbcc/cmd/api/search/criteria/executions"
	"ahbcc/cmd/api/search/criteria/executions/summary"
	"ahbcc/cmd/api/tweets"
	"ahbcc/cmd/api/tweets/categorized"
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

	// POST /migrations/run/v1 dependencies
	createMigrationsTable := migrations.MakeCreateMigrationsTable(db)
	isMigrationApplied := migrations.MakeIsMigrationApplied(db)
	insertAppliedMigration := migrations.MakeInsertAppliedMigration(db)
	runMigrations := migrations.MakeRun(db, createMigrationsTable, isMigrationApplied, insertAppliedMigration)

	// POST /auth/signup/v1 dependencies
	userExists := user.MakeExists(db)
	insertUser := user.MakeInsert(db)
	signUp := auth.MakeSignUp(userExists, insertUser)

	// POST /auth/login/v1 dependencies
	selectUserByUsername := user.MakeSelectByUsername(db)
	deleteExpiredUserSessions := session.MakeDeleteExpiredSessions(db)
	insertUserSession := session.MakeInsert(db)
	createSessionToken := session.MakeCreateToken(insertUserSession)
	logIn := auth.MakeLogIn(selectUserByUsername, deleteExpiredUserSessions, createSessionToken)

	// POST /auth/logout/v1 dependencies
	deleteUserSession := session.MakeDelete(db)
	logOut := auth.MakeLogOut(deleteUserSession)

	// POST /tweets/v1 dependencies
	insertSingleQuote := quotes.MakeInsertSingle(db)
	deleteOrphanQuotes := quotes.MakeDeleteOrphans(db)
	insertTweets := tweets.MakeInsert(db, insertSingleQuote, deleteOrphanQuotes)

	// POST /criteria/{criteria_id}/enqueue/v1 dependencies
	selectCriteriaByID := criteria.MakeSelectByID(db)
	collectExecutionDAORows := database.MakeCollectRows[executions.ExecutionDAO]()
	selectExecutionsByStatuses := executions.MakeSelectExecutionsByStatuses(db, collectExecutionDAORows)
	insertCriteriaExecution := executions.MakeInsertExecution(db)
	scrapperEnqueueCriteria := scrapper.MakeEnqueueCriteria(httpClient, os.Getenv("ENQUEUE_CRITERIA_API_URL"))
	enqueueCriteria := criteria.MakeEnqueue(selectCriteriaByID, selectExecutionsByStatuses, insertCriteriaExecution, scrapperEnqueueCriteria)

	// POST /criteria/init/v1 dependencies
	selectLastDayExecutedByCriteriaID := executions.MakeSelectLastDayExecutedByCriteriaID(db)
	resumeCriteria := criteria.MakeResume(selectCriteriaByID, selectLastDayExecutedByCriteriaID, selectExecutionsByStatuses, scrapperEnqueueCriteria)
	initCriteria := criteria.MakeInit(selectExecutionsByStatuses, resumeCriteria)

	// GET /criteria/executions/{execution_id}/v1 dependencies
	selectExecutionByID := executions.MakeSelectExecutionByID(db)

	// PUT /criteria/executions/{execution_id}/v1 dependencies
	updateCriteriaExecution := executions.MakeUpdateExecution(db)

	// POST /criteria/executions/{execution_id}/day/v1 dependencies
	insertCriteriaExecutionDay := executions.MakeInsertExecutionDay(db)

	// POST /criteria/executions/summarize/v1 dependencies
	collectSummaryDAORows := database.MakeCollectRows[summary.DAO]()
	selectMonthlyTweetsCountsByYearByCriteriaID := summary.MakeSelectMonthlyTweetsCountsByYearByCriteriaID(db, collectSummaryDAORows)
	selectIDBySearchCriteriaIDYearAndMonth := summary.MakeSelectIDBySearchCriteriaIDYearAndMonth(db)
	insertExecutionSummary := summary.MakeInsert(db)
	updateSummaryTotalTweets := summary.MakeUpdateTotalTweets(db)
	upsertExecutionSummary := summary.MakeUpsert(selectIDBySearchCriteriaIDYearAndMonth, insertExecutionSummary, updateSummaryTotalTweets)
	summarizeCriteriaExecutions := executions.MakeSummarize(db, selectExecutionsByStatuses, selectMonthlyTweetsCountsByYearByCriteriaID, upsertExecutionSummary)

	// POST /users/{user_id}/criteria/v1
	collectCriteriaDAORows := database.MakeCollectRows[criteria.DAO]()
	selectAllCriteriaExecutionsSummaries := summary.MakeSelectAll(db, collectSummaryDAORows)
	selectAllSearchCriteria := criteria.MakeSelectAll(db, collectCriteriaDAORows)
	collectCategorizedTweetsDAORows := database.MakeCollectRows[categorized.DAO]()
	selectAllCategorizedTweets := categorized.MakeSelectAllByUserID(db, collectCategorizedTweetsDAORows)
	information := criteria.MakeInformation(selectAllCriteriaExecutionsSummaries, selectAllSearchCriteria, selectAllCategorizedTweets)

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
	router.HandleFunc("POST /criteria/executions/summarize/v1", executions.SummarizeV1(summarizeCriteriaExecutions))
	router.HandleFunc("GET /users/{user_id}/criteria/v1", criteria.InformationV1(information))
	log.Info(ctx, "Router initialized!")

	/* --- Server --- */
	port := fmt.Sprintf(":%s", os.Getenv("API_PORT"))
	log.Info(ctx, fmt.Sprintf("AHBCC server is ready to receive request on port %s", port))
	setup.Must(http.ListenAndServe(port, router))
}
