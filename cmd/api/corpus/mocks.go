package corpus

import "context"

// MockInsert mocks Insert function
func MockInsert(err error) Insert {
	return func(ctx context.Context, entry DTO) (int, error) {
		return 1, err
	}
}

// MockSelectAll mocks SelectAll function
func MockSelectAll(entries []DAO, err error) SelectAll {
	return func(ctx context.Context) ([]DAO, error) {
		return entries, err
	}
}

// MockDeleteAll mocks DeleteAll function
func MockDeleteAll(err error) DeleteAll {
	return func(ctx context.Context) error {
		return err
	}
}

// MockCreate mocks Create function
func MockCreate(err error) Create {
	return func(ctx context.Context, perfectlyBalanced bool, applyCleaningRules bool) error {
		return err
	}
}

// MockExportCorpus mocks ExportCorpus function
func MockExportCorpus(result *ExportResult, err error) ExportCorpus {
	return func(ctx context.Context, format string) (*ExportResult, error) {
		return result, err
	}
}

// MockExportDataToJSON mocks ExportDataToJSON function
func MockExportDataToJSON(result *ExportResult, err error) ExportDataToJSON {
	return func(ctx context.Context, corpusData []DAO) (*ExportResult, error) {
		return result, err
	}
}

// MockExportDataToCSV mocks ExportDataToCSV function
func MockExportDataToCSV(result *ExportResult, err error) ExportDataToCSV {
	return func(ctx context.Context, corpusData []DAO) (*ExportResult, error) {
		return result, err
	}
}

// MockJSONExportResult creates a mock JSON export result
func MockJSONExportResult() *ExportResult {
	return &ExportResult{
		Data:        []byte(MockJSONData()),
		ContentType: "application/json",
		Filename:    "corpus.json",
	}
}

// MockCSVExportResult creates a mock CSV export result
func MockCSVExportResult() *ExportResult {
	return &ExportResult{
		Data:        []byte(MockCSVData()),
		ContentType: "text/csv",
		Filename:    "corpus.csv",
	}
}

// MockDTO mocks a corpus.DTO
func MockDTO() DTO {
	tweetAvatar := "test_avatar"
	tweetText := "test_text"
	quoteAuthor := "quote_author"
	quoteAvatar := "quote_avatar"
	quoteText := "quote_text"
	isQuoteAReply := true

	return DTO{
		TweetAuthor:    "test_author",
		TweetAvatar:    &tweetAvatar,
		TweetText:      &tweetText,
		TweetImages:    []string{"image1.jpg", "image2.jpg"},
		IsTweetAReply:  false,
		QuoteAuthor:    &quoteAuthor,
		QuoteAvatar:    &quoteAvatar,
		QuoteText:      &quoteText,
		QuoteImages:    []string{"quote_image1.jpg"},
		IsQuoteAReply:  &isQuoteAReply,
		Categorization: "POSITIVE",
	}
}

// MockDAO mocks a corpus.DAO
func MockDAO() DAO {
	tweetAvatar := "test_avatar"
	tweetText := "test_text"
	quoteAuthor := "quote_author"
	quoteAvatar := "quote_avatar"
	quoteText := "quote_text"
	isQuoteAReply := true

	return DAO{
		ID:             1,
		TweetAuthor:    "test_author",
		TweetAvatar:    &tweetAvatar,
		TweetText:      &tweetText,
		TweetImages:    []string{"image1.jpg"},
		IsTweetAReply:  false,
		QuoteAuthor:    &quoteAuthor,
		QuoteAvatar:    &quoteAvatar,
		QuoteText:      &quoteText,
		QuoteImages:    []string{"quote_image1.jpg"},
		IsQuoteAReply:  &isQuoteAReply,
		Categorization: "POSITIVE",
	}
}

// MockCSVData mocks the string result of a CSV file
func MockCSVData() string {
	return "ID,TweetAuthor,TweetAvatar,TweetText,TweetImages,IsTweetAReply,QuoteAuthor,QuoteAvatar,QuoteText,QuoteImages,IsQuoteAReply,Categorization\n" +
		"1,test_author,test_avatar,test_text,image1.jpg,false,quote_author,quote_avatar,quote_text,quote_image1.jpg,true,POSITIVE\n"
}

// MockJSONData mocks the string result of a JSON file
func MockJSONData() string {
	return `[
	  {
		"ID": 1,
		"TweetAuthor": "test_author",
		"TweetAvatar": "test_avatar",
		"TweetText": "test_text",
		"TweetImages": ["image1.jpg"],
		"IsTweetAReply": false,
		"QuoteAuthor": "quote_author",
		"QuoteAvatar": "quote_avatar",
		"QuoteText": "quote_text",
		"QuoteImages": ["quote_image1.jpg"],
		"IsQuoteAReply": true,
		"Categorization": "POSITIVE"
	  }
	]`
}
