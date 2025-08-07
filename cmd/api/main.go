package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/rs/zerolog"

	"ahbcc/cmd/api/auth"
	"ahbcc/cmd/api/corpus"
	"ahbcc/cmd/api/middleware"
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

	// POST /tweets/categorized/v1 dependencies
	selectUserIDByToken := session.MakeSelectUserIDByToken(db)
	selectTweetByID := tweets.MakeSelectByID(db)
	selectByUserIDTweetIDAndSearchCriteriaID := categorized.MakeSelectByUserIDTweetIDAndSearchCriteriaID(db)
	insertSingle := categorized.MakeInsertSingle(db)
	insertCategorizedTweet := categorized.MakeInsertCategorizedTweet(selectUserIDByToken, selectTweetByID, selectByUserIDTweetIDAndSearchCriteriaID, insertSingle)

	// GET /criteria/v1
	collectSummaryDAORows := database.MakeCollectRows[summary.DAO](nil)
	selectAllCriteriaExecutionsSummaries := summary.MakeSelectAll(db, collectSummaryDAORows)
	collectCriteriaDAORows := database.MakeCollectRows[criteria.DAO](nil)
	selectAllSearchCriteria := criteria.MakeSelectAll(db, collectCriteriaDAORows)
	collectAnalyzedTweetsDTO := database.MakeCollectRows[categorized.AnalyzedTweetsDTO](nil)
	selectAllCategorizedTweets := categorized.MakeSelectAllByUserID(db, collectAnalyzedTweetsDTO)
	information := criteria.MakeInformation(selectUserIDByToken, selectAllCriteriaExecutionsSummaries, selectAllSearchCriteria, selectAllCategorizedTweets)

	// GET /criteria/{criteria_id}/v1
	selectCriteriaByID := criteria.MakeSelectByID(db)
	summarizedInformation := criteria.MakeSummarizedInformation(selectUserIDByToken, selectCriteriaByID, selectAllCriteriaExecutionsSummaries, selectAllCategorizedTweets)

	// POST /criteria/init/v1 dependencies
	collectExecutionDAORows := database.MakeCollectRows[executions.ExecutionDAO](nil)
	selectExecutionsByStatuses := executions.MakeSelectExecutionsByStatuses(db, collectExecutionDAORows)
	selectLastDayExecutedByCriteriaID := executions.MakeSelectLastDayExecutedByCriteriaID(db)
	scrapperEnqueueCriteria := scrapper.MakeEnqueueCriteria(httpClient, os.Getenv("ENQUEUE_CRITERIA_API_URL"))
	resumeCriteria := criteria.MakeResume(selectCriteriaByID, selectLastDayExecutedByCriteriaID, selectExecutionsByStatuses, scrapperEnqueueCriteria)
	initCriteria := criteria.MakeInit(selectExecutionsByStatuses, resumeCriteria)

	// GET /criteria/{criteria_id}/tweets/v1 dependencies
	tweetsCustomScanner := tweets.CustomScanner()
	collectTweetsDTORows := database.MakeCollectRows[tweets.CustomTweetDTO](tweetsCustomScanner)
	selectBySearchCriteriaIDYearAndMonth := tweets.MakeSelectBySearchCriteriaIDYearAndMonth(db, collectTweetsDTORows, selectUserIDByToken)

	// POST /criteria/{criteria_id}/enqueue/v1 dependencies
	insertCriteriaExecution := executions.MakeInsertExecution(db)
	enqueueCriteria := criteria.MakeEnqueue(selectCriteriaByID, selectExecutionsByStatuses, insertCriteriaExecution, scrapperEnqueueCriteria)

	// POST /criteria-executions/summarize/v1 dependencies
	selectMonthlyTweetsCountsByYearByCriteriaID := summary.MakeSelectMonthlyTweetsCountsByYearByCriteriaID(db, collectSummaryDAORows)
	selectIDBySearchCriteriaIDYearAndMonth := summary.MakeSelectIDBySearchCriteriaIDYearAndMonth(db)
	insertExecutionSummary := summary.MakeInsert(db)
	updateSummaryTotalTweets := summary.MakeUpdateTotalTweets(db)
	upsertExecutionSummary := summary.MakeUpsert(selectIDBySearchCriteriaIDYearAndMonth, insertExecutionSummary, updateSummaryTotalTweets)
	summarizeCriteriaExecutions := executions.MakeSummarize(db, selectExecutionsByStatuses, selectMonthlyTweetsCountsByYearByCriteriaID, upsertExecutionSummary)

	// GET /criteria-executions/{execution_id}/v1 dependencies
	selectExecutionByID := executions.MakeSelectExecutionByID(db)

	// PUT /criteria-executions/{execution_id}/v1 dependencies
	updateCriteriaExecution := executions.MakeUpdateExecution(db)

	// POST /criteria-executions/{execution_id}/day/v1 dependencies
	insertCriteriaExecutionDay := executions.MakeInsertExecutionDay(db)

	//selectTweetByID tweets.SelectByID,
	//selectTweetQuoteByID quotes.SelectByID,
	//deleteAllCorpusRows DeleteAll,
	//insertCorpusRow Insert

	collectCategorizedTweetsDAORows := database.MakeCollectRows[categorized.DAO](nil)
	selectCategorizedTweetsByCategorizations := categorized.MakeSelectByCategorizations(db, collectCategorizedTweetsDAORows)
	selectTweetQuoteByID := quotes.MakeSelectByID(db)
	deleteAllCorpusRows := corpus.MockDeleteAll(nil)
	insertCorpusRow := corpus.MockInsert(nil)
	createCorpus := corpus.MakeCreate(selectCategorizedTweetsByCategorizations, selectTweetByID, selectTweetQuoteByID, deleteAllCorpusRows, insertCorpusRow)

	/* --- Router --- */
	log.Info(ctx, "Initializing router...")
	router := http.NewServeMux()
	router.HandleFunc("GET /ping/v1", ping.HandlerV1())
	router.HandleFunc("POST /migrations/run/v1", migrations.RunHandlerV1(runMigrations))
	router.HandleFunc("POST /auth/signup/v1", auth.SignUpHandlerV1(signUp))
	router.HandleFunc("POST /auth/login/v1", auth.LogInHandlerV1(logIn))
	router.HandleFunc("POST /auth/logout/v1", auth.LogOutHandlerV1(logOut))
	router.HandleFunc("POST /tweets/v1", tweets.InsertHandlerV1(insertTweets))
	router.HandleFunc("POST /tweets/{tweet_id}/categorize/v1", categorized.InsertSingleHandlerV1(insertCategorizedTweet))
	router.HandleFunc("GET /criteria/v1", criteria.InformationHandlerV1(information))
	router.HandleFunc("GET /criteria/{criteria_id}/summarize/v1", criteria.SummarizedInformationHandlerV1(summarizedInformation))
	router.HandleFunc("POST /criteria/init/v1", criteria.InitHandlerV1(initCriteria))
	router.HandleFunc("GET /criteria/{criteria_id}/tweets/v1", tweets.CriteriaTweetsHandlerV1(selectBySearchCriteriaIDYearAndMonth))
	router.HandleFunc("POST /criteria/{criteria_id}/enqueue/v1", criteria.EnqueueHandlerV1(enqueueCriteria))
	router.HandleFunc("POST /criteria-executions/summarize/v1", executions.SummarizeHandlerV1(summarizeCriteriaExecutions))
	router.HandleFunc("GET /criteria-executions/{execution_id}/v1", executions.GetExecutionByIDHandlerV1(selectExecutionByID))
	router.HandleFunc("PUT /criteria-executions/{execution_id}/v1", executions.UpdateExecutionHandlerV1(updateCriteriaExecution))
	router.HandleFunc("POST /criteria-executions/{execution_id}/day/v1", executions.CreateExecutionDayHandlerV1(insertCriteriaExecutionDay))
	router.HandleFunc("POST /corpus/v1", corpus.CreateCorpusHandlerV1(createCorpus))
	log.Info(ctx, "Router initialized!")

	/* --- Middlewares --- */
	handler := middleware.CORS(router)

	/* --- Server --- */
	port := fmt.Sprintf(":%s", os.Getenv("API_PORT"))
	log.Info(ctx, fmt.Sprintf("AHBCC server is ready to receive request on port %s", port))
	setup.Must(http.ListenAndServe(port, handler))
}
