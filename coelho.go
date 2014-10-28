package coelho

import (
	"encoding/json"
	"errors"
	"math/rand"
	"os"
)

// Paolo type, gives access to random Paolo quote
type Paolo struct {
	quotes []Quote
}

//RandomQuote is returned
func (paolo *Paolo) RandomQuote() Quote {
	return paolo.quotes[rand.Intn(len(paolo.quotes))]
}

// NewPaolo is born
func NewPaolo(quotesFile string) (*Paolo, error) {
	newQuotes, err := NewQuotes(quotesFile)
	if err != nil {
		return nil, err
	}
	return &Paolo{quotes: newQuotes}, nil
}

// Quote db
type Quote struct {
	BookTitle string `json:"book_title"`
	Sentence  string `json:"sentence"`
}

// NewQuotes factory
func NewQuotes(quotesFilePath string) ([]Quote, error) {
	fileContent, err := os.Open(quotesFilePath)
	if err != nil {
		return nil, err
	}
	defer fileContent.Close()

	decoder := json.NewDecoder(fileContent)
	var quotes []Quote
	if err := decoder.Decode(&quotes); err != nil {
		return nil, errors.New("Failed to decode json")
	}
	return quotes, nil
}
