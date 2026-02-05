package corpus

import "ahbcc/cmd/api/corpus/cleaner"

// DTO represents a corpus entry to be inserted into the 'corpus' table
type DTO struct {
	TweetAuthor    string   `json:"tweet_author"`
	TweetAvatar    *string  `json:"tweet_avatar,omitempty"`
	TweetText      *string  `json:"tweet_text,omitempty"`
	TweetImages    []string `json:"tweet_images,omitempty"`
	IsTweetAReply  bool     `json:"is_tweet_a_reply"`
	QuoteAuthor    *string  `json:"quote_author"`
	QuoteAvatar    *string  `json:"quote_avatar,omitempty"`
	QuoteText      *string  `json:"quote_text,omitempty"`
	QuoteImages    []string `json:"quote_images,omitempty"`
	IsQuoteAReply  *bool    `json:"is_quote_a_reply,omitempty"`
	Categorization string   `json:"categorization"`
}

// convertTweetsToCleanToDTOs converts a slice of cleaner.TweetToClean to a slice of DTOs
func convertTweetsToCleanToDTOs(tweetsToClean []cleaner.TweetToClean) []DTO {
	var corpusDTOs []DTO
	for _, ttc := range tweetsToClean {
		corpusDTOs = append(corpusDTOs, toDTO(ttc))
	}
	return corpusDTOs
}

// toDTO converts a cleaner.TweetToClean to a DTO
func toDTO(tweetToClean cleaner.TweetToClean) DTO {
	return DTO{
		TweetAuthor:    tweetToClean.TweetAuthor,
		TweetAvatar:    tweetToClean.TweetAvatar,
		TweetText:      tweetToClean.TweetText,
		TweetImages:    tweetToClean.TweetImages,
		IsTweetAReply:  tweetToClean.IsTweetAReply,
		QuoteAuthor:    tweetToClean.QuoteAuthor,
		QuoteAvatar:    tweetToClean.QuoteAvatar,
		QuoteText:      tweetToClean.QuoteText,
		QuoteImages:    tweetToClean.QuoteImages,
		IsQuoteAReply:  tweetToClean.IsQuoteAReply,
		Categorization: tweetToClean.Categorization,
	}
}
