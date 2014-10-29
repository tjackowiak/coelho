package main

import (
	"fmt"
	"github.com/tjackowiak/coelho"
	"io"
	"net/http"
)

func main() {

	upQuotes := 0

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		quote := coelho.NewHeartQuotes().RandomHeartQuote()
		io.WriteString(w, quote.Sentence)
		upQuotes++
	})
	http.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, fmt.Sprintf("Service is alive, served %d tweets so far", upQuotes))
	})
	http.ListenAndServe(":8080", nil)
}
