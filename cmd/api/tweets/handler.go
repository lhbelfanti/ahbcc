package tweets

import (
	"encoding/json"
	"net/http"

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
