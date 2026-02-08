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

	"github.com/lhbelfanti/corpus-creator/cmd/api/tweets"
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

func TestInsertHandlerV1_failsWhenTheBodyCannotBeParsed(t *testing.T) {
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
	mockTweets[0].StatusID = ""
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

func TestCriteriaTweetsHandlerV1_success(t *testing.T) {
	mockTweets := tweets.MockCustomTweetDTOs()
	mockSelectBySearchCriteriaIDYearAndMonth := tweets.MockSelectBySearchCriteriaIDYearAndMonth(mockTweets, nil)
	mockResponseWriter := httptest.NewRecorder()
	mockRequest, _ := http.NewRequestWithContext(context.Background(), http.MethodGet, "/criteria/{criteria_id}/tweets/v1", nil)
	mockRequest.SetPathValue("criteria_id", "1")
	mockRequest.Header.Set("X-Session-Token", "token")
	mockURLQuery := mockRequest.URL.Query()
	mockURLQuery.Add("year", "2025")
	mockURLQuery.Add("month", "1")
	mockURLQuery.Add("limit", "2")
	mockRequest.URL.RawQuery = mockURLQuery.Encode()

	criteriaTweetsV1 := tweets.CriteriaTweetsHandlerV1(mockSelectBySearchCriteriaIDYearAndMonth)

	criteriaTweetsV1(mockResponseWriter, mockRequest)

	want := http.StatusOK
	got := mockResponseWriter.Result().StatusCode

	assert.Equal(t, want, got)
}

func TestCriteriaTweetsHandlerV1_successWithoutQueryParamsOrWithWrongLimitQueryParam(t *testing.T) {
	tests := []struct {
		params map[string]string
	}{
		{params: map[string]string{}},
		{params: map[string]string{"limit": "wrong"}},
	}

	mockTweets := tweets.MockCustomTweetDTOs()
	mockSelectBySearchCriteriaIDYearAndMonth := tweets.MockSelectBySearchCriteriaIDYearAndMonth(mockTweets, nil)
	mockResponseWriter := httptest.NewRecorder()

	for _, tt := range tests {
		mockRequest, _ := http.NewRequestWithContext(context.Background(), http.MethodGet, "/criteria/{criteria_id}/tweets/v1", nil)
		mockRequest.SetPathValue("criteria_id", "1")
		mockRequest.Header.Set("X-Session-Token", "token")
		mockURLQuery := mockRequest.URL.Query()
		for k, v := range tt.params {
			mockURLQuery.Add(k, v)
		}
		mockRequest.URL.RawQuery = mockURLQuery.Encode()

		criteriaTweetsV1 := tweets.CriteriaTweetsHandlerV1(mockSelectBySearchCriteriaIDYearAndMonth)

		criteriaTweetsV1(mockResponseWriter, mockRequest)

		want := http.StatusOK
		got := mockResponseWriter.Result().StatusCode

		assert.Equal(t, want, got)
	}
}

func TestCriteriaTweetsHandlerV1_failsWhenSessionTokenHeaderWasNotFound(t *testing.T) {
	mockTweets := tweets.MockCustomTweetDTOs()
	mockSelectBySearchCriteriaIDYearAndMonth := tweets.MockSelectBySearchCriteriaIDYearAndMonth(mockTweets, nil)
	mockResponseWriter := httptest.NewRecorder()
	mockRequest, _ := http.NewRequestWithContext(context.Background(), http.MethodGet, "/criteria/{criteria_id}/tweets/v1", nil)
	mockRequest.SetPathValue("criteria_id", "1")
	mockURLQuery := mockRequest.URL.Query()
	mockURLQuery.Add("year", "2025")
	mockURLQuery.Add("month", "1")
	mockURLQuery.Add("limit", "2")
	mockRequest.URL.RawQuery = mockURLQuery.Encode()

	criteriaTweetsV1 := tweets.CriteriaTweetsHandlerV1(mockSelectBySearchCriteriaIDYearAndMonth)

	criteriaTweetsV1(mockResponseWriter, mockRequest)

	want := http.StatusUnauthorized
	got := mockResponseWriter.Result().StatusCode

	assert.Equal(t, want, got)
}

func TestCriteriaTweetsHandlerV1_failsWhenTheURLParamCannotBeParsed(t *testing.T) {
	mockTweets := tweets.MockCustomTweetDTOs()
	mockSelectBySearchCriteriaIDYearAndMonth := tweets.MockSelectBySearchCriteriaIDYearAndMonth(mockTweets, nil)
	mockResponseWriter := httptest.NewRecorder()
	mockRequest, _ := http.NewRequestWithContext(context.Background(), http.MethodGet, "/criteria/{criteria_id}/tweets/v1", nil)
	mockRequest.SetPathValue("criteria_id", "wrong")
	mockRequest.Header.Set("X-Session-Token", "token")

	criteriaTweetsV1 := tweets.CriteriaTweetsHandlerV1(mockSelectBySearchCriteriaIDYearAndMonth)

	criteriaTweetsV1(mockResponseWriter, mockRequest)

	want := http.StatusBadRequest
	got := mockResponseWriter.Result().StatusCode

	assert.Equal(t, want, got)
}

func TestCriteriaTweetsHandlerV1_failsWhenYearOrMonthQueryParamsCannotBeParsed(t *testing.T) {
	tests := []struct {
		params map[string]string
	}{
		{params: map[string]string{"year": "wrong", "month": "1"}},
		{params: map[string]string{"year": "2025", "month": "wrong"}},
	}

	mockTweets := tweets.MockCustomTweetDTOs()
	mockSelectBySearchCriteriaIDYearAndMonth := tweets.MockSelectBySearchCriteriaIDYearAndMonth(mockTweets, nil)
	mockResponseWriter := httptest.NewRecorder()

	for _, tt := range tests {
		mockRequest, _ := http.NewRequestWithContext(context.Background(), http.MethodGet, "/criteria/{criteria_id}/tweets/v1", nil)
		mockRequest.SetPathValue("criteria_id", "1")
		mockRequest.Header.Set("X-Session-Token", "token")
		mockURLQuery := mockRequest.URL.Query()
		for k, v := range tt.params {
			mockURLQuery.Add(k, v)
		}
		mockRequest.URL.RawQuery = mockURLQuery.Encode()

		criteriaTweetsV1 := tweets.CriteriaTweetsHandlerV1(mockSelectBySearchCriteriaIDYearAndMonth)

		criteriaTweetsV1(mockResponseWriter, mockRequest)

		want := http.StatusBadRequest
		got := mockResponseWriter.Result().StatusCode

		assert.Equal(t, want, got)
	}
}

func TestCriteriaTweetsHandlerV1_failsWhenSelectBySearchCriteriaIDYearAndMonthThrowsError(t *testing.T) {
	mockTweets := tweets.MockCustomTweetDTOs()
	mockSelectBySearchCriteriaIDYearAndMonth := tweets.MockSelectBySearchCriteriaIDYearAndMonth(mockTweets, errors.New("failed to retrieve tweets"))
	mockResponseWriter := httptest.NewRecorder()
	mockRequest, _ := http.NewRequestWithContext(context.Background(), http.MethodGet, "/criteria/{criteria_id}/tweets/v1", nil)
	mockRequest.SetPathValue("criteria_id", "1")
	mockRequest.Header.Set("X-Session-Token", "token")
	mockURLQuery := mockRequest.URL.Query()
	mockURLQuery.Add("year", "2025")
	mockURLQuery.Add("month", "1")
	mockURLQuery.Add("limit", "2")
	mockRequest.URL.RawQuery = mockURLQuery.Encode()

	criteriaTweetsV1 := tweets.CriteriaTweetsHandlerV1(mockSelectBySearchCriteriaIDYearAndMonth)

	criteriaTweetsV1(mockResponseWriter, mockRequest)

	want := http.StatusInternalServerError
	got := mockResponseWriter.Result().StatusCode

	assert.Equal(t, want, got)
}
