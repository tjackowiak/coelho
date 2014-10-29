package coelho

import (
	"encoding/json"
	"errors"
	"math/rand"
	"os"
)

// Quote
type Quote struct {
	Source   string `json:"source"`
	Sentence string `json:"sentence"`
}

// Paolo type, gives access to random Paolo quote
type Paolo struct {
	Quotes []Quote
}

// NewPaolo is born
func NewPaolo(quotesFile string) (*Paolo, error) {
	newQuotes, err := newQuotes(quotesFile)
	if err != nil {
		return nil, err
	}
	return &Paolo{Quotes: newQuotes}, nil
}

//RandomQuote is returned
func (paolo *Paolo) RandomQuote() Quote {
	return paolo.Quotes[rand.Intn(len(paolo.Quotes))]
}

// newQuotes factory
func newQuotes(quotesFilePath string) ([]Quote, error) {
	fileContent, err := os.Open(quotesFilePath)
	if err != nil {
		return nil, err
	}
	defer fileContent.Close()

	decoder := json.NewDecoder(fileContent)
	var quotes []Quote
	if err := decoder.Decode(&quotes); err != nil {
		return nil, errors.New("Failed to decode json: " + err.Error())
	}
	return quotes, nil
}
