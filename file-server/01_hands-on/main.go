package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("dog.gohtml"))
}

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/dog/", dog)

	http.HandleFunc("/dog.jpg", photo)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func foo(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "foo ran")
}

func dog(res http.ResponseWriter, req *http.Request) {
	err := tpl.Execute(res, "")
	if err != nil {
		log.Fatal(err)
	}
}

func photo(res http.ResponseWriter, req *http.Request) {
	http.ServeFile(res, req, "./dog.jpg")
}
