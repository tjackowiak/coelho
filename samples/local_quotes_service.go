package main

import (
	"fmt"
	"github.com/tjackowiak/coelho"
	"io"
	"net/http"
)

func main() {
	paolo, err := coelho.NewPaolo("quotes.json")
	if err != nil {
		println("Failed to load quotes")
		return
	}

	upQuoute := 0

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, paolo.RandomQuote().Sentence)
		upQuoute++
	})
	http.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, fmt.Sprintf("Service is alive, served %d quotes so far", upQuoute))
	})
	http.ListenAndServe(":8080", nil)
}
