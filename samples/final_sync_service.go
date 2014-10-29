package main

import (
	"github.com/tjackowiak/coelho"
	"io"
	"net/http"
)

func main() {

	wisePaolo := coelho.NewHeartQuotes()
	twittingPaolo := coelho.DefaultTwettingPaolo()
	localPaolo, err := coelho.NewPaolo("quotes.json")
	if err != nil {
		println("Failed to load quotes")
		return
	}

	http.HandleFunc("/local", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, localPaolo.RandomQuote().Sentence)
	})

	http.HandleFunc("/heart", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, wisePaolo.RandomHeartQuote().Sentence)
	})

	http.HandleFunc("/twitter", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, twittingPaolo.RandomTweet().Sentence)
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		localQuote := localPaolo.RandomQuote().Sentence
		twitterQuote := twittingPaolo.RandomTweet().Sentence
		heartQuote := wisePaolo.RandomHeartQuote().Sentence
		io.WriteString(w, "First tweet\n")
		io.WriteString(w, localQuote)
		io.WriteString(w, "\nSecond tweet\n")
		io.WriteString(w, twitterQuote)
		io.WriteString(w, "\nThird tweet\n")
		io.WriteString(w, heartQuote)
	})

	http.ListenAndServe(":8080", nil)
}
