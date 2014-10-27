package coelho

import (
	"errors"
	"github.com/ChimeraCoder/anaconda"
	"net/url"
)

const (
	consumerKey    = "NTGWeHADBaAWfMWlm9G61RR0o"
	consumerSecret = "2hvDVHvNCWzKfj0x11hANV3iTTuv8ZDWaTV70XFCgEPwrghhdK"
	accessToken    = "13347922-N6OZ1bGeEJG8LuhDcXjTCLo9lYYWr0Hzx0areWMLk"
	accessSecret   = "PabfY3LWWBLJHN9VFMTeZW1q92CB4hQR1CGwihbNNo5iT"
)

//TwitterClient interface
type TwitterClient interface {
	Tweets(screenName string) ([]string, error)
}

//AnacondaTwitterClient gets recent posts from user
type AnacondaTwitterClient struct {
	api *anaconda.TwitterApi
}

//NewTwitterClient with default configuration
func NewTwitterClient() *AnacondaTwitterClient {
	twitter := &AnacondaTwitterClient{}
	anaconda.SetConsumerKey(consumerKey)
	anaconda.SetConsumerSecret(consumerSecret)
	twitter.api = anaconda.NewTwitterApi(accessToken, accessSecret)
	return twitter
}

//Tweets from given screen name
func (twitter *AnacondaTwitterClient) Tweets(screenName string) ([]string, error) {
	if len(screenName) == 0 {
		return nil, errors.New("screenName cannot be empty")
	}

	v := url.Values{}
	v.Set("screen_name", screenName)
	results, err := twitter.api.GetUserTimeline(v)
	if err != nil {
		print("Failed to load Paolos timeline" + err.Error())
		return nil, err
	}

	tweets := make([]string, len(results))
	for i, tweet := range results {
		tweets[i] = tweet.Text
	}

	return tweets, nil
}
