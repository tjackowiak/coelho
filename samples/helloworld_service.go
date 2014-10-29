package main

import (
	"github.com/tjackowiak/coelho"
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "hello, world\n")
	})
	http.ListenAndServe(":8080", nil)
}
