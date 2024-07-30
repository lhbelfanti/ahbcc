package quotes

import "context"

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

	return QuoteDTO{
		IsAReply:    true,
		TextContent: &textContent,
		Images:      []string{"test1", "test2"},
	}
}
