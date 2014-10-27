package coelho

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_shouldReturnDefaultTweetingPaolo(t *testing.T) {
	paolo := DefaultTwettingPaolo()
	assert.NotNil(t, paolo)
}

func Test_shouldCreateNewPaolo(t *testing.T) {
	twitterClient := &FakeTwitterClient{}
	paolo := NewTweetingPaolo(twitterClient)
	assert.NotNil(t, paolo)
}

func Test_shouldGetRandomTweet(t *testing.T) {
	twitterClient := &FakeTwitterClient{}
	paolo := NewTweetingPaolo(twitterClient)

	randomTweet := paolo.RandomTweet()

	assert.NotNil(t, randomTweet)
	assert.Equal(t, "test tweet", randomTweet)
}

type FakeTwitterClient struct{}

func (fake *FakeTwitterClient) Tweets(screenName string) ([]string, error) {
	return []string{"test tweet"}, nil
}
