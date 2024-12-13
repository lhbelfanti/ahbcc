package tweets_test

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"ahbcc/cmd/api/tweets"
)

func TestInsertHandlerV1_success(t *testing.T) {
	mockInsert := tweets.MockInsert(nil)
	mockResponseWriter := httptest.NewRecorder()
	mockTweets := tweets.MockTweetsDTOs()
	mockBody, _ := json.Marshal(mockTweets)
	mockRequest, _ := http.NewRequestWithContext(context.Background(), http.MethodPost, "/tweets/v1", bytes.NewReader(mockBody))

	insertHandlerV1 := tweets.InsertHandlerV1(mockInsert)

	insertHandlerV1(mockResponseWriter, mockRequest)

	want := http.StatusOK
	got := mockResponseWriter.Result().StatusCode

	assert.Equal(t, want, got)
}

func TestInsertHandlerV1_failsWhenTheBodyCantBeParsed(t *testing.T) {
	mockInsert := tweets.MockInsert(nil)
	mockResponseWriter := httptest.NewRecorder()
	mockBody, _ := json.Marshal(`{"wrong": "body"}`)
	mockRequest, _ := http.NewRequestWithContext(context.Background(), http.MethodPost, "/tweets/v1", bytes.NewReader(mockBody))

	insertHandlerV1 := tweets.InsertHandlerV1(mockInsert)

	insertHandlerV1(mockResponseWriter, mockRequest)

	want := http.StatusBadRequest
	got := mockResponseWriter.Result().StatusCode

	assert.Equal(t, want, got)
}

func TestInsertHandlerV1_failsWhenTweetIDIsNotPresentInBody(t *testing.T) {
	mockInsert := tweets.MockInsert(nil)
	mockResponseWriter := httptest.NewRecorder()
	mockTweets := tweets.MockTweetsDTOs()
	mockTweets[0].ID = ""
	mockBody, _ := json.Marshal(mockTweets)
	mockRequest, _ := http.NewRequestWithContext(context.Background(), http.MethodPost, "/tweets/v1", bytes.NewReader(mockBody))

	insertHandlerV1 := tweets.InsertHandlerV1(mockInsert)

	insertHandlerV1(mockResponseWriter, mockRequest)

	want := http.StatusBadRequest
	got := mockResponseWriter.Result().StatusCode

	assert.Equal(t, want, got)
}

func TestInsertHandlerV1_failsWhenTweetSearchCriteriaIDIsNotPresentInBody(t *testing.T) {
	mockInsert := tweets.MockInsert(nil)
	mockResponseWriter := httptest.NewRecorder()
	mockTweets := tweets.MockTweetsDTOs()
	mockTweets[0].SearchCriteriaID = nil
	mockBody, _ := json.Marshal(mockTweets)
	mockRequest, _ := http.NewRequestWithContext(context.Background(), http.MethodPost, "/tweets/v1", bytes.NewReader(mockBody))

	insertHandlerV1 := tweets.InsertHandlerV1(mockInsert)

	insertHandlerV1(mockResponseWriter, mockRequest)

	want := http.StatusBadRequest
	got := mockResponseWriter.Result().StatusCode

	assert.Equal(t, want, got)
}

func TestInsertHandlerV1_failsWhenInsertTweetsThrowsError(t *testing.T) {
	mockInsert := tweets.MockInsert(errors.New("failed to insert tweets"))
	mockResponseWriter := httptest.NewRecorder()
	mockTweets := tweets.MockTweetsDTOs()
	mockBody, _ := json.Marshal(mockTweets)
	mockRequest, _ := http.NewRequestWithContext(context.Background(), http.MethodPost, "/tweets/v1", bytes.NewReader(mockBody))

	insertHandlerV1 := tweets.InsertHandlerV1(mockInsert)

	insertHandlerV1(mockResponseWriter, mockRequest)

	want := http.StatusInternalServerError
	got := mockResponseWriter.Result().StatusCode

	assert.Equal(t, want, got)
}
