package main

import "net/http"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello Go!"))
	})
	// nil uses Default mux which is basically default instance of server.
	http.ListenAndServe(":8080", nil)
}