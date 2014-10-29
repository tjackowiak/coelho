package coelho

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"errors"
)

func Test_shouldReturnDefaultTweetingPaolo(t *testing.T) {
	paolo := DefaultTwettingPaolo()
	assert.NotNil(t, paolo)
}

func Test_shouldCreateNewPaolo(t *testing.T) {
	twitterClient := &FakeTwitterClient{Msg:"test tweet"}
	paolo := NewTweetingPaolo(twitterClient)
	assert.NotNil(t, paolo)
}

func Test_shouldGetRandomTweet(t *testing.T) {
	twitterClient := &FakeTwitterClient{Msg:"test tweet"}
	paolo := NewTweetingPaolo(twitterClient)

	randomTweet := paolo.RandomTweet()

	assert.NotNil(t, randomTweet)
	assert.Equal(t, "test tweet", randomTweet.Sentence)
}

func Test_shouldBringDarknessOnNoTweet(t *testing.T) {
	twitterClient := &FakeTwitterClient{Err: errors.New("No tweet for you")}
	paolo := NewTweetingPaolo(twitterClient)

	randomTweet := paolo.RandomTweet()

	assert.Equal(t, "darkness", randomTweet.Sentence)
}

type FakeTwitterClient struct{
	Msg string
	Err error
}

func (fake *FakeTwitterClient) Tweets(screenName string) ([]string, error) {
	return []string{fake.Msg}, fake.Err
}
