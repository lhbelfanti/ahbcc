package categorized

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/lhbelfanti/corpus-creator/internal/http/response"
	"github.com/lhbelfanti/corpus-creator/internal/log"
)

// InsertSingleHandlerV1 HTTP Handler of the endpoint /tweets/{tweet_id}/categorize/v1
func InsertSingleHandlerV1(insertCategorizedTweet InsertCategorizedTweet) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		token := r.Header.Get("X-Session-Token")
		if token == "" {
			response.Send(ctx, w, http.StatusUnauthorized, AuthorizationTokenRequired, nil, AuthorizationTokenIsRequired)
			return
		}
		ctx = log.With(ctx, log.Param("token", token))

		tweetIDParam := r.PathValue("tweet_id")
		tweetID, err := strconv.Atoi(tweetIDParam)
		if err != nil {
			response.Send(ctx, w, http.StatusBadRequest, InvalidURLParameter, nil, err)
			return
		}

		ctx = log.With(ctx, log.Param("tweet_id", tweetID))

		var body InsertSingleBodyDTO
		err = json.NewDecoder(r.Body).Decode(&body)
		if err != nil {
			response.Send(ctx, w, http.StatusBadRequest, InvalidRequestBody, nil, err)
			return
		}
		ctx = log.With(ctx, log.Param("body", body))

		if body.Categorization != VerdictPositive &&
			body.Categorization != VerdictIndeterminate &&
			body.Categorization != VerdictNegative {
			response.Send(ctx, w, http.StatusBadRequest, InvalidRequestBody, nil, InvalidCategorization)
			return
		}

		categorizedTweetID, err := insertCategorizedTweet(ctx, token, tweetID, body)
		if err != nil {
			if errors.Is(err, TweetAlreadyCategorized) {
				response.Send(ctx, w, http.StatusConflict, FailedToCategorizeAnAlreadyCategorizedTweet, nil, TweetAlreadyCategorized)
				return
			}

			response.Send(ctx, w, http.StatusInternalServerError, FailedToInsertCategorizedTweet, nil, err)
			return
		}

		response.Send(ctx, w, http.StatusOK, "Tweet successfully categorized", InsertSingleResponseDTO{ID: categorizedTweetID}, nil)
	}
}
