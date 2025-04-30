package quotes

import (
	"context"
	"time"
)

// MockInsertSingle mocks InsertSingle function
func MockInsertSingle(quoteID int, err error) InsertSingle {
	return func(ctx context.Context, q *QuoteDTO) (int, error) {
		return quoteID, err
	}
}

// MockDeleteOrphans mocks DeleteOrphans function
func MockDeleteOrphans(err error) DeleteOrphans {
	return func(ctx context.Context, ids []int) error {
		return err
	}
}

// MockQuoteDTO mocks a QuoteDTO
func MockQuoteDTO() QuoteDTO {
	textContent := "test"
	avatar := "https://testquoteavatar.com"

	return QuoteDTO{
		IsAReply:    true,
		Author:      "TestQuoteAuthor",
		Avatar:      &avatar,
		PostedAt:    "2022-11-18T15:04:05Z",
		TextContent: &textContent,
		Images:      []string{"test1", "test2"},
	}
}

// MockCustomQuoteDTO mocks a CustomQuoteDTO
func MockCustomQuoteDTO() CustomQuoteDTO {
	textContent := "test"
	avatar := "https://testquoteavatar.com"

	return CustomQuoteDTO{
		IsAReply:    true,
		Author:      "TestQuoteAuthor",
		Avatar:      &avatar,
		PostedAt:    time.Now(),
		TextContent: &textContent,
		Images:      []string{"test1", "test2"},
	}
}
