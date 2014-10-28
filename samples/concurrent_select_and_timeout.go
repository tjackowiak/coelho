package main

import (
	"bitbucket.org/maneo/coelho"
	"io"
	"net/http"
	"strconv"
	"time"
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

		timeout := time.After(500 * time.Millisecond)
		for i := 0; i < 3; i++ {
			select {
			case tweet := <-quote:
				io.WriteString(w, "#"+strconv.Itoa(i+1)+" tweet\n")
				io.WriteString(w, tweet+"\n")
			case <-timeout:
				io.WriteString(w, "\nWhy so slow?")
				return
			}
		}

	})

	http.ListenAndServe(":8080", nil)
}
