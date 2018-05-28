package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		io.WriteString(res, "Welcome!")
	})

	http.HandleFunc("/dog", func(res http.ResponseWriter, req *http.Request) {
		io.WriteString(res, "Bark! Bark! Bark!")
	})

	http.HandleFunc("/me", func(res http.ResponseWriter, req *http.Request) {
		io.WriteString(res, "Hi, I'm Rem!")
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
