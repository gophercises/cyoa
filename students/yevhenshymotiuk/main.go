package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"cyoa/parser"
)

type WebServer int

func (ws WebServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	check := func(err error) {
		if err != nil {
			log.Fatalln(err)
		}
	}

	err := r.ParseForm()
	check(err)
	place := r.Form.Get("place")

	if place == "" {
		place = "intro"
	}

	pages, err := parser.GetPages("gopher.json")
	check(err)

	page := pages[place]

	t, err := template.ParseFiles("page.html")
	check(err)

	t.Execute(w, page)
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
	}
}

func run() error {
	var ws WebServer
	port := 8080
	log.Println(fmt.Sprintf("Listening on :%d", port))
	return http.ListenAndServe(fmt.Sprintf(":%d", port), ws)
}
