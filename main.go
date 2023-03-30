package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	// Naive fileserver implementation
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		f, err := os.Open("public" + r.URL.Path)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
		}
		defer f.Close()

		var contentType string
		switch {
		case strings.HasSuffix(r.URL.Path, ".css"):
			contentType = "text/css"
		case strings.HasSuffix(r.URL.Path, ".html"):
			contentType = "text/html"
		default:
			contentType = "text/plain"
		}
		w.Header().Add("Content-Type", contentType)
		io.Copy(w, f)
	})
	// nil uses Default mux which is basically default instance of server.
	http.ListenAndServe(":8080", nil)
}