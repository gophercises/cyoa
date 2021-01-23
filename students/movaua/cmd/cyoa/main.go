package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"../../pkg/handler"
	"../../pkg/model"
)

func main() {
	filename := flag.String("book", "goper.json", "a book JSON file name")
	port := flag.Int("port", 8080, "port to listen on for the HTTP requests")
	flag.Parse()

	f, err := os.Open(*filename)
	check(err)
	defer f.Close()

	story, err := model.JSONStory(f)
	check(err)
	f.Close()

	h := handler.New(story)

	fmt.Printf("starting server on :%d...\n", *port)
	err = http.ListenAndServe(fmt.Sprintf(":%d", *port), h)
	check(err)
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
