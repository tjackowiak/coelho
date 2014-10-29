package coelho

import (
	"math/rand"
)

const (
	paoloScreenName = "paulocoelho"
)

//TweetingPaolo Coelho
type TweetingPaolo struct {
	api TwitterClient
}

//DefaultTwettingPaolo return paolo
func DefaultTwettingPaolo() *TweetingPaolo {
	return &TweetingPaolo{api: NewTwitterClient()}
}

//NewTweetingPaolo factory
func NewTweetingPaolo(twitter TwitterClient) *TweetingPaolo {
	return &TweetingPaolo{api: twitter}
}

//RandomTweet is returned
func (paolo *TweetingPaolo) RandomTweet() Quote {
	quote := Quote{Source: "twitter"}
	listOfTweets, err := paolo.api.Tweets(paoloScreenName)
	if err != nil {
		quote.Sentence = "darkness"
		return quote
	}

	quote.Sentence = listOfTweets[rand.Intn(len(listOfTweets))]
	return quote
}
