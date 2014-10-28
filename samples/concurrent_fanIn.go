package main

import (
	"bitbucket.org/maneo/coelho"
	"io"
	"net/http"
	"strconv"
)

func main() {

	localPaolo, err := coelho.NewPaolo("quotes.json")
	if err != nil {
		println("Failed to load quotes")
		return
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		quote := make(chan string)
		go func() { quote <- "[Local]" + localPaolo.RandomQuote().Sentence }()
		go func() { quote <- "[Twitter]" + coelho.DefaultTwettingPaolo().RandomTweet() }()
		go func() { q, _ := coelho.NewHeartQuotes().RandomHeartQuote(); quote <- "[HeartQuote]" + q }()

		for i := 0; i < 3; i++ {
			io.WriteString(w, "#"+strconv.Itoa(i+1)+" tweet\n")
			io.WriteString(w, <-quote+"\n")
		}

	})

	http.ListenAndServe(":8080", nil)
}
