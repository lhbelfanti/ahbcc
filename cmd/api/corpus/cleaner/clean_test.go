package cleaner_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"ahbcc/cmd/api/corpus/cleaner"
	"ahbcc/cmd/api/corpus/cleaner/rules"
	"ahbcc/cmd/api/tweets"
)

func TestCleanTweet_successWithRuleReplacementWithRegex(t *testing.T) {
	textContent := "Hello 123 World"
	tweet := tweets.MockTweetDTO()
	tweet.TextContent = &textContent
	target := "NUM"
	cleaningRules := rules.MockRulesDTOs(rules.MockRuleDTO(rules.RuleReplacement, `\d+`, &target, 1))
	cleanTweet := cleaner.MakeCleanTweets()

	err := cleanTweet(context.Background(), []tweets.TweetDTO{tweet}, cleaningRules)

	want := "Hello NUM World"
	got := *tweet.TextContent

	assert.Nil(t, err)
	assert.Equal(t, want, got)
}

func TestCleanTweet_successWithRuleDeleteWithRegex(t *testing.T) {
	cleanTweet := cleaner.MakeCleanTweets()
	textContent := "Hello 123 World"
	tweet := tweets.MockTweetDTO()
	tweet.TextContent = &textContent
	cleaningRules := rules.MockRulesDTOs(rules.MockRuleDTO(rules.RuleDelete, `\s\d+\s`, nil, 1))

	err := cleanTweet(context.Background(), []tweets.TweetDTO{tweet}, cleaningRules)

	want := "HelloWorld"
	got := *tweet.TextContent

	assert.Nil(t, err)
	assert.Equal(t, want, got)
}

func TestCleanTweet_successWithRuleBadWordWithRegex(t *testing.T) {
	cleanTweet := cleaner.MakeCleanTweets()
	textContent := "Hello badword123 World"
	tweet := tweets.MockTweetDTO()
	tweet.TextContent = &textContent
	cleaningRules := rules.MockRulesDTOs(rules.MockRuleDTO(rules.RuleBadWord, `badword\d+`, nil, 1))

	err := cleanTweet(context.Background(), []tweets.TweetDTO{tweet}, cleaningRules)

	want := "Hello *** World"
	got := *tweet.TextContent

	assert.Nil(t, err)
	assert.Equal(t, want, got)
}

func TestCleanTweet_successWithLiteralStringReplacement(t *testing.T) {
	cleanTweet := cleaner.MakeCleanTweets()
	textContent := "Hello World"
	tweet := tweets.MockTweetDTO()
	tweet.TextContent = &textContent
	target := "Universe"
	cleaningRules := rules.MockRulesDTOs(rules.MockRuleDTO(rules.RuleReplacement, "World", &target, 1))

	err := cleanTweet(context.Background(), []tweets.TweetDTO{tweet}, cleaningRules)

	want := "Hello Universe"
	got := *tweet.TextContent

	assert.Nil(t, err)
	assert.Equal(t, want, got)
}

func TestCleanTweet_failsWhenRegexCompileThrowsError(t *testing.T) {
	cleanTweet := cleaner.MakeCleanTweets()
	textContent := "Hello World"
	tweet := tweets.MockTweetDTO()
	tweet.TextContent = &textContent
	cleaningRules := rules.MockRulesDTOs(rules.MockRuleDTO(rules.RuleReplacement, "[", nil, 1))

	want := cleaner.CannotParseRegex
	got := cleanTweet(context.Background(), []tweets.TweetDTO{tweet}, cleaningRules)

	assert.Equal(t, want, got)
}

func TestCleanTweets_multipleTweets(t *testing.T) {
	cleanTweets := cleaner.MakeCleanTweets()
	tweetContent1 := "Hello 123"
	tweet1 := tweets.MockTweetDTO()
	tweet1.TextContent = &tweetContent1
	tweetContent2 := "World 456"
	tweet2 := tweets.MockTweetDTO()
	tweet2.TextContent = &tweetContent2
	target := "NUM"
	cleaningRules := rules.MockRulesDTOs(rules.MockRuleDTO(rules.RuleReplacement, `\d+`, &target, 1))

	err := cleanTweets(context.Background(), []tweets.TweetDTO{tweet1, tweet2}, cleaningRules)

	assert.Nil(t, err)
	assert.Equal(t, "Hello NUM", *tweet1.TextContent)
	assert.Equal(t, "World NUM", *tweet2.TextContent)
}
