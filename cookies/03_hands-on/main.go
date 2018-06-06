package main

import (
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", indexHandler)
	http.Handle("/favicon.ico", http.NotFoundHandler())

	http.ListenAndServe("localhost:8080", nil)
}

func indexHandler(res http.ResponseWriter, req *http.Request) {
	timesVisited, err := req.Cookie("times-visited")
	if err != nil || timesVisited.Value == "" {
		http.SetCookie(res, &http.Cookie{
			Name:  "times-visited",
			Value: "1",
		})
	} else {
		timesVisited, err := strconv.Atoi(timesVisited.Value)
		if err != nil {
			log.Fatal(err)
		}

		http.SetCookie(res, &http.Cookie{
			Name:  "times-visited",
			Value: strconv.Itoa(timesVisited + 1),
		})
	}
}
