package cleaner

import (
	"context"
	"regexp"

	"ahbcc/cmd/api/corpus/cleaner/rules"
	"ahbcc/internal/log"
)

// CleanTweets cleans tweets according to cleaning rules
type CleanTweets func(ctx context.Context, tweets []TweetToClean) error

// MakeCleanTweets creates a new CleanTweets
func MakeCleanTweets(selectCleaningRulesByPriority rules.SelectAllByPriority) CleanTweets {
	return func(ctx context.Context, tweets []TweetToClean) error {
		// The priorities go from 1 to 10, being 1 the first highest priority.
		// This means the cleaning rule will be applied first.
		cleaningRulesSlice := make([]rules.DAO, 0, 10)
		for i := 1; i <= 10; i++ {
			cleaningRules, err := selectCleaningRulesByPriority(ctx, i)
			if err != nil {
				log.Error(ctx, err.Error())
				return FailedToRetrieveCleaningRulesByPriority
			}

			cleaningRulesSlice = append(cleaningRulesSlice, cleaningRules...)
		}

		for _, rule := range cleaningRulesSlice {
			re, err := regexp.Compile(rule.SourceText)
			if err != nil {
				log.Error(ctx, err.Error())
				return CannotParseRegex
			}

			for _, tweet := range tweets {
				textContent := *tweet.TweetText
				switch rule.RuleType {
				case rules.RuleReplacement:
					textContent = re.ReplaceAllString(textContent, *rule.TargetText)
				case rules.RuleDelete:
					textContent = re.ReplaceAllString(textContent, "")
				case rules.RuleBadWord:
					textContent = re.ReplaceAllString(textContent, "***")
				}

				*tweet.TweetText = textContent
			}
		}

		return nil
	}
}
