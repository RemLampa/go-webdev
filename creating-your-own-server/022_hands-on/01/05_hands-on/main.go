package main

import (
	"html/template"
	"log"
	"net/http"
)

type httpHandler int

func (s httpHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	err := tpl.Execute(res, "Welcome!")
	if err != nil {
		log.Fatalln(err)
	}
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	var handler httpHandler

	http.Handle("/", handler)

	http.HandleFunc("/dog", http.HandlerFunc(dogHandler))

	http.HandleFunc("/me", http.HandlerFunc(meHandler))

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func rootHandler(res http.ResponseWriter, req *http.Request) {
	err := tpl.Execute(res, "Welcome!")
	if err != nil {
		log.Fatal(err)
	}
}

func dogHandler(res http.ResponseWriter, req *http.Request) {
	err := tpl.Execute(res, "Bark! Bark! Bark")
	if err != nil {
		log.Fatal(err)
	}
}

func meHandler(res http.ResponseWriter, req *http.Request) {
	err := tpl.Execute(res, "Hi, I'm Rem!")
	if err != nil {
		log.Fatal(err)
	}
}
