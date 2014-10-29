package main

import (
	"bitbucket.org/maneo/coelho"
	"encoding/json"
	"net/http"
)

func main() {

	wisePaolo := coelho.NewHeartQuotes()
	twittingPaolo := coelho.DefaultTwettingPaolo()
	localPaolo, err := coelho.NewPaolo("quotes.json")
	if err != nil {
		println("Failed to load quotes: " + err.Error())
		return
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		paolo := coelho.Paolo{}
		quote := make(chan coelho.Quote)
		go func() { quote <- localPaolo.RandomQuote() }()
		go func() { quote <- twittingPaolo.RandomTweet() }()
		go func() { quote <- wisePaolo.RandomHeartQuote() }()

		for i := 0; i < 3; i++ {
			paolo.Quotes = append(paolo.Quotes, <-quote)
		}

		quotes, _ := json.Marshal(paolo.Quotes)
		w.Write(quotes)
	})

	http.ListenAndServe(":8080", nil)
}
