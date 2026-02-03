package cleaner

import (
	"ahbcc/internal/log"
	"context"
	"regexp"

	"ahbcc/cmd/api/corpus/cleaner/rules"
	"ahbcc/cmd/api/tweets"
)

// CleanTweets cleans tweets according to cleaning rules
type CleanTweets func(ctx context.Context, tweets []tweets.TweetDTO, cleaningRules []rules.DTO) error

// MakeCleanTweets creates a new CleanTweets
func MakeCleanTweets() CleanTweets {
	return func(ctx context.Context, tweets []tweets.TweetDTO, cleaningRules []rules.DTO) error {
		for _, rule := range cleaningRules {
			re, err := regexp.Compile(rule.SourceText)
			if err != nil {
				log.Error(ctx, err.Error())
				return CannotParseRegex
			}

			for _, tweet := range tweets {
				textContent := *tweet.TextContent
				switch rule.RuleType {
				case rules.RuleReplacement:
					textContent = re.ReplaceAllString(textContent, *rule.TargetText)
				case rules.RuleDelete:
					textContent = re.ReplaceAllString(textContent, "")
				case rules.RuleBadWord:
					textContent = re.ReplaceAllString(textContent, "***")
				}

				*tweet.TextContent = textContent
			}
		}

		return nil
	}
}
