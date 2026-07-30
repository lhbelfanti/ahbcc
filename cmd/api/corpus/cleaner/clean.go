package cleaner

import (
	"context"
	"regexp"
	"strings"

	"github.com/lhbelfanti/corpus-creator/cmd/api/corpus/cleaner/rules"
	"github.com/lhbelfanti/corpus-creator/internal/log"
)

type (
	// CleanTweets cleans tweets according to cleaning rules
	CleanTweets func(ctx context.Context, tweets []TweetToClean) error

	// compiledRule is a cleaning rule whose pattern and replacement have already
	// been compiled/accent-escaped, ready to be applied to every tweet
	compiledRule struct {
		re         *regexp.Regexp
		ruleType   string
		targetText string
	}
)

// accentToASCII holds a reversible mapping between Spanish accented characters and
// two-character ASCII placeholders (an underscore plus a letter/digit, both of which
// Go's RE2 engine treats as "word" characters).
//
// Go's regexp package (RE2) always treats \b/\w as ASCII-only ([0-9A-Za-z_]), with no
// flag to make it Unicode-aware. Accented vowels (á, é, í, ó, ú, ...) are therefore NOT
// considered word characters by \b, which breaks any cleaning rule using \b next to an
// accented letter: e.g. the rule `\bm\b -> me` sees a word boundary between the "m" and
// the "á" in "más" (because RE2 thinks "á" ends the word), matches the lone "m", and
// corrupts "más" into "meás". The same happened with "día" -> "deía", "tú" -> "teú", and
// with longer rules like `\bNas\b -> Unas`, which fired inside "escénas" and produced
// "escéUnas".
//
// The fix: before applying any \b-based rule, temporarily replace every accented
// character (in both the tweet text and the rule patterns/replacements) with a stand-in
// ASCII sequence that RE2 DOES treat as part of a word. This makes \b behave correctly
// around what were originally accented letters. Once every rule has been applied, the
// placeholders are mapped back to the original accented characters.
var accentToASCII = map[string]string{
	"á": "_1", "é": "_2", "í": "_3", "ó": "_4", "ú": "_5", "ñ": "_6", "ü": "_7",
	"Á": "_8", "É": "_9", "Í": "_A", "Ó": "_B", "Ú": "_C", "Ñ": "_D", "Ü": "_E",
}

// invertStringMap swaps the keys and values of a string-to-string map
func invertStringMap(m map[string]string) map[string]string {
	inverted := make(map[string]string, len(m))
	for k, v := range m {
		inverted[v] = k
	}

	return inverted
}

// toSafeASCII replaces accented characters with reversible ASCII placeholders
func toSafeASCII(text string) string {
	for accented, ascii := range accentToASCII {
		text = strings.ReplaceAll(text, accented, ascii)
	}

	return text
}

// fromSafeASCII restores the original accented characters from their placeholders
func fromSafeASCII(text string) string {
	for ascii, accented := range invertStringMap(accentToASCII) {
		text = strings.ReplaceAll(text, ascii, accented)
	}

	return text
}

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

		// Compile every rule once, up front, against the accent-safe version of its
		// pattern/replacement.
		compiledRules := make([]compiledRule, 0, len(cleaningRulesSlice))
		for _, rule := range cleaningRulesSlice {
			re, err := regexp.Compile(toSafeASCII(rule.SourceText))
			if err != nil {
				log.Error(ctx, err.Error())
				return CannotParseRegex
			}

			target := ""
			if rule.TargetText != nil {
				target = toSafeASCII(*rule.TargetText)
			}
			compiledRules = append(compiledRules, compiledRule{re: re, ruleType: rule.RuleType, targetText: target})
		}

		for _, tweet := range tweets {
			textContent := toSafeASCII(*tweet.TweetText)

			for _, rule := range compiledRules {
				switch rule.ruleType {
				case rules.RuleReplacement:
					textContent = rule.re.ReplaceAllString(textContent, rule.targetText)
				case rules.RuleDelete:
					textContent = rule.re.ReplaceAllString(textContent, "")
				case rules.RuleBadWord:
					defaultBadWordTag := "[BAD_WORD]"
					if rule.targetText != "" {
						defaultBadWordTag = rule.targetText
					}
					textContent = rule.re.ReplaceAllString(textContent, defaultBadWordTag)
				}
			}

			*tweet.TweetText = fromSafeASCII(textContent)
		}

		return nil
	}
}
