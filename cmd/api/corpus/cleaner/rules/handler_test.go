package rules_test

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"ahbcc/cmd/api/corpus/cleaner/rules"
)

func TestInsertRulesHandlerV1_success(t *testing.T) {
	mockInsert := rules.MockInsert(nil)
	mockResponseWriter := httptest.NewRecorder()
	mockRules := rules.MockRulesDTOSlice()
	mockBody, _ := json.Marshal(mockRules)
	mockRequest, _ := http.NewRequestWithContext(context.Background(), http.MethodPost, "/corpus/cleaning-rules/v1", bytes.NewReader(mockBody))

	insertRulesHandlerV1 := rules.InsertRulesHandlerV1(mockInsert)

	insertRulesHandlerV1(mockResponseWriter, mockRequest)

	want := http.StatusOK
	got := mockResponseWriter.Result().StatusCode

	assert.Equal(t, want, got)
}

func TestInsertRulesHandlerV1_failsWhenTheBodyCannotBeParsed(t *testing.T) {
	mockInsert := rules.MockInsert(nil)
	mockResponseWriter := httptest.NewRecorder()
	mockBody, _ := json.Marshal(`{"wrong": "body"}`)
	mockRequest, _ := http.NewRequestWithContext(context.Background(), http.MethodPost, "/corpus/cleaning-rules/v1", bytes.NewReader(mockBody))

	insertRulesHandlerV1 := rules.InsertRulesHandlerV1(mockInsert)

	insertRulesHandlerV1(mockResponseWriter, mockRequest)

	want := http.StatusBadRequest
	got := mockResponseWriter.Result().StatusCode

	assert.Equal(t, want, got)
}

func TestInsertRulesHandlerV1_failsWhenTheRuleTypeIsInvalid(t *testing.T) {
	mockInsert := rules.MockInsert(nil)
	mockResponseWriter := httptest.NewRecorder()
	mockRules := rules.MockRulesDTOSlice()
	mockRules[0].RuleType = "wrong"
	mockBody, _ := json.Marshal(mockRules)
	mockRequest, _ := http.NewRequestWithContext(context.Background(), http.MethodPost, "/corpus/cleaning-rules/v1", bytes.NewReader(mockBody))

	insertRulesHandlerV1 := rules.InsertRulesHandlerV1(mockInsert)

	insertRulesHandlerV1(mockResponseWriter, mockRequest)

	want := http.StatusBadRequest
	got := mockResponseWriter.Result().StatusCode

	assert.Equal(t, want, got)
}

func TestInsertRulesHandlerV1_failsWhenInsertRulesThrowsError(t *testing.T) {
	mockInsert := rules.MockInsert(errors.New("failed to insert rules"))
	mockResponseWriter := httptest.NewRecorder()
	mockRules := rules.MockRulesDTOSlice()
	mockBody, _ := json.Marshal(mockRules)
	mockRequest, _ := http.NewRequestWithContext(context.Background(), http.MethodPost, "/corpus/cleaning-rules/v1", bytes.NewReader(mockBody))

	insertRulesHandlerV1 := rules.InsertRulesHandlerV1(mockInsert)

	insertRulesHandlerV1(mockResponseWriter, mockRequest)

	want := http.StatusInternalServerError
	got := mockResponseWriter.Result().StatusCode

	assert.Equal(t, want, got)
}
