package tweets

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

// InsertHandlerV1 HTTP Handler of the endpoint /tweets/v1
func InsertHandlerV1(insertTweets Insert) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var tweets []TweetDTO
		err := json.NewDecoder(r.Body).Decode(&tweets)
		if err != nil {
			slog.Error(err.Error())
			http.Error(w, InvalidRequestBody, http.StatusBadRequest)
			return
		}

		err = validateBody(tweets)
		if err != nil {
			slog.Error(err.Error())
			http.Error(w, InvalidRequestBody, http.StatusBadRequest)
		}

		err = insertTweets(tweets)
		if err != nil {
			slog.Error(err.Error())
			http.Error(w, FailedToInsertTweetsIntoDatabase, http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

// validateBody validates that mandatory fields are present
func validateBody(body []TweetDTO) error {
	for _, tweet := range body {
		if tweet.Hash == nil {
			return MissingTweetHash
		}

		if tweet.SearchCriteriaID == nil {
			return MissingTweetSearchCriteriaID
		}
	}

	return nil
}
