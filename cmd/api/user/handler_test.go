package user_test

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"ahbcc/cmd/api/search/criteria"
	"ahbcc/cmd/api/user"
)

func TestInformationV1_success(t *testing.T) {
	mockInformationDTOs := criteria.MockInformationDTOs()
	mockInformation := user.MockInformation(mockInformationDTOs, nil)
	mockBody, _ := json.Marshal(mockInformation)
	mockResponseWriter := httptest.NewRecorder()
	mockRequest, _ := http.NewRequestWithContext(context.Background(), http.MethodGet, "/users/{user_id}/criteria/v1", bytes.NewReader(mockBody))
	mockRequest.SetPathValue("user_id", "1")

	handlerV1 := user.InformationV1(mockInformation)

	handlerV1(mockResponseWriter, mockRequest)

	want := http.StatusOK
	got := mockResponseWriter.Result().StatusCode

	assert.Equal(t, want, got)
}

func TestInformationV1_failsWhenTheURLParamIsEmpty(t *testing.T) {
	mockInformationDTOs := criteria.MockInformationDTOs()
	mockInformation := user.MockInformation(mockInformationDTOs, nil)
	mockBody, _ := json.Marshal(mockInformation)
	mockResponseWriter := httptest.NewRecorder()
	mockRequest, _ := http.NewRequestWithContext(context.Background(), http.MethodGet, "/users/{user_id}/criteria/v1", bytes.NewReader(mockBody))

	handlerV1 := user.InformationV1(mockInformation)

	handlerV1(mockResponseWriter, mockRequest)

	want := http.StatusBadRequest
	got := mockResponseWriter.Result().StatusCode

	assert.Equal(t, want, got)
}

func TestInformationV1_failsWhenInformationThrowsError(t *testing.T) {
	mockInformationDTOs := criteria.MockInformationDTOs()
	mockInformation := user.MockInformation(mockInformationDTOs, errors.New("failed to retrieve information"))
	mockBody, _ := json.Marshal(mockInformation)
	mockResponseWriter := httptest.NewRecorder()
	mockRequest, _ := http.NewRequestWithContext(context.Background(), http.MethodGet, "/users/{user_id}/criteria/v1", bytes.NewReader(mockBody))
	mockRequest.SetPathValue("user_id", "1")

	handlerV1 := user.InformationV1(mockInformation)

	handlerV1(mockResponseWriter, mockRequest)

	want := http.StatusInternalServerError
	got := mockResponseWriter.Result().StatusCode

	assert.Equal(t, want, got)
}
