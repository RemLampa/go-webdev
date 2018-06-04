package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("starting-files/templates/index.gohtml"))
}

func main() {
	http.HandleFunc("/", handleIndex)
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("starting-files/public"))))

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func handleIndex(res http.ResponseWriter, req *http.Request) {
	err := tpl.Execute(res, nil)
	if err != nil {
		log.Fatal(err)
	}
}
