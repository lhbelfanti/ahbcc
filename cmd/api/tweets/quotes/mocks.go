package quotes

// MockInsertSingle mocks InsertSingle function
func MockInsertSingle(quoteID int, err error) InsertSingle {
	return func(q QuoteDTO) (int, error) {
		return quoteID, err
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
