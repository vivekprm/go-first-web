package main

import (
	"fmt"
	"net/http"
)

type myHandler struct {
	greeting string
}

func (mh myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fmt.Sprintf("%v world", mh.greeting)))
}
func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello world!"))
	})
	http.Handle("/", &myHandler{greeting: "Hello"})
	http.ListenAndServe(":8080", nil)
}