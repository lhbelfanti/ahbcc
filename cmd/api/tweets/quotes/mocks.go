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

// MockSelectByID mocks SelectByID function
func MockSelectByID(quoteDAO DAO, err error) SelectByID {
	return func(ctx context.Context, id int) (DAO, error) {
		return quoteDAO, err
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

// MockTweetQuoteDAO mocks a quotes.DAO
func MockTweetQuoteDAO() DAO {
	textContent := "test"
	avatar := "https://testquoteavatar.com"

	return DAO{
		ID:          1,
		Author:      "TestQuoteAuthor",
		Avatar:      &avatar,
		PostedAt:    time.Now(),
		IsAReply:    false,
		TextContent: &textContent,
		Images:      []string{"test1", "test2"},
	}
}

// MockScanTweetQuoteDAOValues mocks the properties of DAO to be used in the Scan function
func MockScanTweetQuoteDAOValues(dao DAO) []any {
	return []any{
		dao.ID,
		dao.Author,
		dao.Avatar,
		dao.PostedAt,
		dao.IsAReply,
		dao.TextContent,
		dao.Images,
	}
}
