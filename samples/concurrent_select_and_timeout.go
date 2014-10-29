package main

import (
	"bitbucket.org/maneo/coelho"
	"encoding/json"
	"net/http"
	"time"
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

		timeout := time.After(500 * time.Millisecond)
	Quotes:
		for i := 0; i < 3; i++ {
			select {
			case q := <-quote:
				paolo.Quotes = append(paolo.Quotes, q)
			case <-timeout:
				q := coelho.Quote{Source: "timeout", Sentence: "why so slow?"}
				paolo.Quotes = append(paolo.Quotes, q)
				break Quotes
			}
		}

		quotes, _ := json.Marshal(paolo.Quotes)
		w.Write(quotes)
	})

	http.ListenAndServe(":8080", nil)
}
