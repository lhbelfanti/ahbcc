package tweets

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/lhbelfanti/corpus-creator/internal/http/response"
	"github.com/lhbelfanti/corpus-creator/internal/log"
)

// InsertHandlerV1 HTTP Handler of the endpoint /tweets/v1
func InsertHandlerV1(insertTweets Insert) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		var tweets []TweetDTO
		err := json.NewDecoder(r.Body).Decode(&tweets)
		if err != nil {
			response.Send(ctx, w, http.StatusBadRequest, InvalidRequestBody, nil, err)
			return
		}
		ctx = log.With(ctx, log.Param("tweets", tweets))

		err = validateBody(tweets)
		if err != nil {
			response.Send(ctx, w, http.StatusBadRequest, InvalidRequestBody, nil, err)
			return
		}

		err = insertTweets(ctx, tweets)
		if err != nil {
			response.Send(ctx, w, http.StatusInternalServerError, FailedToInsertTweetsIntoDatabase, nil, err)
			return
		}

		response.Send(ctx, w, http.StatusOK, "Tweets successfully inserted", nil, nil)
	}
}

// validateBody validates that mandatory fields are present
func validateBody(body []TweetDTO) error {
	for _, tweet := range body {
		if tweet.StatusID == "" {
			return MissingTweetStatusID
		}

		if tweet.SearchCriteriaID == nil {
			return MissingTweetSearchCriteriaID
		}
	}

	return nil
}

// CriteriaTweetsHandlerV1 HTTP Handler of the endpoint /criteria/{criteria_id}/tweets/v1
func CriteriaTweetsHandlerV1(selectBySearchCriteriaIDYearAndMonth SelectBySearchCriteriaIDYearAndMonth) http.HandlerFunc {
	const defaultLimit int = 10

	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		token := r.Header.Get("X-Session-Token")
		if token == "" {
			response.Send(ctx, w, http.StatusUnauthorized, AuthorizationTokenRequired, nil, AuthorizationTokenIsRequired)
			return
		}

		criteriaIDParam := r.PathValue("criteria_id")
		criteriaID, err := strconv.Atoi(criteriaIDParam)
		if err != nil {
			response.Send(ctx, w, http.StatusBadRequest, InvalidURLParameter, nil, err)
			return
		}
		ctx = log.With(ctx, log.Param("criteria_id", criteriaIDParam))

		var year, month, limit int
		yearQueryParamStr := r.URL.Query().Get("year")
		if yearQueryParamStr != "" {
			year, err = strconv.Atoi(yearQueryParamStr)
			if err != nil {
				response.Send(ctx, w, http.StatusBadRequest, InvalidQueryParameterFormat, nil, err)
				return
			}

			// Only retrieve the month if the year is present. Otherwise, the default value is 0, which means all months.
			monthQueryParamStr := r.URL.Query().Get("month")
			if monthQueryParamStr != "" {
				month, err = strconv.Atoi(monthQueryParamStr)
				if err != nil {
					response.Send(ctx, w, http.StatusBadRequest, InvalidQueryParameterFormat, nil, err)
					return
				}
				ctx = log.With(ctx, log.Param("month", month))
			}
			ctx = log.With(ctx, log.Param("year", year))
		}

		limitQueryParamStr := r.URL.Query().Get("limit")
		if limitQueryParamStr != "" {
			limit, err = strconv.Atoi(limitQueryParamStr)
			if err != nil {
				limit = defaultLimit
			}
			ctx = log.With(ctx, log.Param("limit", limitQueryParamStr))
		}

		uncategorizedTweets, err := selectBySearchCriteriaIDYearAndMonth(ctx, criteriaID, year, month, limit, token)
		if err != nil {
			response.Send(ctx, w, http.StatusInternalServerError, FailedToRetrieveTweets, nil, err)
			return
		}

		response.Send(ctx, w, http.StatusOK, "Criteria tweets successfully retrieved", uncategorizedTweets, nil)
	}
}
