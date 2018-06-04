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
	http.Handle("/resources/", http.StripPrefix("/resources/", http.FileServer(http.Dir("starting-files/public"))))

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleIndex(res http.ResponseWriter, req *http.Request) {
	error := tpl.Execute(res, nil)
	if error != nil {
		log.Fatal(error)
	}
}
