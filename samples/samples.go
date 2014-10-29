package main

// simple sample webservice
// import (
// 	"io"
// 	"net/http"
// )

// func main() {
// 	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 		io.WriteString(w, "hello, world\n")
// 	})
// 	http.ListenAndServe(":8080", nil)
// }

// *******************
// Paolo coelho
// *******************
// import (
// 	"bitbucket.org/maneo/coelho"
// 	"fmt"
// 	"io"
// 	"net/http"
// )

// func main() {
// 	paolo, err := coelho.NewPaolo("quotes.json")
// 	if err != nil {
// 		println("Failed to load quotes")
// 		return
// 	}

// 	upQuoute := 0

// 	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 		io.WriteString(w, paolo.RandomQuote().Sentence)
// 		upQuoute++
// 	})
// 	http.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
// 		io.WriteString(w, fmt.Sprintf("Service is alive, served %d quotes so far", upQuoute))
// 	})
// 	http.ListenAndServe(":8080", nil)
// }

// *******************
// Tweeting Paolo coelho
// *******************
// import (
// 	"bitbucket.org/maneo/coelho"
// 	"fmt"
// 	"io"
// 	"net/http"
// )

// func main() {

// 	paolo := coelho.DefaultTwettingPaolo()
// 	upTweets := 0

// 	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 		io.WriteString(w, paolo.RandomTweet().Sentence)
// 		upTweets++
// 	})
// 	http.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
// 		io.WriteString(w, fmt.Sprintf("Service is alive, served %d tweets so far", upTweets))
// 	})
// 	http.ListenAndServe(":8080", nil)
// }

// *******************
// Reading iheart quotes
// *******************
// import (
// 	"bitbucket.org/maneo/coelho"
// 	"fmt"
// 	"io"
// 	"net/http"
// )

// func main() {

// 	upQuotes := 0

// 	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 		quote := coelho.NewHeartQuotes().RandomHeartQuote()
// 		io.WriteString(w, quote.Sentence)
// 		upQuotes++
// 	})
// 	http.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
// 		io.WriteString(w, fmt.Sprintf("Service is alive, served %d tweets so far", upQuotes))
// 	})
// 	http.ListenAndServe(":8080", nil)
// }

// *******************
// Final sample
// *******************
import (
	"bitbucket.org/maneo/coelho"
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
