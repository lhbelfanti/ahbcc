package corpus

import (
	"bytes"
	"context"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/lhbelfanti/corpus-creator/internal/log"
)

const (
	JSONFormat string = "json"
	CSVFormat  string = "csv"
)

type (
	// ExportCorpus exports the corpus in a given format
	ExportCorpus func(ctx context.Context, format string) (*ExportResult, error)

	// ExportDataToJSON exports the corpus in a JSON format
	ExportDataToJSON func(ctx context.Context, corpusData []DAO) (*ExportResult, error)

	// ExportDataToCSV exports the corpus in a CSV format
	ExportDataToCSV func(ctx context.Context, corpusData []DAO) (*ExportResult, error)
)

// MakeExportCorpus creates a new ExportCorpus function
func MakeExportCorpus(selectAll SelectAll, toJSON ExportDataToJSON, toCSV ExportDataToCSV) ExportCorpus {
	return func(ctx context.Context, format string) (*ExportResult, error) {
		if format != JSONFormat && format != CSVFormat {
			log.Error(ctx, fmt.Sprintf("Invalid export format: %s", format))
			return nil, InvalidExportFormat
		}

		corpusData, err := selectAll(ctx)
		if err != nil {
			log.Error(ctx, err.Error())
			return nil, FailedToExecuteSelectAll
		}

		var exportResult *ExportResult
		switch format {
		case JSONFormat:
			exportResult, err = toJSON(ctx, corpusData)
		case CSVFormat:
			exportResult, err = toCSV(ctx, corpusData)
		}

		return exportResult, err
	}
}

// MakeExportDataToJSON creates a new ExportDataToJSON function
func MakeExportDataToJSON() ExportDataToJSON {
	return func(ctx context.Context, corpusData []DAO) (*ExportResult, error) {
		var data []byte
		var err error

		if corpusData == nil {
			data = []byte("[]")
		} else {
			data, err = json.MarshalIndent(corpusData, "", "  ")
			if err != nil {
				log.Error(ctx, "Failed to encode JSON data: "+err.Error())
				return nil, err
			}
		}

		return &ExportResult{
			Data:        data,
			ContentType: "application/json",
			Filename:    "corpus.json",
		}, nil
	}
}

// MakeExportDataToCSV creates a new ExportDataToCSV function
func MakeExportDataToCSV() ExportDataToCSV {
	return func(ctx context.Context, corpusData []DAO) (*ExportResult, error) {
		var buf bytes.Buffer
		writer := csv.NewWriter(&buf)

		header := []string{
			"ID", "TweetAuthor", "TweetAvatar", "TweetText", "TweetImages", "IsTweetAReply",
			"QuoteAuthor", "QuoteAvatar", "QuoteText", "QuoteImages", "IsQuoteAReply", "Categorization",
		}

		err := writer.Write(header)
		if err != nil {
			log.Error(ctx, "Failed to write CSV header: "+err.Error())
			return nil, err
		}

		for _, entry := range corpusData {
			tweetAvatar := ""
			if entry.TweetAvatar != nil {
				tweetAvatar = *entry.TweetAvatar
			}

			tweetText := ""
			if entry.TweetText != nil {
				tweetText = *entry.TweetText
			}

			tweetImages := ""
			if len(entry.TweetImages) > 0 {
				tweetImages = strings.Join(entry.TweetImages, ",")
			}

			quoteAuthor := ""
			if entry.QuoteAuthor != nil {
				quoteAuthor = *entry.QuoteAuthor
			}

			quoteAvatar := ""
			if entry.QuoteAvatar != nil {
				quoteAvatar = *entry.QuoteAvatar
			}

			quoteText := ""
			if entry.QuoteText != nil {
				quoteText = *entry.QuoteText
			}

			quoteImages := ""
			if len(entry.QuoteImages) > 0 {
				quoteImages = strings.Join(entry.QuoteImages, ",")
			}

			isQuoteAReply := ""
			if entry.IsQuoteAReply != nil {
				isQuoteAReply = fmt.Sprintf("%v", *entry.IsQuoteAReply)
			}

			row := []string{
				strconv.Itoa(entry.ID),
				entry.TweetAuthor,
				tweetAvatar,
				tweetText,
				tweetImages,
				fmt.Sprintf("%v", entry.IsTweetAReply),
				quoteAuthor,
				quoteAvatar,
				quoteText,
				quoteImages,
				isQuoteAReply,
				entry.Categorization,
			}

			err = writer.Write(row)
			if err != nil {
				log.Error(ctx, "Failed to write CSV row: "+err.Error())
				return nil, err
			}
		}

		writer.Flush()
		err = writer.Error()
		if err != nil {
			log.Error(ctx, "Failed to flush CSV writer: "+err.Error())
			return nil, err
		}

		return &ExportResult{
			Data:        buf.Bytes(),
			ContentType: "text/csv",
			Filename:    "corpus.csv",
		}, nil
	}
}
