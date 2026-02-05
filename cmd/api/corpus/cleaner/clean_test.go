package cleaner_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"ahbcc/cmd/api/corpus/cleaner"
	"ahbcc/cmd/api/corpus/cleaner/rules"
)

func TestCleanTweet_successWithRuleReplacementWithRegex(t *testing.T) {
	textContent := "Hello 123 World"
	mockTweet := cleaner.MockTweetToClean(textContent)
	target := "NUM"
	mockCleaningRule := rules.MockRuleDAO(rules.RuleReplacement, `\d+`, &target, 1)
	mockSelectAllByPriority := rules.MockSelectAllByPriority([]rules.DAO{mockCleaningRule}, nil)

	cleanTweet := cleaner.MakeCleanTweets(mockSelectAllByPriority)

	err := cleanTweet(context.Background(), []cleaner.TweetToClean{mockTweet})

	want := "Hello NUM World"
	got := *mockTweet.TweetText

	assert.Nil(t, err)
	assert.Equal(t, want, got)
}

func TestCleanTweet_successWithRuleDeleteWithRegex(t *testing.T) {
	textContent := "Hello 123 World"
	mockTweet := cleaner.MockTweetToClean(textContent)
	mockCleaningRule := rules.MockRuleDAO(rules.RuleDelete, `\s\d+\s`, nil, 1)
	mockSelectAllByPriority := rules.MockSelectAllByPriority([]rules.DAO{mockCleaningRule}, nil)

	cleanTweet := cleaner.MakeCleanTweets(mockSelectAllByPriority)

	err := cleanTweet(context.Background(), []cleaner.TweetToClean{mockTweet})

	want := "HelloWorld"
	got := *mockTweet.TweetText

	assert.Nil(t, err)
	assert.Equal(t, want, got)
}

func TestCleanTweet_successWithRuleBadWordWithRegex(t *testing.T) {
	textContent := "Hello badword123 World"
	mockTweet := cleaner.MockTweetToClean(textContent)
	mockCleaningRule := rules.MockRuleDAO(rules.RuleBadWord, `badword\d+`, nil, 1)
	mockSelectAllByPriority := rules.MockSelectAllByPriority([]rules.DAO{mockCleaningRule}, nil)

	cleanTweet := cleaner.MakeCleanTweets(mockSelectAllByPriority)

	err := cleanTweet(context.Background(), []cleaner.TweetToClean{mockTweet})

	want := "Hello *** World"
	got := *mockTweet.TweetText

	assert.Nil(t, err)
	assert.Equal(t, want, got)
}

func TestCleanTweet_successWithLiteralStringReplacement(t *testing.T) {
	textContent := "Hello World"
	mockTweet := cleaner.MockTweetToClean(textContent)
	target := "Universe"
	mockCleaningRule := rules.MockRuleDAO(rules.RuleReplacement, "World", &target, 1)
	mockSelectAllByPriority := rules.MockSelectAllByPriority([]rules.DAO{mockCleaningRule}, nil)

	cleanTweet := cleaner.MakeCleanTweets(mockSelectAllByPriority)

	err := cleanTweet(context.Background(), []cleaner.TweetToClean{mockTweet})

	want := "Hello Universe"
	got := *mockTweet.TweetText

	assert.Nil(t, err)
	assert.Equal(t, want, got)
}

func TestCleanTweets_successWithMultipleTweets(t *testing.T) {
	tweetContent1 := "Hello 123"
	mockTweet1 := cleaner.MockTweetToClean(tweetContent1)
	tweetContent2 := "World 456"
	mockTweet2 := cleaner.MockTweetToClean(tweetContent2)
	target := "NUM"
	mockCleaningRule := rules.MockRuleDAO(rules.RuleReplacement, `\d+`, &target, 1)
	mockSelectAllByPriority := rules.MockSelectAllByPriority([]rules.DAO{mockCleaningRule}, nil)

	cleanTweets := cleaner.MakeCleanTweets(mockSelectAllByPriority)

	err := cleanTweets(context.Background(), []cleaner.TweetToClean{mockTweet1, mockTweet2})

	assert.Nil(t, err)
	assert.Equal(t, "Hello NUM", *mockTweet1.TweetText)
	assert.Equal(t, "World NUM", *mockTweet2.TweetText)
}

func TestCleanTweet_failsWhenSelectCleaningRulesByPriorityThrowsError(t *testing.T) {
	textContent := "Hello 123 World"
	mockTweet := cleaner.MockTweetToClean(textContent)
	target := "NUM"
	mockCleaningRule := rules.MockRuleDAO(rules.RuleReplacement, `\d+`, &target, 1)
	mockSelectAllByPriority := rules.MockSelectAllByPriority([]rules.DAO{mockCleaningRule}, errors.New("failed to select cleaning rules by priority"))

	cleanTweet := cleaner.MakeCleanTweets(mockSelectAllByPriority)

	want := cleaner.FailedToRetrieveCleaningRulesByPriority
	got := cleanTweet(context.Background(), []cleaner.TweetToClean{mockTweet})

	assert.Equal(t, want, got)
}

func TestCleanTweet_failsWhenRegexCompileThrowsError(t *testing.T) {
	textContent := "Hello World"
	mockTweet := cleaner.MockTweetToClean(textContent)
	mockCleaningRule := rules.MockRuleDAO(rules.RuleReplacement, "[", nil, 1)
	mockSelectAllByPriority := rules.MockSelectAllByPriority([]rules.DAO{mockCleaningRule}, nil)

	cleanTweet := cleaner.MakeCleanTweets(mockSelectAllByPriority)

	want := cleaner.CannotParseRegex
	got := cleanTweet(context.Background(), []cleaner.TweetToClean{mockTweet})

	assert.Equal(t, want, got)
}
