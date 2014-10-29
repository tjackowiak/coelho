package main

import (
	"fmt"
	"github.com/tjackowiak/coelho"
	"io"
	"net/http"
)

func main() {

	paolo := coelho.DefaultTwettingPaolo()
	upTweets := 0

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, paolo.RandomTweet().Sentence)
		upTweets++
	})
	http.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, fmt.Sprintf("Service is alive, served %d tweets so far", upTweets))
	})
	http.ListenAndServe(":8080", nil)
}
