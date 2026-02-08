package categorized_test

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/lhbelfanti/corpus-creator/cmd/api/tweets/categorized"
)

func TestInsertSingleHandlerV1_success(t *testing.T) {
	mockInsertCategorizedTweet := categorized.MockInsertCategorizedTweet(1, nil)
	mockResponseWriter := httptest.NewRecorder()
	mockBody := categorized.MockInsertSingleBodyDTO(categorized.VerdictPositive)
	bodyBytes, _ := json.Marshal(mockBody)
	mockRequest, _ := http.NewRequestWithContext(context.Background(), http.MethodPost, "/tweets/{tweet_id}/categorize/v1", bytes.NewReader(bodyBytes))
	mockRequest.Header.Set("X-Session-Token", "token")
	mockRequest.SetPathValue("tweet_id", "123")

	insertSingleHandlerV1 := categorized.InsertSingleHandlerV1(mockInsertCategorizedTweet)

	insertSingleHandlerV1(mockResponseWriter, mockRequest)

	want := http.StatusOK
	got := mockResponseWriter.Result().StatusCode

	assert.Equal(t, want, got)

	var response struct {
		Data categorized.InsertSingleResponseDTO `json:"data"`
	}
	err := json.NewDecoder(mockResponseWriter.Body).Decode(&response)
	assert.NoError(t, err)

	assert.Equal(t, 1, response.Data.ID)
}

func TestInsertSingleHandlerV1_failsWhenTheBodyCannotBeParsed(t *testing.T) {
	mockInsertCategorizedTweet := categorized.MockInsertCategorizedTweet(1, nil)
	mockResponseWriter := httptest.NewRecorder()
	mockBody, _ := json.Marshal(`{"wrong": "body"}`)
	mockRequest, _ := http.NewRequestWithContext(context.Background(), http.MethodPost, "/tweets/{tweet_id}/categorize/v1", bytes.NewReader(mockBody))
	mockRequest.Header.Set("X-Session-Token", "token")
	mockRequest.SetPathValue("tweet_id", "123")

	insertSingleHandlerV1 := categorized.InsertSingleHandlerV1(mockInsertCategorizedTweet)

	insertSingleHandlerV1(mockResponseWriter, mockRequest)

	want := http.StatusBadRequest
	got := mockResponseWriter.Result().StatusCode

	assert.Equal(t, want, got)
}

func TestInsertSingleHandlerV1_failsWhenTweetIDIsInvalid(t *testing.T) {
	mockInsertCategorizedTweet := categorized.MockInsertCategorizedTweet(-1, categorized.InvalidTweetID)
	mockResponseWriter := httptest.NewRecorder()
	mockBody := categorized.MockInsertSingleBodyDTO(categorized.VerdictPositive)
	bodyBytes, _ := json.Marshal(mockBody)
	mockRequest, _ := http.NewRequestWithContext(context.Background(), http.MethodPost, "/tweets/{tweet_id}/categorize/v1", bytes.NewReader(bodyBytes))
	mockRequest.Header.Set("X-Session-Token", "token")
	mockRequest.SetPathValue("tweet_id", "invalid")

	insertSingleHandlerV1 := categorized.InsertSingleHandlerV1(mockInsertCategorizedTweet)

	insertSingleHandlerV1(mockResponseWriter, mockRequest)

	want := http.StatusBadRequest
	got := mockResponseWriter.Result().StatusCode

	assert.Equal(t, want, got)
}

func TestInsertSingleHandlerV1_failsWhenTokenIsMissing(t *testing.T) {
	mockInsertCategorizedTweet := categorized.MockInsertCategorizedTweet(-1, errors.New("token is missing"))
	mockResponseWriter := httptest.NewRecorder()
	mockBody := categorized.MockInsertSingleBodyDTO(categorized.VerdictPositive)
	bodyBytes, _ := json.Marshal(mockBody)
	mockRequest, _ := http.NewRequestWithContext(context.Background(), http.MethodPost, "/tweets/{tweet_id}/categorize/v1", bytes.NewReader(bodyBytes))
	mockRequest.SetPathValue("tweet_id", "123")

	insertSingleHandlerV1 := categorized.InsertSingleHandlerV1(mockInsertCategorizedTweet)

	insertSingleHandlerV1(mockResponseWriter, mockRequest)

	want := http.StatusUnauthorized
	got := mockResponseWriter.Result().StatusCode

	assert.Equal(t, want, got)
}

func TestInsertSingleHandlerV1_failsWhenCategorizationIsInvalid(t *testing.T) {
	mockInsertCategorizedTweet := categorized.MockInsertCategorizedTweet(-1, errors.New("invalid categorization"))
	mockResponseWriter := httptest.NewRecorder()
	mockBody := categorized.MockInsertSingleBodyDTO("INVALID")
	bodyBytes, _ := json.Marshal(mockBody)
	mockRequest, _ := http.NewRequestWithContext(context.Background(), http.MethodPost, "/tweets/123/categorize/v1", bytes.NewReader(bodyBytes))
	mockRequest.Header.Set("X-Session-Token", "token")
	mockRequest.SetPathValue("tweet_id", "123")

	insertSingleHandlerV1 := categorized.InsertSingleHandlerV1(mockInsertCategorizedTweet)

	insertSingleHandlerV1(mockResponseWriter, mockRequest)

	want := http.StatusBadRequest
	got := mockResponseWriter.Result().StatusCode

	assert.Equal(t, want, got)
}

func TestInsertSingleHandlerV1_failsWhenInsertThrowsTweetAlreadyCategorizedError(t *testing.T) {
	mockInsertCategorizedTweet := categorized.MockInsertCategorizedTweet(-1, categorized.TweetAlreadyCategorized)
	mockResponseWriter := httptest.NewRecorder()
	mockBody := categorized.MockInsertSingleBodyDTO(categorized.VerdictPositive)
	bodyBytes, _ := json.Marshal(mockBody)
	mockRequest, _ := http.NewRequestWithContext(context.Background(), http.MethodPost, "/tweets/123/categorize/v1", bytes.NewReader(bodyBytes))
	mockRequest.Header.Set("X-Session-Token", "token")
	mockRequest.SetPathValue("tweet_id", "123")

	insertSingleHandlerV1 := categorized.InsertSingleHandlerV1(mockInsertCategorizedTweet)

	insertSingleHandlerV1(mockResponseWriter, mockRequest)

	want := http.StatusConflict
	got := mockResponseWriter.Result().StatusCode

	assert.Equal(t, want, got)
}

func TestInsertSingleHandlerV1_failsWhenInsertThrowsError(t *testing.T) {
	mockInsertCategorizedTweet := categorized.MockInsertCategorizedTweet(-1, errors.New("failed to insert categorized tweet"))
	mockResponseWriter := httptest.NewRecorder()
	mockBody := categorized.MockInsertSingleBodyDTO(categorized.VerdictPositive)
	bodyBytes, _ := json.Marshal(mockBody)
	mockRequest, _ := http.NewRequestWithContext(context.Background(), http.MethodPost, "/tweets/123/categorize/v1", bytes.NewReader(bodyBytes))
	mockRequest.Header.Set("X-Session-Token", "token")
	mockRequest.SetPathValue("tweet_id", "123")

	insertSingleHandlerV1 := categorized.InsertSingleHandlerV1(mockInsertCategorizedTweet)

	insertSingleHandlerV1(mockResponseWriter, mockRequest)

	want := http.StatusInternalServerError
	got := mockResponseWriter.Result().StatusCode

	assert.Equal(t, want, got)
}
