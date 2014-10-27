package coelho

import (
	"html"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	endpointURL = "http://www.iheartquotes.com/api/v1/random"
)

//HeartQuotes endpoint
type HeartQuotes struct {
	QuoteFetcher func() (string, error)
}

//NewHeartQuotes is ready
func NewHeartQuotes() *HeartQuotes {
	instance := &HeartQuotes{}
	instance.QuoteFetcher = func() (string, error) {
		resp, err := http.Get(endpointURL)
		if err != nil {
			return "", err
		}

		defer resp.Body.Close()
		result, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return "", err
		}
		quote := string(result[:])
		return quote, nil
	}
	return instance
}

//RandomHeartQuote returns random quote
func (hq *HeartQuotes) RandomHeartQuote() (string, error) {
	resp, err := hq.QuoteFetcher()
	return hq.cleanQuote(resp), err
}

func (hq *HeartQuotes) cleanQuote(quote string) string {
	quote = html.UnescapeString(quote)
	quote = strings.Replace(quote, "\n", "", -1)

	if len(quote) > 140 {
		quote = quote[:139]
	}
	return quote
}
