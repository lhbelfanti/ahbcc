package corpus_test

import (
	"context"
	"encoding/json"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/lhbelfanti/corpus-creator/cmd/api/corpus"
)

func TestMakeExportCorpus_successWithJSON(t *testing.T) {
	corpusData := []corpus.DAO{{ID: 1, TweetAuthor: "author1"}}
	mockSelectAll := corpus.MockSelectAll(corpusData, nil)
	mockExportToJSON := corpus.MockExportDataToJSON(corpus.MockJSONExportResult(), nil)
	mockExportToCSV := corpus.MockExportDataToCSV(nil, nil)

	exportCorpus := corpus.MakeExportCorpus(mockSelectAll, mockExportToJSON, mockExportToCSV)

	got, err := exportCorpus(context.Background(), corpus.JSONFormat)

	assert.Nil(t, err)
	assert.Equal(t, "application/json", got.ContentType)
	assert.Equal(t, "corpus.json", got.Filename)
}

func TestMakeExportCorpus_successWithCSV(t *testing.T) {
	corpusData := []corpus.DAO{{ID: 1, TweetAuthor: "author1"}}
	mockSelectAll := corpus.MockSelectAll(corpusData, nil)
	mockExportToJSON := corpus.MockExportDataToJSON(nil, nil)
	mockExportToCSV := corpus.MockExportDataToCSV(corpus.MockCSVExportResult(), nil)

	exportCorpus := corpus.MakeExportCorpus(mockSelectAll, mockExportToJSON, mockExportToCSV)

	got, err := exportCorpus(context.Background(), corpus.CSVFormat)

	assert.Nil(t, err)
	assert.Equal(t, "text/csv", got.ContentType)
	assert.Equal(t, "corpus.csv", got.Filename)
}

func TestMakeExportCorpus_failsWhenTheFormatRequestedIsInvalid(t *testing.T) {
	mockSelectAll := corpus.MockSelectAll(nil, nil)
	mockExportToJSON := corpus.MockExportDataToJSON(nil, nil)
	mockExportToCSV := corpus.MockExportDataToCSV(nil, nil)

	exportCorpus := corpus.MakeExportCorpus(mockSelectAll, mockExportToJSON, mockExportToCSV)

	want := corpus.InvalidExportFormat
	_, got := exportCorpus(context.Background(), "invalid-format")

	assert.Equal(t, want, got)
}

func TestMakeExportCorpus_failsWhenSelectAllThrowsError(t *testing.T) {
	mockSelectAll := corpus.MockSelectAll(nil, errors.New("failed to select all"))
	mockExportToJSON := corpus.MockExportDataToJSON(nil, nil)
	mockExportToCSV := corpus.MockExportDataToCSV(nil, nil)

	exportCorpus := corpus.MakeExportCorpus(mockSelectAll, mockExportToJSON, mockExportToCSV)

	want := corpus.FailedToExecuteSelectAll
	_, got := exportCorpus(context.Background(), corpus.JSONFormat)

	assert.Equal(t, want, got)
}

func TestMakeExportCorpus_failsWhenJSONExportThrowsError(t *testing.T) {
	corpusData := []corpus.DAO{{ID: 1, TweetAuthor: "author1"}}
	want := errors.New("JSON encoding error")
	mockSelectAll := corpus.MockSelectAll(corpusData, nil)
	mockExportToJSON := corpus.MockExportDataToJSON(nil, want)
	mockExportToCSV := corpus.MockExportDataToCSV(nil, nil)

	exportCorpus := corpus.MakeExportCorpus(mockSelectAll, mockExportToJSON, mockExportToCSV)

	_, got := exportCorpus(context.Background(), corpus.JSONFormat)

	assert.Equal(t, want, got)
}

func TestMakeExportDataToJSON_success(t *testing.T) {
	corpusData := []corpus.DAO{{ID: 1, TweetAuthor: "author1"}}

	exportDataToJSON := corpus.MakeExportDataToJSON()

	got, err := exportDataToJSON(context.Background(), corpusData)

	assert.Nil(t, err)
	var decodedData []corpus.DAO
	err = json.Unmarshal(got.Data, &decodedData)
	assert.Nil(t, err)
	assert.Equal(t, corpusData, decodedData)
	assert.Equal(t, "application/json", got.ContentType)
	assert.Equal(t, "corpus.json", got.Filename)
}

func TestMakeExportDataToJSON_emptyCorpus(t *testing.T) {
	var corpusData []corpus.DAO

	exportDataToJSON := corpus.MakeExportDataToJSON()

	got, err := exportDataToJSON(context.Background(), corpusData)

	assert.Nil(t, err)
	assert.Equal(t, "[]", string(got.Data))
	assert.Equal(t, "application/json", got.ContentType)
	assert.Equal(t, "corpus.json", got.Filename)
}

func TestMakeExportDataToCSV_success(t *testing.T) {
	corpusData := []corpus.DAO{corpus.MockDAO()}

	exportDataToCSV := corpus.MakeExportDataToCSV()

	want := corpus.MockCSVData()
	got, err := exportDataToCSV(context.Background(), corpusData)

	assert.Nil(t, err)
	assert.Equal(t, "text/csv", got.ContentType)
	assert.Equal(t, "corpus.csv", got.Filename)
	assert.Equal(t, string(got.Data), want)
}

func TestMakeExportDataToCSV_emptyCorpus(t *testing.T) {
	var corpusData []corpus.DAO

	exportDataToCSV := corpus.MakeExportDataToCSV()

	got, err := exportDataToCSV(context.Background(), corpusData)

	assert.Nil(t, err)
	assert.Equal(t, "text/csv", got.ContentType)
	assert.Equal(t, "corpus.csv", got.Filename)
	// Verify the CSV contains only a header row for empty corpus
	assert.Contains(t, string(got.Data), "ID,TweetAuthor")
	// Ensure there's no data row
	assert.NotContains(t, string(got.Data), "\n1,")
}

func TestMakeExportCorpus_CSVExportError(t *testing.T) {
	corpusData := []corpus.DAO{{ID: 1, TweetAuthor: "author1"}}
	want := errors.New("CSV encoding error")
	mockSelectAll := corpus.MockSelectAll(corpusData, nil)
	mockExportToJSON := corpus.MockExportDataToJSON(nil, nil)
	mockExportToCSV := corpus.MockExportDataToCSV(nil, want)

	exportCorpus := corpus.MakeExportCorpus(mockSelectAll, mockExportToJSON, mockExportToCSV)

	_, got := exportCorpus(context.Background(), corpus.CSVFormat)

	assert.Equal(t, want, got)
}
