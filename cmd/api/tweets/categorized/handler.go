package categorized

import (
	"encoding/json"
	"net/http"

	"ahbcc/internal/http/response"
	"ahbcc/internal/log"
)

// InsertSingleHandlerV1 HTTP Handler of the endpoint /tweets/categorized/v1
func InsertSingleHandlerV1(insertCategorizedTweet InsertCategorizedTweet) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		token := r.Header.Get("X-Session-Token")
		if token == "" {
			response.Send(ctx, w, http.StatusUnauthorized, AuthorizationTokenRequired, nil, AuthorizationTokenIsRequired)
			return
		}
		ctx = log.With(ctx, log.Param("token", token))

		var categorizedTweet DTO
		err := json.NewDecoder(r.Body).Decode(&categorizedTweet)
		if err != nil {
			response.Send(ctx, w, http.StatusBadRequest, InvalidRequestBody, nil, err)
			return
		}
		ctx = log.With(ctx, log.Param("categorized_tweet", categorizedTweet))

		err = validateBody(categorizedTweet)
		if err != nil {
			response.Send(ctx, w, http.StatusBadRequest, InvalidRequestBody, nil, err)
			return
		}

		categorizedTweetID, err := insertCategorizedTweet(ctx, token, categorizedTweet)
		if err != nil {
			response.Send(ctx, w, http.StatusInternalServerError, FailedToInsertCategorizedTweet, nil, err)
			return
		}

		response.Send(ctx, w, http.StatusOK, "Tweet successfully categorized", InsertSingleResponseDTO{ID: categorizedTweetID}, nil)
	}
}

// validateBody validates that mandatory fields are present
func validateBody(body DTO) error {
	if body.SearchCriteriaID <= 0 {
		return InvalidSearchCriteriaID
	}

	if body.TweetID <= 0 {
		return InvalidTweetID
	}

	if body.Categorization != VerdictPositive &&
		body.Categorization != VerdictIndeterminate &&
		body.Categorization != VerdictNegative {
		return InvalidCategorization
	}

	return nil
}
