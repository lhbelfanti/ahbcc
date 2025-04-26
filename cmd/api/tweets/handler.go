package tweets

import (
	"encoding/json"
	"net/http"
	"strconv"

	"ahbcc/internal/http/response"
	"ahbcc/internal/log"
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
		if tweet.ID == "" {
			return MissingTweetID
		}

		if tweet.SearchCriteriaID == nil {
			return MissingTweetSearchCriteriaID
		}
	}

	return nil
}

// GetCriteriaTweetsHandlerV1 HTTP Handler of the endpoint /criteria/{criteria_id}/tweets/v1
func GetCriteriaTweetsHandlerV1(selectBySearchCriteriaIDYearAndMonth SelectBySearchCriteriaIDYearAndMonth) http.HandlerFunc {
	const defaultLimit int = 10

	return func(w http.ResponseWriter, r *http.Request) {
		var year, month, limit int

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

		yearQueryParamStr := r.URL.Query().Get("year")
		if yearQueryParamStr != "" {
			year, err = strconv.Atoi(yearQueryParamStr)
			if err != nil {
				response.Send(ctx, w, http.StatusBadRequest, InvalidQueryParameterFormat, nil, err)
				return
			} else {
				// Only retrieve the month if the year is present. Otherwise, the default value is 0, which means all months.
				monthQueryParamStr := r.URL.Query().Get("month")
				if monthQueryParamStr != "" {
					month, err = strconv.Atoi(monthQueryParamStr)
					if err != nil {
						response.Send(ctx, w, http.StatusBadRequest, InvalidQueryParameterFormat, nil, err)
						return
					}
					ctx = log.With(ctx, log.Param("month", monthQueryParamStr))
				}
			}
			ctx = log.With(ctx, log.Param("year", yearQueryParamStr))
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

		response.Send(ctx, w, http.StatusOK, "Tweets successfully retrieved", uncategorizedTweets, nil)
	}
}
